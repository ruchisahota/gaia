package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	// NamespaceAttributeNameID represents the attribute ID.
	NamespaceAttributeNameID elemental.AttributeSpecificationNameKey = "namespace/ID"

	// NamespaceAttributeNameAnnotation represents the attribute annotation.
	NamespaceAttributeNameAnnotation elemental.AttributeSpecificationNameKey = "namespace/annotation"

	// NamespaceAttributeNameAssociatedTags represents the attribute associatedTags.
	NamespaceAttributeNameAssociatedTags elemental.AttributeSpecificationNameKey = "namespace/associatedTags"

	// NamespaceAttributeNameAuthenticator represents the attribute authenticator.
	NamespaceAttributeNameAuthenticator elemental.AttributeSpecificationNameKey = "namespace/authenticator"

	// NamespaceAttributeNameCreatedAt represents the attribute createdAt.
	NamespaceAttributeNameCreatedAt elemental.AttributeSpecificationNameKey = "namespace/createdAt"

	// NamespaceAttributeNameDeleted represents the attribute deleted.
	NamespaceAttributeNameDeleted elemental.AttributeSpecificationNameKey = "namespace/deleted"

	// NamespaceAttributeNameDescription represents the attribute description.
	NamespaceAttributeNameDescription elemental.AttributeSpecificationNameKey = "namespace/description"

	// NamespaceAttributeNameName represents the attribute name.
	NamespaceAttributeNameName elemental.AttributeSpecificationNameKey = "namespace/name"

	// NamespaceAttributeNameNamespace represents the attribute namespace.
	NamespaceAttributeNameNamespace elemental.AttributeSpecificationNameKey = "namespace/namespace"

	// NamespaceAttributeNameParentID represents the attribute parentID.
	NamespaceAttributeNameParentID elemental.AttributeSpecificationNameKey = "namespace/parentID"

	// NamespaceAttributeNameParentType represents the attribute parentType.
	NamespaceAttributeNameParentType elemental.AttributeSpecificationNameKey = "namespace/parentType"

	// NamespaceAttributeNameStatus represents the attribute status.
	NamespaceAttributeNameStatus elemental.AttributeSpecificationNameKey = "namespace/status"

	// NamespaceAttributeNameUpdatedAt represents the attribute updatedAt.
	NamespaceAttributeNameUpdatedAt elemental.AttributeSpecificationNameKey = "namespace/updatedAt"
)

// NamespaceIdentity represents the Identity of the object
var NamespaceIdentity = elemental.Identity{
	Name:     "namespace",
	Category: "namespaces",
}

// NamespacesList represents a list of Namespaces
type NamespacesList []*Namespace

// Namespace represents the model of a namespace
type Namespace struct {
	// ID is the identifier of the object.
	ID string `json:"ID,omitempty" cql:"id,omitempty"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation,omitempty" cql:"annotation,omitempty"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags,omitempty" cql:"associatedtags,omitempty"`

	// Authenticator for this namespace
	Authenticator string `json:"authenticator,omitempty" cql:"authenticator,omitempty"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt,omitempty" cql:"createdat,omitempty"`

	// Deleted marks if the entity has been deleted.
	Deleted bool `json:"-" cql:"deleted,omitempty"`

	// Description is the description of the object.
	Description string `json:"description,omitempty" cql:"description,omitempty"`

	// Name is the name of the namespace.
	Name string `json:"name,omitempty" cql:"name,primarykey,omitempty"`

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
}

// NewNamespace returns a new *Namespace
func NewNamespace() *Namespace {

	return &Namespace{
		Status: enum.Active,
	}
}

// Identity returns the Identity of the object.
func (o *Namespace) Identity() elemental.Identity {

	return NamespaceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Namespace) Identifier() string {

	return o.ID
}

func (o *Namespace) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Namespace) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *Namespace) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *Namespace) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *Namespace) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *Namespace) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *Namespace) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetName returns the name of the receiver
func (o *Namespace) GetName() string {
	return o.Name
}

// GetNamespace returns the namespace of the receiver
func (o *Namespace) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *Namespace) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetParentID returns the parentID of the receiver
func (o *Namespace) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *Namespace) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *Namespace) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *Namespace) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetStatus returns the status of the receiver
func (o *Namespace) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *Namespace) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *Namespace) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *Namespace) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Candidate", "Disabled"}); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o Namespace) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return NamespaceAttributesMap[name]
}

// NamespaceAttributesMap represents the map of attribute for Namespace.
var NamespaceAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	NamespaceAttributeNameID: elemental.AttributeSpecification{
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
	NamespaceAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	NamespaceAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	NamespaceAttributeNameAuthenticator: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "authenticator",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	NamespaceAttributeNameCreatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	NamespaceAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Filterable:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	NamespaceAttributeNameDescription: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	NamespaceAttributeNameName: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		PrimaryKey:     true,
		Required:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	NamespaceAttributeNameNamespace: elemental.AttributeSpecification{
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
	NamespaceAttributeNameParentID: elemental.AttributeSpecification{
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
	NamespaceAttributeNameParentType: elemental.AttributeSpecification{
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
	NamespaceAttributeNameStatus: elemental.AttributeSpecification{
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
	NamespaceAttributeNameUpdatedAt: elemental.AttributeSpecification{
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
