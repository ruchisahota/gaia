package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CounterReportIdentity represents the Identity of the object.
var CounterReportIdentity = elemental.Identity{
	Name:     "counterreport",
	Category: "counterreports",
	Package:  "zack",
	Private:  false,
}

// CounterReportsList represents a list of CounterReports
type CounterReportsList []*CounterReport

// Identity returns the identity of the objects in the list.
func (o CounterReportsList) Identity() elemental.Identity {

	return CounterReportIdentity
}

// Copy returns a pointer to a copy the CounterReportsList.
func (o CounterReportsList) Copy() elemental.Identifiables {

	copy := append(CounterReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CounterReportsList.
func (o CounterReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CounterReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CounterReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CounterReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CounterReportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CounterReportsList converted to SparseCounterReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CounterReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCounterReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCounterReport)
	}

	return out
}

// Version returns the version of the content.
func (o CounterReportsList) Version() int {

	return 1
}

// CounterReport represents the model of a counterreport
type CounterReport struct {
	// Counter for sending finack ack received in uknown connection state.
	CounterAckInUnknownState int `json:"counterAckInUnknownState" msgpack:"counterAckInUnknownState" bson:"-" mapstructure:"counterAckInUnknownState,omitempty"`

	// Counter for ack packet dropped because of invalid format.
	CounterAckInvalidFormat int `json:"counterAckInvalidFormat" msgpack:"counterAckInvalidFormat" bson:"-" mapstructure:"counterAckInvalidFormat,omitempty"`

	// Counter for reject ack packet as per policy.
	CounterAckRejected int `json:"counterAckRejected" msgpack:"counterAckRejected" bson:"-" mapstructure:"counterAckRejected,omitempty"`

	// Counter for ack packet dropped because signature validation failed.
	CounterAckSigValidationFailed int `json:"counterAckSigValidationFailed" msgpack:"counterAckSigValidationFailed" bson:"-" mapstructure:"counterAckSigValidationFailed,omitempty"`

	// Counter for tcp authentication option not found.
	CounterAckTCPNoTCPAuthOption int `json:"counterAckTCPNoTCPAuthOption" msgpack:"counterAckTCPNoTCPAuthOption" bson:"-" mapstructure:"counterAckTCPNoTCPAuthOption,omitempty"`

	// Counter for ".
	CounterConnectionsProcessed int `json:"counterConnectionsProcessed" msgpack:"counterConnectionsProcessed" bson:"-" mapstructure:"counterConnectionsProcessed,omitempty"`

	// Counter for unable to find contextid.
	CounterContextIDNotFound int `json:"counterContextIDNotFound" msgpack:"counterContextIDNotFound" bson:"-" mapstructure:"counterContextIDNotFound,omitempty"`

	// Counter for no acls found for external services. dropping application syn
	// packet.
	CounterDroppedExternalService int `json:"counterDroppedExternalService" msgpack:"counterDroppedExternalService" bson:"-" mapstructure:"counterDroppedExternalService,omitempty"`

	// Counter for invalid connection state.
	CounterInvalidConnState int `json:"counterInvalidConnState" msgpack:"counterInvalidConnState" bson:"-" mapstructure:"counterInvalidConnState,omitempty"`

	// Counter for invalid net state.
	CounterInvalidNetState int `json:"counterInvalidNetState" msgpack:"counterInvalidNetState" bson:"-" mapstructure:"counterInvalidNetState,omitempty"`

	// Counter for invalid protocol.
	CounterInvalidProtocol int `json:"counterInvalidProtocol" msgpack:"counterInvalidProtocol" bson:"-" mapstructure:"counterInvalidProtocol,omitempty"`

	// Counter for pu is already dead - drop synack packet.
	CounterInvalidSynAck int `json:"counterInvalidSynAck" msgpack:"counterInvalidSynAck" bson:"-" mapstructure:"counterInvalidSynAck,omitempty"`

	// Counter for pu mark not found.
	CounterMarkNotFound int `json:"counterMarkNotFound" msgpack:"counterMarkNotFound" bson:"-" mapstructure:"counterMarkNotFound,omitempty"`

	// Counter for network syn packet was not seen.
	CounterNetSynNotSeen int `json:"counterNetSynNotSeen" msgpack:"counterNetSynNotSeen" bson:"-" mapstructure:"counterNetSynNotSeen,omitempty"`

	// Counter for no context or connection found.
	CounterNoConnFound int `json:"counterNoConnFound" msgpack:"counterNoConnFound" bson:"-" mapstructure:"counterNoConnFound,omitempty"`

	// Counter for traffic that belongs to a non PU process.
	CounterNonPUTraffic int `json:"counterNonPUTraffic" msgpack:"counterNonPUTraffic" bson:"-" mapstructure:"counterNonPUTraffic,omitempty"`

	// Counter for synack for flow with processed finack.
	CounterOutOfOrderSynAck int `json:"counterOutOfOrderSynAck" msgpack:"counterOutOfOrderSynAck" bson:"-" mapstructure:"counterOutOfOrderSynAck,omitempty"`

	// Counter for port not found.
	CounterPortNotFound int `json:"counterPortNotFound" msgpack:"counterPortNotFound" bson:"-" mapstructure:"counterPortNotFound,omitempty"`

	// Counter for reject the packet as per policy.
	CounterRejectPacket int `json:"counterRejectPacket" msgpack:"counterRejectPacket" bson:"-" mapstructure:"counterRejectPacket,omitempty"`

	// Counter for post service processing failed for network packet.
	CounterServicePostprocessorFailed int `json:"counterServicePostprocessorFailed" msgpack:"counterServicePostprocessorFailed" bson:"-" mapstructure:"counterServicePostprocessorFailed,omitempty"`

	// Counter for pre service processing failed for network packet.
	CounterServicePreprocessorFailed int `json:"counterServicePreprocessorFailed" msgpack:"counterServicePreprocessorFailed" bson:"-" mapstructure:"counterServicePreprocessorFailed,omitempty"`

	// Counter for synack packet dropped because of bad claims.
	CounterSynAckBadClaims int `json:"counterSynAckBadClaims" msgpack:"counterSynAckBadClaims" bson:"-" mapstructure:"counterSynAckBadClaims,omitempty"`

	// Counter for syn/ack packet dropped because of encryption mismatch.
	CounterSynAckClaimsMisMatch int `json:"counterSynAckClaimsMisMatch" msgpack:"counterSynAckClaimsMisMatch" bson:"-" mapstructure:"counterSynAckClaimsMisMatch,omitempty"`

	// Counter for synack from external service dropped.
	CounterSynAckDroppedExternalService int `json:"counterSynAckDroppedExternalService" msgpack:"counterSynAckDroppedExternalService" bson:"-" mapstructure:"counterSynAckDroppedExternalService,omitempty"`

	// Counter for synack packet dropped because of invalid format.
	CounterSynAckInvalidFormat int `json:"counterSynAckInvalidFormat" msgpack:"counterSynAckInvalidFormat" bson:"-" mapstructure:"counterSynAckInvalidFormat,omitempty"`

	// Counter for synack packet dropped because of no claims.
	CounterSynAckMissingClaims int `json:"counterSynAckMissingClaims" msgpack:"counterSynAckMissingClaims" bson:"-" mapstructure:"counterSynAckMissingClaims,omitempty"`

	// Counter for synack packet dropped because of missing token.
	CounterSynAckMissingToken int `json:"counterSynAckMissingToken" msgpack:"counterSynAckMissingToken" bson:"-" mapstructure:"counterSynAckMissingToken,omitempty"`

	// Counter for tcp authentication option not found.
	CounterSynAckNoTCPAuthOption int `json:"counterSynAckNoTCPAuthOption" msgpack:"counterSynAckNoTCPAuthOption" bson:"-" mapstructure:"counterSynAckNoTCPAuthOption,omitempty"`

	// Counter for dropping because of reject rule on transmitter.
	CounterSynAckRejected int `json:"counterSynAckRejected" msgpack:"counterSynAckRejected" bson:"-" mapstructure:"counterSynAckRejected,omitempty"`

	// Counter for syn packet dropped because of invalid format.
	CounterSynDroppedInvalidFormat int `json:"counterSynDroppedInvalidFormat" msgpack:"counterSynDroppedInvalidFormat" bson:"-" mapstructure:"counterSynDroppedInvalidFormat,omitempty"`

	// Counter for syn packet dropped because of invalid token.
	CounterSynDroppedInvalidToken int `json:"counterSynDroppedInvalidToken" msgpack:"counterSynDroppedInvalidToken" bson:"-" mapstructure:"counterSynDroppedInvalidToken,omitempty"`

	// Counter for syn packet dropped because of no claims.
	CounterSynDroppedNoClaims int `json:"counterSynDroppedNoClaims" msgpack:"counterSynDroppedNoClaims" bson:"-" mapstructure:"counterSynDroppedNoClaims,omitempty"`

	// Counter for tcp authentication option not found.
	CounterSynDroppedTCPOption int `json:"counterSynDroppedTCPOption" msgpack:"counterSynDroppedTCPOption" bson:"-" mapstructure:"counterSynDroppedTCPOption,omitempty"`

	// Counter for syn dropped due to policy.
	CounterSynRejectPacket int `json:"counterSynRejectPacket" msgpack:"counterSynRejectPacket" bson:"-" mapstructure:"counterSynRejectPacket,omitempty"`

	// Counter for received syn packet from unknown pu.
	CounterSynUnexpectedPacket int `json:"counterSynUnexpectedPacket" msgpack:"counterSynUnexpectedPacket" bson:"-" mapstructure:"counterSynUnexpectedPacket,omitempty"`

	// Counter for tcp authentication option not found.
	CounterTCPAuthNotFound int `json:"counterTCPAuthNotFound" msgpack:"counterTCPAuthNotFound" bson:"-" mapstructure:"counterTCPAuthNotFound,omitempty"`

	// Counter for unknown error.
	CounterUnknownError int `json:"counterUnknownError" msgpack:"counterUnknownError" bson:"-" mapstructure:"counterUnknownError,omitempty"`

	// Identifier of the enforcer sending the report.
	EnforcerID string `json:"enforcerID" msgpack:"enforcerID" bson:"enforcerid" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer sending the report.
	EnforcerNamespace string `json:"enforcerNamespace" msgpack:"enforcerNamespace" bson:"enforcernamespace" mapstructure:"enforcerNamespace,omitempty"`

	// PUID is the ID of the PU reporting the counter.
	ProcessingUnitID string `json:"processingUnitID" msgpack:"processingUnitID" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the PU reporting the counter.
	ProcessingUnitNamespace string `json:"processingUnitNamespace" msgpack:"processingUnitNamespace" bson:"-" mapstructure:"processingUnitNamespace,omitempty"`

	// Timestamp is the date of the report.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCounterReport returns a new *CounterReport
func NewCounterReport() *CounterReport {

	return &CounterReport{
		ModelVersion:                        1,
		CounterSynAckBadClaims:              0,
		CounterConnectionsProcessed:         0,
		CounterSynAckClaimsMisMatch:         0,
		CounterAckRejected:                  0,
		CounterSynAckDroppedExternalService: 0,
		CounterAckTCPNoTCPAuthOption:        0,
		CounterSynAckInvalidFormat:          0,
		CounterAckInvalidFormat:             0,
		CounterSynAckMissingClaims:          0,
		CounterInvalidConnState:             0,
		CounterSynAckMissingToken:           0,
		CounterInvalidProtocol:              0,
		CounterSynDroppedInvalidToken:       0,
		CounterSynUnexpectedPacket:          0,
		CounterAckSigValidationFailed:       0,
		CounterSynAckRejected:               0,
		EnforcerNamespace:                   "/my/namespace",
		CounterServicePreprocessorFailed:    0,
		CounterSynDroppedTCPOption:          0,
		CounterOutOfOrderSynAck:             0,
		CounterUnknownError:                 0,
		CounterDroppedExternalService:       0,
		CounterSynAckNoTCPAuthOption:        0,
		CounterServicePostprocessorFailed:   0,
		CounterSynDroppedInvalidFormat:      0,
		CounterNonPUTraffic:                 0,
		CounterSynDroppedNoClaims:           0,
		CounterContextIDNotFound:            0,
		CounterSynRejectPacket:              0,
		CounterNetSynNotSeen:                0,
		CounterTCPAuthNotFound:              0,
		CounterAckInUnknownState:            0,
		EnforcerID:                          "xxxx-xxx-xxxx",
		CounterPortNotFound:                 0,
		CounterInvalidNetState:              0,
		CounterMarkNotFound:                 0,
		CounterNoConnFound:                  0,
		CounterRejectPacket:                 0,
		CounterInvalidSynAck:                0,
	}
}

// Identity returns the Identity of the object.
func (o *CounterReport) Identity() elemental.Identity {

	return CounterReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CounterReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CounterReport) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *CounterReport) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *CounterReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CounterReport) Doc() string {

	return `Post a new counter tracing report.`
}

