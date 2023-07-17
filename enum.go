// Code generated by gen.go; DO NOT EDIT.

package opslevel

// AlertSourceTypeEnum The monitor status level.
type AlertSourceStatusTypeEnum string

const (
	AlertSourceStatusTypeEnumAlert        AlertSourceStatusTypeEnum = "alert"         // Monitor is reporting an alert.
	AlertSourceStatusTypeEnumFetchingData AlertSourceStatusTypeEnum = "fetching_data" // Monitor currently being updated.
	AlertSourceStatusTypeEnumNoData       AlertSourceStatusTypeEnum = "no_data"       // No data received yet. Ensure your monitors are configured correctly.
	AlertSourceStatusTypeEnumOk           AlertSourceStatusTypeEnum = "ok"            // Monitor is not reporting any warnings or alerts.
	AlertSourceStatusTypeEnumWarn         AlertSourceStatusTypeEnum = "warn"          // Monitor is reporting a warning.
)

// All AlertSourceStatusTypeEnum as []string
var AllAlertSourceStatusTypeEnum = []string{
	string(AlertSourceStatusTypeEnumAlert),
	string(AlertSourceStatusTypeEnumFetchingData),
	string(AlertSourceStatusTypeEnumNoData),
	string(AlertSourceStatusTypeEnumOk),
	string(AlertSourceStatusTypeEnumWarn),
}

// AlertSourceTypeEnum represents the type of the alert source.
type AlertSourceTypeEnum string

const (
	AlertSourceTypeEnumDatadog   AlertSourceTypeEnum = "datadog"   // A Datadog alert source (aka monitor).
	AlertSourceTypeEnumOpsgenie  AlertSourceTypeEnum = "opsgenie"  // An Opsgenie alert source (aka service).
	AlertSourceTypeEnumPagerduty AlertSourceTypeEnum = "pagerduty" // A PagerDuty alert source (aka service).
	AlertSourceTypeEnumNewRelic   AlertSourceTypeEnum = "new_relic"   // A New Relic alert source (aka monitor).
)

// All AlertSourceTypeEnum as []string
var AllAlertSourceTypeEnum = []string{
	string(AlertSourceTypeEnumDatadog),
	string(AlertSourceTypeEnumOpsgenie),
	string(AlertSourceTypeEnumPagerduty),
	string(AlertSourceTypeEnumNewRelic),
}

// AliasOwnerTypeEnum represents the owner type an alias is assigned to.
type AliasOwnerTypeEnum string

const (
	AliasOwnerTypeEnumService                 AliasOwnerTypeEnum = "service" // Aliases that are assigned to services.
	AliasOwnerTypeEnumTeam                    AliasOwnerTypeEnum = "team"    // Aliases that are assigned to teams.
	AliasOwnerTypeEnumSystem                  AliasOwnerTypeEnum = "system"  // Aliases that are assigned to systems.
	AliasOwnerTypeEnumDomain                  AliasOwnerTypeEnum = "domain"  // Aliases that are assigned to domains.
	AliasOwnerTypeEnumGroup                   AliasOwnerTypeEnum = "group"   // Aliases that are assigned to groups.
	AliasOwnerTypeEnumInfrastructureResource  AliasOwnerTypeEnum = "infrastructure_resource"   // Aliases that are assigned to infrastructure resources
)

// All AliasOwnerTypeEnum as []string
var AllAliasOwnerTypeEnum = []string{
	string(AliasOwnerTypeEnumService),
	string(AliasOwnerTypeEnumTeam),
	string(AliasOwnerTypeEnumSystem),
	string(AliasOwnerTypeEnumDomain),
	string(AliasOwnerTypeEnumGroup),
	string(AliasOwnerTypeEnumInfrastructureResource),
}

// ApiDocumentSourceEnum represents the source used to determine the preferred API document.
type ApiDocumentSourceEnum string

const (
	ApiDocumentSourceEnumPull ApiDocumentSourceEnum = "PULL" // Use the document that was pulled by OpsLevel via a repo.
	ApiDocumentSourceEnumPush ApiDocumentSourceEnum = "PUSH" // Use the document that was pushed to OpsLevel via an API Docs integration.
)

// All ApiDocumentSourceEnum as []string
var AllApiDocumentSourceEnum = []string{
	string(ApiDocumentSourceEnumPull),
	string(ApiDocumentSourceEnumPush),
}

// BasicTypeEnum
type BasicTypeEnum string

const (
	BasicTypeEnumDoesNotEqual BasicTypeEnum = "does_not_equal" //
	BasicTypeEnumEquals       BasicTypeEnum = "equals"         //
)

// All BasicTypeEnum as []string
var AllBasicTypeEnum = []string{
	string(BasicTypeEnumDoesNotEqual),
	string(BasicTypeEnumEquals),
}

