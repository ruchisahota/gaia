package vincemodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"

// AccountStatusValue represents the possible values for attribute "status".
type AccountStatusValue string

const (
	// AccountStatusActive represents the value Active.
	AccountStatusActive AccountStatusValue = "Active"

	// AccountStatusDisabled represents the value Disabled.
	AccountStatusDisabled AccountStatusValue = "Disabled"

	// AccountStatusInvited represents the value Invited.
	AccountStatusInvited AccountStatusValue = "Invited"

	// AccountStatusPending represents the value Pending.
	AccountStatusPending AccountStatusValue = "Pending"
)

// AccountIdentity represents the Identity of the object
var AccountIdentity = elemental.Identity{
	Name:     "account",
	Category: "accounts",
}

// AccountsList represents a list of Accounts
type AccountsList []*Account

// Account represents the model of a account
type Account struct {
	// ID of the object.
	ID string `json:"ID" cql:"id,primarykey,omitempty" bson:"_id"`

	// LDAPAddress holds the account authentication account's private ldap server.
	LDAPAddress string `json:"LDAPAddress" cql:"ldapaddress,omitempty" bson:"ldapaddress"`

	// LDAPBaseDN holds the base DN to use to ldap queries.
	LDAPBaseDN string `json:"LDAPBaseDN" cql:"ldapbasedn,omitempty" bson:"ldapbasedn"`

	// LDAPBindDN holds the account's internal LDAP bind dn for querying auth.
	LDAPBindDN string `json:"LDAPBindDN" cql:"ldapbinddn,omitempty" bson:"ldapbinddn"`

	// LDAPBindPassword holds the password to the LDAPBindDN.
	LDAPBindPassword string `json:"LDAPBindPassword" cql:"ldapbindpassword,omitempty" bson:"ldapbindpassword"`

	// LDAPEnabled triggers if the account uses it's own LDAP for authentication.
	LDAPEnabled bool `json:"LDAPEnabled" cql:"ldapenabled,omitempty" bson:"ldapenabled"`

	// AccessEnabled defines if the account holder should have access to the systems.
	AccessEnabled bool `json:"accessEnabled" cql:"accessenabled,omitempty" bson:"accessenabled"`

	// ActivationExpiration contains the expiration date of the activation token.
	ActivationExpiration time.Time `json:"-" cql:"activationexpiration,omitempty" bson:"activationexpiration"`

	// ActivationToken contains the activation token.
	ActivationToken string `json:"-" cql:"activationtoken,omitempty" bson:"activationtoken"`

	// AssociatedAPIAuthPolicyID holds the ID of the associated API auth policy.
	AssociatedAPIAuthPolicyID string `json:"associatedAPIAuthPolicyID" cql:"associatedapiauthpolicyid,omitempty" bson:"associatedapiauthpolicyid"`

	// associatedAgentReadAPIAuthPolicyID holds the ID of the associated agent read auth policy.
	AssociatedAgentReadAPIAuthPolicyID string `json:"associatedAgentReadAPIAuthPolicyID" cql:"associatedagentreadapiauthpolicyid,omitempty" bson:"associatedagentreadapiauthpolicyid"`

	// associatedAgentWriteAPIAuthPolicyID holds the ID of the associated agent write auth policy.
	AssociatedAgentWriteAPIAuthPolicyID string `json:"associatedAgentWriteAPIAuthPolicyID" cql:"associatedagentwriteapiauthpolicyid,omitempty" bson:"associatedagentwriteapiauthpolicyid"`

	// AssociatedNSMappingPolicyID contains the ID of the associated Namespace Mapping Policy.
	AssociatedNSMappingPolicyID string `json:"associatedNSMappingPolicyID" cql:"associatednsmappingpolicyid,omitempty" bson:"associatednsmappingpolicyid"`

	// AssociatedNamespaceID contains the ID of the associated namespace.
	AssociatedNamespaceID string `json:"associatedNamespaceID" cql:"associatednamespaceid,omitempty" bson:"associatednamespaceid"`

	// BumpToken contains the tag processing unit must use to be placed in the account's namespace.
	BumpToken string `json:"bumpToken" cql:"bumptoken,omitempty" bson:"bumptoken"`

	// CreatedAt represents the creation date of the objct.
	CreatedAt time.Time `json:"createdAt" cql:"createdat,omitempty" bson:"createdat"`

	// Email of the account holder.
	Email string `json:"email" cql:"email,omitempty" bson:"email"`

	// Name of the account.
	Name string `json:"name" cql:"name,omitempty" bson:"name"`

	// Password for the account.
	Password string `json:"password" cql:"password,omitempty" bson:"password"`

	// Status of the account.
	Status AccountStatusValue `json:"status" cql:"status,omitempty" bson:"status"`

	// UpdatedAt represents the last update date of the objct.
	UpdatedAt time.Time `json:"updatedAt" cql:"updatedat,omitempty" bson:"updatedat"`
}

