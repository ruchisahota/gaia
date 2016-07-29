package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	// TagAttributeNameID represents the attribute ID.
	TagAttributeNameID elemental.AttributeSpecificationNameKey = "ID"

	// TagAttributeNameAnnotation represents the attribute annotation.
	TagAttributeNameAnnotation elemental.AttributeSpecificationNameKey = "Annotation"

	// TagAttributeNameAssociatedTags represents the attribute associatedTags.
	TagAttributeNameAssociatedTags elemental.AttributeSpecificationNameKey = "AssociatedTags"

	// TagAttributeNameCreatedAt represents the attribute createdAt.
	TagAttributeNameCreatedAt elemental.AttributeSpecificationNameKey = "CreatedAt"

	// TagAttributeNameDeleted represents the attribute deleted.
	TagAttributeNameDeleted elemental.AttributeSpecificationNameKey = "Deleted"

	// TagAttributeNameDescription represents the attribute description.
	TagAttributeNameDescription elemental.AttributeSpecificationNameKey = "Description"

	// TagAttributeNameNamespace represents the attribute namespace.
	TagAttributeNameNamespace elemental.AttributeSpecificationNameKey = "Namespace"

	// TagAttributeNameParentID represents the attribute parentID.
	TagAttributeNameParentID elemental.AttributeSpecificationNameKey = "ParentID"

	// TagAttributeNameParentType represents the attribute parentType.
	TagAttributeNameParentType elemental.AttributeSpecificationNameKey = "ParentType"

	// TagAttributeNameStatus represents the attribute status.
	TagAttributeNameStatus elemental.AttributeSpecificationNameKey = "Status"

	// TagAttributeNameUpdatedAt represents the attribute updatedAt.
	TagAttributeNameUpdatedAt elemental.AttributeSpecificationNameKey = "UpdatedAt"

	// TagAttributeNameValue represents the attribute value.
	TagAttributeNameValue elemental.AttributeSpecificationNameKey = "Value"
)

// TagIdentity represents the Identity of the object
var TagIdentity = elemental.Identity{
	Name:     "tag",
	Category: "tags",
}

// TagsList represents a list of Tags
type TagsList []*Tag

// Tag represents the model of a tag
type Tag struct {
	// ID is the identifier of the object.
	ID string `json:"ID,omitempty" cql:"id,omitempty"`

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

	// Namespace tag attached to an entity
	Namespace string `json:"namespace,omitempty" cql:"namespace,primarykey,omitempty"`

	// ParentID is the ID of the parent, if any,
	ParentID string `json:"parentID,omitempty" cql:"parentid,omitempty"`

	// ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.
	ParentType string `json:"parentType,omitempty" cql:"parenttype,omitempty"`

	// Status of an entity
	Status enum.EntityStatus `json:"status,omitempty" cql:"status,omitempty"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt,omitempty" cql:"updatedat,omitempty"`

	// Value stores the value of the tag in `<key>=<value>` format.
	Value string `json:"value,omitempty" cql:"value,primarykey,omitempty"`
}

// NewTag returns a new *Tag
func NewTag() *Tag {

	return &Tag{
		Status: enum.Active,
	}
}

// Identity returns the Identity of the object.
func (o *Tag) Identity() elemental.Identity {

	return TagIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Tag) Identifier() string {

	return o.ID
}

func (o *Tag) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Tag) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *Tag) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *Tag) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *Tag) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *Tag) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *Tag) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetNamespace returns the namespace of the receiver
func (o *Tag) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *Tag) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetParentID returns the parentID of the receiver
func (o *Tag) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *Tag) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *Tag) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *Tag) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetStatus returns the status of the receiver
func (o *Tag) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *Tag) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *Tag) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *Tag) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Candidate", "Disabled"}); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("value", o.Value, `^[^\s=]+=[^\s=]+$`); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("value", o.Value); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o Tag) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return TagAttributesMap[name]
}

// TagAttributesMap represents the map of attribute for Tag.
var TagAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	TagAttributeNameID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	TagAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	TagAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	TagAttributeNameCreatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	TagAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Filterable:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	TagAttributeNameDescription: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	TagAttributeNameNamespace: elemental.AttributeSpecification{
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
	TagAttributeNameParentID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Format:         "free",
		Name:           "parentID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	TagAttributeNameParentType: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "parentType",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	TagAttributeNameStatus: elemental.AttributeSpecification{
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
	TagAttributeNameUpdatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "updatedAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	TagAttributeNameValue: elemental.AttributeSpecification{
		AllowedChars:   `^[^\s=]+=[^\s=]+$`,
		AllowedChoices: []string{},
		Exposed:        true,
		Format:         "free",
		Name:           "value",
		PrimaryKey:     true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}
