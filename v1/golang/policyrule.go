package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// PolicyRuleIdentity represents the Identity of the object.
var PolicyRuleIdentity = elemental.Identity{
	Name:     "policyrule",
	Category: "policyrules",
	Private:  false,
}

// PolicyRulesList represents a list of PolicyRules
type PolicyRulesList []*PolicyRule

// ContentIdentity returns the identity of the objects in the list.
func (o PolicyRulesList) ContentIdentity() elemental.Identity {

	return PolicyRuleIdentity
}

// Copy returns a pointer to a copy the PolicyRulesList.
func (o PolicyRulesList) Copy() elemental.ContentIdentifiable {

	copy := append(PolicyRulesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PolicyRulesList.
func (o PolicyRulesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(PolicyRulesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PolicyRule))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PolicyRulesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PolicyRulesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o PolicyRulesList) Version() int {

	return 1
}

// PolicyRule represents the model of a policyrule
type PolicyRule struct {
	// Action defines set of actions that must be enforced when a dependency is met.
	Action map[string]map[string]interface{} `json:"action" bson:"-" mapstructure:"action,omitempty"`

	// EnforcerProfiles provides the information about the server profile.
	EnforcerProfiles EnforcerProfilesList `json:"enforcerProfiles" bson:"-" mapstructure:"enforcerProfiles,omitempty"`

	// Policy target networks
	ExternalServices ExternalServicesList `json:"externalServices" bson:"-" mapstructure:"externalServices,omitempty"`

	// Policy target networks
	FilePaths FilePathsList `json:"filePaths" bson:"-" mapstructure:"filePaths,omitempty"`

	// IsolationProfiles are the isolation profiles of the rule.
	IsolationProfiles IsolationProfilesList `json:"isolationProfiles" bson:"isolationprofiles" mapstructure:"isolationProfiles,omitempty"`

	// Policy target networks
	Namespaces NamespacesList `json:"namespaces" bson:"-" mapstructure:"namespaces,omitempty"`

	// List of external services the policy mandate to pass through before reaching the destination.
	PassthroughExternalServices ExternalServicesList `json:"passthroughExternalServices" bson:"-" mapstructure:"passthroughExternalServices,omitempty"`

	// Propagated indicates if the policy is propagated.
	Propagated bool `json:"propagated" bson:"-" mapstructure:"propagated,omitempty"`

	// Relation describes the required operation to be performed between subjects and objects
	Relation []string `json:"relation" bson:"-" mapstructure:"relation,omitempty"`

	// Policy target tags
	TagClauses [][]string `json:"tagClauses" bson:"-" mapstructure:"tagClauses,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewPolicyRule returns a new *PolicyRule
func NewPolicyRule() *PolicyRule {

	return &PolicyRule{
		ModelVersion:                1,
		EnforcerProfiles:            EnforcerProfilesList{},
		ExternalServices:            ExternalServicesList{},
		FilePaths:                   FilePathsList{},
		IsolationProfiles:           IsolationProfilesList{},
		Namespaces:                  NamespacesList{},
		PassthroughExternalServices: ExternalServicesList{},
	}
}

// Identity returns the Identity of the object.
func (o *PolicyRule) Identity() elemental.Identity {

	return PolicyRuleIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PolicyRule) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PolicyRule) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *PolicyRule) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *PolicyRule) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *PolicyRule) Doc() string {
	return nodocString
}

func (o *PolicyRule) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetName returns the Name of the receiver.
func (o *PolicyRule) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *PolicyRule) SetName(name string) {

	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *PolicyRule) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*PolicyRule) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PolicyRuleAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PolicyRuleLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PolicyRule) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PolicyRuleAttributesMap
}

// PolicyRuleAttributesMap represents the map of attribute for PolicyRule.
var PolicyRuleAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
		Unique:         true,
	},
	"Action": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Action",
		Description:    `Action defines set of actions that must be enforced when a dependency is met.`,
		Exposed:        true,
		Name:           "action",
		SubType:        "actions_list",
		Type:           "external",
	},
	"EnforcerProfiles": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerProfiles",
		Description:    `EnforcerProfiles provides the information about the server profile.`,
		Exposed:        true,
		Name:           "enforcerProfiles",
		SubType:        "enforcerprofiles_list",
		Type:           "external",
	},
	"ExternalServices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExternalServices",
		Description:    `Policy target networks `,
		Exposed:        true,
		Name:           "externalServices",
		SubType:        "network_entities",
		Type:           "external",
	},
	"FilePaths": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FilePaths",
		Description:    `Policy target networks `,
		Exposed:        true,
		Name:           "filePaths",
		SubType:        "file_entities",
		Type:           "external",
	},
	"IsolationProfiles": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IsolationProfiles",
		Description:    `IsolationProfiles are the isolation profiles of the rule.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "isolationProfiles",
		Orderable:      true,
		Stored:         true,
		SubType:        "isolation_profile_entities",
		Type:           "external",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Namespaces": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespaces",
		Description:    `Policy target networks `,
		Exposed:        true,
		Name:           "namespaces",
		SubType:        "namespace_entities",
		Type:           "external",
	},
	"PassthroughExternalServices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PassthroughExternalServices",
		Description:    `List of external services the policy mandate to pass through before reaching the destination.`,
		Exposed:        true,
		Name:           "passthroughExternalServices",
		SubType:        "network_entities",
		Type:           "external",
	},
	"Propagated": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagated",
		Description:    `Propagated indicates if the policy is propagated.`,
		Exposed:        true,
		Name:           "propagated",
		Type:           "boolean",
	},
	"Relation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Relation",
		Description:    `Relation describes the required operation to be performed between subjects and objects`,
		Exposed:        true,
		Name:           "relation",
		SubType:        "relations_list",
		Type:           "external",
	},
	"TagClauses": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TagClauses",
		Description:    `Policy target tags`,
		Exposed:        true,
		Name:           "tagClauses",
		SubType:        "target_tags",
		Type:           "external",
	},
}

