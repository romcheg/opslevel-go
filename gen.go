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
	"github.com/opslevel/opslevel-go/v2024"
)

const (
	// connectionFile  string = "pkg/gen/connection.go"
	enumFile        string = "enum.go"
	inputObjectFile string = "input.go"
	// interfaceFile   string = "pkg/gen/interface.go"
	// mutationFile    string = "pkg/gen/mutation.go"
	// objectFile      string = "pkg/gen/object.go"
	// payloadFile     string = "pkg/gen/payload.go"
	// queryFile       string = "pkg/gen/query.go"
	// scalarFile      string = "pkg/gen/scalar.go"
	// unionFile       string = "pkg/gen/union.go"
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
	"filteredcount":                  "int",
	"memberships":                    "TeamMembership",
	"node":                           "any",
	"nodes":                          "[]any",
	"notupdatedrepositories":         "RepositoryOperationErrorPayload",
	"promotedchecks":                 "Check",
	"relationship":                   "RelationshipType",
	"teamsbeingnotified":             "CampaignSendReminderOutcomeTeams",
	"teamsbeingnotifiedcount":        "int",
	"teamsmissingcontactmethod":      "int",
	"teamsmissingcontactmethodcount": "int",
	"type":                           "any",
	"totalcount":                     "int",
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
		// case connectionFile:
		// 	subSchema = objectSchema
		case enumFile:
			subSchema = enumSchema
		case inputObjectFile:
			subSchema = inputObjectSchema
		// case interfaceFile:
		// 	subSchema = interfaceSchema
		// case mutationFile:
		// 	subSchema = objectSchema
		// case objectFile:
		// 	subSchema = objectSchema
		// case payloadFile:
		// 	subSchema = objectSchema
		// case queryFile:
		// 	subSchema = objectSchema
		// case scalarFile:
		// 	subSchema = scalarSchema
		// case unionFile:
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

const (
	convertedTypeTmpl = `
{{- define "converted_type" -}}
  {{ .Name | title | convertPayloadType }}
{{- end }}`
	descriptionTmpl = `
{{- define "description" -}}
 {{.Description | clean | endSentence}}
{{- end }}`
	graphqlStructTagTmpl = `
{{- define "graphql_struct_tag" -}}` + "`" + `graphql:"
  {{- .Name | lowerFirst }}"` + "`" + `
{{- end }}`
	graphqlStructTagWithArgsTmpl = `
{{- define "graphql_struct_tag_with_args" -}}` + "`" + `graphql:"
  {{- .Name}}( {{- range $index, $element := .Args }}
    {{- if gt $index 0 }}, {{ end -}}
    {{- .Name}}: ${{.Name}}
  {{- end}})"` + "`" + `
{{- end }}`
	nameToSingularTmpl = `
{{- define "name_to_singular" -}}
  {{- .Name | title | makeSingular }}
{{- end }}`
	typeCommentDescriptionTmpl = `
{{- define "type_comment_description" -}}
  // {{.Name | title}} {{ template "description" . }}
{{- end }}`
	fieldCommentDescriptionTmpl = `
{{- define "field_comment_description" -}}
  // {{ .Description | clean | fullSentence }}
{{- end }}`
)

// Filename -> Template.
var templates = map[string]*template.Template{
	// 	connectionFile: t(header + `
	// {{range .Types | sortByName}}
	//   {{if and (eq .Kind "OBJECT") (not (internal .Name)) }}
	//     {{ if hasSuffix "Connection" .Name }}
	//       {{- template "object" . }}
	//     {{end}}
	//   {{- end}}
	// {{- end}}

	// {{- define "object" -}}
	// {{ if hasSuffix "Connection" .Name }}
	// {{ template "type_comment_description" . }}
	// type {{.Name}} struct {
	//   Nodes []{{- if eq .Name "AncestorGroupsConnection"}}Group
	//           {{- else}}{{.Name | trimSuffix "Connection" | trimSuffix "V2" }} ` + "`" + `graphql:"nodes"` + "`" + `
	//           {{- end }}
	//   Edges []{{.Name | trimSuffix "Connection" }}Edge ` + "`" + `graphql:"edges"` + "`" + `
	// {{ range .Fields }} {{ if and (ne "edges" .Name) (ne "nodes" .Name) }}
	//     {{- .Name | title}} {{ template "converted_type" . }} {{ template "graphql_struct_tag" . }} {{ template "field_comment_description" . }}
	//   {{- end }}
	// {{- end }}
	// }
	// {{- end }}{{- end -}}
	//   `),
	enumFile: t(header + `
{{range .Types | sortByName}}{{if and (eq .Kind "ENUM") (not (internal .Name))}}
{{template "enum" .}}
{{end}}{{end}}


{{- define "enum" -}}
{{ template "type_comment_description" . }}
type {{.Name}} string

const ({{range .EnumValues}}
	{{$.Name}}{{.Name | enumIdentifier}} {{$.Name}} = {{.Name | quote}} {{ template "field_comment_description" . }}{{end}}
)
// All {{$.Name}} as []string
var All{{$.Name}} = []string {
	{{range .EnumValues}}string({{$.Name}}{{.Name | enumIdentifier}}),
	{{end}}
}
{{- end -}}
`),
	inputObjectFile: t(header + `
import "github.com/relvacode/iso8601"

{{range .Types | sortByName}}{{if and (eq .Kind "INPUT_OBJECT") (not (internal .Name))}}
{{ if and (not (hasPrefix "Campaign" .Name)) (not (hasPrefix "Group" .Name)) -}}
{{template "input_object" .}}
{{end}}{{end}}{{end}}

{{- define "input_object" -}}
{{ template "type_comment_description" . }}
type {{.Name}} struct { {{range .InputFields }}
  {{.Name | title}} {{if ne .Type.Kind "NON_NULL"}}*{{end -}}
    {{- if isListType .Name }}[]{{ end -}}
    {{- if and (hasSuffix "Id" .Name) (not (eq .Name "externalId")) }}ID
    {{- else if hasSuffix "Access" .Name }}IdentifierInput
    {{- else if eq .Name "predicates" }}FilterPredicateInput
    {{- else if eq .Name "tags" }}TagInput
    {{- else if eq .Name "members" }}TeamMembershipUserInput
    {{- else if eq .Name "contacts" }}ContactInput
    {{- else if eq .Type.Name "UserRole" }}UserRole
    {{- else if .Type.Name }}{{ template "converted_type" .Type }}
    {{- else }}{{ .Type.OfType.OfTypeName | convertPayloadType  }}{{ end -}} ` + "`" +
		`json:"{{.Name | lowerFirst }}{{if ne .Type.Kind "NON_NULL"}},omitempty{{end}}"` +
		` yaml:"{{.Name | lowerFirst }}{{if ne .Type.Kind "NON_NULL"}},omitempty{{end}}"` + `
  {{-  if and (not (hasSuffix "Input" .Type.Name)) (not (hasSuffix "Input" .Type.OfType.OfTypeName)) }} example:"
   {{- if isListType .Name }}[{{ end -}}
    {{- if or (eq .Type.Name "Boolean") (eq .Type.OfType.OfTypeName "Boolean") }}false
      {{- else if or (eq .Type.Name "Int") (eq .Type.OfType.OfTypeName "Int") }}3
      {{- else if eq .Type.Name "JSON" }}{\"name\":\"my-big-query\",\"engine\":\"BigQuery\",\"endpoint\":\"https://google.com\",\"replica\":false}
      {{- else if or (hasSuffix "Time" .Type.Name) (hasSuffix "Time" .Type.OfType.OfTypeName) }}2024-01-05T01:00:00.000Z
      {{- else if or (eq "FrequencyTimeScale" .Type.Name) (eq "FrequencyTimeScale" .Type.OfType.OfTypeName) }}week
      {{- else if or (eq "ContactType" .Type.Name) (eq "ContactType" .Type.OfType.OfTypeName) }}slack
      {{- else if or (eq "AlertSourceTypeEnum" .Type.Name) (hasSuffix "AlertSourceTypeEnum" .Type.OfType.OfTypeName) }}pagerduty
      {{- else if or (eq "AliasOwnerTypeEnum" .Type.Name) (hasSuffix "AliasOwnerTypeEnum" .Type.OfType.OfTypeName) }}scorecard
      {{- else if or (eq "BasicTypeEnum" .Type.Name) (hasSuffix "BasicTypeEnum" .Type.OfType.OfTypeName) }}does_not_equal
      {{- else if or (eq "ConnectiveEnum" .Type.Name) (hasSuffix "ConnectiveEnum" .Type.OfType.OfTypeName) }}or
      {{- else if or (eq "CustomActionsEntityTypeEnum" .Type.Name) (hasSuffix "CustomActionsEntityTypeEnum" .Type.OfType.OfTypeName) }}GLOBAL
      {{- else if or (eq "CustomActionsHttpMethodEnum" .Type.Name) (hasSuffix "CustomActionsHttpMethodEnum" .Type.OfType.OfTypeName) }}GET
      {{- else if or (eq "CustomActionsTriggerDefinitionAccessControlEnum" .Type.Name) (hasSuffix "CustomActionsTriggerDefinitionAccessControlEnum" .Type.OfType.OfTypeName) }}service_owners
      {{- else if or (eq "HasDocumentationTypeEnum" .Type.Name) (hasSuffix "HasDocumentationTypeEnum" .Type.OfType.OfTypeName) }}api
      {{- else if eq .Name "documentSubtype" }}openapi
      {{- else if or (eq "RelationshipTypeEnum" .Type.Name) (hasSuffix "RelationshipTypeEnum" .Type.OfType.OfTypeName) }}depends_on
      {{- else if or (eq "PredicateKeyEnum" .Type.Name) (hasSuffix "PredicateKeyEnum" .Type.OfType.OfTypeName) }}filter_id
      {{- else if or (eq "PredicateTypeEnum" .Type.Name) (hasSuffix "PredicateTypeEnum" .Type.OfType.OfTypeName) }}satisfies_jq_expression
      {{- else if or (eq "ServicePropertyTypeEnum" .Type.Name) (hasSuffix "ServicePropertyTypeEnum" .Type.OfType.OfTypeName) }}language
      {{- else if or (eq "UsersFilterEnum" .Type.Name) (hasSuffix "UsersFilterEnum" .Type.OfType.OfTypeName) }}last_sign_in_at
      {{- else if or (eq "UserRole" .Type.Name) (hasSuffix "UserRole" .Type.OfType.OfTypeName) }}admin
      {{- else if or (hasSuffix "Enum" .Type.Name) (hasSuffix "Enum" .Type.OfType.OfTypeName) }}NEW_ENUM_SET_DEFAULT
      {{- else if or (hasSuffix "ToolCategory" .Type.Name) (hasSuffix "ToolCategory" .Type.OfType.OfTypeName) }}api_documentation
      {{- else if or (eq "type" .Name) (hasSuffix "Type" .Name) }}example_type
      {{- else if eq "address" .Name }}support@company.com
      {{- else if or (eq "id" .Name) (hasSuffix "Id" .Name) }}Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk
      {{- else if or (eq "definition" .Name) (hasSuffix "Definition" .Name) }}example_definition
      {{- else if hasSuffix "Template" .Name }}{\"token\": \"XXX\", \"ref\":\"main\", \"action\": \"rollback\"}
      {{- else if or (eq "name" .Name) (hasSuffix "Name" .Name) }}example_name
      {{- else if or (eq "language" .Name) (hasSuffix "Language" .Name) }}example_language
      {{- else if or (eq "alias" .Name) (hasSuffix "Alias" .Name) }}example_alias
      {{- else if or (eq "description" .Name) (hasSuffix "Description" .Name) }}example_description
      {{- else if or (eq "key" .Name) (hasSuffix "Key" .Name) }}XXX_example_key_XXX
      {{- else if or (eq "email" .Name) (hasSuffix "Email" .Name) }}first.last@domain.com
      {{- else if or (eq "data" .Name) (hasSuffix "Data" .Name) }}example_data
      {{- else if or (eq "note" .Name) (hasSuffix "Note" .Name) }}example_note
      {{- else if or (eq "role" .Name) (hasSuffix "Role" .Name) }}example_role
      {{- else if or (eq "notes" .Name) (hasSuffix "Notes" .Name) }}example_notes
      {{- else if or (eq "value" .Name) (hasSuffix "Value" .Name) }}example_value
      {{- else if or (eq "product" .Name) (hasSuffix "Product" .Name) }}example_product
      {{- else if or (eq "framework" .Name) (hasSuffix "Framework" .Name) }}example_framework
      {{- else if or (eq "url" .Name) (hasSuffix "Url" .Name) }}john.doe@example.com
      {{- else if eq "baseDirectory" .Name }}/home/opslevel.yaml
      {{- else if eq "externalUrl" .Name }}https://google.com
      {{- else if eq "responsibilities" .Name }}example description of responsibilities
      {{- else if eq "environment" .Name }}environment that tool belongs to
      {{- else if eq "arg" .Name }}example_arg
      {{- else if hasSuffix "Extensions" .Name }}'go', 'py', 'rb'
      {{- else if hasSuffix "Paths" .Name }}'/usr/local/bin', '/home/opslevel'
      {{- else if hasSuffix "Ids" .Name }}'Z2lkOi8vc2VydmljZS8xMjM0NTY3ODk', 'Z2lkOi8vc2VydmljZS85ODc2NTQzMjE'
      {{- else if hasSuffix "TagKeys" .Name }}'tag_key1', 'tag_key2'
      {{- else if hasSuffix "Selector" .Name }}example_selector
      {{- else if hasSuffix "Condition" .Name }}example_condition
      {{- else if hasSuffix "Message" .Name }}example_message
      {{- else if hasSuffix "Method" .Name }}example_method
      {{- else if hasSuffix "Identifier" .Name}}example_identifier
    {{- end}}
   {{- if isListType .Name }}]{{ end -}}"{{- end}}` +
		"`" + `{{ template "field_comment_description" . }} {{if eq .Type.Kind "NON_NULL"}}(Required.){{else}}(Optional.){{end}}
  {{- end}}
}
{{- end -}}
`),
	// 	interfaceFile: t(header + `
	// {{range .Types | sortByName}}{{if and (eq .Kind "INTERFACE") (not (internal .Name))}}
	// {{template "interface_object" .}}
	// {{end}}{{end}}

	// {{- define "interface_object" -}}
	// {{ template "type_comment_description" . }}
	// type {{.Name}} interface { {{range .Fields }}
	//   {{ template "field_comment_description" . }}
	//   {{.Name | title}}() {{.Name | title}}
	// {{end}}
	// }
	// {{- end -}}
	// 	`),
	// 	payloadFile: t(header + `
	// {{range .Types | sortByName}}{{if and (eq .Kind "OBJECT") (not (internal .Name)) }}
	// {{template "payload_object" .}}
	// {{- end}}{{- end}}

	// {{- define "payload_object" -}}
	// {{ if hasSuffix "Payload" .Name }}
	// {{ template "type_comment_description" . }}
	// type {{.Name}} struct {
	// {{ range .Fields }}
	//   {{.Name | title}} {{ if isListType .Name }}[]{{- end -}}
	//     {{ template "converted_type" . }} {{ template "field_comment_description" . }}
	// {{- end }}
	// }
	// {{- end }}{{ end -}}
	// `),
	// NOTE: "account" == objectSchema.Types[0]
	// NOTE: "mutation" == objectSchema.Types[134]
	// NOTE: may have to use interfaceSchema to derive details for objects
	// 	queryFile: t(header + `
	// {{range .Types | sortByName}}
	//   {{if and (eq .Kind "OBJECT") (not (internal .Name)) }}
	//     {{- if eq .Name "Account" }}
	//       {{ template "account_queries" . }}
	//     {{- end}}
	//   {{- end}}
	// {{- end}}

	// {{ define "account_queries" -}}
	//     {{- range .Fields }}
	// {{ template "type_comment_description" . }}
	// func (client *Client) {{ if isListType .Name }}List{{ .Name | title }}(input any) ({{ template "name_to_singular" . }}Connection, error) {
	//     {{- else }}Get{{ .Name | title }}(input any) ({{.Name | title}}, error) {
	//     {{end -}}
	//     var q struct {
	//       Account struct {
	//         {{ .Name | title }} {{ if isListType .Name }}{{ template "name_to_singular" . }}Connection
	//                             {{- else }}Get{{ template "name_to_singular" . }}{{end -}}` + "`" + `graphql:"{{.Name}}(input: $input)"` + "`" + `
	//       }
	//     }
	//     v := PayloadVariables{ {{ range .Args }}
	//       "{{.Name}}": input, {{ end}}
	//     }
	//     err := client.Query(&q, v, WithName("{{ template "name_to_singular" . }}{{ if isListType .Name }}List{{else}}Get{{end}}"))
	//     return &q.Account.{{ .Name | title }}, HandleErrors(err, nil)
	// }
	// {{- end}}{{- end}}

	// {{- define "object" -}}
	// {{ if not (hasSuffix "Payload" .Name) }}
	// {{ template "type_comment_description" . }}
	// type {{.Name}} struct {
	//     {{ range .Fields }}
	//   {{.Name | title}} string  {{ template "field_comment_description" . }}
	//     {{- end }}
	// }
	// {{- end }}{{- end -}}
	// `),
	// 	mutationFile: t(header + `
	// {{range .Types | sortByName}}
	//   {{if and (eq .Kind "OBJECT") (not (internal .Name)) }}
	//     {{- if eq .Name "Mutation" }}
	//       {{- template "mutation" .}}
	//     {{end}}
	//   {{- end}}
	// {{- end}}

	// {{ define "mutation" -}}
	// {{- range .Fields }}
	// // {{ .Name | title | renameMutation }} {{ template "description" . }}
	// func (client *Client) {{ .Name | title | renameMutation }}(
	//   {{- range $index, $element := .Args }}{{- if gt $index 0 }}, {{ end -}}
	//     {{- if eq "IdentifierInput" .Type.OfType.OfTypeName }}identifier string
	//     {{- else }}{{- .Name }} {{ with .Type.OfType.OfTypeName }}{{.}}{{else}}any{{end}}
	//     {{- end }}
	//   {{- end -}} ) (*{{.Name | title | renameMutationReturnType}}, error) {
	//     var m struct {
	//       Payload struct {
	//         {{ .Name | title | renameMutationReturnType}} {{ .Name | title | renameMutationReturnType}}
	//         Errors []OpsLevelErrors
	//       }{{ template "graphql_struct_tag_with_args" . }}
	//     }
	//     v := PayloadVariables{ {{ range .Args }}
	//       "{{.Name}}": {{- if eq "IdentifierInput" .Type.OfType.OfTypeName }}*NewIdentifier(identifier),
	//                    {{- else}}{{.Name}},{{ end }}
	//                            {{- end}}
	//     }
	//     err := client.Mutate(&m, v, WithName("{{ .Name | title }}"))
	//     return &m.Account.{{ .Name | title | renameMutationReturnType}}, HandleErrors(err, m.Payload.Errors)
	// }
	// {{- end}}
	// {{- end}}
	// `),
	// 	objectFile: t(header + `
	// {{range .Types | sortByName}}
	//   {{if and (eq .Kind "OBJECT") (not (internal .Name)) }}
	//     {{- if eq .Name "Account" }}
	//       {{- template "account_struct" . }}
	//     {{- else}}{{template "object" .}}{{end}}
	//   {{- end}}
	// {{- end}}

	// {{ define "account_struct" -}}
	// {{ template "type_comment_description" . }}
	// type {{.Name}} struct { {{range .Fields }}
	//   {{.Name | title}} *{{ if isListType .Name }}[]{{ end }}{{ template "converted_type" . }}  {{ template "field_comment_description" . }}
	//  {{- end }}
	// }
	// {{- end }}

	// {{- define "object" -}}
	// {{ if and (not (hasSuffix "Payload" .Name)) (not (hasSuffix "Connection" .Name)) }}
	// {{ template "type_comment_description" . }}
	// type {{.Name}} struct {
	//   {{ range .Fields -}}
	//     {{ if not (len .Args) }}{{.Name | title}} {{ template "converted_type" . }} {{ template "graphql_struct_tag" . }} {{ template "field_comment_description" . }}
	//     {{- end}}
	//   {{ end -}}
	// }
	// {{- end }}{{- end -}}
	// 	`),
	// 	scalarFile: t(header + `
	// import (
	// 	"encoding/base64"
	// 	"strconv"
	// 	"strings"
	// )

	// {{range .Types | sortByName}}{{if and (eq .Kind "SCALAR") (not (internal .Name))}}
	// {{template "scalar" .}}
	// {{end}}{{end}}

	// {{- define "scalar" -}}
	// {{ template "type_comment_description" . }}
	// type {{.Name}}
	// {{- if eq .Name "Boolean" }} bool
	// {{- else if eq .Name "Float" }} float64
	// {{- else if eq .Name "ID" }} string
	// {{- else if eq .Name "ISO8601DateTime" }} string
	// {{- else if eq .Name "Int" }} int
	// {{- else if eq .Name "JSON" }} map[string]any
	// {{- else if eq .Name "JSONSchema" }} map[string]any
	// {{- else if eq .Name "String" }} string
	// {{- end -}}{{end}}

	// func NewID(id ...string) *ID {
	// 	var output ID
	// 	if len(id) == 1 {
	// 		output = ID(id[0])
	// 	}
	// 	return &output
	// }

	// func (s ID) GetGraphQLType() string { return "ID" }

	// func (s *ID) MarshalJSON() ([]byte, error) {
	// 	if *s == "" {
	// 		return []byte("null"), nil
	// 	}
	// 	return []byte(strconv.Quote(string(*s))), nil
	// }

	// type Identifier struct {
	// 	Id      ID       ` + "`" + `graphql:"id"` + "`" + `
	// 	Aliases []string ` + "`" + `graphql:"aliases"` + "`" + `
	// }

	// func NewIdentifier(value string) *IdentifierInput {
	// 	if IsID(value) {
	// 		return &IdentifierInput{
	// 			Id: NewID(value),
	// 		}
	// 	}
	// 	return &IdentifierInput{
	// 		Alias: NewString(value),
	// 	}
	// }

	// func NewIdentifierArray(values []string) []IdentifierInput {
	// 	output := []IdentifierInput{}
	// 	for _, value := range values {
	// 		output = append(output, *NewIdentifier(value))
	// 	}
	// 	return output
	// }

	// func IsID(value string) bool {
	// 	decoded, err := base64.RawURLEncoding.DecodeString(value)
	// 	if err != nil {
	// 		return false
	// 	}
	// 	return strings.HasPrefix(string(decoded), "gid://")
	// }

	// func NewString(value string) *string {
	// 	return &value
	// }`),
	// 	unionFile: t(header + `
	// {{range .Types | sortByName}}{{if and (eq .Kind "UNION") (not (internal .Name))}}
	// {{template "union_object" .}}
	// {{end}}{{end}}

	// {{- define "union_object" -}}
	// // Union{{.Name}} {{ template "description" . }}
	// type Union{{.Name}} interface { {{range .PossibleTypes }}
	//
	//	    {{.Name}}Fragment() {{.Name}}Fragment{{end}}
	//	}
	//
	// {{- end -}}
	//
	//	`),
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
	genTemplate.Parse(convertedTypeTmpl)
	genTemplate.Parse(descriptionTmpl)
	genTemplate.Parse(fieldCommentDescriptionTmpl)
	genTemplate.Parse(graphqlStructTagTmpl)
	genTemplate.Parse(graphqlStructTagWithArgsTmpl)
	genTemplate.Parse(nameToSingularTmpl)
	genTemplate.Parse(typeCommentDescriptionTmpl)
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
	switch s {
	case "Boolean":
		return "bool"
	case "Int":
		return "int"
	case "String":
		return "string"
	case "ISO8601DateTime":
		return "iso8601.Time"
	case "":
		return "string"
	}
	value := strings.ToLower(s)
	if strings.HasSuffix(value, "id") {
		return "ID"
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

// TODO fix up later
func renameMutationReturnType(s string) string {
	create := "Create"
	delete := "Delete"
	update := "Update"
	if strings.HasSuffix(s, create) {
		s = strings.TrimSuffix(s, create)
	} else if strings.HasSuffix(s, delete) {
		s = strings.TrimSuffix(s, delete)
	} else if strings.HasSuffix(s, update) {
		s = strings.TrimSuffix(s, update)
	}
	return s
}

// TODO fix up later
func renameMutation(s string) string {
	create := "Create"
	delete := "Delete"
	update := "Update"
	if strings.HasSuffix(s, create) {
		s = strings.TrimSuffix(s, create)
		s = fmt.Sprintf("%s%s", create, s)
	} else if strings.HasSuffix(s, delete) {
		s = strings.TrimSuffix(s, delete)
		s = fmt.Sprintf("%s%s", delete, s)
	} else if strings.HasSuffix(s, update) {
		s = strings.TrimSuffix(s, update)
		s = fmt.Sprintf("%s%s", update, s)
	}
	return s
}

func isPlural(s string) bool {
	value := strings.ToLower(s)
	// Examples: "alias", "address", "status", "levels", "responsibilities"
	if value == "notes" || value == "days" || value == "headers" ||
		strings.HasSuffix(value, "ies") ||
		strings.HasSuffix(value, "ias") ||
		strings.HasSuffix(value, "ls") ||
		(!strings.HasSuffix(value, "access") && strings.HasSuffix(value, "ss")) ||
		strings.HasSuffix(value, "us") {
		return false
	}
	if strings.HasSuffix(value, "s") {
		return true
	}
	return false
}

var templFuncMap = template.FuncMap{
	"internal":                 func(s string) bool { return strings.HasPrefix(s, "__") },
	"quote":                    strconv.Quote,
	"join":                     strings.Join,
	"isListType":               isPlural,
	"renameMutation":           renameMutation,
	"renameMutationReturnType": renameMutationReturnType,
	"convertPayloadType":       convertPayloadType,
	"makeSingular":             makeSingular,
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
