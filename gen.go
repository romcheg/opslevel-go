//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/hasura/go-graphql-client/ident"
	"github.com/opslevel/opslevel-go/v2023"
)

type GraphQLSchema struct {
	Types []GraphQLTypes `graphql:"types" json:"types"`
}

type GraphQLTypes struct {
	Name          string                `graphql:"name" json:"name"`
	Kind          string                `graphql:"kind" json:"kind"`
	Description   string                `graphql:"description" json:"description"`
	PossibleTypes []GraphQLPossibleType `graphql:"possibleTypes"`
	// Fields ?
	// InputFields ?
	EnumValues []GraphQLEnumValues `graphql:"enumValues" json:"enumValues"`
}

type GraphQLEnumValues struct {
	Name        string `graphql:"name" json:"name"`
	Description string `graphql:"description" json:"description"`
}

type GraphQLPossibleType struct {
	Name   string
	Kind   string
	OfType GraphQLOfType
}

type GraphQLOfType struct {
	Name string
	Kind string
}

func GetSchema(client *opslevel.Client) (*GraphQLSchema, error) {
	var q struct {
		Schema GraphQLSchema `graphql:"__schema"`
	}
	if err := client.Query(&q, nil); err != nil {
		return nil, err
	}
	return &q.Schema, nil
}

func main() {
	flag.Parse()

	err := run()
	if err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	token, ok := os.LookupEnv("OPSLEVEL_API_TOKEN")
	if !ok {
		return fmt.Errorf("OPSLEVEL_API_TOKEN environment variable not set")
	}
	client := opslevel.NewClient(token, opslevel.SetAPIVisibility("public"))
	schema, err := GetSchema(client)
	if err != nil {
		return err
	}

	for filename, t := range templates {
		var buf bytes.Buffer
		err := t.Execute(&buf, schema)
		if err != nil {
			return err
		}
		out, err := format.Source(buf.Bytes())
		if err != nil {
			log.Println(err)
			out = []byte("// gofmt error: " + err.Error() + "\n\n" + buf.String())
		}
		fmt.Println("writing", filename)
		err = ioutil.WriteFile(filename, out, 0o644)
		if err != nil {
			return err
		}
	}

	return nil
}

// Filename -> Template.
var templates = map[string]*template.Template{
	"enum.go": t(`// Code generated by gen.go; DO NOT EDIT.

package opslevel
{{range .Types | sortByName}}{{if and (eq .Kind "ENUM") (not (internal .Name))}}
{{template "enum" .}}
{{end}}{{end}}


{{- define "enum" -}}
// {{.Name}} {{.Description | clean | endSentence}}
type {{.Name}} string

const ({{range .EnumValues}}
	{{$.Name}}{{.Name | enumIdentifier}} {{$.Name}} = {{.Name | quote}} // {{.Description | clean | fullSentence}}{{end}}
)
// All {{$.Name}} as []string
var All{{$.Name}} = []string {
	{{range .EnumValues}}string({{$.Name}}{{.Name | enumIdentifier}}),
	{{end}}
}
{{- end -}}
`),
	/*
	   	"input.go": t(`// Code generated by gen.go; DO NOT EDIT.

	   	package opslevel

	   	type Input interface{}

	   	{{range .Types | sortByName}}{{if eq .Kind "INPUT_OBJECT"}}
	   	{{template "inputObject" .}}
	   	{{end}}{{end}}

	   	{{- define "inputObject" -}}
	   	// {{.Name}} {{.Description | clean | endSentence}}
	   	type {{.Name}} struct {}

	   	{{- end -}}
	   `),
	*/
	// TODO: fix this to generate all Input structs
	// 	"input.go": t(`// Code generated by gen.go; DO NOT EDIT.

	// package opslevel

	// // Input represents one of the Input structs:
	// //
	// // {{join (inputObjects .data.__schema.types) ", "}}.
	// type Input interface{}
	// {{range .data.__schema.types | sortByName}}{{if eq .kind "INPUT_OBJECT"}}
	// {{template "inputObject" .}}
	// {{end}}{{end}}

	// {{- define "inputObject" -}}
	// // {{.name}} {{.description | clean | endSentence}}
	// type {{.name}} struct {{"{"}}{{range .inputFields}}{{if eq .type.kind "NON_NULL"}}
	// 	// {{.description | clean | fullSentence}} (Required.)
	// 	{{.name | identifier}} {{.type | type}} ` + "`" + `json:"{{.name}}"` + "`" + `{{end}}{{end}}
	// {{range .inputFields}}{{if ne .type.kind "NON_NULL"}}
	// 	// {{.description | clean | fullSentence}} (Optional.)
	// 	{{.name | identifier}} {{.type | type}} ` + "`" + `json:"{{.name}},omitempty"` + "`" + `{{end}}{{end}}
	// }
	// {{- end -}}
	// `),
}

func t(text string) *template.Template {
	// typeString returns a string representation of GraphQL type t.
	var typeString func(t map[string]interface{}) string
	typeString = func(t map[string]interface{}) string {
		switch t["kind"] {
		case "NON_NULL":
			s := typeString(t["ofType"].(map[string]interface{}))
			if !strings.HasPrefix(s, "*") {
				panic(fmt.Errorf("nullable type %q doesn't begin with '*'", s))
			}
			return s[1:] // Strip star from nullable type to make it non-null.
		case "LIST":
			return "*[]" + typeString(t["ofType"].(map[string]interface{}))
		default:
			return "*" + t["name"].(string)
		}
	}

	return template.Must(template.New("").Funcs(template.FuncMap{
		"internal": func(s string) bool { return strings.HasPrefix(s, "__") },
		"quote":    strconv.Quote,
		"join":     strings.Join,
		"sortByName": func(types []GraphQLTypes) []GraphQLTypes {
			sort.Slice(types, func(i, j int) bool {
				ni := types[i].Name
				nj := types[j].Name
				return ni < nj
			})
			return types
		},
		"inputObjects": func(types []interface{}) []string {
			var names []string
			for _, t := range types {
				t := t.(map[string]interface{})
				if t["kind"].(string) != "INPUT_OBJECT" {
					continue
				}
				names = append(names, t["name"].(string))
			}
			sort.Strings(names)
			return names
		},
		"identifier":     func(name string) string { return ident.ParseLowerCamelCase(name).ToMixedCaps() },
		"enumIdentifier": func(name string) string { return ident.ParseScreamingSnakeCase(name).ToMixedCaps() },
		"type":           typeString,
		"clean":          func(s string) string { return strings.Join(strings.Fields(s), " ") },
		"endSentence": func(s string) string {
			s = strings.ToLower(s[0:1]) + s[1:]
			switch {
			default:
				s = "represents " + s
			case strings.HasPrefix(s, "autogenerated "):
				s = "is an " + s
			case strings.HasPrefix(s, "specifies "):
				// Do nothing.
			}
			if !strings.HasSuffix(s, ".") {
				s += "."
			}
			return s
		},
		"fullSentence": func(s string) string {
			if !strings.HasSuffix(s, ".") {
				s += "."
			}
			return s
		},
	}).Parse(text))
}