func (o *CounterReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CounterReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCounterReport{
			CounterAckInUnknownState:            &o.CounterAckInUnknownState,
			CounterAckInvalidFormat:             &o.CounterAckInvalidFormat,
			CounterAckRejected:                  &o.CounterAckRejected,
			CounterAckSigValidationFailed:       &o.CounterAckSigValidationFailed,
			CounterAckTCPNoTCPAuthOption:        &o.CounterAckTCPNoTCPAuthOption,
			CounterConnectionsProcessed:         &o.CounterConnectionsProcessed,
			CounterContextIDNotFound:            &o.CounterContextIDNotFound,
			CounterDroppedExternalService:       &o.CounterDroppedExternalService,
			CounterInvalidConnState:             &o.CounterInvalidConnState,
			CounterInvalidNetState:              &o.CounterInvalidNetState,
			CounterInvalidProtocol:              &o.CounterInvalidProtocol,
			CounterInvalidSynAck:                &o.CounterInvalidSynAck,
			CounterMarkNotFound:                 &o.CounterMarkNotFound,
			CounterNetSynNotSeen:                &o.CounterNetSynNotSeen,
			CounterNoConnFound:                  &o.CounterNoConnFound,
			CounterNonPUTraffic:                 &o.CounterNonPUTraffic,
			CounterOutOfOrderSynAck:             &o.CounterOutOfOrderSynAck,
			CounterPortNotFound:                 &o.CounterPortNotFound,
			CounterRejectPacket:                 &o.CounterRejectPacket,
			CounterServicePostprocessorFailed:   &o.CounterServicePostprocessorFailed,
			CounterServicePreprocessorFailed:    &o.CounterServicePreprocessorFailed,
			CounterSynAckBadClaims:              &o.CounterSynAckBadClaims,
			CounterSynAckClaimsMisMatch:         &o.CounterSynAckClaimsMisMatch,
			CounterSynAckDroppedExternalService: &o.CounterSynAckDroppedExternalService,
			CounterSynAckInvalidFormat:          &o.CounterSynAckInvalidFormat,
			CounterSynAckMissingClaims:          &o.CounterSynAckMissingClaims,
			CounterSynAckMissingToken:           &o.CounterSynAckMissingToken,
			CounterSynAckNoTCPAuthOption:        &o.CounterSynAckNoTCPAuthOption,
			CounterSynAckRejected:               &o.CounterSynAckRejected,
			CounterSynDroppedInvalidFormat:      &o.CounterSynDroppedInvalidFormat,
			CounterSynDroppedInvalidToken:       &o.CounterSynDroppedInvalidToken,
			CounterSynDroppedNoClaims:           &o.CounterSynDroppedNoClaims,
			CounterSynDroppedTCPOption:          &o.CounterSynDroppedTCPOption,
			CounterSynRejectPacket:              &o.CounterSynRejectPacket,
			CounterSynUnexpectedPacket:          &o.CounterSynUnexpectedPacket,
			CounterTCPAuthNotFound:              &o.CounterTCPAuthNotFound,
			CounterUnknownError:                 &o.CounterUnknownError,
			EnforcerID:                          &o.EnforcerID,
			EnforcerNamespace:                   &o.EnforcerNamespace,
			ProcessingUnitID:                    &o.ProcessingUnitID,
			ProcessingUnitNamespace:             &o.ProcessingUnitNamespace,
			Timestamp:                           &o.Timestamp,
		}
	}

	sp := &SparseCounterReport{}
	for _, f := range fields {
		switch f {
		case "counterAckInUnknownState":
			sp.CounterAckInUnknownState = &(o.CounterAckInUnknownState)
		case "counterAckInvalidFormat":
			sp.CounterAckInvalidFormat = &(o.CounterAckInvalidFormat)
		case "counterAckRejected":
			sp.CounterAckRejected = &(o.CounterAckRejected)
		case "counterAckSigValidationFailed":
			sp.CounterAckSigValidationFailed = &(o.CounterAckSigValidationFailed)
		case "counterAckTCPNoTCPAuthOption":
			sp.CounterAckTCPNoTCPAuthOption = &(o.CounterAckTCPNoTCPAuthOption)
		case "counterConnectionsProcessed":
			sp.CounterConnectionsProcessed = &(o.CounterConnectionsProcessed)
		case "counterContextIDNotFound":
			sp.CounterContextIDNotFound = &(o.CounterContextIDNotFound)
		case "counterDroppedExternalService":
			sp.CounterDroppedExternalService = &(o.CounterDroppedExternalService)
		case "counterInvalidConnState":
			sp.CounterInvalidConnState = &(o.CounterInvalidConnState)
		case "counterInvalidNetState":
			sp.CounterInvalidNetState = &(o.CounterInvalidNetState)
		case "counterInvalidProtocol":
			sp.CounterInvalidProtocol = &(o.CounterInvalidProtocol)
		case "counterInvalidSynAck":
			sp.CounterInvalidSynAck = &(o.CounterInvalidSynAck)
		case "counterMarkNotFound":
			sp.CounterMarkNotFound = &(o.CounterMarkNotFound)
		case "counterNetSynNotSeen":
			sp.CounterNetSynNotSeen = &(o.CounterNetSynNotSeen)
		case "counterNoConnFound":
			sp.CounterNoConnFound = &(o.CounterNoConnFound)
		case "counterNonPUTraffic":
			sp.CounterNonPUTraffic = &(o.CounterNonPUTraffic)
		case "counterOutOfOrderSynAck":
			sp.CounterOutOfOrderSynAck = &(o.CounterOutOfOrderSynAck)
		case "counterPortNotFound":
			sp.CounterPortNotFound = &(o.CounterPortNotFound)
		case "counterRejectPacket":
			sp.CounterRejectPacket = &(o.CounterRejectPacket)
		case "counterServicePostprocessorFailed":
			sp.CounterServicePostprocessorFailed = &(o.CounterServicePostprocessorFailed)
		case "counterServicePreprocessorFailed":
			sp.CounterServicePreprocessorFailed = &(o.CounterServicePreprocessorFailed)
		case "counterSynAckBadClaims":
			sp.CounterSynAckBadClaims = &(o.CounterSynAckBadClaims)
		case "counterSynAckClaimsMisMatch":
			sp.CounterSynAckClaimsMisMatch = &(o.CounterSynAckClaimsMisMatch)
		case "counterSynAckDroppedExternalService":
			sp.CounterSynAckDroppedExternalService = &(o.CounterSynAckDroppedExternalService)
		case "counterSynAckInvalidFormat":
			sp.CounterSynAckInvalidFormat = &(o.CounterSynAckInvalidFormat)
		case "counterSynAckMissingClaims":
			sp.CounterSynAckMissingClaims = &(o.CounterSynAckMissingClaims)
		case "counterSynAckMissingToken":
			sp.CounterSynAckMissingToken = &(o.CounterSynAckMissingToken)
		case "counterSynAckNoTCPAuthOption":
			sp.CounterSynAckNoTCPAuthOption = &(o.CounterSynAckNoTCPAuthOption)
		case "counterSynAckRejected":
			sp.CounterSynAckRejected = &(o.CounterSynAckRejected)
		case "counterSynDroppedInvalidFormat":
			sp.CounterSynDroppedInvalidFormat = &(o.CounterSynDroppedInvalidFormat)
		case "counterSynDroppedInvalidToken":
			sp.CounterSynDroppedInvalidToken = &(o.CounterSynDroppedInvalidToken)
		case "counterSynDroppedNoClaims":
			sp.CounterSynDroppedNoClaims = &(o.CounterSynDroppedNoClaims)
		case "counterSynDroppedTCPOption":
			sp.CounterSynDroppedTCPOption = &(o.CounterSynDroppedTCPOption)
		case "counterSynRejectPacket":
			sp.CounterSynRejectPacket = &(o.CounterSynRejectPacket)
		case "counterSynUnexpectedPacket":
			sp.CounterSynUnexpectedPacket = &(o.CounterSynUnexpectedPacket)
		case "counterTCPAuthNotFound":
			sp.CounterTCPAuthNotFound = &(o.CounterTCPAuthNotFound)
		case "counterUnknownError":
			sp.CounterUnknownError = &(o.CounterUnknownError)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "enforcerNamespace":
			sp.EnforcerNamespace = &(o.EnforcerNamespace)
		case "processingUnitID":
			sp.ProcessingUnitID = &(o.ProcessingUnitID)
		case "processingUnitNamespace":
			sp.ProcessingUnitNamespace = &(o.ProcessingUnitNamespace)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCounterReport to the object.
func (o *CounterReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCounterReport)
	if so.CounterAckInUnknownState != nil {
		o.CounterAckInUnknownState = *so.CounterAckInUnknownState
	}
	if so.CounterAckInvalidFormat != nil {
		o.CounterAckInvalidFormat = *so.CounterAckInvalidFormat
	}
	if so.CounterAckRejected != nil {
		o.CounterAckRejected = *so.CounterAckRejected
	}
	if so.CounterAckSigValidationFailed != nil {
		o.CounterAckSigValidationFailed = *so.CounterAckSigValidationFailed
	}
	if so.CounterAckTCPNoTCPAuthOption != nil {
		o.CounterAckTCPNoTCPAuthOption = *so.CounterAckTCPNoTCPAuthOption
	}
	if so.CounterConnectionsProcessed != nil {
		o.CounterConnectionsProcessed = *so.CounterConnectionsProcessed
	}
	if so.CounterContextIDNotFound != nil {
		o.CounterContextIDNotFound = *so.CounterContextIDNotFound
	}
	if so.CounterDroppedExternalService != nil {
		o.CounterDroppedExternalService = *so.CounterDroppedExternalService
	}
	if so.CounterInvalidConnState != nil {
		o.CounterInvalidConnState = *so.CounterInvalidConnState
	}
	if so.CounterInvalidNetState != nil {
		o.CounterInvalidNetState = *so.CounterInvalidNetState
	}
	if so.CounterInvalidProtocol != nil {
		o.CounterInvalidProtocol = *so.CounterInvalidProtocol
	}
	if so.CounterInvalidSynAck != nil {
		o.CounterInvalidSynAck = *so.CounterInvalidSynAck
	}
	if so.CounterMarkNotFound != nil {
		o.CounterMarkNotFound = *so.CounterMarkNotFound
	}
	if so.CounterNetSynNotSeen != nil {
		o.CounterNetSynNotSeen = *so.CounterNetSynNotSeen
	}
	if so.CounterNoConnFound != nil {
		o.CounterNoConnFound = *so.CounterNoConnFound
	}
	if so.CounterNonPUTraffic != nil {
		o.CounterNonPUTraffic = *so.CounterNonPUTraffic
	}
	if so.CounterOutOfOrderSynAck != nil {
		o.CounterOutOfOrderSynAck = *so.CounterOutOfOrderSynAck
	}
	if so.CounterPortNotFound != nil {
		o.CounterPortNotFound = *so.CounterPortNotFound
	}
	if so.CounterRejectPacket != nil {
		o.CounterRejectPacket = *so.CounterRejectPacket
	}
	if so.CounterServicePostprocessorFailed != nil {
		o.CounterServicePostprocessorFailed = *so.CounterServicePostprocessorFailed
	}
	if so.CounterServicePreprocessorFailed != nil {
		o.CounterServicePreprocessorFailed = *so.CounterServicePreprocessorFailed
	}
	if so.CounterSynAckBadClaims != nil {
		o.CounterSynAckBadClaims = *so.CounterSynAckBadClaims
	}
	if so.CounterSynAckClaimsMisMatch != nil {
		o.CounterSynAckClaimsMisMatch = *so.CounterSynAckClaimsMisMatch
	}
	if so.CounterSynAckDroppedExternalService != nil {
		o.CounterSynAckDroppedExternalService = *so.CounterSynAckDroppedExternalService
	}
	if so.CounterSynAckInvalidFormat != nil {
		o.CounterSynAckInvalidFormat = *so.CounterSynAckInvalidFormat
	}
	if so.CounterSynAckMissingClaims != nil {
		o.CounterSynAckMissingClaims = *so.CounterSynAckMissingClaims
	}
	if so.CounterSynAckMissingToken != nil {
		o.CounterSynAckMissingToken = *so.CounterSynAckMissingToken
	}
	if so.CounterSynAckNoTCPAuthOption != nil {
		o.CounterSynAckNoTCPAuthOption = *so.CounterSynAckNoTCPAuthOption
	}
	if so.CounterSynAckRejected != nil {
		o.CounterSynAckRejected = *so.CounterSynAckRejected
	}
	if so.CounterSynDroppedInvalidFormat != nil {
		o.CounterSynDroppedInvalidFormat = *so.CounterSynDroppedInvalidFormat
	}
	if so.CounterSynDroppedInvalidToken != nil {
		o.CounterSynDroppedInvalidToken = *so.CounterSynDroppedInvalidToken
	}
	if so.CounterSynDroppedNoClaims != nil {
		o.CounterSynDroppedNoClaims = *so.CounterSynDroppedNoClaims
	}
	if so.CounterSynDroppedTCPOption != nil {
		o.CounterSynDroppedTCPOption = *so.CounterSynDroppedTCPOption
	}
	if so.CounterSynRejectPacket != nil {
		o.CounterSynRejectPacket = *so.CounterSynRejectPacket
	}
	if so.CounterSynUnexpectedPacket != nil {
		o.CounterSynUnexpectedPacket = *so.CounterSynUnexpectedPacket
	}
	if so.CounterTCPAuthNotFound != nil {
		o.CounterTCPAuthNotFound = *so.CounterTCPAuthNotFound
	}
	if so.CounterUnknownError != nil {
		o.CounterUnknownError = *so.CounterUnknownError
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.EnforcerNamespace != nil {
		o.EnforcerNamespace = *so.EnforcerNamespace
	}
	if so.ProcessingUnitID != nil {
		o.ProcessingUnitID = *so.ProcessingUnitID
	}
	if so.ProcessingUnitNamespace != nil {
		o.ProcessingUnitNamespace = *so.ProcessingUnitNamespace
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
}

// DeepCopy returns a deep copy if the CounterReport.
func (o *CounterReport) DeepCopy() *CounterReport {

	if o == nil {
		return nil
	}

	out := &CounterReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CounterReport.
func (o *CounterReport) DeepCopyInto(out *CounterReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CounterReport: %s", err))
	}

	*out = *target.(*CounterReport)
}

// Validate valides the current information stored into the structure.
func (o *CounterReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredInt("counterAckInUnknownState", o.CounterAckInUnknownState); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterAckInvalidFormat", o.CounterAckInvalidFormat); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterAckRejected", o.CounterAckRejected); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterAckSigValidationFailed", o.CounterAckSigValidationFailed); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterAckTCPNoTCPAuthOption", o.CounterAckTCPNoTCPAuthOption); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterConnectionsProcessed", o.CounterConnectionsProcessed); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterContextIDNotFound", o.CounterContextIDNotFound); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterDroppedExternalService", o.CounterDroppedExternalService); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterInvalidConnState", o.CounterInvalidConnState); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterInvalidNetState", o.CounterInvalidNetState); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterInvalidProtocol", o.CounterInvalidProtocol); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterInvalidSynAck", o.CounterInvalidSynAck); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterMarkNotFound", o.CounterMarkNotFound); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterNetSynNotSeen", o.CounterNetSynNotSeen); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterNoConnFound", o.CounterNoConnFound); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterNonPUTraffic", o.CounterNonPUTraffic); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterOutOfOrderSynAck", o.CounterOutOfOrderSynAck); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterPortNotFound", o.CounterPortNotFound); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterRejectPacket", o.CounterRejectPacket); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterServicePostprocessorFailed", o.CounterServicePostprocessorFailed); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterServicePreprocessorFailed", o.CounterServicePreprocessorFailed); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterSynAckBadClaims", o.CounterSynAckBadClaims); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterSynAckClaimsMisMatch", o.CounterSynAckClaimsMisMatch); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterSynAckDroppedExternalService", o.CounterSynAckDroppedExternalService); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterSynAckInvalidFormat", o.CounterSynAckInvalidFormat); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterSynAckMissingClaims", o.CounterSynAckMissingClaims); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterSynAckMissingToken", o.CounterSynAckMissingToken); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterSynAckNoTCPAuthOption", o.CounterSynAckNoTCPAuthOption); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterSynAckRejected", o.CounterSynAckRejected); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterSynDroppedInvalidFormat", o.CounterSynDroppedInvalidFormat); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterSynDroppedInvalidToken", o.CounterSynDroppedInvalidToken); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterSynDroppedNoClaims", o.CounterSynDroppedNoClaims); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterSynDroppedTCPOption", o.CounterSynDroppedTCPOption); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterSynRejectPacket", o.CounterSynRejectPacket); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterSynUnexpectedPacket", o.CounterSynUnexpectedPacket); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterTCPAuthNotFound", o.CounterTCPAuthNotFound); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("counterUnknownError", o.CounterUnknownError); err != nil {
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
func (*CounterReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CounterReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CounterReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CounterReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CounterReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CounterReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "counterAckInUnknownState":
		return o.CounterAckInUnknownState
	case "counterAckInvalidFormat":
		return o.CounterAckInvalidFormat
	case "counterAckRejected":
		return o.CounterAckRejected
	case "counterAckSigValidationFailed":
		return o.CounterAckSigValidationFailed
	case "counterAckTCPNoTCPAuthOption":
		return o.CounterAckTCPNoTCPAuthOption
	case "counterConnectionsProcessed":
		return o.CounterConnectionsProcessed
	case "counterContextIDNotFound":
		return o.CounterContextIDNotFound
	case "counterDroppedExternalService":
		return o.CounterDroppedExternalService
	case "counterInvalidConnState":
		return o.CounterInvalidConnState
	case "counterInvalidNetState":
		return o.CounterInvalidNetState
	case "counterInvalidProtocol":
		return o.CounterInvalidProtocol
	case "counterInvalidSynAck":
		return o.CounterInvalidSynAck
	case "counterMarkNotFound":
		return o.CounterMarkNotFound
	case "counterNetSynNotSeen":
		return o.CounterNetSynNotSeen
	case "counterNoConnFound":
		return o.CounterNoConnFound
	case "counterNonPUTraffic":
		return o.CounterNonPUTraffic
	case "counterOutOfOrderSynAck":
		return o.CounterOutOfOrderSynAck
	case "counterPortNotFound":
		return o.CounterPortNotFound
	case "counterRejectPacket":
		return o.CounterRejectPacket
	case "counterServicePostprocessorFailed":
		return o.CounterServicePostprocessorFailed
	case "counterServicePreprocessorFailed":
		return o.CounterServicePreprocessorFailed
	case "counterSynAckBadClaims":
		return o.CounterSynAckBadClaims
	case "counterSynAckClaimsMisMatch":
		return o.CounterSynAckClaimsMisMatch
	case "counterSynAckDroppedExternalService":
		return o.CounterSynAckDroppedExternalService
	case "counterSynAckInvalidFormat":
		return o.CounterSynAckInvalidFormat
	case "counterSynAckMissingClaims":
		return o.CounterSynAckMissingClaims
	case "counterSynAckMissingToken":
		return o.CounterSynAckMissingToken
	case "counterSynAckNoTCPAuthOption":
		return o.CounterSynAckNoTCPAuthOption
	case "counterSynAckRejected":
		return o.CounterSynAckRejected
	case "counterSynDroppedInvalidFormat":
		return o.CounterSynDroppedInvalidFormat
	case "counterSynDroppedInvalidToken":
		return o.CounterSynDroppedInvalidToken
	case "counterSynDroppedNoClaims":
		return o.CounterSynDroppedNoClaims
	case "counterSynDroppedTCPOption":
		return o.CounterSynDroppedTCPOption
	case "counterSynRejectPacket":
		return o.CounterSynRejectPacket
	case "counterSynUnexpectedPacket":
		return o.CounterSynUnexpectedPacket
	case "counterTCPAuthNotFound":
		return o.CounterTCPAuthNotFound
	case "counterUnknownError":
		return o.CounterUnknownError
	case "enforcerID":
		return o.EnforcerID
	case "enforcerNamespace":
		return o.EnforcerNamespace
	case "processingUnitID":
		return o.ProcessingUnitID
	case "processingUnitNamespace":
		return o.ProcessingUnitNamespace
	case "timestamp":
		return o.Timestamp
	}

	return nil
}

