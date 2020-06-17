package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// FlowReportActionValue represents the possible values for attribute "action".
type FlowReportActionValue string

const (
	// FlowReportActionAccept represents the value Accept.
	FlowReportActionAccept FlowReportActionValue = "Accept"

	// FlowReportActionReject represents the value Reject.
	FlowReportActionReject FlowReportActionValue = "Reject"
)

// FlowReportDestinationTypeValue represents the possible values for attribute "destinationType".
type FlowReportDestinationTypeValue string

const (
	// FlowReportDestinationTypeClaims represents the value Claims.
	FlowReportDestinationTypeClaims FlowReportDestinationTypeValue = "Claims"

	// FlowReportDestinationTypeExternalNetwork represents the value ExternalNetwork.
	FlowReportDestinationTypeExternalNetwork FlowReportDestinationTypeValue = "ExternalNetwork"

	// FlowReportDestinationTypeProcessingUnit represents the value ProcessingUnit.
	FlowReportDestinationTypeProcessingUnit FlowReportDestinationTypeValue = "ProcessingUnit"
)

// FlowReportObservedActionValue represents the possible values for attribute "observedAction".
type FlowReportObservedActionValue string

const (
	// FlowReportObservedActionAccept represents the value Accept.
	FlowReportObservedActionAccept FlowReportObservedActionValue = "Accept"

	// FlowReportObservedActionNotApplicable represents the value NotApplicable.
	FlowReportObservedActionNotApplicable FlowReportObservedActionValue = "NotApplicable"

	// FlowReportObservedActionReject represents the value Reject.
	FlowReportObservedActionReject FlowReportObservedActionValue = "Reject"
)

// FlowReportServiceTypeValue represents the possible values for attribute "serviceType".
type FlowReportServiceTypeValue string

const (
	// FlowReportServiceTypeHTTP represents the value HTTP.
	FlowReportServiceTypeHTTP FlowReportServiceTypeValue = "HTTP"

	// FlowReportServiceTypeL3 represents the value L3.
	FlowReportServiceTypeL3 FlowReportServiceTypeValue = "L3"

	// FlowReportServiceTypeNotApplicable represents the value NotApplicable.
	FlowReportServiceTypeNotApplicable FlowReportServiceTypeValue = "NotApplicable"

	// FlowReportServiceTypeTCP represents the value TCP.
	FlowReportServiceTypeTCP FlowReportServiceTypeValue = "TCP"
)

// FlowReportSourceTypeValue represents the possible values for attribute "sourceType".
type FlowReportSourceTypeValue string

const (
	// FlowReportSourceTypeClaims represents the value Claims.
	FlowReportSourceTypeClaims FlowReportSourceTypeValue = "Claims"

	// FlowReportSourceTypeExternalNetwork represents the value ExternalNetwork.
	FlowReportSourceTypeExternalNetwork FlowReportSourceTypeValue = "ExternalNetwork"

	// FlowReportSourceTypeProcessingUnit represents the value ProcessingUnit.
	FlowReportSourceTypeProcessingUnit FlowReportSourceTypeValue = "ProcessingUnit"
)

// FlowReportIdentity represents the Identity of the object.
var FlowReportIdentity = elemental.Identity{
	Name:     "flowreport",
	Category: "flowreports",
	Package:  "zack",
	Private:  false,
}

// FlowReportsList represents a list of FlowReports
type FlowReportsList []*FlowReport

// Identity returns the identity of the objects in the list.
func (o FlowReportsList) Identity() elemental.Identity {

	return FlowReportIdentity
}

