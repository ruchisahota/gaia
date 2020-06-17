package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ImportRequestStatusValue represents the possible values for attribute "status".
type ImportRequestStatusValue string

const (
	// ImportRequestStatusApproved represents the value Approved.
	ImportRequestStatusApproved ImportRequestStatusValue = "Approved"

	// ImportRequestStatusDraft represents the value Draft.
	ImportRequestStatusDraft ImportRequestStatusValue = "Draft"

	// ImportRequestStatusRejected represents the value Rejected.
	ImportRequestStatusRejected ImportRequestStatusValue = "Rejected"

	// ImportRequestStatusSubmitted represents the value Submitted.
	ImportRequestStatusSubmitted ImportRequestStatusValue = "Submitted"
)

// ImportRequestIdentity represents the Identity of the object.
var ImportRequestIdentity = elemental.Identity{
	Name:     "importrequest",
	Category: "importrequests",
	Package:  "vivi",
	Private:  false,
}

// ImportRequestsList represents a list of ImportRequests
type ImportRequestsList []*ImportRequest

// Identity returns the identity of the objects in the list.
func (o ImportRequestsList) Identity() elemental.Identity {

	return ImportRequestIdentity
}

// Copy returns a pointer to a copy the ImportRequestsList.
func (o ImportRequestsList) Copy() elemental.Identifiables {

	copy := append(ImportRequestsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ImportRequestsList.
func (o ImportRequestsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ImportRequestsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ImportRequest))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ImportRequestsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ImportRequestsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ImportRequestsList converted to SparseImportRequestsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ImportRequestsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseImportRequestsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseImportRequest)
	}

	return out
}

// Version returns the version of the content.
func (o ImportRequestsList) Version() int {

	return 1
}

// ImportRequest represents the model of a importrequest
type ImportRequest struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// A new comment that will be added to `commentFeed`.
	Comment string `json:"comment" msgpack:"comment" bson:"-" mapstructure:"comment,omitempty"`

	// List of comments that have been added to that request.
	CommentFeed []*Comment `json:"commentFeed" msgpack:"commentFeed" bson:"commentfeed" mapstructure:"commentFeed,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Data to import.
	Data map[string][]map[string]interface{} `json:"data" msgpack:"data" bson:"data" mapstructure:"data,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// The identity claims of the requester; populated by Segment Console.
	RequesterClaims []string `json:"requesterClaims" msgpack:"requesterClaims" bson:"requesterclaims" mapstructure:"requesterClaims,omitempty"`

	// The namespace from which the request originated; populated by Segment Console.
	RequesterNamespace string `json:"requesterNamespace" msgpack:"requesterNamespace" bson:"requesternamespace" mapstructure:"requesterNamespace,omitempty"`

	// Allows the content to be changed. `Submitted`: the request moves to the target
	// namespace
	// for approval. `Approved`: the data will be created immediately. `Rejected`: the
	// request
	// cannot be changed anymore and can be deleted.
	Status ImportRequestStatusValue `json:"status" msgpack:"status" bson:"status" mapstructure:"status,omitempty"`

	// Internal field to know if the request has been submitted once.
	SubmittedOnce bool `json:"-" msgpack:"-" bson:"submittedonce" mapstructure:"-,omitempty"`

	// The namespace where the request will be sent. The requester can set any
	// namespace but
	// needs to have an authorization to post the request in that namespace.
	TargetNamespace string `json:"targetNamespace" msgpack:"targetNamespace" bson:"targetnamespace" mapstructure:"targetNamespace,omitempty"`

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