// CampaignFilterEnum
type CampaignFilterEnum string

const (
	CampaignFilterEnumID     CampaignFilterEnum = "id"     //
	CampaignFilterEnumOwner  CampaignFilterEnum = "owner"  //
	CampaignFilterEnumStatus CampaignFilterEnum = "status" //
)

// All CampaignFilterEnum as []string
var AllCampaignFilterEnum = []string{
	string(CampaignFilterEnumID),
	string(CampaignFilterEnumOwner),
	string(CampaignFilterEnumStatus),
}

// CampaignReminderTypeEnum
type CampaignReminderTypeEnum string

const (
	CampaignReminderTypeEnumEmail CampaignReminderTypeEnum = "email" //
	CampaignReminderTypeEnumSlack CampaignReminderTypeEnum = "slack" //
)

// All CampaignReminderTypeEnum as []string
var AllCampaignReminderTypeEnum = []string{
	string(CampaignReminderTypeEnumEmail),
	string(CampaignReminderTypeEnumSlack),
}

// CampaignServiceStatusEnum
type CampaignServiceStatusEnum string

const (
	CampaignServiceStatusEnumFailing CampaignServiceStatusEnum = "failing" //
	CampaignServiceStatusEnumPassing CampaignServiceStatusEnum = "passing" //
)

// All CampaignServiceStatusEnum as []string
var AllCampaignServiceStatusEnum = []string{
	string(CampaignServiceStatusEnumFailing),
	string(CampaignServiceStatusEnumPassing),
}

// CampaignSortEnum
type CampaignSortEnum string

const (
	CampaignSortEnumChecksPassingAsc     CampaignSortEnum = "checks_passing_ASC"     //
	CampaignSortEnumChecksPassingDesc    CampaignSortEnum = "checks_passing_DESC"    //
	CampaignSortEnumEndedDateAsc         CampaignSortEnum = "ended_date_ASC"         //
	CampaignSortEnumEndedDateDesc        CampaignSortEnum = "ended_date_DESC"        //
	CampaignSortEnumFilterAsc            CampaignSortEnum = "filter_ASC"             //
	CampaignSortEnumFilterDesc           CampaignSortEnum = "filter_DESC"            //
	CampaignSortEnumNameAsc              CampaignSortEnum = "name_ASC"               //
	CampaignSortEnumNameDesc             CampaignSortEnum = "name_DESC"              //
	CampaignSortEnumOwnerAsc             CampaignSortEnum = "owner_ASC"              //
	CampaignSortEnumOwnerDesc            CampaignSortEnum = "owner_DESC"             //
	CampaignSortEnumServicesCompleteAsc  CampaignSortEnum = "services_complete_ASC"  //
	CampaignSortEnumServicesCompleteDesc CampaignSortEnum = "services_complete_DESC" //
	CampaignSortEnumStartDateAsc         CampaignSortEnum = "start_date_ASC"         //
	CampaignSortEnumStartDateDesc        CampaignSortEnum = "start_date_DESC"        //
	CampaignSortEnumStatusAsc            CampaignSortEnum = "status_ASC"             //
	CampaignSortEnumStatusDesc           CampaignSortEnum = "status_DESC"            //
	CampaignSortEnumTargetDateAsc        CampaignSortEnum = "target_date_ASC"        //
	CampaignSortEnumTargetDateDesc       CampaignSortEnum = "target_date_DESC"       //
)

// All CampaignSortEnum as []string
var AllCampaignSortEnum = []string{
	string(CampaignSortEnumChecksPassingAsc),
	string(CampaignSortEnumChecksPassingDesc),
	string(CampaignSortEnumEndedDateAsc),
	string(CampaignSortEnumEndedDateDesc),
	string(CampaignSortEnumFilterAsc),
	string(CampaignSortEnumFilterDesc),
	string(CampaignSortEnumNameAsc),
	string(CampaignSortEnumNameDesc),
	string(CampaignSortEnumOwnerAsc),
	string(CampaignSortEnumOwnerDesc),
	string(CampaignSortEnumServicesCompleteAsc),
	string(CampaignSortEnumServicesCompleteDesc),
	string(CampaignSortEnumStartDateAsc),
	string(CampaignSortEnumStartDateDesc),
	string(CampaignSortEnumStatusAsc),
	string(CampaignSortEnumStatusDesc),
	string(CampaignSortEnumTargetDateAsc),
	string(CampaignSortEnumTargetDateDesc),
}

// CampaignStatusEnum
type CampaignStatusEnum string