// Copy returns a pointer to a copy the FlowReportsList.
func (o FlowReportsList) Copy() elemental.Identifiables {

	copy := append(FlowReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the FlowReportsList.
func (o FlowReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(FlowReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*FlowReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o FlowReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o FlowReportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the FlowReportsList converted to SparseFlowReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o FlowReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseFlowReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseFlowReport)
	}

	return out
}

// Version returns the version of the content.
func (o FlowReportsList) Version() int {

	return 1
}

// FlowReport represents the model of a flowreport
type FlowReport struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Action applied to the flow.
	Action FlowReportActionValue `json:"action" msgpack:"action" bson:"action" mapstructure:"action,omitempty"`

	// Identifier of the destination controller.
	DestinationController string `json:"destinationController" msgpack:"destinationController" bson:"-" mapstructure:"destinationController,omitempty"`

	// ID of the destination.
	DestinationID string `json:"destinationID" msgpack:"destinationID" bson:"destinationid" mapstructure:"destinationID,omitempty"`

	// Destination IP address.
	DestinationIP string `json:"destinationIP" msgpack:"destinationIP" bson:"destinationip" mapstructure:"destinationIP,omitempty"`

	// Namespace of the destination. This is deprecated. Use `remoteNamespace`. This
	// property does nothing.
	DestinationNamespace string `json:"destinationNamespace,omitempty" msgpack:"destinationNamespace,omitempty" bson:"destinationnamespace" mapstructure:"destinationNamespace,omitempty"`

	// Identifier of the destination platform.
	DestinationPlatform string `json:"destinationPlatform" msgpack:"destinationPlatform" bson:"destinationplatform" mapstructure:"destinationPlatform,omitempty"`

	// Port of the destination.
	DestinationPort int `json:"destinationPort" msgpack:"destinationPort" bson:"destinationport" mapstructure:"destinationPort,omitempty"`

	// Destination type.
	DestinationType FlowReportDestinationTypeValue `json:"destinationType" msgpack:"destinationType" bson:"destinationtype" mapstructure:"destinationType,omitempty"`

	// This field is only set if `action` is set to `Reject`. It specifies the reason
	// for the rejection.
	DropReason string `json:"dropReason" msgpack:"dropReason" bson:"dropreason" mapstructure:"dropReason,omitempty"`

	// If `true`, the flow was encrypted.
	Encrypted bool `json:"encrypted" msgpack:"encrypted" bson:"encrypted" mapstructure:"encrypted,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// This is here for backward compatibility.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// If `true`, design mode is on.
	Observed bool `json:"observed" msgpack:"observed" bson:"observed" mapstructure:"observed,omitempty"`

	// Action observed on the flow.
	ObservedAction FlowReportObservedActionValue `json:"observedAction" msgpack:"observedAction" bson:"observedaction" mapstructure:"observedAction,omitempty"`

	// Specifies the reason for a rejection. Only set if `observedAction` is set
	// to `Reject`.
	ObservedDropReason string `json:"observedDropReason" msgpack:"observedDropReason" bson:"observeddropreason" mapstructure:"observedDropReason,omitempty"`

	// Value of the encryption of the network policy that observed the flow.
	ObservedEncrypted bool `json:"observedEncrypted" msgpack:"observedEncrypted" bson:"observedencrypted" mapstructure:"observedEncrypted,omitempty"`

	// ID of the network policy that observed the flow.
	ObservedPolicyID string `json:"observedPolicyID" msgpack:"observedPolicyID" bson:"observedpolicyid" mapstructure:"observedPolicyID,omitempty"`

	// Namespace of the network policy that observed the flow.
	ObservedPolicyNamespace string `json:"observedPolicyNamespace" msgpack:"observedPolicyNamespace" bson:"observedpolicynamespace" mapstructure:"observedPolicyNamespace,omitempty"`

	// ID of the network policy that accepted the flow.
	PolicyID string `json:"policyID" msgpack:"policyID" bson:"policyid" mapstructure:"policyID,omitempty"`

	// Namespace of the network policy that accepted the flow.
	PolicyNamespace string `json:"policyNamespace" msgpack:"policyNamespace" bson:"policynamespace" mapstructure:"policyNamespace,omitempty"`

	// Protocol number.
	Protocol int `json:"protocol" msgpack:"protocol" bson:"protocol" mapstructure:"protocol,omitempty"`

	// Namespace of the object at the other end of the flow.
	RemoteNamespace string `json:"remoteNamespace,omitempty" msgpack:"remoteNamespace,omitempty" bson:"remotenamespace" mapstructure:"remoteNamespace,omitempty"`

	// Hash of the claims used to communicate.
	ServiceClaimHash string `json:"serviceClaimHash" msgpack:"serviceClaimHash" bson:"serviceclaimhash" mapstructure:"serviceClaimHash,omitempty"`

	// ID of the service.
	ServiceID string `json:"serviceID" msgpack:"serviceID" bson:"serviceid" mapstructure:"serviceID,omitempty"`

	// Namespace of Service accessed.
	ServiceNamespace string `json:"serviceNamespace" msgpack:"serviceNamespace" bson:"servicenamespace" mapstructure:"serviceNamespace,omitempty"`

	// ID of the service.
	ServiceType FlowReportServiceTypeValue `json:"serviceType" msgpack:"serviceType" bson:"servicetype" mapstructure:"serviceType,omitempty"`

	// Service URL accessed.
	ServiceURL string `json:"serviceURL" msgpack:"serviceURL" bson:"serviceurl" mapstructure:"serviceURL,omitempty"`

	// Identifier of the source controller.
	SourceController string `json:"sourceController" msgpack:"sourceController" bson:"-" mapstructure:"sourceController,omitempty"`

	// ID of the source.
	SourceID string `json:"sourceID" msgpack:"sourceID" bson:"sourceid" mapstructure:"sourceID,omitempty"`

	// Type of the source.
	SourceIP string `json:"sourceIP" msgpack:"sourceIP" bson:"sourceip" mapstructure:"sourceIP,omitempty"`

	// Namespace of the source. This is deprecated. Use `remoteNamespace`. This
	// property does nothing.
	SourceNamespace string `json:"sourceNamespace,omitempty" msgpack:"sourceNamespace,omitempty" bson:"sourcenamespace" mapstructure:"sourceNamespace,omitempty"`

	// Identifier of the source platform.
	SourcePlatform string `json:"sourcePlatform" msgpack:"sourcePlatform" bson:"sourceplatform" mapstructure:"sourcePlatform,omitempty"`

	// Type of the source.
	SourceType FlowReportSourceTypeValue `json:"sourceType" msgpack:"sourceType" bson:"sourcetype" mapstructure:"sourceType,omitempty"`

	// Time and date of the log.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"timestamp" mapstructure:"timestamp,omitempty"`

	// Number of flows in the log.
	Value int `json:"value" msgpack:"value" bson:"value" mapstructure:"value,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewFlowReport returns a new *FlowReport
func NewFlowReport() *FlowReport {

	return &FlowReport{
		ModelVersion:   1,
		ServiceType:    FlowReportServiceTypeNotApplicable,
		ObservedAction: FlowReportObservedActionNotApplicable,
		MigrationsLog:  map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *FlowReport) Identity() elemental.Identity {

	return FlowReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FlowReport) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FlowReport) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FlowReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesFlowReport{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Action = o.Action
	s.DestinationID = o.DestinationID
	s.DestinationIP = o.DestinationIP
	s.DestinationNamespace = o.DestinationNamespace
	s.DestinationPlatform = o.DestinationPlatform
	s.DestinationPort = o.DestinationPort
	s.DestinationType = o.DestinationType
	s.DropReason = o.DropReason
	s.Encrypted = o.Encrypted
	s.MigrationsLog = o.MigrationsLog
	s.Namespace = o.Namespace
	s.Observed = o.Observed
	s.ObservedAction = o.ObservedAction
	s.ObservedDropReason = o.ObservedDropReason
	s.ObservedEncrypted = o.ObservedEncrypted
	s.ObservedPolicyID = o.ObservedPolicyID
	s.ObservedPolicyNamespace = o.ObservedPolicyNamespace
	s.PolicyID = o.PolicyID
	s.PolicyNamespace = o.PolicyNamespace
	s.Protocol = o.Protocol
	s.RemoteNamespace = o.RemoteNamespace
	s.ServiceClaimHash = o.ServiceClaimHash
	s.ServiceID = o.ServiceID
	s.ServiceNamespace = o.ServiceNamespace
	s.ServiceType = o.ServiceType
	s.ServiceURL = o.ServiceURL
	s.SourceID = o.SourceID
	s.SourceIP = o.SourceIP
	s.SourceNamespace = o.SourceNamespace
	s.SourcePlatform = o.SourcePlatform
	s.SourceType = o.SourceType
	s.Timestamp = o.Timestamp
	s.Value = o.Value
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *FlowReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesFlowReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Action = s.Action
	o.DestinationID = s.DestinationID
	o.DestinationIP = s.DestinationIP
	o.DestinationNamespace = s.DestinationNamespace
	o.DestinationPlatform = s.DestinationPlatform
	o.DestinationPort = s.DestinationPort
	o.DestinationType = s.DestinationType
	o.DropReason = s.DropReason
	o.Encrypted = s.Encrypted
	o.MigrationsLog = s.MigrationsLog
	o.Namespace = s.Namespace
	o.Observed = s.Observed
	o.ObservedAction = s.ObservedAction
	o.ObservedDropReason = s.ObservedDropReason
	o.ObservedEncrypted = s.ObservedEncrypted
	o.ObservedPolicyID = s.ObservedPolicyID
	o.ObservedPolicyNamespace = s.ObservedPolicyNamespace
	o.PolicyID = s.PolicyID
	o.PolicyNamespace = s.PolicyNamespace
	o.Protocol = s.Protocol
	o.RemoteNamespace = s.RemoteNamespace
	o.ServiceClaimHash = s.ServiceClaimHash
	o.ServiceID = s.ServiceID
	o.ServiceNamespace = s.ServiceNamespace
	o.ServiceType = s.ServiceType
	o.ServiceURL = s.ServiceURL
	o.SourceID = s.SourceID
	o.SourceIP = s.SourceIP
	o.SourceNamespace = s.SourceNamespace
	o.SourcePlatform = s.SourcePlatform
	o.SourceType = s.SourceType
	o.Timestamp = s.Timestamp
	o.Value = s.Value
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *FlowReport) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *FlowReport) BleveType() string {

	return "flowreport"
}

// DefaultOrder returns the list of default ordering fields.
func (o *FlowReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *FlowReport) Doc() string {

	return `Post a new flow log.`
}

func (o *FlowReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *FlowReport) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *FlowReport) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetZHash returns the ZHash of the receiver.
func (o *FlowReport) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *FlowReport) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *FlowReport) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *FlowReport) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *FlowReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseFlowReport{
			ID:                      &o.ID,
			Action:                  &o.Action,
			DestinationController:   &o.DestinationController,
			DestinationID:           &o.DestinationID,
			DestinationIP:           &o.DestinationIP,
			DestinationNamespace:    &o.DestinationNamespace,
			DestinationPlatform:     &o.DestinationPlatform,
			DestinationPort:         &o.DestinationPort,
			DestinationType:         &o.DestinationType,
			DropReason:              &o.DropReason,
			Encrypted:               &o.Encrypted,
			MigrationsLog:           &o.MigrationsLog,
			Namespace:               &o.Namespace,
			Observed:                &o.Observed,
			ObservedAction:          &o.ObservedAction,
			ObservedDropReason:      &o.ObservedDropReason,
			ObservedEncrypted:       &o.ObservedEncrypted,
			ObservedPolicyID:        &o.ObservedPolicyID,
			ObservedPolicyNamespace: &o.ObservedPolicyNamespace,
			PolicyID:                &o.PolicyID,
			PolicyNamespace:         &o.PolicyNamespace,
			Protocol:                &o.Protocol,
			RemoteNamespace:         &o.RemoteNamespace,
			ServiceClaimHash:        &o.ServiceClaimHash,
			ServiceID:               &o.ServiceID,
			ServiceNamespace:        &o.ServiceNamespace,
			ServiceType:             &o.ServiceType,
			ServiceURL:              &o.ServiceURL,
			SourceController:        &o.SourceController,
			SourceID:                &o.SourceID,
			SourceIP:                &o.SourceIP,
			SourceNamespace:         &o.SourceNamespace,
			SourcePlatform:          &o.SourcePlatform,
			SourceType:              &o.SourceType,
			Timestamp:               &o.Timestamp,
			Value:                   &o.Value,
			ZHash:                   &o.ZHash,
			Zone:                    &o.Zone,
		}
	}

	sp := &SparseFlowReport{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "action":
			sp.Action = &(o.Action)
		case "destinationController":
			sp.DestinationController = &(o.DestinationController)
		case "destinationID":
			sp.DestinationID = &(o.DestinationID)
		case "destinationIP":
			sp.DestinationIP = &(o.DestinationIP)
		case "destinationNamespace":
			sp.DestinationNamespace = &(o.DestinationNamespace)
		case "destinationPlatform":
			sp.DestinationPlatform = &(o.DestinationPlatform)
		case "destinationPort":
			sp.DestinationPort = &(o.DestinationPort)
		case "destinationType":
			sp.DestinationType = &(o.DestinationType)
		case "dropReason":
			sp.DropReason = &(o.DropReason)
		case "encrypted":
			sp.Encrypted = &(o.Encrypted)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "observed":
			sp.Observed = &(o.Observed)
		case "observedAction":
			sp.ObservedAction = &(o.ObservedAction)
		case "observedDropReason":
			sp.ObservedDropReason = &(o.ObservedDropReason)
		case "observedEncrypted":
			sp.ObservedEncrypted = &(o.ObservedEncrypted)
		case "observedPolicyID":
			sp.ObservedPolicyID = &(o.ObservedPolicyID)
		case "observedPolicyNamespace":
			sp.ObservedPolicyNamespace = &(o.ObservedPolicyNamespace)
		case "policyID":
			sp.PolicyID = &(o.PolicyID)
		case "policyNamespace":
			sp.PolicyNamespace = &(o.PolicyNamespace)
		case "protocol":
			sp.Protocol = &(o.Protocol)
		case "remoteNamespace":
			sp.RemoteNamespace = &(o.RemoteNamespace)
		case "serviceClaimHash":
			sp.ServiceClaimHash = &(o.ServiceClaimHash)
		case "serviceID":
			sp.ServiceID = &(o.ServiceID)
		case "serviceNamespace":
			sp.ServiceNamespace = &(o.ServiceNamespace)
		case "serviceType":
			sp.ServiceType = &(o.ServiceType)
		case "serviceURL":
			sp.ServiceURL = &(o.ServiceURL)
		case "sourceController":
			sp.SourceController = &(o.SourceController)
		case "sourceID":
			sp.SourceID = &(o.SourceID)
		case "sourceIP":
			sp.SourceIP = &(o.SourceIP)
		case "sourceNamespace":
			sp.SourceNamespace = &(o.SourceNamespace)
		case "sourcePlatform":
			sp.SourcePlatform = &(o.SourcePlatform)
		case "sourceType":
			sp.SourceType = &(o.SourceType)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		case "value":
			sp.Value = &(o.Value)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseFlowReport to the object.
func (o *FlowReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseFlowReport)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Action != nil {
		o.Action = *so.Action
	}
	if so.DestinationController != nil {
		o.DestinationController = *so.DestinationController
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
	if so.DestinationPlatform != nil {
		o.DestinationPlatform = *so.DestinationPlatform
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
	if so.Encrypted != nil {
		o.Encrypted = *so.Encrypted
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Observed != nil {
		o.Observed = *so.Observed
	}
	if so.ObservedAction != nil {
		o.ObservedAction = *so.ObservedAction
	}
	if so.ObservedDropReason != nil {
		o.ObservedDropReason = *so.ObservedDropReason
	}
	if so.ObservedEncrypted != nil {
		o.ObservedEncrypted = *so.ObservedEncrypted
	}
	if so.ObservedPolicyID != nil {
		o.ObservedPolicyID = *so.ObservedPolicyID
	}
	if so.ObservedPolicyNamespace != nil {
		o.ObservedPolicyNamespace = *so.ObservedPolicyNamespace
	}
	if so.PolicyID != nil {
		o.PolicyID = *so.PolicyID
	}
	if so.PolicyNamespace != nil {
		o.PolicyNamespace = *so.PolicyNamespace
	}
	if so.Protocol != nil {
		o.Protocol = *so.Protocol
	}
	if so.RemoteNamespace != nil {
		o.RemoteNamespace = *so.RemoteNamespace
	}
	if so.ServiceClaimHash != nil {
		o.ServiceClaimHash = *so.ServiceClaimHash
	}
	if so.ServiceID != nil {
		o.ServiceID = *so.ServiceID
	}
	if so.ServiceNamespace != nil {
		o.ServiceNamespace = *so.ServiceNamespace
	}
	if so.ServiceType != nil {
		o.ServiceType = *so.ServiceType
	}
	if so.ServiceURL != nil {
		o.ServiceURL = *so.ServiceURL
	}
	if so.SourceController != nil {
		o.SourceController = *so.SourceController
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
	if so.SourcePlatform != nil {
		o.SourcePlatform = *so.SourcePlatform
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
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the FlowReport.
func (o *FlowReport) DeepCopy() *FlowReport {

	if o == nil {
		return nil
	}

	out := &FlowReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *FlowReport.
func (o *FlowReport) DeepCopyInto(out *FlowReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy FlowReport: %s", err))
	}

	*out = *target.(*FlowReport)
}

// Validate valides the current information stored into the structure.
func (o *FlowReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("action", string(o.Action)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("action", string(o.Action), []string{"Accept", "Reject"}, false); err != nil {
		errors = errors.Append(err)
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

	if err := elemental.ValidateStringInList("observedAction", string(o.ObservedAction), []string{"Accept", "Reject", "NotApplicable"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("policyID", o.PolicyID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("policyNamespace", o.PolicyNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("protocol", o.Protocol); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("serviceType", string(o.ServiceType), []string{"L3", "HTTP", "TCP", "NotApplicable"}, false); err != nil {
		errors = errors.Append(err)
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
func (*FlowReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FlowReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FlowReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*FlowReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FlowReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *FlowReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "action":
		return o.Action
	case "destinationController":
		return o.DestinationController
	case "destinationID":
		return o.DestinationID
	case "destinationIP":
		return o.DestinationIP
	case "destinationNamespace":
		return o.DestinationNamespace
	case "destinationPlatform":
		return o.DestinationPlatform
	case "destinationPort":
		return o.DestinationPort
	case "destinationType":
		return o.DestinationType
	case "dropReason":
		return o.DropReason
	case "encrypted":
		return o.Encrypted
	case "migrationsLog":
		return o.MigrationsLog
	case "namespace":
		return o.Namespace
	case "observed":
		return o.Observed
	case "observedAction":
		return o.ObservedAction
	case "observedDropReason":
		return o.ObservedDropReason
	case "observedEncrypted":
		return o.ObservedEncrypted
	case "observedPolicyID":
		return o.ObservedPolicyID
	case "observedPolicyNamespace":
		return o.ObservedPolicyNamespace
	case "policyID":
		return o.PolicyID
	case "policyNamespace":
		return o.PolicyNamespace
	case "protocol":
		return o.Protocol
	case "remoteNamespace":
		return o.RemoteNamespace
	case "serviceClaimHash":
		return o.ServiceClaimHash
	case "serviceID":
		return o.ServiceID
	case "serviceNamespace":
		return o.ServiceNamespace
	case "serviceType":
		return o.ServiceType
	case "serviceURL":
		return o.ServiceURL
	case "sourceController":
		return o.SourceController
	case "sourceID":
		return o.SourceID
	case "sourceIP":
		return o.SourceIP
	case "sourceNamespace":
		return o.SourceNamespace
	case "sourcePlatform":
		return o.SourcePlatform
	case "sourceType":
		return o.SourceType
	case "timestamp":
		return o.Timestamp
	case "value":
		return o.Value
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// FlowReportAttributesMap represents the map of attribute for FlowReport.
var FlowReportAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Action": {
		AllowedChoices: []string{"Accept", "Reject"},
		ConvertedName:  "Action",
		Description:    `Action applied to the flow.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"DestinationController": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationController",
		Description:    `Identifier of the destination controller.`,
		Exposed:        true,
		Name:           "destinationController",
		Type:           "string",
	},
	"DestinationID": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination.`,
		Exposed:        true,
		Name:           "destinationID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"DestinationIP": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationIP",
		Description:    `Destination IP address.`,
		Exposed:        true,
		Name:           "destinationIP",
		Stored:         true,
		Type:           "string",
	},
	"DestinationNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationNamespace",
		Deprecated:     true,
		Description: `Namespace of the destination. This is deprecated. Use ` + "`" + `remoteNamespace` + "`" + `. This
