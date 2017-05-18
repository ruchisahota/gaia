package rufusmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// RemoteProcessorTypeValue represents the possible values for attribute "type".
type RemoteProcessorTypeValue string

const (
	// RemoteProcessorTypePosthook represents the value PostHook.
	RemoteProcessorTypePosthook RemoteProcessorTypeValue = "PostHook"

	// RemoteProcessorTypePrehook represents the value PreHook.
	RemoteProcessorTypePrehook RemoteProcessorTypeValue = "PreHook"
)

// RemoteProcessorIdentity represents the Identity of the object
var RemoteProcessorIdentity = elemental.Identity{
	Name:     "remoteprocessor",
	Category: "remoteprocessors",
}

// RemoteProcessorsList represents a list of RemoteProcessors
type RemoteProcessorsList []*RemoteProcessor

// ContentIdentity returns the identity of the objects in the list.
func (o RemoteProcessorsList) ContentIdentity() elemental.Identity {

	return RemoteProcessorIdentity
}

// List converts the object to and elemental.IdentifiablesList.
func (o RemoteProcessorsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o RemoteProcessorsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// RemoteProcessor represents the model of a remoteprocessor
type RemoteProcessor struct {
	// Represents the claims of the currently managed object
	Claims []string `json:"claims" bson:"claims"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description"`

	// Disabled defines if the propert is disabled.
	Disabled bool `json:"disabled" bson:"disabled"`

	// Represents data received from the service
	InputData interface{} `json:"inputData" bson:"inputdata"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name"`

	// Define the operation that is currently handled by the service
	Operation elemental.Operation `json:"operation" bson:"operation"`

	// Returns the OutputData filled with the processor information
	OutputData interface{} `json:"outputData" bson:"outputdata"`

	// Represents the Identity name of the managed object
	TargetIdentity string `json:"targetIdentity" bson:"targetidentity"`

	// Defines the type of the hook
	Type RemoteProcessorTypeValue `json:"type" bson:"type"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewRemoteProcessor returns a new *RemoteProcessor
func NewRemoteProcessor() *RemoteProcessor {

	return &RemoteProcessor{
		ModelVersion: 1.0,
	}
}

// Identity returns the Identity of the object.
func (o *RemoteProcessor) Identity() elemental.Identity {

	return RemoteProcessorIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *RemoteProcessor) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *RemoteProcessor) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *RemoteProcessor) Version() float64 {

	return 1.0
}

// DefaultOrder returns the list of default ordering fields.
func (o *RemoteProcessor) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *RemoteProcessor) Doc() string {
	return `Hook to integrate in an Aporeto service`
}

func (o *RemoteProcessor) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetDisabled returns the disabled of the receiver
func (o *RemoteProcessor) GetDisabled() bool {
	return o.Disabled
}

// SetDisabled set the given disabled of the receiver
func (o *RemoteProcessor) SetDisabled(disabled bool) {
	o.Disabled = disabled
}

// GetName returns the name of the receiver
func (o *RemoteProcessor) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *RemoteProcessor) SetName(name string) {
	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *RemoteProcessor) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredExternal("operation", o.Operation); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("operation", o.Operation); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"PostHook", "PreHook"}, false); err != nil {
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
func (*RemoteProcessor) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return RemoteProcessorAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*RemoteProcessor) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RemoteProcessorAttributesMap
}

// RemoteProcessorAttributesMap represents the map of attribute for RemoteProcessor.
var RemoteProcessorAttributesMap = map[string]elemental.AttributeSpecification{
	"Claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Represents the claims of the currently managed object`,
		Exposed:        true,
		Filterable:     true,
		Name:           "claims",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"Disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"InputData": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Represents data received from the service`,
		Exposed:        true,
		Filterable:     true,
		Name:           "inputData",
		Orderable:      true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		DefaultOrder:   true,
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
	"Operation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Define the operation that is currently handled by the service`,
		Exposed:        true,
		Filterable:     true,
		Name:           "operation",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		SubType:        "elemental_operation",
		Type:           "external",
	},
	"OutputData": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Returns the OutputData filled with the processor information`,
		Exposed:        true,
		Filterable:     true,
		Name:           "outputData",
		Orderable:      true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"TargetIdentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Represents the Identity name of the managed object`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "targetIdentity",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"PostHook", "PreHook"},
		Description:    `Defines the type of the hook`,
		Exposed:        true,
		Filterable:     true,
		Name:           "type",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "enum",
	},
}
