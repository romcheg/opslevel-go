package opslevel_test

import (
	"testing"

	ol "github.com/opslevel/opslevel-go/v2023"
	"github.com/rocktavious/autopilot/v2023"
)

var checkCreateInput = ol.CheckCreateInput{
	Name:     "Hello World",
	Enabled:  true,
	Category: "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1",
	Level:    "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3",
	Notes:    "Hello World Check",
}

var checkUpdateNotes = "Hello World Check"

var checkUpdateInput = ol.CheckUpdateInput{
	Id:       "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4",
	Name:     "Hello World",
	Enabled:  ol.Bool(true),
	Category: "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1",
	Level:    "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3",
	Notes:    &checkUpdateNotes,
}

// Temporary solution until repo wide testing is standardized
type TmpCheckTestCase struct {
	fixture  TestRequest
	endpoint string
	body     func(c *ol.Client) (*ol.Check, error)
}

func getCheckTestCases() map[string]TmpCheckTestCase {
	testcases := map[string]TmpCheckTestCase{
		"CreateAlertSourceUsage": {
			fixture: NewTestRequest(
				`"mutation CheckAlertSourceUsageCreate($input:CheckAlertSourceUsageCreateInput!){checkAlertSourceUsageCreate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "alertSourceNamePredicate": {"type":"equals", "value":"Requests"}, "alertSourceType":"datadog" } }`,
				`{ "data": { "checkAlertSourceUsageCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_alert_source_usage",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.CreateCheckAlertSourceUsage(ol.CheckAlertSourceUsageCreateInput{
					CheckCreateInput: checkCreateInput,
					AlertSourceType:  ol.AlertSourceTypeEnumDatadog,
					AlertSourceNamePredicate: &ol.PredicateInput{
						Type:  ol.PredicateTypeEnumEquals,
						Value: "Requests",
					},
				})
			},
		},
		"UpdateAlertSourceUsage": {
			fixture: NewTestRequest(
				`"mutation CheckAlertSourceUsageUpdate($input:CheckAlertSourceUsageUpdateInput!){checkAlertSourceUsageUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "alertSourceNamePredicate": {"type":"equals", "value":"Requests"}, "alertSourceType":"datadog" } }`,
				`{ "data": { "checkAlertSourceUsageUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_alert_source_usage",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.UpdateCheckAlertSourceUsage(ol.CheckAlertSourceUsageUpdateInput{
					CheckUpdateInput: checkUpdateInput,
					AlertSourceType:  ol.AlertSourceTypeEnumDatadog,
					AlertSourceNamePredicate: &ol.PredicateUpdateInput{
						Type:  ol.PredicateTypeEnumEquals,
						Value: "Requests",
					},
				})
			},
		},

		"CreateGitBranchProtection": {
			fixture: NewTestRequest(
				`"mutation CheckGitBranchProtectionCreate($input:CheckGitBranchProtectionCreateInput!){checkGitBranchProtectionCreate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check" } }`,
				`{ "data": { "checkGitBranchProtectionCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_git_branch_protection",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.CreateCheckGitBranchProtection(ol.CheckGitBranchProtectionCreateInput{
					CheckCreateInput: checkCreateInput,
				})
			},
		},
		"UpdateGitBranchProtection": {
			fixture: NewTestRequest(
				`"mutation CheckGitBranchProtectionUpdate($input:CheckGitBranchProtectionUpdateInput!){checkGitBranchProtectionUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }} } }`,
				`{ "data": { "checkGitBranchProtectionUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_git_branch_protection",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.UpdateCheckGitBranchProtection(ol.CheckGitBranchProtectionUpdateInput{
					CheckUpdateInput: checkUpdateInput,
				})
			},
		},

		"CreateHasRecentDeploy": {
			fixture: NewTestRequest(
				`"mutation CheckHasRecentDeployCreate($input:CheckHasRecentDeployCreateInput!){checkHasRecentDeployCreate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "days": 5 } }`,
				`{ "data": { "checkHasRecentDeployCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_has_recent_deploy",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.CreateCheckHasRecentDeploy(ol.CheckHasRecentDeployCreateInput{
					CheckCreateInput: checkCreateInput,
					Days:             5,
				})
			},
		},
		"UpdateHasRecentDeploy": {
			fixture: NewTestRequest(
				`"mutation CheckHasRecentDeployUpdate($input:CheckHasRecentDeployUpdateInput!){checkHasRecentDeployUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "days": 5 } }`,
				`{ "data": { "checkHasRecentDeployUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_has_recent_deploy",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.UpdateCheckHasRecentDeploy(ol.CheckHasRecentDeployUpdateInput{
					CheckUpdateInput: checkUpdateInput,
					Days:             ol.NewInt(5),
				})
			},
		},

		"CreateHasDocumentation": {
			fixture: NewTestRequest(
				`"mutation CheckHasDocumentationCreate($input:CheckHasDocumentationCreateInput!){checkHasDocumentationCreate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "documentType": "api", "documentSubtype": "openapi", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check" } }`,
				`{ "data": { "checkHasDocumentationCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has valid documentation.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World", "notes": "Hello World Check", "documentType": "api", "documentSubtype": "openapi" }, "errors": [] } } }`,
			),
			endpoint: "check/create_has_documentation",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.CreateCheckHasDocumentation(ol.CheckHasDocumentationCreateInput{
					CheckCreateInput: checkCreateInput,
					DocumentType:     "api",
					DocumentSubtype:  "openapi",
				})
			},
		},
		"UpdateHasDocumentation": {
			fixture: NewTestRequest(
				`"mutation CheckHasDocumentationUpdate($input:CheckHasDocumentationUpdateInput!){checkHasDocumentationUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "documentType": "api", "documentSubtype": "openapi", {{ template "check_base_vars" }} } }`,
				`{ "data": { "checkHasDocumentationUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World Update", "notes": "Hello World Update", "documentType": "api", "documentSubtype": "openapi" }, "errors": [] } } }`,
			),
			endpoint: "check/update_has_documentation",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.UpdateCheckHasDocumentation(ol.CheckHasDocumentationUpdateInput{
					CheckUpdateInput: checkUpdateInput,
					DocumentType:     "api",
					DocumentSubtype:  "openapi",
				})
			},
		},

		"CreateCustomEvent": {
			fixture: NewTestRequest(
				`"mutation CheckCustomEventCreate($input:CheckCustomEventCreateInput!){checkCustomEventCreate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "passPending": false, "serviceSelector": ".metadata.name", "successCondition": ".metadata.name", "resultMessage": "#Hello World", "integrationId": "Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg" } }`,
				`{ "data": { "checkCustomEventCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_custom_event",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.CreateCheckCustomEvent(ol.CheckCustomEventCreateInput{
					CheckCreateInput: checkCreateInput,
					ServiceSelector:  ".metadata.name",
					SuccessCondition: ".metadata.name",
					Message:          "#Hello World",
					Integration:      "Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg",
					PassPending:      ol.Bool(false),
				})
			},
		},
		"UpdateCustomEvent": {
			fixture: NewTestRequest(
				` "mutation CheckCustomEventUpdate($input:CheckCustomEventUpdateInput!){checkCustomEventUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "passPending": false, "serviceSelector": ".metadata.name", "successCondition": ".metadata.name", "resultMessage": "#Hello World", "integrationId": "Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg" } }`,
				`{ "data": { "checkCustomEventUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_custom_event",
			body: func(c *ol.Client) (*ol.Check, error) {
				message := "#Hello World"
				return c.UpdateCheckCustomEvent(ol.CheckCustomEventUpdateInput{
					CheckUpdateInput: checkUpdateInput,
					ServiceSelector:  ".metadata.name",
					SuccessCondition: ".metadata.name",
					Message:          &message,
					Integration:      ol.NewID("Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg"),
					PassPending:      ol.Bool(false),
				})
			},
		},
		"UpdateCustomEventNoMessage": {
			fixture: NewTestRequest(
				`"mutation CheckCustomEventUpdate($input:CheckCustomEventUpdateInput!){checkCustomEventUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "passPending": false, "serviceSelector": ".metadata.name", "successCondition": ".metadata.name", "resultMessage": "", "integrationId": "Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg" } }`,
				`{ "data": { "checkCustomEventUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a JSON payload to be sent to the integration endpoint to complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_custom_event_no_message",
			body: func(c *ol.Client) (*ol.Check, error) {
				message := ""
				return c.UpdateCheckCustomEvent(ol.CheckCustomEventUpdateInput{
					CheckUpdateInput: checkUpdateInput,
					ServiceSelector:  ".metadata.name",
					SuccessCondition: ".metadata.name",
					Message:          &message,
					Integration:      ol.NewID("Z2lkOi8vb3BzbGV2ZWwvSW50ZWdyYXRpb25zOjpFdmVudHM6OkdlbmVyaWNJbnRlZ3JhdGlvbi81Njg"),
					PassPending:      ol.Bool(false),
				})
			},
		},
		"CreateManual": {
			fixture: NewTestRequest(
				`"mutation CheckManualCreate($input:CheckManualCreateInput!){checkManualCreate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "updateFrequency": { "startingDate": "2021-07-26T20:22:44.427Z", "frequencyTimeScale": "week", "frequencyValue": 1 }, "updateRequiresComment": false } }`,
				`{ "data": { "checkManualCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a service owner to manually complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_manual",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.CreateCheckManual(ol.CheckManualCreateInput{
					CheckCreateInput: checkCreateInput,
					UpdateFrequency:  ol.NewManualCheckFrequencyInput("2021-07-26T20:22:44.427Z", ol.FrequencyTimeScaleWeek, 1),
				})
			},
		},
		"UpdateManual": {
			fixture: NewTestRequest(
				`"mutation CheckManualUpdate($input:CheckManualUpdateInput!){checkManualUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "updateFrequency": { "startingDate": "2021-07-26T20:22:44.427Z", "frequencyTimeScale": "week", "frequencyValue": 1 } } }`,
				`{ "data": { "checkManualUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Requires a service owner to manually complete a check for the service.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_manual",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.UpdateCheckManual(ol.CheckManualUpdateInput{
					CheckUpdateInput: checkUpdateInput,
					UpdateFrequency:  ol.NewManualCheckFrequencyInput("2021-07-26T20:22:44.427Z", ol.FrequencyTimeScaleWeek, 1),
				})
			},
		},
		"CreateRepositoryFile": {
			fixture: NewTestRequest(
				`"mutation CheckRepositoryFileCreate($input:CheckRepositoryFileCreateInput!){checkRepositoryFileCreate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "directorySearch": true, "filePaths": [ "/src", "/test" ], "fileContentsPredicate": { "type": "equals", "value": "postgres" }, "useAbsoluteRoot": true } }`,
				`{ "data": { "checkRepositoryFileCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Repo directorys ''/src' or '/test'' equals 'postgres'.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_repo_file",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.CreateCheckRepositoryFile(ol.CheckRepositoryFileCreateInput{
					CheckCreateInput: checkCreateInput,
					DirectorySearch:  true,
					Filepaths:        []string{"/src", "/test"},
					FileContentsPredicate: &ol.PredicateInput{
						Type:  ol.PredicateTypeEnumEquals,
						Value: "postgres",
					},
					UseAbsoluteRoot: true,
				})
			},
		},
		"UpdateRepositoryFile": {
			fixture: NewTestRequest(
				`"mutation CheckRepositoryFileUpdate($input:CheckRepositoryFileUpdateInput!){checkRepositoryFileUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "directorySearch": true, "filePaths": [ "/src", "/test" ], "fileContentsPredicate": { "type": "equals", "value": "postgres" }, "useAbsoluteRoot": false } }`,
				`{ "data": { "checkRepositoryFileUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Repo directorys ''/src' or '/test'' equals 'postgres'.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_repo_file",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.UpdateCheckRepositoryFile(ol.CheckRepositoryFileUpdateInput{
					CheckUpdateInput: checkUpdateInput,
					DirectorySearch:  true,
					Filepaths:        []string{"/src", "/test"},
					FileContentsPredicate: &ol.PredicateInput{
						Type:  ol.PredicateTypeEnumEquals,
						Value: "postgres",
					},
					UseAbsoluteRoot: false,
				})
			},
		},
		"CreateRepositoryGrep": {
			fixture: NewTestRequest(
				`"mutation CheckRepositoryGrepCreate($input:CheckRepositoryGrepCreateInput!){checkRepositoryGrepCreate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "directorySearch": true, "filePaths": [ "**/hello.go" ], "fileContentsPredicate": { "type": "exists" } } }`,
				`{ "data": { "checkRepositoryGrepCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNzEw", "name": "Performance" }, "description": "Verifies the existence and/or contents of files in a service's attached Git repositories.", "enabled": true, "enableOn": null, "filter": null, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDQ1", "index": 1, "name": "Bronze" }, "name": "Hello World", "notes": "Hello World Check", "owner": null, "type": "repo_grep" }, "errors": [] } } }`,
			),
			endpoint: "check/create_repo_grep",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.CreateCheckRepositoryGrep(ol.CheckRepositoryGrepCreateInput{
					CheckCreateInput: checkCreateInput,
					DirectorySearch:  true,
					Filepaths:        []string{"**/hello.go"},
					FileContentsPredicate: &ol.PredicateInput{
						Type:  ol.PredicateTypeEnumExists,
						Value: "",
					},
				})
			},
		},
		"UpdateRepositoryGrep": {
			fixture: NewTestRequest(
				` "mutation CheckRepositoryGrepUpdate($input:CheckRepositoryGrepUpdateInput!){checkRepositoryGrepUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "directorySearch": true, "filePaths": [ "**/go.mod" ], "fileContentsPredicate": { "type": "exists" } } }`,
				`{ "data": { "checkRepositoryGrepUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNzEw", "name": "Performance" }, "description": "Verifies the existence and/or contents of files in a service's attached Git repositories.", "enabled": true, "enableOn": null, "filter": null, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDQ1", "index": 1, "name": "Bronze" }, "name": "Hello World", "notes": "Hello World Check", "owner": null, "type": "repo_grep" }, "errors": [] } } }`,
			),
			endpoint: "check/update_repo_grep",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.UpdateCheckRepositoryGrep(ol.CheckRepositoryGrepUpdateInput{
					CheckUpdateInput: checkUpdateInput,
					DirectorySearch:  true,
					Filepaths:        []string{"**/go.mod"},
					FileContentsPredicate: &ol.PredicateInput{
						Type:  ol.PredicateTypeEnumExists,
						Value: "",
					},
				})
			},
		},
		"UpdateRepositoryGrepMissingDirectorySearch": {
			fixture: NewTestRequest(
				`"mutation CheckRepositoryGrepUpdate($input:CheckRepositoryGrepUpdateInput!){checkRepositoryGrepUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "filePaths": [ "**/go.mod" ], "directorySearch": false, "fileContentsPredicate": { "type": "exists" } } }`,
				`{ "data": { "checkRepositoryGrepUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNzEw", "name": "Performance" }, "description": "Verifies the existence and/or contents of files in a service's attached Git repositories.", "enabled": true, "enableOn": null, "filter": null, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvNDQ1", "index": 1, "name": "Bronze" }, "name": "Hello World", "notes": "Hello World Check", "owner": null, "type": "repo_grep" }, "errors": [] } } }`,
			),
			endpoint: "check/update_repo_grep_missing_directory_search",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.UpdateCheckRepositoryGrep(ol.CheckRepositoryGrepUpdateInput{
					CheckUpdateInput: checkUpdateInput,
					Filepaths:        []string{"**/go.mod"},
					FileContentsPredicate: &ol.PredicateInput{
						Type:  ol.PredicateTypeEnumExists,
						Value: "",
					},
				})
			},
		},
		"CreateRepositoryIntegrated": {
			fixture: NewTestRequest(
				`"mutation CheckRepositoryIntegratedCreate($input:CheckRepositoryIntegratedCreateInput!){checkRepositoryIntegratedCreate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check" } }`,
				`{ "data": { "checkRepositoryIntegratedCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has a repository integrated.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_repo_integrated",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.CreateCheckRepositoryIntegrated(ol.CheckRepositoryIntegratedCreateInput{
					CheckCreateInput: checkCreateInput,
				})
			},
		},
		"UpdateRepositoryIntegrated": {
			fixture: NewTestRequest(
				`"mutation CheckRepositoryIntegratedUpdate($input:CheckRepositoryIntegratedUpdateInput!){checkRepositoryIntegratedUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }} } }`,
				`{ "data": { "checkRepositoryIntegratedUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has a repository integrated.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_repo_integrated",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.UpdateCheckRepositoryIntegrated(ol.CheckRepositoryIntegratedUpdateInput{
					CheckUpdateInput: checkUpdateInput,
				})
			},
		},
		"CreateRepositorySearch": {
			fixture: NewTestRequest(
				`"mutation CheckRepositorySearchCreate($input:CheckRepositorySearchCreateInput!){checkRepositorySearchCreate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "fileExtensions": [ "sbt", "py" ], "fileContentsPredicate": { "type": "contains", "value": "postgres" } } }`,
				`{ "data": { "checkRepositorySearchCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Repo contains search term 'postgres' in at least one .sbt or .py file.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_repo_search",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.CreateCheckRepositorySearch(ol.CheckRepositorySearchCreateInput{
					CheckCreateInput: checkCreateInput,
					FileExtensions:   []string{"sbt", "py"},
					FileContentsPredicate: ol.PredicateInput{
						Type:  ol.PredicateTypeEnumContains,
						Value: "postgres",
					},
				})
			},
		},
		"UpdateRepositorySearch": {
			fixture: NewTestRequest(
				`"mutation CheckRepositorySearchUpdate($input:CheckRepositorySearchUpdateInput!){checkRepositorySearchUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "fileExtensions": [ "sbt", "py" ], "fileContentsPredicate": { "type": "contains", "value": "postgres" } } }`,
				`{ "data": { "checkRepositorySearchUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Repo contains search term 'postgres' in at least one .sbt or .py file.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_repo_search",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.UpdateCheckRepositorySearch(ol.CheckRepositorySearchUpdateInput{
					CheckUpdateInput: checkUpdateInput,
					FileExtensions:   []string{"sbt", "py"},
					FileContentsPredicate: &ol.PredicateInput{
						Type:  ol.PredicateTypeEnumContains,
						Value: "postgres",
					},
				})
			},
		},
		"CreateServiceDependency": {
			fixture: NewTestRequest(
				`"mutation CheckServiceDependencyCreate($input:CheckServiceDependencyCreateInput!){checkServiceDependencyCreate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check" } }`,
				`{ "data": { "checkServiceDependencyCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has either a dependent or dependency.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_service_dependency",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.CreateCheckServiceDependency(ol.CheckServiceDependencyCreateInput{
					CheckCreateInput: checkCreateInput,
				})
			},
		},
		"UpdateServiceDependency": {
			fixture: NewTestRequest(
				`"mutation CheckServiceDependencyUpdate($input:CheckServiceDependencyUpdateInput!){checkServiceDependencyUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }} } }`,
				`{ "data": { "checkServiceDependencyUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has either a dependent or dependency.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_service_dependency",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.UpdateCheckServiceDependency(ol.CheckServiceDependencyUpdateInput{
					CheckUpdateInput: checkUpdateInput,
				})
			},
		},
		"CreateServiceConfiguration": {
			fixture: NewTestRequest(
				`"mutation CheckServiceConfigurationCreate($input:CheckServiceConfigurationCreateInput!){checkServiceConfigurationCreate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check" } }`,
				`{ "data": { "checkServiceConfigurationCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service is maintained though the use of an opslevel.yml service config.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_service_configuration",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.CreateCheckServiceConfiguration(ol.CheckServiceConfigurationCreateInput{
					CheckCreateInput: checkCreateInput,
				})
			},
		},
		"UpdateServiceConfiguration": {
			fixture: NewTestRequest(
				`"mutation CheckServiceConfigurationUpdate($input:CheckServiceConfigurationUpdateInput!){checkServiceConfigurationUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }} } }`,
				`{ "data": { "checkServiceConfigurationUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service is maintained though the use of an opslevel.yml service config.", "enabled": true, "filter": null, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World", "notes": "Hello World Check" }, "errors": [] } } }`,
			),
			endpoint: "check/update_service_configuration",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.UpdateCheckServiceConfiguration(ol.CheckServiceConfigurationUpdateInput{
					CheckUpdateInput: checkUpdateInput,
				})
			},
		},
		"CreateServiceOwnership": {
			fixture: NewTestRequest(
				`"mutation CheckServiceOwnershipCreate($input:CheckServiceOwnershipCreateInput!){checkServiceOwnershipCreate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "requireContactMethod": true, "contactMethod": "slack", "tagKey": "updated_at", "tagPredicate": { "type": "equals", "value": "2-11-2022" } } }`,
				`{ "data": { "checkServiceOwnershipCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has an owner defined.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_service_ownership",
			body: func(c *ol.Client) (*ol.Check, error) {
				slackType := ol.ContactTypeSlack
				return c.CreateCheckServiceOwnership(ol.CheckServiceOwnershipCreateInput{
					CheckCreateInput:     checkCreateInput,
					RequireContactMethod: ol.Bool(true),
					ContactMethod:        &slackType,
					TeamTagKey:           "updated_at",
					TeamTagPredicate: &ol.PredicateInput{
						Type:  ol.PredicateTypeEnumEquals,
						Value: "2-11-2022",
					},
				})
			},
		},
		"UpdateServiceOwnership": {
			fixture: NewTestRequest(
				`"mutation CheckServiceOwnershipUpdate($input:CheckServiceOwnershipUpdateInput!){checkServiceOwnershipUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "requireContactMethod": true, "contactMethod": "email", "tagKey": "updated_at", "tagPredicate": { "type": "equals", "value": "2-11-2022" } } }`,
				`{ "data": { "checkServiceOwnershipUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has an owner defined.", "enabled": true, "filter": null, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World", "notes": "Hello World Check" }, "errors": [] } } }`,
			),
			endpoint: "check/update_service_ownership",
			body: func(c *ol.Client) (*ol.Check, error) {
				emailType := ol.ContactTypeEmail
				return c.UpdateCheckServiceOwnership(ol.CheckServiceOwnershipUpdateInput{
					CheckUpdateInput:     checkUpdateInput,
					RequireContactMethod: ol.Bool(true),
					ContactMethod:        &emailType,
					TeamTagKey:           "updated_at",
					TeamTagPredicate: &ol.PredicateUpdateInput{
						Type:  ol.PredicateTypeEnumEquals,
						Value: "2-11-2022",
					},
				})
			},
		},
		"CreateServiceProperty": {
			fixture: NewTestRequest(
				`"mutation CheckServicePropertyCreate($input:CheckServicePropertyCreateInput!){checkServicePropertyCreate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "serviceProperty": "framework", "propertyValuePredicate": { "type": "equals", "value": "postgres" } } }`,
				`{ "data": { "checkServicePropertyCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "The service has a framework that equals <code>postgres</code>", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_service_property",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.CreateCheckServiceProperty(ol.CheckServicePropertyCreateInput{
					CheckCreateInput: checkCreateInput,
					Property:         ol.ServicePropertyTypeEnumFramework,
					Predicate: &ol.PredicateInput{
						Type:  ol.PredicateTypeEnumEquals,
						Value: "postgres",
					},
				})
			},
		},
		"UpdateServiceProperty": {
			fixture: NewTestRequest(
				`"mutation CheckServicePropertyUpdate($input:CheckServicePropertyUpdateInput!){checkServicePropertyUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "serviceProperty": "framework", "propertyValuePredicate": { "type": "equals", "value": "postgres" } } }`,
				`{ "data": { "checkServicePropertyUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "The service has a framework that equals <code>postgres</code>", "enabled": true, "filter": null, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World", "notes": "Hello World Check" }, "errors": [] } } }`,
			),
			endpoint: "check/update_service_property",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.UpdateCheckServiceProperty(ol.CheckServicePropertyUpdateInput{
					CheckUpdateInput: checkUpdateInput,
					Property:         ol.ServicePropertyTypeEnumFramework,
					Predicate: &ol.PredicateInput{
						Type:  ol.PredicateTypeEnumEquals,
						Value: "postgres",
					},
				})
			},
		},
		"CreateTagDefined": {
			fixture: NewTestRequest(
				`"mutation CheckTagDefinedCreate($input:CheckTagDefinedCreateInput!){checkTagDefinedCreate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "tagKey": "app", "tagPredicate": { "type": "equals", "value": "postgres" } } }`,
				`{ "data": { "checkTagDefinedCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has the specified tag defined.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_tag_defined",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.CreateCheckTagDefined(ol.CheckTagDefinedCreateInput{
					CheckCreateInput: checkCreateInput,
					TagKey:           "app",
					TagPredicate: &ol.PredicateInput{
						Type:  ol.PredicateTypeEnumEquals,
						Value: "postgres",
					},
				})
			},
		},
		"UpdateTagDefined": {
			fixture: NewTestRequest(
				`"mutation CheckTagDefinedUpdate($input:CheckTagDefinedUpdateInput!){checkTagDefinedUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "tagKey": "app", "tagPredicate": { "type": "equals", "value": "postgres" } } }`,
				`{ "data": { "checkTagDefinedUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has the specified tag defined.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_tag_defined",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.UpdateCheckTagDefined(ol.CheckTagDefinedUpdateInput{
					CheckUpdateInput: checkUpdateInput,
					TagKey:           "app",
					TagPredicate: &ol.PredicateInput{
						Type:  ol.PredicateTypeEnumEquals,
						Value: "postgres",
					},
				})
			},
		},
		"CreateToolUsage": {
			fixture: NewTestRequest(
				`"mutation CheckToolUsageCreate($input:CheckToolUsageCreateInput!){checkToolUsageCreate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "name": "Hello World", "enabled": true, "categoryId": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "levelId": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "notes": "Hello World Check", "toolCategory": "metrics", "toolNamePredicate": { "type": "equals", "value": "datadog" }, "toolUrlPredicate": { "type": "contains", "value": "https://" }, "environmentPredicate": { "type": "equals", "value": "production" } } }`,
				`{ "data": { "checkToolUsageCreate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "The service is using 'datadog' as a metrics tool in the 'production' environment.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/create_tool_usage",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.CreateCheckToolUsage(ol.CheckToolUsageCreateInput{
					CheckCreateInput: checkCreateInput,
					ToolCategory:     ol.ToolCategoryMetrics,
					ToolNamePredicate: &ol.PredicateInput{
						Type:  ol.PredicateTypeEnumEquals,
						Value: "datadog",
					},
					ToolUrlPredicate: &ol.PredicateInput{
						Type:  ol.PredicateTypeEnumContains,
						Value: "https://",
					},
					EnvironmentPredicate: &ol.PredicateInput{
						Type:  ol.PredicateTypeEnumEquals,
						Value: "production",
					},
				})
			},
		},
		"UpdateToolUsage": {
			fixture: NewTestRequest(
				`"mutation CheckToolUsageUpdate($input:CheckToolUsageUpdateInput!){checkToolUsageUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
				`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", {{ template "check_base_vars" }}, "toolCategory": "metrics", "toolNamePredicate": { "type": "equals", "value": "datadog" }, "toolUrlPredicate": { "type": "contains", "value": "https://" }, "environmentPredicate": { "type": "equals", "value": "production" } } }`,
				`{ "data": { "checkToolUsageUpdate": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "The service is using 'datadog' as a metrics tool in the 'production' environment.", "enabled": true, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Hello World" }, "errors": [] } } }`,
			),
			endpoint: "check/update_tool_usage",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.UpdateCheckToolUsage(ol.CheckToolUsageUpdateInput{
					CheckUpdateInput: checkUpdateInput,
					ToolCategory:     ol.ToolCategoryMetrics,
					ToolNamePredicate: &ol.PredicateInput{
						Type:  ol.PredicateTypeEnumEquals,
						Value: "datadog",
					},
					ToolUrlPredicate: &ol.PredicateInput{
						Type:  ol.PredicateTypeEnumContains,
						Value: "https://",
					},
					EnvironmentPredicate: &ol.PredicateInput{
						Type:  ol.PredicateTypeEnumEquals,
						Value: "production",
					},
				})
			},
		},
		"GetCheck": {
			fixture: NewTestRequest(
				`"query CheckGet($id:ID!){account{check(id: $id){category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}}}}"`,
				`{ "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4" }`,
				`{ "data": { "account": { "check": { "category": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1", "name": "Performance" }, "description": "Verifies that the service has an owner defined.", "enabled": true, "filter": null, "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "level": { "alias": "bronze", "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.", "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3", "index": 1, "name": "Bronze" }, "name": "Owner Defined", "notes": null } } } }`,
			),
			endpoint: "check/get",
			body: func(c *ol.Client) (*ol.Check, error) {
				return c.GetCheck("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4")
			},
		},
	}
	return testcases
}