// CounterReportAttributesMap represents the map of attribute for CounterReport.
var CounterReportAttributesMap = map[string]elemental.AttributeSpecification{
	"CounterAckInUnknownState": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterAckInUnknownState",
		Description:    `Counter for sending finack ack received in uknown connection state.`,
		Exposed:        true,
		Name:           "counterAckInUnknownState",
		Required:       true,
		Type:           "integer",
	},
	"CounterAckInvalidFormat": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterAckInvalidFormat",
		Description:    `Counter for ack packet dropped because of invalid format.`,
		Exposed:        true,
		Name:           "counterAckInvalidFormat",
		Required:       true,
		Type:           "integer",
	},
	"CounterAckRejected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterAckRejected",
		Description:    `Counter for reject ack packet as per policy.`,
		Exposed:        true,
		Name:           "counterAckRejected",
		Required:       true,
		Type:           "integer",
	},
	"CounterAckSigValidationFailed": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterAckSigValidationFailed",
		Description:    `Counter for ack packet dropped because signature validation failed.`,
		Exposed:        true,
		Name:           "counterAckSigValidationFailed",
		Required:       true,
		Type:           "integer",
	},
	"CounterAckTCPNoTCPAuthOption": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterAckTCPNoTCPAuthOption",
		Description:    `Counter for tcp authentication option not found.`,
		Exposed:        true,
		Name:           "counterAckTCPNoTCPAuthOption",
		Required:       true,
		Type:           "integer",
	},
	"CounterConnectionsProcessed": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterConnectionsProcessed",
		Description:    `Counter for ".`,
		Exposed:        true,
		Name:           "counterConnectionsProcessed",
		Required:       true,
		Type:           "integer",
	},
	"CounterContextIDNotFound": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterContextIDNotFound",
		Description:    `Counter for unable to find contextid.`,
		Exposed:        true,
		Name:           "counterContextIDNotFound",
		Required:       true,
		Type:           "integer",
	},
	"CounterDroppedExternalService": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterDroppedExternalService",
		Description: `Counter for no acls found for external services. dropping application syn
packet.`,
		Exposed:  true,
		Name:     "counterDroppedExternalService",
		Required: true,
		Type:     "integer",
	},
	"CounterInvalidConnState": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterInvalidConnState",
		Description:    `Counter for invalid connection state.`,
		Exposed:        true,
		Name:           "counterInvalidConnState",
		Required:       true,
		Type:           "integer",
	},
	"CounterInvalidNetState": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterInvalidNetState",
		Description:    `Counter for invalid net state.`,
		Exposed:        true,
		Name:           "counterInvalidNetState",
		Required:       true,
		Type:           "integer",
	},
	"CounterInvalidProtocol": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterInvalidProtocol",
		Description:    `Counter for invalid protocol.`,
		Exposed:        true,
		Name:           "counterInvalidProtocol",
		Required:       true,
		Type:           "integer",
	},
	"CounterInvalidSynAck": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterInvalidSynAck",
		Description:    `Counter for pu is already dead - drop synack packet.`,
		Exposed:        true,
		Name:           "counterInvalidSynAck",
		Required:       true,
		Type:           "integer",
	},
	"CounterMarkNotFound": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterMarkNotFound",
		Description:    `Counter for pu mark not found.`,
		Exposed:        true,
		Name:           "counterMarkNotFound",
		Required:       true,
		Type:           "integer",
	},
	"CounterNetSynNotSeen": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterNetSynNotSeen",
		Description:    `Counter for network syn packet was not seen.`,
		Exposed:        true,
		Name:           "counterNetSynNotSeen",
		Required:       true,
		Type:           "integer",
	},
	"CounterNoConnFound": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterNoConnFound",
		Description:    `Counter for no context or connection found.`,
		Exposed:        true,
		Name:           "counterNoConnFound",
		Required:       true,
		Type:           "integer",
	},
	"CounterNonPUTraffic": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterNonPUTraffic",
		Description:    `Counter for traffic that belongs to a non PU process.`,
		Exposed:        true,
		Name:           "counterNonPUTraffic",
		Required:       true,
		Type:           "integer",
	},
	"CounterOutOfOrderSynAck": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterOutOfOrderSynAck",
		Description:    `Counter for synack for flow with processed finack.`,
		Exposed:        true,
		Name:           "counterOutOfOrderSynAck",
		Required:       true,
		Type:           "integer",
	},
	"CounterPortNotFound": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterPortNotFound",
		Description:    `Counter for port not found.`,
		Exposed:        true,
		Name:           "counterPortNotFound",
		Required:       true,
		Type:           "integer",
	},
	"CounterRejectPacket": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterRejectPacket",
		Description:    `Counter for reject the packet as per policy.`,
		Exposed:        true,
		Name:           "counterRejectPacket",
		Required:       true,
		Type:           "integer",
	},
	"CounterServicePostprocessorFailed": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterServicePostprocessorFailed",
		Description:    `Counter for post service processing failed for network packet.`,
		Exposed:        true,
		Name:           "counterServicePostprocessorFailed",
		Required:       true,
		Type:           "integer",
	},
	"CounterServicePreprocessorFailed": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterServicePreprocessorFailed",
		Description:    `Counter for pre service processing failed for network packet.`,
		Exposed:        true,
		Name:           "counterServicePreprocessorFailed",
		Required:       true,
		Type:           "integer",
	},
	"CounterSynAckBadClaims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynAckBadClaims",
		Description:    `Counter for synack packet dropped because of bad claims.`,
		Exposed:        true,
		Name:           "counterSynAckBadClaims",
		Required:       true,
		Type:           "integer",
	},
	"CounterSynAckClaimsMisMatch": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynAckClaimsMisMatch",
		Description:    `Counter for syn/ack packet dropped because of encryption mismatch.`,
		Exposed:        true,
		Name:           "counterSynAckClaimsMisMatch",
		Required:       true,
		Type:           "integer",
	},
	"CounterSynAckDroppedExternalService": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynAckDroppedExternalService",
		Description:    `Counter for synack from external service dropped.`,
		Exposed:        true,
		Name:           "counterSynAckDroppedExternalService",
		Required:       true,
		Type:           "integer",
	},
	"CounterSynAckInvalidFormat": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynAckInvalidFormat",
		Description:    `Counter for synack packet dropped because of invalid format.`,
		Exposed:        true,
		Name:           "counterSynAckInvalidFormat",
		Required:       true,
		Type:           "integer",
	},
	"CounterSynAckMissingClaims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynAckMissingClaims",
		Description:    `Counter for synack packet dropped because of no claims.`,
		Exposed:        true,
		Name:           "counterSynAckMissingClaims",
		Required:       true,
		Type:           "integer",
	},
	"CounterSynAckMissingToken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynAckMissingToken",
		Description:    `Counter for synack packet dropped because of missing token.`,
		Exposed:        true,
		Name:           "counterSynAckMissingToken",
		Required:       true,
		Type:           "integer",
	},
	"CounterSynAckNoTCPAuthOption": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynAckNoTCPAuthOption",
		Description:    `Counter for tcp authentication option not found.`,
		Exposed:        true,
		Name:           "counterSynAckNoTCPAuthOption",
		Required:       true,
		Type:           "integer",
	},
	"CounterSynAckRejected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynAckRejected",
		Description:    `Counter for dropping because of reject rule on transmitter.`,
		Exposed:        true,
		Name:           "counterSynAckRejected",
		Required:       true,
		Type:           "integer",
	},
	"CounterSynDroppedInvalidFormat": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynDroppedInvalidFormat",
		Description:    `Counter for syn packet dropped because of invalid format.`,
		Exposed:        true,
		Name:           "counterSynDroppedInvalidFormat",
		Required:       true,
		Type:           "integer",
	},
	"CounterSynDroppedInvalidToken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynDroppedInvalidToken",
		Description:    `Counter for syn packet dropped because of invalid token.`,
		Exposed:        true,
		Name:           "counterSynDroppedInvalidToken",
		Required:       true,
		Type:           "integer",
	},
	"CounterSynDroppedNoClaims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynDroppedNoClaims",
		Description:    `Counter for syn packet dropped because of no claims.`,
		Exposed:        true,
		Name:           "counterSynDroppedNoClaims",
		Required:       true,
		Type:           "integer",
	},
	"CounterSynDroppedTCPOption": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynDroppedTCPOption",
		Description:    `Counter for tcp authentication option not found.`,
		Exposed:        true,
		Name:           "counterSynDroppedTCPOption",
		Required:       true,
		Type:           "integer",
	},
	"CounterSynRejectPacket": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynRejectPacket",
		Description:    `Counter for syn dropped due to policy.`,
		Exposed:        true,
		Name:           "counterSynRejectPacket",
		Required:       true,
		Type:           "integer",
	},
	"CounterSynUnexpectedPacket": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynUnexpectedPacket",
		Description:    `Counter for received syn packet from unknown pu.`,
		Exposed:        true,
		Name:           "counterSynUnexpectedPacket",
		Required:       true,
		Type:           "integer",
	},
	"CounterTCPAuthNotFound": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterTCPAuthNotFound",
		Description:    `Counter for tcp authentication option not found.`,
		Exposed:        true,
		Name:           "counterTCPAuthNotFound",
		Required:       true,
		Type:           "integer",
	},
	"CounterUnknownError": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterUnknownError",
		Description:    `Counter for unknown error.`,
		Exposed:        true,
		Name:           "counterUnknownError",
		Required:       true,
		Type:           "integer",
	},
	"EnforcerID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		DefaultValue:   "xxxx-xxx-xxxx",
		Description:    `Identifier of the enforcer sending the report.`,
		Exposed:        true,
		Name:           "enforcerID",
		Stored:         true,
		Type:           "string",
	},
	"EnforcerNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		DefaultValue:   "/my/namespace",
		Description:    `Namespace of the enforcer sending the report.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Stored:         true,
		Type:           "string",
	},
	"ProcessingUnitID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `PUID is the ID of the PU reporting the counter.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "processingUnitID",
		Type:           "string",
	},
	"ProcessingUnitNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the PU reporting the counter.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "processingUnitNamespace",
		Type:           "string",
	},
	"Timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Timestamp is the date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
}

