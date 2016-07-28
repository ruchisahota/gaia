package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	// CertificateAttributeNameID represents the attribute ID.
	CertificateAttributeNameID elemental.AttributeSpecificationNameKey = "certificate/ID"

	// CertificateAttributeNameAnnotation represents the attribute annotation.
	CertificateAttributeNameAnnotation elemental.AttributeSpecificationNameKey = "certificate/annotation"

	// CertificateAttributeNameAssociatedTags represents the attribute associatedTags.
	CertificateAttributeNameAssociatedTags elemental.AttributeSpecificationNameKey = "certificate/associatedTags"

	// CertificateAttributeNameCertificate represents the attribute certificate.
	CertificateAttributeNameCertificate elemental.AttributeSpecificationNameKey = "certificate/certificate"

	// CertificateAttributeNameCreatedAt represents the attribute createdAt.
	CertificateAttributeNameCreatedAt elemental.AttributeSpecificationNameKey = "certificate/createdAt"

	// CertificateAttributeNameDeleted represents the attribute deleted.
	CertificateAttributeNameDeleted elemental.AttributeSpecificationNameKey = "certificate/deleted"

	// CertificateAttributeNameDescription represents the attribute description.
	CertificateAttributeNameDescription elemental.AttributeSpecificationNameKey = "certificate/description"

	// CertificateAttributeNameExpirationDate represents the attribute expirationDate.
	CertificateAttributeNameExpirationDate elemental.AttributeSpecificationNameKey = "certificate/expirationDate"

	// CertificateAttributeNameKey represents the attribute key.
	CertificateAttributeNameKey elemental.AttributeSpecificationNameKey = "certificate/key"

	// CertificateAttributeNameName represents the attribute name.
	CertificateAttributeNameName elemental.AttributeSpecificationNameKey = "certificate/name"

	// CertificateAttributeNameNamespace represents the attribute namespace.
	CertificateAttributeNameNamespace elemental.AttributeSpecificationNameKey = "certificate/namespace"

	// CertificateAttributeNameParentID represents the attribute parentID.
	CertificateAttributeNameParentID elemental.AttributeSpecificationNameKey = "certificate/parentID"

	// CertificateAttributeNameStatus represents the attribute status.
	CertificateAttributeNameStatus elemental.AttributeSpecificationNameKey = "certificate/status"

	// CertificateAttributeNameUpdatedAt represents the attribute updatedAt.
	CertificateAttributeNameUpdatedAt elemental.AttributeSpecificationNameKey = "certificate/updatedAt"
)

// CertificateIdentity represents the Identity of the object
var CertificateIdentity = elemental.Identity{
	Name:     "certificate",
	Category: "certificates",
}

// CertificatesList represents a list of Certificates
type CertificatesList []*Certificate

// Certificate represents the model of a certificate
type Certificate struct {
	// ID is the identifier of the object.
	ID string `json:"ID,omitempty" cql:"id,primarykey,omitempty"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation,omitempty" cql:"annotation,omitempty"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags,omitempty" cql:"associatedtags,omitempty"`

	// Certificate for the user
	Certificate string `json:"certificate,omitempty" cql:"certificate,omitempty"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt,omitempty" cql:"createdat,omitempty"`

	// Deleted marks if the entity has been deleted.
	Deleted bool `json:"-" cql:"deleted,omitempty"`

	// Description is the description of the object.
	Description string `json:"description,omitempty" cql:"description,omitempty"`

	// ExpirationDate is the date that the certificate must expire.
	ExpirationDate time.Time `json:"expirationDate,omitempty" cql:"expirationdate,omitempty"`

	// Key generated for the user
	Key string `json:"key,omitempty" cql:"-"`

	// Name is the name of the entity
	Name string `json:"name,omitempty" cql:"name,omitempty"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace,omitempty" cql:"namespace,primarykey,omitempty"`

	// ParentID is an internal object that refers to the parrent of the certificate.
	ParentID string `json:"parentID,omitempty" cql:"parentid,omitempty"`

	// Status of an entity
	Status enum.EntityStatus `json:"status,omitempty" cql:"status,omitempty"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt,omitempty" cql:"updatedat,omitempty"`
}

// NewCertificate returns a new *Certificate
func NewCertificate() *Certificate {

	return &Certificate{
		Status: enum.Active,
	}
}

// Identity returns the Identity of the object.
func (o *Certificate) Identity() elemental.Identity {

	return CertificateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Certificate) Identifier() string {

	return o.ID
}

func (o *Certificate) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Certificate) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *Certificate) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *Certificate) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *Certificate) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *Certificate) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *Certificate) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetName returns the name of the receiver
func (o *Certificate) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *Certificate) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *Certificate) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *Certificate) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetStatus returns the status of the receiver
func (o *Certificate) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *Certificate) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *Certificate) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *Certificate) Validate() elemental.Errors {

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
func (o Certificate) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return CertificateAttributesMap[name]
}

// CertificateAttributesMap represents the map of attribute for Certificate.
var CertificateAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	CertificateAttributeNameID: elemental.AttributeSpecification{
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
	CertificateAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	CertificateAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	CertificateAttributeNameCertificate: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Name:           "certificate",
		Stored:         true,
		Type:           "string",
	},
	CertificateAttributeNameCreatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	CertificateAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Filterable:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	CertificateAttributeNameDescription: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	CertificateAttributeNameExpirationDate: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Name:           "expirationDate",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "time",
	},
	CertificateAttributeNameKey: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Name:           "key",
		Type:           "string",
	},
	CertificateAttributeNameName: elemental.AttributeSpecification{
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
	CertificateAttributeNameNamespace: elemental.AttributeSpecification{
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
	CertificateAttributeNameParentID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Exposed:        true,
		ForeignKey:     true,
		Format:         "free",
		Name:           "parentID",
		Stored:         true,
		Type:           "string",
	},
	CertificateAttributeNameStatus: elemental.AttributeSpecification{
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
	CertificateAttributeNameUpdatedAt: elemental.AttributeSpecification{
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
