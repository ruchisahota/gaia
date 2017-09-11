package barretmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

import "time"

// TriremeCertIdentity represents the Identity of the object
var TriremeCertIdentity = elemental.Identity{
	Name:     "triremecert",
	Category: "triremecerts",
}

// TriremeCertsList represents a list of TriremeCerts
type TriremeCertsList []*TriremeCert

// ContentIdentity returns the identity of the objects in the list.
func (o TriremeCertsList) ContentIdentity() elemental.Identity {

	return TriremeCertIdentity
}

// List converts the object to an elemental.IdentifiablesList.
func (o TriremeCertsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TriremeCertsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o TriremeCertsList) Version() int {

	return 1
}

// TriremeCert represents the model of a triremecert
type TriremeCert struct {
	// DNSNames contains the list of requested DNS names.
	DNSNames []string `json:"DNSNames" bson:"-"`

	// CommonName contains the requested common name.
	CommonName string `json:"commonName" bson:"-"`

	// ExpirationDate contains the requested expiration date.
	ExpirationDate time.Time `json:"expirationDate" bson:"-"`

	// KeyUsage contains the generated key usage,
	KeyUsage string `json:"keyUsage" bson:"-"`

	// OrganizationalUnits contains the list of requested organizational units.
	OrganizationalUnits []string `json:"organizationalUnits" bson:"-"`

	// Organizations contains the list of requested organizations.
	Organizations []string `json:"organizations" bson:"-"`

	// SerialNumber contains the generated certificate serial number.
	SerialNumber string `json:"serialNumber" bson:"-"`

	// SigningRequest contains the eventual signing request.
	SigningRequest string `json:"signingRequest" bson:"-"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewTriremeCert returns a new *TriremeCert
func NewTriremeCert() *TriremeCert {

	return &TriremeCert{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *TriremeCert) Identity() elemental.Identity {

	return TriremeCertIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *TriremeCert) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *TriremeCert) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *TriremeCert) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *TriremeCert) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *TriremeCert) Doc() string {
	return `This api allows to retrieve a free short lived certificate for open source Trireme based product.`
}

func (o *TriremeCert) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *TriremeCert) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("commonName", o.CommonName); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("commonName", o.CommonName); err != nil {
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
func (*TriremeCert) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TriremeCertAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TriremeCertLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*TriremeCert) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TriremeCertAttributesMap
}

// TriremeCertAttributesMap represents the map of attribute for TriremeCert.
var TriremeCertAttributesMap = map[string]elemental.AttributeSpecification{
	"DNSNames": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `DNSNames contains the list of requested DNS names.`,
		Exposed:        true,
		Name:           "DNSNames",
		SubType:        "string",
		Type:           "list",
	},
	"CommonName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `CommonName contains the requested common name.`,
		Exposed:        true,
		Format:         "free",
		Name:           "commonName",
		Required:       true,
		Type:           "string",
	},
	"ExpirationDate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `ExpirationDate contains the requested expiration date.`,
		Exposed:        true,
		Name:           "expirationDate",
		SubType:        "string",
		Type:           "time",
	},
	"KeyUsage": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `KeyUsage contains the generated key usage,`,
		Exposed:        true,
		Format:         "free",
		Name:           "keyUsage",
		ReadOnly:       true,
		Type:           "string",
	},
	"OrganizationalUnits": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `OrganizationalUnits contains the list of requested organizational units.`,
		Exposed:        true,
		Name:           "organizationalUnits",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"Organizations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Organizations contains the list of requested organizations.`,
		Exposed:        true,
		Name:           "organizations",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"SerialNumber": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `SerialNumber contains the generated certificate serial number.`,
		Exposed:        true,
		Format:         "free",
		Name:           "serialNumber",
		ReadOnly:       true,
		Type:           "string",
	},
	"SigningRequest": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `SigningRequest contains the eventual signing request.`,
		Exposed:        true,
		Format:         "free",
		Name:           "signingRequest",
		Type:           "string",
	},
}

// TriremeCertLowerCaseAttributesMap represents the map of attribute for TriremeCert.
var TriremeCertLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"dnsnames": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `DNSNames contains the list of requested DNS names.`,
		Exposed:        true,
		Name:           "DNSNames",
		SubType:        "string",
		Type:           "list",
	},
	"commonname": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `CommonName contains the requested common name.`,
		Exposed:        true,
		Format:         "free",
		Name:           "commonName",
		Required:       true,
		Type:           "string",
	},
	"expirationdate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `ExpirationDate contains the requested expiration date.`,
		Exposed:        true,
		Name:           "expirationDate",
		SubType:        "string",
		Type:           "time",
	},
	"keyusage": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `KeyUsage contains the generated key usage,`,
		Exposed:        true,
		Format:         "free",
		Name:           "keyUsage",
		ReadOnly:       true,
		Type:           "string",
	},
	"organizationalunits": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `OrganizationalUnits contains the list of requested organizational units.`,
		Exposed:        true,
		Name:           "organizationalUnits",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"organizations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Organizations contains the list of requested organizations.`,
		Exposed:        true,
		Name:           "organizations",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"serialnumber": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `SerialNumber contains the generated certificate serial number.`,
		Exposed:        true,
		Format:         "free",
		Name:           "serialNumber",
		ReadOnly:       true,
		Type:           "string",
	},
	"signingrequest": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `SigningRequest contains the eventual signing request.`,
		Exposed:        true,
		Format:         "free",
		Name:           "signingRequest",
		Type:           "string",
	},
}