// NewImportRequest returns a new *ImportRequest
func NewImportRequest() *ImportRequest {

	return &ImportRequest{
		ModelVersion:    1,
		Annotations:     map[string][]string{},
		AssociatedTags:  []string{},
		CommentFeed:     []*Comment{},
		Data:            map[string][]map[string]interface{}{},
		NormalizedTags:  []string{},
		Status:          ImportRequestStatusDraft,
		RequesterClaims: []string{},
		MigrationsLog:   map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *ImportRequest) Identity() elemental.Identity {

	return ImportRequestIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ImportRequest) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ImportRequest) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ImportRequest) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesImportRequest{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CommentFeed = o.CommentFeed
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.Data = o.Data
	s.Description = o.Description
	s.MigrationsLog = o.MigrationsLog
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.Protected = o.Protected
	s.RequesterClaims = o.RequesterClaims
	s.RequesterNamespace = o.RequesterNamespace
	s.Status = o.Status
	s.SubmittedOnce = o.SubmittedOnce
	s.TargetNamespace = o.TargetNamespace
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ImportRequest) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesImportRequest{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CommentFeed = s.CommentFeed
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.Data = s.Data
	o.Description = s.Description
	o.MigrationsLog = s.MigrationsLog
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.Protected = s.Protected
	o.RequesterClaims = s.RequesterClaims
	o.RequesterNamespace = s.RequesterNamespace
	o.Status = s.Status
	o.SubmittedOnce = s.SubmittedOnce
	o.TargetNamespace = s.TargetNamespace
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *ImportRequest) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *ImportRequest) BleveType() string {

	return "importrequest"
}

// DefaultOrder returns the list of default ordering fields.
func (o *ImportRequest) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *ImportRequest) Doc() string {

	return `Allows you to send an import request to create objects to a namespace where
the requester doesn't normally have the permission to do so (other than creating
import requests).

The requester must have the permission to create the request in their namespace
and the target namespace.

When the request is created, the status is set to ` + "`" + `Draft` + "`" + `. The requester can
edit the content as much as desired.
When ready to send the request, update the status to ` + "`" + `Submitted` + "`" + `.
The request will then be moved to the target namespace.
At that point nobody can edit the content of the requests other than adding
comments.

The requestee will now see the request, and will either

- Set the status as ` + "`" + `Approved` + "`" + `. This will create the objects in the target
  namespace.

- Set the status as ` + "`" + `Rejected` + "`" + `. The request cannot be edited anymore and can be
  deleted.

- Set the status back as ` + "`" + `Draft` + "`" + `. The request will go back to the requester
  namespace so that the requester can make changes. Once the change are ready,
the requester
  will set back the status as ` + "`" + `Submitted` + "`" + `.

The ` + "`" + `data` + "`" + ` format is the same as ` + "`" + `Export` + "`" + `.`
}

