package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/shared/golang/gaiaconstants"

// ServerCertificateStatusValue represents the possible values for attribute "certificateStatus".
type ServerCertificateStatusValue string

const (
	// ServerCertificateStatusRenew represents the value RENEW.
	ServerCertificateStatusRenew ServerCertificateStatusValue = "RENEW"

	// ServerCertificateStatusRevoked represents the value REVOKED.
	ServerCertificateStatusRevoked ServerCertificateStatusValue = "REVOKED"

	// ServerCertificateStatusValid represents the value VALID.
	ServerCertificateStatusValid ServerCertificateStatusValue = "VALID"
)

// ServerOperationalStatusValue represents the possible values for attribute "operationalStatus".
type ServerOperationalStatusValue string

const (
	// ServerOperationalStatusConnected represents the value Connected.
	ServerOperationalStatusConnected ServerOperationalStatusValue = "Connected"

	// ServerOperationalStatusDisconnected represents the value Disconnected.
	ServerOperationalStatusDisconnected ServerOperationalStatusValue = "Disconnected"

	// ServerOperationalStatusInitialized represents the value Initialized.
	ServerOperationalStatusInitialized ServerOperationalStatusValue = "Initialized"

	// ServerOperationalStatusUnknown represents the value Unknown.
	ServerOperationalStatusUnknown ServerOperationalStatusValue = "Unknown"
)

// ServerIdentity represents the Identity of the object
var ServerIdentity = elemental.Identity{
	Name:     "server",
	Category: "servers",
}

// ServersList represents a list of Servers
type ServersList []*Server

// ContentIdentity returns the identity of the objects in the list.
func (o ServersList) ContentIdentity() elemental.Identity {
	return ServerIdentity
}

// List convert the object to and elemental.IdentifiablesList.
func (o ServersList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// Server represents the model of a server
type Server struct {
	// FQDN contains the fqdn of the server.
	FQDN string `json:"FQDN" bson:"fqdn"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags"`

	// Certificate is the certificate of the server
	Certificate string `json:"certificate" bson:"certificate"`

	// CertificateExpirationDate is the expiration date of the certificate.
	CertificateExpirationDate time.Time `json:"certificateExpirationDate" bson:"certificateexpirationdate"`

	// CertificateKey is the secret key of the server. Returned only when a server is created or the certificate is updated.
	CertificateKey string `json:"certificateKey" bson:"-"`

	// CertificateStatus indicates if the certificate is valid.
	CertificateStatus ServerCertificateStatusValue `json:"certificateStatus" bson:"certificatestatus"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt" bson:"createdat"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description"`

	// LastSyncTime holds the last heart beat time.
	LastSyncTime time.Time `json:"lastSyncTime" bson:"lastsynctime"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" bson:"namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags"`

	// OperationalStatus tells the status of the server
	OperationalStatus ServerOperationalStatusValue `json:"operationalStatus" bson:"-"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected"`

	// Status of an entity
	Status gaiaconstants.EntityStatus `json:"status" bson:"status"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedat"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`
}

// NewServer returns a new *Server
func NewServer() *Server {

	return &Server{
		ModelVersion:      1.0,
		AssociatedTags:    []string{},
		CertificateStatus: "VALID",
		NormalizedTags:    []string{},
		OperationalStatus: "Initialized",
		Status:            gaiaconstants.Active,
	}
}

// Identity returns the Identity of the object.
func (o *Server) Identity() elemental.Identity {

	return ServerIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Server) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Server) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *Server) Version() float64 {

	return 1.0
}

func (o *Server) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *Server) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *Server) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *Server) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetName returns the name of the receiver
func (o *Server) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *Server) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *Server) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *Server) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *Server) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *Server) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetProtected returns the protected of the receiver
func (o *Server) GetProtected() bool {
	return o.Protected
}

// GetStatus returns the status of the receiver
func (o *Server) GetStatus() gaiaconstants.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *Server) SetStatus(status gaiaconstants.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *Server) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *Server) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("FQDN", o.FQDN); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("FQDN", o.FQDN); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("certificateStatus", string(o.CertificateStatus), []string{"RENEW", "REVOKED", "VALID"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredTime("lastSyncTime", o.LastSyncTime); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredTime("lastSyncTime", o.LastSyncTime); err != nil {
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
func (Server) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return ServerAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (Server) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ServerAttributesMap
}

// ServerAttributesMap represents the map of attribute for Server.
var ServerAttributesMap = map[string]elemental.AttributeSpecification{
	"FQDN": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `FQDN contains the fqdn of the server.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "FQDN",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
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
		Description:    `Certificate is the certificate of the server `,
		Exposed:        true,
		Format:         "free",
		Name:           "certificate",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"CertificateExpirationDate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `CertificateExpirationDate is the expiration date of the certificate.`,
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
		Description:    `CertificateKey is the secret key of the server. Returned only when a server is created or the certificate is updated.`,
		Exposed:        true,
		Format:         "free",
		Name:           "certificateKey",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"CertificateStatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"RENEW", "REVOKED", "VALID"},
		DefaultValue:   ServerCertificateStatusValue("VALID"),
		Description:    `CertificateStatus indicates if the certificate is valid.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "certificateStatus",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
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
	"LastSyncTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LastSyncTime holds the last heart beat time.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "lastSyncTime",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "time",
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
	"OperationalStatus": elemental.AttributeSpecification{
		AllowedChoices: []string{"Connected", "Disconnected", "Initialized", "Unknown"},
		Autogenerated:  true,
		DefaultValue:   ServerOperationalStatusValue("Initialized"),
		Description:    `OperationalStatus tells the status of the server`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operationalStatus",
		ReadOnly:       true,
		Transient:      true,
		Type:           "enum",
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