func TestChecks(t *testing.T) {
	for name, tc := range getCheckTestCases() {
		t.Run(name, func(t *testing.T) {
			// Arrange
			client := BestTestClient(t, name, tc.fixture)
			// Act
			result, err := tc.body(client)
			// Assert
			autopilot.Equals(t, nil, err)
			autopilot.Equals(t, "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", string(result.Id))
			autopilot.Equals(t, "Performance", result.Category.Name)
			autopilot.Equals(t, "Bronze", result.Level.Name)
		})
	}
}

func TestCanUpdateFilterToNull(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation CheckCustomEventUpdate($input:CheckCustomEventUpdateInput!){checkCustomEventUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
		`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "filterId": null }}`,
		`{ "data": {
      "checkCustomEventUpdate": {
        "check": {
          "category": {
            "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1",
            "name": "Performance"
          },
          "enabled": true,
          "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4",
          "level": {
            "alias": "bronze",
            "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.",
            "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3",
            "index": 1,
            "name": "Bronze"
          },
          "notes": "Hello World Notes",
          "name": "Hello World"
        },
        "errors": []
      }}}`,
	)
	client := BestTestClient(t, "check/can_update_filter_to_null", testRequest)
	// Act
	result, err := client.UpdateCheckCustomEvent(ol.CheckCustomEventUpdateInput{
		CheckUpdateInput: ol.CheckUpdateInput{
			Id:     "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4",
			Filter: ol.NewID(),
		},
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Hello World", result.Name)
	autopilot.Equals(t, ol.ID(""), result.Filter.Id)
}

func TestCanUpdateNotesToNull(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation CheckCustomEventUpdate($input:CheckCustomEventUpdateInput!){checkCustomEventUpdate(input: $input){check{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},errors{message,path}}}"`,
		`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4", "notes": "" }}`,
		`{ "data": {
      "checkCustomEventUpdate": {
        "check": {
          "category": {
            "id": "Z2lkOi8vb3BzbGV2ZWwvQ2F0ZWdvcnkvNjA1",
            "name": "Performance"
          },
          "enabled": true,
          "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4",
          "level": {
            "alias": "bronze",
            "description": "Services in this level satisfy critical checks. This is the minimum standard to ship to production.",
            "id": "Z2lkOi8vb3BzbGV2ZWwvTGV2ZWwvMzE3",
            "index": 1,
            "name": "Bronze"
          },
          "name": "Hello World"
        },
        "errors": []
      }}}`,
	)
	client := BestTestClient(t, "check/can_update_notes_to_null", testRequest)
	// Act
	result, err := client.UpdateCheckCustomEvent(ol.CheckCustomEventUpdateInput{
		CheckUpdateInput: ol.CheckUpdateInput{
			Id:    "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDE4",
			Notes: ol.NewString(""),
		},
	})
	// Assert
	autopilot.Equals(t, nil, err)
	autopilot.Equals(t, "Hello World", result.Name)
	autopilot.Equals(t, "", result.Notes)
}

func TestListChecks(t *testing.T) {
	// Arrange
	testRequestOne := NewTestRequest(
		`"query CheckList($after:String!$first:Int!){account{rubric{checks(after: $after, first: $first){nodes{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},{{ template "pagination_request" }},totalCount}}}}"`,
		`{{ template "pagination_initial_query_variables" }}`,
		`{ "data": { "account": { "rubric": { "checks": { "nodes": [ { {{ template "common_check_response" }} }, { {{ template "metrics_tool_check" }} } ], {{ template "pagination_initial_pageInfo_response" }}, "totalCount": 2 }}}}}`,
	)
	testRequestTwo := NewTestRequest(
		`"query CheckList($after:String!$first:Int!){account{rubric{checks(after: $after, first: $first){nodes{category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}},{{ template "pagination_request" }},totalCount}}}}"`,
		`{{ template "pagination_second_query_variables" }}`,
		`{ "data": { "account": { "rubric": { "checks": { "nodes": [ { {{ template "owner_defined_check" }} } ], {{ template "pagination_second_pageInfo_response" }}, "totalCount": 1 }}}}}`,
	)
	requests := []TestRequest{testRequestOne, testRequestTwo}

	client := BestTestClient(t, "check/list", requests...)
	// Act
	response, err := client.ListChecks(nil)
	result := response.Nodes
	// Assert
	autopilot.Ok(t, err)
	autopilot.Equals(t, 3, response.TotalCount)
	autopilot.Equals(t, "Metrics Tool", result[1].Name)
	autopilot.Equals(t, "Tier 1 Services", result[1].Filter.Name)
	autopilot.Equals(t, "Owner Defined", result[2].Name)
	autopilot.Equals(t, "Verifies that the service has an owner defined.", result[2].Description)
}

func TestGetMissingCheck(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"query CheckGet($id:ID!){account{check(id: $id){category{id,name},description,enabled,enableOn,filter{connective,htmlUrl,id,name,predicates{key,keyData,type,value,caseSensitive}},id,level{alias,description,id,index,name},name,notes: rawNotes,owner{... on Team{alias,id}},type,... on AlertSourceUsageCheck{alertSourceNamePredicate{type,value},alertSourceType},... on CustomEventCheck{integration{id,name,type},passPending,resultMessage,serviceSelector,successCondition},... on HasRecentDeployCheck{days},... on ManualCheck{updateFrequency{startingDate,frequencyTimeScale,frequencyValue},updateRequiresComment},... on RepositoryFileCheck{directorySearch,filePaths,fileContentsPredicate{type,value},useAbsoluteRoot},... on RepositoryGrepCheck{directorySearch,filePaths,fileContentsPredicate{type,value}},... on RepositorySearchCheck{fileExtensions,fileContentsPredicate{type,value}},... on ServiceOwnershipCheck{requireContactMethod,contactMethod,tagKey,tagPredicate{type,value}},... on ServicePropertyCheck{serviceProperty,propertyValuePredicate{type,value}},... on TagDefinedCheck{tagKey,tagPredicate{type,value}},... on ToolUsageCheck{toolCategory,toolNamePredicate{type,value},toolUrlPredicate{type,value},environmentPredicate{type,value}},... on HasDocumentationCheck{documentType,documentSubtype}}}}"`,
		`{ "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf" }`,
		`{ "data": { "account": { "check": null } } }`,
	)
	client := BestTestClient(t, "check/get_missing", testRequest)
	// Act
	_, err := client.GetCheck("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpIYXNPd25lci8yNDEf")
	// Assert
	autopilot.Assert(t, err != nil, "This test should throw an error.")
}

func TestDeleteCheck(t *testing.T) {
	// Arrange
	testRequest := NewTestRequest(
		`"mutation CheckDelete($input:CheckDeleteInput!){checkDelete(input: $input){errors{message,path}}}"`,
		`{ "input": { "id": "Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpHZW5lcmljLzIxNzI" } }`,
		`{ "data": { "checkDelete": { "errors": [] } } }`,
	)
	client := BestTestClient(t, "check/delete", testRequest)
	// Act
	err := client.DeleteCheck("Z2lkOi8vb3BzbGV2ZWwvQ2hlY2tzOjpHZW5lcmljLzIxNzI")
	// Assert
	autopilot.Equals(t, nil, err)
}

func TestJsonUnmarshalCreateCheck(t *testing.T) {
	// Arrange
	data := `{
	"name": "Example",
	"notes": "Example Notes",
	"updateRequiresComment": false
}`
	output := ol.CheckManualCreateInput{
		CheckCreateInput: ol.CheckCreateInput{
			Name:  "Example",
			Notes: "Example Notes",
		},
		UpdateRequiresComment: false,
	}
	// Act
	buf1, err1 := ol.UnmarshalCheckCreateInput(ol.CheckTypeManual, []byte(data))
	// Assert
	autopilot.Ok(t, err1)
	autopilot.Equals(t, &output, buf1.(*ol.CheckManualCreateInput))
}

func TestJsonUnmarshalUpdateCheck(t *testing.T) {
	// Arrange
	data := `{
	"name": "Example",
	"notes": "Example Notes",
	"updateRequiresComment": true
}`
	output := ol.CheckManualUpdateInput{
		CheckUpdateInput: ol.CheckUpdateInput{
			Name:  "Example",
			Notes: ol.NewString("Example Notes"),
		},
		UpdateRequiresComment: true,
	}
	// Act
	buf1, err1 := ol.UnmarshalCheckUpdateInput(ol.CheckTypeManual, []byte(data))
	// Assert
	autopilot.Ok(t, err1)
	autopilot.Equals(t, &output, buf1.(*ol.CheckManualUpdateInput))
}
