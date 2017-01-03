package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

// APICheckIdentity represents the Identity of the object
var APICheckIdentity = elemental.Identity{
	Name:     "apicheck",
	Category: "apichecks",
}

// APIChecksList represents a list of APIChecks
type APIChecksList []*APICheck

// APICheck represents the model of a apicheck
type APICheck struct {
	// API is the API identity to use to use to check the api authentication
	API string `json:"API" cql:"-" bson:"-"`

	// Namespace is the namespace to use to check the api authentication.
	Namespace string `json:"namespace" cql:"-" bson:"-"`

	// Token is the token to use to check api authentication
	Token string `json:"token" cql:"-" bson:"-"`
}

// NewAPICheck returns a new *APICheck
func NewAPICheck() *APICheck {

	return &APICheck{}
}

// Identity returns the Identity of the object.
func (o *APICheck) Identity() elemental.Identity {

	return APICheckIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *APICheck) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *APICheck) SetIdentifier(ID string) {

}

func (o *APICheck) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *APICheck) Validate() error {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("namespace", o.Namespace); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("token", o.Token); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (APICheck) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return APICheckAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (APICheck) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return APICheckAttributesMap
}

// APICheckAttributesMap represents the map of attribute for APICheck.
var APICheckAttributesMap = map[string]elemental.AttributeSpecification{
	"API": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Format:         "free",
		Name:           "API",
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Format:         "free",
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"Token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Format:         "free",
		Name:           "token",
		Required:       true,
		Type:           "string",
	},
}
