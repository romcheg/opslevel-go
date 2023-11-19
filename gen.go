//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"github.com/Masterminds/sprig/v3"
	"github.com/hasura/go-graphql-client/ident"
	"github.com/opslevel/opslevel-go/v2023"
)

const (
	enumFile        string = "gen/enum.go"
	inputObjectFile string = "gen/input.go"
	interfaceFile   string = "gen/interface.go"
	objectFile      string = "gen/object.go"
	payloadFile     string = "gen/payload.go"
	queryFile       string = "gen/query.go"
	scalarFile      string = "gen/scalar_gen.go"
	unionFile       string = "gen/union.go"
)

var stringTypeSuffixes = []string{
	"actionmessage",
	"alias",
	"aliases",
	"apidocsdefaultpath",
	"createdat",
	"cursor",
	"email",
	"externaluuid",
	"htmlurl",
	"id",
	"kind",
	"message",
	"name",
	"processedat",
	"role",
	"queryparams",
	"updatedat",
	"userdeletepayload",
	"yaml",
}

var knownTypeMappings = map[string]string{
	"data":                           "JSON",
	"deletedmembers":                 "User",
	"edges":                          "any",
	"filteredcount":                  "Int",
	"memberships":                    "TeamMembership",
	"node":                           "any",
	"nodes":                          "any",
	"notupdatedrepositories":         "RepositoryOperationErrorPayload",
	"promotedchecks":                 "Check",
	"relationship":                   "RelationshipType",
	"teamsbeingnotified":             "CampaignSendReminderOutcomeTeams",
	"teamsbeingnotifiedcount":        "Int",
	"teamsmissingcontactmethod":      "Int",
	"teamsmissingcontactmethodcount": "Int",
	"type":                           "any",
	"totalcount":                     "Int",
	"triggerdefinition":              "CustomActionsTriggerDefinition",
	"updatedrepositories":            "Repository",
	"webhookaction":                  "CustomActionsWebhookAction",
}

const header = `// Code generated by gen.go; DO NOT EDIT.

package opslevel`

type GraphQLSchema struct {
	Types []GraphQLTypes `graphql:"types" json:"types"`
}

type IntrospectiveType struct {
	Name   string `graphql:"name" json:"name"`
	Kind   string `graphql:"kind" json:"kind"`
	OfType struct {
		OfTypeName string `graphql:"name" json:"name"`
	} `graphql:"ofType" json:"ofType"`
}

type GraphQLInputValue struct {
	Name         string            `graphql:"name" json:"name"`
	DefaultValue string            `graphql:"defaultValue" json:"defaultValue"`
	Description  string            `graphql:"description" json:"description"`
	Type         IntrospectiveType `graphql:"type" json:"type"`
}

type GraphQLField struct {
	Args         []GraphQLInputValue `graphql:"args" json:"args"`
	Description  string              `graphql:"description" json:"description"`
	IsDeprecated bool                `graphql:"isDeprecated" json:"isDeprecated"`
	Name         string              `graphql:"name" json:"name"`
}

