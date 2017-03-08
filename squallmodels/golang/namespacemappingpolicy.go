package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/shared/golang/gaiaconstants"

// NamespaceMappingPolicyIdentity represents the Identity of the object
var NamespaceMappingPolicyIdentity = elemental.Identity{
	Name:     "namespacemappingpolicy",
	Category: "namespacemappingpolicies",
}

// NamespaceMappingPoliciesList represents a list of NamespaceMappingPolicies
type NamespaceMappingPoliciesList []*NamespaceMappingPolicy

// ContentIdentity returns the identity of the objects in the list.
func (o NamespaceMappingPoliciesList) ContentIdentity() elemental.Identity {
	return NamespaceMappingPolicyIdentity
}

// List convert the object to and elemental.IdentifiablesList.
func (o NamespaceMappingPoliciesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// NamespaceMappingPolicy represents the model of a namespacemappingpolicy
type NamespaceMappingPolicy struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt" bson:"createdat"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description"`

	// mappedNamespace is the mapped namespace
	MappedNamespace string `json:"mappedNamespace" bson:"mappednamespace"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" bson:"namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags"`

	// ParentID is the ID of the parent, if any,
	ParentID string `json:"parentID" bson:"parentid"`

	// ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.
	ParentType string `json:"parentType" bson:"parenttype"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected"`

	// Status of an entity
	Status gaiaconstants.EntityStatus `json:"status" bson:"status"`

	// Subject is the subject.
	Subject [][]string `json:"subject" bson:"-"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedat"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`
}

// NewNamespaceMappingPolicy returns a new *NamespaceMappingPolicy
func NewNamespaceMappingPolicy() *NamespaceMappingPolicy {

	return &NamespaceMappingPolicy{
		ModelVersion:   1.0,
		AssociatedTags: []string{},
		NormalizedTags: []string{},
		Status:         gaiaconstants.Active,
	}
}

// Identity returns the Identity of the object.
func (o *NamespaceMappingPolicy) Identity() elemental.Identity {

	return NamespaceMappingPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NamespaceMappingPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NamespaceMappingPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *NamespaceMappingPolicy) Version() float64 {

	return 1.0
}

func (o *NamespaceMappingPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *NamespaceMappingPolicy) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *NamespaceMappingPolicy) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *NamespaceMappingPolicy) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetName returns the name of the receiver
func (o *NamespaceMappingPolicy) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *NamespaceMappingPolicy) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *NamespaceMappingPolicy) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *NamespaceMappingPolicy) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *NamespaceMappingPolicy) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *NamespaceMappingPolicy) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetParentID returns the parentID of the receiver
func (o *NamespaceMappingPolicy) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *NamespaceMappingPolicy) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *NamespaceMappingPolicy) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *NamespaceMappingPolicy) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetProtected returns the protected of the receiver
func (o *NamespaceMappingPolicy) GetProtected() bool {
	return o.Protected
}

// GetStatus returns the status of the receiver
func (o *NamespaceMappingPolicy) GetStatus() gaiaconstants.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *NamespaceMappingPolicy) SetStatus(status gaiaconstants.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *NamespaceMappingPolicy) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *NamespaceMappingPolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("mappedNamespace", o.MappedNamespace); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("mappedNamespace", o.MappedNamespace); err != nil {
		errors = append(errors, err)
	}

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
func (NamespaceMappingPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return NamespaceMappingPolicyAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (NamespaceMappingPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return NamespaceMappingPolicyAttributesMap
}

// NamespaceMappingPolicyAttributesMap represents the map of attribute for NamespaceMappingPolicy.
var NamespaceMappingPolicyAttributesMap = map[string]elemental.AttributeSpecification{
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
	"CreatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `CreatedAt is the time at which an entity was created`,
		Exposed:        true,
		Name:           "createdAt",
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
	"MappedNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `mappedNamespace is the mapped namespace`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "mappedNamespace",
		Orderable:      true,
		Required:       true,
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
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ParentID is the ID of the parent, if any,`,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Format:         "free",
		Getter:         true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ParentType": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "parentType",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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
		Autogenerated:  true,
		DefaultValue:   gaiaconstants.Active,
		Description:    `Status of an entity`,
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
		Description:    `Subject is the subject.`,
		Exposed:        true,
		Name:           "subject",
		Orderable:      true,
		SubType:        "policies_list",
		Type:           "external",
	},
	"UpdatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `UpdatedAt is the time at which an entity was updated.`,
		Exposed:        true,
		Name:           "updatedAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}