func (o *ImportRequest) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *ImportRequest) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *ImportRequest) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *ImportRequest) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *ImportRequest) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *ImportRequest) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *ImportRequest) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *ImportRequest) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *ImportRequest) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *ImportRequest) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *ImportRequest) SetDescription(description string) {

	o.Description = description
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *ImportRequest) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *ImportRequest) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *ImportRequest) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *ImportRequest) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *ImportRequest) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *ImportRequest) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *ImportRequest) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *ImportRequest) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *ImportRequest) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *ImportRequest) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *ImportRequest) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *ImportRequest) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *ImportRequest) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *ImportRequest) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *ImportRequest) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *ImportRequest) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *ImportRequest) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseImportRequest{
			ID:                   &o.ID,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			Comment:              &o.Comment,
			CommentFeed:          &o.CommentFeed,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Data:                 &o.Data,
			Description:          &o.Description,
			MigrationsLog:        &o.MigrationsLog,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Protected:            &o.Protected,
			RequesterClaims:      &o.RequesterClaims,
			RequesterNamespace:   &o.RequesterNamespace,
			Status:               &o.Status,
			SubmittedOnce:        &o.SubmittedOnce,
			TargetNamespace:      &o.TargetNamespace,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseImportRequest{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "comment":
			sp.Comment = &(o.Comment)
		case "commentFeed":
			sp.CommentFeed = &(o.CommentFeed)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "data":
			sp.Data = &(o.Data)
		case "description":
			sp.Description = &(o.Description)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "protected":
			sp.Protected = &(o.Protected)
		case "requesterClaims":
			sp.RequesterClaims = &(o.RequesterClaims)
		case "requesterNamespace":
			sp.RequesterNamespace = &(o.RequesterNamespace)
		case "status":
			sp.Status = &(o.Status)
		case "submittedOnce":
			sp.SubmittedOnce = &(o.SubmittedOnce)
		case "targetNamespace":
			sp.TargetNamespace = &(o.TargetNamespace)
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

// Patch apply the non nil value of a *SparseImportRequest to the object.
func (o *ImportRequest) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseImportRequest)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.Comment != nil {
		o.Comment = *so.Comment
	}
	if so.CommentFeed != nil {
		o.CommentFeed = *so.CommentFeed
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Data != nil {
		o.Data = *so.Data
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.RequesterClaims != nil {
		o.RequesterClaims = *so.RequesterClaims
	}
	if so.RequesterNamespace != nil {
		o.RequesterNamespace = *so.RequesterNamespace
	}
	if so.Status != nil {
		o.Status = *so.Status
	}
	if so.SubmittedOnce != nil {
		o.SubmittedOnce = *so.SubmittedOnce
	}
	if so.TargetNamespace != nil {
		o.TargetNamespace = *so.TargetNamespace
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

// DeepCopy returns a deep copy if the ImportRequest.
func (o *ImportRequest) DeepCopy() *ImportRequest {

	if o == nil {
		return nil
	}

	out := &ImportRequest{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ImportRequest.
func (o *ImportRequest) DeepCopyInto(out *ImportRequest) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ImportRequest: %s", err))
	}

	*out = *target.(*ImportRequest)
}

// Validate valides the current information stored into the structure.
func (o *ImportRequest) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredExternal("data", o.Data); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Draft", "Submitted", "Approved", "Rejected"}, false); err != nil {
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
func (*ImportRequest) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ImportRequestAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ImportRequestLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ImportRequest) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ImportRequestAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *ImportRequest) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "comment":
		return o.Comment
	case "commentFeed":
		return o.CommentFeed
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "data":
		return o.Data
	case "description":
		return o.Description
	case "migrationsLog":
		return o.MigrationsLog
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "protected":
		return o.Protected
	case "requesterClaims":
		return o.RequesterClaims
	case "requesterNamespace":
		return o.RequesterNamespace
	case "status":
		return o.Status
	case "submittedOnce":
		return o.SubmittedOnce
	case "targetNamespace":
		return o.TargetNamespace
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

// ImportRequestAttributesMap represents the map of attribute for ImportRequest.
var ImportRequestAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Comment": {
		AllowedChoices: []string{},
		ConvertedName:  "Comment",
		Description:    `A new comment that will be added to ` + "`" + `commentFeed` + "`" + `.`,
		Exposed:        true,
		Name:           "comment",
		Transient:      true,
		Type:           "string",
	},
	"CommentFeed": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CommentFeed",
		Description:    `List of comments that have been added to that request.`,
		Exposed:        true,
		Name:           "commentFeed",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "comment",
		Transient:      true,
		Type:           "refList",
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
	"Data": {
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `Data to import.`,
		Exposed:        true,
		Name:           "data",
		Required:       true,
		Stored:         true,
		SubType:        "map[string][]map[string]interface{}",
		Type:           "external",
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
	"RequesterClaims": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "RequesterClaims",
		Description:    `The identity claims of the requester; populated by Segment Console.`,
		Exposed:        true,
		Name:           "requesterClaims",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"RequesterNamespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "RequesterNamespace",
		Description:    `The namespace from which the request originated; populated by Segment Console.`,
		Exposed:        true,
		Name:           "requesterNamespace",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Status": {
		AllowedChoices: []string{"Draft", "Submitted", "Approved", "Rejected"},
		ConvertedName:  "Status",
		DefaultValue:   ImportRequestStatusDraft,
		Description: `Allows the content to be changed. ` + "`" + `Submitted` + "`" + `: the request moves to the target
namespace
for approval. ` + "`" + `Approved` + "`" + `: the data will be created immediately. ` + "`" + `Rejected` + "`" + `: the
request
cannot be changed anymore and can be deleted.`,
		Exposed: true,
		Name:    "status",
		Stored:  true,
		Type:    "enum",
	},
	"SubmittedOnce": {
		AllowedChoices: []string{},
		ConvertedName:  "SubmittedOnce",
		Description:    `Internal field to know if the request has been submitted once.`,
		Name:           "submittedOnce",
		Stored:         true,
		Type:           "boolean",
	},
	"TargetNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetNamespace",
		CreationOnly:   true,
		Description: `The namespace where the request will be sent. The requester can set any
namespace but
needs to have an authorization to post the request in that namespace.`,
		Exposed:   true,
		Name:      "targetNamespace",
		Required:  true,
		Stored:    true,
		Transient: true,
		Type:      "string",
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

// ImportRequestLowerCaseAttributesMap represents the map of attribute for ImportRequest.
var ImportRequestLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"comment": {
		AllowedChoices: []string{},
		ConvertedName:  "Comment",
		Description:    `A new comment that will be added to ` + "`" + `commentFeed` + "`" + `.`,
		Exposed:        true,
		Name:           "comment",
		Transient:      true,
		Type:           "string",
	},
	"commentfeed": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CommentFeed",
		Description:    `List of comments that have been added to that request.`,
		Exposed:        true,
		Name:           "commentFeed",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "comment",
		Transient:      true,
		Type:           "refList",
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
	"data": {
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `Data to import.`,
		Exposed:        true,
		Name:           "data",
		Required:       true,
		Stored:         true,
		SubType:        "map[string][]map[string]interface{}",
		Type:           "external",
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
	"requesterclaims": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "RequesterClaims",
		Description:    `The identity claims of the requester; populated by Segment Console.`,
		Exposed:        true,
		Name:           "requesterClaims",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"requesternamespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "RequesterNamespace",
		Description:    `The namespace from which the request originated; populated by Segment Console.`,
		Exposed:        true,
		Name:           "requesterNamespace",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"status": {
		AllowedChoices: []string{"Draft", "Submitted", "Approved", "Rejected"},
		ConvertedName:  "Status",
		DefaultValue:   ImportRequestStatusDraft,
		Description: `Allows the content to be changed. ` + "`" + `Submitted` + "`" + `: the request moves to the target
namespace
for approval. ` + "`" + `Approved` + "`" + `: the data will be created immediately. ` + "`" + `Rejected` + "`" + `: the
request
cannot be changed anymore and can be deleted.`,
		Exposed: true,
		Name:    "status",
		Stored:  true,
		Type:    "enum",
	},
	"submittedonce": {
		AllowedChoices: []string{},
		ConvertedName:  "SubmittedOnce",
		Description:    `Internal field to know if the request has been submitted once.`,
		Name:           "submittedOnce",
		Stored:         true,
		Type:           "boolean",
	},
	"targetnamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "TargetNamespace",
		CreationOnly:   true,
		Description: `The namespace where the request will be sent. The requester can set any
namespace but
needs to have an authorization to post the request in that namespace.`,
		Exposed:   true,
		Name:      "targetNamespace",
		Required:  true,
		Stored:    true,
		Transient: true,
		Type:      "string",
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

// SparseImportRequestsList represents a list of SparseImportRequests
type SparseImportRequestsList []*SparseImportRequest

// Identity returns the identity of the objects in the list.
func (o SparseImportRequestsList) Identity() elemental.Identity {

	return ImportRequestIdentity
}

// Copy returns a pointer to a copy the SparseImportRequestsList.
func (o SparseImportRequestsList) Copy() elemental.Identifiables {

	copy := append(SparseImportRequestsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseImportRequestsList.
func (o SparseImportRequestsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseImportRequestsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseImportRequest))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseImportRequestsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseImportRequestsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseImportRequestsList converted to ImportRequestsList.
func (o SparseImportRequestsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseImportRequestsList) Version() int {

	return 1
}

// SparseImportRequest represents the sparse version of a importrequest.
type SparseImportRequest struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// A new comment that will be added to `commentFeed`.
	Comment *string `json:"comment,omitempty" msgpack:"comment,omitempty" bson:"-" mapstructure:"comment,omitempty"`

	// List of comments that have been added to that request.
	CommentFeed *[]*Comment `json:"commentFeed,omitempty" msgpack:"commentFeed,omitempty" bson:"commentfeed,omitempty" mapstructure:"commentFeed,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Data to import.
	Data *map[string][]map[string]interface{} `json:"data,omitempty" msgpack:"data,omitempty" bson:"data,omitempty" mapstructure:"data,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// The identity claims of the requester; populated by Segment Console.
	RequesterClaims *[]string `json:"requesterClaims,omitempty" msgpack:"requesterClaims,omitempty" bson:"requesterclaims,omitempty" mapstructure:"requesterClaims,omitempty"`

	// The namespace from which the request originated; populated by Segment Console.
	RequesterNamespace *string `json:"requesterNamespace,omitempty" msgpack:"requesterNamespace,omitempty" bson:"requesternamespace,omitempty" mapstructure:"requesterNamespace,omitempty"`

	// Allows the content to be changed. `Submitted`: the request moves to the target
	// namespace
	// for approval. `Approved`: the data will be created immediately. `Rejected`: the
	// request
	// cannot be changed anymore and can be deleted.
	Status *ImportRequestStatusValue `json:"status,omitempty" msgpack:"status,omitempty" bson:"status,omitempty" mapstructure:"status,omitempty"`

	// Internal field to know if the request has been submitted once.
	SubmittedOnce *bool `json:"-" msgpack:"-" bson:"submittedonce,omitempty" mapstructure:"-,omitempty"`

	// The namespace where the request will be sent. The requester can set any
	// namespace but
	// needs to have an authorization to post the request in that namespace.
	TargetNamespace *string `json:"targetNamespace,omitempty" msgpack:"targetNamespace,omitempty" bson:"targetnamespace,omitempty" mapstructure:"targetNamespace,omitempty"`

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

// NewSparseImportRequest returns a new  SparseImportRequest.
func NewSparseImportRequest() *SparseImportRequest {
	return &SparseImportRequest{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseImportRequest) Identity() elemental.Identity {

	return ImportRequestIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseImportRequest) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseImportRequest) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseImportRequest) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseImportRequest{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.CommentFeed != nil {
		s.CommentFeed = o.CommentFeed
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.Data != nil {
		s.Data = o.Data
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.RequesterClaims != nil {
		s.RequesterClaims = o.RequesterClaims
	}
	if o.RequesterNamespace != nil {
		s.RequesterNamespace = o.RequesterNamespace
	}
	if o.Status != nil {
		s.Status = o.Status
	}
	if o.SubmittedOnce != nil {
		s.SubmittedOnce = o.SubmittedOnce
	}
	if o.TargetNamespace != nil {
		s.TargetNamespace = o.TargetNamespace
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
func (o *SparseImportRequest) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseImportRequest{}
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
	if s.CommentFeed != nil {
		o.CommentFeed = s.CommentFeed
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.Data != nil {
		o.Data = s.Data
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.RequesterClaims != nil {
		o.RequesterClaims = s.RequesterClaims
	}
	if s.RequesterNamespace != nil {
		o.RequesterNamespace = s.RequesterNamespace
	}
	if s.Status != nil {
		o.Status = s.Status
	}
	if s.SubmittedOnce != nil {
		o.SubmittedOnce = s.SubmittedOnce
	}
	if s.TargetNamespace != nil {
		o.TargetNamespace = s.TargetNamespace
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
func (o *SparseImportRequest) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseImportRequest) ToPlain() elemental.PlainIdentifiable {

	out := NewImportRequest()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.Comment != nil {
		out.Comment = *o.Comment
	}
	if o.CommentFeed != nil {
		out.CommentFeed = *o.CommentFeed
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Data != nil {
		out.Data = *o.Data
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.RequesterClaims != nil {
		out.RequesterClaims = *o.RequesterClaims
	}
	if o.RequesterNamespace != nil {
		out.RequesterNamespace = *o.RequesterNamespace
	}
	if o.Status != nil {
		out.Status = *o.Status
	}
	if o.SubmittedOnce != nil {
		out.SubmittedOnce = *o.SubmittedOnce
	}
	if o.TargetNamespace != nil {
		out.TargetNamespace = *o.TargetNamespace
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
func (o *SparseImportRequest) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseImportRequest) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseImportRequest) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseImportRequest) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseImportRequest) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseImportRequest) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseImportRequest) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseImportRequest) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseImportRequest) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseImportRequest) SetDescription(description string) {

	o.Description = &description
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseImportRequest) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseImportRequest) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseImportRequest) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseImportRequest) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseImportRequest) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseImportRequest) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseImportRequest) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseImportRequest) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseImportRequest) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseImportRequest) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseImportRequest) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseImportRequest) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseImportRequest) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseImportRequest) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseImportRequest) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseImportRequest) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseImportRequest.
func (o *SparseImportRequest) DeepCopy() *SparseImportRequest {

	if o == nil {
		return nil
	}

	out := &SparseImportRequest{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseImportRequest.
func (o *SparseImportRequest) DeepCopyInto(out *SparseImportRequest) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseImportRequest: %s", err))
	}

	*out = *target.(*SparseImportRequest)
}

