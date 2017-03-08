package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"

// UserCertificateStatusValue represents the possible values for attribute "certificateStatus".
type UserCertificateStatusValue string

const (
	// UserCertificateStatusRenew represents the value RENEW.
	UserCertificateStatusRenew UserCertificateStatusValue = "RENEW"

	// UserCertificateStatusRevoked represents the value REVOKED.
	UserCertificateStatusRevoked UserCertificateStatusValue = "REVOKED"

	// UserCertificateStatusValid represents the value VALID.
	UserCertificateStatusValid UserCertificateStatusValue = "VALID"
)

// UserIdentity represents the Identity of the object
var UserIdentity = elemental.Identity{
	Name:     "user",
	Category: "users",
}

// UsersList represents a list of Users
type UsersList []*User

// ContentIdentity returns the identity of the objects in the list.
func (o UsersList) ContentIdentity() elemental.Identity {
	return UserIdentity
}

// List convert the object to and elemental.IdentifiablesList.
func (o UsersList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// User represents the model of a user
type User struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags"`

	// Certificate provides a certificate for the user
	Certificate string `json:"certificate" bson:"certificate"`

	// CertificateExpirationDate indicates the expiration day for the certificate.
	CertificateExpirationDate time.Time `json:"certificateExpirationDate" bson:"certificateexpirationdate"`

	// CertificateKey provides the key for the user. Only available at create or update time.
	CertificateKey string `json:"certificateKey" bson:"-"`

	// CertificateStatus provides the status of the certificate. Update with RENEW to get a new certificate.
	CertificateStatus UserCertificateStatusValue `json:"certificateStatus" bson:"certificatestatus"`

	// CreatedTime is the time at which the object was created
	CreateTime time.Time `json:"createTime" bson:"createtime"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description"`

	// e-mail address of the user
	Email string `json:"email" bson:"email"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" bson:"namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags"`

	// ParentAuthenticator is an Internal attribute that points to the parent authenticator.
	ParentAuthenticator string `json:"-" bson:"parentauthenticator"`

	// parentID contains the ID of the parent.
	ParentID string `json:"parentID" bson:"parentid"`

	// parentType contains the identiy name of the parent.
	ParentType string `json:"parentType" bson:"parenttype"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected"`

	// OU attribute for the generated certificates
	SubOrganizations []string `json:"subOrganizations" bson:"suborganizations"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime"`

	// CommonName (CN) for the user certificate
	UserName string `json:"userName" bson:"username"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`
}

// NewUser returns a new *User
func NewUser() *User {

	return &User{
		ModelVersion:      1.0,
		AssociatedTags:    []string{},
		CertificateStatus: "VALID",
		NormalizedTags:    []string{},
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

// SetIdentifier sets the value of the object's unique identifier.
func (o *User) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *User) Version() float64 {

	return 1.0
}

func (o *User) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *User) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *User) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreateTime set the given createTime of the receiver
func (o *User) SetCreateTime(createTime time.Time) {
	o.CreateTime = createTime
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

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *User) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *User) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
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

// GetProtected returns the protected of the receiver
func (o *User) GetProtected() bool {
	return o.Protected
}

// SetUpdateTime set the given updateTime of the receiver
func (o *User) SetUpdateTime(updateTime time.Time) {
	o.UpdateTime = updateTime
}

// Validate valides the current information stored into the structure.
func (o *User) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("certificateStatus", string(o.CertificateStatus), []string{"RENEW", "REVOKED", "VALID"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumLength("userName", o.UserName, 64, false); err != nil {
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
func (User) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return UserAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (User) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return UserAttributesMap
}

// UserAttributesMap represents the map of attribute for User.
var UserAttributesMap = map[string]elemental.AttributeSpecification{
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
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
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
	"Certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Certificate provides a certificate for the user`,
		Exposed:        true,
		Format:         "free",
		Name:           "certificate",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"CertificateExpirationDate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `CertificateExpirationDate indicates the expiration day for the certificate.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "certificateExpirationDate",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"CertificateKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `CertificateKey provides the key for the user. Only available at create or update time.`,
		Exposed:        true,
		Format:         "free",
		Name:           "certificateKey",
		ReadOnly:       true,
		Type:           "string",
	},
	"CertificateStatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"RENEW", "REVOKED", "VALID"},
		DefaultValue:   UserCertificateStatusValue("VALID"),
		Description:    `CertificateStatus provides the status of the certificate. Update with RENEW to get a new certificate.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "certificateStatus",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
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
	"Email": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `e-mail address of the user`,
		Exposed:        true,
		Filterable:     true,
		Format:         "email",
		Name:           "email",
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
	"ParentAuthenticator": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ParentAuthenticator is an Internal attribute that points to the parent authenticator.`,
		Filterable:     true,
		Format:         "free",
		Name:           "parentAuthenticator",
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `parentID contains the ID of the parent.`,
		Exposed:        true,
		Filterable:     true,
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
		Description:    `parentType contains the identiy name of the parent.`,
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
	"SubOrganizations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `OU attribute for the generated certificates`,
		Exposed:        true,
		Filterable:     true,
		Name:           "subOrganizations",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"UserName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `CommonName (CN) for the user certificate`,
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
