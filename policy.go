package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	// PolicyAttributeNameID represents the attribute ID.
	PolicyAttributeNameID elemental.AttributeSpecificationNameKey = "ID"

	// PolicyAttributeNameAction represents the attribute action.
	PolicyAttributeNameAction elemental.AttributeSpecificationNameKey = "Action"

	// PolicyAttributeNameAllObjectTags represents the attribute allObjectTags.
	PolicyAttributeNameAllObjectTags elemental.AttributeSpecificationNameKey = "AllObjectTags"

	// PolicyAttributeNameAllSubjectTags represents the attribute allSubjectTags.
	PolicyAttributeNameAllSubjectTags elemental.AttributeSpecificationNameKey = "AllSubjectTags"

	// PolicyAttributeNameAnnotation represents the attribute annotation.
	PolicyAttributeNameAnnotation elemental.AttributeSpecificationNameKey = "Annotation"

	// PolicyAttributeNameAssociatedTags represents the attribute associatedTags.
	PolicyAttributeNameAssociatedTags elemental.AttributeSpecificationNameKey = "AssociatedTags"

	// PolicyAttributeNameCreatedAt represents the attribute createdAt.
	PolicyAttributeNameCreatedAt elemental.AttributeSpecificationNameKey = "CreatedAt"

	// PolicyAttributeNameDeleted represents the attribute deleted.
	PolicyAttributeNameDeleted elemental.AttributeSpecificationNameKey = "Deleted"

	// PolicyAttributeNameDescription represents the attribute description.
	PolicyAttributeNameDescription elemental.AttributeSpecificationNameKey = "Description"

	// PolicyAttributeNameName represents the attribute name.
	PolicyAttributeNameName elemental.AttributeSpecificationNameKey = "Name"

	// PolicyAttributeNameNamespace represents the attribute namespace.
	PolicyAttributeNameNamespace elemental.AttributeSpecificationNameKey = "Namespace"

	// PolicyAttributeNameObject represents the attribute object.
	PolicyAttributeNameObject elemental.AttributeSpecificationNameKey = "Object"

	// PolicyAttributeNameParentID represents the attribute parentID.
	PolicyAttributeNameParentID elemental.AttributeSpecificationNameKey = "ParentID"

	// PolicyAttributeNameParentType represents the attribute parentType.
	PolicyAttributeNameParentType elemental.AttributeSpecificationNameKey = "ParentType"

	// PolicyAttributeNameRelation represents the attribute relation.
	PolicyAttributeNameRelation elemental.AttributeSpecificationNameKey = "Relation"

	// PolicyAttributeNameStatus represents the attribute status.
	PolicyAttributeNameStatus elemental.AttributeSpecificationNameKey = "Status"

	// PolicyAttributeNameSubject represents the attribute subject.
	PolicyAttributeNameSubject elemental.AttributeSpecificationNameKey = "Subject"

	// PolicyAttributeNameType represents the attribute type.
	PolicyAttributeNameType elemental.AttributeSpecificationNameKey = "Type"

	// PolicyAttributeNameUpdatedAt represents the attribute updatedAt.
	PolicyAttributeNameUpdatedAt elemental.AttributeSpecificationNameKey = "UpdatedAt"
)

// PolicyIdentity represents the Identity of the object
var PolicyIdentity = elemental.Identity{
	Name:     "policy",
	Category: "policies",
}

// PoliciesList represents a list of Policies
type PoliciesList []*Policy