const (
	CampaignStatusEnumDelayed    CampaignStatusEnum = "delayed"     //
	CampaignStatusEnumDraft      CampaignStatusEnum = "draft"       //
	CampaignStatusEnumEnded      CampaignStatusEnum = "ended"       //
	CampaignStatusEnumInProgress CampaignStatusEnum = "in_progress" //
	CampaignStatusEnumScheduled  CampaignStatusEnum = "scheduled"   //
)

// All CampaignStatusEnum as []string
var AllCampaignStatusEnum = []string{
	string(CampaignStatusEnumDelayed),
	string(CampaignStatusEnumDraft),
	string(CampaignStatusEnumEnded),
	string(CampaignStatusEnumInProgress),
	string(CampaignStatusEnumScheduled),
}

// CheckStatus represents the evaluation status of the check.
type CheckStatus string

const (
	CheckStatusFailed  CheckStatus = "failed"  // The check evaluated to a falsy value based on some conditions.
	CheckStatusPassed  CheckStatus = "passed"  // The check evaluated to a truthy value based on some conditions.
	CheckStatusPending CheckStatus = "pending" // The check has not been evaluated yet..
)

// All CheckStatus as []string
var AllCheckStatus = []string{
	string(CheckStatusFailed),
	string(CheckStatusPassed),
	string(CheckStatusPending),
}

// CheckType represents the type of check.
type CheckType string

const (
	CheckTypeAlertSourceUsage    CheckType = "alert_source_usage"    // Verifies that the service has an alert source of a particular type or name.
	CheckTypeCustom              CheckType = "custom"                // Allows for the creation of programmatic checks that use an API to mark the status as passing or failing.
	CheckTypeGeneric             CheckType = "generic"               // Requires a generic integration api call to complete a series of checks for multiple services.
	CheckTypeGitBranchProtection CheckType = "git_branch_protection" // Verifies that all the repositories on the service have branch protection enabled.
	CheckTypeHasDocumentation    CheckType = "has_documentation"     // Verifies that the service has visible documentation of a particular type and subtype.
	CheckTypeHasOwner            CheckType = "has_owner"             // Verifies that the service has an owner defined.
	CheckTypeHasRecentDeploy     CheckType = "has_recent_deploy"     // Verified that the services has received a deploy within a specified number of days.
	CheckTypeHasRepository       CheckType = "has_repository"        // Verifies that the service has a repository integrated.
	CheckTypeHasServiceConfig    CheckType = "has_service_config"    // Verifies that the service is maintained though the use of an opslevel.yml service config.
	CheckTypeManual              CheckType = "manual"                // Requires a service owner to manually complete a check for the service.
	CheckTypePayload             CheckType = "payload"               // Requires a payload integration api call to complete a check for the service.
	CheckTypeRepoFile            CheckType = "repo_file"             // Verifies that the service's repository contains a file with a certain path. (Optional: Can also be used to check for specific file contents).
	CheckTypeRepoGrep            CheckType = "repo_grep"             // Runs a comprehensive search across the service's repository with advanced search parameters.
	CheckTypeRepoSearch          CheckType = "repo_search"           // Searches the service's repository and verifies if any file matches the given contents.
	CheckTypeServiceDependency   CheckType = "service_dependency"    // Verifies that the service has either a dependent or dependency.
	CheckTypeServiceProperty     CheckType = "service_property"      // Verifies that a service property is set or matches a specified format.
	CheckTypeTagDefined          CheckType = "tag_defined"           // Verifies that the service has the specified tag defined.
	CheckTypeToolUsage           CheckType = "tool_usage"            // Verifies that the service is using a tool of a particular category or name.
)

// All CheckType as []string
var AllCheckType = []string{
	string(CheckTypeAlertSourceUsage),
	string(CheckTypeCustom),
	string(CheckTypeGeneric),
	string(CheckTypeGitBranchProtection),
	string(CheckTypeHasDocumentation),
	string(CheckTypeHasOwner),
	string(CheckTypeHasRecentDeploy),
	string(CheckTypeHasRepository),
	string(CheckTypeHasServiceConfig),
	string(CheckTypeManual),
	string(CheckTypePayload),
	string(CheckTypeRepoFile),
	string(CheckTypeRepoGrep),
	string(CheckTypeRepoSearch),
	string(CheckTypeServiceDependency),
	string(CheckTypeServiceProperty),
	string(CheckTypeTagDefined),
	string(CheckTypeToolUsage),
}

// ConnectiveEnum represents the logical operator to be used in conjunction with multiple filters (requires filters to be supplied).
type ConnectiveEnum string

const (
	ConnectiveEnumAnd ConnectiveEnum = "and" // Used to ensure **all** filters match for a given resource.
	ConnectiveEnumOr  ConnectiveEnum = "or"  // Used to ensure **any** filters match for a given resource.
)

