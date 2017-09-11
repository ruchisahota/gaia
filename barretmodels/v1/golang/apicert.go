package barretmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

import "time"

// APICertIdentity represents the Identity of the object
var APICertIdentity = elemental.Identity{
	Name:     "apicert",
	Category: "apicerts",
}

// APICertsList represents a list of APICerts
type APICertsList []*APICert

// ContentIdentity returns the identity of the objects in the list.
func (o APICertsList) ContentIdentity() elemental.Identity {

	return APICertIdentity
}

// List converts the object to an elemental.IdentifiablesList.
func (o APICertsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o APICertsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o APICertsList) Version() int {

	return 1
}

// APICert represents the model of a apicert
type APICert struct {
	// CommonName contains the requested common name.
	CommonName string `json:"commonName" bson:"-"`

	// Data contains the certificate data.
	Data string `json:"data" bson:"-"`

	// Emails contains the list of requested emails.
	Emails []string `json:"emails" bson:"-"`

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

	// P12 contains the certificate in the p12 fomat.
	P12 string `json:"p12" bson:"-"`

	// Password contains the requested p12 password.
	Password string `json:"password" bson:"-"`

	// SerialNumber contains the generated certificate serial number.
	SerialNumber string `json:"serialNumber" bson:"-"`

	// SigningRequest contains the eventual signing request.
	SigningRequest string `json:"signingRequest" bson:"-"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewAPICert returns a new *APICert
func NewAPICert() *APICert {

	return &APICert{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *APICert) Identity() elemental.Identity {

	return APICertIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *APICert) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *APICert) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *APICert) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *APICert) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *APICert) Doc() string {
	return `This API allows to retrieve an client certifcate for api authentication.`
}

func (o *APICert) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *APICert) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("commonName", o.CommonName); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("commonName", o.CommonName); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("password", o.Password); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("password", o.Password); err != nil {
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
func (*APICert) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := APICertAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return APICertLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*APICert) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return APICertAttributesMap
}

// APICertAttributesMap represents the map of attribute for APICert.
var APICertAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Emails": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Emails contains the list of requested emails.`,
		Exposed:        true,
		Name:           "emails",
		Required:       true,
		SubType:        "string",
		Type:           "list",
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
	"P12": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `P12 contains the certificate in the p12 fomat.`,
		Exposed:        true,
		Format:         "free",
		Name:           "p12",
		ReadOnly:       true,
		Type:           "string",
	},
	"Password": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Password contains the requested p12 password. `,
		Exposed:        true,
		Format:         "free",
		Name:           "password",
		Required:       true,
		Type:           "string",
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

// APICertLowerCaseAttributesMap represents the map of attribute for APICert.
var APICertLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"emails": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Emails contains the list of requested emails.`,
		Exposed:        true,
		Name:           "emails",
		Required:       true,
		SubType:        "string",
		Type:           "list",
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
	"p12": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `P12 contains the certificate in the p12 fomat.`,
		Exposed:        true,
		Format:         "free",
		Name:           "p12",
		ReadOnly:       true,
		Type:           "string",
	},
	"password": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Password contains the requested p12 password. `,
		Exposed:        true,
		Format:         "free",
		Name:           "password",
		Required:       true,
		Type:           "string",
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
