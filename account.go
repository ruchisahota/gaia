package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

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

// AccountIdentity represents the Identity of the object.
var AccountIdentity = elemental.Identity{
	Name:     "account",
	Category: "accounts",
	Package:  "vince",
	Private:  false,
}

// AccountsList represents a list of Accounts
type AccountsList []*Account

// Identity returns the identity of the objects in the list.
func (o AccountsList) Identity() elemental.Identity {

	return AccountIdentity
}

// Copy returns a pointer to a copy the AccountsList.
func (o AccountsList) Copy() elemental.Identifiables {

	copy := append(AccountsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AccountsList.
func (o AccountsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AccountsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Account))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AccountsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AccountsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the AccountsList converted to SparseAccountsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AccountsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAccountsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAccount)
	}

	return out
}

// Version returns the version of the content.
func (o AccountsList) Version() int {

	return 1
}

// Account represents the model of a account
type Account struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Enable or disable two-factor authentication.
	OTPEnabled bool `json:"OTPEnabled" msgpack:"OTPEnabled" bson:"otpenabled" mapstructure:"OTPEnabled,omitempty"`

	// Returns the base64-encoded QR code for setting up two-factor authentication.
	OTPQRCode string `json:"OTPQRCode" msgpack:"OTPQRCode" bson:"-" mapstructure:"OTPQRCode,omitempty"`

	// Stores the two-factor authentication secret.
	OTPSecret string `json:"-" msgpack:"-" bson:"otpsecret" mapstructure:"-,omitempty"`

	// Holds the SSH certificate authority used by the account namespace.
	SSHCA string `json:"SSHCA" msgpack:"SSHCA" bson:"-" mapstructure:"SSHCA,omitempty"`

	// Set to `true` to renew the SSH certificate authority of the account namespace.
	SSHCARenew bool `json:"SSHCARenew" msgpack:"SSHCARenew" bson:"-" mapstructure:"SSHCARenew,omitempty"`

	// Defines if the account holder should have access to the system.
	AccessEnabled bool `json:"accessEnabled" msgpack:"accessEnabled" bson:"accessenabled" mapstructure:"accessEnabled,omitempty"`

	// Contains the expiration date of the activation token.
	ActivationExpiration time.Time `json:"-" msgpack:"-" bson:"activationexpiration" mapstructure:"-,omitempty"`

	// Contains the activation token.
	ActivationToken string `json:"activationToken,omitempty" msgpack:"activationToken,omitempty" bson:"activationtoken,omitempty" mapstructure:"activationToken,omitempty"`

	// Holds the ID of the associated API authorization.
	AssociatedAPIAuthPolicyID string `json:"-" msgpack:"-" bson:"associatedapiauthpolicyid" mapstructure:"-,omitempty"`

	// Contains a map of associated AWS defender policies.
	AssociatedAWSPolicies map[string]string `json:"-" msgpack:"-" bson:"associatedawspolicies" mapstructure:"-,omitempty"`

	// Holds the ID of the associated billing customer.
	AssociatedBillingID string `json:"associatedBillingID" msgpack:"associatedBillingID" bson:"associatedbillingid" mapstructure:"associatedBillingID,omitempty"`

	// Contains the ID of the associated namespace.
	AssociatedNamespaceID string `json:"-" msgpack:"-" bson:"associatednamespaceid" mapstructure:"-,omitempty"`

	// Contains the plan key associated with this account.
	AssociatedPlanKey string `json:"associatedPlanKey" msgpack:"associatedPlanKey" bson:"associatedplankey" mapstructure:"associatedPlanKey,omitempty"`

	// Contains a map of the associated quotas.
	AssociatedQuotaPolicies map[string]string `json:"-" msgpack:"-" bson:"associatedquotapolicies" mapstructure:"-,omitempty"`

	// Company of the account user.
	Company string `json:"company" msgpack:"company" bson:"company" mapstructure:"company,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Email of the account holder.
	Email string `json:"email" msgpack:"email" bson:"email" mapstructure:"email,omitempty"`

	// Internally keeps track of the number of failed attempt.
	FailedAuth int `json:"-" msgpack:"-" bson:"failedauth" mapstructure:"-,omitempty"`

	// Internally keeps track of the time of the last failed attempt.
	FailedTime time.Time `json:"-" msgpack:"-" bson:"failedtime" mapstructure:"-,omitempty"`

	// First name of the account user.
	FirstName string `json:"firstName" msgpack:"firstName" bson:"firstname" mapstructure:"firstName,omitempty"`

	// Last name of the account user.
	LastName string `json:"lastName" msgpack:"lastName" bson:"lastname" mapstructure:"lastName,omitempty"`

	// The certificate authority used by this namespace.
	LocalCA string `json:"localCA" msgpack:"localCA" bson:"-" mapstructure:"localCA,omitempty"`

	// Set to `true` to renew the local certificate authority of the account namespace.
	LocalCARenew bool `json:"localCARenew" msgpack:"localCARenew" bson:"-" mapstructure:"localCARenew,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the account.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// New password for the account. If set the previous password must be given through
	// the property `password`.
	NewPassword string `json:"newPassword" msgpack:"newPassword" bson:"-" mapstructure:"newPassword,omitempty"`

	// Password for the account.
	Password string `json:"password" msgpack:"password" bson:"password" mapstructure:"password,omitempty"`

	// Contains the completely automated public Turing test (CAPTCHA)
	// validation if reCAPTCHA is enabled.
	ReCAPTCHAKey string `json:"reCAPTCHAKey" msgpack:"reCAPTCHAKey" bson:"-" mapstructure:"reCAPTCHAKey,omitempty"`

	// Contains the expiration time for resetting the password.
	ResetPasswordExpiration time.Time `json:"-" msgpack:"-" bson:"resetpasswordexpiration" mapstructure:"-,omitempty"`

	// Contains the token to use for resetting password.
	ResetPasswordToken string `json:"-" msgpack:"-" bson:"resetpasswordtoken" mapstructure:"-,omitempty"`

	// Status of the account.
	Status AccountStatusValue `json:"status" msgpack:"status" bson:"status" mapstructure:"status,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAccount returns a new *Account
func NewAccount() *Account {

	return &Account{
		ModelVersion:            1,
		AssociatedAWSPolicies:   map[string]string{},
		AssociatedQuotaPolicies: map[string]string{},
		MigrationsLog:           map[string]string{},
		Status:                  AccountStatusPending,
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
func (o *Account) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Account) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAccount{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.OTPEnabled = o.OTPEnabled
	s.OTPSecret = o.OTPSecret
	s.AccessEnabled = o.AccessEnabled
	s.ActivationExpiration = o.ActivationExpiration
	s.ActivationToken = o.ActivationToken
	s.AssociatedAPIAuthPolicyID = o.AssociatedAPIAuthPolicyID
	s.AssociatedAWSPolicies = o.AssociatedAWSPolicies
	s.AssociatedBillingID = o.AssociatedBillingID
	s.AssociatedNamespaceID = o.AssociatedNamespaceID
	s.AssociatedPlanKey = o.AssociatedPlanKey
	s.AssociatedQuotaPolicies = o.AssociatedQuotaPolicies
	s.Company = o.Company
	s.CreateTime = o.CreateTime
	s.Email = o.Email
	s.FailedAuth = o.FailedAuth
	s.FailedTime = o.FailedTime
	s.FirstName = o.FirstName
	s.LastName = o.LastName
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Password = o.Password
	s.ResetPasswordExpiration = o.ResetPasswordExpiration
	s.ResetPasswordToken = o.ResetPasswordToken
	s.Status = o.Status
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Account) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAccount{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.OTPEnabled = s.OTPEnabled
	o.OTPSecret = s.OTPSecret
	o.AccessEnabled = s.AccessEnabled
	o.ActivationExpiration = s.ActivationExpiration
	o.ActivationToken = s.ActivationToken
	o.AssociatedAPIAuthPolicyID = s.AssociatedAPIAuthPolicyID
	o.AssociatedAWSPolicies = s.AssociatedAWSPolicies
	o.AssociatedBillingID = s.AssociatedBillingID
	o.AssociatedNamespaceID = s.AssociatedNamespaceID
	o.AssociatedPlanKey = s.AssociatedPlanKey
	o.AssociatedQuotaPolicies = s.AssociatedQuotaPolicies
	o.Company = s.Company
	o.CreateTime = s.CreateTime
	o.Email = s.Email
	o.FailedAuth = s.FailedAuth
	o.FailedTime = s.FailedTime
	o.FirstName = s.FirstName
	o.LastName = s.LastName
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Password = s.Password
	o.ResetPasswordExpiration = s.ResetPasswordExpiration
	o.ResetPasswordToken = s.ResetPasswordToken
	o.Status = s.Status
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Account) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Account) BleveType() string {

	return "account"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Account) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Account) Doc() string {

	return `Allows you to view and manage basic information about your account like