property does nothing.`,
		Exposed: true,
		Name:    "destinationNamespace",
		Stored:  true,
		Type:    "string",
	},
	"DestinationPlatform": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationPlatform",
		Description:    `Identifier of the destination platform.`,
		Exposed:        true,
		Name:           "destinationPlatform",
		Stored:         true,
		Type:           "string",
	},
	"DestinationPort": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationPort",
		Description:    `Port of the destination.`,
		Exposed:        true,
		Name:           "destinationPort",
		Stored:         true,
		Type:           "integer",
	},
	"DestinationType": {
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Claims"},
		ConvertedName:  "DestinationType",
		Description:    `Destination type.`,
		Exposed:        true,
		Name:           "destinationType",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"DropReason": {
		AllowedChoices: []string{},
		ConvertedName:  "DropReason",
		Description: `This field is only set if ` + "`" + `action` + "`" + ` is set to ` + "`" + `Reject` + "`" + `. It specifies the reason
for the rejection.`,
		Exposed: true,
		Name:    "dropReason",
		Stored:  true,
		Type:    "string",
	},
	"Encrypted": {
		AllowedChoices: []string{},
		ConvertedName:  "Encrypted",
		Description:    `If ` + "`" + `true` + "`" + `, the flow was encrypted.`,
		Exposed:        true,
		Name:           "encrypted",
		Stored:         true,
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
	"Namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Deprecated:     true,
		Description:    `This is here for backward compatibility.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Observed": {
		AllowedChoices: []string{},
		ConvertedName:  "Observed",
		Description:    `If ` + "`" + `true` + "`" + `, design mode is on.`,
		Exposed:        true,
		Name:           "observed",
		Stored:         true,
		Type:           "boolean",
	},
	"ObservedAction": {
		AllowedChoices: []string{"Accept", "Reject", "NotApplicable"},
		ConvertedName:  "ObservedAction",
		DefaultValue:   FlowReportObservedActionNotApplicable,
		Description:    `Action observed on the flow.`,
		Exposed:        true,
		Name:           "observedAction",
		Stored:         true,
		Type:           "enum",
	},
	"ObservedDropReason": {
		AllowedChoices: []string{},
		ConvertedName:  "ObservedDropReason",
		Description: `Specifies the reason for a rejection. Only set if ` + "`" + `observedAction` + "`" + ` is set
to ` + "`" + `Reject` + "`" + `.`,
		Exposed: true,
		Name:    "observedDropReason",
		Stored:  true,
		Type:    "string",
	},
	"ObservedEncrypted": {
		AllowedChoices: []string{},
		ConvertedName:  "ObservedEncrypted",
		Description:    `Value of the encryption of the network policy that observed the flow.`,
		Exposed:        true,
		Name:           "observedEncrypted",
		Stored:         true,
		Type:           "boolean",
	},
	"ObservedPolicyID": {
		AllowedChoices: []string{},
		ConvertedName:  "ObservedPolicyID",
		Description:    `ID of the network policy that observed the flow.`,
		Exposed:        true,
		Name:           "observedPolicyID",
		Stored:         true,
		Type:           "string",
	},
	"ObservedPolicyNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ObservedPolicyNamespace",
		Description:    `Namespace of the network policy that observed the flow.`,
		Exposed:        true,
		Name:           "observedPolicyNamespace",
		Stored:         true,
		Type:           "string",
	},
	"PolicyID": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `ID of the network policy that accepted the flow.`,
		Exposed:        true,
		Name:           "policyID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"PolicyNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyNamespace",
		Description:    `Namespace of the network policy that accepted the flow.`,
		Exposed:        true,
		Name:           "policyNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Protocol": {
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol number.`,
		Exposed:        true,
		Name:           "protocol",
		Required:       true,
		Stored:         true,
		Type:           "integer",
	},
	"RemoteNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "RemoteNamespace",
		Description:    `Namespace of the object at the other end of the flow.`,
		Exposed:        true,
		Name:           "remoteNamespace",
		Stored:         true,
		Type:           "string",
	},
	"ServiceClaimHash": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceClaimHash",
		Description:    `Hash of the claims used to communicate.`,
		Exposed:        true,
		Name:           "serviceClaimHash",
		Stored:         true,
		Type:           "string",
	},
	"ServiceID": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceID",
		Description:    `ID of the service.`,
		Exposed:        true,
		Name:           "serviceID",
		Stored:         true,
		Type:           "string",
	},
	"ServiceNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceNamespace",
		Description:    `Namespace of Service accessed.`,
		Exposed:        true,
		Name:           "serviceNamespace",
		Stored:         true,
		Type:           "string",
	},
	"ServiceType": {
		AllowedChoices: []string{"L3", "HTTP", "TCP", "NotApplicable"},
		ConvertedName:  "ServiceType",
		DefaultValue:   FlowReportServiceTypeNotApplicable,
		Description:    `ID of the service.`,
		Exposed:        true,
		Name:           "serviceType",
		Stored:         true,
		Type:           "enum",
	},
	"ServiceURL": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceURL",
		Description:    `Service URL accessed.`,
		Exposed:        true,
		Name:           "serviceURL",
		Stored:         true,
		Type:           "string",
	},
	"SourceController": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceController",
		Description:    `Identifier of the source controller.`,
		Exposed:        true,
		Name:           "sourceController",
		Type:           "string",
	},
	"SourceID": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `ID of the source.`,
		Exposed:        true,
		Name:           "sourceID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"SourceIP": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceIP",
		Stored:         true,
		Type:           "string",
	},
	"SourceNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceNamespace",
		Deprecated:     true,
		Description: `Namespace of the source. This is deprecated. Use ` + "`" + `remoteNamespace` + "`" + `. This
property does nothing.`,
		Exposed: true,
		Name:    "sourceNamespace",
		Stored:  true,
		Type:    "string",
	},
	"SourcePlatform": {
		AllowedChoices: []string{},
		ConvertedName:  "SourcePlatform",
		Description:    `Identifier of the source platform.`,
		Exposed:        true,
		Name:           "sourcePlatform",
		Stored:         true,
		Type:           "string",
	},
	"SourceType": {
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Claims"},
		ConvertedName:  "SourceType",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceType",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"Timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Time and date of the log.`,
		Exposed:        true,
		Name:           "timestamp",
		Stored:         true,
		Type:           "time",
	},
	"Value": {
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Number of flows in the log.`,
		Exposed:        true,
		Name:           "value",
		Required:       true,
		Stored:         true,
		Type:           "integer",
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

// FlowReportLowerCaseAttributesMap represents the map of attribute for FlowReport.
var FlowReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"action": {
		AllowedChoices: []string{"Accept", "Reject"},
		ConvertedName:  "Action",
		Description:    `Action applied to the flow.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"destinationcontroller": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationController",
		Description:    `Identifier of the destination controller.`,
		Exposed:        true,
		Name:           "destinationController",
		Type:           "string",
	},
	"destinationid": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination.`,
		Exposed:        true,
		Name:           "destinationID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"destinationip": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationIP",
		Description:    `Destination IP address.`,
		Exposed:        true,
		Name:           "destinationIP",
		Stored:         true,
		Type:           "string",
	},
	"destinationnamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationNamespace",
		Deprecated:     true,
		Description: `Namespace of the destination. This is deprecated. Use ` + "`" + `remoteNamespace` + "`" + `. This
