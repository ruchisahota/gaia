package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PingProbeClaimsTypeValue represents the possible values for attribute "claimsType".
type PingProbeClaimsTypeValue string

const (
	// PingProbeClaimsTypeReceived represents the value Received.
	PingProbeClaimsTypeReceived PingProbeClaimsTypeValue = "Received"

	// PingProbeClaimsTypeTransmitted represents the value Transmitted.
	PingProbeClaimsTypeTransmitted PingProbeClaimsTypeValue = "Transmitted"
)

// PingProbeRemoteEndpointTypeValue represents the possible values for attribute "remoteEndpointType".
type PingProbeRemoteEndpointTypeValue string

const (
	// PingProbeRemoteEndpointTypeExternal represents the value External.
	PingProbeRemoteEndpointTypeExternal PingProbeRemoteEndpointTypeValue = "External"

	// PingProbeRemoteEndpointTypeProcessingUnit represents the value ProcessingUnit.
	PingProbeRemoteEndpointTypeProcessingUnit PingProbeRemoteEndpointTypeValue = "ProcessingUnit"
)

// PingProbeRemoteNamespaceTypeValue represents the possible values for attribute "remoteNamespaceType".
type PingProbeRemoteNamespaceTypeValue string

const (
	// PingProbeRemoteNamespaceTypeHash represents the value Hash.
	PingProbeRemoteNamespaceTypeHash PingProbeRemoteNamespaceTypeValue = "Hash"

	// PingProbeRemoteNamespaceTypePlain represents the value Plain.
	PingProbeRemoteNamespaceTypePlain PingProbeRemoteNamespaceTypeValue = "Plain"
)

// PingProbeTypeValue represents the possible values for attribute "type".
type PingProbeTypeValue string

const (
	// PingProbeTypeRequest represents the value Request.
	PingProbeTypeRequest PingProbeTypeValue = "Request"

	// PingProbeTypeResponse represents the value Response.
	PingProbeTypeResponse PingProbeTypeValue = "Response"
)

// PingProbeIdentity represents the Identity of the object.
var PingProbeIdentity = elemental.Identity{
	Name:     "pingprobe",
	Category: "pingprobes",
	Package:  "guy",
	Private:  false,
}

// PingProbesList represents a list of PingProbes
type PingProbesList []*PingProbe

// Identity returns the identity of the objects in the list.
func (o PingProbesList) Identity() elemental.Identity {

	return PingProbeIdentity
}

// Copy returns a pointer to a copy the PingProbesList.
func (o PingProbesList) Copy() elemental.Identifiables {

	copy := append(PingProbesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PingProbesList.
func (o PingProbesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PingProbesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PingProbe))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PingProbesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PingProbesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PingProbesList converted to SparsePingProbesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PingProbesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePingProbesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePingProbe)
	}

	return out
}

// Version returns the version of the content.
func (o PingProbesList) Version() int {

	return 1
}

// PingProbe represents the model of a pingprobe
type PingProbe struct {
	// Action of the ACL policy.
	ACLPolicyAction string `json:"ACLPolicyAction" msgpack:"ACLPolicyAction" bson:"aclpolicyaction" mapstructure:"ACLPolicyAction,omitempty"`

	// ID of the ACL policy.
	ACLPolicyID string `json:"ACLPolicyID" msgpack:"ACLPolicyID" bson:"aclpolicyid" mapstructure:"ACLPolicyID,omitempty"`

	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Time taken for a single request-response to complete.
	RTT string `json:"RTT" msgpack:"RTT" bson:"rtt" mapstructure:"RTT,omitempty"`

	// If true, application responded to the request.
	ApplicationListening bool `json:"applicationListening" msgpack:"applicationListening" bson:"applicationlistening" mapstructure:"applicationListening,omitempty"`

	// Claims of the processing unit.
	Claims []string `json:"claims" msgpack:"claims" bson:"claims" mapstructure:"claims,omitempty"`

	// Type of claims reported.
	ClaimsType PingProbeClaimsTypeValue `json:"claimsType" msgpack:"claimsType" bson:"claimstype" mapstructure:"claimsType,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// ID of the defender.
	EnforcerID string `json:"enforcerID" msgpack:"enforcerID" bson:"enforcerid" mapstructure:"enforcerID,omitempty"`

	// Namespace of the defender.
	EnforcerNamespace string `json:"enforcerNamespace" msgpack:"enforcerNamespace" bson:"enforcernamespace" mapstructure:"enforcerNamespace,omitempty"`

	// Semantic version of the defender.
	EnforcerVersion string `json:"enforcerVersion" msgpack:"enforcerVersion" bson:"enforcerversion" mapstructure:"enforcerVersion,omitempty"`

	// A non-empty error indicates a failure.
	Error string `json:"error" msgpack:"error" bson:"error" mapstructure:"error,omitempty"`

	// If true, destination IP is in `excludedNetworks`.
	ExcludedNetworks bool `json:"excludedNetworks" msgpack:"excludedNetworks" bson:"excludednetworks" mapstructure:"excludedNetworks,omitempty"`

	// Four tuple in the format <sip:dip:spt:dpt>.
	FourTuple string `json:"fourTuple" msgpack:"fourTuple" bson:"fourtuple" mapstructure:"fourTuple,omitempty"`

	// If true, the report was generated by the server.
	IsServer bool `json:"isServer" msgpack:"isServer" bson:"isserver" mapstructure:"isServer,omitempty"`

	// Holds the iteration number this probe is attached to.
	IterationIndex int `json:"iterationIndex" msgpack:"iterationIndex" bson:"iterationindex" mapstructure:"iterationIndex,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Size of the payload attached to the packet.
	PayloadSize int `json:"payloadSize" msgpack:"payloadSize" bson:"payloadsize" mapstructure:"payloadSize,omitempty"`

	// Represents the expiry of the peer certificate.
	PeerCertExpiry string `json:"peerCertExpiry" msgpack:"peerCertExpiry" bson:"peercertexpiry" mapstructure:"peerCertExpiry,omitempty"`

	// Represents the issuer of the peer certificate.
	PeerCertIssuer string `json:"peerCertIssuer" msgpack:"peerCertIssuer" bson:"peercertissuer" mapstructure:"peerCertIssuer,omitempty"`

	// Represents the subject of the peer certificate.
	PeerCertSubject string `json:"peerCertSubject" msgpack:"peerCertSubject" bson:"peercertsubject" mapstructure:"peerCertSubject,omitempty"`

	// PingID unique to a single ping control.
	PingID string `json:"pingID" msgpack:"pingID" bson:"pingid" mapstructure:"pingID,omitempty"`

	// Action of the policy.
	PolicyAction string `json:"policyAction" msgpack:"policyAction" bson:"policyaction" mapstructure:"policyAction,omitempty"`

	// ID of the policy.
	PolicyID string `json:"policyID" msgpack:"policyID" bson:"policyid" mapstructure:"policyID,omitempty"`

	// ID of the policy.
	PolicyNamespace string `json:"policyNamespace" msgpack:"policyNamespace" bson:"policynamespace" mapstructure:"policyNamespace,omitempty"`

	// ID of the reporting processing unit.
	ProcessingUnitID string `json:"processingUnitID" msgpack:"processingUnitID" bson:"processingunitid" mapstructure:"processingUnitID,omitempty"`

	// Protocol used for the communication.
	Protocol int `json:"protocol" msgpack:"protocol" bson:"protocol" mapstructure:"protocol,omitempty"`

	// Controller of the remote endpoint.
	RemoteController string `json:"remoteController" msgpack:"remoteController" bson:"remotecontroller" mapstructure:"remoteController,omitempty"`

	// Represents the remote endpoint type.
	RemoteEndpointType PingProbeRemoteEndpointTypeValue `json:"remoteEndpointType" msgpack:"remoteEndpointType" bson:"remoteendpointtype" mapstructure:"remoteEndpointType,omitempty"`

	// Namespace of the remote processing unit.
	RemoteNamespace string `json:"remoteNamespace" msgpack:"remoteNamespace" bson:"remotenamespace" mapstructure:"remoteNamespace,omitempty"`

	// Type of the namespace reported.
	RemoteNamespaceType PingProbeRemoteNamespaceTypeValue `json:"remoteNamespaceType" msgpack:"remoteNamespaceType" bson:"remotenamespacetype" mapstructure:"remoteNamespaceType,omitempty"`

	// ID of the remote processing unit.
	RemoteProcessingUnitID string `json:"remoteProcessingUnitID" msgpack:"remoteProcessingUnitID" bson:"remoteprocessingunitid" mapstructure:"remoteProcessingUnitID,omitempty"`

	// Sequence number of the TCP packet. number.
	SeqNum int `json:"seqNum" msgpack:"seqNum" bson:"seqnum" mapstructure:"seqNum,omitempty"`

	// ID of the service If the service type is a proxy.
	ServiceID string `json:"serviceID" msgpack:"serviceID" bson:"serviceid" mapstructure:"serviceID,omitempty"`

	// Type of the service.
	ServiceType string `json:"serviceType" msgpack:"serviceType" bson:"servicetype" mapstructure:"serviceType,omitempty"`

	// If true, destination IP is in `targetTCPNetworks`.
	TargetTCPNetworks bool `json:"targetTCPNetworks" msgpack:"targetTCPNetworks" bson:"targettcpnetworks" mapstructure:"targetTCPNetworks,omitempty"`

	// Type of the report.
	Type PingProbeTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPingProbe returns a new *PingProbe
func NewPingProbe() *PingProbe {

	return &PingProbe{
		ModelVersion:  1,
		Claims:        []string{},
		MigrationsLog: map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *PingProbe) Identity() elemental.Identity {

	return PingProbeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PingProbe) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PingProbe) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PingProbe) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPingProbe{}

	s.ACLPolicyAction = o.ACLPolicyAction
	s.ACLPolicyID = o.ACLPolicyID
	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.RTT = o.RTT
	s.ApplicationListening = o.ApplicationListening
	s.Claims = o.Claims
	s.ClaimsType = o.ClaimsType
	s.CreateTime = o.CreateTime
	s.EnforcerID = o.EnforcerID
	s.EnforcerNamespace = o.EnforcerNamespace
	s.EnforcerVersion = o.EnforcerVersion
	s.Error = o.Error
	s.ExcludedNetworks = o.ExcludedNetworks
	s.FourTuple = o.FourTuple
	s.IsServer = o.IsServer
	s.IterationIndex = o.IterationIndex
	s.MigrationsLog = o.MigrationsLog
	s.Namespace = o.Namespace
	s.PayloadSize = o.PayloadSize
	s.PeerCertExpiry = o.PeerCertExpiry
	s.PeerCertIssuer = o.PeerCertIssuer
	s.PeerCertSubject = o.PeerCertSubject
	s.PingID = o.PingID
	s.PolicyAction = o.PolicyAction
	s.PolicyID = o.PolicyID
	s.PolicyNamespace = o.PolicyNamespace
	s.ProcessingUnitID = o.ProcessingUnitID
	s.Protocol = o.Protocol
	s.RemoteController = o.RemoteController
	s.RemoteEndpointType = o.RemoteEndpointType
	s.RemoteNamespace = o.RemoteNamespace
	s.RemoteNamespaceType = o.RemoteNamespaceType
	s.RemoteProcessingUnitID = o.RemoteProcessingUnitID
	s.SeqNum = o.SeqNum
	s.ServiceID = o.ServiceID
	s.ServiceType = o.ServiceType
	s.TargetTCPNetworks = o.TargetTCPNetworks
	s.Type = o.Type
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PingProbe) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPingProbe{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ACLPolicyAction = s.ACLPolicyAction
	o.ACLPolicyID = s.ACLPolicyID
	o.ID = s.ID.Hex()
	o.RTT = s.RTT
	o.ApplicationListening = s.ApplicationListening
	o.Claims = s.Claims
	o.ClaimsType = s.ClaimsType
	o.CreateTime = s.CreateTime
	o.EnforcerID = s.EnforcerID
	o.EnforcerNamespace = s.EnforcerNamespace
	o.EnforcerVersion = s.EnforcerVersion
	o.Error = s.Error
	o.ExcludedNetworks = s.ExcludedNetworks
	o.FourTuple = s.FourTuple
	o.IsServer = s.IsServer
	o.IterationIndex = s.IterationIndex
	o.MigrationsLog = s.MigrationsLog
	o.Namespace = s.Namespace
	o.PayloadSize = s.PayloadSize
	o.PeerCertExpiry = s.PeerCertExpiry
	o.PeerCertIssuer = s.PeerCertIssuer
	o.PeerCertSubject = s.PeerCertSubject
	o.PingID = s.PingID
	o.PolicyAction = s.PolicyAction
	o.PolicyID = s.PolicyID
	o.PolicyNamespace = s.PolicyNamespace
	o.ProcessingUnitID = s.ProcessingUnitID
	o.Protocol = s.Protocol
	o.RemoteController = s.RemoteController
	o.RemoteEndpointType = s.RemoteEndpointType
	o.RemoteNamespace = s.RemoteNamespace
	o.RemoteNamespaceType = s.RemoteNamespaceType
	o.RemoteProcessingUnitID = s.RemoteProcessingUnitID
	o.SeqNum = s.SeqNum
	o.ServiceID = s.ServiceID
	o.ServiceType = s.ServiceType
	o.TargetTCPNetworks = s.TargetTCPNetworks
	o.Type = s.Type
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *PingProbe) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PingProbe) BleveType() string {

	return "pingprobe"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PingProbe) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PingProbe) Doc() string {

	return `Represents the result of a unique ping probe. They are aggregated into a
PingResult.`
}