// All ConnectiveEnum as []string
var AllConnectiveEnum = []string{
	string(ConnectiveEnumAnd),
	string(ConnectiveEnumOr),
}

// ContactType represents the method of contact.
type ContactType string

const (
	ContactTypeEmail       ContactType = "email"        // An email contact method.
	ContactTypeGitHub      ContactType = "github"       //
	ContactTypeSlack       ContactType = "slack"        // A Slack channel contact method.
	ContactTypeSlackHandle ContactType = "slack_handle" // A Slack handle contact method.
	ContactTypeWeb         ContactType = "web"          // A website contact method.
)

// All ContactType as []string
var AllContactType = []string{
	string(ContactTypeEmail),
	string(ContactTypeGitHub),
	string(ContactTypeSlack),
	string(ContactTypeSlackHandle),
	string(ContactTypeWeb),
}

// TODO: This appears to be duplicative of the above and i'm not sure why we need it
// ContactType represents the method of contact.
type ServiceOwnershipCheckContactType string

const (
	ServiceOwnershipCheckContactTypeAny         ServiceOwnershipCheckContactType = "any"          // Any contact method.
	ServiceOwnershipCheckContactTypeSlack       ServiceOwnershipCheckContactType = "slack"        // A Slack channel contact method.
	ServiceOwnershipCheckContactTypeSlackHandle ServiceOwnershipCheckContactType = "slack_handle" // A Slack handle contact method.
	ServiceOwnershipCheckContactTypeEmail       ServiceOwnershipCheckContactType = "email"        // An email contact method.
	ServiceOwnershipCheckContactTypeWeb         ServiceOwnershipCheckContactType = "web"          // A website contact method.
)

// All ServiceOwnershipContactType as []string
var AllServiceOwnershipCheckContactType = []string {
	string(ServiceOwnershipCheckContactTypeAny),
	string(ServiceOwnershipCheckContactTypeSlack),
	string(ServiceOwnershipCheckContactTypeSlackHandle),
	string(ServiceOwnershipCheckContactTypeEmail),
	string(ServiceOwnershipCheckContactTypeWeb),
}

// CustomActionsEntityTypeEnum
type CustomActionsEntityTypeEnum string

const (
	CustomActionsEntityTypeEnumGlobal  CustomActionsEntityTypeEnum = "GLOBAL"  //
	CustomActionsEntityTypeEnumService CustomActionsEntityTypeEnum = "SERVICE" //
)

// All CustomActionsEntityTypeEnum as []string
var AllCustomActionsEntityTypeEnum = []string{
	string(CustomActionsEntityTypeEnumGlobal),
	string(CustomActionsEntityTypeEnumService),
}

// CustomActionsHttpMethodEnum An HTTP request method
type CustomActionsHttpMethodEnum string

const (
	CustomActionsHttpMethodEnumDelete CustomActionsHttpMethodEnum = "DELETE" // An HTTP DELETE request
	CustomActionsHttpMethodEnumGet    CustomActionsHttpMethodEnum = "GET"    //
	CustomActionsHttpMethodEnumPatch  CustomActionsHttpMethodEnum = "PATCH"  //
	CustomActionsHttpMethodEnumPost   CustomActionsHttpMethodEnum = "POST"   // An HTTP POST request
	CustomActionsHttpMethodEnumPut    CustomActionsHttpMethodEnum = "PUT"    // An HTTP PUT request
)

// All CustomActionsHttpMethodEnum as []string
var AllCustomActionsHttpMethodEnum = []string{
	string(CustomActionsHttpMethodEnumDelete),
	string(CustomActionsHttpMethodEnumGet),
	string(CustomActionsHttpMethodEnumPatch),
	string(CustomActionsHttpMethodEnumPost),
	string(CustomActionsHttpMethodEnumPut),
}

// CustomActionsTriggerDefinitionAccessControlEnum Who can see and use the trigger definition
type CustomActionsTriggerDefinitionAccessControlEnum string

const (
	CustomActionsTriggerDefinitionAccessControlEnumAdmins        CustomActionsTriggerDefinitionAccessControlEnum = "admins"         // Admin users
	CustomActionsTriggerDefinitionAccessControlEnumEveryone      CustomActionsTriggerDefinitionAccessControlEnum = "everyone"       // All users of OpsLevel
	CustomActionsTriggerDefinitionAccessControlEnumServiceOwners CustomActionsTriggerDefinitionAccessControlEnum = "service_owners" // The owners of a service
)

// All CustomActionsTriggerDefinitionAccessControlEnum as []string
var AllCustomActionsTriggerDefinitionAccessControlEnum = []string{
	string(CustomActionsTriggerDefinitionAccessControlEnumAdmins),
	string(CustomActionsTriggerDefinitionAccessControlEnumEveryone),
	string(CustomActionsTriggerDefinitionAccessControlEnumServiceOwners),
}

