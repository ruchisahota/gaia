package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	// FilePathAttributeNameID represents the attribute ID.
	FilePathAttributeNameID elemental.AttributeSpecificationNameKey = "filepath/ID"

	// FilePathAttributeNameAnnotation represents the attribute annotation.
	FilePathAttributeNameAnnotation elemental.AttributeSpecificationNameKey = "filepath/annotation"

	// FilePathAttributeNameAssociatedTags represents the attribute associatedTags.
	FilePathAttributeNameAssociatedTags elemental.AttributeSpecificationNameKey = "filepath/associatedTags"

	// FilePathAttributeNameCreatedAt represents the attribute createdAt.
	FilePathAttributeNameCreatedAt elemental.AttributeSpecificationNameKey = "filepath/createdAt"

	// FilePathAttributeNameDeleted represents the attribute deleted.
	FilePathAttributeNameDeleted elemental.AttributeSpecificationNameKey = "filepath/deleted"

	// FilePathAttributeNameDescription represents the attribute description.
	FilePathAttributeNameDescription elemental.AttributeSpecificationNameKey = "filepath/description"

	// FilePathAttributeNameFilepath represents the attribute filepath.
	FilePathAttributeNameFilepath elemental.AttributeSpecificationNameKey = "filepath/filepath"

	// FilePathAttributeNameName represents the attribute name.
	FilePathAttributeNameName elemental.AttributeSpecificationNameKey = "filepath/name"

	// FilePathAttributeNameNamespace represents the attribute namespace.
	FilePathAttributeNameNamespace elemental.AttributeSpecificationNameKey = "filepath/namespace"

	// FilePathAttributeNameServer represents the attribute server.
	FilePathAttributeNameServer elemental.AttributeSpecificationNameKey = "filepath/server"

	// FilePathAttributeNameStatus represents the attribute status.
	FilePathAttributeNameStatus elemental.AttributeSpecificationNameKey = "filepath/status"

	// FilePathAttributeNameUpdatedAt represents the attribute updatedAt.
	FilePathAttributeNameUpdatedAt elemental.AttributeSpecificationNameKey = "filepath/updatedAt"
)

// FilePathIdentity represents the Identity of the object
var FilePathIdentity = elemental.Identity{
	Name:     "filepath",
	Category: "filepaths",
}

// FilePathsList represents a list of FilePaths
type FilePathsList []*FilePath

// FilePath represents the model of a filepath
type FilePath struct {
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

	// FilePath refer to the file mount path
	Filepath string `json:"filepath,omitempty" cql:"filepath,omitempty"`

	// Name is the name of the entity
	Name string `json:"name,omitempty" cql:"name,omitempty"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace,omitempty" cql:"namespace,primarykey,omitempty"`

	// server is the server name/ID/IP associated with the file path
	Server string `json:"server,omitempty" cql:"server,omitempty"`

	// Status of an entity
	Status enum.EntityStatus `json:"status,omitempty" cql:"status,omitempty"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt,omitempty" cql:"updatedat,omitempty"`
}

// NewFilePath returns a new *FilePath
func NewFilePath() *FilePath {

	return &FilePath{
		Status: enum.Active,
	}
}

// Identity returns the Identity of the object.
func (o *FilePath) Identity() elemental.Identity {

	return FilePathIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FilePath) Identifier() string {

	return o.ID
}

func (o *FilePath) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FilePath) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *FilePath) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *FilePath) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *FilePath) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *FilePath) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *FilePath) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetName returns the name of the receiver
func (o *FilePath) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *FilePath) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *FilePath) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *FilePath) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetStatus returns the status of the receiver
func (o *FilePath) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *FilePath) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *FilePath) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *FilePath) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("filepath", o.Filepath); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("server", o.Server); err != nil {
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
func (o FilePath) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return FilePathAttributesMap[name]
}

// FilePathAttributesMap represents the map of attribute for FilePath.
var FilePathAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	FilePathAttributeNameID: elemental.AttributeSpecification{
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
	FilePathAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	FilePathAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	FilePathAttributeNameCreatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	FilePathAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Filterable:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	FilePathAttributeNameDescription: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	FilePathAttributeNameFilepath: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Name:           "filepath",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	FilePathAttributeNameName: elemental.AttributeSpecification{
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
	FilePathAttributeNameNamespace: elemental.AttributeSpecification{
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
	FilePathAttributeNameServer: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Name:           "server",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	FilePathAttributeNameStatus: elemental.AttributeSpecification{
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
	FilePathAttributeNameUpdatedAt: elemental.AttributeSpecification{
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
