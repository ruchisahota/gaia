package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// DependencyMapSubviewIdentity represents the Identity of the object
var DependencyMapSubviewIdentity = elemental.Identity{
	Name:     "dependencymapsubview",
	Category: "dependencymapsubviews",
}

// DependencyMapSubviewsList represents a list of DependencyMapSubviews
type DependencyMapSubviewsList []*DependencyMapSubview

// ContentIdentity returns the identity of the objects in the list.
func (o DependencyMapSubviewsList) ContentIdentity() elemental.Identity {
	return DependencyMapSubviewIdentity
}

// List convert the object to and elemental.IdentifiablesList.
func (o DependencyMapSubviewsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DependencyMapSubview represents the model of a dependencymapsubview
type DependencyMapSubview struct {
	// Selector is the main selector for the DependencyMapSubview.
	Selector []string `json:"selector" bson:"selector"`

	// SubSelectors are the selector to apply inside the main selector.
	SubSelectors map[string][]string `json:"subSelectors" bson:"subselectors"`

	// Tonality sets the color tonality to use for the DependencyMapSubView.
	Tonality string `json:"tonality" bson:"tonality"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewDependencyMapSubview returns a new *DependencyMapSubview
func NewDependencyMapSubview() *DependencyMapSubview {

	return &DependencyMapSubview{
		ModelVersion: 1.0,
		Selector:     []string{},
		SubSelectors: map[string][]string{},
	}
}

// Identity returns the Identity of the object.
func (o *DependencyMapSubview) Identity() elemental.Identity {

	return DependencyMapSubviewIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DependencyMapSubview) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DependencyMapSubview) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *DependencyMapSubview) Version() float64 {

	return 1.0
}

// Doc returns the documentation for the object
func (o *DependencyMapSubview) Doc() string {

	return nodocString

}

func (o *DependencyMapSubview) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *DependencyMapSubview) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*DependencyMapSubview) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return DependencyMapSubviewAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*DependencyMapSubview) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return DependencyMapSubviewAttributesMap
}

// DependencyMapSubviewAttributesMap represents the map of attribute for DependencyMapSubview.
var DependencyMapSubviewAttributesMap = map[string]elemental.AttributeSpecification{
	"Selector": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Selector is the main selector for the DependencyMapSubview.`,
		Exposed:        true,
		Name:           "selector",
		Orderable:      true,
		Stored:         true,
		SubType:        "dependencymapview_selector",
		Type:           "external",
	},
	"SubSelectors": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `SubSelectors are the selector to apply inside the main selector.`,
		Exposed:        true,
		Name:           "subSelectors",
		Orderable:      true,
		Stored:         true,
		SubType:        "dependencymapview_subselector",
		Type:           "external",
	},
	"Tonality": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Tonality sets the color tonality to use for the DependencyMapSubView.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "tonality",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}
