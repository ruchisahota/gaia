package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/shared/golang/gaiaconstants"

// ServerPolicyIdentity represents the Identity of the object
var ServerPolicyIdentity = elemental.Identity{
	Name:     "serverpolicy",
	Category: "serverpolicies",
}

// ServerPoliciesList represents a list of ServerPolicies
type ServerPoliciesList []*ServerPolicy

// ContentIdentity returns the identity of the objects in the list.
func (o ServerPoliciesList) ContentIdentity() elemental.Identity {
	return ServerPolicyIdentity
}

// List convert the object to and elemental.IdentifiablesList.
func (o ServerPoliciesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// ServerPolicy represents the model of a serverpolicy
type ServerPolicy struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags"`

	// CreatedTime is the time at which the object was created
	CreateTime time.Time `json:"createTime" bson:"createtime"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" bson:"namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags"`

	// Object is the list of tags to use to find a server profile.
	Object [][]string `json:"object" bson:"object"`

	// Propagate will propagate the policy to all of its children.
	Propagate bool `json:"propagate" bson:"propagate"`

	// If set to true while the policy is propagating, it won't be visible to children namespace, but still used for policy resolution.
	PropagationHidden bool `json:"propagationHidden" bson:"propagationhidden"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected"`

	// Status represents the status of the object.
	Status gaiaconstants.EntityStatus `json:"status" bson:"status"`

	// Subject is the subject of the policy.
	Subject [][]string `json:"subject" bson:"subject"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`
}

// NewServerPolicy returns a new *ServerPolicy
func NewServerPolicy() *ServerPolicy {

	return &ServerPolicy{
		ModelVersion:   1.0,
		AssociatedTags: []string{},
		NormalizedTags: []string{},
		Status:         gaiaconstants.Active,
	}
}

// Identity returns the Identity of the object.
func (o *ServerPolicy) Identity() elemental.Identity {

	return ServerPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ServerPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ServerPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *ServerPolicy) Version() float64 {

	return 1.0
}

func (o *ServerPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *ServerPolicy) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *ServerPolicy) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreateTime set the given createTime of the receiver
func (o *ServerPolicy) SetCreateTime(createTime time.Time) {
	o.CreateTime = createTime
}

// GetName returns the name of the receiver
func (o *ServerPolicy) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *ServerPolicy) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *ServerPolicy) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *ServerPolicy) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *ServerPolicy) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *ServerPolicy) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the propagate of the receiver
func (o *ServerPolicy) GetPropagate() bool {
	return o.Propagate
}

// SetPropagate set the given propagate of the receiver
func (o *ServerPolicy) SetPropagate(propagate bool) {
	o.Propagate = propagate
}

// GetPropagationHidden returns the propagationHidden of the receiver
func (o *ServerPolicy) GetPropagationHidden() bool {
	return o.PropagationHidden
}

// SetPropagationHidden set the given propagationHidden of the receiver
func (o *ServerPolicy) SetPropagationHidden(propagationHidden bool) {
	o.PropagationHidden = propagationHidden
}

// GetProtected returns the protected of the receiver
func (o *ServerPolicy) GetProtected() bool {
	return o.Protected
}

// GetStatus returns the status of the receiver
func (o *ServerPolicy) GetStatus() gaiaconstants.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *ServerPolicy) SetStatus(status gaiaconstants.EntityStatus) {
	o.Status = status
}

// SetUpdateTime set the given updateTime of the receiver
func (o *ServerPolicy) SetUpdateTime(updateTime time.Time) {
	o.UpdateTime = updateTime
}

// Validate valides the current information stored into the structure.
func (o *ServerPolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredExternal("object", o.Object); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("object", o.Object); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredExternal("subject", o.Subject); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("subject", o.Subject); err != nil {
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
func (ServerPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return ServerPolicyAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (ServerPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ServerPolicyAttributesMap
}

// ServerPolicyAttributesMap represents the map of attribute for ServerPolicy.
var ServerPolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"Annotation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Annotation stores additional information about an entity`,
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AssociatedTags are the list of tags attached to an entity`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `CreatedTime is the time at which the object was created`,
		Exposed:        true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Index:          true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `NormalizedTags contains the list of normalized tags of the entities`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Transient:      true,
		Type:           "external",
	},
	"Object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Object is the list of tags to use to find a server profile.`,
		Exposed:        true,
		Name:           "object",
		Required:       true,
		Stored:         true,
		SubType:        "policies_list",
		Type:           "external",
	},
	"Propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Propagate will propagate the policy to all of its children.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"PropagationHidden": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `If set to true while the policy is propagating, it won't be visible to children namespace, but still used for policy resolution.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "propagationHidden",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		DefaultValue:   gaiaconstants.Active,
		Description:    `Status represents the status of the object.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "status",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		SubType:        "status_enum",
		Type:           "external",
	},
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Subject is the subject of the policy.`,
		Exposed:        true,
		Name:           "subject",
		Required:       true,
		Stored:         true,
		SubType:        "policies_list",
		Type:           "external",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `UpdateTime is the time at which an entity was updated.`,
		Exposed:        true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}
