package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	// UserAttributeNameID represents the attribute ID.
	UserAttributeNameID elemental.AttributeSpecificationNameKey = "ID"

	// UserAttributeNameAnnotation represents the attribute annotation.
	UserAttributeNameAnnotation elemental.AttributeSpecificationNameKey = "Annotation"

	// UserAttributeNameAssociatedTags represents the attribute associatedTags.
	UserAttributeNameAssociatedTags elemental.AttributeSpecificationNameKey = "AssociatedTags"

	// UserAttributeNameCreatedAt represents the attribute createdAt.
	UserAttributeNameCreatedAt elemental.AttributeSpecificationNameKey = "CreatedAt"

	// UserAttributeNameDeleted represents the attribute deleted.
	UserAttributeNameDeleted elemental.AttributeSpecificationNameKey = "Deleted"

	// UserAttributeNameEmail represents the attribute email.
	UserAttributeNameEmail elemental.AttributeSpecificationNameKey = "Email"

	// UserAttributeNameName represents the attribute name.
	UserAttributeNameName elemental.AttributeSpecificationNameKey = "Name"

	// UserAttributeNameNamespace represents the attribute namespace.
	UserAttributeNameNamespace elemental.AttributeSpecificationNameKey = "Namespace"

	// UserAttributeNameParentAuthenticator represents the attribute parentAuthenticator.
	UserAttributeNameParentAuthenticator elemental.AttributeSpecificationNameKey = "ParentAuthenticator"

	// UserAttributeNameParentID represents the attribute parentID.
	UserAttributeNameParentID elemental.AttributeSpecificationNameKey = "ParentID"

	// UserAttributeNameParentType represents the attribute parentType.
	UserAttributeNameParentType elemental.AttributeSpecificationNameKey = "ParentType"

	// UserAttributeNameStatus represents the attribute status.
	UserAttributeNameStatus elemental.AttributeSpecificationNameKey = "Status"

	// UserAttributeNameSubOrganizations represents the attribute subOrganizations.
	UserAttributeNameSubOrganizations elemental.AttributeSpecificationNameKey = "SubOrganizations"

	// UserAttributeNameUpdatedAt represents the attribute updatedAt.
	UserAttributeNameUpdatedAt elemental.AttributeSpecificationNameKey = "UpdatedAt"

	// UserAttributeNameUserName represents the attribute userName.
	UserAttributeNameUserName elemental.AttributeSpecificationNameKey = "UserName"
)

// UserIdentity represents the Identity of the object
var UserIdentity = elemental.Identity{
	Name:     "user",
	Category: "users",
}

// UsersList represents a list of Users
type UsersList []*User

// User represents the model of a user
type User struct {
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

	// e-mail address of the user
	Email string `json:"email,omitempty" cql:"email,primarykey,omitempty"`

	// Name is the name of the entity
	Name string `json:"name,omitempty" cql:"name,omitempty"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace,omitempty" cql:"namespace,primarykey,omitempty"`

	// ParentAuthenticator is an Internal attribute that points to the parent authenticator.
	ParentAuthenticator string `json:"-" cql:"parentauthenticator,primarykey,omitempty"`

	// ParentID is the ID of the parent, if any,
	ParentID string `json:"parentID,omitempty" cql:"parentid,omitempty"`

	// ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.
	ParentType string `json:"parentType,omitempty" cql:"parenttype,omitempty"`

	// Status of an entity
	Status enum.EntityStatus `json:"status,omitempty" cql:"status,omitempty"`

	// OU attribute for the generated certificates
	SubOrganizations []string `json:"subOrganizations,omitempty" cql:"suborganizations,omitempty"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt,omitempty" cql:"updatedat,omitempty"`

	// CommonName (CN) for the user certificate
	UserName string `json:"userName,omitempty" cql:"username,omitempty"`
}

// NewUser returns a new *User
func NewUser() *User {

	return &User{
		Status: enum.Active,
	}
}

// Identity returns the Identity of the object.
func (o *User) Identity() elemental.Identity {

	return UserIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *User) Identifier() string {

	return o.ID
}

func (o *User) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *User) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *User) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *User) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *User) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *User) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *User) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetName returns the name of the receiver
func (o *User) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *User) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *User) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *User) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetParentID returns the parentID of the receiver
func (o *User) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *User) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *User) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *User) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetStatus returns the status of the receiver
func (o *User) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *User) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *User) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *User) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Candidate", "Disabled"}); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumLength("userName", o.UserName, 64, false); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o User) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return UserAttributesMap[name]
}

// UserAttributesMap represents the map of attribute for User.
var UserAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	UserAttributeNameID: elemental.AttributeSpecification{
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
	UserAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	UserAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	UserAttributeNameCreatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	UserAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Filterable:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	UserAttributeNameEmail: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "email",
		Name:           "email",
		Orderable:      true,
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
	},
	UserAttributeNameName: elemental.AttributeSpecification{
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
	UserAttributeNameNamespace: elemental.AttributeSpecification{
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
	UserAttributeNameParentAuthenticator: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Format:         "free",
		Name:           "parentAuthenticator",
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
	},
	UserAttributeNameParentID: elemental.AttributeSpecification{
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
	UserAttributeNameParentType: elemental.AttributeSpecification{
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
	UserAttributeNameStatus: elemental.AttributeSpecification{
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
	UserAttributeNameSubOrganizations: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Name:           "subOrganizations",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	UserAttributeNameUpdatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "updatedAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	UserAttributeNameUserName: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		MaxLength:      64,
		Name:           "userName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}