your name, password, and whether or not two-factor authentication is enabled.`
}

func (o *Account) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Account) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *Account) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *Account) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *Account) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Account) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *Account) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *Account) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Account) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Account) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Account) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Account) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAccount{
			ID:                        &o.ID,
			OTPEnabled:                &o.OTPEnabled,
			OTPQRCode:                 &o.OTPQRCode,
			OTPSecret:                 &o.OTPSecret,
			SSHCA:                     &o.SSHCA,
			SSHCARenew:                &o.SSHCARenew,
			AccessEnabled:             &o.AccessEnabled,
			ActivationExpiration:      &o.ActivationExpiration,
			ActivationToken:           &o.ActivationToken,
			AssociatedAPIAuthPolicyID: &o.AssociatedAPIAuthPolicyID,
			AssociatedAWSPolicies:     &o.AssociatedAWSPolicies,
			AssociatedBillingID:       &o.AssociatedBillingID,
			AssociatedNamespaceID:     &o.AssociatedNamespaceID,
			AssociatedPlanKey:         &o.AssociatedPlanKey,
			AssociatedQuotaPolicies:   &o.AssociatedQuotaPolicies,
			Company:                   &o.Company,
			CreateTime:                &o.CreateTime,
			Email:                     &o.Email,
			FailedAuth:                &o.FailedAuth,
			FailedTime:                &o.FailedTime,
			FirstName:                 &o.FirstName,
			LastName:                  &o.LastName,
			LocalCA:                   &o.LocalCA,
			LocalCARenew:              &o.LocalCARenew,
			MigrationsLog:             &o.MigrationsLog,
			Name:                      &o.Name,
			NewPassword:               &o.NewPassword,
			Password:                  &o.Password,
			ReCAPTCHAKey:              &o.ReCAPTCHAKey,
			ResetPasswordExpiration:   &o.ResetPasswordExpiration,
			ResetPasswordToken:        &o.ResetPasswordToken,
			Status:                    &o.Status,
			UpdateTime:                &o.UpdateTime,
			ZHash:                     &o.ZHash,
			Zone:                      &o.Zone,
		}
	}

	sp := &SparseAccount{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "OTPEnabled":
			sp.OTPEnabled = &(o.OTPEnabled)
		case "OTPQRCode":
			sp.OTPQRCode = &(o.OTPQRCode)
		case "OTPSecret":
			sp.OTPSecret = &(o.OTPSecret)
		case "SSHCA":
			sp.SSHCA = &(o.SSHCA)
		case "SSHCARenew":
			sp.SSHCARenew = &(o.SSHCARenew)
		case "accessEnabled":
			sp.AccessEnabled = &(o.AccessEnabled)
		case "activationExpiration":
			sp.ActivationExpiration = &(o.ActivationExpiration)
		case "activationToken":
			sp.ActivationToken = &(o.ActivationToken)
		case "associatedAPIAuthPolicyID":
			sp.AssociatedAPIAuthPolicyID = &(o.AssociatedAPIAuthPolicyID)
		case "associatedAWSPolicies":
			sp.AssociatedAWSPolicies = &(o.AssociatedAWSPolicies)
		case "associatedBillingID":
			sp.AssociatedBillingID = &(o.AssociatedBillingID)
		case "associatedNamespaceID":
			sp.AssociatedNamespaceID = &(o.AssociatedNamespaceID)
		case "associatedPlanKey":
			sp.AssociatedPlanKey = &(o.AssociatedPlanKey)
		case "associatedQuotaPolicies":
			sp.AssociatedQuotaPolicies = &(o.AssociatedQuotaPolicies)
		case "company":
			sp.Company = &(o.Company)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "email":
			sp.Email = &(o.Email)
		case "failedAuth":
			sp.FailedAuth = &(o.FailedAuth)
		case "failedTime":
			sp.FailedTime = &(o.FailedTime)
		case "firstName":
			sp.FirstName = &(o.FirstName)
		case "lastName":
			sp.LastName = &(o.LastName)
		case "localCA":
			sp.LocalCA = &(o.LocalCA)
		case "localCARenew":
			sp.LocalCARenew = &(o.LocalCARenew)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "newPassword":
			sp.NewPassword = &(o.NewPassword)
		case "password":
			sp.Password = &(o.Password)
		case "reCAPTCHAKey":
			sp.ReCAPTCHAKey = &(o.ReCAPTCHAKey)
		case "resetPasswordExpiration":
			sp.ResetPasswordExpiration = &(o.ResetPasswordExpiration)
		case "resetPasswordToken":
			sp.ResetPasswordToken = &(o.ResetPasswordToken)
		case "status":
			sp.Status = &(o.Status)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAccount to the object.
func (o *Account) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAccount)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.OTPEnabled != nil {
		o.OTPEnabled = *so.OTPEnabled
	}
	if so.OTPQRCode != nil {
		o.OTPQRCode = *so.OTPQRCode
	}
	if so.OTPSecret != nil {
		o.OTPSecret = *so.OTPSecret
	}
	if so.SSHCA != nil {
		o.SSHCA = *so.SSHCA
	}
	if so.SSHCARenew != nil {
		o.SSHCARenew = *so.SSHCARenew
	}
	if so.AccessEnabled != nil {
		o.AccessEnabled = *so.AccessEnabled
	}
	if so.ActivationExpiration != nil {
		o.ActivationExpiration = *so.ActivationExpiration
	}
	if so.ActivationToken != nil {
		o.ActivationToken = *so.ActivationToken
	}
	if so.AssociatedAPIAuthPolicyID != nil {
		o.AssociatedAPIAuthPolicyID = *so.AssociatedAPIAuthPolicyID
	}
	if so.AssociatedAWSPolicies != nil {
		o.AssociatedAWSPolicies = *so.AssociatedAWSPolicies
	}
	if so.AssociatedBillingID != nil {
		o.AssociatedBillingID = *so.AssociatedBillingID
	}
	if so.AssociatedNamespaceID != nil {
		o.AssociatedNamespaceID = *so.AssociatedNamespaceID
	}
	if so.AssociatedPlanKey != nil {
		o.AssociatedPlanKey = *so.AssociatedPlanKey
	}
	if so.AssociatedQuotaPolicies != nil {
		o.AssociatedQuotaPolicies = *so.AssociatedQuotaPolicies
	}
	if so.Company != nil {
		o.Company = *so.Company
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Email != nil {
		o.Email = *so.Email
	}
	if so.FailedAuth != nil {
		o.FailedAuth = *so.FailedAuth
	}
	if so.FailedTime != nil {
		o.FailedTime = *so.FailedTime
	}
	if so.FirstName != nil {
		o.FirstName = *so.FirstName
	}
	if so.LastName != nil {
		o.LastName = *so.LastName
	}
	if so.LocalCA != nil {
		o.LocalCA = *so.LocalCA
	}
	if so.LocalCARenew != nil {
		o.LocalCARenew = *so.LocalCARenew
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.NewPassword != nil {
		o.NewPassword = *so.NewPassword
	}
	if so.Password != nil {
		o.Password = *so.Password
	}
	if so.ReCAPTCHAKey != nil {
		o.ReCAPTCHAKey = *so.ReCAPTCHAKey
	}
	if so.ResetPasswordExpiration != nil {
		o.ResetPasswordExpiration = *so.ResetPasswordExpiration
	}
	if so.ResetPasswordToken != nil {
		o.ResetPasswordToken = *so.ResetPasswordToken
	}
	if so.Status != nil {
		o.Status = *so.Status
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the Account.
func (o *Account) DeepCopy() *Account {

	if o == nil {
		return nil
	}

	out := &Account{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Account.
func (o *Account) DeepCopyInto(out *Account) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Account: %s", err))
	}

	*out = *target.(*Account)
}

// Validate valides the current information stored into the structure.
func (o *Account) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("email", o.Email); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidatePattern("name", o.Name, `^[^\*\=]*$`, `must not contain any '*' or '='`, true); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Disabled", "Invited", "Pending"}, true); err != nil {
		errors = errors.Append(err)
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
func (*Account) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AccountAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AccountLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Account) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AccountAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Account) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "OTPEnabled":
		return o.OTPEnabled
	case "OTPQRCode":
		return o.OTPQRCode
	case "OTPSecret":
		return o.OTPSecret
	case "SSHCA":
		return o.SSHCA
	case "SSHCARenew":
		return o.SSHCARenew
	case "accessEnabled":
		return o.AccessEnabled
	case "activationExpiration":
		return o.ActivationExpiration
	case "activationToken":
		return o.ActivationToken
	case "associatedAPIAuthPolicyID":
		return o.AssociatedAPIAuthPolicyID
	case "associatedAWSPolicies":
		return o.AssociatedAWSPolicies
	case "associatedBillingID":
		return o.AssociatedBillingID
	case "associatedNamespaceID":
		return o.AssociatedNamespaceID
	case "associatedPlanKey":
		return o.AssociatedPlanKey
	case "associatedQuotaPolicies":
		return o.AssociatedQuotaPolicies
	case "company":
		return o.Company
	case "createTime":
		return o.CreateTime
	case "email":
		return o.Email
	case "failedAuth":
		return o.FailedAuth
	case "failedTime":
		return o.FailedTime
	case "firstName":
		return o.FirstName
	case "lastName":
		return o.LastName
	case "localCA":
		return o.LocalCA
	case "localCARenew":
		return o.LocalCARenew
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "newPassword":
		return o.NewPassword
	case "password":
		return o.Password
	case "reCAPTCHAKey":
		return o.ReCAPTCHAKey
	case "resetPasswordExpiration":
		return o.ResetPasswordExpiration
	case "resetPasswordToken":
		return o.ResetPasswordToken
	case "status":
		return o.Status
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// AccountAttributesMap represents the map of attribute for Account.
var AccountAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"OTPEnabled": {
		AllowedChoices: []string{},
		ConvertedName:  "OTPEnabled",
		Description:    `Enable or disable two-factor authentication.`,
		Exposed:        true,
		Name:           "OTPEnabled",
		Stored:         true,
		Type:           "boolean",
	},
	"OTPQRCode": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "OTPQRCode",
		Description:    `Returns the base64-encoded QR code for setting up two-factor authentication.`,
		Exposed:        true,
		Name:           "OTPQRCode",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"OTPSecret": {
		AllowedChoices: []string{},
		ConvertedName:  "OTPSecret",
		Description:    `Stores the two-factor authentication secret.`,
		Name:           "OTPSecret",
		Stored:         true,
		Type:           "string",
	},
	"SSHCA": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SSHCA",
		Description:    `Holds the SSH certificate authority used by the account namespace.`,
		Exposed:        true,
		Name:           "SSHCA",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"SSHCARenew": {
		AllowedChoices: []string{},
		ConvertedName:  "SSHCARenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the SSH certificate authority of the account namespace.`,
		Exposed:        true,
		Name:           "SSHCARenew",
		Type:           "boolean",
	},
	"AccessEnabled": {
		AllowedChoices: []string{},
		ConvertedName:  "AccessEnabled",
		Description:    `Defines if the account holder should have access to the system.`,
		Exposed:        true,
		Name:           "accessEnabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"ActivationExpiration": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ActivationExpiration",
		Description:    `Contains the expiration date of the activation token.`,
		Name:           "activationExpiration",
		Stored:         true,
		Type:           "time",
	},
	"ActivationToken": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ActivationToken",
		Description:    `Contains the activation token.`,
		Exposed:        true,
		Name:           "activationToken",
		Stored:         true,
		Type:           "string",
	},
	"AssociatedAPIAuthPolicyID": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedAPIAuthPolicyID",
		Description:    `Holds the ID of the associated API authorization.`,
		Name:           "associatedAPIAuthPolicyID",
		Stored:         true,
		Type:           "string",
	},
	"AssociatedAWSPolicies": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedAWSPolicies",
		Description:    `Contains a map of associated AWS defender policies.`,
		Name:           "associatedAWSPolicies",
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"AssociatedBillingID": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedBillingID",
		Description:    `Holds the ID of the associated billing customer.`,
		Exposed:        true,
		Name:           "associatedBillingID",
		Stored:         true,
		Type:           "string",
	},
	"AssociatedNamespaceID": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedNamespaceID",
		Description:    `Contains the ID of the associated namespace.`,
		Name:           "associatedNamespaceID",
		Stored:         true,
		Type:           "string",
	},
	"AssociatedPlanKey": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedPlanKey",
		CreationOnly:   true,
		Description:    `Contains the plan key associated with this account.`,
		Exposed:        true,
		Name:           "associatedPlanKey",
		Stored:         true,
		Type:           "string",
	},
	"AssociatedQuotaPolicies": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedQuotaPolicies",
		Description:    `Contains a map of the associated quotas.`,
		Name:           "associatedQuotaPolicies",
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Company": {
		AllowedChoices: []string{},
		ConvertedName:  "Company",
		Description:    `Company of the account user.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "company",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Email": {
		AllowedChoices: []string{},
		ConvertedName:  "Email",
		Description:    `Email of the account holder.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "email",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"FailedAuth": {
		AllowedChoices: []string{},
		ConvertedName:  "FailedAuth",
		Description:    `Internally keeps track of the number of failed attempt.`,
		Name:           "failedAuth",
		Stored:         true,
		Type:           "integer",
	},
	"FailedTime": {
		AllowedChoices: []string{},
		ConvertedName:  "FailedTime",
		Description:    `Internally keeps track of the time of the last failed attempt.`,
		Name:           "failedTime",
		Stored:         true,
		Type:           "time",
	},
	"FirstName": {
		AllowedChoices: []string{},
		ConvertedName:  "FirstName",
		Description:    `First name of the account user.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "firstName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LastName": {
		AllowedChoices: []string{},
		ConvertedName:  "LastName",
		Description:    `Last name of the account user.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "lastName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LocalCA": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LocalCA",
		Description:    `The certificate authority used by this namespace.`,
		Exposed:        true,
		Name:           "localCA",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"LocalCARenew": {
		AllowedChoices: []string{},
		ConvertedName:  "LocalCARenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the local certificate authority of the account namespace.`,
		Exposed:        true,
		Name:           "localCARenew",
		Type:           "boolean",
	},
	"MigrationsLog": {
		AllowedChoices: []string{},
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"Name": {
		AllowedChars:   `^[^\*\=]*$`,
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		CreationOnly:   true,
		Description:    `Name of the account.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"NewPassword": {
		AllowedChoices: []string{},
		ConvertedName:  "NewPassword",
		Description: `New password for the account. If set the previous password must be given through
the property ` + "`" + `password` + "`" + `.`,
		Exposed: true,
		Name:    "newPassword",
		Type:    "string",
	},
	"Password": {
		AllowedChoices: []string{},
		ConvertedName:  "Password",
		Description:    `Password for the account.`,
		Exposed:        true,
		Name:           "password",
		Stored:         true,
		Type:           "string",
	},
	"ReCAPTCHAKey": {
		AllowedChoices: []string{},
		ConvertedName:  "ReCAPTCHAKey",
		CreationOnly:   true,
		Description: `Contains the completely automated public Turing test (CAPTCHA)
validation if reCAPTCHA is enabled.`,
		Exposed: true,
		Name:    "reCAPTCHAKey",
		Type:    "string",
	},
	"ResetPasswordExpiration": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ResetPasswordExpiration",
		Description:    `Contains the expiration time for resetting the password.`,
		Name:           "resetPasswordExpiration",
		Stored:         true,
		Type:           "time",
	},
	"ResetPasswordToken": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ResetPasswordToken",
		Description:    `Contains the token to use for resetting password.`,
		Name:           "resetPasswordToken",
		Stored:         true,
		Type:           "string",
	},
	"Status": {
		AllowedChoices: []string{"Active", "Disabled", "Invited", "Pending"},
		Autogenerated:  true,
		ConvertedName:  "Status",
		DefaultValue:   AccountStatusPending,
		Description:    `Status of the account.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
	"UpdateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"ZHash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// AccountLowerCaseAttributesMap represents the map of attribute for Account.
var AccountLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"otpenabled": {
		AllowedChoices: []string{},
		ConvertedName:  "OTPEnabled",
		Description:    `Enable or disable two-factor authentication.`,
		Exposed:        true,
		Name:           "OTPEnabled",
		Stored:         true,
		Type:           "boolean",
	},
	"otpqrcode": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "OTPQRCode",
		Description:    `Returns the base64-encoded QR code for setting up two-factor authentication.`,
		Exposed:        true,
		Name:           "OTPQRCode",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"otpsecret": {
		AllowedChoices: []string{},
		ConvertedName:  "OTPSecret",
		Description:    `Stores the two-factor authentication secret.`,
		Name:           "OTPSecret",
		Stored:         true,
		Type:           "string",
	},
	"sshca": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SSHCA",
		Description:    `Holds the SSH certificate authority used by the account namespace.`,
		Exposed:        true,
		Name:           "SSHCA",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"sshcarenew": {
		AllowedChoices: []string{},
		ConvertedName:  "SSHCARenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the SSH certificate authority of the account namespace.`,
		Exposed:        true,
		Name:           "SSHCARenew",
		Type:           "boolean",
	},
	"accessenabled": {
		AllowedChoices: []string{},
		ConvertedName:  "AccessEnabled",
		Description:    `Defines if the account holder should have access to the system.`,
		Exposed:        true,
		Name:           "accessEnabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"activationexpiration": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ActivationExpiration",
		Description:    `Contains the expiration date of the activation token.`,
		Name:           "activationExpiration",
		Stored:         true,
		Type:           "time",
	},
	"activationtoken": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ActivationToken",
		Description:    `Contains the activation token.`,
		Exposed:        true,
		Name:           "activationToken",
		Stored:         true,
		Type:           "string",
	},
	"associatedapiauthpolicyid": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedAPIAuthPolicyID",
		Description:    `Holds the ID of the associated API authorization.`,
		Name:           "associatedAPIAuthPolicyID",
		Stored:         true,
		Type:           "string",
	},
	"associatedawspolicies": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedAWSPolicies",
		Description:    `Contains a map of associated AWS defender policies.`,
		Name:           "associatedAWSPolicies",
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"associatedbillingid": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedBillingID",
		Description:    `Holds the ID of the associated billing customer.`,
		Exposed:        true,
		Name:           "associatedBillingID",
		Stored:         true,
		Type:           "string",
	},
	"associatednamespaceid": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedNamespaceID",
		Description:    `Contains the ID of the associated namespace.`,
		Name:           "associatedNamespaceID",
		Stored:         true,
		Type:           "string",
	},
	"associatedplankey": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedPlanKey",
		CreationOnly:   true,
		Description:    `Contains the plan key associated with this account.`,
		Exposed:        true,
		Name:           "associatedPlanKey",
		Stored:         true,
		Type:           "string",
	},
	"associatedquotapolicies": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedQuotaPolicies",
		Description:    `Contains a map of the associated quotas.`,
		Name:           "associatedQuotaPolicies",
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"company": {
		AllowedChoices: []string{},
		ConvertedName:  "Company",
		Description:    `Company of the account user.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "company",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"email": {
		AllowedChoices: []string{},
		ConvertedName:  "Email",
		Description:    `Email of the account holder.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "email",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"failedauth": {
		AllowedChoices: []string{},
		ConvertedName:  "FailedAuth",
		Description:    `Internally keeps track of the number of failed attempt.`,
		Name:           "failedAuth",
		Stored:         true,
		Type:           "integer",
	},
	"failedtime": {
		AllowedChoices: []string{},
		ConvertedName:  "FailedTime",
		Description:    `Internally keeps track of the time of the last failed attempt.`,
		Name:           "failedTime",
		Stored:         true,
		Type:           "time",
	},
	"firstname": {
		AllowedChoices: []string{},
		ConvertedName:  "FirstName",
		Description:    `First name of the account user.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "firstName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"lastname": {
		AllowedChoices: []string{},
		ConvertedName:  "LastName",
		Description:    `Last name of the account user.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "lastName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"localca": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LocalCA",
		Description:    `The certificate authority used by this namespace.`,
		Exposed:        true,
		Name:           "localCA",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"localcarenew": {
		AllowedChoices: []string{},
		ConvertedName:  "LocalCARenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the local certificate authority of the account namespace.`,
		Exposed:        true,
		Name:           "localCARenew",
		Type:           "boolean",
	},
	"migrationslog": {
		AllowedChoices: []string{},
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"name": {
		AllowedChars:   `^[^\*\=]*$`,
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		CreationOnly:   true,
		Description:    `Name of the account.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"newpassword": {
		AllowedChoices: []string{},
		ConvertedName:  "NewPassword",
		Description: `New password for the account. If set the previous password must be given through
the property ` + "`" + `password` + "`" + `.`,
		Exposed: true,
		Name:    "newPassword",
		Type:    "string",
	},
	"password": {
		AllowedChoices: []string{},
		ConvertedName:  "Password",
		Description:    `Password for the account.`,
		Exposed:        true,
		Name:           "password",
		Stored:         true,
		Type:           "string",
	},
	"recaptchakey": {
		AllowedChoices: []string{},
		ConvertedName:  "ReCAPTCHAKey",
		CreationOnly:   true,
		Description: `Contains the completely automated public Turing test (CAPTCHA)
validation if reCAPTCHA is enabled.`,
		Exposed: true,
		Name:    "reCAPTCHAKey",
		Type:    "string",
	},
	"resetpasswordexpiration": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ResetPasswordExpiration",
		Description:    `Contains the expiration time for resetting the password.`,
		Name:           "resetPasswordExpiration",
		Stored:         true,
		Type:           "time",
	},
	"resetpasswordtoken": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ResetPasswordToken",
		Description:    `Contains the token to use for resetting password.`,
		Name:           "resetPasswordToken",
		Stored:         true,
		Type:           "string",
	},
	"status": {
		AllowedChoices: []string{"Active", "Disabled", "Invited", "Pending"},
		Autogenerated:  true,
		ConvertedName:  "Status",
		DefaultValue:   AccountStatusPending,
		Description:    `Status of the account.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
	"updatetime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"zhash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseAccountsList represents a list of SparseAccounts
type SparseAccountsList []*SparseAccount

// Identity returns the identity of the objects in the list.
func (o SparseAccountsList) Identity() elemental.Identity {

	return AccountIdentity
}

// Copy returns a pointer to a copy the SparseAccountsList.
func (o SparseAccountsList) Copy() elemental.Identifiables {

	copy := append(SparseAccountsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAccountsList.
func (o SparseAccountsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAccountsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAccount))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAccountsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAccountsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseAccountsList converted to AccountsList.
func (o SparseAccountsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAccountsList) Version() int {

	return 1
}

// SparseAccount represents the sparse version of a account.
type SparseAccount struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Enable or disable two-factor authentication.
	OTPEnabled *bool `json:"OTPEnabled,omitempty" msgpack:"OTPEnabled,omitempty" bson:"otpenabled,omitempty" mapstructure:"OTPEnabled,omitempty"`

	// Returns the base64-encoded QR code for setting up two-factor authentication.
	OTPQRCode *string `json:"OTPQRCode,omitempty" msgpack:"OTPQRCode,omitempty" bson:"-" mapstructure:"OTPQRCode,omitempty"`

	// Stores the two-factor authentication secret.
	OTPSecret *string `json:"-" msgpack:"-" bson:"otpsecret,omitempty" mapstructure:"-,omitempty"`

	// Holds the SSH certificate authority used by the account namespace.
	SSHCA *string `json:"SSHCA,omitempty" msgpack:"SSHCA,omitempty" bson:"-" mapstructure:"SSHCA,omitempty"`

	// Set to `true` to renew the SSH certificate authority of the account namespace.
	SSHCARenew *bool `json:"SSHCARenew,omitempty" msgpack:"SSHCARenew,omitempty" bson:"-" mapstructure:"SSHCARenew,omitempty"`

	// Defines if the account holder should have access to the system.
	AccessEnabled *bool `json:"accessEnabled,omitempty" msgpack:"accessEnabled,omitempty" bson:"accessenabled,omitempty" mapstructure:"accessEnabled,omitempty"`

	// Contains the expiration date of the activation token.
	ActivationExpiration *time.Time `json:"-" msgpack:"-" bson:"activationexpiration,omitempty" mapstructure:"-,omitempty"`

	// Contains the activation token.
	ActivationToken *string `json:"activationToken,omitempty" msgpack:"activationToken,omitempty" bson:"activationtoken,omitempty" mapstructure:"activationToken,omitempty"`

	// Holds the ID of the associated API authorization.
	AssociatedAPIAuthPolicyID *string `json:"-" msgpack:"-" bson:"associatedapiauthpolicyid,omitempty" mapstructure:"-,omitempty"`

	// Contains a map of associated AWS defender policies.
	AssociatedAWSPolicies *map[string]string `json:"-" msgpack:"-" bson:"associatedawspolicies,omitempty" mapstructure:"-,omitempty"`

	// Holds the ID of the associated billing customer.
	AssociatedBillingID *string `json:"associatedBillingID,omitempty" msgpack:"associatedBillingID,omitempty" bson:"associatedbillingid,omitempty" mapstructure:"associatedBillingID,omitempty"`

	// Contains the ID of the associated namespace.
	AssociatedNamespaceID *string `json:"-" msgpack:"-" bson:"associatednamespaceid,omitempty" mapstructure:"-,omitempty"`

	// Contains the plan key associated with this account.
	AssociatedPlanKey *string `json:"associatedPlanKey,omitempty" msgpack:"associatedPlanKey,omitempty" bson:"associatedplankey,omitempty" mapstructure:"associatedPlanKey,omitempty"`

	// Contains a map of the associated quotas.
	AssociatedQuotaPolicies *map[string]string `json:"-" msgpack:"-" bson:"associatedquotapolicies,omitempty" mapstructure:"-,omitempty"`

	// Company of the account user.
	Company *string `json:"company,omitempty" msgpack:"company,omitempty" bson:"company,omitempty" mapstructure:"company,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Email of the account holder.
	Email *string `json:"email,omitempty" msgpack:"email,omitempty" bson:"email,omitempty" mapstructure:"email,omitempty"`

	// Internally keeps track of the number of failed attempt.
	FailedAuth *int `json:"-" msgpack:"-" bson:"failedauth,omitempty" mapstructure:"-,omitempty"`

	// Internally keeps track of the time of the last failed attempt.
	FailedTime *time.Time `json:"-" msgpack:"-" bson:"failedtime,omitempty" mapstructure:"-,omitempty"`

	// First name of the account user.
	FirstName *string `json:"firstName,omitempty" msgpack:"firstName,omitempty" bson:"firstname,omitempty" mapstructure:"firstName,omitempty"`

	// Last name of the account user.
	LastName *string `json:"lastName,omitempty" msgpack:"lastName,omitempty" bson:"lastname,omitempty" mapstructure:"lastName,omitempty"`

	// The certificate authority used by this namespace.
	LocalCA *string `json:"localCA,omitempty" msgpack:"localCA,omitempty" bson:"-" mapstructure:"localCA,omitempty"`

	// Set to `true` to renew the local certificate authority of the account namespace.
	LocalCARenew *bool `json:"localCARenew,omitempty" msgpack:"localCARenew,omitempty" bson:"-" mapstructure:"localCARenew,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the account.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// New password for the account. If set the previous password must be given through
	// the property `password`.
	NewPassword *string `json:"newPassword,omitempty" msgpack:"newPassword,omitempty" bson:"-" mapstructure:"newPassword,omitempty"`

	// Password for the account.
	Password *string `json:"password,omitempty" msgpack:"password,omitempty" bson:"password,omitempty" mapstructure:"password,omitempty"`

	// Contains the completely automated public Turing test (CAPTCHA)
	// validation if reCAPTCHA is enabled.
	ReCAPTCHAKey *string `json:"reCAPTCHAKey,omitempty" msgpack:"reCAPTCHAKey,omitempty" bson:"-" mapstructure:"reCAPTCHAKey,omitempty"`

	// Contains the expiration time for resetting the password.
	ResetPasswordExpiration *time.Time `json:"-" msgpack:"-" bson:"resetpasswordexpiration,omitempty" mapstructure:"-,omitempty"`

	// Contains the token to use for resetting password.
	ResetPasswordToken *string `json:"-" msgpack:"-" bson:"resetpasswordtoken,omitempty" mapstructure:"-,omitempty"`

	// Status of the account.
	Status *AccountStatusValue `json:"status,omitempty" msgpack:"status,omitempty" bson:"status,omitempty" mapstructure:"status,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAccount returns a new  SparseAccount.
func NewSparseAccount() *SparseAccount {
	return &SparseAccount{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAccount) Identity() elemental.Identity {

	return AccountIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAccount) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAccount) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAccount) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseAccount{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.OTPEnabled != nil {
		s.OTPEnabled = o.OTPEnabled
	}
	if o.OTPSecret != nil {
		s.OTPSecret = o.OTPSecret
	}
	if o.AccessEnabled != nil {
		s.AccessEnabled = o.AccessEnabled
	}
	if o.ActivationExpiration != nil {
		s.ActivationExpiration = o.ActivationExpiration
	}
	if o.ActivationToken != nil {
		s.ActivationToken = o.ActivationToken
	}
	if o.AssociatedAPIAuthPolicyID != nil {
		s.AssociatedAPIAuthPolicyID = o.AssociatedAPIAuthPolicyID
	}
	if o.AssociatedAWSPolicies != nil {
		s.AssociatedAWSPolicies = o.AssociatedAWSPolicies
	}
	if o.AssociatedBillingID != nil {
		s.AssociatedBillingID = o.AssociatedBillingID
	}
	if o.AssociatedNamespaceID != nil {
		s.AssociatedNamespaceID = o.AssociatedNamespaceID
	}
	if o.AssociatedPlanKey != nil {
		s.AssociatedPlanKey = o.AssociatedPlanKey
	}
	if o.AssociatedQuotaPolicies != nil {
		s.AssociatedQuotaPolicies = o.AssociatedQuotaPolicies
	}
	if o.Company != nil {
		s.Company = o.Company
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.Email != nil {
		s.Email = o.Email
	}
	if o.FailedAuth != nil {
		s.FailedAuth = o.FailedAuth
	}
	if o.FailedTime != nil {
		s.FailedTime = o.FailedTime
	}
	if o.FirstName != nil {
		s.FirstName = o.FirstName
	}
	if o.LastName != nil {
		s.LastName = o.LastName
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Password != nil {
		s.Password = o.Password
	}
	if o.ResetPasswordExpiration != nil {
		s.ResetPasswordExpiration = o.ResetPasswordExpiration
	}
	if o.ResetPasswordToken != nil {
		s.ResetPasswordToken = o.ResetPasswordToken
	}
	if o.Status != nil {
		s.Status = o.Status
	}
	if o.UpdateTime != nil {
		s.UpdateTime = o.UpdateTime
	}
	if o.ZHash != nil {
		s.ZHash = o.ZHash
	}
	if o.Zone != nil {
		s.Zone = o.Zone
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAccount) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseAccount{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.OTPEnabled != nil {
		o.OTPEnabled = s.OTPEnabled
	}
	if s.OTPSecret != nil {
		o.OTPSecret = s.OTPSecret
	}
	if s.AccessEnabled != nil {
		o.AccessEnabled = s.AccessEnabled
	}
	if s.ActivationExpiration != nil {
		o.ActivationExpiration = s.ActivationExpiration
	}
	if s.ActivationToken != nil {
		o.ActivationToken = s.ActivationToken
	}
	if s.AssociatedAPIAuthPolicyID != nil {
		o.AssociatedAPIAuthPolicyID = s.AssociatedAPIAuthPolicyID
	}
	if s.AssociatedAWSPolicies != nil {
		o.AssociatedAWSPolicies = s.AssociatedAWSPolicies
	}
	if s.AssociatedBillingID != nil {
		o.AssociatedBillingID = s.AssociatedBillingID
	}
	if s.AssociatedNamespaceID != nil {
		o.AssociatedNamespaceID = s.AssociatedNamespaceID
	}
	if s.AssociatedPlanKey != nil {
		o.AssociatedPlanKey = s.AssociatedPlanKey
	}
	if s.AssociatedQuotaPolicies != nil {
		o.AssociatedQuotaPolicies = s.AssociatedQuotaPolicies
	}
	if s.Company != nil {
		o.Company = s.Company
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.Email != nil {
		o.Email = s.Email
	}
	if s.FailedAuth != nil {
		o.FailedAuth = s.FailedAuth
	}
	if s.FailedTime != nil {
		o.FailedTime = s.FailedTime
	}
	if s.FirstName != nil {
		o.FirstName = s.FirstName
	}
	if s.LastName != nil {
		o.LastName = s.LastName
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Password != nil {
		o.Password = s.Password
	}
	if s.ResetPasswordExpiration != nil {
		o.ResetPasswordExpiration = s.ResetPasswordExpiration
	}
	if s.ResetPasswordToken != nil {
		o.ResetPasswordToken = s.ResetPasswordToken
	}
	if s.Status != nil {
		o.Status = s.Status
	}
	if s.UpdateTime != nil {
		o.UpdateTime = s.UpdateTime
	}
	if s.ZHash != nil {
		o.ZHash = s.ZHash
	}
	if s.Zone != nil {
		o.Zone = s.Zone
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseAccount) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAccount) ToPlain() elemental.PlainIdentifiable {

	out := NewAccount()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.OTPEnabled != nil {
		out.OTPEnabled = *o.OTPEnabled
	}
	if o.OTPQRCode != nil {
		out.OTPQRCode = *o.OTPQRCode
	}
	if o.OTPSecret != nil {
		out.OTPSecret = *o.OTPSecret
	}
	if o.SSHCA != nil {
		out.SSHCA = *o.SSHCA
	}
	if o.SSHCARenew != nil {
		out.SSHCARenew = *o.SSHCARenew
	}
	if o.AccessEnabled != nil {
		out.AccessEnabled = *o.AccessEnabled
	}
	if o.ActivationExpiration != nil {
		out.ActivationExpiration = *o.ActivationExpiration
	}
	if o.ActivationToken != nil {
		out.ActivationToken = *o.ActivationToken
	}
	if o.AssociatedAPIAuthPolicyID != nil {
		out.AssociatedAPIAuthPolicyID = *o.AssociatedAPIAuthPolicyID
	}
	if o.AssociatedAWSPolicies != nil {
		out.AssociatedAWSPolicies = *o.AssociatedAWSPolicies
	}
	if o.AssociatedBillingID != nil {
		out.AssociatedBillingID = *o.AssociatedBillingID
	}
	if o.AssociatedNamespaceID != nil {
		out.AssociatedNamespaceID = *o.AssociatedNamespaceID
	}
	if o.AssociatedPlanKey != nil {
		out.AssociatedPlanKey = *o.AssociatedPlanKey
	}
	if o.AssociatedQuotaPolicies != nil {
		out.AssociatedQuotaPolicies = *o.AssociatedQuotaPolicies
	}
	if o.Company != nil {
		out.Company = *o.Company
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Email != nil {
		out.Email = *o.Email
	}
	if o.FailedAuth != nil {
		out.FailedAuth = *o.FailedAuth
	}
	if o.FailedTime != nil {
		out.FailedTime = *o.FailedTime
	}
	if o.FirstName != nil {
		out.FirstName = *o.FirstName
	}
	if o.LastName != nil {
		out.LastName = *o.LastName
	}
	if o.LocalCA != nil {
		out.LocalCA = *o.LocalCA
	}
	if o.LocalCARenew != nil {
		out.LocalCARenew = *o.LocalCARenew
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.NewPassword != nil {
		out.NewPassword = *o.NewPassword
	}
	if o.Password != nil {
		out.Password = *o.Password
	}
	if o.ReCAPTCHAKey != nil {
		out.ReCAPTCHAKey = *o.ReCAPTCHAKey
	}
	if o.ResetPasswordExpiration != nil {
		out.ResetPasswordExpiration = *o.ResetPasswordExpiration
	}
	if o.ResetPasswordToken != nil {
		out.ResetPasswordToken = *o.ResetPasswordToken
	}
	if o.Status != nil {
		out.Status = *o.Status
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseAccount) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseAccount) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseAccount) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseAccount) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseAccount) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseAccount) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseAccount) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseAccount) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseAccount) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseAccount) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseAccount.
func (o *SparseAccount) DeepCopy() *SparseAccount {

	if o == nil {
		return nil
	}

	out := &SparseAccount{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAccount.
func (o *SparseAccount) DeepCopyInto(out *SparseAccount) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAccount: %s", err))
	}

	*out = *target.(*SparseAccount)
}

type mongoAttributesAccount struct {
	ID                        bson.ObjectId      `bson:"_id,omitempty"`
	OTPEnabled                bool               `bson:"otpenabled"`
	OTPSecret                 string             `bson:"otpsecret"`
	AccessEnabled             bool               `bson:"accessenabled"`
	ActivationExpiration      time.Time          `bson:"activationexpiration"`
	ActivationToken           string             `bson:"activationtoken,omitempty"`
	AssociatedAPIAuthPolicyID string             `bson:"associatedapiauthpolicyid"`
	AssociatedAWSPolicies     map[string]string  `bson:"associatedawspolicies"`
	AssociatedBillingID       string             `bson:"associatedbillingid"`
	AssociatedNamespaceID     string             `bson:"associatednamespaceid"`
	AssociatedPlanKey         string             `bson:"associatedplankey"`
	AssociatedQuotaPolicies   map[string]string  `bson:"associatedquotapolicies"`
	Company                   string             `bson:"company"`
	CreateTime                time.Time          `bson:"createtime"`
	Email                     string             `bson:"email"`
	FailedAuth                int                `bson:"failedauth"`
	FailedTime                time.Time          `bson:"failedtime"`
	FirstName                 string             `bson:"firstname"`
	LastName                  string             `bson:"lastname"`
	MigrationsLog             map[string]string  `bson:"migrationslog,omitempty"`
	Name                      string             `bson:"name"`
	Password                  string             `bson:"password"`
	ResetPasswordExpiration   time.Time          `bson:"resetpasswordexpiration"`
	ResetPasswordToken        string             `bson:"resetpasswordtoken"`
	Status                    AccountStatusValue `bson:"status"`
	UpdateTime                time.Time          `bson:"updatetime"`
	ZHash                     int                `bson:"zhash"`
	Zone                      int                `bson:"zone"`
}
type mongoAttributesSparseAccount struct {
	ID                        bson.ObjectId       `bson:"_id,omitempty"`
	OTPEnabled                *bool               `bson:"otpenabled,omitempty"`
	OTPSecret                 *string             `bson:"otpsecret,omitempty"`
	AccessEnabled             *bool               `bson:"accessenabled,omitempty"`
	ActivationExpiration      *time.Time          `bson:"activationexpiration,omitempty"`
	ActivationToken           *string             `bson:"activationtoken,omitempty"`
	AssociatedAPIAuthPolicyID *string             `bson:"associatedapiauthpolicyid,omitempty"`
	AssociatedAWSPolicies     *map[string]string  `bson:"associatedawspolicies,omitempty"`
	AssociatedBillingID       *string             `bson:"associatedbillingid,omitempty"`
	AssociatedNamespaceID     *string             `bson:"associatednamespaceid,omitempty"`
	AssociatedPlanKey         *string             `bson:"associatedplankey,omitempty"`
	AssociatedQuotaPolicies   *map[string]string  `bson:"associatedquotapolicies,omitempty"`
	Company                   *string             `bson:"company,omitempty"`
	CreateTime                *time.Time          `bson:"createtime,omitempty"`
	Email                     *string             `bson:"email,omitempty"`
	FailedAuth                *int                `bson:"failedauth,omitempty"`
	FailedTime                *time.Time          `bson:"failedtime,omitempty"`
	FirstName                 *string             `bson:"firstname,omitempty"`
	LastName                  *string             `bson:"lastname,omitempty"`
	MigrationsLog             *map[string]string  `bson:"migrationslog,omitempty"`
	Name                      *string             `bson:"name,omitempty"`
	Password                  *string             `bson:"password,omitempty"`
	ResetPasswordExpiration   *time.Time          `bson:"resetpasswordexpiration,omitempty"`
	ResetPasswordToken        *string             `bson:"resetpasswordtoken,omitempty"`
	Status                    *AccountStatusValue `bson:"status,omitempty"`
	UpdateTime                *time.Time          `bson:"updatetime,omitempty"`
	ZHash                     *int                `bson:"zhash,omitempty"`
	Zone                      *int                `bson:"zone,omitempty"`
}
