package models

import "fmt"
import "github.com/aporeto-inc/elemental"

// ActivateIdentity represents the Identity of the object
var ActivateIdentity = elemental.Identity{
	Name:     "activate",
	Category: "activate",
}

// ActivatesList represents a list of Activates
type ActivatesList []*Activate

// Activate represents the model of a activate
type Activate struct {
	// Token contains the activation token
	Token string `json:"token" cql:"-" bson:"-"`
}

// NewActivate returns a new *Activate
func NewActivate() *Activate {

	return &Activate{}
}

// Identity returns the Identity of the object.
func (o *Activate) Identity() elemental.Identity {

	return ActivateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Activate) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Activate) SetIdentifier(ID string) {

}

func (o *Activate) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Activate) Validate() error {

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
func (Activate) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return ActivateAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (Activate) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ActivateAttributesMap
}

// ActivateAttributesMap represents the map of attribute for Activate.
var ActivateAttributesMap = map[string]elemental.AttributeSpecification{
	"Token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Token contains the activation token`,
		Exposed:        true,
		Format:         "free",
		Name:           "token",
		Type:           "string",
	},
}
