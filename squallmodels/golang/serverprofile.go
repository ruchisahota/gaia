package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"

// ServerProfileDockerSocketTypeValue represents the possible values for attribute "dockerSocketType".
type ServerProfileDockerSocketTypeValue string

const (
	// ServerProfileDockerSocketTypeTcp represents the value tcp.
	ServerProfileDockerSocketTypeTcp ServerProfileDockerSocketTypeValue = "tcp"

	// ServerProfileDockerSocketTypeUnix represents the value unix.
	ServerProfileDockerSocketTypeUnix ServerProfileDockerSocketTypeValue = "unix"
)

// ServerProfileIdentity represents the Identity of the object
var ServerProfileIdentity = elemental.Identity{
	Name:     "serverprofile",
	Category: "serverprofiles",
}

// ServerProfilesList represents a list of ServerProfiles
type ServerProfilesList []*ServerProfile

// ContentIdentity returns the identity of the objects in the list.
func (o ServerProfilesList) ContentIdentity() elemental.Identity {
	return ServerProfileIdentity
}

// List convert the object to and elemental.IdentifiablesList.
func (o ServerProfilesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// ServerProfile represents the model of a serverprofile
type ServerProfile struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id"`

	// IptablesMarkValue is the mark value to be used in an iptables implementation.
	IPTablesMarkValue int `json:"IPTablesMarkValue" bson:"iptablesmarkvalue"`

	// IgnoreExpression allows to set a tag expression that will make Aporeto to ignore docker container started with labels matching the rule.
	IgnoreExpression [][]string `json:"IgnoreExpression" bson:"ignoreexpression"`

	// PUBookkeepingInterval configures how often the PU will be synchronized.
	PUBookkeepingInterval string `json:"PUBookkeepingInterval" bson:"pubookkeepinginterval"`

	// PUHeartbeatInterval configures the heart beat interval.
	PUHeartbeatInterval string `json:"PUHeartbeatInterval" bson:"puheartbeatinterval"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags"`

	// CreatedTime is the time at which the object was created
	CreateTime time.Time `json:"createTime" bson:"createtime"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description"`

	// DockerSocketAddress is the address of the docker daemon.
	DockerSocketAddress string `json:"dockerSocketAddress" bson:"dockersocketaddress"`

	// DockerSocketType is the type of socket to use to talk to the docker daemon.
	DockerSocketType ServerProfileDockerSocketTypeValue `json:"dockerSocketType" bson:"dockersockettype"`

	// kubernetesEnable enables  kubernetes mode for the agent.
	KubernetesEnable bool `json:"kubernetesEnable" bson:"kubernetesenable"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" bson:"namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags"`

	// PolicySynchronizationInterval configures how often the policy will be resynchronized.
	PolicySynchronizationInterval string `json:"policySynchronizationInterval" bson:"policysynchronizationinterval"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected"`

	// AgentPort is the port the agent should use to listen for API calls
	ProxyListenAddress string `json:"proxyListenAddress" bson:"proxylistenaddress"`

	// ReceiverNumberOfQueues is the number of queues for the NFQUEUE of the network receiver starting at the ReceiverQueue
	ReceiverNumberOfQueues int `json:"receiverNumberOfQueues" bson:"receivernumberofqueues"`

	// ReceiverQueue is the base queue number for traffic from the network.
	ReceiverQueue int `json:"receiverQueue" bson:"receiverqueue"`

	// ReceiverQueueSize is the queue size of the receiver
	ReceiverQueueSize int `json:"receiverQueueSize" bson:"receiverqueuesize"`

	// RemoteEnforcer inidicates whether a single enforcer should be used or a distributed enforcer. True means distributed.
	RemoteEnforcer bool `json:"remoteEnforcer" bson:"remoteenforcer"`

	// SupportLinuxProcesses configures support for Linux processes.
	SupportLinuxProcesses bool `json:"supportLinuxProcesses" bson:"supportlinuxprocesses"`

	// TargetNetworks is the list of networks that authorization should be applied.
	TargetNetworks []string `json:"targetNetworks" bson:"targetnetworks"`

	// TransmitterNumberOfQueues is the number of queues for application traffic.
	TransmitterNumberOfQueues int `json:"transmitterNumberOfQueues" bson:"transmitternumberofqueues"`

	// TransmitterQueue is the queue number for traffic from the applications starting at the transmitterQueue
	TransmitterQueue int `json:"transmitterQueue" bson:"transmitterqueue"`

	// TransmitterQueueSize is the size of the queue for application traffic.
	TransmitterQueueSize int `json:"transmitterQueueSize" bson:"transmitterqueuesize"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`
}