// CounterReportLowerCaseAttributesMap represents the map of attribute for CounterReport.
var CounterReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"counterackinunknownstate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterAckInUnknownState",
		Description:    `Counter for sending finack ack received in uknown connection state.`,
		Exposed:        true,
		Name:           "counterAckInUnknownState",
		Required:       true,
		Type:           "integer",
	},
	"counterackinvalidformat": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterAckInvalidFormat",
		Description:    `Counter for ack packet dropped because of invalid format.`,
		Exposed:        true,
		Name:           "counterAckInvalidFormat",
		Required:       true,
		Type:           "integer",
	},
	"counterackrejected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterAckRejected",
		Description:    `Counter for reject ack packet as per policy.`,
		Exposed:        true,
		Name:           "counterAckRejected",
		Required:       true,
		Type:           "integer",
	},
	"counteracksigvalidationfailed": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterAckSigValidationFailed",
		Description:    `Counter for ack packet dropped because signature validation failed.`,
		Exposed:        true,
		Name:           "counterAckSigValidationFailed",
		Required:       true,
		Type:           "integer",
	},
	"counteracktcpnotcpauthoption": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterAckTCPNoTCPAuthOption",
		Description:    `Counter for tcp authentication option not found.`,
		Exposed:        true,
		Name:           "counterAckTCPNoTCPAuthOption",
		Required:       true,
		Type:           "integer",
	},
	"counterconnectionsprocessed": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterConnectionsProcessed",
		Description:    `Counter for ".`,
		Exposed:        true,
		Name:           "counterConnectionsProcessed",
		Required:       true,
		Type:           "integer",
	},
	"countercontextidnotfound": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterContextIDNotFound",
		Description:    `Counter for unable to find contextid.`,
		Exposed:        true,
		Name:           "counterContextIDNotFound",
		Required:       true,
		Type:           "integer",
	},
	"counterdroppedexternalservice": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterDroppedExternalService",
		Description: `Counter for no acls found for external services. dropping application syn
packet.`,
		Exposed:  true,
		Name:     "counterDroppedExternalService",
		Required: true,
		Type:     "integer",
	},
	"counterinvalidconnstate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterInvalidConnState",
		Description:    `Counter for invalid connection state.`,
		Exposed:        true,
		Name:           "counterInvalidConnState",
		Required:       true,
		Type:           "integer",
	},
	"counterinvalidnetstate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterInvalidNetState",
		Description:    `Counter for invalid net state.`,
		Exposed:        true,
		Name:           "counterInvalidNetState",
		Required:       true,
		Type:           "integer",
	},
	"counterinvalidprotocol": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterInvalidProtocol",
		Description:    `Counter for invalid protocol.`,
		Exposed:        true,
		Name:           "counterInvalidProtocol",
		Required:       true,
		Type:           "integer",
	},
	"counterinvalidsynack": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterInvalidSynAck",
		Description:    `Counter for pu is already dead - drop synack packet.`,
		Exposed:        true,
		Name:           "counterInvalidSynAck",
		Required:       true,
		Type:           "integer",
	},
	"countermarknotfound": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterMarkNotFound",
		Description:    `Counter for pu mark not found.`,
		Exposed:        true,
		Name:           "counterMarkNotFound",
		Required:       true,
		Type:           "integer",
	},
	"counternetsynnotseen": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterNetSynNotSeen",
		Description:    `Counter for network syn packet was not seen.`,
		Exposed:        true,
		Name:           "counterNetSynNotSeen",
		Required:       true,
		Type:           "integer",
	},
	"counternoconnfound": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterNoConnFound",
		Description:    `Counter for no context or connection found.`,
		Exposed:        true,
		Name:           "counterNoConnFound",
		Required:       true,
		Type:           "integer",
	},
	"counternonputraffic": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterNonPUTraffic",
		Description:    `Counter for traffic that belongs to a non PU process.`,
		Exposed:        true,
		Name:           "counterNonPUTraffic",
		Required:       true,
		Type:           "integer",
	},
	"counteroutofordersynack": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterOutOfOrderSynAck",
		Description:    `Counter for synack for flow with processed finack.`,
		Exposed:        true,
		Name:           "counterOutOfOrderSynAck",
		Required:       true,
		Type:           "integer",
	},
	"counterportnotfound": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterPortNotFound",
		Description:    `Counter for port not found.`,
		Exposed:        true,
		Name:           "counterPortNotFound",
		Required:       true,
		Type:           "integer",
	},
	"counterrejectpacket": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterRejectPacket",
		Description:    `Counter for reject the packet as per policy.`,
		Exposed:        true,
		Name:           "counterRejectPacket",
		Required:       true,
		Type:           "integer",
	},
	"counterservicepostprocessorfailed": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterServicePostprocessorFailed",
		Description:    `Counter for post service processing failed for network packet.`,
		Exposed:        true,
		Name:           "counterServicePostprocessorFailed",
		Required:       true,
		Type:           "integer",
	},
	"counterservicepreprocessorfailed": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterServicePreprocessorFailed",
		Description:    `Counter for pre service processing failed for network packet.`,
		Exposed:        true,
		Name:           "counterServicePreprocessorFailed",
		Required:       true,
		Type:           "integer",
	},
	"countersynackbadclaims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynAckBadClaims",
		Description:    `Counter for synack packet dropped because of bad claims.`,
		Exposed:        true,
		Name:           "counterSynAckBadClaims",
		Required:       true,
		Type:           "integer",
	},
	"countersynackclaimsmismatch": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynAckClaimsMisMatch",
		Description:    `Counter for syn/ack packet dropped because of encryption mismatch.`,
		Exposed:        true,
		Name:           "counterSynAckClaimsMisMatch",
		Required:       true,
		Type:           "integer",
	},
	"countersynackdroppedexternalservice": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynAckDroppedExternalService",
		Description:    `Counter for synack from external service dropped.`,
		Exposed:        true,
		Name:           "counterSynAckDroppedExternalService",
		Required:       true,
		Type:           "integer",
	},
	"countersynackinvalidformat": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynAckInvalidFormat",
		Description:    `Counter for synack packet dropped because of invalid format.`,
		Exposed:        true,
		Name:           "counterSynAckInvalidFormat",
		Required:       true,
		Type:           "integer",
	},
	"countersynackmissingclaims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynAckMissingClaims",
		Description:    `Counter for synack packet dropped because of no claims.`,
		Exposed:        true,
		Name:           "counterSynAckMissingClaims",
		Required:       true,
		Type:           "integer",
	},
	"countersynackmissingtoken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynAckMissingToken",
		Description:    `Counter for synack packet dropped because of missing token.`,
		Exposed:        true,
		Name:           "counterSynAckMissingToken",
		Required:       true,
		Type:           "integer",
	},
	"countersynacknotcpauthoption": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynAckNoTCPAuthOption",
		Description:    `Counter for tcp authentication option not found.`,
		Exposed:        true,
		Name:           "counterSynAckNoTCPAuthOption",
		Required:       true,
		Type:           "integer",
	},
	"countersynackrejected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynAckRejected",
		Description:    `Counter for dropping because of reject rule on transmitter.`,
		Exposed:        true,
		Name:           "counterSynAckRejected",
		Required:       true,
		Type:           "integer",
	},
	"countersyndroppedinvalidformat": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynDroppedInvalidFormat",
		Description:    `Counter for syn packet dropped because of invalid format.`,
		Exposed:        true,
		Name:           "counterSynDroppedInvalidFormat",
		Required:       true,
		Type:           "integer",
	},
	"countersyndroppedinvalidtoken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynDroppedInvalidToken",
		Description:    `Counter for syn packet dropped because of invalid token.`,
		Exposed:        true,
		Name:           "counterSynDroppedInvalidToken",
		Required:       true,
		Type:           "integer",
	},
	"countersyndroppednoclaims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynDroppedNoClaims",
		Description:    `Counter for syn packet dropped because of no claims.`,
		Exposed:        true,
		Name:           "counterSynDroppedNoClaims",
		Required:       true,
		Type:           "integer",
	},
	"countersyndroppedtcpoption": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynDroppedTCPOption",
		Description:    `Counter for tcp authentication option not found.`,
		Exposed:        true,
		Name:           "counterSynDroppedTCPOption",
		Required:       true,
		Type:           "integer",
	},
	"countersynrejectpacket": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynRejectPacket",
		Description:    `Counter for syn dropped due to policy.`,
		Exposed:        true,
		Name:           "counterSynRejectPacket",
		Required:       true,
		Type:           "integer",
	},
	"countersynunexpectedpacket": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterSynUnexpectedPacket",
		Description:    `Counter for received syn packet from unknown pu.`,
		Exposed:        true,
		Name:           "counterSynUnexpectedPacket",
		Required:       true,
		Type:           "integer",
	},
	"countertcpauthnotfound": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterTCPAuthNotFound",
		Description:    `Counter for tcp authentication option not found.`,
		Exposed:        true,
		Name:           "counterTCPAuthNotFound",
		Required:       true,
		Type:           "integer",
	},
	"counterunknownerror": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CounterUnknownError",
		Description:    `Counter for unknown error.`,
		Exposed:        true,
		Name:           "counterUnknownError",
		Required:       true,
		Type:           "integer",
	},
	"enforcerid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		DefaultValue:   "xxxx-xxx-xxxx",
		Description:    `Identifier of the enforcer sending the report.`,
		Exposed:        true,
		Name:           "enforcerID",
		Stored:         true,
		Type:           "string",
	},
	"enforcernamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		DefaultValue:   "/my/namespace",
		Description:    `Namespace of the enforcer sending the report.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Stored:         true,
		Type:           "string",
	},
	"processingunitid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `PUID is the ID of the PU reporting the counter.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "processingUnitID",
		Type:           "string",
	},
	"processingunitnamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the PU reporting the counter.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "processingUnitNamespace",
		Type:           "string",
	},
	"timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Timestamp is the date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
}

