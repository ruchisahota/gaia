package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"

// ActivityIdentity represents the Identity of the object
var ActivityIdentity = elemental.Identity{
	Name:     "activity",
	Category: "activities",
}

// ActivitiesList represents a list of Activities
type ActivitiesList []*Activity

// Activity represents the model of a activity
type Activity struct {
	// ID is the identifier of the object.
	ID string `json:"ID" cql:"id,primarykey,omitempty" bson:"_id"`

	// Claims of the user who performed the operation.
	Claims interface{} `json:"claims" cql:"claims,omitempty" bson:"claims"`

	// Data of the notification.
	Data interface{} `json:"data" cql:"data,omitempty" bson:"data"`

	// Date of the notification.
	Date time.Time `json:"date" cql:"date,omitempty" bson:"date"`

	// Error contains the eventual error.
	Error interface{} `json:"error" cql:"error,omitempty" bson:"error"`

	// Message of the notification.
	Message string `json:"message" cql:"message,omitempty" bson:"message"`

	// Namespace of the notification.
	Namespace string `json:"namespace" cql:"namespace,omitempty" bson:"namespace"`

	// Operation describe what kind of operation the notification represents.
	Operation string `json:"operation" cql:"operation,omitempty" bson:"operation"`

	// TargetIdentity is the Identity of the related object.
	TargetIdentity string `json:"targetIdentity" cql:"targetidentity,omitempty" bson:"targetidentity"`
}

// NewActivity returns a new *Activity
func NewActivity() *Activity {

	return &Activity{}
}

// Identity returns the Identity of the object.
func (o *Activity) Identity() elemental.Identity {

	return ActivityIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Activity) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Activity) SetIdentifier(ID string) {

	o.ID = ID
}

func (o *Activity) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Activity) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (Activity) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return ActivityAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (Activity) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ActivityAttributesMap
}

// ActivityAttributesMap represents the map of attribute for Activity.
var ActivityAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Claims of the user who performed the operation.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "claims",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"Data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Data of the notification.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "data",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"Date": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Date of the notification.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "date",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"Error": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Error contains the eventual error.`,
		Exposed:        true,
		Name:           "error",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		SubType:        "raw_data",
		Type:           "external",
	},
	"Message": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Message of the notification.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "message",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Namespace of the notification.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Operation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Operation describe what kind of operation the notification represents.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "operation",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"TargetIdentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `TargetIdentity is the Identity of the related object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "targetIdentity",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
}
