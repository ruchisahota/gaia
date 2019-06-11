package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AccessReportActionValue represents the possible values for attribute "action".
type AccessReportActionValue string

const (
	// AccessReportActionAccept represents the value Accept.
	AccessReportActionAccept AccessReportActionValue = "Accept"

	// AccessReportActionReject represents the value Reject.
	AccessReportActionReject AccessReportActionValue = "Reject"
)

// AccessReportDestinationTypeValue represents the possible values for attribute "destinationType".
type AccessReportDestinationTypeValue string

const (
	// AccessReportDestinationTypeClaims represents the value Claims.
	AccessReportDestinationTypeClaims AccessReportDestinationTypeValue = "Claims"

	// AccessReportDestinationTypeExternalNetwork represents the value ExternalNetwork.
	AccessReportDestinationTypeExternalNetwork AccessReportDestinationTypeValue = "ExternalNetwork"

	// AccessReportDestinationTypeProcessingUnit represents the value ProcessingUnit.
	AccessReportDestinationTypeProcessingUnit AccessReportDestinationTypeValue = "ProcessingUnit"
)

// AccessReportSourceTypeValue represents the possible values for attribute "sourceType".
type AccessReportSourceTypeValue string

const (
	// AccessReportSourceTypeClaims represents the value Claims.
	AccessReportSourceTypeClaims AccessReportSourceTypeValue = "Claims"

	// AccessReportSourceTypeExternalNetwork represents the value ExternalNetwork.
	AccessReportSourceTypeExternalNetwork AccessReportSourceTypeValue = "ExternalNetwork"

	// AccessReportSourceTypeProcessingUnit represents the value ProcessingUnit.
	AccessReportSourceTypeProcessingUnit AccessReportSourceTypeValue = "ProcessingUnit"
)

// AccessReportIdentity represents the Identity of the object.
var AccessReportIdentity = elemental.Identity{
	Name:     "accessreport",
	Category: "accessreports",
	Package:  "zack",
	Private:  false,
}

// AccessReportsList represents a list of AccessReports
type AccessReportsList []*AccessReport

// Identity returns the identity of the objects in the list.
func (o AccessReportsList) Identity() elemental.Identity {

	return AccessReportIdentity
}

// Copy returns a pointer to a copy the AccessReportsList.
func (o AccessReportsList) Copy() elemental.Identifiables {

	copy := append(AccessReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AccessReportsList.
func (o AccessReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AccessReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AccessReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AccessReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AccessReportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the AccessReportsList converted to SparseAccessReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AccessReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAccessReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAccessReport)
	}

	return out
}

// Version returns the version of the content.
func (o AccessReportsList) Version() int {

	return 1
}

