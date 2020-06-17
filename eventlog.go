package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// EventLogLevelValue represents the possible values for attribute "level".
type EventLogLevelValue string

const (
	// EventLogLevelCritical represents the value Critical.
	EventLogLevelCritical EventLogLevelValue = "Critical"

	// EventLogLevelDebug represents the value Debug.
	EventLogLevelDebug EventLogLevelValue = "Debug"

	// EventLogLevelError represents the value Error.
	EventLogLevelError EventLogLevelValue = "Error"

	// EventLogLevelInfo represents the value Info.
	EventLogLevelInfo EventLogLevelValue = "Info"

	// EventLogLevelWarning represents the value Warning.
	EventLogLevelWarning EventLogLevelValue = "Warning"
)

// EventLogIdentity represents the Identity of the object.
var EventLogIdentity = elemental.Identity{
	Name:     "eventlog",
	Category: "eventlogs",
	Package:  "leon",
	Private:  false,
}

// EventLogsList represents a list of EventLogs
type EventLogsList []*EventLog

// Identity returns the identity of the objects in the list.
func (o EventLogsList) Identity() elemental.Identity {

	return EventLogIdentity
}

// Copy returns a pointer to a copy the EventLogsList.
func (o EventLogsList) Copy() elemental.Identifiables {

	copy := append(EventLogsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EventLogsList.
func (o EventLogsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(EventLogsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*EventLog))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EventLogsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EventLogsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the EventLogsList converted to SparseEventLogsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o EventLogsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseEventLogsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseEventLog)
	}

	return out
}

// Version returns the version of the content.
func (o EventLogsList) Version() int {

	return 1
}

// EventLog represents the model of a eventlog
type EventLog struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Category of the event log.
	Category string `json:"category" msgpack:"category" bson:"category" mapstructure:"category,omitempty"`

	// Content of the event log.
	Content string `json:"content" msgpack:"content" bson:"content" mapstructure:"content,omitempty"`

	// Creation date of the event log.
	Date time.Time `json:"date" msgpack:"date" bson:"date" mapstructure:"date,omitempty"`

	// Sets the log level.
	Level EventLogLevelValue `json:"level" msgpack:"level" bson:"level" mapstructure:"level,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Namespace tag attached to the event log.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Opaque data that can be attached to the event log, for further machine
	// processing.
	Opaque string `json:"opaque" msgpack:"opaque" bson:"opaque" mapstructure:"opaque,omitempty"`

	// ID of the object this event log is attached to. The object must be in the same
	// namespace than the event log.
	TargetID string `json:"targetID" msgpack:"targetID" bson:"targetid" mapstructure:"targetID,omitempty"`

	// Identity of the object this event log is attached to.
	TargetIdentity string `json:"targetIdentity" msgpack:"targetIdentity" bson:"targetidentity" mapstructure:"targetIdentity,omitempty"`

	// Title of the event log.
	Title string `json:"title" msgpack:"title" bson:"title" mapstructure:"title,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone int `json:"zone" msgpack:"zone" bson:"zone" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewEventLog returns a new *EventLog
func NewEventLog() *EventLog {

	return &EventLog{
		ModelVersion:  1,
		Level:         EventLogLevelInfo,
		MigrationsLog: map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *EventLog) Identity() elemental.Identity {

	return EventLogIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EventLog) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EventLog) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *EventLog) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesEventLog{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Category = o.Category
	s.Content = o.Content
	s.Date = o.Date
	s.Level = o.Level
	s.MigrationsLog = o.MigrationsLog
	s.Namespace = o.Namespace
	s.Opaque = o.Opaque
	s.TargetID = o.TargetID
	s.TargetIdentity = o.TargetIdentity
	s.Title = o.Title
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *EventLog) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesEventLog{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Category = s.Category
	o.Content = s.Content
	o.Date = s.Date
	o.Level = s.Level
	o.MigrationsLog = s.MigrationsLog
	o.Namespace = s.Namespace
	o.Opaque = s.Opaque
	o.TargetID = s.TargetID
	o.TargetIdentity = s.TargetIdentity
	o.Title = s.Title
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *EventLog) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *EventLog) BleveType() string {

	return "eventlog"
}

// DefaultOrder returns the list of default ordering fields.
func (o *EventLog) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *EventLog) Doc() string {

	return `Allows you to report various events on any object.`
}