// NewAccount returns a new *Account
func NewAccount() *Account {

	return &Account{
		Status: "Pending",
	}
}

// Identity returns the Identity of the object.
func (o *Account) Identity() elemental.Identity {

	return AccountIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Account) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Account) SetIdentifier(ID string) {

	o.ID = ID
}

func (o *Account) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Account) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("email", o.Email); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("email", o.Email); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Disabled", "Invited", "Pending"}, true); err != nil {
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
func (Account) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return AccountAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (Account) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AccountAttributesMap
}

// AccountAttributesMap represents the map of attribute for Account.
var AccountAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID of the object.`,
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
	},
	"LDAPAddress": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LDAPAddress holds the account authentication account's private ldap server.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPAddress",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LDAPBaseDN": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LDAPBaseDN holds the base DN to use to ldap queries.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPBaseDN",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LDAPBindDN": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LDAPBindDN holds the account's internal LDAP bind dn for querying auth.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPBindDN",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LDAPBindPassword": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LDAPBindPassword holds the password to the LDAPBindDN. `,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPBindPassword",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LDAPEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LDAPEnabled triggers if the account uses it's own LDAP for authentication.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "LDAPEnabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"AccessEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AccessEnabled defines if the account holder should have access to the systems.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "accessEnabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"ActivationExpiration": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ActivationExpiration contains the expiration date of the activation token.`,
		Name:           "activationExpiration",
		Stored:         true,
		Type:           "time",
	},
	"ActivationToken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ActivationToken contains the activation token.`,
		Format:         "free",
		Name:           "activationToken",
		Stored:         true,
		Type:           "string",
	},
	"AssociatedAPIAuthPolicyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `AssociatedAPIAuthPolicyID holds the ID of the associated API auth policy.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "associatedAPIAuthPolicyID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"AssociatedAgentReadAPIAuthPolicyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `associatedAgentReadAPIAuthPolicyID holds the ID of the associated agent read auth policy.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "associatedAgentReadAPIAuthPolicyID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"AssociatedAgentWriteAPIAuthPolicyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `associatedAgentWriteAPIAuthPolicyID holds the ID of the associated agent write auth policy.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "associatedAgentWriteAPIAuthPolicyID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"AssociatedNSMappingPolicyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `AssociatedNSMappingPolicyID contains the ID of the associated Namespace Mapping Policy.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "associatedNSMappingPolicyID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"AssociatedNamespaceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `AssociatedNamespaceID contains the ID of the associated namespace.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "associatedNamespaceID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"BumpToken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `BumpToken contains the tag processing unit must use to be placed in the account's namespace.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "bumpToken",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"CreatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `CreatedAt represents the creation date of the objct.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "createdAt",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"Email": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Email of the account holder.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "email",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Name of the account.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Password": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Password for the account.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "password",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{"Active", "Disabled", "Invited", "Pending"},
		Autogenerated:  true,
		Description:    `Status of the account.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
	"UpdatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `UpdatedAt represents the last update date of the objct.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "updatedAt",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}
