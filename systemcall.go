package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	// SystemCallAttributeNameID represents the attribute ID.
	SystemCallAttributeNameID elemental.AttributeSpecificationNameKey = "systemcall/ID"

	// SystemCallAttributeNameAnnotation represents the attribute annotation.
	SystemCallAttributeNameAnnotation elemental.AttributeSpecificationNameKey = "systemcall/annotation"

	// SystemCallAttributeNameAssociatedTags represents the attribute associatedTags.
	SystemCallAttributeNameAssociatedTags elemental.AttributeSpecificationNameKey = "systemcall/associatedTags"

	// SystemCallAttributeNameCreatedAt represents the attribute createdAt.
	SystemCallAttributeNameCreatedAt elemental.AttributeSpecificationNameKey = "systemcall/createdAt"

	// SystemCallAttributeNameDeleted represents the attribute deleted.
	SystemCallAttributeNameDeleted elemental.AttributeSpecificationNameKey = "systemcall/deleted"

	// SystemCallAttributeNameDescription represents the attribute description.
	SystemCallAttributeNameDescription elemental.AttributeSpecificationNameKey = "systemcall/description"

	// SystemCallAttributeNameName represents the attribute name.
	SystemCallAttributeNameName elemental.AttributeSpecificationNameKey = "systemcall/name"

	// SystemCallAttributeNameNamespace represents the attribute namespace.
	SystemCallAttributeNameNamespace elemental.AttributeSpecificationNameKey = "systemcall/namespace"

	// SystemCallAttributeNameStatus represents the attribute status.
	SystemCallAttributeNameStatus elemental.AttributeSpecificationNameKey = "systemcall/status"

	// SystemCallAttributeNameUpdatedAt represents the attribute updatedAt.
	SystemCallAttributeNameUpdatedAt elemental.AttributeSpecificationNameKey = "systemcall/updatedAt"
)

// SystemCallIdentity represents the Identity of the object
var SystemCallIdentity = elemental.Identity{
	Name:     "systemcall",
	Category: "systemcalls",
}

// SystemCallsList represents a list of SystemCalls
type SystemCallsList []*SystemCall

// SystemCall represents the model of a systemcall
type SystemCall struct {
	// ID is the identifier of the object.
	ID string `json:"ID,omitempty" cql:"id,primarykey,omitempty"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation,omitempty" cql:"annotation,omitempty"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags,omitempty" cql:"associatedtags,omitempty"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt,omitempty" cql:"createdat,omitempty"`

	// Deleted marks if the entity has been deleted.
	Deleted bool `json:"-" cql:"deleted,omitempty"`

	// Description is the description of the object.
	Description string `json:"description,omitempty" cql:"description,omitempty"`

	// Name is the name of the entity
	Name string `json:"name,omitempty" cql:"name,omitempty"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace,omitempty" cql:"namespace,primarykey,omitempty"`

	// Status of an entity
	Status enum.EntityStatus `json:"status,omitempty" cql:"status,omitempty"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt,omitempty" cql:"updatedat,omitempty"`
}

// NewSystemCall returns a new *SystemCall
func NewSystemCall() *SystemCall {

	return &SystemCall{
		Status: enum.Active,
	}
}

// Identity returns the Identity of the object.
func (o *SystemCall) Identity() elemental.Identity {

	return SystemCallIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SystemCall) Identifier() string {

	return o.ID
}

func (o *SystemCall) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SystemCall) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *SystemCall) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *SystemCall) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *SystemCall) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *SystemCall) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *SystemCall) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetName returns the name of the receiver
func (o *SystemCall) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *SystemCall) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *SystemCall) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *SystemCall) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetStatus returns the status of the receiver
func (o *SystemCall) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *SystemCall) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *SystemCall) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *SystemCall) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Candidate", "Disabled"}); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o SystemCall) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return SystemCallAttributesMap[name]
}

// SystemCallAttributesMap represents the map of attribute for SystemCall.
var SystemCallAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	SystemCallAttributeNameID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	SystemCallAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	SystemCallAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	SystemCallAttributeNameCreatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	SystemCallAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Filterable:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	SystemCallAttributeNameDescription: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	SystemCallAttributeNameName: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	SystemCallAttributeNameNamespace: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	SystemCallAttributeNameStatus: elemental.AttributeSpecification{
		AllowedChoices: []string{"Active", "Candidate", "Disabled"},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		Stored:         true,
		SubType:        "status_enum",
		Type:           "external",
	},
	SystemCallAttributeNameUpdatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "updatedAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
}