// AccessReport represents the model of a accessreport
type AccessReport struct {
	// Action applied to the access.
	Action AccessReportActionValue `json:"action" msgpack:"action" bson:"-" mapstructure:"action,omitempty"`

	// Hash of the claims used to communicate.
	ClaimHash string `json:"claimHash" msgpack:"claimHash" bson:"-" mapstructure:"claimHash,omitempty"`

	// content of the report.
	Content string `json:"content" msgpack:"content" bson:"-" mapstructure:"content,omitempty"`

	// ID of the destination.
	DestinationID string `json:"destinationID" msgpack:"destinationID" bson:"-" mapstructure:"destinationID,omitempty"`

	// Type of the destination.
	DestinationIP string `json:"destinationIP" msgpack:"destinationIP" bson:"-" mapstructure:"destinationIP,omitempty"`

	// Namespace of the receiver.
	DestinationNamespace string `json:"destinationNamespace" msgpack:"destinationNamespace" bson:"-" mapstructure:"destinationNamespace,omitempty"`

	// Port of the destination.
	DestinationPort int `json:"destinationPort" msgpack:"destinationPort" bson:"-" mapstructure:"destinationPort,omitempty"`

	// Type of the source.
	DestinationType AccessReportDestinationTypeValue `json:"destinationType" msgpack:"destinationType" bson:"-" mapstructure:"destinationType,omitempty"`

	// This field is only set if 'action' is set to 'Reject' and specifies the reason
	// for the rejection.
	DropReason string `json:"dropReason" msgpack:"dropReason" bson:"-" mapstructure:"dropReason,omitempty"`

	// This is here for backward compatibility.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// ID of the policy that accepted the access.
	PolicyID string `json:"policyID" msgpack:"policyID" bson:"-" mapstructure:"policyID,omitempty"`

	// Namespace of the policy that accepted the access.
	PolicyNamespace string `json:"policyNamespace" msgpack:"policyNamespace" bson:"-" mapstructure:"policyNamespace,omitempty"`

	// ID of the source.
	SourceID string `json:"sourceID" msgpack:"sourceID" bson:"-" mapstructure:"sourceID,omitempty"`

	// Type of the source.
	SourceIP string `json:"sourceIP" msgpack:"sourceIP" bson:"-" mapstructure:"sourceIP,omitempty"`

	// Namespace of the receiver.
	SourceNamespace string `json:"sourceNamespace" msgpack:"sourceNamespace" bson:"-" mapstructure:"sourceNamespace,omitempty"`

	// Type of the source.
	SourceType AccessReportSourceTypeValue `json:"sourceType" msgpack:"sourceType" bson:"-" mapstructure:"sourceType,omitempty"`

	// Date of the report.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	// Number of access in the report.
	Value int `json:"value" msgpack:"value" bson:"-" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAccessReport returns a new *AccessReport
func NewAccessReport() *AccessReport {

	return &AccessReport{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *AccessReport) Identity() elemental.Identity {

	return AccessReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AccessReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AccessReport) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *AccessReport) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *AccessReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *AccessReport) Doc() string {

	return `Post a new access report.`
}

func (o *AccessReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *AccessReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAccessReport{
			Action:               &o.Action,
			ClaimHash:            &o.ClaimHash,
			Content:              &o.Content,
			DestinationID:        &o.DestinationID,
			DestinationIP:        &o.DestinationIP,
			DestinationNamespace: &o.DestinationNamespace,
			DestinationPort:      &o.DestinationPort,
			DestinationType:      &o.DestinationType,
			DropReason:           &o.DropReason,
			Namespace:            &o.Namespace,
			PolicyID:             &o.PolicyID,
			PolicyNamespace:      &o.PolicyNamespace,
			SourceID:             &o.SourceID,
			SourceIP:             &o.SourceIP,
			SourceNamespace:      &o.SourceNamespace,
			SourceType:           &o.SourceType,
			Timestamp:            &o.Timestamp,
			Value:                &o.Value,
		}
	}

	sp := &SparseAccessReport{}
	for _, f := range fields {
		switch f {
		case "action":
			sp.Action = &(o.Action)
		case "claimHash":
			sp.ClaimHash = &(o.ClaimHash)
		case "content":
			sp.Content = &(o.Content)
		case "destinationID":
			sp.DestinationID = &(o.DestinationID)
		case "destinationIP":
			sp.DestinationIP = &(o.DestinationIP)
		case "destinationNamespace":
			sp.DestinationNamespace = &(o.DestinationNamespace)
		case "destinationPort":
			sp.DestinationPort = &(o.DestinationPort)
		case "destinationType":
			sp.DestinationType = &(o.DestinationType)
		case "dropReason":
			sp.DropReason = &(o.DropReason)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "policyID":
			sp.PolicyID = &(o.PolicyID)
		case "policyNamespace":
			sp.PolicyNamespace = &(o.PolicyNamespace)
		case "sourceID":
			sp.SourceID = &(o.SourceID)
		case "sourceIP":
			sp.SourceIP = &(o.SourceIP)
		case "sourceNamespace":
			sp.SourceNamespace = &(o.SourceNamespace)
		case "sourceType":
			sp.SourceType = &(o.SourceType)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		case "value":
			sp.Value = &(o.Value)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAccessReport to the object.
func (o *AccessReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAccessReport)
	if so.Action != nil {
		o.Action = *so.Action
	}
	if so.ClaimHash != nil {
		o.ClaimHash = *so.ClaimHash
	}
	if so.Content != nil {
		o.Content = *so.Content
	}
	if so.DestinationID != nil {
		o.DestinationID = *so.DestinationID
	}
	if so.DestinationIP != nil {
		o.DestinationIP = *so.DestinationIP
	}
	if so.DestinationNamespace != nil {
		o.DestinationNamespace = *so.DestinationNamespace
	}
	if so.DestinationPort != nil {
		o.DestinationPort = *so.DestinationPort
	}
	if so.DestinationType != nil {
		o.DestinationType = *so.DestinationType
	}
	if so.DropReason != nil {
		o.DropReason = *so.DropReason
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.PolicyID != nil {
		o.PolicyID = *so.PolicyID
	}
	if so.PolicyNamespace != nil {
		o.PolicyNamespace = *so.PolicyNamespace
	}
	if so.SourceID != nil {
		o.SourceID = *so.SourceID
	}
	if so.SourceIP != nil {
		o.SourceIP = *so.SourceIP
	}
	if so.SourceNamespace != nil {
		o.SourceNamespace = *so.SourceNamespace
	}
	if so.SourceType != nil {
		o.SourceType = *so.SourceType
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
	if so.Value != nil {
		o.Value = *so.Value
	}
}

// DeepCopy returns a deep copy if the AccessReport.
func (o *AccessReport) DeepCopy() *AccessReport {

	if o == nil {
		return nil
	}

	out := &AccessReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AccessReport.
func (o *AccessReport) DeepCopyInto(out *AccessReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AccessReport: %s", err))
	}

	*out = *target.(*AccessReport)
}

// Validate valides the current information stored into the structure.
func (o *AccessReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("action", string(o.Action)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("action", string(o.Action), []string{"Accept", "Reject"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("content", o.Content); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("destinationID", o.DestinationID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("destinationType", string(o.DestinationType)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("destinationType", string(o.DestinationType), []string{"ProcessingUnit", "ExternalNetwork", "Claims"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("namespace", o.Namespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("policyID", o.PolicyID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("policyNamespace", o.PolicyNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("sourceID", o.SourceID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("sourceType", string(o.SourceType)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("sourceType", string(o.SourceType), []string{"ProcessingUnit", "ExternalNetwork", "Claims"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("value", o.Value); err != nil {
		requiredErrors = requiredErrors.Append(err)
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
func (*AccessReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AccessReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AccessReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AccessReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AccessReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AccessReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "action":
		return o.Action
	case "claimHash":
		return o.ClaimHash
	case "content":
		return o.Content
	case "destinationID":
		return o.DestinationID
	case "destinationIP":
		return o.DestinationIP
	case "destinationNamespace":
		return o.DestinationNamespace
	case "destinationPort":
		return o.DestinationPort
	case "destinationType":
		return o.DestinationType
	case "dropReason":
		return o.DropReason
	case "namespace":
		return o.Namespace
	case "policyID":
		return o.PolicyID
	case "policyNamespace":
		return o.PolicyNamespace
	case "sourceID":
		return o.SourceID
	case "sourceIP":
		return o.SourceIP
	case "sourceNamespace":
		return o.SourceNamespace
	case "sourceType":
		return o.SourceType
	case "timestamp":
		return o.Timestamp
	case "value":
		return o.Value
	}

	return nil
}

// AccessReportAttributesMap represents the map of attribute for AccessReport.
var AccessReportAttributesMap = map[string]elemental.AttributeSpecification{
	"Action": elemental.AttributeSpecification{
		AllowedChoices: []string{"Accept", "Reject"},
		ConvertedName:  "Action",
		Description:    `Action applied to the access.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Type:           "enum",
	},
	"ClaimHash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ClaimHash",
		Description:    `Hash of the claims used to communicate.`,
		Exposed:        true,
		Name:           "claimHash",
		Type:           "string",
	},
	"Content": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Content",
		Description:    `content of the report.`,
		Exposed:        true,
		Name:           "content",
		Required:       true,
		Type:           "string",
	},
	"DestinationID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination.`,
		Exposed:        true,
		Name:           "destinationID",
		Required:       true,
		Type:           "string",
	},
	"DestinationIP": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationIP",
		Description:    `Type of the destination.`,
		Exposed:        true,
		Name:           "destinationIP",
		Type:           "string",
	},
	"DestinationNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationNamespace",
		Description:    `Namespace of the receiver.`,
		Exposed:        true,
		Name:           "destinationNamespace",
		Type:           "string",
	},
	"DestinationPort": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationPort",
		Description:    `Port of the destination.`,
		Exposed:        true,
		Name:           "destinationPort",
		Type:           "integer",
	},
	"DestinationType": elemental.AttributeSpecification{
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Claims"},
		ConvertedName:  "DestinationType",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "destinationType",
		Required:       true,
		Type:           "enum",
	},
	"DropReason": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DropReason",
		Description: `This field is only set if 'action' is set to 'Reject' and specifies the reason
for the rejection.`,
		Exposed: true,
		Name:    "dropReason",
		Type:    "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Deprecated:     true,
		Description:    `This is here for backward compatibility.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"PolicyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `ID of the policy that accepted the access.`,
		Exposed:        true,
		Name:           "policyID",
		Required:       true,
		Type:           "string",
	},
	"PolicyNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PolicyNamespace",
		Description:    `Namespace of the policy that accepted the access.`,
		Exposed:        true,
		Name:           "policyNamespace",
		Required:       true,
		Type:           "string",
	},
	"SourceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `ID of the source.`,
		Exposed:        true,
		Name:           "sourceID",
		Required:       true,
		Type:           "string",
	},
	"SourceIP": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceIP",
		Type:           "string",
	},
	"SourceNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceNamespace",
		Description:    `Namespace of the receiver.`,
		Exposed:        true,
		Name:           "sourceNamespace",
		Type:           "string",
	},
	"SourceType": elemental.AttributeSpecification{
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Claims"},
		ConvertedName:  "SourceType",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceType",
		Required:       true,
		Type:           "enum",
	},
	"Timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
	"Value": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Number of access in the report.`,
		Exposed:        true,
		Name:           "value",
		Required:       true,
		Type:           "integer",
	},
}

