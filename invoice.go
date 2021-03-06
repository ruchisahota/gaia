package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// InvoiceBilledToProviderValue represents the possible values for attribute "billedToProvider".
type InvoiceBilledToProviderValue string

const (
	// InvoiceBilledToProviderAWS represents the value AWS.
	InvoiceBilledToProviderAWS InvoiceBilledToProviderValue = "AWS"

	// InvoiceBilledToProviderAporeto represents the value Aporeto.
	InvoiceBilledToProviderAporeto InvoiceBilledToProviderValue = "Aporeto"
)

// InvoiceIdentity represents the Identity of the object.
var InvoiceIdentity = elemental.Identity{
	Name:     "invoice",
	Category: "invoices",
	Package:  "bill",
	Private:  false,
}

// InvoicesList represents a list of Invoices
type InvoicesList []*Invoice

// Identity returns the identity of the objects in the list.
func (o InvoicesList) Identity() elemental.Identity {

	return InvoiceIdentity
}

// Copy returns a pointer to a copy the InvoicesList.
func (o InvoicesList) Copy() elemental.Identifiables {

	copy := append(InvoicesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the InvoicesList.
func (o InvoicesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(InvoicesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Invoice))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o InvoicesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o InvoicesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the InvoicesList converted to SparseInvoicesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o InvoicesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseInvoicesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseInvoice)
	}

	return out
}

// Version returns the version of the content.
func (o InvoicesList) Version() int {

	return 1
}

// Invoice represents the model of a invoice
type Invoice struct {
	// The ID of the invoice.
	ID string `json:"ID" msgpack:"ID" bson:"id" mapstructure:"ID,omitempty"`

	// The ID of the customer that this invoice belongs to.
	AccountID string `json:"accountID" msgpack:"accountID" bson:"accountid" mapstructure:"accountID,omitempty"`

	// The name of the provider that this invoice was billed to.
	BilledToProvider InvoiceBilledToProviderValue `json:"billedToProvider" msgpack:"billedToProvider" bson:"billedtoprovider" mapstructure:"billedToProvider,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// The end date of the invoice.
	EndDate time.Time `json:"endDate" msgpack:"endDate" bson:"enddate" mapstructure:"endDate,omitempty"`

	// The start date of this invoice.
	StartDate time.Time `json:"startDate" msgpack:"startDate" bson:"startdate" mapstructure:"startDate,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewInvoice returns a new *Invoice
func NewInvoice() *Invoice {

	return &Invoice{
		ModelVersion:     1,
		BilledToProvider: InvoiceBilledToProviderAporeto,
	}
}

// Identity returns the Identity of the object.
func (o *Invoice) Identity() elemental.Identity {

	return InvoiceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Invoice) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Invoice) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Invoice) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesInvoice{}

	s.ID = o.ID
	s.AccountID = o.AccountID
	s.BilledToProvider = o.BilledToProvider
	s.CreateTime = o.CreateTime
	s.EndDate = o.EndDate
	s.StartDate = o.StartDate
	s.UpdateTime = o.UpdateTime

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Invoice) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesInvoice{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID
	o.AccountID = s.AccountID
	o.BilledToProvider = s.BilledToProvider
	o.CreateTime = s.CreateTime
	o.EndDate = s.EndDate
	o.StartDate = s.StartDate
	o.UpdateTime = s.UpdateTime

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Invoice) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Invoice) BleveType() string {

	return "invoice"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Invoice) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Invoice) Doc() string {

	return `Provides access to Segment customer invoices.`
}

