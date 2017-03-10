package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"

// DependencyMapViewTypeValue represents the possible values for attribute "type".
type DependencyMapViewTypeValue string

const (
	// DependencyMapViewTypeAutomatic represents the value Automatic.
	DependencyMapViewTypeAutomatic DependencyMapViewTypeValue = "Automatic"

	// DependencyMapViewTypeManual represents the value Manual.
	DependencyMapViewTypeManual DependencyMapViewTypeValue = "Manual"
)

// DependencyMapViewIdentity represents the Identity of the object
var DependencyMapViewIdentity = elemental.Identity{
	Name:     "dependencymapview",
	Category: "dependencymapviews",
}

// DependencyMapViewsList represents a list of DependencyMapViews
type DependencyMapViewsList []*DependencyMapView

// ContentIdentity returns the identity of the objects in the list.
func (o DependencyMapViewsList) ContentIdentity() elemental.Identity {
	return DependencyMapViewIdentity
}

// List convert the object to and elemental.IdentifiablesList.
func (o DependencyMapViewsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DependencyMapView represents the model of a dependencymapview
type DependencyMapView struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags"`

	// Boolean to know if the dependency map view was computed by the system or not
	Computed bool `json:"computed" bson:"computed"`

	// CreatedTime is the time at which the object was created
	CreateTime time.Time `json:"createTime" bson:"createtime"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" bson:"namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags"`

	// A map of the tags to apply to processing units
	ProcessingUnitTags map[string][]string `json:"processingUnitTags" bson:"processingunittags"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected"`

	// Values used by the dependency map group
	Subviews DependencyMapSubviewsList `json:"subviews" bson:"subviews"`

	// Type represents the type of the dependency map. It could be manual or automatic
	Type DependencyMapViewTypeValue `json:"type" bson:"type"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`
}

// NewDependencyMapView returns a new *DependencyMapView
func NewDependencyMapView() *DependencyMapView {

	return &DependencyMapView{
		ModelVersion:   1.0,
		AssociatedTags: []string{},
		NormalizedTags: []string{},
		Subviews:       DependencyMapSubviewsList{},
		Type:           "Manual",
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

// SetIdentifier sets the value of the object's unique identifier.
func (o *DependencyMapView) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *DependencyMapView) Version() float64 {

	return 1.0
}

// Doc returns the documentation for the object
func (o *DependencyMapView) Doc() string {
	return `Allows to create a view to apply on a dependency map view. It is possible to ask the dependency map view to group all the processing units based on a primary selector, which is a tag's key, then to subgroup them based on a secondary selector.`
}

func (o *DependencyMapView) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *DependencyMapView) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *DependencyMapView) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreateTime set the given createTime of the receiver
func (o *DependencyMapView) SetCreateTime(createTime time.Time) {
	o.CreateTime = createTime
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

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *DependencyMapView) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *DependencyMapView) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetProtected returns the protected of the receiver
func (o *DependencyMapView) GetProtected() bool {
	return o.Protected
}

// SetUpdateTime set the given updateTime of the receiver
func (o *DependencyMapView) SetUpdateTime(updateTime time.Time) {
	o.UpdateTime = updateTime
}

// Validate valides the current information stored into the structure.
func (o *DependencyMapView) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredExternal("subviews", o.Subviews); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("subviews", o.Subviews); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Automatic", "Manual"}, false); err != nil {
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
func (DependencyMapView) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return DependencyMapViewAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (DependencyMapView) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return DependencyMapViewAttributesMap
}

// DependencyMapViewAttributesMap represents the map of attribute for DependencyMapView.
var DependencyMapViewAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Computed": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Boolean to know if the dependency map view was computed by the system or not`,
		Exposed:        true,
		Filterable:     true,
		Name:           "computed",
		Orderable:      true,
		Stored:         true,
		Transient:      true,
		Type:           "boolean",
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
	"ProcessingUnitTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `A map of the tags to apply to processing units`,
		Exposed:        true,
		Name:           "processingUnitTags",
		Orderable:      true,
		Stored:         true,
		SubType:        "processingunit_transient_tags_map",
		Type:           "external",
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
	"Subviews": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Values used by the dependency map group`,
		Exposed:        true,
		Name:           "subviews",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		SubType:        "dependencymapsubviews_entities",
		Type:           "external",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"Automatic", "Manual"},
		DefaultValue:   DependencyMapViewTypeValue("Manual"),
		Description:    `Type represents the type of the dependency map. It could be manual or automatic`,
		Exposed:        true,
		Filterable:     true,
		Name:           "type",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
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
}
