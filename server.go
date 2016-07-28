package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	// ServerAttributeNameID represents the attribute ID.
	ServerAttributeNameID elemental.AttributeSpecificationNameKey = "server/ID"

	// ServerAttributeNameAddress represents the attribute address.
	ServerAttributeNameAddress elemental.AttributeSpecificationNameKey = "server/address"

	// ServerAttributeNameAnnotation represents the attribute annotation.
	ServerAttributeNameAnnotation elemental.AttributeSpecificationNameKey = "server/annotation"

	// ServerAttributeNameAssociatedTags represents the attribute associatedTags.
	ServerAttributeNameAssociatedTags elemental.AttributeSpecificationNameKey = "server/associatedTags"

	// ServerAttributeNameCreatedAt represents the attribute createdAt.
	ServerAttributeNameCreatedAt elemental.AttributeSpecificationNameKey = "server/createdAt"

	// ServerAttributeNameDeleted represents the attribute deleted.
	ServerAttributeNameDeleted elemental.AttributeSpecificationNameKey = "server/deleted"

	// ServerAttributeNameDescription represents the attribute description.
	ServerAttributeNameDescription elemental.AttributeSpecificationNameKey = "server/description"

	// ServerAttributeNameDomain represents the attribute domain.
	ServerAttributeNameDomain elemental.AttributeSpecificationNameKey = "server/domain"

	// ServerAttributeNameEnvironment represents the attribute environment.
	ServerAttributeNameEnvironment elemental.AttributeSpecificationNameKey = "server/environment"

	// ServerAttributeNameName represents the attribute name.
	ServerAttributeNameName elemental.AttributeSpecificationNameKey = "server/name"

	// ServerAttributeNameNamespace represents the attribute namespace.
	ServerAttributeNameNamespace elemental.AttributeSpecificationNameKey = "server/namespace"

	// ServerAttributeNameOperationalStatus represents the attribute operationalStatus.
	ServerAttributeNameOperationalStatus elemental.AttributeSpecificationNameKey = "server/operationalStatus"

	// ServerAttributeNameStatus represents the attribute status.
	ServerAttributeNameStatus elemental.AttributeSpecificationNameKey = "server/status"

	// ServerAttributeNameUpdatedAt represents the attribute updatedAt.
	ServerAttributeNameUpdatedAt elemental.AttributeSpecificationNameKey = "server/updatedAt"
)

// ServerEnvironmentValue represents the possible values for attribute "environment".
type ServerEnvironmentValue string

const (
	// ServerEnvironmentAws represents the value AWS.
	ServerEnvironmentAws ServerEnvironmentValue = "AWS"

	// ServerEnvironmentGcp represents the value GCP.
	ServerEnvironmentGcp ServerEnvironmentValue = "GCP"

	// ServerEnvironmentPrivate represents the value Private.
	ServerEnvironmentPrivate ServerEnvironmentValue = "Private"
)

// ServerOperationalStatusValue represents the possible values for attribute "operationalStatus".
type ServerOperationalStatusValue string

const (
	// ServerOperationalStatusConnected represents the value CONNECTED.
	ServerOperationalStatusConnected ServerOperationalStatusValue = "CONNECTED"

	// ServerOperationalStatusInitialized represents the value INITIALIZED.
	ServerOperationalStatusInitialized ServerOperationalStatusValue = "INITIALIZED"

	// ServerOperationalStatusUnknown represents the value UNKNOWN.
	ServerOperationalStatusUnknown ServerOperationalStatusValue = "UNKNOWN"
)

// ServerIdentity represents the Identity of the object
var ServerIdentity = elemental.Identity{
	Name:     "server",
	Category: "servers",
}

// ServersList represents a list of Servers
type ServersList []*Server

// Server represents the model of a server
type Server struct {
	// ID is the identifier of the object.
	ID string `json:"ID,omitempty" cql:"id,primarykey,omitempty"`

	// Address provides the current IP address of the server after its initialized.
	Address string `json:"address,omitempty" cql:"address,omitempty"`

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

	// Domain refers to the discovered domain name of the server
	Domain string `json:"domain,omitempty" cql:"domain,omitempty"`

	// Environment describes where the server will be running.
	Environment ServerEnvironmentValue `json:"environment,omitempty" cql:"environment,omitempty"`

	// Name is the name of the entity
	Name string `json:"name,omitempty" cql:"name,omitempty"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace,omitempty" cql:"namespace,primarykey,omitempty"`

	// Operational status of the server
	OperationalStatus ServerOperationalStatusValue `json:"operationalStatus,omitempty" cql:"operationalstatus,omitempty"`

	// Status of an entity
	Status enum.EntityStatus `json:"status,omitempty" cql:"status,omitempty"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt,omitempty" cql:"updatedat,omitempty"`
}

// NewServer returns a new *Server
func NewServer() *Server {

	return &Server{
		OperationalStatus: "UNKNOWN",
		Status:            enum.Active,
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

func (o *Server) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Server) SetIdentifier(ID string) {

	o.ID = ID
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

// GetDeleted returns the deleted of the receiver
func (o *Server) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *Server) SetDeleted(deleted bool) {
	o.Deleted = deleted
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

// GetStatus returns the status of the receiver
func (o *Server) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *Server) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *Server) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *Server) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateStringInList("environment", string(o.Environment), []string{"AWS", "GCP", "Private"}); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("operationalStatus", string(o.OperationalStatus), []string{"CONNECTED", "INITIALIZED", "UNKNOWN"}); err != nil {
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
func (o Server) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return ServerAttributesMap[name]
}

// ServerAttributesMap represents the map of attribute for Server.
var ServerAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	ServerAttributeNameID: elemental.AttributeSpecification{
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
	ServerAttributeNameAddress: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "address",
		Stored:         true,
		Type:           "string",
	},
	ServerAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	ServerAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	ServerAttributeNameCreatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	ServerAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Filterable:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	ServerAttributeNameDescription: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	ServerAttributeNameDomain: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "domain",
		Stored:         true,
		Type:           "string",
	},
	ServerAttributeNameEnvironment: elemental.AttributeSpecification{
		AllowedChoices: []string{"AWS", "GCP", "Private"},
		CreationOnly:   true,
		Exposed:        true,
		Name:           "environment",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	ServerAttributeNameName: elemental.AttributeSpecification{
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
	ServerAttributeNameNamespace: elemental.AttributeSpecification{
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
	ServerAttributeNameOperationalStatus: elemental.AttributeSpecification{
		AllowedChoices: []string{"CONNECTED", "INITIALIZED", "UNKNOWN"},
		Exposed:        true,
		Name:           "operationalStatus",
		Stored:         true,
		Type:           "enum",
	},
	ServerAttributeNameStatus: elemental.AttributeSpecification{
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
	ServerAttributeNameUpdatedAt: elemental.AttributeSpecification{
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