// AccessReportLowerCaseAttributesMap represents the map of attribute for AccessReport.
var AccessReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"action": elemental.AttributeSpecification{
		AllowedChoices: []string{"Accept", "Reject"},
		ConvertedName:  "Action",
		Description:    `Action applied to the access.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Type:           "enum",
	},
	"claimhash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ClaimHash",
		Description:    `Hash of the claims used to communicate.`,
		Exposed:        true,
		Name:           "claimHash",
		Type:           "string",
	},
	"content": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Content",
		Description:    `content of the report.`,
		Exposed:        true,
		Name:           "content",
		Required:       true,
		Type:           "string",
	},
	"destinationid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination.`,
		Exposed:        true,
		Name:           "destinationID",
		Required:       true,
		Type:           "string",
	},
	"destinationip": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationIP",
		Description:    `Type of the destination.`,
		Exposed:        true,
		Name:           "destinationIP",
		Type:           "string",
	},
	"destinationnamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationNamespace",
		Description:    `Namespace of the receiver.`,
		Exposed:        true,
		Name:           "destinationNamespace",
		Type:           "string",
	},
	"destinationport": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationPort",
		Description:    `Port of the destination.`,
		Exposed:        true,
		Name:           "destinationPort",
		Type:           "integer",
	},
	"destinationtype": elemental.AttributeSpecification{
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Claims"},
		ConvertedName:  "DestinationType",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "destinationType",
		Required:       true,
		Type:           "enum",
	},
	"dropreason": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DropReason",
		Description: `This field is only set if 'action' is set to 'Reject' and specifies the reason
for the rejection.`,
		Exposed: true,
		Name:    "dropReason",
		Type:    "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Deprecated:     true,
		Description:    `This is here for backward compatibility.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"policyid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `ID of the policy that accepted the access.`,
		Exposed:        true,
		Name:           "policyID",
		Required:       true,
		Type:           "string",
	},
	"policynamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PolicyNamespace",
		Description:    `Namespace of the policy that accepted the access.`,
		Exposed:        true,
		Name:           "policyNamespace",
		Required:       true,
		Type:           "string",
	},
	"sourceid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `ID of the source.`,
		Exposed:        true,
		Name:           "sourceID",
		Required:       true,
		Type:           "string",
	},
	"sourceip": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceIP",
		Type:           "string",
	},
	"sourcenamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceNamespace",
		Description:    `Namespace of the receiver.`,
		Exposed:        true,
		Name:           "sourceNamespace",
		Type:           "string",
	},
	"sourcetype": elemental.AttributeSpecification{
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Claims"},
		ConvertedName:  "SourceType",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceType",
		Required:       true,
		Type:           "enum",
	},
	"timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
	"value": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Number of access in the report.`,
		Exposed:        true,
		Name:           "value",
		Required:       true,
		Type:           "integer",
	},
}