property does nothing.`,
		Exposed: true,
		Name:    "destinationNamespace",
		Stored:  true,
		Type:    "string",
	},
	"destinationplatform": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationPlatform",
		Description:    `Identifier of the destination platform.`,
		Exposed:        true,
		Name:           "destinationPlatform",
		Stored:         true,
		Type:           "string",
	},
	"destinationport": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationPort",
		Description:    `Port of the destination.`,
		Exposed:        true,
		Name:           "destinationPort",
		Stored:         true,
		Type:           "integer",
	},
	"destinationtype": {
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Claims"},
		ConvertedName:  "DestinationType",
		Description:    `Destination type.`,
		Exposed:        true,
		Name:           "destinationType",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"dropreason": {
		AllowedChoices: []string{},
		ConvertedName:  "DropReason",
		Description: `This field is only set if ` + "`" + `action` + "`" + ` is set to ` + "`" + `Reject` + "`" + `. It specifies the reason
for the rejection.`,
		Exposed: true,
		Name:    "dropReason",
		Stored:  true,
		Type:    "string",
	},
	"encrypted": {
		AllowedChoices: []string{},
		ConvertedName:  "Encrypted",
		Description:    `If ` + "`" + `true` + "`" + `, the flow was encrypted.`,
		Exposed:        true,
		Name:           "encrypted",
		Stored:         true,
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
	"namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Deprecated:     true,
		Description:    `This is here for backward compatibility.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"observed": {
		AllowedChoices: []string{},
		ConvertedName:  "Observed",
		Description:    `If ` + "`" + `true` + "`" + `, design mode is on.`,
		Exposed:        true,
		Name:           "observed",
		Stored:         true,
		Type:           "boolean",
	},
	"observedaction": {
		AllowedChoices: []string{"Accept", "Reject", "NotApplicable"},
		ConvertedName:  "ObservedAction",
		DefaultValue:   FlowReportObservedActionNotApplicable,
		Description:    `Action observed on the flow.`,
		Exposed:        true,
		Name:           "observedAction",
		Stored:         true,
		Type:           "enum",
	},
	"observeddropreason": {
		AllowedChoices: []string{},
		ConvertedName:  "ObservedDropReason",
		Description: `Specifies the reason for a rejection. Only set if ` + "`" + `observedAction` + "`" + ` is set
to ` + "`" + `Reject` + "`" + `.`,
		Exposed: true,
		Name:    "observedDropReason",
		Stored:  true,
		Type:    "string",
	},
	"observedencrypted": {
		AllowedChoices: []string{},
		ConvertedName:  "ObservedEncrypted",
		Description:    `Value of the encryption of the network policy that observed the flow.`,
		Exposed:        true,
		Name:           "observedEncrypted",
		Stored:         true,
		Type:           "boolean",
	},
	"observedpolicyid": {
		AllowedChoices: []string{},
		ConvertedName:  "ObservedPolicyID",
		Description:    `ID of the network policy that observed the flow.`,
		Exposed:        true,
		Name:           "observedPolicyID",
		Stored:         true,
		Type:           "string",
	},
	"observedpolicynamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ObservedPolicyNamespace",
		Description:    `Namespace of the network policy that observed the flow.`,
		Exposed:        true,
		Name:           "observedPolicyNamespace",
		Stored:         true,
		Type:           "string",
	},
	"policyid": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `ID of the network policy that accepted the flow.`,
		Exposed:        true,
		Name:           "policyID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"policynamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyNamespace",
		Description:    `Namespace of the network policy that accepted the flow.`,
		Exposed:        true,
		Name:           "policyNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"protocol": {
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol number.`,
		Exposed:        true,
		Name:           "protocol",
		Required:       true,
		Stored:         true,
		Type:           "integer",
	},
	"remotenamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "RemoteNamespace",
		Description:    `Namespace of the object at the other end of the flow.`,
		Exposed:        true,
		Name:           "remoteNamespace",
		Stored:         true,
		Type:           "string",
	},
	"serviceclaimhash": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceClaimHash",
		Description:    `Hash of the claims used to communicate.`,
		Exposed:        true,
		Name:           "serviceClaimHash",
		Stored:         true,
		Type:           "string",
	},
	"serviceid": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceID",
		Description:    `ID of the service.`,
		Exposed:        true,
		Name:           "serviceID",
		Stored:         true,
		Type:           "string",
	},
	"servicenamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceNamespace",
		Description:    `Namespace of Service accessed.`,
		Exposed:        true,
		Name:           "serviceNamespace",
		Stored:         true,
		Type:           "string",
	},
	"servicetype": {
		AllowedChoices: []string{"L3", "HTTP", "TCP", "NotApplicable"},
		ConvertedName:  "ServiceType",
		DefaultValue:   FlowReportServiceTypeNotApplicable,
		Description:    `ID of the service.`,
		Exposed:        true,
		Name:           "serviceType",
		Stored:         true,
		Type:           "enum",
	},
	"serviceurl": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceURL",
		Description:    `Service URL accessed.`,
		Exposed:        true,
		Name:           "serviceURL",
		Stored:         true,
		Type:           "string",
	},
	"sourcecontroller": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceController",
		Description:    `Identifier of the source controller.`,
		Exposed:        true,
		Name:           "sourceController",
		Type:           "string",
	},
	"sourceid": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `ID of the source.`,
		Exposed:        true,
		Name:           "sourceID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"sourceip": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceIP",
		Stored:         true,
		Type:           "string",
	},
	"sourcenamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceNamespace",
		Deprecated:     true,
		Description: `Namespace of the source. This is deprecated. Use ` + "`" + `remoteNamespace` + "`" + `. This