type mongoAttributesImportRequest struct {
	ID                   bson.ObjectId                       `bson:"_id,omitempty"`
	Annotations          map[string][]string                 `bson:"annotations"`
	AssociatedTags       []string                            `bson:"associatedtags"`
	CommentFeed          []*Comment                          `bson:"commentfeed"`
	CreateIdempotencyKey string                              `bson:"createidempotencykey"`
	CreateTime           time.Time                           `bson:"createtime"`
	Data                 map[string][]map[string]interface{} `bson:"data"`
	Description          string                              `bson:"description"`
	MigrationsLog        map[string]string                   `bson:"migrationslog,omitempty"`
	Namespace            string                              `bson:"namespace"`
	NormalizedTags       []string                            `bson:"normalizedtags"`
	Protected            bool                                `bson:"protected"`
	RequesterClaims      []string                            `bson:"requesterclaims"`
	RequesterNamespace   string                              `bson:"requesternamespace"`
	Status               ImportRequestStatusValue            `bson:"status"`
	SubmittedOnce        bool                                `bson:"submittedonce"`
	TargetNamespace      string                              `bson:"targetnamespace"`
	UpdateIdempotencyKey string                              `bson:"updateidempotencykey"`
	UpdateTime           time.Time                           `bson:"updatetime"`
	ZHash                int                                 `bson:"zhash"`
	Zone                 int                                 `bson:"zone"`
}
type mongoAttributesSparseImportRequest struct {
	ID                   bson.ObjectId                        `bson:"_id,omitempty"`
	Annotations          *map[string][]string                 `bson:"annotations,omitempty"`
	AssociatedTags       *[]string                            `bson:"associatedtags,omitempty"`
	CommentFeed          *[]*Comment                          `bson:"commentfeed,omitempty"`
	CreateIdempotencyKey *string                              `bson:"createidempotencykey,omitempty"`
	CreateTime           *time.Time                           `bson:"createtime,omitempty"`
	Data                 *map[string][]map[string]interface{} `bson:"data,omitempty"`
	Description          *string                              `bson:"description,omitempty"`
	MigrationsLog        *map[string]string                   `bson:"migrationslog,omitempty"`
	Namespace            *string                              `bson:"namespace,omitempty"`
	NormalizedTags       *[]string                            `bson:"normalizedtags,omitempty"`
	Protected            *bool                                `bson:"protected,omitempty"`
	RequesterClaims      *[]string                            `bson:"requesterclaims,omitempty"`
	RequesterNamespace   *string                              `bson:"requesternamespace,omitempty"`
	Status               *ImportRequestStatusValue            `bson:"status,omitempty"`
	SubmittedOnce        *bool                                `bson:"submittedonce,omitempty"`
	TargetNamespace      *string                              `bson:"targetnamespace,omitempty"`
	UpdateIdempotencyKey *string                              `bson:"updateidempotencykey,omitempty"`
	UpdateTime           *time.Time                           `bson:"updatetime,omitempty"`
	ZHash                *int                                 `bson:"zhash,omitempty"`
	Zone                 *int                                 `bson:"zone,omitempty"`
}
