package barretmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

import "time"

// EnforcerCertIdentity represents the Identity of the object
var EnforcerCertIdentity = elemental.Identity{
	Name:     "enforcercert",
	Category: "enforcercerts",
}

// EnforcerCertsList represents a list of EnforcerCerts
type EnforcerCertsList []*EnforcerCert

// ContentIdentity returns the identity of the objects in the list.
func (o EnforcerCertsList) ContentIdentity() elemental.Identity {

	return EnforcerCertIdentity
}

// List converts the object to an elemental.IdentifiablesList.
func (o EnforcerCertsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EnforcerCertsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o EnforcerCertsList) Version() int {

	return 1
}

// EnforcerCert represents the model of a enforcercert
type EnforcerCert struct {
	// DNSNames contains the list of requested DNS names.
	DNSNames []string `json:"DNSNames" bson:"-"`

	// CommonName contains the requested common name.
	CommonName string `json:"commonName" bson:"-"`

	// Data contains the certificate data.
	Data string `json:"data" bson:"-"`

	// ExpirationDate contains the requested expiration date.
	ExpirationDate time.Time `json:"expirationDate" bson:"-"`

	// Key contains the certificate key.
	Key string `json:"key" bson:"-"`

	// KeyUsage contains the generated key usage,
	KeyUsage string `json:"keyUsage" bson:"-"`

	// OrganizationalUnits contains the list of requested organizational units.
	OrganizationalUnits []string `json:"organizationalUnits" bson:"-"`

	// Organizations contains the list of requested organizations.
	Organizations []string `json:"organizations" bson:"-"`

	// SerialNumber contains the generated certificate serial number.
	SerialNumber string `json:"serialNumber" bson:"-"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewEnforcerCert returns a new *EnforcerCert
func NewEnforcerCert() *EnforcerCert {

	return &EnforcerCert{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *EnforcerCert) Identity() elemental.Identity {

	return EnforcerCertIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnforcerCert) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnforcerCert) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *EnforcerCert) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *EnforcerCert) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *EnforcerCert) Doc() string {
	return `This api allows to create an enforcerd certificate.`
}

func (o *EnforcerCert) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *EnforcerCert) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("commonName", o.CommonName); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("commonName", o.CommonName); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredTime("expirationDate", o.ExpirationDate); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredTime("expirationDate", o.ExpirationDate); err != nil {
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
func (*EnforcerCert) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EnforcerCertAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EnforcerCertLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*EnforcerCert) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EnforcerCertAttributesMap
}

// EnforcerCertAttributesMap represents the map of attribute for EnforcerCert.
var EnforcerCertAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Data contains the certificate data.`,
		Exposed:        true,
		Format:         "free",
		Name:           "data",
		ReadOnly:       true,
		Type:           "string",
	},
	"ExpirationDate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `ExpirationDate contains the requested expiration date.`,
		Exposed:        true,
		Name:           "expirationDate",
		Required:       true,
		SubType:        "string",
		Type:           "time",
	},
	"Key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Key contains the certificate key.`,
		Exposed:        true,
		Format:         "free",
		Name:           "key",
		ReadOnly:       true,
		Type:           "string",
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
}

// EnforcerCertLowerCaseAttributesMap represents the map of attribute for EnforcerCert.
var EnforcerCertLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Data contains the certificate data.`,
		Exposed:        true,
		Format:         "free",
		Name:           "data",
		ReadOnly:       true,
		Type:           "string",
	},
	"expirationdate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `ExpirationDate contains the requested expiration date.`,
		Exposed:        true,
		Name:           "expirationDate",
		Required:       true,
		SubType:        "string",
		Type:           "time",
	},
	"key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Key contains the certificate key.`,
		Exposed:        true,
		Format:         "free",
		Name:           "key",
		ReadOnly:       true,
		Type:           "string",
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
}
