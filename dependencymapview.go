package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/enum"

const (
	// DependencyMapViewAttributeNameID represents the attribute ID.
	DependencyMapViewAttributeNameID elemental.AttributeSpecificationNameKey = "dependencymapview/ID"

	// DependencyMapViewAttributeNameAnnotation represents the attribute annotation.
	DependencyMapViewAttributeNameAnnotation elemental.AttributeSpecificationNameKey = "dependencymapview/annotation"

	// DependencyMapViewAttributeNameAssociatedTags represents the attribute associatedTags.
	DependencyMapViewAttributeNameAssociatedTags elemental.AttributeSpecificationNameKey = "dependencymapview/associatedTags"

	// DependencyMapViewAttributeNameCreatedAt represents the attribute createdAt.
	DependencyMapViewAttributeNameCreatedAt elemental.AttributeSpecificationNameKey = "dependencymapview/createdAt"

	// DependencyMapViewAttributeNameDeleted represents the attribute deleted.
	DependencyMapViewAttributeNameDeleted elemental.AttributeSpecificationNameKey = "dependencymapview/deleted"

	// DependencyMapViewAttributeNameDescription represents the attribute description.
	DependencyMapViewAttributeNameDescription elemental.AttributeSpecificationNameKey = "dependencymapview/description"

	// DependencyMapViewAttributeNameName represents the attribute name.
	DependencyMapViewAttributeNameName elemental.AttributeSpecificationNameKey = "dependencymapview/name"

	// DependencyMapViewAttributeNameNamespace represents the attribute namespace.
	DependencyMapViewAttributeNameNamespace elemental.AttributeSpecificationNameKey = "dependencymapview/namespace"

	// DependencyMapViewAttributeNameStatus represents the attribute status.
	DependencyMapViewAttributeNameStatus elemental.AttributeSpecificationNameKey = "dependencymapview/status"

	// DependencyMapViewAttributeNameSubviews represents the attribute subviews.
	DependencyMapViewAttributeNameSubviews elemental.AttributeSpecificationNameKey = "dependencymapview/subviews"

	// DependencyMapViewAttributeNameUpdatedAt represents the attribute updatedAt.
	DependencyMapViewAttributeNameUpdatedAt elemental.AttributeSpecificationNameKey = "dependencymapview/updatedAt"
)

// DependencyMapViewIdentity represents the Identity of the object
var DependencyMapViewIdentity = elemental.Identity{
	Name:     "dependencymapview",
	Category: "dependencymapviews",
}

// DependencyMapViewsList represents a list of DependencyMapViews
type DependencyMapViewsList []*DependencyMapView

// DependencyMapView represents the model of a dependencymapview
type DependencyMapView struct {
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

	// Values used by the dependency map group
	Subviews DependencyMapSubviewsList `json:"subviews,omitempty" cql:"subviews,omitempty"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt,omitempty" cql:"updatedat,omitempty"`
}

// NewDependencyMapView returns a new *DependencyMapView
func NewDependencyMapView() *DependencyMapView {

	return &DependencyMapView{
		Status:   enum.Active,
		Subviews: DependencyMapSubviewsList{},
	}
}

// Identity returns the Identity of the object.
func (o *DependencyMapView) Identity() elemental.Identity {

	return DependencyMapViewIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DependencyMapView) Identifier() string {

	return o.ID
}

func (o *DependencyMapView) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DependencyMapView) SetIdentifier(ID string) {

	o.ID = ID
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *DependencyMapView) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *DependencyMapView) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *DependencyMapView) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetDeleted returns the deleted of the receiver
func (o *DependencyMapView) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *DependencyMapView) SetDeleted(deleted bool) {
	o.Deleted = deleted
}

// GetName returns the name of the receiver
func (o *DependencyMapView) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *DependencyMapView) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *DependencyMapView) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *DependencyMapView) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetStatus returns the status of the receiver
func (o *DependencyMapView) GetStatus() enum.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *DependencyMapView) SetStatus(status enum.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *DependencyMapView) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *DependencyMapView) Validate() elemental.Errors {

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
func (o DependencyMapView) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return DependencyMapViewAttributesMap[name]
}

// DependencyMapViewAttributesMap represents the map of attribute for DependencyMapView.
var DependencyMapViewAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	DependencyMapViewAttributeNameID: elemental.AttributeSpecification{
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
	DependencyMapViewAttributeNameAnnotation: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	DependencyMapViewAttributeNameAssociatedTags: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "associatedTags",
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	DependencyMapViewAttributeNameCreatedAt: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	DependencyMapViewAttributeNameDeleted: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Filterable:     true,
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	DependencyMapViewAttributeNameDescription: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	DependencyMapViewAttributeNameName: elemental.AttributeSpecification{
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
	DependencyMapViewAttributeNameNamespace: elemental.AttributeSpecification{
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
	DependencyMapViewAttributeNameStatus: elemental.AttributeSpecification{
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
	DependencyMapViewAttributeNameSubviews: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Name:           "subviews",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		SubType:        "dependencymapsubviews_entities",
		Type:           "external",
	},
	DependencyMapViewAttributeNameUpdatedAt: elemental.AttributeSpecification{
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