// SparseCounterReportsList represents a list of SparseCounterReports
type SparseCounterReportsList []*SparseCounterReport

// Identity returns the identity of the objects in the list.
func (o SparseCounterReportsList) Identity() elemental.Identity {

	return CounterReportIdentity
}

// Copy returns a pointer to a copy the SparseCounterReportsList.
func (o SparseCounterReportsList) Copy() elemental.Identifiables {

	copy := append(SparseCounterReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCounterReportsList.
func (o SparseCounterReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCounterReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCounterReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCounterReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCounterReportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCounterReportsList converted to CounterReportsList.
func (o SparseCounterReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCounterReportsList) Version() int {

	return 1
}

// SparseCounterReport represents the sparse version of a counterreport.
type SparseCounterReport struct {
	// Counter for sending finack ack received in uknown connection state.
	CounterAckInUnknownState *int `json:"counterAckInUnknownState,omitempty" msgpack:"counterAckInUnknownState,omitempty" bson:"-" mapstructure:"counterAckInUnknownState,omitempty"`

	// Counter for ack packet dropped because of invalid format.
	CounterAckInvalidFormat *int `json:"counterAckInvalidFormat,omitempty" msgpack:"counterAckInvalidFormat,omitempty" bson:"-" mapstructure:"counterAckInvalidFormat,omitempty"`

	// Counter for reject ack packet as per policy.
	CounterAckRejected *int `json:"counterAckRejected,omitempty" msgpack:"counterAckRejected,omitempty" bson:"-" mapstructure:"counterAckRejected,omitempty"`

	// Counter for ack packet dropped because signature validation failed.
	CounterAckSigValidationFailed *int `json:"counterAckSigValidationFailed,omitempty" msgpack:"counterAckSigValidationFailed,omitempty" bson:"-" mapstructure:"counterAckSigValidationFailed,omitempty"`

	// Counter for tcp authentication option not found.
	CounterAckTCPNoTCPAuthOption *int `json:"counterAckTCPNoTCPAuthOption,omitempty" msgpack:"counterAckTCPNoTCPAuthOption,omitempty" bson:"-" mapstructure:"counterAckTCPNoTCPAuthOption,omitempty"`

	// Counter for ".
	CounterConnectionsProcessed *int `json:"counterConnectionsProcessed,omitempty" msgpack:"counterConnectionsProcessed,omitempty" bson:"-" mapstructure:"counterConnectionsProcessed,omitempty"`

	// Counter for unable to find contextid.
	CounterContextIDNotFound *int `json:"counterContextIDNotFound,omitempty" msgpack:"counterContextIDNotFound,omitempty" bson:"-" mapstructure:"counterContextIDNotFound,omitempty"`

	// Counter for no acls found for external services. dropping application syn
	// packet.
	CounterDroppedExternalService *int `json:"counterDroppedExternalService,omitempty" msgpack:"counterDroppedExternalService,omitempty" bson:"-" mapstructure:"counterDroppedExternalService,omitempty"`

	// Counter for invalid connection state.
	CounterInvalidConnState *int `json:"counterInvalidConnState,omitempty" msgpack:"counterInvalidConnState,omitempty" bson:"-" mapstructure:"counterInvalidConnState,omitempty"`

	// Counter for invalid net state.
	CounterInvalidNetState *int `json:"counterInvalidNetState,omitempty" msgpack:"counterInvalidNetState,omitempty" bson:"-" mapstructure:"counterInvalidNetState,omitempty"`

	// Counter for invalid protocol.
	CounterInvalidProtocol *int `json:"counterInvalidProtocol,omitempty" msgpack:"counterInvalidProtocol,omitempty" bson:"-" mapstructure:"counterInvalidProtocol,omitempty"`

	// Counter for pu is already dead - drop synack packet.
	CounterInvalidSynAck *int `json:"counterInvalidSynAck,omitempty" msgpack:"counterInvalidSynAck,omitempty" bson:"-" mapstructure:"counterInvalidSynAck,omitempty"`

	// Counter for pu mark not found.
	CounterMarkNotFound *int `json:"counterMarkNotFound,omitempty" msgpack:"counterMarkNotFound,omitempty" bson:"-" mapstructure:"counterMarkNotFound,omitempty"`

	// Counter for network syn packet was not seen.
	CounterNetSynNotSeen *int `json:"counterNetSynNotSeen,omitempty" msgpack:"counterNetSynNotSeen,omitempty" bson:"-" mapstructure:"counterNetSynNotSeen,omitempty"`

	// Counter for no context or connection found.
	CounterNoConnFound *int `json:"counterNoConnFound,omitempty" msgpack:"counterNoConnFound,omitempty" bson:"-" mapstructure:"counterNoConnFound,omitempty"`

	// Counter for traffic that belongs to a non PU process.
	CounterNonPUTraffic *int `json:"counterNonPUTraffic,omitempty" msgpack:"counterNonPUTraffic,omitempty" bson:"-" mapstructure:"counterNonPUTraffic,omitempty"`

	// Counter for synack for flow with processed finack.
	CounterOutOfOrderSynAck *int `json:"counterOutOfOrderSynAck,omitempty" msgpack:"counterOutOfOrderSynAck,omitempty" bson:"-" mapstructure:"counterOutOfOrderSynAck,omitempty"`

	// Counter for port not found.
	CounterPortNotFound *int `json:"counterPortNotFound,omitempty" msgpack:"counterPortNotFound,omitempty" bson:"-" mapstructure:"counterPortNotFound,omitempty"`

	// Counter for reject the packet as per policy.
	CounterRejectPacket *int `json:"counterRejectPacket,omitempty" msgpack:"counterRejectPacket,omitempty" bson:"-" mapstructure:"counterRejectPacket,omitempty"`

	// Counter for post service processing failed for network packet.
	CounterServicePostprocessorFailed *int `json:"counterServicePostprocessorFailed,omitempty" msgpack:"counterServicePostprocessorFailed,omitempty" bson:"-" mapstructure:"counterServicePostprocessorFailed,omitempty"`

	// Counter for pre service processing failed for network packet.
	CounterServicePreprocessorFailed *int `json:"counterServicePreprocessorFailed,omitempty" msgpack:"counterServicePreprocessorFailed,omitempty" bson:"-" mapstructure:"counterServicePreprocessorFailed,omitempty"`

	// Counter for synack packet dropped because of bad claims.
	CounterSynAckBadClaims *int `json:"counterSynAckBadClaims,omitempty" msgpack:"counterSynAckBadClaims,omitempty" bson:"-" mapstructure:"counterSynAckBadClaims,omitempty"`

	// Counter for syn/ack packet dropped because of encryption mismatch.
	CounterSynAckClaimsMisMatch *int `json:"counterSynAckClaimsMisMatch,omitempty" msgpack:"counterSynAckClaimsMisMatch,omitempty" bson:"-" mapstructure:"counterSynAckClaimsMisMatch,omitempty"`

	// Counter for synack from external service dropped.
	CounterSynAckDroppedExternalService *int `json:"counterSynAckDroppedExternalService,omitempty" msgpack:"counterSynAckDroppedExternalService,omitempty" bson:"-" mapstructure:"counterSynAckDroppedExternalService,omitempty"`

	// Counter for synack packet dropped because of invalid format.
	CounterSynAckInvalidFormat *int `json:"counterSynAckInvalidFormat,omitempty" msgpack:"counterSynAckInvalidFormat,omitempty" bson:"-" mapstructure:"counterSynAckInvalidFormat,omitempty"`

	// Counter for synack packet dropped because of no claims.
	CounterSynAckMissingClaims *int `json:"counterSynAckMissingClaims,omitempty" msgpack:"counterSynAckMissingClaims,omitempty" bson:"-" mapstructure:"counterSynAckMissingClaims,omitempty"`

	// Counter for synack packet dropped because of missing token.
	CounterSynAckMissingToken *int `json:"counterSynAckMissingToken,omitempty" msgpack:"counterSynAckMissingToken,omitempty" bson:"-" mapstructure:"counterSynAckMissingToken,omitempty"`

	// Counter for tcp authentication option not found.
	CounterSynAckNoTCPAuthOption *int `json:"counterSynAckNoTCPAuthOption,omitempty" msgpack:"counterSynAckNoTCPAuthOption,omitempty" bson:"-" mapstructure:"counterSynAckNoTCPAuthOption,omitempty"`

	// Counter for dropping because of reject rule on transmitter.
	CounterSynAckRejected *int `json:"counterSynAckRejected,omitempty" msgpack:"counterSynAckRejected,omitempty" bson:"-" mapstructure:"counterSynAckRejected,omitempty"`

	// Counter for syn packet dropped because of invalid format.
	CounterSynDroppedInvalidFormat *int `json:"counterSynDroppedInvalidFormat,omitempty" msgpack:"counterSynDroppedInvalidFormat,omitempty" bson:"-" mapstructure:"counterSynDroppedInvalidFormat,omitempty"`

	// Counter for syn packet dropped because of invalid token.
	CounterSynDroppedInvalidToken *int `json:"counterSynDroppedInvalidToken,omitempty" msgpack:"counterSynDroppedInvalidToken,omitempty" bson:"-" mapstructure:"counterSynDroppedInvalidToken,omitempty"`

	// Counter for syn packet dropped because of no claims.
	CounterSynDroppedNoClaims *int `json:"counterSynDroppedNoClaims,omitempty" msgpack:"counterSynDroppedNoClaims,omitempty" bson:"-" mapstructure:"counterSynDroppedNoClaims,omitempty"`

	// Counter for tcp authentication option not found.
	CounterSynDroppedTCPOption *int `json:"counterSynDroppedTCPOption,omitempty" msgpack:"counterSynDroppedTCPOption,omitempty" bson:"-" mapstructure:"counterSynDroppedTCPOption,omitempty"`

	// Counter for syn dropped due to policy.
	CounterSynRejectPacket *int `json:"counterSynRejectPacket,omitempty" msgpack:"counterSynRejectPacket,omitempty" bson:"-" mapstructure:"counterSynRejectPacket,omitempty"`

	// Counter for received syn packet from unknown pu.
	CounterSynUnexpectedPacket *int `json:"counterSynUnexpectedPacket,omitempty" msgpack:"counterSynUnexpectedPacket,omitempty" bson:"-" mapstructure:"counterSynUnexpectedPacket,omitempty"`

	// Counter for tcp authentication option not found.
	CounterTCPAuthNotFound *int `json:"counterTCPAuthNotFound,omitempty" msgpack:"counterTCPAuthNotFound,omitempty" bson:"-" mapstructure:"counterTCPAuthNotFound,omitempty"`

	// Counter for unknown error.
	CounterUnknownError *int `json:"counterUnknownError,omitempty" msgpack:"counterUnknownError,omitempty" bson:"-" mapstructure:"counterUnknownError,omitempty"`

	// Identifier of the enforcer sending the report.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"enforcerid,omitempty" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer sending the report.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"enforcernamespace,omitempty" mapstructure:"enforcerNamespace,omitempty"`

	// PUID is the ID of the PU reporting the counter.
	ProcessingUnitID *string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the PU reporting the counter.
	ProcessingUnitNamespace *string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"-" mapstructure:"processingUnitNamespace,omitempty"`

	// Timestamp is the date of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCounterReport returns a new  SparseCounterReport.
func NewSparseCounterReport() *SparseCounterReport {
	return &SparseCounterReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCounterReport) Identity() elemental.Identity {

	return CounterReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCounterReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCounterReport) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseCounterReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCounterReport) ToPlain() elemental.PlainIdentifiable {

	out := NewCounterReport()
	if o.CounterAckInUnknownState != nil {
		out.CounterAckInUnknownState = *o.CounterAckInUnknownState
	}
	if o.CounterAckInvalidFormat != nil {
		out.CounterAckInvalidFormat = *o.CounterAckInvalidFormat
	}
	if o.CounterAckRejected != nil {
		out.CounterAckRejected = *o.CounterAckRejected
	}
	if o.CounterAckSigValidationFailed != nil {
		out.CounterAckSigValidationFailed = *o.CounterAckSigValidationFailed
	}
	if o.CounterAckTCPNoTCPAuthOption != nil {
		out.CounterAckTCPNoTCPAuthOption = *o.CounterAckTCPNoTCPAuthOption
	}
	if o.CounterConnectionsProcessed != nil {
		out.CounterConnectionsProcessed = *o.CounterConnectionsProcessed
	}
	if o.CounterContextIDNotFound != nil {
		out.CounterContextIDNotFound = *o.CounterContextIDNotFound
	}
	if o.CounterDroppedExternalService != nil {
		out.CounterDroppedExternalService = *o.CounterDroppedExternalService
	}
	if o.CounterInvalidConnState != nil {
		out.CounterInvalidConnState = *o.CounterInvalidConnState
	}
	if o.CounterInvalidNetState != nil {
		out.CounterInvalidNetState = *o.CounterInvalidNetState
	}
	if o.CounterInvalidProtocol != nil {
		out.CounterInvalidProtocol = *o.CounterInvalidProtocol
	}
	if o.CounterInvalidSynAck != nil {
		out.CounterInvalidSynAck = *o.CounterInvalidSynAck
	}
	if o.CounterMarkNotFound != nil {
		out.CounterMarkNotFound = *o.CounterMarkNotFound
	}
	if o.CounterNetSynNotSeen != nil {
		out.CounterNetSynNotSeen = *o.CounterNetSynNotSeen
	}
	if o.CounterNoConnFound != nil {
		out.CounterNoConnFound = *o.CounterNoConnFound
	}
	if o.CounterNonPUTraffic != nil {
		out.CounterNonPUTraffic = *o.CounterNonPUTraffic
	}
	if o.CounterOutOfOrderSynAck != nil {
		out.CounterOutOfOrderSynAck = *o.CounterOutOfOrderSynAck
	}
	if o.CounterPortNotFound != nil {
		out.CounterPortNotFound = *o.CounterPortNotFound
	}
	if o.CounterRejectPacket != nil {
		out.CounterRejectPacket = *o.CounterRejectPacket
	}
	if o.CounterServicePostprocessorFailed != nil {
		out.CounterServicePostprocessorFailed = *o.CounterServicePostprocessorFailed
	}
	if o.CounterServicePreprocessorFailed != nil {
		out.CounterServicePreprocessorFailed = *o.CounterServicePreprocessorFailed
	}
	if o.CounterSynAckBadClaims != nil {
		out.CounterSynAckBadClaims = *o.CounterSynAckBadClaims
	}
	if o.CounterSynAckClaimsMisMatch != nil {
		out.CounterSynAckClaimsMisMatch = *o.CounterSynAckClaimsMisMatch
	}
	if o.CounterSynAckDroppedExternalService != nil {
		out.CounterSynAckDroppedExternalService = *o.CounterSynAckDroppedExternalService
	}
	if o.CounterSynAckInvalidFormat != nil {
		out.CounterSynAckInvalidFormat = *o.CounterSynAckInvalidFormat
	}
	if o.CounterSynAckMissingClaims != nil {
		out.CounterSynAckMissingClaims = *o.CounterSynAckMissingClaims
	}
	if o.CounterSynAckMissingToken != nil {
		out.CounterSynAckMissingToken = *o.CounterSynAckMissingToken
	}
	if o.CounterSynAckNoTCPAuthOption != nil {
		out.CounterSynAckNoTCPAuthOption = *o.CounterSynAckNoTCPAuthOption
	}
	if o.CounterSynAckRejected != nil {
		out.CounterSynAckRejected = *o.CounterSynAckRejected
	}
	if o.CounterSynDroppedInvalidFormat != nil {
		out.CounterSynDroppedInvalidFormat = *o.CounterSynDroppedInvalidFormat
	}
	if o.CounterSynDroppedInvalidToken != nil {
		out.CounterSynDroppedInvalidToken = *o.CounterSynDroppedInvalidToken
	}
	if o.CounterSynDroppedNoClaims != nil {
		out.CounterSynDroppedNoClaims = *o.CounterSynDroppedNoClaims
	}
	if o.CounterSynDroppedTCPOption != nil {
		out.CounterSynDroppedTCPOption = *o.CounterSynDroppedTCPOption
	}
	if o.CounterSynRejectPacket != nil {
		out.CounterSynRejectPacket = *o.CounterSynRejectPacket
	}
	if o.CounterSynUnexpectedPacket != nil {
		out.CounterSynUnexpectedPacket = *o.CounterSynUnexpectedPacket
	}
	if o.CounterTCPAuthNotFound != nil {
		out.CounterTCPAuthNotFound = *o.CounterTCPAuthNotFound
	}
	if o.CounterUnknownError != nil {
		out.CounterUnknownError = *o.CounterUnknownError
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		out.EnforcerNamespace = *o.EnforcerNamespace
	}
	if o.ProcessingUnitID != nil {
		out.ProcessingUnitID = *o.ProcessingUnitID
	}
	if o.ProcessingUnitNamespace != nil {
		out.ProcessingUnitNamespace = *o.ProcessingUnitNamespace
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCounterReport.
func (o *SparseCounterReport) DeepCopy() *SparseCounterReport {

	if o == nil {
		return nil
	}

	out := &SparseCounterReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCounterReport.
func (o *SparseCounterReport) DeepCopyInto(out *SparseCounterReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCounterReport: %s", err))
	}

	*out = *target.(*SparseCounterReport)
}