// NewServerProfile returns a new *ServerProfile
func NewServerProfile() *ServerProfile {

	return &ServerProfile{
		ModelVersion:                  1.0,
		IPTablesMarkValue:             1000,
		PUBookkeepingInterval:         "5m",
		PUHeartbeatInterval:           "5s",
		AssociatedTags:                []string{},
		DockerSocketAddress:           "/var/run/docker.sock",
		DockerSocketType:              "unix",
		KubernetesEnable:              false,
		NormalizedTags:                []string{},
		PolicySynchronizationInterval: "10m",
		ProxyListenAddress:            ":9443",
		ReceiverNumberOfQueues:        4,
		ReceiverQueue:                 0,
		ReceiverQueueSize:             500,
		RemoteEnforcer:                false,
		SupportLinuxProcesses:         false,
		TransmitterNumberOfQueues:     4,
		TransmitterQueue:              4,
		TransmitterQueueSize:          500,
	}
}

// Identity returns the Identity of the object.
func (o *ServerProfile) Identity() elemental.Identity {

	return ServerProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ServerProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ServerProfile) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *ServerProfile) Version() float64 {

	return 1.0
}

func (o *ServerProfile) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *ServerProfile) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *ServerProfile) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreateTime set the given createTime of the receiver
func (o *ServerProfile) SetCreateTime(createTime time.Time) {
	o.CreateTime = createTime
}

// GetName returns the name of the receiver
func (o *ServerProfile) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *ServerProfile) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *ServerProfile) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *ServerProfile) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *ServerProfile) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *ServerProfile) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetProtected returns the protected of the receiver
func (o *ServerProfile) GetProtected() bool {
	return o.Protected
}

// SetUpdateTime set the given updateTime of the receiver
func (o *ServerProfile) SetUpdateTime(updateTime time.Time) {
	o.UpdateTime = updateTime
}