// Policy represents the model of a policy
type Policy struct {
	// ID is the identifier of the object.
	ID string `json:"ID,omitempty" cql:"id,primarykey,omitempty"`

	// Action defines set of actions that must be enforced when a dependency is met.
	Action map[string]map[string]string `json:"action,omitempty" cql:"action,omitempty"`

	// This is a set of all object tags for matching in the DB
	AllObjectTags []string `json:"-" cql:"allobjecttags,omitempty"`

	// This is a set of all subject tags for matching in the DB
	AllSubjectTags []string `json:"-" cql:"allsubjecttags,omitempty"`

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

	// Object represents set of entities that another entity depends on. As subjects, objects are identified as logical operations on tags when a policy is defined.
	Object [][]string `json:"object,omitempty" cql:"object,omitempty"`

	// ParentID is the ID of the parent, if any,
	ParentID string `json:"parentID,omitempty" cql:"parentid,omitempty"`

	// ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.
	ParentType string `json:"parentType,omitempty" cql:"parenttype,omitempty"`

	// Relation describes the required operation to be performed between subjects and objects
	Relation []string `json:"relation,omitempty" cql:"relation,omitempty"`

	// Status of an entity
	Status enum.EntityStatus `json:"status,omitempty" cql:"status,omitempty"`

	// Subject represent sets of entities that will have a dependency other entities. Subjects are defined as logical operations on tags. Logical operations can includes AND/OR
	Subject [][]string `json:"subject,omitempty" cql:"subject,omitempty"`

	// Type of the policy
	Type enum.PolicyType `json:"type,omitempty" cql:"type,primarykey,omitempty"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt,omitempty" cql:"updatedat,omitempty"`
}

// NewPolicy returns a new *Policy
func NewPolicy() *Policy {

	return &Policy{
		Status: enum.Active,
	}
}

// Identity returns the Identity of the object.
func (o *Policy) Identity() elemental.Identity {

	return PolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Policy) Identifier() string {

	return o.ID
}

func (o *Policy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Policy) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *Policy) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *Policy) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *Policy) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *Policy) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *Policy) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetName returns the name of the receiver
func (o *Policy) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *Policy) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *Policy) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *Policy) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetParentID returns the parentID of the receiver
func (o *Policy) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *Policy) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *Policy) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *Policy) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetStatus returns the status of the receiver
func (o *Policy) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *Policy) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *Policy) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *Policy) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Candidate", "Disabled"}); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"APIAuthorization", "ExtendTags", "File", "NamespaceMapping", "Network", "Statistics", "Syscall"}); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o Policy) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return PolicyAttributesMap[name]
}

// PolicyAttributesMap represents the map of attribute for Policy.
var PolicyAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	PolicyAttributeNameID: elemental.AttributeSpecification{
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
	PolicyAttributeNameAction: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Stored:         true,
		SubType:        "actions_list",
		Type:           "external",
	},
	PolicyAttributeNameAllObjectTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Name:           "allObjectTags",
		Required:       true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	PolicyAttributeNameAllSubjectTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Name:           "allSubjectTags",
		Required:       true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	PolicyAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	PolicyAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	PolicyAttributeNameCreatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	PolicyAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Filterable:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	PolicyAttributeNameDescription: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	PolicyAttributeNameName: elemental.AttributeSpecification{
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
	PolicyAttributeNameNamespace: elemental.AttributeSpecification{
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
	PolicyAttributeNameObject: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "object",
		Stored:         true,
		SubType:        "policies_list",
		Type:           "external",
	},
	PolicyAttributeNameParentID: elemental.AttributeSpecification{
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
	PolicyAttributeNameParentType: elemental.AttributeSpecification{
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
	PolicyAttributeNameRelation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "relation",
		Required:       true,
		Stored:         true,
		SubType:        "relations_list",
		Type:           "external",
	},
	PolicyAttributeNameStatus: elemental.AttributeSpecification{
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
	PolicyAttributeNameSubject: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "subject",
		Required:       true,
		Stored:         true,
		SubType:        "policies_list",
		Type:           "external",
	},
	PolicyAttributeNameType: elemental.AttributeSpecification{
		AllowedChoices: []string{"APIAuthorization", "ExtendTags", "File", "NamespaceMapping", "Network", "Statistics", "Syscall"},
		Exposed:        true,
		Format:         "free",
		Name:           "type",
		PrimaryKey:     true,
		Required:       true,
		Stored:         true,
		SubType:        "policytype_enum",
		Type:           "external",
	},
	PolicyAttributeNameUpdatedAt: elemental.AttributeSpecification{
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