// PolicyRuleLowerCaseAttributesMap represents the map of attribute for PolicyRule.
var PolicyRuleLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
		Unique:         true,
	},
	"action": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Action",
		Description:    `Action defines set of actions that must be enforced when a dependency is met.`,
		Exposed:        true,
		Name:           "action",
		SubType:        "actions_list",
		Type:           "external",
	},
	"enforcerprofiles": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerProfiles",
		Description:    `EnforcerProfiles provides the information about the server profile.`,
		Exposed:        true,
		Name:           "enforcerProfiles",
		SubType:        "enforcerprofiles_list",
		Type:           "external",
	},
	"externalservices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExternalServices",
		Description:    `Policy target networks `,
		Exposed:        true,
		Name:           "externalServices",
		SubType:        "network_entities",
		Type:           "external",
	},
	"filepaths": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FilePaths",
		Description:    `Policy target networks `,
		Exposed:        true,
		Name:           "filePaths",
		SubType:        "file_entities",
		Type:           "external",
	},
	"isolationprofiles": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IsolationProfiles",
		Description:    `IsolationProfiles are the isolation profiles of the rule.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "isolationProfiles",
		Orderable:      true,
		Stored:         true,
		SubType:        "isolation_profile_entities",
		Type:           "external",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"namespaces": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespaces",
		Description:    `Policy target networks `,
		Exposed:        true,
		Name:           "namespaces",
		SubType:        "namespace_entities",
		Type:           "external",
	},
	"passthroughexternalservices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PassthroughExternalServices",
		Description:    `List of external services the policy mandate to pass through before reaching the destination.`,
		Exposed:        true,
		Name:           "passthroughExternalServices",
		SubType:        "network_entities",
		Type:           "external",
	},
	"propagated": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagated",
		Description:    `Propagated indicates if the policy is propagated.`,
		Exposed:        true,
		Name:           "propagated",
		Type:           "boolean",
	},
	"relation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Relation",
		Description:    `Relation describes the required operation to be performed between subjects and objects`,
		Exposed:        true,
		Name:           "relation",
		SubType:        "relations_list",
		Type:           "external",
	},
	"tagclauses": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TagClauses",
		Description:    `Policy target tags`,
		Exposed:        true,
		Name:           "tagClauses",
		SubType:        "target_tags",
		Type:           "external",
	},
}