func (o *PingProbe) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *PingProbe) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *PingProbe) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *PingProbe) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *PingProbe) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *PingProbe) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *PingProbe) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *PingProbe) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *PingProbe) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *PingProbe) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *PingProbe) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *PingProbe) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *PingProbe) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PingProbe) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePingProbe{
			ACLPolicyAction:        &o.ACLPolicyAction,
			ACLPolicyID:            &o.ACLPolicyID,
			ID:                     &o.ID,
			RTT:                    &o.RTT,
			ApplicationListening:   &o.ApplicationListening,
			Claims:                 &o.Claims,
			ClaimsType:             &o.ClaimsType,
			CreateTime:             &o.CreateTime,
			EnforcerID:             &o.EnforcerID,
			EnforcerNamespace:      &o.EnforcerNamespace,
			EnforcerVersion:        &o.EnforcerVersion,
			Error:                  &o.Error,
			ExcludedNetworks:       &o.ExcludedNetworks,
			FourTuple:              &o.FourTuple,
			IsServer:               &o.IsServer,
			IterationIndex:         &o.IterationIndex,
			MigrationsLog:          &o.MigrationsLog,
			Namespace:              &o.Namespace,
			PayloadSize:            &o.PayloadSize,
			PeerCertExpiry:         &o.PeerCertExpiry,
			PeerCertIssuer:         &o.PeerCertIssuer,
			PeerCertSubject:        &o.PeerCertSubject,
			PingID:                 &o.PingID,
			PolicyAction:           &o.PolicyAction,
			PolicyID:               &o.PolicyID,
			PolicyNamespace:        &o.PolicyNamespace,
			ProcessingUnitID:       &o.ProcessingUnitID,
			Protocol:               &o.Protocol,
			RemoteController:       &o.RemoteController,
			RemoteEndpointType:     &o.RemoteEndpointType,
			RemoteNamespace:        &o.RemoteNamespace,
			RemoteNamespaceType:    &o.RemoteNamespaceType,
			RemoteProcessingUnitID: &o.RemoteProcessingUnitID,
			SeqNum:                 &o.SeqNum,
			ServiceID:              &o.ServiceID,
			ServiceType:            &o.ServiceType,
			TargetTCPNetworks:      &o.TargetTCPNetworks,
			Type:                   &o.Type,
			UpdateTime:             &o.UpdateTime,
			ZHash:                  &o.ZHash,
			Zone:                   &o.Zone,
		}
	}

	sp := &SparsePingProbe{}
	for _, f := range fields {
		switch f {
		case "ACLPolicyAction":
			sp.ACLPolicyAction = &(o.ACLPolicyAction)
		case "ACLPolicyID":
			sp.ACLPolicyID = &(o.ACLPolicyID)
		case "ID":
			sp.ID = &(o.ID)
		case "RTT":
			sp.RTT = &(o.RTT)
		case "applicationListening":
			sp.ApplicationListening = &(o.ApplicationListening)
		case "claims":
			sp.Claims = &(o.Claims)
		case "claimsType":
			sp.ClaimsType = &(o.ClaimsType)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "enforcerNamespace":
			sp.EnforcerNamespace = &(o.EnforcerNamespace)
		case "enforcerVersion":
			sp.EnforcerVersion = &(o.EnforcerVersion)
		case "error":
			sp.Error = &(o.Error)
		case "excludedNetworks":
			sp.ExcludedNetworks = &(o.ExcludedNetworks)
		case "fourTuple":
			sp.FourTuple = &(o.FourTuple)
		case "isServer":
			sp.IsServer = &(o.IsServer)
		case "iterationIndex":
			sp.IterationIndex = &(o.IterationIndex)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "payloadSize":
			sp.PayloadSize = &(o.PayloadSize)
		case "peerCertExpiry":
			sp.PeerCertExpiry = &(o.PeerCertExpiry)
		case "peerCertIssuer":
			sp.PeerCertIssuer = &(o.PeerCertIssuer)
		case "peerCertSubject":
			sp.PeerCertSubject = &(o.PeerCertSubject)
		case "pingID":
			sp.PingID = &(o.PingID)
		case "policyAction":
			sp.PolicyAction = &(o.PolicyAction)
		case "policyID":
			sp.PolicyID = &(o.PolicyID)
		case "policyNamespace":
			sp.PolicyNamespace = &(o.PolicyNamespace)
		case "processingUnitID":
			sp.ProcessingUnitID = &(o.ProcessingUnitID)
		case "protocol":
			sp.Protocol = &(o.Protocol)
		case "remoteController":
			sp.RemoteController = &(o.RemoteController)
		case "remoteEndpointType":
			sp.RemoteEndpointType = &(o.RemoteEndpointType)
		case "remoteNamespace":
			sp.RemoteNamespace = &(o.RemoteNamespace)
		case "remoteNamespaceType":
			sp.RemoteNamespaceType = &(o.RemoteNamespaceType)
		case "remoteProcessingUnitID":
			sp.RemoteProcessingUnitID = &(o.RemoteProcessingUnitID)
		case "seqNum":
			sp.SeqNum = &(o.SeqNum)
		case "serviceID":
			sp.ServiceID = &(o.ServiceID)
		case "serviceType":
			sp.ServiceType = &(o.ServiceType)
		case "targetTCPNetworks":
			sp.TargetTCPNetworks = &(o.TargetTCPNetworks)
		case "type":
			sp.Type = &(o.Type)
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

// Patch apply the non nil value of a *SparsePingProbe to the object.
func (o *PingProbe) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePingProbe)
	if so.ACLPolicyAction != nil {
		o.ACLPolicyAction = *so.ACLPolicyAction
	}
	if so.ACLPolicyID != nil {
		o.ACLPolicyID = *so.ACLPolicyID
	}
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.RTT != nil {
		o.RTT = *so.RTT
	}
	if so.ApplicationListening != nil {
		o.ApplicationListening = *so.ApplicationListening
	}
	if so.Claims != nil {
		o.Claims = *so.Claims
	}
	if so.ClaimsType != nil {
		o.ClaimsType = *so.ClaimsType
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.EnforcerNamespace != nil {
		o.EnforcerNamespace = *so.EnforcerNamespace
	}
	if so.EnforcerVersion != nil {
		o.EnforcerVersion = *so.EnforcerVersion
	}
	if so.Error != nil {
		o.Error = *so.Error
	}
	if so.ExcludedNetworks != nil {
		o.ExcludedNetworks = *so.ExcludedNetworks
	}
	if so.FourTuple != nil {
		o.FourTuple = *so.FourTuple
	}
	if so.IsServer != nil {
		o.IsServer = *so.IsServer
	}
	if so.IterationIndex != nil {
		o.IterationIndex = *so.IterationIndex
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.PayloadSize != nil {
		o.PayloadSize = *so.PayloadSize
	}
	if so.PeerCertExpiry != nil {
		o.PeerCertExpiry = *so.PeerCertExpiry
	}
	if so.PeerCertIssuer != nil {
		o.PeerCertIssuer = *so.PeerCertIssuer
	}
	if so.PeerCertSubject != nil {
		o.PeerCertSubject = *so.PeerCertSubject
	}
	if so.PingID != nil {
		o.PingID = *so.PingID
	}
	if so.PolicyAction != nil {
		o.PolicyAction = *so.PolicyAction
	}
	if so.PolicyID != nil {
		o.PolicyID = *so.PolicyID
	}
	if so.PolicyNamespace != nil {
		o.PolicyNamespace = *so.PolicyNamespace
	}
	if so.ProcessingUnitID != nil {
		o.ProcessingUnitID = *so.ProcessingUnitID
	}
	if so.Protocol != nil {
		o.Protocol = *so.Protocol
	}
	if so.RemoteController != nil {
		o.RemoteController = *so.RemoteController
	}
	if so.RemoteEndpointType != nil {
		o.RemoteEndpointType = *so.RemoteEndpointType
	}
	if so.RemoteNamespace != nil {
		o.RemoteNamespace = *so.RemoteNamespace
	}
	if so.RemoteNamespaceType != nil {
		o.RemoteNamespaceType = *so.RemoteNamespaceType
	}
	if so.RemoteProcessingUnitID != nil {
		o.RemoteProcessingUnitID = *so.RemoteProcessingUnitID
	}
	if so.SeqNum != nil {
		o.SeqNum = *so.SeqNum
	}
	if so.ServiceID != nil {
		o.ServiceID = *so.ServiceID
	}
	if so.ServiceType != nil {
		o.ServiceType = *so.ServiceType
	}
	if so.TargetTCPNetworks != nil {
		o.TargetTCPNetworks = *so.TargetTCPNetworks
	}
	if so.Type != nil {
		o.Type = *so.Type
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

// DeepCopy returns a deep copy if the PingProbe.
func (o *PingProbe) DeepCopy() *PingProbe {

	if o == nil {
		return nil
	}

	out := &PingProbe{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PingProbe.
func (o *PingProbe) DeepCopyInto(out *PingProbe) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PingProbe: %s", err))
	}

	*out = *target.(*PingProbe)
}

// Validate valides the current information stored into the structure.
func (o *PingProbe) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("claimsType", string(o.ClaimsType), []string{"Transmitted", "Received"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerID", o.EnforcerID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerNamespace", o.EnforcerNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("pingID", o.PingID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("remoteEndpointType", string(o.RemoteEndpointType), []string{"ProcessingUnit", "External"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("remoteNamespaceType", string(o.RemoteNamespaceType), []string{"Plain", "Hash"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Request", "Response"}, false); err != nil {
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
func (*PingProbe) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PingProbeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PingProbeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PingProbe) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PingProbeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PingProbe) ValueForAttribute(name string) interface{} {

	switch name {
	case "ACLPolicyAction":
		return o.ACLPolicyAction
	case "ACLPolicyID":
		return o.ACLPolicyID
	case "ID":
		return o.ID
	case "RTT":
		return o.RTT
	case "applicationListening":
		return o.ApplicationListening
	case "claims":
		return o.Claims
	case "claimsType":
		return o.ClaimsType
	case "createTime":
		return o.CreateTime
	case "enforcerID":
		return o.EnforcerID
	case "enforcerNamespace":
		return o.EnforcerNamespace
	case "enforcerVersion":
		return o.EnforcerVersion
	case "error":
		return o.Error
	case "excludedNetworks":
		return o.ExcludedNetworks
	case "fourTuple":
		return o.FourTuple
	case "isServer":
		return o.IsServer
	case "iterationIndex":
		return o.IterationIndex
	case "migrationsLog":
		return o.MigrationsLog
	case "namespace":
		return o.Namespace
	case "payloadSize":
		return o.PayloadSize
	case "peerCertExpiry":
		return o.PeerCertExpiry
	case "peerCertIssuer":
		return o.PeerCertIssuer
	case "peerCertSubject":
		return o.PeerCertSubject
	case "pingID":
		return o.PingID
	case "policyAction":
		return o.PolicyAction
	case "policyID":
		return o.PolicyID
	case "policyNamespace":
		return o.PolicyNamespace
	case "processingUnitID":
		return o.ProcessingUnitID
	case "protocol":
		return o.Protocol
	case "remoteController":
		return o.RemoteController
	case "remoteEndpointType":
		return o.RemoteEndpointType
	case "remoteNamespace":
		return o.RemoteNamespace
	case "remoteNamespaceType":
		return o.RemoteNamespaceType
	case "remoteProcessingUnitID":
		return o.RemoteProcessingUnitID
	case "seqNum":
		return o.SeqNum
	case "serviceID":
		return o.ServiceID
	case "serviceType":
		return o.ServiceType
	case "targetTCPNetworks":
		return o.TargetTCPNetworks
	case "type":
		return o.Type
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// PingProbeAttributesMap represents the map of attribute for PingProbe.
var PingProbeAttributesMap = map[string]elemental.AttributeSpecification{
	"ACLPolicyAction": {
		AllowedChoices: []string{},
		ConvertedName:  "ACLPolicyAction",
		Description:    `Action of the ACL policy.`,
		Exposed:        true,
		Name:           "ACLPolicyAction",
		Stored:         true,
		Type:           "string",
	},
	"ACLPolicyID": {
		AllowedChoices: []string{},
		ConvertedName:  "ACLPolicyID",
		Description:    `ID of the ACL policy.`,
		Exposed:        true,
		Name:           "ACLPolicyID",
		Stored:         true,
		Type:           "string",
	},
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
	"RTT": {
		AllowedChoices: []string{},
		ConvertedName:  "RTT",
		Description:    `Time taken for a single request-response to complete.`,
		Exposed:        true,
		Name:           "RTT",
		Stored:         true,
		Type:           "string",
	},
	"ApplicationListening": {
		AllowedChoices: []string{},
		ConvertedName:  "ApplicationListening",
		Description:    `If true, application responded to the request.`,
		Exposed:        true,
		Name:           "applicationListening",
		Stored:         true,
		Type:           "boolean",
	},
	"Claims": {
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `Claims of the processing unit.`,
		Exposed:        true,
		Name:           "claims",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"ClaimsType": {
		AllowedChoices: []string{"Transmitted", "Received"},
		ConvertedName:  "ClaimsType",
		Description:    `Type of claims reported.`,
		Exposed:        true,
		Name:           "claimsType",
		Stored:         true,
		SubType:        "string",
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
	"EnforcerID": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the defender.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"EnforcerNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the defender.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"EnforcerVersion": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerVersion",
		Description:    `Semantic version of the defender.`,
		Exposed:        true,
		Name:           "enforcerVersion",
		Stored:         true,
		Type:           "string",
	},
	"Error": {
		AllowedChoices: []string{},
		ConvertedName:  "Error",
		Description:    `A non-empty error indicates a failure.`,
		Exposed:        true,
		Name:           "error",
		Stored:         true,
		Type:           "string",
	},
	"ExcludedNetworks": {
		AllowedChoices: []string{},
		ConvertedName:  "ExcludedNetworks",
		Description:    `If true, destination IP is in ` + "`" + `excludedNetworks` + "`" + `.`,
		Exposed:        true,
		Name:           "excludedNetworks",
		Stored:         true,
		Type:           "boolean",
	},
	"FourTuple": {
		AllowedChoices: []string{},
		ConvertedName:  "FourTuple",
		Description:    `Four tuple in the format <sip:dip:spt:dpt>.`,
		Exposed:        true,
		Name:           "fourTuple",
		Stored:         true,
		Type:           "string",
	},
	"IsServer": {
		AllowedChoices: []string{},
		ConvertedName:  "IsServer",
		Description:    `If true, the report was generated by the server.`,
		Exposed:        true,
		Name:           "isServer",
		Stored:         true,
		Type:           "boolean",
	},
	"IterationIndex": {
		AllowedChoices: []string{},
		ConvertedName:  "IterationIndex",
		Description:    `Holds the iteration number this probe is attached to.`,
		Exposed:        true,
		Name:           "iterationIndex",
		Stored:         true,
		Type:           "integer",
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
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"PayloadSize": {
		AllowedChoices: []string{},
		ConvertedName:  "PayloadSize",
		Description:    `Size of the payload attached to the packet.`,
		Exposed:        true,
		Name:           "payloadSize",
		Stored:         true,
		Type:           "integer",
	},
	"PeerCertExpiry": {
		AllowedChoices: []string{},
		ConvertedName:  "PeerCertExpiry",
		Description:    `Represents the expiry of the peer certificate.`,
		Exposed:        true,
		Name:           "peerCertExpiry",
		Stored:         true,
		Type:           "string",
	},
	"PeerCertIssuer": {
		AllowedChoices: []string{},
		ConvertedName:  "PeerCertIssuer",
		Description:    `Represents the issuer of the peer certificate.`,
		Exposed:        true,
		Name:           "peerCertIssuer",
		Stored:         true,
		Type:           "string",
	},
	"PeerCertSubject": {
		AllowedChoices: []string{},
		ConvertedName:  "PeerCertSubject",
		Description:    `Represents the subject of the peer certificate.`,
		Exposed:        true,
		Name:           "peerCertSubject",
		Stored:         true,
		Type:           "string",
	},
	"PingID": {
		AllowedChoices: []string{},
		ConvertedName:  "PingID",
		Description:    `PingID unique to a single ping control.`,
		Exposed:        true,
		Name:           "pingID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"PolicyAction": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyAction",
		Description:    `Action of the policy.`,
		Exposed:        true,
		Name:           "policyAction",
		Stored:         true,
		Type:           "string",
	},
	"PolicyID": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `ID of the policy.`,
		Exposed:        true,
		Name:           "policyID",
		Stored:         true,
		Type:           "string",
	},
	"PolicyNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyNamespace",
		Description:    `ID of the policy.`,
		Exposed:        true,
		Name:           "policyNamespace",
		Stored:         true,
		Type:           "string",
	},
	"ProcessingUnitID": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the reporting processing unit.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Stored:         true,
		Type:           "string",
	},
	"Protocol": {
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol used for the communication.`,
		Exposed:        true,
		Name:           "protocol",
		Stored:         true,
		Type:           "integer",
	},
	"RemoteController": {
		AllowedChoices: []string{},
		ConvertedName:  "RemoteController",
		Description:    `Controller of the remote endpoint.`,
		Exposed:        true,
		Name:           "remoteController",
		Stored:         true,
		Type:           "string",
	},
	"RemoteEndpointType": {
		AllowedChoices: []string{"ProcessingUnit", "External"},
		ConvertedName:  "RemoteEndpointType",
		Description:    `Represents the remote endpoint type.`,
		Exposed:        true,
		Name:           "remoteEndpointType",
		Stored:         true,
		Type:           "enum",
	},
	"RemoteNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "RemoteNamespace",
		Description:    `Namespace of the remote processing unit.`,
		Exposed:        true,
		Name:           "remoteNamespace",
		Stored:         true,
		Type:           "string",
	},
	"RemoteNamespaceType": {
		AllowedChoices: []string{"Plain", "Hash"},
		ConvertedName:  "RemoteNamespaceType",
		Description:    `Type of the namespace reported.`,
		Exposed:        true,
		Name:           "remoteNamespaceType",
		Stored:         true,
		Type:           "enum",
	},
	"RemoteProcessingUnitID": {
		AllowedChoices: []string{},
		ConvertedName:  "RemoteProcessingUnitID",
		Description:    `ID of the remote processing unit.`,
		Exposed:        true,
		Name:           "remoteProcessingUnitID",
		Stored:         true,
		Type:           "string",
	},
	"SeqNum": {
		AllowedChoices: []string{},
		ConvertedName:  "SeqNum",
		Description:    `Sequence number of the TCP packet. number.`,
		Exposed:        true,
		Name:           "seqNum",
		Stored:         true,
		Type:           "integer",
	},
	"ServiceID": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceID",
		Description:    `ID of the service If the service type is a proxy.`,
		Exposed:        true,
		Name:           "serviceID",
		Stored:         true,
		Type:           "string",
	},
	"ServiceType": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ServiceType",
		Description:    `Type of the service.`,
		Exposed:        true,
		Name:           "serviceType",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"TargetTCPNetworks": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetTCPNetworks",
		Description:    `If true, destination IP is in ` + "`" + `targetTCPNetworks` + "`" + `.`,
		Exposed:        true,
		Name:           "targetTCPNetworks",
		Stored:         true,
		Type:           "boolean",
	},
	"Type": {
		AllowedChoices: []string{"Request", "Response"},
		ConvertedName:  "Type",
		Description:    `Type of the report.`,
		Exposed:        true,
		Name:           "type",
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

// PingProbeLowerCaseAttributesMap represents the map of attribute for PingProbe.
var PingProbeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"aclpolicyaction": {
		AllowedChoices: []string{},
		ConvertedName:  "ACLPolicyAction",
		Description:    `Action of the ACL policy.`,
		Exposed:        true,
		Name:           "ACLPolicyAction",
		Stored:         true,
		Type:           "string",
	},
	"aclpolicyid": {
		AllowedChoices: []string{},
		ConvertedName:  "ACLPolicyID",
		Description:    `ID of the ACL policy.`,
		Exposed:        true,
		Name:           "ACLPolicyID",
		Stored:         true,
		Type:           "string",
	},
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
	"rtt": {
		AllowedChoices: []string{},
		ConvertedName:  "RTT",
		Description:    `Time taken for a single request-response to complete.`,
		Exposed:        true,
		Name:           "RTT",
		Stored:         true,
		Type:           "string",
	},
	"applicationlistening": {
		AllowedChoices: []string{},
		ConvertedName:  "ApplicationListening",
		Description:    `If true, application responded to the request.`,
		Exposed:        true,
		Name:           "applicationListening",
		Stored:         true,
		Type:           "boolean",
	},
	"claims": {
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `Claims of the processing unit.`,
		Exposed:        true,
		Name:           "claims",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"claimstype": {
		AllowedChoices: []string{"Transmitted", "Received"},
		ConvertedName:  "ClaimsType",
		Description:    `Type of claims reported.`,
		Exposed:        true,
		Name:           "claimsType",
		Stored:         true,
		SubType:        "string",
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
	"enforcerid": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the defender.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"enforcernamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the defender.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"enforcerversion": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerVersion",
		Description:    `Semantic version of the defender.`,
		Exposed:        true,
		Name:           "enforcerVersion",
		Stored:         true,
		Type:           "string",
	},
	"error": {
		AllowedChoices: []string{},
		ConvertedName:  "Error",
		Description:    `A non-empty error indicates a failure.`,
		Exposed:        true,
		Name:           "error",
		Stored:         true,
		Type:           "string",
	},
	"excludednetworks": {
		AllowedChoices: []string{},
		ConvertedName:  "ExcludedNetworks",
		Description:    `If true, destination IP is in ` + "`" + `excludedNetworks` + "`" + `.`,
		Exposed:        true,
		Name:           "excludedNetworks",
		Stored:         true,
		Type:           "boolean",
	},
	"fourtuple": {
		AllowedChoices: []string{},
		ConvertedName:  "FourTuple",
		Description:    `Four tuple in the format <sip:dip:spt:dpt>.`,
		Exposed:        true,
		Name:           "fourTuple",
		Stored:         true,
		Type:           "string",
	},
	"isserver": {
		AllowedChoices: []string{},
		ConvertedName:  "IsServer",
		Description:    `If true, the report was generated by the server.`,
		Exposed:        true,
		Name:           "isServer",
		Stored:         true,
		Type:           "boolean",
	},
	"iterationindex": {
		AllowedChoices: []string{},
		ConvertedName:  "IterationIndex",
		Description:    `Holds the iteration number this probe is attached to.`,
		Exposed:        true,
		Name:           "iterationIndex",
		Stored:         true,
		Type:           "integer",
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
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"payloadsize": {
		AllowedChoices: []string{},
		ConvertedName:  "PayloadSize",
		Description:    `Size of the payload attached to the packet.`,
		Exposed:        true,
		Name:           "payloadSize",
		Stored:         true,
		Type:           "integer",
	},
	"peercertexpiry": {
		AllowedChoices: []string{},
		ConvertedName:  "PeerCertExpiry",
		Description:    `Represents the expiry of the peer certificate.`,
		Exposed:        true,
		Name:           "peerCertExpiry",
		Stored:         true,
		Type:           "string",
	},
	"peercertissuer": {
		AllowedChoices: []string{},
		ConvertedName:  "PeerCertIssuer",
		Description:    `Represents the issuer of the peer certificate.`,
		Exposed:        true,
		Name:           "peerCertIssuer",
		Stored:         true,
		Type:           "string",
	},
	"peercertsubject": {
		AllowedChoices: []string{},
		ConvertedName:  "PeerCertSubject",
		Description:    `Represents the subject of the peer certificate.`,
		Exposed:        true,
		Name:           "peerCertSubject",
		Stored:         true,
		Type:           "string",
	},
	"pingid": {
		AllowedChoices: []string{},
		ConvertedName:  "PingID",
		Description:    `PingID unique to a single ping control.`,
		Exposed:        true,
		Name:           "pingID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"policyaction": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyAction",
		Description:    `Action of the policy.`,
		Exposed:        true,
		Name:           "policyAction",
		Stored:         true,
		Type:           "string",
	},
	"policyid": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `ID of the policy.`,
		Exposed:        true,
		Name:           "policyID",
		Stored:         true,
		Type:           "string",
	},
	"policynamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyNamespace",
		Description:    `ID of the policy.`,
		Exposed:        true,
		Name:           "policyNamespace",
		Stored:         true,
		Type:           "string",
	},
	"processingunitid": {
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the reporting processing unit.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Stored:         true,
		Type:           "string",
	},
	"protocol": {
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol used for the communication.`,
		Exposed:        true,
		Name:           "protocol",
		Stored:         true,
		Type:           "integer",
	},
	"remotecontroller": {
		AllowedChoices: []string{},
		ConvertedName:  "RemoteController",
		Description:    `Controller of the remote endpoint.`,
		Exposed:        true,
		Name:           "remoteController",
		Stored:         true,
		Type:           "string",
	},
	"remoteendpointtype": {
		AllowedChoices: []string{"ProcessingUnit", "External"},
		ConvertedName:  "RemoteEndpointType",
		Description:    `Represents the remote endpoint type.`,
		Exposed:        true,
		Name:           "remoteEndpointType",
		Stored:         true,
		Type:           "enum",
	},
	"remotenamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "RemoteNamespace",
		Description:    `Namespace of the remote processing unit.`,
		Exposed:        true,
		Name:           "remoteNamespace",
		Stored:         true,
		Type:           "string",
	},
	"remotenamespacetype": {
		AllowedChoices: []string{"Plain", "Hash"},
		ConvertedName:  "RemoteNamespaceType",
		Description:    `Type of the namespace reported.`,
		Exposed:        true,
		Name:           "remoteNamespaceType",
		Stored:         true,
		Type:           "enum",
	},
	"remoteprocessingunitid": {
		AllowedChoices: []string{},
		ConvertedName:  "RemoteProcessingUnitID",
		Description:    `ID of the remote processing unit.`,
		Exposed:        true,
		Name:           "remoteProcessingUnitID",
		Stored:         true,
		Type:           "string",
	},
	"seqnum": {
		AllowedChoices: []string{},
		ConvertedName:  "SeqNum",
		Description:    `Sequence number of the TCP packet. number.`,
		Exposed:        true,
		Name:           "seqNum",
		Stored:         true,
		Type:           "integer",
	},
	"serviceid": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceID",
		Description:    `ID of the service If the service type is a proxy.`,
		Exposed:        true,
		Name:           "serviceID",
		Stored:         true,
		Type:           "string",
	},
	"servicetype": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ServiceType",
		Description:    `Type of the service.`,
		Exposed:        true,
		Name:           "serviceType",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"targettcpnetworks": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetTCPNetworks",
		Description:    `If true, destination IP is in ` + "`" + `targetTCPNetworks` + "`" + `.`,
		Exposed:        true,
		Name:           "targetTCPNetworks",
		Stored:         true,
		Type:           "boolean",
	},
	"type": {
		AllowedChoices: []string{"Request", "Response"},
		ConvertedName:  "Type",
		Description:    `Type of the report.`,
		Exposed:        true,
		Name:           "type",
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

// SparsePingProbesList represents a list of SparsePingProbes
type SparsePingProbesList []*SparsePingProbe

// Identity returns the identity of the objects in the list.
func (o SparsePingProbesList) Identity() elemental.Identity {

	return PingProbeIdentity
}

// Copy returns a pointer to a copy the SparsePingProbesList.
func (o SparsePingProbesList) Copy() elemental.Identifiables {

	copy := append(SparsePingProbesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePingProbesList.
func (o SparsePingProbesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePingProbesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePingProbe))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePingProbesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePingProbesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePingProbesList converted to PingProbesList.
func (o SparsePingProbesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePingProbesList) Version() int {

	return 1
}

// SparsePingProbe represents the sparse version of a pingprobe.
type SparsePingProbe struct {
	// Action of the ACL policy.
	ACLPolicyAction *string `json:"ACLPolicyAction,omitempty" msgpack:"ACLPolicyAction,omitempty" bson:"aclpolicyaction,omitempty" mapstructure:"ACLPolicyAction,omitempty"`

	// ID of the ACL policy.
	ACLPolicyID *string `json:"ACLPolicyID,omitempty" msgpack:"ACLPolicyID,omitempty" bson:"aclpolicyid,omitempty" mapstructure:"ACLPolicyID,omitempty"`

	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Time taken for a single request-response to complete.
	RTT *string `json:"RTT,omitempty" msgpack:"RTT,omitempty" bson:"rtt,omitempty" mapstructure:"RTT,omitempty"`

	// If true, application responded to the request.
	ApplicationListening *bool `json:"applicationListening,omitempty" msgpack:"applicationListening,omitempty" bson:"applicationlistening,omitempty" mapstructure:"applicationListening,omitempty"`

	// Claims of the processing unit.
	Claims *[]string `json:"claims,omitempty" msgpack:"claims,omitempty" bson:"claims,omitempty" mapstructure:"claims,omitempty"`

	// Type of claims reported.
	ClaimsType *PingProbeClaimsTypeValue `json:"claimsType,omitempty" msgpack:"claimsType,omitempty" bson:"claimstype,omitempty" mapstructure:"claimsType,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// ID of the defender.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"enforcerid,omitempty" mapstructure:"enforcerID,omitempty"`

	// Namespace of the defender.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"enforcernamespace,omitempty" mapstructure:"enforcerNamespace,omitempty"`

	// Semantic version of the defender.
	EnforcerVersion *string `json:"enforcerVersion,omitempty" msgpack:"enforcerVersion,omitempty" bson:"enforcerversion,omitempty" mapstructure:"enforcerVersion,omitempty"`

	// A non-empty error indicates a failure.
	Error *string `json:"error,omitempty" msgpack:"error,omitempty" bson:"error,omitempty" mapstructure:"error,omitempty"`

	// If true, destination IP is in `excludedNetworks`.
	ExcludedNetworks *bool `json:"excludedNetworks,omitempty" msgpack:"excludedNetworks,omitempty" bson:"excludednetworks,omitempty" mapstructure:"excludedNetworks,omitempty"`

	// Four tuple in the format <sip:dip:spt:dpt>.
	FourTuple *string `json:"fourTuple,omitempty" msgpack:"fourTuple,omitempty" bson:"fourtuple,omitempty" mapstructure:"fourTuple,omitempty"`

	// If true, the report was generated by the server.
	IsServer *bool `json:"isServer,omitempty" msgpack:"isServer,omitempty" bson:"isserver,omitempty" mapstructure:"isServer,omitempty"`

	// Holds the iteration number this probe is attached to.
	IterationIndex *int `json:"iterationIndex,omitempty" msgpack:"iterationIndex,omitempty" bson:"iterationindex,omitempty" mapstructure:"iterationIndex,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Size of the payload attached to the packet.
	PayloadSize *int `json:"payloadSize,omitempty" msgpack:"payloadSize,omitempty" bson:"payloadsize,omitempty" mapstructure:"payloadSize,omitempty"`

	// Represents the expiry of the peer certificate.
	PeerCertExpiry *string `json:"peerCertExpiry,omitempty" msgpack:"peerCertExpiry,omitempty" bson:"peercertexpiry,omitempty" mapstructure:"peerCertExpiry,omitempty"`

	// Represents the issuer of the peer certificate.
	PeerCertIssuer *string `json:"peerCertIssuer,omitempty" msgpack:"peerCertIssuer,omitempty" bson:"peercertissuer,omitempty" mapstructure:"peerCertIssuer,omitempty"`

	// Represents the subject of the peer certificate.
	PeerCertSubject *string `json:"peerCertSubject,omitempty" msgpack:"peerCertSubject,omitempty" bson:"peercertsubject,omitempty" mapstructure:"peerCertSubject,omitempty"`

	// PingID unique to a single ping control.
	PingID *string `json:"pingID,omitempty" msgpack:"pingID,omitempty" bson:"pingid,omitempty" mapstructure:"pingID,omitempty"`

	// Action of the policy.
	PolicyAction *string `json:"policyAction,omitempty" msgpack:"policyAction,omitempty" bson:"policyaction,omitempty" mapstructure:"policyAction,omitempty"`

	// ID of the policy.
	PolicyID *string `json:"policyID,omitempty" msgpack:"policyID,omitempty" bson:"policyid,omitempty" mapstructure:"policyID,omitempty"`

	// ID of the policy.
	PolicyNamespace *string `json:"policyNamespace,omitempty" msgpack:"policyNamespace,omitempty" bson:"policynamespace,omitempty" mapstructure:"policyNamespace,omitempty"`

	// ID of the reporting processing unit.
	ProcessingUnitID *string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"processingunitid,omitempty" mapstructure:"processingUnitID,omitempty"`

	// Protocol used for the communication.
	Protocol *int `json:"protocol,omitempty" msgpack:"protocol,omitempty" bson:"protocol,omitempty" mapstructure:"protocol,omitempty"`

	// Controller of the remote endpoint.
	RemoteController *string `json:"remoteController,omitempty" msgpack:"remoteController,omitempty" bson:"remotecontroller,omitempty" mapstructure:"remoteController,omitempty"`

	// Represents the remote endpoint type.
	RemoteEndpointType *PingProbeRemoteEndpointTypeValue `json:"remoteEndpointType,omitempty" msgpack:"remoteEndpointType,omitempty" bson:"remoteendpointtype,omitempty" mapstructure:"remoteEndpointType,omitempty"`

	// Namespace of the remote processing unit.
	RemoteNamespace *string `json:"remoteNamespace,omitempty" msgpack:"remoteNamespace,omitempty" bson:"remotenamespace,omitempty" mapstructure:"remoteNamespace,omitempty"`

	// Type of the namespace reported.
	RemoteNamespaceType *PingProbeRemoteNamespaceTypeValue `json:"remoteNamespaceType,omitempty" msgpack:"remoteNamespaceType,omitempty" bson:"remotenamespacetype,omitempty" mapstructure:"remoteNamespaceType,omitempty"`

	// ID of the remote processing unit.
	RemoteProcessingUnitID *string `json:"remoteProcessingUnitID,omitempty" msgpack:"remoteProcessingUnitID,omitempty" bson:"remoteprocessingunitid,omitempty" mapstructure:"remoteProcessingUnitID,omitempty"`

	// Sequence number of the TCP packet. number.
	SeqNum *int `json:"seqNum,omitempty" msgpack:"seqNum,omitempty" bson:"seqnum,omitempty" mapstructure:"seqNum,omitempty"`

	// ID of the service If the service type is a proxy.
	ServiceID *string `json:"serviceID,omitempty" msgpack:"serviceID,omitempty" bson:"serviceid,omitempty" mapstructure:"serviceID,omitempty"`

	// Type of the service.
	ServiceType *string `json:"serviceType,omitempty" msgpack:"serviceType,omitempty" bson:"servicetype,omitempty" mapstructure:"serviceType,omitempty"`

	// If true, destination IP is in `targetTCPNetworks`.
	TargetTCPNetworks *bool `json:"targetTCPNetworks,omitempty" msgpack:"targetTCPNetworks,omitempty" bson:"targettcpnetworks,omitempty" mapstructure:"targetTCPNetworks,omitempty"`

	// Type of the report.
	Type *PingProbeTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"type,omitempty" mapstructure:"type,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePingProbe returns a new  SparsePingProbe.
func NewSparsePingProbe() *SparsePingProbe {
	return &SparsePingProbe{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePingProbe) Identity() elemental.Identity {

	return PingProbeIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePingProbe) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePingProbe) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePingProbe) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePingProbe{}

	if o.ACLPolicyAction != nil {
		s.ACLPolicyAction = o.ACLPolicyAction
	}
	if o.ACLPolicyID != nil {
		s.ACLPolicyID = o.ACLPolicyID
	}
	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.RTT != nil {
		s.RTT = o.RTT
	}
	if o.ApplicationListening != nil {
		s.ApplicationListening = o.ApplicationListening
	}
	if o.Claims != nil {
		s.Claims = o.Claims
	}
	if o.ClaimsType != nil {
		s.ClaimsType = o.ClaimsType
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.EnforcerID != nil {
		s.EnforcerID = o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		s.EnforcerNamespace = o.EnforcerNamespace
	}
	if o.EnforcerVersion != nil {
		s.EnforcerVersion = o.EnforcerVersion
	}
	if o.Error != nil {
		s.Error = o.Error
	}
	if o.ExcludedNetworks != nil {
		s.ExcludedNetworks = o.ExcludedNetworks
	}
	if o.FourTuple != nil {
		s.FourTuple = o.FourTuple
	}
	if o.IsServer != nil {
		s.IsServer = o.IsServer
	}
	if o.IterationIndex != nil {
		s.IterationIndex = o.IterationIndex
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.PayloadSize != nil {
		s.PayloadSize = o.PayloadSize
	}
	if o.PeerCertExpiry != nil {
		s.PeerCertExpiry = o.PeerCertExpiry
	}
	if o.PeerCertIssuer != nil {
		s.PeerCertIssuer = o.PeerCertIssuer
	}
	if o.PeerCertSubject != nil {
		s.PeerCertSubject = o.PeerCertSubject
	}
	if o.PingID != nil {
		s.PingID = o.PingID
	}
	if o.PolicyAction != nil {
		s.PolicyAction = o.PolicyAction
	}
	if o.PolicyID != nil {
		s.PolicyID = o.PolicyID
	}
	if o.PolicyNamespace != nil {
		s.PolicyNamespace = o.PolicyNamespace
	}
	if o.ProcessingUnitID != nil {
		s.ProcessingUnitID = o.ProcessingUnitID
	}
	if o.Protocol != nil {
		s.Protocol = o.Protocol
	}
	if o.RemoteController != nil {
		s.RemoteController = o.RemoteController
	}
	if o.RemoteEndpointType != nil {
		s.RemoteEndpointType = o.RemoteEndpointType
	}
	if o.RemoteNamespace != nil {
		s.RemoteNamespace = o.RemoteNamespace
	}
	if o.RemoteNamespaceType != nil {
		s.RemoteNamespaceType = o.RemoteNamespaceType
	}
	if o.RemoteProcessingUnitID != nil {
		s.RemoteProcessingUnitID = o.RemoteProcessingUnitID
	}
	if o.SeqNum != nil {
		s.SeqNum = o.SeqNum
	}
	if o.ServiceID != nil {
		s.ServiceID = o.ServiceID
	}
	if o.ServiceType != nil {
		s.ServiceType = o.ServiceType
	}
	if o.TargetTCPNetworks != nil {
		s.TargetTCPNetworks = o.TargetTCPNetworks
	}
	if o.Type != nil {
		s.Type = o.Type
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
func (o *SparsePingProbe) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePingProbe{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.ACLPolicyAction != nil {
		o.ACLPolicyAction = s.ACLPolicyAction
	}
	if s.ACLPolicyID != nil {
		o.ACLPolicyID = s.ACLPolicyID
	}
	id := s.ID.Hex()
	o.ID = &id
	if s.RTT != nil {
		o.RTT = s.RTT
	}
	if s.ApplicationListening != nil {
		o.ApplicationListening = s.ApplicationListening
	}
	if s.Claims != nil {
		o.Claims = s.Claims
	}
	if s.ClaimsType != nil {
		o.ClaimsType = s.ClaimsType
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.EnforcerID != nil {
		o.EnforcerID = s.EnforcerID
	}
	if s.EnforcerNamespace != nil {
		o.EnforcerNamespace = s.EnforcerNamespace
	}
	if s.EnforcerVersion != nil {
		o.EnforcerVersion = s.EnforcerVersion
	}
	if s.Error != nil {
		o.Error = s.Error
	}
	if s.ExcludedNetworks != nil {
		o.ExcludedNetworks = s.ExcludedNetworks
	}
	if s.FourTuple != nil {
		o.FourTuple = s.FourTuple
	}
	if s.IsServer != nil {
		o.IsServer = s.IsServer
	}
	if s.IterationIndex != nil {
		o.IterationIndex = s.IterationIndex
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.PayloadSize != nil {
		o.PayloadSize = s.PayloadSize
	}
	if s.PeerCertExpiry != nil {
		o.PeerCertExpiry = s.PeerCertExpiry
	}
	if s.PeerCertIssuer != nil {
		o.PeerCertIssuer = s.PeerCertIssuer
	}
	if s.PeerCertSubject != nil {
		o.PeerCertSubject = s.PeerCertSubject
	}
	if s.PingID != nil {
		o.PingID = s.PingID
	}
	if s.PolicyAction != nil {
		o.PolicyAction = s.PolicyAction
	}
	if s.PolicyID != nil {
		o.PolicyID = s.PolicyID
	}
	if s.PolicyNamespace != nil {
		o.PolicyNamespace = s.PolicyNamespace
	}
	if s.ProcessingUnitID != nil {
		o.ProcessingUnitID = s.ProcessingUnitID
	}
	if s.Protocol != nil {
		o.Protocol = s.Protocol
	}
	if s.RemoteController != nil {
		o.RemoteController = s.RemoteController
	}
	if s.RemoteEndpointType != nil {
		o.RemoteEndpointType = s.RemoteEndpointType
	}
	if s.RemoteNamespace != nil {
		o.RemoteNamespace = s.RemoteNamespace
	}
	if s.RemoteNamespaceType != nil {
		o.RemoteNamespaceType = s.RemoteNamespaceType
	}
	if s.RemoteProcessingUnitID != nil {
		o.RemoteProcessingUnitID = s.RemoteProcessingUnitID
	}
	if s.SeqNum != nil {
		o.SeqNum = s.SeqNum
	}
	if s.ServiceID != nil {
		o.ServiceID = s.ServiceID
	}
	if s.ServiceType != nil {
		o.ServiceType = s.ServiceType
	}
	if s.TargetTCPNetworks != nil {
		o.TargetTCPNetworks = s.TargetTCPNetworks
	}
	if s.Type != nil {
		o.Type = s.Type
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
func (o *SparsePingProbe) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePingProbe) ToPlain() elemental.PlainIdentifiable {

	out := NewPingProbe()
	if o.ACLPolicyAction != nil {
		out.ACLPolicyAction = *o.ACLPolicyAction
	}
	if o.ACLPolicyID != nil {
		out.ACLPolicyID = *o.ACLPolicyID
	}
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.RTT != nil {
		out.RTT = *o.RTT
	}
	if o.ApplicationListening != nil {
		out.ApplicationListening = *o.ApplicationListening
	}
	if o.Claims != nil {
		out.Claims = *o.Claims
	}
	if o.ClaimsType != nil {
		out.ClaimsType = *o.ClaimsType
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		out.EnforcerNamespace = *o.EnforcerNamespace
	}
	if o.EnforcerVersion != nil {
		out.EnforcerVersion = *o.EnforcerVersion
	}
	if o.Error != nil {
		out.Error = *o.Error
	}
	if o.ExcludedNetworks != nil {
		out.ExcludedNetworks = *o.ExcludedNetworks
	}
	if o.FourTuple != nil {
		out.FourTuple = *o.FourTuple
	}
	if o.IsServer != nil {
		out.IsServer = *o.IsServer
	}
	if o.IterationIndex != nil {
		out.IterationIndex = *o.IterationIndex
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.PayloadSize != nil {
		out.PayloadSize = *o.PayloadSize
	}
	if o.PeerCertExpiry != nil {
		out.PeerCertExpiry = *o.PeerCertExpiry
	}
	if o.PeerCertIssuer != nil {
		out.PeerCertIssuer = *o.PeerCertIssuer
	}
	if o.PeerCertSubject != nil {
		out.PeerCertSubject = *o.PeerCertSubject
	}
	if o.PingID != nil {
		out.PingID = *o.PingID
	}
	if o.PolicyAction != nil {
		out.PolicyAction = *o.PolicyAction
	}
	if o.PolicyID != nil {
		out.PolicyID = *o.PolicyID
	}
	if o.PolicyNamespace != nil {
		out.PolicyNamespace = *o.PolicyNamespace
	}
	if o.ProcessingUnitID != nil {
		out.ProcessingUnitID = *o.ProcessingUnitID
	}
	if o.Protocol != nil {
		out.Protocol = *o.Protocol
	}
	if o.RemoteController != nil {
		out.RemoteController = *o.RemoteController
	}
	if o.RemoteEndpointType != nil {
		out.RemoteEndpointType = *o.RemoteEndpointType
	}
	if o.RemoteNamespace != nil {
		out.RemoteNamespace = *o.RemoteNamespace
	}
	if o.RemoteNamespaceType != nil {
		out.RemoteNamespaceType = *o.RemoteNamespaceType
	}
	if o.RemoteProcessingUnitID != nil {
		out.RemoteProcessingUnitID = *o.RemoteProcessingUnitID
	}
	if o.SeqNum != nil {
		out.SeqNum = *o.SeqNum
	}
	if o.ServiceID != nil {
		out.ServiceID = *o.ServiceID
	}
	if o.ServiceType != nil {
		out.ServiceType = *o.ServiceType
	}
	if o.TargetTCPNetworks != nil {
		out.TargetTCPNetworks = *o.TargetTCPNetworks
	}
	if o.Type != nil {
		out.Type = *o.Type
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
func (o *SparsePingProbe) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparsePingProbe) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparsePingProbe) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparsePingProbe) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparsePingProbe) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparsePingProbe) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparsePingProbe) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparsePingProbe) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparsePingProbe) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparsePingProbe) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparsePingProbe) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparsePingProbe) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparsePingProbe.
func (o *SparsePingProbe) DeepCopy() *SparsePingProbe {

	if o == nil {
		return nil
	}

	out := &SparsePingProbe{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePingProbe.
func (o *SparsePingProbe) DeepCopyInto(out *SparsePingProbe) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePingProbe: %s", err))
	}

	*out = *target.(*SparsePingProbe)
}

type mongoAttributesPingProbe struct {
	ACLPolicyAction        string                            `bson:"aclpolicyaction"`
	ACLPolicyID            string                            `bson:"aclpolicyid"`
	ID                     bson.ObjectId                     `bson:"_id,omitempty"`
	RTT                    string                            `bson:"rtt"`
	ApplicationListening   bool                              `bson:"applicationlistening"`
	Claims                 []string                          `bson:"claims"`
	ClaimsType             PingProbeClaimsTypeValue          `bson:"claimstype"`
	CreateTime             time.Time                         `bson:"createtime"`
	EnforcerID             string                            `bson:"enforcerid"`
	EnforcerNamespace      string                            `bson:"enforcernamespace"`
	EnforcerVersion        string                            `bson:"enforcerversion"`
	Error                  string                            `bson:"error"`
	ExcludedNetworks       bool                              `bson:"excludednetworks"`
	FourTuple              string                            `bson:"fourtuple"`
	IsServer               bool                              `bson:"isserver"`
	IterationIndex         int                               `bson:"iterationindex"`
	MigrationsLog          map[string]string                 `bson:"migrationslog,omitempty"`
	Namespace              string                            `bson:"namespace"`
	PayloadSize            int                               `bson:"payloadsize"`
	PeerCertExpiry         string                            `bson:"peercertexpiry"`
	PeerCertIssuer         string                            `bson:"peercertissuer"`
	PeerCertSubject        string                            `bson:"peercertsubject"`
	PingID                 string                            `bson:"pingid"`
	PolicyAction           string                            `bson:"policyaction"`
	PolicyID               string                            `bson:"policyid"`
	PolicyNamespace        string                            `bson:"policynamespace"`
	ProcessingUnitID       string                            `bson:"processingunitid"`
	Protocol               int                               `bson:"protocol"`
	RemoteController       string                            `bson:"remotecontroller"`
	RemoteEndpointType     PingProbeRemoteEndpointTypeValue  `bson:"remoteendpointtype"`
	RemoteNamespace        string                            `bson:"remotenamespace"`
	RemoteNamespaceType    PingProbeRemoteNamespaceTypeValue `bson:"remotenamespacetype"`
	RemoteProcessingUnitID string                            `bson:"remoteprocessingunitid"`
	SeqNum                 int                               `bson:"seqnum"`
	ServiceID              string                            `bson:"serviceid"`
	ServiceType            string                            `bson:"servicetype"`
	TargetTCPNetworks      bool                              `bson:"targettcpnetworks"`
	Type                   PingProbeTypeValue                `bson:"type"`
	UpdateTime             time.Time                         `bson:"updatetime"`
	ZHash                  int                               `bson:"zhash"`
	Zone                   int                               `bson:"zone"`
}
type mongoAttributesSparsePingProbe struct {
	ACLPolicyAction        *string                            `bson:"aclpolicyaction,omitempty"`
	ACLPolicyID            *string                            `bson:"aclpolicyid,omitempty"`
	ID                     bson.ObjectId                      `bson:"_id,omitempty"`
	RTT                    *string                            `bson:"rtt,omitempty"`
	ApplicationListening   *bool                              `bson:"applicationlistening,omitempty"`
	Claims                 *[]string                          `bson:"claims,omitempty"`
	ClaimsType             *PingProbeClaimsTypeValue          `bson:"claimstype,omitempty"`
	CreateTime             *time.Time                         `bson:"createtime,omitempty"`
	EnforcerID             *string                            `bson:"enforcerid,omitempty"`
	EnforcerNamespace      *string                            `bson:"enforcernamespace,omitempty"`
	EnforcerVersion        *string                            `bson:"enforcerversion,omitempty"`
	Error                  *string                            `bson:"error,omitempty"`
	ExcludedNetworks       *bool                              `bson:"excludednetworks,omitempty"`
	FourTuple              *string                            `bson:"fourtuple,omitempty"`
	IsServer               *bool                              `bson:"isserver,omitempty"`
	IterationIndex         *int                               `bson:"iterationindex,omitempty"`
	MigrationsLog          *map[string]string                 `bson:"migrationslog,omitempty"`
	Namespace              *string                            `bson:"namespace,omitempty"`
	PayloadSize            *int                               `bson:"payloadsize,omitempty"`
	PeerCertExpiry         *string                            `bson:"peercertexpiry,omitempty"`
	PeerCertIssuer         *string                            `bson:"peercertissuer,omitempty"`
	PeerCertSubject        *string                            `bson:"peercertsubject,omitempty"`
	PingID                 *string                            `bson:"pingid,omitempty"`
	PolicyAction           *string                            `bson:"policyaction,omitempty"`
	PolicyID               *string                            `bson:"policyid,omitempty"`
	PolicyNamespace        *string                            `bson:"policynamespace,omitempty"`
	ProcessingUnitID       *string                            `bson:"processingunitid,omitempty"`
	Protocol               *int                               `bson:"protocol,omitempty"`
	RemoteController       *string                            `bson:"remotecontroller,omitempty"`
	RemoteEndpointType     *PingProbeRemoteEndpointTypeValue  `bson:"remoteendpointtype,omitempty"`
	RemoteNamespace        *string                            `bson:"remotenamespace,omitempty"`
	RemoteNamespaceType    *PingProbeRemoteNamespaceTypeValue `bson:"remotenamespacetype,omitempty"`
	RemoteProcessingUnitID *string                            `bson:"remoteprocessingunitid,omitempty"`
	SeqNum                 *int                               `bson:"seqnum,omitempty"`
	ServiceID              *string                            `bson:"serviceid,omitempty"`
	ServiceType            *string                            `bson:"servicetype,omitempty"`
	TargetTCPNetworks      *bool                              `bson:"targettcpnetworks,omitempty"`
	Type                   *PingProbeTypeValue                `bson:"type,omitempty"`
	UpdateTime             *time.Time                         `bson:"updatetime,omitempty"`
	ZHash                  *int                               `bson:"zhash,omitempty"`
	Zone                   *int                               `bson:"zone,omitempty"`
}