func (o *Invoice) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Invoice) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *Invoice) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Invoice) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *Invoice) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Invoice) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseInvoice{
			ID:               &o.ID,
			AccountID:        &o.AccountID,
			BilledToProvider: &o.BilledToProvider,
			CreateTime:       &o.CreateTime,
			EndDate:          &o.EndDate,
			StartDate:        &o.StartDate,
			UpdateTime:       &o.UpdateTime,
		}
	}

	sp := &SparseInvoice{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "accountID":
			sp.AccountID = &(o.AccountID)
		case "billedToProvider":
			sp.BilledToProvider = &(o.BilledToProvider)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "endDate":
			sp.EndDate = &(o.EndDate)
		case "startDate":
			sp.StartDate = &(o.StartDate)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseInvoice to the object.
func (o *Invoice) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseInvoice)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.AccountID != nil {
		o.AccountID = *so.AccountID
	}
	if so.BilledToProvider != nil {
		o.BilledToProvider = *so.BilledToProvider
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.EndDate != nil {
		o.EndDate = *so.EndDate
	}
	if so.StartDate != nil {
		o.StartDate = *so.StartDate
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
}

// DeepCopy returns a deep copy if the Invoice.
func (o *Invoice) DeepCopy() *Invoice {

	if o == nil {
		return nil
	}

	out := &Invoice{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Invoice.
func (o *Invoice) DeepCopyInto(out *Invoice) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Invoice: %s", err))
	}

	*out = *target.(*Invoice)
}

// Validate valides the current information stored into the structure.
func (o *Invoice) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("billedToProvider", string(o.BilledToProvider), []string{"Aporeto", "AWS"}, false); err != nil {
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
func (*Invoice) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := InvoiceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return InvoiceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Invoice) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return InvoiceAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Invoice) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "accountID":
		return o.AccountID
	case "billedToProvider":
		return o.BilledToProvider
	case "createTime":
		return o.CreateTime
	case "endDate":
		return o.EndDate
	case "startDate":
		return o.StartDate
	case "updateTime":
		return o.UpdateTime
	}

	return nil
}