// CustomActionsTriggerEventStatusEnum
type CustomActionsTriggerEventStatusEnum string

const (
	CustomActionsTriggerEventStatusEnumFailure CustomActionsTriggerEventStatusEnum = "FAILURE" //
	CustomActionsTriggerEventStatusEnumPending CustomActionsTriggerEventStatusEnum = "PENDING" //
	CustomActionsTriggerEventStatusEnumSuccess CustomActionsTriggerEventStatusEnum = "SUCCESS" //
)

// All CustomActionsTriggerEventStatusEnum as []string
var AllCustomActionsTriggerEventStatusEnum = []string{
	string(CustomActionsTriggerEventStatusEnumFailure),
	string(CustomActionsTriggerEventStatusEnumPending),
	string(CustomActionsTriggerEventStatusEnumSuccess),
}

// FrequencyTimeScale represents the time scale type for the frequency.
type FrequencyTimeScale string

const (
	FrequencyTimeScaleDay   FrequencyTimeScale = "day"   // Consider the time scale of days.
	FrequencyTimeScaleMonth FrequencyTimeScale = "month" // Consider the time scale of months.
	FrequencyTimeScaleWeek  FrequencyTimeScale = "week"  // Consider the time scale of weeks.
	FrequencyTimeScaleYear  FrequencyTimeScale = "year"  // Consider the time scale of years.
)

// All FrequencyTimeScale as []string
var AllFrequencyTimeScale = []string{
	string(FrequencyTimeScaleDay),
	string(FrequencyTimeScaleMonth),
	string(FrequencyTimeScaleWeek),
	string(FrequencyTimeScaleYear),
}

// HasDocumentationSubtypeEnum represents the subtype of the document.
type HasDocumentationSubtypeEnum string

const (
	HasDocumentationSubtypeEnumOpenapi HasDocumentationSubtypeEnum = "openapi" // Document is an OpenAPI document.
)

// All HasDocumentationSubtypeEnum as []string
var AllHasDocumentationSubtypeEnum = []string{
	string(HasDocumentationSubtypeEnumOpenapi),
}

// HasDocumentationTypeEnum represents the type of the document.
type HasDocumentationTypeEnum string

const (
	HasDocumentationTypeEnumAPI  HasDocumentationTypeEnum = "api"  // Document is an API document.
	HasDocumentationTypeEnumTech HasDocumentationTypeEnum = "tech" //
)

// All HasDocumentationTypeEnum as []string
var AllHasDocumentationTypeEnum = []string{
	string(HasDocumentationTypeEnumAPI),
	string(HasDocumentationTypeEnumTech),
}

// PayloadSortEnum
type PayloadSortEnum string

const (
	PayloadSortEnumCreatedAtAsc    PayloadSortEnum = "created_at_ASC"    //
	PayloadSortEnumCreatedAtDesc   PayloadSortEnum = "created_at_DESC"   //
	PayloadSortEnumProcessedAtAsc  PayloadSortEnum = "processed_at_ASC"  //
	PayloadSortEnumProcessedAtDesc PayloadSortEnum = "processed_at_DESC" //
)

// All PayloadSortEnum as []string
var AllPayloadSortEnum = []string{
	string(PayloadSortEnumCreatedAtAsc),
	string(PayloadSortEnumCreatedAtDesc),
	string(PayloadSortEnumProcessedAtAsc),
	string(PayloadSortEnumProcessedAtDesc),
}

// PredicateKeyEnum represents fields that can be used as part of filter for services.
type PredicateKeyEnum string

const (
	PredicateKeyEnumCreationSource PredicateKeyEnum = "creation_source" // Filter by the creation source.
	PredicateKeyEnumFramework      PredicateKeyEnum = "framework"       // Filter by `framework` field.
	PredicateKeyEnumGroupIDs       PredicateKeyEnum = "group_ids"       // Filter by group hierarchy. Will return resources who's owner is in the group ancestry chain.
	PredicateKeyEnumLanguage       PredicateKeyEnum = "language"        // Filter by `language` field.
	PredicateKeyEnumLifecycleIndex PredicateKeyEnum = "lifecycle_index" // Filter by `lifecycle` field.
	PredicateKeyEnumName           PredicateKeyEnum = "name"            // Filter by `name` field.
	PredicateKeyEnumOwnerID        PredicateKeyEnum = "owner_id"        // Filter by `owner` field.
	PredicateKeyEnumProduct        PredicateKeyEnum = "product"         // Filter by `product` field.
	PredicateKeyEnumTags           PredicateKeyEnum = "tags"            // Filter by `tags` field.
	PredicateKeyEnumTierIndex      PredicateKeyEnum = "tier_index"      // Filter by `tier` field.
	PredicateKeyEnumDomainID       PredicateKeyEnum = "domain_id"       // Filter by Domain that includes the System this service is assigned to, if any.
	PredicateKeyEnumSystemID       PredicateKeyEnum = "system_id"       // Filter by System that this service is assigned to, if any.
)