// SparseAccessReportsList represents a list of SparseAccessReports
type SparseAccessReportsList []*SparseAccessReport

// Identity returns the identity of the objects in the list.
func (o SparseAccessReportsList) Identity() elemental.Identity {

	return AccessReportIdentity
}

// Copy returns a pointer to a copy the SparseAccessReportsList.
func (o SparseAccessReportsList) Copy() elemental.Identifiables {

	copy := append(SparseAccessReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAccessReportsList.
func (o SparseAccessReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAccessReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAccessReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAccessReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAccessReportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseAccessReportsList converted to AccessReportsList.
func (o SparseAccessReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAccessReportsList) Version() int {

	return 1
}

// SparseAccessReport represents the sparse version of a accessreport.
type SparseAccessReport struct {
	// Action applied to the access.
	Action *AccessReportActionValue `json:"action,omitempty" msgpack:"action,omitempty" bson:"-" mapstructure:"action,omitempty"`

	// Hash of the claims used to communicate.
	ClaimHash *string `json:"claimHash,omitempty" msgpack:"claimHash,omitempty" bson:"-" mapstructure:"claimHash,omitempty"`

	// content of the report.
	Content *string `json:"content,omitempty" msgpack:"content,omitempty" bson:"-" mapstructure:"content,omitempty"`

	// ID of the destination.
	DestinationID *string `json:"destinationID,omitempty" msgpack:"destinationID,omitempty" bson:"-" mapstructure:"destinationID,omitempty"`

	// Type of the destination.
	DestinationIP *string `json:"destinationIP,omitempty" msgpack:"destinationIP,omitempty" bson:"-" mapstructure:"destinationIP,omitempty"`

	// Namespace of the receiver.
	DestinationNamespace *string `json:"destinationNamespace,omitempty" msgpack:"destinationNamespace,omitempty" bson:"-" mapstructure:"destinationNamespace,omitempty"`

	// Port of the destination.
	DestinationPort *int `json:"destinationPort,omitempty" msgpack:"destinationPort,omitempty" bson:"-" mapstructure:"destinationPort,omitempty"`

	// Type of the source.
	DestinationType *AccessReportDestinationTypeValue `json:"destinationType,omitempty" msgpack:"destinationType,omitempty" bson:"-" mapstructure:"destinationType,omitempty"`

	// This field is only set if 'action' is set to 'Reject' and specifies the reason
	// for the rejection.
	DropReason *string `json:"dropReason,omitempty" msgpack:"dropReason,omitempty" bson:"-" mapstructure:"dropReason,omitempty"`

	// This is here for backward compatibility.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	// ID of the policy that accepted the access.
	PolicyID *string `json:"policyID,omitempty" msgpack:"policyID,omitempty" bson:"-" mapstructure:"policyID,omitempty"`

	// Namespace of the policy that accepted the access.
	PolicyNamespace *string `json:"policyNamespace,omitempty" msgpack:"policyNamespace,omitempty" bson:"-" mapstructure:"policyNamespace,omitempty"`

	// ID of the source.
	SourceID *string `json:"sourceID,omitempty" msgpack:"sourceID,omitempty" bson:"-" mapstructure:"sourceID,omitempty"`

	// Type of the source.
	SourceIP *string `json:"sourceIP,omitempty" msgpack:"sourceIP,omitempty" bson:"-" mapstructure:"sourceIP,omitempty"`

	// Namespace of the receiver.
	SourceNamespace *string `json:"sourceNamespace,omitempty" msgpack:"sourceNamespace,omitempty" bson:"-" mapstructure:"sourceNamespace,omitempty"`

	// Type of the source.
	SourceType *AccessReportSourceTypeValue `json:"sourceType,omitempty" msgpack:"sourceType,omitempty" bson:"-" mapstructure:"sourceType,omitempty"`

	// Date of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"-" mapstructure:"timestamp,omitempty"`

	// Number of access in the report.
	Value *int `json:"value,omitempty" msgpack:"value,omitempty" bson:"-" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAccessReport returns a new  SparseAccessReport.
func NewSparseAccessReport() *SparseAccessReport {
	return &SparseAccessReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAccessReport) Identity() elemental.Identity {

	return AccessReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAccessReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAccessReport) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseAccessReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAccessReport) ToPlain() elemental.PlainIdentifiable {

	out := NewAccessReport()
	if o.Action != nil {
		out.Action = *o.Action
	}
	if o.ClaimHash != nil {
		out.ClaimHash = *o.ClaimHash
	}
	if o.Content != nil {
		out.Content = *o.Content
	}
	if o.DestinationID != nil {
		out.DestinationID = *o.DestinationID
	}
	if o.DestinationIP != nil {
		out.DestinationIP = *o.DestinationIP
	}
	if o.DestinationNamespace != nil {
		out.DestinationNamespace = *o.DestinationNamespace
	}
	if o.DestinationPort != nil {
		out.DestinationPort = *o.DestinationPort
	}
	if o.DestinationType != nil {
		out.DestinationType = *o.DestinationType
	}
	if o.DropReason != nil {
		out.DropReason = *o.DropReason
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.PolicyID != nil {
		out.PolicyID = *o.PolicyID
	}
	if o.PolicyNamespace != nil {
		out.PolicyNamespace = *o.PolicyNamespace
	}
	if o.SourceID != nil {
		out.SourceID = *o.SourceID
	}
	if o.SourceIP != nil {
		out.SourceIP = *o.SourceIP
	}
	if o.SourceNamespace != nil {
		out.SourceNamespace = *o.SourceNamespace
	}
	if o.SourceType != nil {
		out.SourceType = *o.SourceType
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}
	if o.Value != nil {
		out.Value = *o.Value
	}

	return out
}

// DeepCopy returns a deep copy if the SparseAccessReport.
func (o *SparseAccessReport) DeepCopy() *SparseAccessReport {

	if o == nil {
		return nil
	}

	out := &SparseAccessReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAccessReport.
func (o *SparseAccessReport) DeepCopyInto(out *SparseAccessReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAccessReport: %s", err))
	}

	*out = *target.(*SparseAccessReport)
}