// InvoiceAttributesMap represents the map of attribute for Invoice.
var InvoiceAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `The ID of the invoice.`,
		Exposed:        true,
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"AccountID": {
		AllowedChoices: []string{},
		ConvertedName:  "AccountID",
		Description:    `The ID of the customer that this invoice belongs to.`,
		Exposed:        true,
		Name:           "accountID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"BilledToProvider": {
		AllowedChoices: []string{"Aporeto", "AWS"},
		ConvertedName:  "BilledToProvider",
		DefaultValue:   InvoiceBilledToProviderAporeto,
		Description:    `The name of the provider that this invoice was billed to.`,
		Exposed:        true,
		Name:           "billedToProvider",
		Stored:         true,
		Type:           "enum",
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
	"EndDate": {
		AllowedChoices: []string{},
		ConvertedName:  "EndDate",
		Description:    `The end date of the invoice.`,
		Exposed:        true,
		Name:           "endDate",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"StartDate": {
		AllowedChoices: []string{},
		ConvertedName:  "StartDate",
		Description:    `The start date of this invoice.`,
		Exposed:        true,
		Name:           "startDate",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
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
}

// InvoiceLowerCaseAttributesMap represents the map of attribute for Invoice.
var InvoiceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `The ID of the invoice.`,
		Exposed:        true,
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"accountid": {
		AllowedChoices: []string{},
		ConvertedName:  "AccountID",
		Description:    `The ID of the customer that this invoice belongs to.`,
		Exposed:        true,
		Name:           "accountID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"billedtoprovider": {
		AllowedChoices: []string{"Aporeto", "AWS"},
		ConvertedName:  "BilledToProvider",
		DefaultValue:   InvoiceBilledToProviderAporeto,
		Description:    `The name of the provider that this invoice was billed to.`,
		Exposed:        true,
		Name:           "billedToProvider",
		Stored:         true,
		Type:           "enum",
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
	"enddate": {
		AllowedChoices: []string{},
		ConvertedName:  "EndDate",
		Description:    `The end date of the invoice.`,
		Exposed:        true,
		Name:           "endDate",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"startdate": {
		AllowedChoices: []string{},
		ConvertedName:  "StartDate",
		Description:    `The start date of this invoice.`,
		Exposed:        true,
		Name:           "startDate",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
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
}

// SparseInvoicesList represents a list of SparseInvoices
type SparseInvoicesList []*SparseInvoice

// Identity returns the identity of the objects in the list.
func (o SparseInvoicesList) Identity() elemental.Identity {

	return InvoiceIdentity
}

// Copy returns a pointer to a copy the SparseInvoicesList.
func (o SparseInvoicesList) Copy() elemental.Identifiables {

	copy := append(SparseInvoicesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseInvoicesList.
func (o SparseInvoicesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseInvoicesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseInvoice))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseInvoicesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseInvoicesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseInvoicesList converted to InvoicesList.
func (o SparseInvoicesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseInvoicesList) Version() int {

	return 1
}

// SparseInvoice represents the sparse version of a invoice.
type SparseInvoice struct {
	// The ID of the invoice.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"id,omitempty" mapstructure:"ID,omitempty"`

	// The ID of the customer that this invoice belongs to.
	AccountID *string `json:"accountID,omitempty" msgpack:"accountID,omitempty" bson:"accountid,omitempty" mapstructure:"accountID,omitempty"`

	// The name of the provider that this invoice was billed to.
	BilledToProvider *InvoiceBilledToProviderValue `json:"billedToProvider,omitempty" msgpack:"billedToProvider,omitempty" bson:"billedtoprovider,omitempty" mapstructure:"billedToProvider,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// The end date of the invoice.
	EndDate *time.Time `json:"endDate,omitempty" msgpack:"endDate,omitempty" bson:"enddate,omitempty" mapstructure:"endDate,omitempty"`

	// The start date of this invoice.
	StartDate *time.Time `json:"startDate,omitempty" msgpack:"startDate,omitempty" bson:"startdate,omitempty" mapstructure:"startDate,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseInvoice returns a new  SparseInvoice.
func NewSparseInvoice() *SparseInvoice {
	return &SparseInvoice{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseInvoice) Identity() elemental.Identity {

	return InvoiceIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseInvoice) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseInvoice) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseInvoice) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseInvoice{}

	if o.ID != nil {
		s.ID = o.ID
	}
	if o.AccountID != nil {
		s.AccountID = o.AccountID
	}
	if o.BilledToProvider != nil {
		s.BilledToProvider = o.BilledToProvider
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.EndDate != nil {
		s.EndDate = o.EndDate
	}
	if o.StartDate != nil {
		s.StartDate = o.StartDate
	}
	if o.UpdateTime != nil {
		s.UpdateTime = o.UpdateTime
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseInvoice) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseInvoice{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.ID != nil {
		o.ID = s.ID
	}
	if s.AccountID != nil {
		o.AccountID = s.AccountID
	}
	if s.BilledToProvider != nil {
		o.BilledToProvider = s.BilledToProvider
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.EndDate != nil {
		o.EndDate = s.EndDate
	}
	if s.StartDate != nil {
		o.StartDate = s.StartDate
	}
	if s.UpdateTime != nil {
		o.UpdateTime = s.UpdateTime
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseInvoice) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseInvoice) ToPlain() elemental.PlainIdentifiable {

	out := NewInvoice()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.AccountID != nil {
		out.AccountID = *o.AccountID
	}
	if o.BilledToProvider != nil {
		out.BilledToProvider = *o.BilledToProvider
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.EndDate != nil {
		out.EndDate = *o.EndDate
	}
	if o.StartDate != nil {
		out.StartDate = *o.StartDate
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}

	return out
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseInvoice) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseInvoice) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseInvoice) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseInvoice) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// DeepCopy returns a deep copy if the SparseInvoice.
func (o *SparseInvoice) DeepCopy() *SparseInvoice {

	if o == nil {
		return nil
	}

	out := &SparseInvoice{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseInvoice.
func (o *SparseInvoice) DeepCopyInto(out *SparseInvoice) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseInvoice: %s", err))
	}

	*out = *target.(*SparseInvoice)
}

type mongoAttributesInvoice struct {
	ID               string                       `bson:"id"`
	AccountID        string                       `bson:"accountid"`
	BilledToProvider InvoiceBilledToProviderValue `bson:"billedtoprovider"`
	CreateTime       time.Time                    `bson:"createtime"`
	EndDate          time.Time                    `bson:"enddate"`
	StartDate        time.Time                    `bson:"startdate"`
	UpdateTime       time.Time                    `bson:"updatetime"`
}
type mongoAttributesSparseInvoice struct {
	ID               *string                       `bson:"id,omitempty"`
	AccountID        *string                       `bson:"accountid,omitempty"`
	BilledToProvider *InvoiceBilledToProviderValue `bson:"billedtoprovider,omitempty"`
	CreateTime       *time.Time                    `bson:"createtime,omitempty"`
	EndDate          *time.Time                    `bson:"enddate,omitempty"`
	StartDate        *time.Time                    `bson:"startdate,omitempty"`
	UpdateTime       *time.Time                    `bson:"updatetime,omitempty"`
}