// All PredicateKeyEnum as []string
var AllPredicateKeyEnum = []string{
	string(PredicateKeyEnumCreationSource),
	string(PredicateKeyEnumFramework),
	string(PredicateKeyEnumGroupIDs),
	string(PredicateKeyEnumLanguage),
	string(PredicateKeyEnumLifecycleIndex),
	string(PredicateKeyEnumName),
	string(PredicateKeyEnumOwnerID),
	string(PredicateKeyEnumProduct),
	string(PredicateKeyEnumTags),
	string(PredicateKeyEnumTierIndex),
	string(PredicateKeyEnumDomainID),
	string(PredicateKeyEnumSystemID),
}

// PredicateTypeEnum represents operations that can be used on predicates.
type PredicateTypeEnum string

const (
	PredicateTypeEnumBelongsTo                  PredicateTypeEnum = "belongs_to"                   // Belongs to a group's hierarchy.
	PredicateTypeEnumContains                   PredicateTypeEnum = "contains"                     // Contains a specific value.
	PredicateTypeEnumDoesNotContain             PredicateTypeEnum = "does_not_contain"             // Does not contain a specific value.
	PredicateTypeEnumDoesNotEqual               PredicateTypeEnum = "does_not_equal"               // Does not equal a specific value.
	PredicateTypeEnumDoesNotExist               PredicateTypeEnum = "does_not_exist"               // Specific attribute does not exist.
	PredicateTypeEnumEndsWith                   PredicateTypeEnum = "ends_with"                    // Ends with a specific value.
	PredicateTypeEnumEquals                     PredicateTypeEnum = "equals"                       // Equals a specific value.
	PredicateTypeEnumExists                     PredicateTypeEnum = "exists"                       // Specific attribute exists.
	PredicateTypeEnumGreaterThanOrEqualTo       PredicateTypeEnum = "greater_than_or_equal_to"     // Greater than or equal to a specific value (numeric only).
	PredicateTypeEnumLessThanOrEqualTo          PredicateTypeEnum = "less_than_or_equal_to"        // Less than or equal to a specific value (numeric only).
	PredicateTypeEnumMatchesRegex               PredicateTypeEnum = "matches_regex"                // Matches a value using a regular expression.
	PredicateTypeEnumSatisfiesJqExpression      PredicateTypeEnum = "satisfies_jq_expression"      // Satisfies an expression defined in jq.
	PredicateTypeEnumSatisfiesVersionConstraint PredicateTypeEnum = "satisfies_version_constraint" // Satisfies version constraint (tag value only).
	PredicateTypeEnumStartsWith                 PredicateTypeEnum = "starts_with"                  // Starts with a specific value.
)

// All PredicateTypeEnum as []string
var AllPredicateTypeEnum = []string{
	string(PredicateTypeEnumBelongsTo),
	string(PredicateTypeEnumContains),
	string(PredicateTypeEnumDoesNotContain),
	string(PredicateTypeEnumDoesNotEqual),
	string(PredicateTypeEnumDoesNotExist),
	string(PredicateTypeEnumEndsWith),
	string(PredicateTypeEnumEquals),
	string(PredicateTypeEnumExists),
	string(PredicateTypeEnumGreaterThanOrEqualTo),
	string(PredicateTypeEnumLessThanOrEqualTo),
	string(PredicateTypeEnumMatchesRegex),
	string(PredicateTypeEnumSatisfiesJqExpression),
	string(PredicateTypeEnumSatisfiesVersionConstraint),
	string(PredicateTypeEnumStartsWith),
}

// RepositoryVisibilityEnum
type RepositoryVisibilityEnum string

const (
	RepositoryVisibilityEnumInternal RepositoryVisibilityEnum = "INTERNAL" //
	RepositoryVisibilityEnumPrivate  RepositoryVisibilityEnum = "PRIVATE"  //
	RepositoryVisibilityEnumPublic   RepositoryVisibilityEnum = "PUBLIC"   //
)

// All RepositoryVisibilityEnum as []string
var AllRepositoryVisibilityEnum = []string{
	string(RepositoryVisibilityEnumInternal),
	string(RepositoryVisibilityEnumPrivate),
	string(RepositoryVisibilityEnumPublic),
}

// ResourceDocumentStatusTypeEnum
type ResourceDocumentStatusTypeEnum string