property does nothing.`,
		Exposed: true,
		Name:    "sourceNamespace",
		Stored:  true,
		Type:    "string",
	},
	"sourceplatform": {
		AllowedChoices: []string{},
		ConvertedName:  "SourcePlatform",
		Description:    `Identifier of the source platform.`,
		Exposed:        true,
		Name:           "sourcePlatform",
		Stored:         true,
		Type:           "string",
	},
	"sourcetype": {
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Claims"},
		ConvertedName:  "SourceType",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceType",
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
	"timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Time and date of the log.`,
		Exposed:        true,
		Name:           "timestamp",
		Stored:         true,
		Type:           "time",
	},
	"value": {
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Number of flows in the log.`,
		Exposed:        true,
		Name:           "value",
		Required:       true,
		Stored:         true,
		Type:           "integer",
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

// SparseFlowReportsList represents a list of SparseFlowReports
type SparseFlowReportsList []*SparseFlowReport

// Identity returns the identity of the objects in the list.
func (o SparseFlowReportsList) Identity() elemental.Identity {

	return FlowReportIdentity
}

// Copy returns a pointer to a copy the SparseFlowReportsList.
func (o SparseFlowReportsList) Copy() elemental.Identifiables {

	copy := append(SparseFlowReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseFlowReportsList.
func (o SparseFlowReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseFlowReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseFlowReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseFlowReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseFlowReportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseFlowReportsList converted to FlowReportsList.
func (o SparseFlowReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseFlowReportsList) Version() int {

	return 1
}

// SparseFlowReport represents the sparse version of a flowreport.
type SparseFlowReport struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Action applied to the flow.
	Action *FlowReportActionValue `json:"action,omitempty" msgpack:"action,omitempty" bson:"action,omitempty" mapstructure:"action,omitempty"`

	// Identifier of the destination controller.
	DestinationController *string `json:"destinationController,omitempty" msgpack:"destinationController,omitempty" bson:"-" mapstructure:"destinationController,omitempty"`

	// ID of the destination.
	DestinationID *string `json:"destinationID,omitempty" msgpack:"destinationID,omitempty" bson:"destinationid,omitempty" mapstructure:"destinationID,omitempty"`

	// Destination IP address.
	DestinationIP *string `json:"destinationIP,omitempty" msgpack:"destinationIP,omitempty" bson:"destinationip,omitempty" mapstructure:"destinationIP,omitempty"`

	// Namespace of the destination. This is deprecated. Use `remoteNamespace`. This
	// property does nothing.
	DestinationNamespace *string `json:"destinationNamespace,omitempty" msgpack:"destinationNamespace,omitempty" bson:"destinationnamespace,omitempty" mapstructure:"destinationNamespace,omitempty"`

	// Identifier of the destination platform.
	DestinationPlatform *string `json:"destinationPlatform,omitempty" msgpack:"destinationPlatform,omitempty" bson:"destinationplatform,omitempty" mapstructure:"destinationPlatform,omitempty"`

	// Port of the destination.
	DestinationPort *int `json:"destinationPort,omitempty" msgpack:"destinationPort,omitempty" bson:"destinationport,omitempty" mapstructure:"destinationPort,omitempty"`

	// Destination type.
	DestinationType *FlowReportDestinationTypeValue `json:"destinationType,omitempty" msgpack:"destinationType,omitempty" bson:"destinationtype,omitempty" mapstructure:"destinationType,omitempty"`

	// This field is only set if `action` is set to `Reject`. It specifies the reason
	// for the rejection.
	DropReason *string `json:"dropReason,omitempty" msgpack:"dropReason,omitempty" bson:"dropreason,omitempty" mapstructure:"dropReason,omitempty"`

	// If `true`, the flow was encrypted.
	Encrypted *bool `json:"encrypted,omitempty" msgpack:"encrypted,omitempty" bson:"encrypted,omitempty" mapstructure:"encrypted,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// This is here for backward compatibility.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// If `true`, design mode is on.
	Observed *bool `json:"observed,omitempty" msgpack:"observed,omitempty" bson:"observed,omitempty" mapstructure:"observed,omitempty"`

	// Action observed on the flow.
	ObservedAction *FlowReportObservedActionValue `json:"observedAction,omitempty" msgpack:"observedAction,omitempty" bson:"observedaction,omitempty" mapstructure:"observedAction,omitempty"`

	// Specifies the reason for a rejection. Only set if `observedAction` is set
	// to `Reject`.
	ObservedDropReason *string `json:"observedDropReason,omitempty" msgpack:"observedDropReason,omitempty" bson:"observeddropreason,omitempty" mapstructure:"observedDropReason,omitempty"`

	// Value of the encryption of the network policy that observed the flow.
	ObservedEncrypted *bool `json:"observedEncrypted,omitempty" msgpack:"observedEncrypted,omitempty" bson:"observedencrypted,omitempty" mapstructure:"observedEncrypted,omitempty"`

	// ID of the network policy that observed the flow.
	ObservedPolicyID *string `json:"observedPolicyID,omitempty" msgpack:"observedPolicyID,omitempty" bson:"observedpolicyid,omitempty" mapstructure:"observedPolicyID,omitempty"`

	// Namespace of the network policy that observed the flow.
	ObservedPolicyNamespace *string `json:"observedPolicyNamespace,omitempty" msgpack:"observedPolicyNamespace,omitempty" bson:"observedpolicynamespace,omitempty" mapstructure:"observedPolicyNamespace,omitempty"`

	// ID of the network policy that accepted the flow.
	PolicyID *string `json:"policyID,omitempty" msgpack:"policyID,omitempty" bson:"policyid,omitempty" mapstructure:"policyID,omitempty"`

	// Namespace of the network policy that accepted the flow.
	PolicyNamespace *string `json:"policyNamespace,omitempty" msgpack:"policyNamespace,omitempty" bson:"policynamespace,omitempty" mapstructure:"policyNamespace,omitempty"`

	// Protocol number.
	Protocol *int `json:"protocol,omitempty" msgpack:"protocol,omitempty" bson:"protocol,omitempty" mapstructure:"protocol,omitempty"`

	// Namespace of the object at the other end of the flow.
	RemoteNamespace *string `json:"remoteNamespace,omitempty" msgpack:"remoteNamespace,omitempty" bson:"remotenamespace,omitempty" mapstructure:"remoteNamespace,omitempty"`

	// Hash of the claims used to communicate.
	ServiceClaimHash *string `json:"serviceClaimHash,omitempty" msgpack:"serviceClaimHash,omitempty" bson:"serviceclaimhash,omitempty" mapstructure:"serviceClaimHash,omitempty"`

	// ID of the service.
	ServiceID *string `json:"serviceID,omitempty" msgpack:"serviceID,omitempty" bson:"serviceid,omitempty" mapstructure:"serviceID,omitempty"`

	// Namespace of Service accessed.
	ServiceNamespace *string `json:"serviceNamespace,omitempty" msgpack:"serviceNamespace,omitempty" bson:"servicenamespace,omitempty" mapstructure:"serviceNamespace,omitempty"`

	// ID of the service.
	ServiceType *FlowReportServiceTypeValue `json:"serviceType,omitempty" msgpack:"serviceType,omitempty" bson:"servicetype,omitempty" mapstructure:"serviceType,omitempty"`

	// Service URL accessed.
	ServiceURL *string `json:"serviceURL,omitempty" msgpack:"serviceURL,omitempty" bson:"serviceurl,omitempty" mapstructure:"serviceURL,omitempty"`

	// Identifier of the source controller.
	SourceController *string `json:"sourceController,omitempty" msgpack:"sourceController,omitempty" bson:"-" mapstructure:"sourceController,omitempty"`

	// ID of the source.
	SourceID *string `json:"sourceID,omitempty" msgpack:"sourceID,omitempty" bson:"sourceid,omitempty" mapstructure:"sourceID,omitempty"`

	// Type of the source.
	SourceIP *string `json:"sourceIP,omitempty" msgpack:"sourceIP,omitempty" bson:"sourceip,omitempty" mapstructure:"sourceIP,omitempty"`

	// Namespace of the source. This is deprecated. Use `remoteNamespace`. This
	// property does nothing.
	SourceNamespace *string `json:"sourceNamespace,omitempty" msgpack:"sourceNamespace,omitempty" bson:"sourcenamespace,omitempty" mapstructure:"sourceNamespace,omitempty"`

	// Identifier of the source platform.
	SourcePlatform *string `json:"sourcePlatform,omitempty" msgpack:"sourcePlatform,omitempty" bson:"sourceplatform,omitempty" mapstructure:"sourcePlatform,omitempty"`

	// Type of the source.
	SourceType *FlowReportSourceTypeValue `json:"sourceType,omitempty" msgpack:"sourceType,omitempty" bson:"sourcetype,omitempty" mapstructure:"sourceType,omitempty"`

	// Time and date of the log.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"timestamp,omitempty" mapstructure:"timestamp,omitempty"`

	// Number of flows in the log.
	Value *int `json:"value,omitempty" msgpack:"value,omitempty" bson:"value,omitempty" mapstructure:"value,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseFlowReport returns a new  SparseFlowReport.
func NewSparseFlowReport() *SparseFlowReport {
	return &SparseFlowReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseFlowReport) Identity() elemental.Identity {

	return FlowReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseFlowReport) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseFlowReport) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseFlowReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseFlowReport{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Action != nil {
		s.Action = o.Action
	}
	if o.DestinationID != nil {
		s.DestinationID = o.DestinationID
	}
	if o.DestinationIP != nil {
		s.DestinationIP = o.DestinationIP
	}
	if o.DestinationNamespace != nil {
		s.DestinationNamespace = o.DestinationNamespace
	}
	if o.DestinationPlatform != nil {
		s.DestinationPlatform = o.DestinationPlatform
	}
	if o.DestinationPort != nil {
		s.DestinationPort = o.DestinationPort
	}
	if o.DestinationType != nil {
		s.DestinationType = o.DestinationType
	}
	if o.DropReason != nil {
		s.DropReason = o.DropReason
	}
	if o.Encrypted != nil {
		s.Encrypted = o.Encrypted
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.Observed != nil {
		s.Observed = o.Observed
	}
	if o.ObservedAction != nil {
		s.ObservedAction = o.ObservedAction
	}
	if o.ObservedDropReason != nil {
		s.ObservedDropReason = o.ObservedDropReason
	}
	if o.ObservedEncrypted != nil {
		s.ObservedEncrypted = o.ObservedEncrypted
	}
	if o.ObservedPolicyID != nil {
		s.ObservedPolicyID = o.ObservedPolicyID
	}
	if o.ObservedPolicyNamespace != nil {
		s.ObservedPolicyNamespace = o.ObservedPolicyNamespace
	}
	if o.PolicyID != nil {
		s.PolicyID = o.PolicyID
	}
	if o.PolicyNamespace != nil {
		s.PolicyNamespace = o.PolicyNamespace
	}
	if o.Protocol != nil {
		s.Protocol = o.Protocol
	}
	if o.RemoteNamespace != nil {
		s.RemoteNamespace = o.RemoteNamespace
	}
	if o.ServiceClaimHash != nil {
		s.ServiceClaimHash = o.ServiceClaimHash
	}
	if o.ServiceID != nil {
		s.ServiceID = o.ServiceID
	}
	if o.ServiceNamespace != nil {
		s.ServiceNamespace = o.ServiceNamespace
	}
	if o.ServiceType != nil {
		s.ServiceType = o.ServiceType
	}
	if o.ServiceURL != nil {
		s.ServiceURL = o.ServiceURL
	}
	if o.SourceID != nil {
		s.SourceID = o.SourceID
	}
	if o.SourceIP != nil {
		s.SourceIP = o.SourceIP
	}
	if o.SourceNamespace != nil {
		s.SourceNamespace = o.SourceNamespace
	}
	if o.SourcePlatform != nil {
		s.SourcePlatform = o.SourcePlatform
	}
	if o.SourceType != nil {
		s.SourceType = o.SourceType
	}
	if o.Timestamp != nil {
		s.Timestamp = o.Timestamp
	}
	if o.Value != nil {
		s.Value = o.Value
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
func (o *SparseFlowReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseFlowReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Action != nil {
		o.Action = s.Action
	}
	if s.DestinationID != nil {
		o.DestinationID = s.DestinationID
	}
	if s.DestinationIP != nil {
		o.DestinationIP = s.DestinationIP
	}
	if s.DestinationNamespace != nil {
		o.DestinationNamespace = s.DestinationNamespace
	}
	if s.DestinationPlatform != nil {
		o.DestinationPlatform = s.DestinationPlatform
	}
	if s.DestinationPort != nil {
		o.DestinationPort = s.DestinationPort
	}
	if s.DestinationType != nil {
		o.DestinationType = s.DestinationType
	}
	if s.DropReason != nil {
		o.DropReason = s.DropReason
	}
	if s.Encrypted != nil {
		o.Encrypted = s.Encrypted
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.Observed != nil {
		o.Observed = s.Observed
	}
	if s.ObservedAction != nil {
		o.ObservedAction = s.ObservedAction
	}
	if s.ObservedDropReason != nil {
		o.ObservedDropReason = s.ObservedDropReason
	}
	if s.ObservedEncrypted != nil {
		o.ObservedEncrypted = s.ObservedEncrypted
	}
	if s.ObservedPolicyID != nil {
		o.ObservedPolicyID = s.ObservedPolicyID
	}
	if s.ObservedPolicyNamespace != nil {
		o.ObservedPolicyNamespace = s.ObservedPolicyNamespace
	}
	if s.PolicyID != nil {
		o.PolicyID = s.PolicyID
	}
	if s.PolicyNamespace != nil {
		o.PolicyNamespace = s.PolicyNamespace
	}
	if s.Protocol != nil {
		o.Protocol = s.Protocol
	}
	if s.RemoteNamespace != nil {
		o.RemoteNamespace = s.RemoteNamespace
	}
	if s.ServiceClaimHash != nil {
		o.ServiceClaimHash = s.ServiceClaimHash
	}
	if s.ServiceID != nil {
		o.ServiceID = s.ServiceID
	}
	if s.ServiceNamespace != nil {
		o.ServiceNamespace = s.ServiceNamespace
	}
	if s.ServiceType != nil {
		o.ServiceType = s.ServiceType
	}
	if s.ServiceURL != nil {
		o.ServiceURL = s.ServiceURL
	}
	if s.SourceID != nil {
		o.SourceID = s.SourceID
	}
	if s.SourceIP != nil {
		o.SourceIP = s.SourceIP
	}
	if s.SourceNamespace != nil {
		o.SourceNamespace = s.SourceNamespace
	}
	if s.SourcePlatform != nil {
		o.SourcePlatform = s.SourcePlatform
	}
	if s.SourceType != nil {
		o.SourceType = s.SourceType
	}
	if s.Timestamp != nil {
		o.Timestamp = s.Timestamp
	}
	if s.Value != nil {
		o.Value = s.Value
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
func (o *SparseFlowReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseFlowReport) ToPlain() elemental.PlainIdentifiable {

	out := NewFlowReport()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Action != nil {
		out.Action = *o.Action
	}
	if o.DestinationController != nil {
		out.DestinationController = *o.DestinationController
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
	if o.DestinationPlatform != nil {
		out.DestinationPlatform = *o.DestinationPlatform
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
	if o.Encrypted != nil {
		out.Encrypted = *o.Encrypted
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Observed != nil {
		out.Observed = *o.Observed
	}
	if o.ObservedAction != nil {
		out.ObservedAction = *o.ObservedAction
	}
	if o.ObservedDropReason != nil {
		out.ObservedDropReason = *o.ObservedDropReason
	}
	if o.ObservedEncrypted != nil {
		out.ObservedEncrypted = *o.ObservedEncrypted
	}
	if o.ObservedPolicyID != nil {
		out.ObservedPolicyID = *o.ObservedPolicyID
	}
	if o.ObservedPolicyNamespace != nil {
		out.ObservedPolicyNamespace = *o.ObservedPolicyNamespace
	}
	if o.PolicyID != nil {
		out.PolicyID = *o.PolicyID
	}
	if o.PolicyNamespace != nil {
		out.PolicyNamespace = *o.PolicyNamespace
	}
	if o.Protocol != nil {
		out.Protocol = *o.Protocol
	}
	if o.RemoteNamespace != nil {
		out.RemoteNamespace = *o.RemoteNamespace
	}
	if o.ServiceClaimHash != nil {
		out.ServiceClaimHash = *o.ServiceClaimHash
	}
	if o.ServiceID != nil {
		out.ServiceID = *o.ServiceID
	}
	if o.ServiceNamespace != nil {
		out.ServiceNamespace = *o.ServiceNamespace
	}
	if o.ServiceType != nil {
		out.ServiceType = *o.ServiceType
	}
	if o.ServiceURL != nil {
		out.ServiceURL = *o.ServiceURL
	}
	if o.SourceController != nil {
		out.SourceController = *o.SourceController
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
	if o.SourcePlatform != nil {
		out.SourcePlatform = *o.SourcePlatform
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
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseFlowReport) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseFlowReport) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseFlowReport) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseFlowReport) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseFlowReport) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseFlowReport) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseFlowReport.
func (o *SparseFlowReport) DeepCopy() *SparseFlowReport {

	if o == nil {
		return nil
	}

	out := &SparseFlowReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseFlowReport.
func (o *SparseFlowReport) DeepCopyInto(out *SparseFlowReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseFlowReport: %s", err))
	}

	*out = *target.(*SparseFlowReport)
}

type mongoAttributesFlowReport struct {
	ID                      bson.ObjectId                  `bson:"_id,omitempty"`
	Action                  FlowReportActionValue          `bson:"action"`
	DestinationID           string                         `bson:"destinationid"`
	DestinationIP           string                         `bson:"destinationip"`
	DestinationNamespace    string                         `bson:"destinationnamespace,omitempty"`
	DestinationPlatform     string                         `bson:"destinationplatform"`
	DestinationPort         int                            `bson:"destinationport"`
	DestinationType         FlowReportDestinationTypeValue `bson:"destinationtype"`
	DropReason              string                         `bson:"dropreason"`
	Encrypted               bool                           `bson:"encrypted"`
	MigrationsLog           map[string]string              `bson:"migrationslog,omitempty"`
	Namespace               string                         `bson:"namespace"`
	Observed                bool                           `bson:"observed"`
	ObservedAction          FlowReportObservedActionValue  `bson:"observedaction"`
	ObservedDropReason      string                         `bson:"observeddropreason"`
	ObservedEncrypted       bool                           `bson:"observedencrypted"`
	ObservedPolicyID        string                         `bson:"observedpolicyid"`
	ObservedPolicyNamespace string                         `bson:"observedpolicynamespace"`
	PolicyID                string                         `bson:"policyid"`
	PolicyNamespace         string                         `bson:"policynamespace"`
	Protocol                int                            `bson:"protocol"`
	RemoteNamespace         string                         `bson:"remotenamespace,omitempty"`
	ServiceClaimHash        string                         `bson:"serviceclaimhash"`
	ServiceID               string                         `bson:"serviceid"`
	ServiceNamespace        string                         `bson:"servicenamespace"`
	ServiceType             FlowReportServiceTypeValue     `bson:"servicetype"`
	ServiceURL              string                         `bson:"serviceurl"`
	SourceID                string                         `bson:"sourceid"`
	SourceIP                string                         `bson:"sourceip"`
	SourceNamespace         string                         `bson:"sourcenamespace,omitempty"`
	SourcePlatform          string                         `bson:"sourceplatform"`
	SourceType              FlowReportSourceTypeValue      `bson:"sourcetype"`
	Timestamp               time.Time                      `bson:"timestamp"`
	Value                   int                            `bson:"value"`
	ZHash                   int                            `bson:"zhash"`
	Zone                    int                            `bson:"zone"`
}
type mongoAttributesSparseFlowReport struct {
	ID                      bson.ObjectId                   `bson:"_id,omitempty"`
	Action                  *FlowReportActionValue          `bson:"action,omitempty"`
	DestinationID           *string                         `bson:"destinationid,omitempty"`
	DestinationIP           *string                         `bson:"destinationip,omitempty"`
	DestinationNamespace    *string                         `bson:"destinationnamespace,omitempty"`
	DestinationPlatform     *string                         `bson:"destinationplatform,omitempty"`
	DestinationPort         *int                            `bson:"destinationport,omitempty"`
	DestinationType         *FlowReportDestinationTypeValue `bson:"destinationtype,omitempty"`
	DropReason              *string                         `bson:"dropreason,omitempty"`
	Encrypted               *bool                           `bson:"encrypted,omitempty"`
	MigrationsLog           *map[string]string              `bson:"migrationslog,omitempty"`
	Namespace               *string                         `bson:"namespace,omitempty"`
	Observed                *bool                           `bson:"observed,omitempty"`
	ObservedAction          *FlowReportObservedActionValue  `bson:"observedaction,omitempty"`
	ObservedDropReason      *string                         `bson:"observeddropreason,omitempty"`
	ObservedEncrypted       *bool                           `bson:"observedencrypted,omitempty"`
	ObservedPolicyID        *string                         `bson:"observedpolicyid,omitempty"`
	ObservedPolicyNamespace *string                         `bson:"observedpolicynamespace,omitempty"`
	PolicyID                *string                         `bson:"policyid,omitempty"`
	PolicyNamespace         *string                         `bson:"policynamespace,omitempty"`
	Protocol                *int                            `bson:"protocol,omitempty"`
	RemoteNamespace         *string                         `bson:"remotenamespace,omitempty"`
	ServiceClaimHash        *string                         `bson:"serviceclaimhash,omitempty"`
	ServiceID               *string                         `bson:"serviceid,omitempty"`
	ServiceNamespace        *string                         `bson:"servicenamespace,omitempty"`
	ServiceType             *FlowReportServiceTypeValue     `bson:"servicetype,omitempty"`
	ServiceURL              *string                         `bson:"serviceurl,omitempty"`
	SourceID                *string                         `bson:"sourceid,omitempty"`
	SourceIP                *string                         `bson:"sourceip,omitempty"`
	SourceNamespace         *string                         `bson:"sourcenamespace,omitempty"`
	SourcePlatform          *string                         `bson:"sourceplatform,omitempty"`
	SourceType              *FlowReportSourceTypeValue      `bson:"sourcetype,omitempty"`
	Timestamp               *time.Time                      `bson:"timestamp,omitempty"`
	Value                   *int                            `bson:"value,omitempty"`
	ZHash                   *int                            `bson:"zhash,omitempty"`
	Zone                    *int                            `bson:"zone,omitempty"`
}