type GraphQLTypes struct {
	Name          string                `graphql:"name" json:"name"`
	Kind          string                `graphql:"kind" json:"kind"`
	Description   string                `graphql:"description" json:"description"`
	PossibleTypes []GraphQLPossibleType `graphql:"possibleTypes"`
	EnumValues    []GraphQLEnumValues   `graphql:"enumValues" json:"enumValues"`
	Fields        []GraphQLField        `graphql:"fields" json:"fields"`
	InputFields   []GraphQLInputValue   `graphql:"inputFields" json:"inputFields"`
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

func getRootSchema() (*GraphQLSchema, error) {
	token, ok := os.LookupEnv("OPSLEVEL_API_TOKEN")
	if !ok {
		return nil, fmt.Errorf("OPSLEVEL_API_TOKEN environment variable not set")
	}
	client := opslevel.NewGQLClient(opslevel.SetAPIToken(token), opslevel.SetAPIVisibility("public"))
	schema, err := GetSchema(client)
	if err != nil {
		return nil, err
	}
	return schema, nil
}

func run() error {
	schema, err := getRootSchema()
	if err != nil {
		return err
	}

	enumSchema := GraphQLSchema{}
	inputObjectSchema := GraphQLSchema{}
	interfaceSchema := GraphQLSchema{}
	objectSchema := GraphQLSchema{}
	scalarSchema := GraphQLSchema{}
	unionSchema := GraphQLSchema{}
	for _, t := range schema.Types {
		switch t.Kind {
		case "ENUM":
			enumSchema.Types = append(enumSchema.Types, t)
		case "INPUT_OBJECT":
			inputObjectSchema.Types = append(inputObjectSchema.Types, t)
		case "INTERFACE":
			interfaceSchema.Types = append(interfaceSchema.Types, t)
		case "OBJECT":
			objectSchema.Types = append(objectSchema.Types, t)
		case "SCALAR":
			scalarSchema.Types = append(scalarSchema.Types, t)
		case "UNION":
			unionSchema.Types = append(unionSchema.Types, t)
		default:
			panic("Unknown GraphQL type: " + t.Kind)
		}
	}

	var buf bytes.Buffer
	var subSchema GraphQLSchema
	for filename, t := range templates {
		switch filename {
		case enumFile:
			subSchema = enumSchema
		case inputObjectFile:
			subSchema = inputObjectSchema
		case interfaceFile:
			break
		// 	subSchema = interfaceSchema
		case objectFile:
			subSchema = objectSchema
		case payloadFile:
			subSchema = objectSchema
		case queryFile:
			subSchema = objectSchema
		case scalarFile:
			subSchema = scalarSchema
		case unionFile:
			break
		// 	subSchema = unionSchema
		default:
			panic("Unknown file: " + filename)
		}
		err := t.Execute(&buf, subSchema)
		if err != nil {
			return err
		}
		out, err := format.Source(buf.Bytes())
		if err != nil {
			log.Println(err)
			out = []byte("// gofmt error: " + err.Error() + "\n\n" + buf.String())
		}
		buf.Reset()
		fmt.Println("writing", filename)
		err = os.WriteFile(filename, out, 0o644)
		if err != nil {
			return err
		}
	}

	return nil
}

// Filename -> Template.
var templates = map[string]*template.Template{
	enumFile: t(header + `
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
	inputObjectFile: t(header + `
{{range .Types | sortByName}}{{if and (eq .Kind "INPUT_OBJECT") (not (internal .Name))}}
{{template "input_object" .}}
{{end}}{{end}}

{{- define "input_object" -}}
// {{.Name}} {{.Description | clean | endSentence}}
type {{.Name}} struct { {{range .InputFields }}
  {{.Name | title}} {{if ne .Type.Kind "NON_NULL"}}*{{end -}}
    {{- if isListType .Name }}[]{{ end -}}
    {{- with .Type.Name }}{{. | convertPayloadType }}
    {{- else }}{{ .Type.OfType.OfTypeName | convertPayloadType  }}{{ end -}}
    ` + "`" + `json:"{{.Name | lowerFirst }}{{if ne .Type.Kind "NON_NULL"}},omitempty{{end}}"` +
		"`" + `// {{.Description | clean | fullSentence}} {{if eq .Type.Kind "NON_NULL"}}(Required.){{else}}(Optional.){{end}}
  {{- end}}
}
{{- end -}}
`),
	interfaceFile: t(`// Code generated by gen.go; DO NOT EDIT.

package opslevel
{{range .Types | sortByName}}{{if and (eq .Kind "INTERFACE") (not (internal .Name))}}
{{template "interface_object" .}}
{{end}}{{end}}

{{- define "interface_object" -}}
// {{.Name}} {{.Description | clean | endSentence}}
type {{.Name}} interface { {{range .Fields }}
  // {{.Description | clean | fullSentence}}
  {{.Name | title}}() {{.Name | title}}
{{end}}
}
{{- end -}}
	`),
	payloadFile: t(header + `
{{range .Types | sortByName}}{{if and (eq .Kind "OBJECT") (not (internal .Name)) }}
{{template "payload_object" .}}
{{- end}}{{- end}}

{{- define "payload_object" -}}
{{ if hasSuffix "Payload" .Name }}
// {{.Name}} {{.Description | clean | endSentence}}
type {{.Name}} struct {
{{ range .Fields }}
  {{.Name | title}} {{ if isListType .Name }}[]{{- end -}}
    {{ .Name | title | convertPayloadType }} // {{.Description | clean | fullSentence}}
{{- end }}
}
{{- end }}{{ end -}}
`),
	// NOTE: "account" == objectSchema.Types[0]
	// NOTE: "mutation" == objectSchema.Types[134]
	// NOTE: may have to use interfaceSchema to derive details for objects
	queryFile: t(header + `
{{range .Types | sortByName}}
  {{if and (eq .Kind "OBJECT") (not (internal .Name)) }}
    {{- if eq .Name "Account" }}
      {{ template "account_queries" . }}
    {{- end}}
  {{- end}}
{{- end}}

{{ define "account_queries" -}}
    {{- range .Fields }}
// {{ .Name | title }} {{.Description | clean | endSentence}}
func (client *Client) {{ if isListType .Name }}List{{ .Name | title }}(input any) ({{.Name | title | makeSingular}}Connection, error) {
    {{- else }}Get{{ .Name | title }}(input any) ({{.Name | title}}, error) {
    {{end -}}
    var q struct {
      Account struct {
        {{ .Name | title }} {{ if isListType .Name }}{{ .Name | title | makeSingular }}Connection
                            {{- else }}Get{{ .Name | title | makeSingular }}{{end -}}` + "`" + `graphql:"{{.Name}}(input: $input)"` + "`" + `
      }
    }
    v := PayloadVariables{ {{ range .Args }}
      "{{.Name}}": input,
      {{- end}}
    }
    err := client.Query(&q, v, WithName("{{ .Name | title | makeSingular }}{{ if isListType .Name }}List{{else}}Get{{end}}"))
    return &q.Account.{{ .Name | title }}, HandleErrors(err, nil)
}
{{- end}}{{- end}}

{{- define "object" -}}
{{ if not (hasSuffix "Payload" .Name) }}
// {{.Name}} {{.Description | clean | endSentence}}
type {{.Name}} struct {
    {{ range .Fields }}
  {{.Name | title}} string  // {{.Description | clean | fullSentence}}
    {{- end }}
}
{{- end }}{{- end -}}
`),
	objectFile: t(header + `
{{range .Types | sortByName}}
  {{if and (eq .Kind "OBJECT") (not (internal .Name)) }}
    {{- if eq .Name "Mutation" }}
      {{- template "mutation" .}}
    {{- else if eq .Name "Account" }}
      {{- template "account_struct" . }}
    {{- else}}{{template "object" .}}{{end}}
  {{- end}}
{{- end}}

{{ define "account_struct" -}}
// {{.Name}} {{.Description | clean | endSentence}}
type {{.Name}} struct { {{range .Fields }}
  {{.Name | title}} *{{ if isListType .Name }}[]{{ end }}{{ .Name | title | convertPayloadType }}  // {{.Description | clean | fullSentence}}
 {{- end }}
}
{{- end }}

{{ define "mutation" -}}
{{range .Fields }}
// {{.Name | title}} {{.Description | clean | endSentence}}
func {{.Name | title}}(
  {{- range $index, $element := .Args }}{{- if gt $index 0 }}, {{ end -}}
    {{- .Name }} {{ with .Type.OfType.OfTypeName }}{{.}}{{else}}any{{end}}
{{- end -}} ) {
  // TODO
} {{- end}}
{{- end -}}

{{- define "object" -}}
{{ if not (hasSuffix "Payload" .Name) }}
// {{.Name}} {{.Description | clean | endSentence}}
type {{.Name}} struct {
    {{ range .Fields -}}
    {{.Name | title}} {{ .Name | title | convertPayloadType }} ` + "`" + `graphql:"{{.Name | lowerFirst }}"` + "`" + ` // {{.Description | clean | fullSentence}}
    {{ end -}}
}
{{- end }}{{- end -}}
	`),
	scalarFile: t(header + `
{{range .Types | sortByName}}{{if and (eq .Kind "SCALAR") (not (internal .Name))}}
{{template "scalar" .}}
{{end}}{{end}}

{{- define "scalar" -}}
// {{.Name}} {{.Description | clean | endSentence}}
type {{.Name}}
{{- if eq .Name "Boolean" }} bool
{{- else if eq .Name "Float" }} float64
{{- else if eq .Name "ID" }} string
{{- else if eq .Name "ISO8601DateTime" }} string
{{- else if eq .Name "Int" }} int
{{- else if eq .Name "JSON" }} map[string]any
{{- else if eq .Name "JSONSchema" }} map[string]any
{{- else if eq .Name "String" }} string
{{- end -}}{{end}}
	`),
	unionFile: t(header + `
{{range .Types | sortByName}}{{if and (eq .Kind "UNION") (not (internal .Name))}}
{{template "union_object" .}}
{{end}}{{end}}

{{- define "union_object" -}}
// Union{{.Name}} {{.Description | clean | endSentence}}
type Union{{.Name}} interface { {{range .PossibleTypes }}
    {{.Name}}Fragment() {{.Name}}Fragment{{end}}
}
{{- end -}}
	`),
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

	genTemplate := template.New("")
	genTemplate.Funcs(templFuncMap)
	genTemplate.Funcs(sprig.TxtFuncMap())
	genTemplate.Funcs(template.FuncMap{"type": typeString})
	return template.Must(genTemplate.Parse(text))
}

func makeSingular(s string) string {
	value := strings.ToLower(s)
	if strings.HasSuffix(value, "ies") {
		return strings.ReplaceAll(s, "ies", "y")
	}
	if isPlural(s) {
		return strings.TrimSuffix(s, "s")
	}
	return s
}

func convertPayloadType(s string) string {
	if s == "" || s == "String" || s == "ISO8601DateTime" {
		return "string"
	}
	value := strings.ToLower(s)
	if strings.HasSuffix(value, "id") {
		return "ID"
	}
	if value == "boolean" {
		return "bool"
	}
	for k, v := range knownTypeMappings {
		if value == k {
			return v
		}
	}
	for _, knownStringTypeSuffix := range stringTypeSuffixes {
		if strings.HasSuffix(value, knownStringTypeSuffix) {
			return "string"
		}
	}
	return makeSingular(s)
}

func isPlural(s string) bool {
	value := strings.ToLower(s)
	// Examples: "alias", "address"
	if value == "notes" ||
		strings.HasSuffix(value, "ias") ||
		strings.HasSuffix(value, "ss") {
		return false
	}
	if strings.HasSuffix(value, "es") ||
		strings.HasSuffix(value, "as") ||
		strings.HasSuffix(value, "s") {
		return true
	}
	return false
}

var templFuncMap = template.FuncMap{
	"internal":           func(s string) bool { return strings.HasPrefix(s, "__") },
	"quote":              strconv.Quote,
	"join":               strings.Join,
	"isListType":         isPlural,
	"convertPayloadType": convertPayloadType,
	"makeSingular":       makeSingular,
	"lowerFirst": func(value string) string {
		for i, v := range value {
			return string(unicode.ToLower(v)) + value[i+1:]
		}
		return value
	},
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
	"clean":          func(s string) string { return strings.Join(strings.Fields(s), " ") },
	"endSentence": func(s string) string {
		if len(s) == 0 {
			// Do nothing.
			return ""
		}

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
}