const (
	ResourceDocumentStatusTypeEnumHidden  ResourceDocumentStatusTypeEnum = "hidden"  //
	ResourceDocumentStatusTypeEnumPinned  ResourceDocumentStatusTypeEnum = "pinned"  //
	ResourceDocumentStatusTypeEnumVisible ResourceDocumentStatusTypeEnum = "visible" //
)

// All ResourceDocumentStatusTypeEnum as []string
var AllResourceDocumentStatusTypeEnum = []string{
	string(ResourceDocumentStatusTypeEnumHidden),
	string(ResourceDocumentStatusTypeEnumPinned),
	string(ResourceDocumentStatusTypeEnumVisible),
}

// ServicePropertyTypeEnum represents properties of services that can be validated.
type ServicePropertyTypeEnum string

const (
	ServicePropertyTypeEnumDescription    ServicePropertyTypeEnum = "description"     // The description of a service.
	ServicePropertyTypeEnumFramework      ServicePropertyTypeEnum = "framework"       // The primary software development framework of a service.
	ServicePropertyTypeEnumLanguage       ServicePropertyTypeEnum = "language"        // The primary programming language of a service.
	ServicePropertyTypeEnumLifecycleIndex ServicePropertyTypeEnum = "lifecycle_index" // The index of the lifecycle a service belongs to.
	ServicePropertyTypeEnumName           ServicePropertyTypeEnum = "name"            // The name of a service.
	ServicePropertyTypeEnumNote           ServicePropertyTypeEnum = "note"            //
	ServicePropertyTypeEnumProduct        ServicePropertyTypeEnum = "product"         // The product that is associated with a service.
	ServicePropertyTypeEnumTierIndex      ServicePropertyTypeEnum = "tier_index"      // The index of the tier a service belongs to.
)

// All ServicePropertyTypeEnum as []string
var AllServicePropertyTypeEnum = []string{
	string(ServicePropertyTypeEnumDescription),
	string(ServicePropertyTypeEnumFramework),
	string(ServicePropertyTypeEnumLanguage),
	string(ServicePropertyTypeEnumLifecycleIndex),
	string(ServicePropertyTypeEnumName),
	string(ServicePropertyTypeEnumNote),
	string(ServicePropertyTypeEnumProduct),
	string(ServicePropertyTypeEnumTierIndex),
}

// ServiceSortEnum
type ServiceSortEnum string

const (
	ServiceSortEnumAlertStatusAsc    ServiceSortEnum = "alert_status_ASC"    //
	ServiceSortEnumAlertStatusDesc   ServiceSortEnum = "alert_status_DESC"   //
	ServiceSortEnumChecksPassingAsc  ServiceSortEnum = "checks_passing_ASC"  //
	ServiceSortEnumChecksPassingDesc ServiceSortEnum = "checks_passing_DESC" //
	ServiceSortEnumLastDeployAsc     ServiceSortEnum = "last_deploy_ASC"     //
	ServiceSortEnumLastDeployDesc    ServiceSortEnum = "last_deploy_DESC"    //
	ServiceSortEnumLevelIndexAsc     ServiceSortEnum = "level_index_ASC"     //
	ServiceSortEnumLevelIndexDesc    ServiceSortEnum = "level_index_DESC"    //
	ServiceSortEnumLifecycleAsc      ServiceSortEnum = "lifecycle_ASC"       //
	ServiceSortEnumLifecycleDesc     ServiceSortEnum = "lifecycle_DESC"      //
	ServiceSortEnumNameAsc           ServiceSortEnum = "name_ASC"            //
	ServiceSortEnumNameDesc          ServiceSortEnum = "name_DESC"           //
	ServiceSortEnumOwnerAsc          ServiceSortEnum = "owner_ASC"           //
	ServiceSortEnumOwnerDesc         ServiceSortEnum = "owner_DESC"          //
	ServiceSortEnumProductAsc        ServiceSortEnum = "product_ASC"         //
	ServiceSortEnumProductDesc       ServiceSortEnum = "product_DESC"        //
	ServiceSortEnumServiceStatAsc    ServiceSortEnum = "service_stat_ASC"    //
	ServiceSortEnumServiceStatDesc   ServiceSortEnum = "service_stat_DESC"   //
	ServiceSortEnumTierAsc           ServiceSortEnum = "tier_ASC"            //
	ServiceSortEnumTierDesc          ServiceSortEnum = "tier_DESC"           //
)

