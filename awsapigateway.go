package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AWSAPIGatewayIdentity represents the Identity of the object.
var AWSAPIGatewayIdentity = elemental.Identity{
	Name:     "awsapigateway",
	Category: "awsapigateways",
	Package:  "goldrush",
	Private:  false,
}

// AWSAPIGatewaysList represents a list of AWSAPIGateways
type AWSAPIGatewaysList []*AWSAPIGateway

// Identity returns the identity of the objects in the list.
func (o AWSAPIGatewaysList) Identity() elemental.Identity {

	return AWSAPIGatewayIdentity
}

// Copy returns a pointer to a copy the AWSAPIGatewaysList.
func (o AWSAPIGatewaysList) Copy() elemental.Identifiables {

	copy := append(AWSAPIGatewaysList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AWSAPIGatewaysList.
func (o AWSAPIGatewaysList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AWSAPIGatewaysList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AWSAPIGateway))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AWSAPIGatewaysList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AWSAPIGatewaysList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the AWSAPIGatewaysList converted to SparseAWSAPIGatewaysList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AWSAPIGatewaysList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAWSAPIGatewaysList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAWSAPIGateway)
	}

	return out
}

// Version returns the version of the content.
func (o AWSAPIGatewaysList) Version() int {

	return 1
}

// AWSAPIGateway represents the model of a awsapigateway
type AWSAPIGateway struct {
	// API ID as defined on AWS for the API that handled this request.
	APIID string `json:"APIID" msgpack:"APIID" bson:"-" mapstructure:"APIID,omitempty"`

	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// the account ID for the gateway managing this request.
	AccountID string `json:"accountID" msgpack:"accountID" bson:"-" mapstructure:"accountID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// The policy decision for this API flow.
	Authorized bool `json:"authorized" msgpack:"authorized" bson:"-" mapstructure:"authorized,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// API method that handled this request.
	Method string `json:"method" msgpack:"method" bson:"-" mapstructure:"method,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Link to the cluster namespace where the AWS API gateway is defined.
	NamespaceID string `json:"namespaceID" msgpack:"namespaceID" bson:"-" mapstructure:"namespaceID,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// API resource that handled this request.
	Resource string `json:"resource" msgpack:"resource" bson:"-" mapstructure:"resource,omitempty"`

	// the client ip for this request.
	SourceIP string `json:"sourceIP" msgpack:"sourceIP" bson:"-" mapstructure:"sourceIP,omitempty"`

	// the stage name as defined on AWS for the API that handled this request.
	Stage string `json:"stage" msgpack:"stage" bson:"-" mapstructure:"stage,omitempty"`

	// the JWT token that was optionally attached to this request.
	Token string `json:"token" msgpack:"token" bson:"-" mapstructure:"token,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAWSAPIGateway returns a new *AWSAPIGateway
func NewAWSAPIGateway() *AWSAPIGateway {

	return &AWSAPIGateway{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		NormalizedTags: []string{},
		Metadata:       []string{},
		MigrationsLog:  map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *AWSAPIGateway) Identity() elemental.Identity {

	return AWSAPIGatewayIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AWSAPIGateway) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AWSAPIGateway) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AWSAPIGateway) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAWSAPIGateway{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.Description = o.Description
	s.Metadata = o.Metadata
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.NormalizedTags = o.NormalizedTags
	s.Protected = o.Protected
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AWSAPIGateway) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAWSAPIGateway{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.Description = s.Description
	o.Metadata = s.Metadata
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.NormalizedTags = s.NormalizedTags
	o.Protected = s.Protected
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *AWSAPIGateway) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *AWSAPIGateway) BleveType() string {

	return "awsapigateway"
}

// DefaultOrder returns the list of default ordering fields.
func (o *AWSAPIGateway) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *AWSAPIGateway) Doc() string {

	return `managed API decisions for the AWS API Gateway.`
}

func (o *AWSAPIGateway) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *AWSAPIGateway) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *AWSAPIGateway) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *AWSAPIGateway) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *AWSAPIGateway) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *AWSAPIGateway) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *AWSAPIGateway) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *AWSAPIGateway) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *AWSAPIGateway) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *AWSAPIGateway) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *AWSAPIGateway) SetDescription(description string) {

	o.Description = description
}

// GetMetadata returns the Metadata of the receiver.
func (o *AWSAPIGateway) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *AWSAPIGateway) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *AWSAPIGateway) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *AWSAPIGateway) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *AWSAPIGateway) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *AWSAPIGateway) SetName(name string) {

	o.Name = name
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *AWSAPIGateway) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *AWSAPIGateway) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *AWSAPIGateway) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *AWSAPIGateway) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *AWSAPIGateway) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *AWSAPIGateway) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *AWSAPIGateway) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *AWSAPIGateway) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *AWSAPIGateway) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *AWSAPIGateway) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *AWSAPIGateway) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *AWSAPIGateway) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *AWSAPIGateway) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAWSAPIGateway{
			APIID:                &o.APIID,
			ID:                   &o.ID,
			AccountID:            &o.AccountID,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			Authorized:           &o.Authorized,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Description:          &o.Description,
			Metadata:             &o.Metadata,
			Method:               &o.Method,
			MigrationsLog:        &o.MigrationsLog,
			Name:                 &o.Name,
			NamespaceID:          &o.NamespaceID,
			NormalizedTags:       &o.NormalizedTags,
			Protected:            &o.Protected,
			Resource:             &o.Resource,
			SourceIP:             &o.SourceIP,
			Stage:                &o.Stage,
			Token:                &o.Token,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseAWSAPIGateway{}
	for _, f := range fields {
		switch f {
		case "APIID":
			sp.APIID = &(o.APIID)
		case "ID":
			sp.ID = &(o.ID)
		case "accountID":
			sp.AccountID = &(o.AccountID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "authorized":
			sp.Authorized = &(o.Authorized)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "method":
			sp.Method = &(o.Method)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespaceID":
			sp.NamespaceID = &(o.NamespaceID)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "protected":
			sp.Protected = &(o.Protected)
		case "resource":
			sp.Resource = &(o.Resource)
		case "sourceIP":
			sp.SourceIP = &(o.SourceIP)
		case "stage":
			sp.Stage = &(o.Stage)
		case "token":
			sp.Token = &(o.Token)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
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

// Patch apply the non nil value of a *SparseAWSAPIGateway to the object.
func (o *AWSAPIGateway) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAWSAPIGateway)
	if so.APIID != nil {
		o.APIID = *so.APIID
	}
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.AccountID != nil {
		o.AccountID = *so.AccountID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.Authorized != nil {
		o.Authorized = *so.Authorized
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.Method != nil {
		o.Method = *so.Method
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.NamespaceID != nil {
		o.NamespaceID = *so.NamespaceID
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Resource != nil {
		o.Resource = *so.Resource
	}
	if so.SourceIP != nil {
		o.SourceIP = *so.SourceIP
	}
	if so.Stage != nil {
		o.Stage = *so.Stage
	}
	if so.Token != nil {
		o.Token = *so.Token
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
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

// DeepCopy returns a deep copy if the AWSAPIGateway.
func (o *AWSAPIGateway) DeepCopy() *AWSAPIGateway {

	if o == nil {
		return nil
	}

	out := &AWSAPIGateway{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AWSAPIGateway.
func (o *AWSAPIGateway) DeepCopyInto(out *AWSAPIGateway) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AWSAPIGateway: %s", err))
	}

	*out = *target.(*AWSAPIGateway)
}

// Validate valides the current information stored into the structure.
func (o *AWSAPIGateway) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateMetadata("metadata", o.Metadata); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
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
func (*AWSAPIGateway) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AWSAPIGatewayAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AWSAPIGatewayLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AWSAPIGateway) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AWSAPIGatewayAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AWSAPIGateway) ValueForAttribute(name string) interface{} {

	switch name {
	case "APIID":
		return o.APIID
	case "ID":
		return o.ID
	case "accountID":
		return o.AccountID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "authorized":
		return o.Authorized
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "metadata":
		return o.Metadata
	case "method":
		return o.Method
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespaceID":
		return o.NamespaceID
	case "normalizedTags":
		return o.NormalizedTags
	case "protected":
		return o.Protected
	case "resource":
		return o.Resource
	case "sourceIP":
		return o.SourceIP
	case "stage":
		return o.Stage
	case "token":
		return o.Token
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// AWSAPIGatewayAttributesMap represents the map of attribute for AWSAPIGateway.
var AWSAPIGatewayAttributesMap = map[string]elemental.AttributeSpecification{
	"APIID": {
		AllowedChoices: []string{},
		ConvertedName:  "APIID",
		Description:    `API ID as defined on AWS for the API that handled this request.`,
		Exposed:        true,
		Name:           "APIID",
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
	"AccountID": {
		AllowedChoices: []string{},
		ConvertedName:  "AccountID",
		Description:    `the account ID for the gateway managing this request.`,
		Exposed:        true,
		Name:           "accountID",
		Type:           "string",
	},
	"Annotations": {
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"AssociatedTags": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Authorized": {
		AllowedChoices: []string{},
		ConvertedName:  "Authorized",
		Description:    `The policy decision for this API flow.`,
		Exposed:        true,
		Name:           "authorized",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "boolean",
	},
	"CreateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
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
	"Description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Metadata": {
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
	},
	"Method": {
		AllowedChoices: []string{},
		ConvertedName:  "Method",
		Description:    `API method that handled this request.`,
		Exposed:        true,
		Name:           "method",
		Type:           "string",
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
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NamespaceID": {
		AllowedChoices: []string{},
		ConvertedName:  "NamespaceID",
		Description:    `Link to the cluster namespace where the AWS API gateway is defined.`,
		Exposed:        true,
		Name:           "namespaceID",
		Type:           "string",
	},
	"NormalizedTags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"Protected": {
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Resource": {
		AllowedChoices: []string{},
		ConvertedName:  "Resource",
		Description:    `API resource that handled this request.`,
		Exposed:        true,
		Name:           "resource",
		Type:           "string",
	},
	"SourceIP": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `the client ip for this request.`,
		Exposed:        true,
		Name:           "sourceIP",
		Type:           "string",
	},
	"Stage": {
		AllowedChoices: []string{},
		ConvertedName:  "Stage",
		Description:    `the stage name as defined on AWS for the API that handled this request.`,
		Exposed:        true,
		Name:           "stage",
		Type:           "string",
	},
	"Token": {
		AllowedChoices: []string{},
		ConvertedName:  "Token",
		Description:    `the JWT token that was optionally attached to this request.`,
		Exposed:        true,
		Name:           "token",
		Type:           "string",
	},
	"UpdateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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

// AWSAPIGatewayLowerCaseAttributesMap represents the map of attribute for AWSAPIGateway.
var AWSAPIGatewayLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"apiid": {
		AllowedChoices: []string{},
		ConvertedName:  "APIID",
		Description:    `API ID as defined on AWS for the API that handled this request.`,
		Exposed:        true,
		Name:           "APIID",
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
	"accountid": {
		AllowedChoices: []string{},
		ConvertedName:  "AccountID",
		Description:    `the account ID for the gateway managing this request.`,
		Exposed:        true,
		Name:           "accountID",
		Type:           "string",
	},
	"annotations": {
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"associatedtags": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"authorized": {
		AllowedChoices: []string{},
		ConvertedName:  "Authorized",
		Description:    `The policy decision for this API flow.`,
		Exposed:        true,
		Name:           "authorized",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "boolean",
	},
	"createidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
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
	"description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"metadata": {
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Contains tags that can only be set during creation, must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
	},
	"method": {
		AllowedChoices: []string{},
		ConvertedName:  "Method",
		Description:    `API method that handled this request.`,
		Exposed:        true,
		Name:           "method",
		Type:           "string",
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
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"namespaceid": {
		AllowedChoices: []string{},
		ConvertedName:  "NamespaceID",
		Description:    `Link to the cluster namespace where the AWS API gateway is defined.`,
		Exposed:        true,
		Name:           "namespaceID",
		Type:           "string",
	},
	"normalizedtags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"protected": {
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"resource": {
		AllowedChoices: []string{},
		ConvertedName:  "Resource",
		Description:    `API resource that handled this request.`,
		Exposed:        true,
		Name:           "resource",
		Type:           "string",
	},
	"sourceip": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `the client ip for this request.`,
		Exposed:        true,
		Name:           "sourceIP",
		Type:           "string",
	},
	"stage": {
		AllowedChoices: []string{},
		ConvertedName:  "Stage",
		Description:    `the stage name as defined on AWS for the API that handled this request.`,
		Exposed:        true,
		Name:           "stage",
		Type:           "string",
	},
	"token": {
		AllowedChoices: []string{},
		ConvertedName:  "Token",
		Description:    `the JWT token that was optionally attached to this request.`,
		Exposed:        true,
		Name:           "token",
		Type:           "string",
	},
	"updateidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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

// SparseAWSAPIGatewaysList represents a list of SparseAWSAPIGateways
type SparseAWSAPIGatewaysList []*SparseAWSAPIGateway

// Identity returns the identity of the objects in the list.
func (o SparseAWSAPIGatewaysList) Identity() elemental.Identity {

	return AWSAPIGatewayIdentity
}

// Copy returns a pointer to a copy the SparseAWSAPIGatewaysList.
func (o SparseAWSAPIGatewaysList) Copy() elemental.Identifiables {

	copy := append(SparseAWSAPIGatewaysList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAWSAPIGatewaysList.
func (o SparseAWSAPIGatewaysList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAWSAPIGatewaysList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAWSAPIGateway))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAWSAPIGatewaysList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAWSAPIGatewaysList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseAWSAPIGatewaysList converted to AWSAPIGatewaysList.
func (o SparseAWSAPIGatewaysList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAWSAPIGatewaysList) Version() int {

	return 1
}

// SparseAWSAPIGateway represents the sparse version of a awsapigateway.
type SparseAWSAPIGateway struct {
	// API ID as defined on AWS for the API that handled this request.
	APIID *string `json:"APIID,omitempty" msgpack:"APIID,omitempty" bson:"-" mapstructure:"APIID,omitempty"`

	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// the account ID for the gateway managing this request.
	AccountID *string `json:"accountID,omitempty" msgpack:"accountID,omitempty" bson:"-" mapstructure:"accountID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// The policy decision for this API flow.
	Authorized *bool `json:"authorized,omitempty" msgpack:"authorized,omitempty" bson:"-" mapstructure:"authorized,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// API method that handled this request.
	Method *string `json:"method,omitempty" msgpack:"method,omitempty" bson:"-" mapstructure:"method,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Link to the cluster namespace where the AWS API gateway is defined.
	NamespaceID *string `json:"namespaceID,omitempty" msgpack:"namespaceID,omitempty" bson:"-" mapstructure:"namespaceID,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// API resource that handled this request.
	Resource *string `json:"resource,omitempty" msgpack:"resource,omitempty" bson:"-" mapstructure:"resource,omitempty"`

	// the client ip for this request.
	SourceIP *string `json:"sourceIP,omitempty" msgpack:"sourceIP,omitempty" bson:"-" mapstructure:"sourceIP,omitempty"`

	// the stage name as defined on AWS for the API that handled this request.
	Stage *string `json:"stage,omitempty" msgpack:"stage,omitempty" bson:"-" mapstructure:"stage,omitempty"`

	// the JWT token that was optionally attached to this request.
	Token *string `json:"token,omitempty" msgpack:"token,omitempty" bson:"-" mapstructure:"token,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAWSAPIGateway returns a new  SparseAWSAPIGateway.
func NewSparseAWSAPIGateway() *SparseAWSAPIGateway {
	return &SparseAWSAPIGateway{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAWSAPIGateway) Identity() elemental.Identity {

	return AWSAPIGatewayIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAWSAPIGateway) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAWSAPIGateway) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAWSAPIGateway) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseAWSAPIGateway{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Metadata != nil {
		s.Metadata = o.Metadata
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.UpdateIdempotencyKey != nil {
		s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
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
func (o *SparseAWSAPIGateway) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseAWSAPIGateway{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Metadata != nil {
		o.Metadata = s.Metadata
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
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
func (o *SparseAWSAPIGateway) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAWSAPIGateway) ToPlain() elemental.PlainIdentifiable {

	out := NewAWSAPIGateway()
	if o.APIID != nil {
		out.APIID = *o.APIID
	}
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.AccountID != nil {
		out.AccountID = *o.AccountID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.Authorized != nil {
		out.Authorized = *o.Authorized
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Method != nil {
		out.Method = *o.Method
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.NamespaceID != nil {
		out.NamespaceID = *o.NamespaceID
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Resource != nil {
		out.Resource = *o.Resource
	}
	if o.SourceIP != nil {
		out.SourceIP = *o.SourceIP
	}
	if o.Stage != nil {
		out.Stage = *o.Stage
	}
	if o.Token != nil {
		out.Token = *o.Token
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
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

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseAWSAPIGateway) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseAWSAPIGateway) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseAWSAPIGateway) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseAWSAPIGateway) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseAWSAPIGateway) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetDescription(description string) {

	o.Description = &description
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseAWSAPIGateway) GetMetadata() (out []string) {

	if o.Metadata == nil {
		return
	}

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseAWSAPIGateway) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseAWSAPIGateway) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetName(name string) {

	o.Name = &name
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseAWSAPIGateway) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseAWSAPIGateway) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseAWSAPIGateway) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseAWSAPIGateway) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseAWSAPIGateway) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseAWSAPIGateway) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseAWSAPIGateway) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseAWSAPIGateway.
func (o *SparseAWSAPIGateway) DeepCopy() *SparseAWSAPIGateway {

	if o == nil {
		return nil
	}

	out := &SparseAWSAPIGateway{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAWSAPIGateway.
func (o *SparseAWSAPIGateway) DeepCopyInto(out *SparseAWSAPIGateway) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAWSAPIGateway: %s", err))
	}

	*out = *target.(*SparseAWSAPIGateway)
}