func (o *EventLog) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *EventLog) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *EventLog) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *EventLog) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *EventLog) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *EventLog) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *EventLog) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *EventLog) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *EventLog) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *EventLog) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseEventLog{
			ID:             &o.ID,
			Category:       &o.Category,
			Content:        &o.Content,
			Date:           &o.Date,
			Level:          &o.Level,
			MigrationsLog:  &o.MigrationsLog,
			Namespace:      &o.Namespace,
			Opaque:         &o.Opaque,
			TargetID:       &o.TargetID,
			TargetIdentity: &o.TargetIdentity,
			Title:          &o.Title,
			ZHash:          &o.ZHash,
			Zone:           &o.Zone,
		}
	}

	sp := &SparseEventLog{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "category":
			sp.Category = &(o.Category)
		case "content":
			sp.Content = &(o.Content)
		case "date":
			sp.Date = &(o.Date)
		case "level":
			sp.Level = &(o.Level)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "opaque":
			sp.Opaque = &(o.Opaque)
		case "targetID":
			sp.TargetID = &(o.TargetID)
		case "targetIdentity":
			sp.TargetIdentity = &(o.TargetIdentity)
		case "title":
			sp.Title = &(o.Title)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseEventLog to the object.
func (o *EventLog) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseEventLog)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Category != nil {
		o.Category = *so.Category
	}
	if so.Content != nil {
		o.Content = *so.Content
	}
	if so.Date != nil {
		o.Date = *so.Date
	}
	if so.Level != nil {
		o.Level = *so.Level
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Opaque != nil {
		o.Opaque = *so.Opaque
	}
	if so.TargetID != nil {
		o.TargetID = *so.TargetID
	}
	if so.TargetIdentity != nil {
		o.TargetIdentity = *so.TargetIdentity
	}
	if so.Title != nil {
		o.Title = *so.Title
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the EventLog.
func (o *EventLog) DeepCopy() *EventLog {

	if o == nil {
		return nil
	}

	out := &EventLog{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *EventLog.
func (o *EventLog) DeepCopyInto(out *EventLog) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy EventLog: %s", err))
	}

	*out = *target.(*EventLog)
}

// Validate valides the current information stored into the structure.
func (o *EventLog) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("category", o.Category); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("content", o.Content); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("level", string(o.Level), []string{"Debug", "Info", "Warning", "Error", "Critical"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("targetID", o.TargetID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("targetIdentity", o.TargetIdentity); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("title", o.Title); err != nil {
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
func (*EventLog) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EventLogAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EventLogLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*EventLog) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EventLogAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *EventLog) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "category":
		return o.Category
	case "content":
		return o.Content
	case "date":
		return o.Date
	case "level":
		return o.Level
	case "migrationsLog":
		return o.MigrationsLog
	case "namespace":
		return o.Namespace
	case "opaque":
		return o.Opaque
	case "targetID":
		return o.TargetID
	case "targetIdentity":
		return o.TargetIdentity
	case "title":
		return o.Title
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// EventLogAttributesMap represents the map of attribute for EventLog.
var EventLogAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Category": {
		AllowedChoices: []string{},
		ConvertedName:  "Category",
		CreationOnly:   true,
		Description:    `Category of the event log.`,
		Exposed:        true,
		Name:           "category",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Content": {
		AllowedChoices: []string{},
		ConvertedName:  "Content",
		CreationOnly:   true,
		Description:    `Content of the event log.`,
		Exposed:        true,
		Name:           "content",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Date": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Date",
		CreationOnly:   true,
		Description:    `Creation date of the event log.`,
		Exposed:        true,
		Name:           "date",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"Level": {
		AllowedChoices: []string{"Debug", "Info", "Warning", "Error", "Critical"},
		ConvertedName:  "Level",
		CreationOnly:   true,
		DefaultValue:   EventLogLevelInfo,
		Description:    `Sets the log level.`,
		Exposed:        true,
		Name:           "level",
		Stored:         true,
		Type:           "enum",
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
		CreationOnly:   true,
		Description:    `Namespace tag attached to the event log.`,
		Exposed:        true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Opaque": {
		AllowedChoices: []string{},
		ConvertedName:  "Opaque",
		CreationOnly:   true,
		Description: `Opaque data that can be attached to the event log, for further machine
processing.`,
		Exposed:   true,
		Name:      "opaque",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"TargetID": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetID",
		CreationOnly:   true,
		Description: `ID of the object this event log is attached to. The object must be in the same
namespace than the event log.`,
		Exposed:  true,
		Name:     "targetID",
		Required: true,
		Stored:   true,
		Type:     "string",
	},
	"TargetIdentity": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentity",
		CreationOnly:   true,
		Description:    `Identity of the object this event log is attached to.`,
		Exposed:        true,
		Name:           "targetIdentity",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Title": {
		AllowedChoices: []string{},
		ConvertedName:  "Title",
		CreationOnly:   true,
		Description:    `Title of the event log.`,
		Exposed:        true,
		Name:           "title",
		Required:       true,
		Stored:         true,
		Type:           "string",
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
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// EventLogLowerCaseAttributesMap represents the map of attribute for EventLog.
var EventLogLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"category": {
		AllowedChoices: []string{},
		ConvertedName:  "Category",
		CreationOnly:   true,
		Description:    `Category of the event log.`,
		Exposed:        true,
		Name:           "category",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"content": {
		AllowedChoices: []string{},
		ConvertedName:  "Content",
		CreationOnly:   true,
		Description:    `Content of the event log.`,
		Exposed:        true,
		Name:           "content",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"date": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Date",
		CreationOnly:   true,
		Description:    `Creation date of the event log.`,
		Exposed:        true,
		Name:           "date",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"level": {
		AllowedChoices: []string{"Debug", "Info", "Warning", "Error", "Critical"},
		ConvertedName:  "Level",
		CreationOnly:   true,
		DefaultValue:   EventLogLevelInfo,
		Description:    `Sets the log level.`,
		Exposed:        true,
		Name:           "level",
		Stored:         true,
		Type:           "enum",
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
		CreationOnly:   true,
		Description:    `Namespace tag attached to the event log.`,
		Exposed:        true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"opaque": {
		AllowedChoices: []string{},
		ConvertedName:  "Opaque",
		CreationOnly:   true,
		Description: `Opaque data that can be attached to the event log, for further machine
processing.`,
		Exposed:   true,
		Name:      "opaque",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"targetid": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetID",
		CreationOnly:   true,
		Description: `ID of the object this event log is attached to. The object must be in the same
namespace than the event log.`,
		Exposed:  true,
		Name:     "targetID",
		Required: true,
		Stored:   true,
		Type:     "string",
	},
	"targetidentity": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentity",
		CreationOnly:   true,
		Description:    `Identity of the object this event log is attached to.`,
		Exposed:        true,
		Name:           "targetIdentity",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"title": {
		AllowedChoices: []string{},
		ConvertedName:  "Title",
		CreationOnly:   true,
		Description:    `Title of the event log.`,
		Exposed:        true,
		Name:           "title",
		Required:       true,
		Stored:         true,
		Type:           "string",
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
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseEventLogsList represents a list of SparseEventLogs
type SparseEventLogsList []*SparseEventLog

// Identity returns the identity of the objects in the list.
func (o SparseEventLogsList) Identity() elemental.Identity {

	return EventLogIdentity
}

// Copy returns a pointer to a copy the SparseEventLogsList.
func (o SparseEventLogsList) Copy() elemental.Identifiables {

	copy := append(SparseEventLogsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseEventLogsList.
func (o SparseEventLogsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseEventLogsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseEventLog))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseEventLogsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseEventLogsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseEventLogsList converted to EventLogsList.
func (o SparseEventLogsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseEventLogsList) Version() int {

	return 1
}

// SparseEventLog represents the sparse version of a eventlog.
type SparseEventLog struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Category of the event log.
	Category *string `json:"category,omitempty" msgpack:"category,omitempty" bson:"category,omitempty" mapstructure:"category,omitempty"`

	// Content of the event log.
	Content *string `json:"content,omitempty" msgpack:"content,omitempty" bson:"content,omitempty" mapstructure:"content,omitempty"`

	// Creation date of the event log.
	Date *time.Time `json:"date,omitempty" msgpack:"date,omitempty" bson:"date,omitempty" mapstructure:"date,omitempty"`

	// Sets the log level.
	Level *EventLogLevelValue `json:"level,omitempty" msgpack:"level,omitempty" bson:"level,omitempty" mapstructure:"level,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to the event log.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Opaque data that can be attached to the event log, for further machine
	// processing.
	Opaque *string `json:"opaque,omitempty" msgpack:"opaque,omitempty" bson:"opaque,omitempty" mapstructure:"opaque,omitempty"`

	// ID of the object this event log is attached to. The object must be in the same
	// namespace than the event log.
	TargetID *string `json:"targetID,omitempty" msgpack:"targetID,omitempty" bson:"targetid,omitempty" mapstructure:"targetID,omitempty"`

	// Identity of the object this event log is attached to.
	TargetIdentity *string `json:"targetIdentity,omitempty" msgpack:"targetIdentity,omitempty" bson:"targetidentity,omitempty" mapstructure:"targetIdentity,omitempty"`

	// Title of the event log.
	Title *string `json:"title,omitempty" msgpack:"title,omitempty" bson:"title,omitempty" mapstructure:"title,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone *int `json:"zone,omitempty" msgpack:"zone,omitempty" bson:"zone,omitempty" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseEventLog returns a new  SparseEventLog.
func NewSparseEventLog() *SparseEventLog {
	return &SparseEventLog{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseEventLog) Identity() elemental.Identity {

	return EventLogIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseEventLog) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseEventLog) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseEventLog) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseEventLog{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Category != nil {
		s.Category = o.Category
	}
	if o.Content != nil {
		s.Content = o.Content
	}
	if o.Date != nil {
		s.Date = o.Date
	}
	if o.Level != nil {
		s.Level = o.Level
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.Opaque != nil {
		s.Opaque = o.Opaque
	}
	if o.TargetID != nil {
		s.TargetID = o.TargetID
	}
	if o.TargetIdentity != nil {
		s.TargetIdentity = o.TargetIdentity
	}
	if o.Title != nil {
		s.Title = o.Title
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
func (o *SparseEventLog) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseEventLog{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Category != nil {
		o.Category = s.Category
	}
	if s.Content != nil {
		o.Content = s.Content
	}
	if s.Date != nil {
		o.Date = s.Date
	}
	if s.Level != nil {
		o.Level = s.Level
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.Opaque != nil {
		o.Opaque = s.Opaque
	}
	if s.TargetID != nil {
		o.TargetID = s.TargetID
	}
	if s.TargetIdentity != nil {
		o.TargetIdentity = s.TargetIdentity
	}
	if s.Title != nil {
		o.Title = s.Title
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
func (o *SparseEventLog) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseEventLog) ToPlain() elemental.PlainIdentifiable {

	out := NewEventLog()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Category != nil {
		out.Category = *o.Category
	}
	if o.Content != nil {
		out.Content = *o.Content
	}
	if o.Date != nil {
		out.Date = *o.Date
	}
	if o.Level != nil {
		out.Level = *o.Level
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Opaque != nil {
		out.Opaque = *o.Opaque
	}
	if o.TargetID != nil {
		out.TargetID = *o.TargetID
	}
	if o.TargetIdentity != nil {
		out.TargetIdentity = *o.TargetIdentity
	}
	if o.Title != nil {
		out.Title = *o.Title
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
func (o *SparseEventLog) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseEventLog) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseEventLog) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseEventLog) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseEventLog) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseEventLog) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseEventLog) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseEventLog) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseEventLog.
func (o *SparseEventLog) DeepCopy() *SparseEventLog {

	if o == nil {
		return nil
	}

	out := &SparseEventLog{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseEventLog.
func (o *SparseEventLog) DeepCopyInto(out *SparseEventLog) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseEventLog: %s", err))
	}

	*out = *target.(*SparseEventLog)
}

type mongoAttributesEventLog struct {
	ID             bson.ObjectId      `bson:"_id,omitempty"`
	Category       string             `bson:"category"`
	Content        string             `bson:"content"`
	Date           time.Time          `bson:"date"`
	Level          EventLogLevelValue `bson:"level"`
	MigrationsLog  map[string]string  `bson:"migrationslog,omitempty"`
	Namespace      string             `bson:"namespace"`
	Opaque         string             `bson:"opaque"`
	TargetID       string             `bson:"targetid"`
	TargetIdentity string             `bson:"targetidentity"`
	Title          string             `bson:"title"`
	ZHash          int                `bson:"zhash"`
	Zone           int                `bson:"zone"`
}
type mongoAttributesSparseEventLog struct {
	ID             bson.ObjectId       `bson:"_id,omitempty"`
	Category       *string             `bson:"category,omitempty"`
	Content        *string             `bson:"content,omitempty"`
	Date           *time.Time          `bson:"date,omitempty"`
	Level          *EventLogLevelValue `bson:"level,omitempty"`
	MigrationsLog  *map[string]string  `bson:"migrationslog,omitempty"`
	Namespace      *string             `bson:"namespace,omitempty"`
	Opaque         *string             `bson:"opaque,omitempty"`
	TargetID       *string             `bson:"targetid,omitempty"`
	TargetIdentity *string             `bson:"targetidentity,omitempty"`
	Title          *string             `bson:"title,omitempty"`
	ZHash          *int                `bson:"zhash,omitempty"`
	Zone           *int                `bson:"zone,omitempty"`
}