// All ServiceSortEnum as []string
var AllServiceSortEnum = []string{
	string(ServiceSortEnumAlertStatusAsc),
	string(ServiceSortEnumAlertStatusDesc),
	string(ServiceSortEnumChecksPassingAsc),
	string(ServiceSortEnumChecksPassingDesc),
	string(ServiceSortEnumLastDeployAsc),
	string(ServiceSortEnumLastDeployDesc),
	string(ServiceSortEnumLevelIndexAsc),
	string(ServiceSortEnumLevelIndexDesc),
	string(ServiceSortEnumLifecycleAsc),
	string(ServiceSortEnumLifecycleDesc),
	string(ServiceSortEnumNameAsc),
	string(ServiceSortEnumNameDesc),
	string(ServiceSortEnumOwnerAsc),
	string(ServiceSortEnumOwnerDesc),
	string(ServiceSortEnumProductAsc),
	string(ServiceSortEnumProductDesc),
	string(ServiceSortEnumServiceStatAsc),
	string(ServiceSortEnumServiceStatDesc),
	string(ServiceSortEnumTierAsc),
	string(ServiceSortEnumTierDesc),
}

// TaggableResource represents possible types to apply tags to.
type TaggableResource string

const (
	TaggableResourceRepository TaggableResource = "Repository" // Used to identify a Repository.
	TaggableResourceService    TaggableResource = "Service"    // Used to identify a Service.
	TaggableResourceTeam       TaggableResource = "Team"       // Used to identify a Team.
)

// All TaggableResource as []string
var AllTaggableResource = []string{
	string(TaggableResourceRepository),
	string(TaggableResourceService),
	string(TaggableResourceTeam),
}

// ToolCategory represents the specific categories that a tool can belong to.
type ToolCategory string

const (
	ToolCategoryAdmin                 ToolCategory = "admin"                  // Tools used for administrative purposes.
	ToolCategoryAPIDocumentation      ToolCategory = "api_documentation"      // Tools used as API documentation for this service.
	ToolCategoryBacklog               ToolCategory = "backlog"                //
	ToolCategoryCode                  ToolCategory = "code"                   // Tools used for source code.
	ToolCategoryContinuousIntegration ToolCategory = "continuous_integration" // Tools used for building/unit testing a service.
	ToolCategoryDeployment            ToolCategory = "deployment"             // Tools used for deploying changes to a service.
	ToolCategoryErrors                ToolCategory = "errors"                 // Tools used for tracking/reporting errors.
	ToolCategoryFeatureFlag           ToolCategory = "feature_flag"           // Tools used for managing feature flags.
	ToolCategoryHealthChecks          ToolCategory = "health_checks"          // Tools used for tracking/reporting the health of a service.
	ToolCategoryIncidents             ToolCategory = "incidents"              // Tools used to surface incidents on a service.
	ToolCategoryIssueTracking         ToolCategory = "issue_tracking"         // Tools used for tracking issues.
	ToolCategoryLogs                  ToolCategory = "logs"                   // Tools used for displaying logs from services.
	ToolCategoryMetrics               ToolCategory = "metrics"                // Tools used for tracking/reporting service metrics.
	ToolCategoryOrchestrator          ToolCategory = "orchestrator"           // Tools used for orchestrating a service.
	ToolCategoryOther                 ToolCategory = "other"                  // Tools that do not fit into the available categories.
	ToolCategoryResiliency            ToolCategory = "resiliency"             // Tools used for testing the resiliency of a service.
	ToolCategoryRunbooks              ToolCategory = "runbooks"               // Tools used for managing runbooks for a service.
	ToolCategorySecurityScans         ToolCategory = "security_scans"         // Tools used for performing security scans.
	ToolCategoryStatusPage            ToolCategory = "status_page"            // Tools used for reporting the status of a service.
	ToolCategoryWiki                  ToolCategory = "wiki"                   // Tools used as a wiki for this service.
)

// All ToolCategory as []string
var AllToolCategory = []string{
	string(ToolCategoryAdmin),
	string(ToolCategoryAPIDocumentation),
	string(ToolCategoryBacklog),
	string(ToolCategoryCode),
	string(ToolCategoryContinuousIntegration),
	string(ToolCategoryDeployment),
	string(ToolCategoryErrors),
	string(ToolCategoryFeatureFlag),
	string(ToolCategoryHealthChecks),
	string(ToolCategoryIncidents),
	string(ToolCategoryIssueTracking),
	string(ToolCategoryLogs),
	string(ToolCategoryMetrics),
	string(ToolCategoryOrchestrator),
	string(ToolCategoryOther),
	string(ToolCategoryResiliency),
	string(ToolCategoryRunbooks),
	string(ToolCategorySecurityScans),
	string(ToolCategoryStatusPage),
	string(ToolCategoryWiki),
}

// UserRole represents a role that can be assigned to a user.
type UserRole string

const (
	UserRoleAdmin UserRole = "admin" // An administrator on the account.
	UserRoleUser  UserRole = "user"  // A regular user on the account.
)

// All UserRole as []string
var AllUserRole = []string{
	string(UserRoleAdmin),
	string(UserRoleUser),
}