type mongoAttributesAWSAPIGateway struct {
	ID                   bson.ObjectId       `bson:"_id,omitempty"`
	Annotations          map[string][]string `bson:"annotations"`
	AssociatedTags       []string            `bson:"associatedtags"`
	CreateIdempotencyKey string              `bson:"createidempotencykey"`
	CreateTime           time.Time           `bson:"createtime"`
	Description          string              `bson:"description"`
	Metadata             []string            `bson:"metadata"`
	MigrationsLog        map[string]string   `bson:"migrationslog,omitempty"`
	Name                 string              `bson:"name"`
	NormalizedTags       []string            `bson:"normalizedtags"`
	Protected            bool                `bson:"protected"`
	UpdateIdempotencyKey string              `bson:"updateidempotencykey"`
	UpdateTime           time.Time           `bson:"updatetime"`
	ZHash                int                 `bson:"zhash"`
	Zone                 int                 `bson:"zone"`
}
type mongoAttributesSparseAWSAPIGateway struct {
	ID                   bson.ObjectId        `bson:"_id,omitempty"`
	Annotations          *map[string][]string `bson:"annotations,omitempty"`
	AssociatedTags       *[]string            `bson:"associatedtags,omitempty"`
	CreateIdempotencyKey *string              `bson:"createidempotencykey,omitempty"`
	CreateTime           *time.Time           `bson:"createtime,omitempty"`
	Description          *string              `bson:"description,omitempty"`
	Metadata             *[]string            `bson:"metadata,omitempty"`
	MigrationsLog        *map[string]string   `bson:"migrationslog,omitempty"`
	Name                 *string              `bson:"name,omitempty"`
	NormalizedTags       *[]string            `bson:"normalizedtags,omitempty"`
	Protected            *bool                `bson:"protected,omitempty"`
	UpdateIdempotencyKey *string              `bson:"updateidempotencykey,omitempty"`
	UpdateTime           *time.Time           `bson:"updatetime,omitempty"`
	ZHash                *int                 `bson:"zhash,omitempty"`
	Zone                 *int                 `bson:"zone,omitempty"`
}