// Validate valides the current information stored into the structure.
func (o *ServerProfile) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumInt("IPTablesMarkValue", o.IPTablesMarkValue, 65000, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("IPTablesMarkValue", o.IPTablesMarkValue, 0, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("PUBookkeepingInterval", o.PUBookkeepingInterval, `^[0-9]+[smh]$`); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("PUHeartbeatInterval", o.PUHeartbeatInterval, `^[0-9]+[smh]$`); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("dockerSocketType", string(o.DockerSocketType), []string{"tcp", "unix"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("policySynchronizationInterval", o.PolicySynchronizationInterval, `^[0-9]+[smh]$`); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("receiverNumberOfQueues", o.ReceiverNumberOfQueues, 16, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("receiverNumberOfQueues", o.ReceiverNumberOfQueues, 1, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("receiverQueue", o.ReceiverQueue, 1000, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("receiverQueue", o.ReceiverQueue, 0, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("receiverQueueSize", o.ReceiverQueueSize, 5000, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("receiverQueueSize", o.ReceiverQueueSize, 1, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredExternal("targetNetworks", o.TargetNetworks); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("targetNetworks", o.TargetNetworks); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("transmitterNumberOfQueues", o.TransmitterNumberOfQueues, 16, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("transmitterNumberOfQueues", o.TransmitterNumberOfQueues, 1, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("transmitterQueue", o.TransmitterQueue, 1000, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("transmitterQueue", o.TransmitterQueue, 1, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("transmitterQueueSize", o.TransmitterQueueSize, 1000, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("transmitterQueueSize", o.TransmitterQueueSize, 1, false); err != nil {
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
func (ServerProfile) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return ServerProfileAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (ServerProfile) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ServerProfileAttributesMap
}

// ServerProfileAttributesMap represents the map of attribute for ServerProfile.
var ServerProfileAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID is the identifier of the object.`,
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
		Unique:         true,
	},
	"IPTablesMarkValue": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		DefaultValue:   1000,
		Description:    `IptablesMarkValue is the mark value to be used in an iptables implementation.`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       65000,
		Name:           "IPTablesMarkValue",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"IgnoreExpression": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `IgnoreExpression allows to set a tag expression that will make Aporeto to ignore docker container started with labels matching the rule.`,
		Exposed:        true,
		Name:           "IgnoreExpression",
		Stored:         true,
		SubType:        "policies_list",
		Type:           "external",
	},
	"PUBookkeepingInterval": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		DefaultValue:   "5m",
		Description:    `PUBookkeepingInterval configures how often the PU will be synchronized.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "PUBookkeepingInterval",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"PUHeartbeatInterval": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		DefaultValue:   "5s",
		Description:    `PUHeartbeatInterval configures the heart beat interval.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "PUHeartbeatInterval",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Annotation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Annotation stores additional information about an entity`,
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AssociatedTags are the list of tags attached to an entity`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `CreatedTime is the time at which the object was created`,
		Exposed:        true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"DockerSocketAddress": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		DefaultValue:   "/var/run/docker.sock",
		Description:    `DockerSocketAddress is the address of the docker daemon.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "dockerSocketAddress",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"DockerSocketType": elemental.AttributeSpecification{
		AllowedChoices: []string{"tcp", "unix"},
		DefaultValue:   ServerProfileDockerSocketTypeValue("unix"),
		Description:    `DockerSocketType is the type of socket to use to talk to the docker daemon.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "dockerSocketType",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"KubernetesEnable": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		DefaultValue:   false,
		Description:    `kubernetesEnable enables  kubernetes mode for the agent.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "kubernetesEnable",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Name is the name of the entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Index:          true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `NormalizedTags contains the list of normalized tags of the entities`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Transient:      true,
		Type:           "external",
	},
	"PolicySynchronizationInterval": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		DefaultValue:   "10m",
		Description:    `PolicySynchronizationInterval configures how often the policy will be resynchronized.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "policySynchronizationInterval",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"ProxyListenAddress": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		DefaultValue:   ":9443",
		Description:    `AgentPort is the port the agent should use to listen for API calls `,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "proxyListenAddress",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ReceiverNumberOfQueues": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		DefaultValue:   4,
		Description:    `ReceiverNumberOfQueues is the number of queues for the NFQUEUE of the network receiver starting at the ReceiverQueue`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       16,
		MinValue:       1,
		Name:           "receiverNumberOfQueues",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"ReceiverQueue": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		DefaultValue:   0,
		Description:    `ReceiverQueue is the base queue number for traffic from the network.`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       1000,
		Name:           "receiverQueue",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"ReceiverQueueSize": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		DefaultValue:   500,
		Description:    `ReceiverQueueSize is the queue size of the receiver`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       5000,
		MinValue:       1,
		Name:           "receiverQueueSize",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"RemoteEnforcer": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		DefaultValue:   false,
		Description:    `RemoteEnforcer inidicates whether a single enforcer should be used or a distributed enforcer. True means distributed.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "remoteEnforcer",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"SupportLinuxProcesses": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		DefaultValue:   false,
		Description:    `SupportLinuxProcesses configures support for Linux processes.`,
		Exposed:        true,
		Name:           "supportLinuxProcesses",
		Stored:         true,
		Type:           "boolean",
	},
	"TargetNetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `TargetNetworks is the list of networks that authorization should be applied.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "targetNetworks",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		SubType:        "target_networks_list",
		Type:           "external",
	},
	"TransmitterNumberOfQueues": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		DefaultValue:   4,
		Description:    `TransmitterNumberOfQueues is the number of queues for application traffic.`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       16,
		MinValue:       1,
		Name:           "transmitterNumberOfQueues",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"TransmitterQueue": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		DefaultValue:   4,
		Description:    `TransmitterQueue is the queue number for traffic from the applications starting at the transmitterQueue`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       1000,
		MinValue:       1,
		Name:           "transmitterQueue",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"TransmitterQueueSize": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		DefaultValue:   500,
		Description:    `TransmitterQueueSize is the size of the queue for application traffic.`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       1000,
		MinValue:       1,
		Name:           "transmitterQueueSize",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `UpdateTime is the time at which an entity was updated.`,
		Exposed:        true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}
