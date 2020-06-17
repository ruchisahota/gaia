package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// NamespaceJWTCertificateTypeValue represents the possible values for attribute "JWTCertificateType".
type NamespaceJWTCertificateTypeValue string

const (
	// NamespaceJWTCertificateTypeEC represents the value EC.
	NamespaceJWTCertificateTypeEC NamespaceJWTCertificateTypeValue = "EC"

	// NamespaceJWTCertificateTypeNone represents the value None.
	NamespaceJWTCertificateTypeNone NamespaceJWTCertificateTypeValue = "None"

	// NamespaceJWTCertificateTypeRSA represents the value RSA.
	NamespaceJWTCertificateTypeRSA NamespaceJWTCertificateTypeValue = "RSA"
)

// NamespaceIdentity represents the Identity of the object.
var NamespaceIdentity = elemental.Identity{
	Name:     "namespace",
	Category: "namespaces",
	Package:  "squall",
	Private:  false,
}

// NamespacesList represents a list of Namespaces
type NamespacesList []*Namespace

// Identity returns the identity of the objects in the list.
func (o NamespacesList) Identity() elemental.Identity {

	return NamespaceIdentity
}

// Copy returns a pointer to a copy the NamespacesList.
func (o NamespacesList) Copy() elemental.Identifiables {

	copy := append(NamespacesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the NamespacesList.
func (o NamespacesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(NamespacesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Namespace))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o NamespacesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o NamespacesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the NamespacesList converted to SparseNamespacesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o NamespacesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseNamespacesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseNamespace)
	}

	return out
}

// Version returns the version of the content.
func (o NamespacesList) Version() int {

	return 1
}

// Namespace represents the model of a namespace
type Namespace struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// JWTCertificateType defines the JWT signing certificate that must be created
	// for this namespace. If the type is none no certificate will be created.
	JWTCertificateType NamespaceJWTCertificateTypeValue `json:"JWTCertificateType" msgpack:"JWTCertificateType" bson:"jwtcertificatetype" mapstructure:"JWTCertificateType,omitempty"`

	// JWTCertificates hold the certificates used to sign tokens for this namespace.
	// This is map indexed by the ID of the certificate.
	JWTCertificates map[string]string `json:"JWTCertificates" msgpack:"JWTCertificates" bson:"jwtcertificates" mapstructure:"JWTCertificates,omitempty"`

	// The SSH certificate authority used by the namespace.
	SSHCA string `json:"-" msgpack:"-" bson:"sshca" mapstructure:"-,omitempty"`

	// If `true`, an SSH certificate authority (CA) will be generated for the
	// namespace. This CA can be deployed in SSH server to validate SSH certificates
	// issued by the controller.
	SSHCAEnabled bool `json:"SSHCAEnabled" msgpack:"SSHCAEnabled" bson:"sshcaenabled" mapstructure:"SSHCAEnabled,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedLocalCAID holds the remote ID of the certificate authority to use.
	AssociatedLocalCAID string `json:"-" msgpack:"-" bson:"associatedlocalcaid" mapstructure:"-,omitempty"`

	// The remote ID of the SSH certificate authority to use.
	AssociatedSSHCAID string `json:"associatedSSHCAID" msgpack:"associatedSSHCAID" bson:"associatedsshcaid" mapstructure:"associatedSSHCAID,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Defines if the namespace should inherit its parent zone. If this property is set
	// to `false`,  the `zoning` property will be ignored and the namespace will have
	// the same zone as its parent.
	CustomZoning bool `json:"customZoning" msgpack:"customZoning" bson:"customzoning" mapstructure:"customZoning,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// The certificate authority used by this namespace.
	LocalCA string `json:"-" msgpack:"-" bson:"localca" mapstructure:"-,omitempty"`

	// Defines if the namespace should use a local certificate
	// authority (CA). Switching it off and on again will regenerate a new CA.
	LocalCAEnabled bool `json:"localCAEnabled" msgpack:"localCAEnabled" bson:"localcaenabled" mapstructure:"localCAEnabled,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// The name of the namespace.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// List of tags that will be added to every `or` clause of all network access
	// policies in the namespace and its children.
	NetworkAccessPolicyTags []string `json:"networkAccessPolicyTags" msgpack:"networkAccessPolicyTags" bson:"networkaccesspolicytags" mapstructure:"networkAccessPolicyTags,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// List of tags that describe this namespace. All organizational tags are
	// automatically passed to policeable objects (e.g. processing units, external
	// networks, defenders) during their creation.
	OrganizationalMetadata []string `json:"organizationalMetadata" msgpack:"organizationalMetadata" bson:"organizationalmetadata" mapstructure:"organizationalMetadata,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// This flag is deprecated and has no incidence.
	ServiceCertificateValidity string `json:"serviceCertificateValidity" msgpack:"serviceCertificateValidity" bson:"servicecertificatevalidity" mapstructure:"serviceCertificateValidity,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	// Defines what zone the namespace should live in.
	Zoning int `json:"zoning" msgpack:"zoning" bson:"zoning" mapstructure:"zoning,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewNamespace returns a new *Namespace
func NewNamespace() *Namespace {

	return &Namespace{
		ModelVersion:               1,
		Annotations:                map[string][]string{},
		AssociatedTags:             []string{},
		NetworkAccessPolicyTags:    []string{},
		NormalizedTags:             []string{},
		JWTCertificates:            map[string]string{},
		OrganizationalMetadata:     []string{},
		Metadata:                   []string{},
		JWTCertificateType:         NamespaceJWTCertificateTypeNone,
		MigrationsLog:              map[string]string{},
		ServiceCertificateValidity: "168h",
	}
}

// Identity returns the Identity of the object.
func (o *Namespace) Identity() elemental.Identity {

	return NamespaceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Namespace) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Namespace) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Namespace) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesNamespace{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.JWTCertificateType = o.JWTCertificateType
	s.JWTCertificates = o.JWTCertificates
	s.SSHCA = o.SSHCA
	s.SSHCAEnabled = o.SSHCAEnabled
	s.Annotations = o.Annotations
	s.AssociatedLocalCAID = o.AssociatedLocalCAID
	s.AssociatedSSHCAID = o.AssociatedSSHCAID
	s.AssociatedTags = o.AssociatedTags
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.CustomZoning = o.CustomZoning
	s.Description = o.Description
	s.LocalCA = o.LocalCA
	s.LocalCAEnabled = o.LocalCAEnabled
	s.Metadata = o.Metadata
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NetworkAccessPolicyTags = o.NetworkAccessPolicyTags
	s.NormalizedTags = o.NormalizedTags
	s.OrganizationalMetadata = o.OrganizationalMetadata
	s.Protected = o.Protected
	s.ServiceCertificateValidity = o.ServiceCertificateValidity
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone
	s.Zoning = o.Zoning

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Namespace) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesNamespace{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.JWTCertificateType = s.JWTCertificateType
	o.JWTCertificates = s.JWTCertificates
	o.SSHCA = s.SSHCA
	o.SSHCAEnabled = s.SSHCAEnabled
	o.Annotations = s.Annotations
	o.AssociatedLocalCAID = s.AssociatedLocalCAID
	o.AssociatedSSHCAID = s.AssociatedSSHCAID
	o.AssociatedTags = s.AssociatedTags
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.CustomZoning = s.CustomZoning
	o.Description = s.Description
	o.LocalCA = s.LocalCA
	o.LocalCAEnabled = s.LocalCAEnabled
	o.Metadata = s.Metadata
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NetworkAccessPolicyTags = s.NetworkAccessPolicyTags
	o.NormalizedTags = s.NormalizedTags
	o.OrganizationalMetadata = s.OrganizationalMetadata
	o.Protected = s.Protected
	o.ServiceCertificateValidity = s.ServiceCertificateValidity
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone
	o.Zoning = s.Zoning

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Namespace) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Namespace) BleveType() string {

	return "namespace"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Namespace) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *Namespace) Doc() string {

	return `A namespace represents the core organizational unit of the system. All objects
always exist in a single namespace. A namespace can also have child namespaces.
They can be used to split the system into organizations, business units,
applications, services or any combination you like.`
}

func (o *Namespace) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *Namespace) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *Namespace) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Namespace) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *Namespace) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *Namespace) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *Namespace) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Namespace) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *Namespace) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *Namespace) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *Namespace) SetDescription(description string) {

	o.Description = description
}

// GetMetadata returns the Metadata of the receiver.
func (o *Namespace) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *Namespace) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *Namespace) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *Namespace) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *Namespace) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *Namespace) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *Namespace) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Namespace) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Namespace) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *Namespace) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *Namespace) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *Namespace) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *Namespace) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *Namespace) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Namespace) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *Namespace) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *Namespace) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Namespace) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Namespace) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Namespace) SetZone(zone int) {

	o.Zone = zone
}

// GetZoning returns the Zoning of the receiver.
func (o *Namespace) GetZoning() int {

	return o.Zoning
}

// SetZoning sets the property Zoning of the receiver using the given value.
func (o *Namespace) SetZoning(zoning int) {

	o.Zoning = zoning
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Namespace) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseNamespace{
			ID:                         &o.ID,
			JWTCertificateType:         &o.JWTCertificateType,
			JWTCertificates:            &o.JWTCertificates,
			SSHCA:                      &o.SSHCA,
			SSHCAEnabled:               &o.SSHCAEnabled,
			Annotations:                &o.Annotations,
			AssociatedLocalCAID:        &o.AssociatedLocalCAID,
			AssociatedSSHCAID:          &o.AssociatedSSHCAID,
			AssociatedTags:             &o.AssociatedTags,
			CreateIdempotencyKey:       &o.CreateIdempotencyKey,
			CreateTime:                 &o.CreateTime,
			CustomZoning:               &o.CustomZoning,
			Description:                &o.Description,
			LocalCA:                    &o.LocalCA,
			LocalCAEnabled:             &o.LocalCAEnabled,
			Metadata:                   &o.Metadata,
			MigrationsLog:              &o.MigrationsLog,
			Name:                       &o.Name,
			Namespace:                  &o.Namespace,
			NetworkAccessPolicyTags:    &o.NetworkAccessPolicyTags,
			NormalizedTags:             &o.NormalizedTags,
			OrganizationalMetadata:     &o.OrganizationalMetadata,
			Protected:                  &o.Protected,
			ServiceCertificateValidity: &o.ServiceCertificateValidity,
			UpdateIdempotencyKey:       &o.UpdateIdempotencyKey,
			UpdateTime:                 &o.UpdateTime,
			ZHash:                      &o.ZHash,
			Zone:                       &o.Zone,
			Zoning:                     &o.Zoning,
		}
	}

	sp := &SparseNamespace{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "JWTCertificateType":
			sp.JWTCertificateType = &(o.JWTCertificateType)
		case "JWTCertificates":
			sp.JWTCertificates = &(o.JWTCertificates)
		case "SSHCA":
			sp.SSHCA = &(o.SSHCA)
		case "SSHCAEnabled":
			sp.SSHCAEnabled = &(o.SSHCAEnabled)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedLocalCAID":
			sp.AssociatedLocalCAID = &(o.AssociatedLocalCAID)
		case "associatedSSHCAID":
			sp.AssociatedSSHCAID = &(o.AssociatedSSHCAID)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "customZoning":
			sp.CustomZoning = &(o.CustomZoning)
		case "description":
			sp.Description = &(o.Description)
		case "localCA":
			sp.LocalCA = &(o.LocalCA)
		case "localCAEnabled":
			sp.LocalCAEnabled = &(o.LocalCAEnabled)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "networkAccessPolicyTags":
			sp.NetworkAccessPolicyTags = &(o.NetworkAccessPolicyTags)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "organizationalMetadata":
			sp.OrganizationalMetadata = &(o.OrganizationalMetadata)
		case "protected":
			sp.Protected = &(o.Protected)
		case "serviceCertificateValidity":
			sp.ServiceCertificateValidity = &(o.ServiceCertificateValidity)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		case "zoning":
			sp.Zoning = &(o.Zoning)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseNamespace to the object.
func (o *Namespace) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseNamespace)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.JWTCertificateType != nil {
		o.JWTCertificateType = *so.JWTCertificateType
	}
	if so.JWTCertificates != nil {
		o.JWTCertificates = *so.JWTCertificates
	}
	if so.SSHCA != nil {
		o.SSHCA = *so.SSHCA
	}
	if so.SSHCAEnabled != nil {
		o.SSHCAEnabled = *so.SSHCAEnabled
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedLocalCAID != nil {
		o.AssociatedLocalCAID = *so.AssociatedLocalCAID
	}
	if so.AssociatedSSHCAID != nil {
		o.AssociatedSSHCAID = *so.AssociatedSSHCAID
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.CustomZoning != nil {
		o.CustomZoning = *so.CustomZoning
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.LocalCA != nil {
		o.LocalCA = *so.LocalCA
	}
	if so.LocalCAEnabled != nil {
		o.LocalCAEnabled = *so.LocalCAEnabled
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NetworkAccessPolicyTags != nil {
		o.NetworkAccessPolicyTags = *so.NetworkAccessPolicyTags
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.OrganizationalMetadata != nil {
		o.OrganizationalMetadata = *so.OrganizationalMetadata
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.ServiceCertificateValidity != nil {
		o.ServiceCertificateValidity = *so.ServiceCertificateValidity
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
	if so.Zoning != nil {
		o.Zoning = *so.Zoning
	}
}

// DeepCopy returns a deep copy if the Namespace.
func (o *Namespace) DeepCopy() *Namespace {

	if o == nil {
		return nil
	}

	out := &Namespace{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Namespace.
func (o *Namespace) DeepCopyInto(out *Namespace) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Namespace: %s", err))
	}

	*out = *target.(*Namespace)
}

// Validate valides the current information stored into the structure.
func (o *Namespace) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("JWTCertificateType", string(o.JWTCertificateType), []string{"RSA", "EC", "None"}, false); err != nil {
		errors = errors.Append(err)
	}

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

	if err := elemental.ValidatePattern("name", o.Name, `^[a-zA-Z0-9-_/]+$`, `must only contain alpha numerical characters, '-' or '_'`, true); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTags("networkAccessPolicyTags", o.NetworkAccessPolicyTags); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateOrganizationalMetadata("organizationalMetadata", o.OrganizationalMetadata); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTimeDuration("serviceCertificateValidity", o.ServiceCertificateValidity); err != nil {
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
func (*Namespace) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := NamespaceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return NamespaceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Namespace) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return NamespaceAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Namespace) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "JWTCertificateType":
		return o.JWTCertificateType
	case "JWTCertificates":
		return o.JWTCertificates
	case "SSHCA":
		return o.SSHCA
	case "SSHCAEnabled":
		return o.SSHCAEnabled
	case "annotations":
		return o.Annotations
	case "associatedLocalCAID":
		return o.AssociatedLocalCAID
	case "associatedSSHCAID":
		return o.AssociatedSSHCAID
	case "associatedTags":
		return o.AssociatedTags
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "customZoning":
		return o.CustomZoning
	case "description":
		return o.Description
	case "localCA":
		return o.LocalCA
	case "localCAEnabled":
		return o.LocalCAEnabled
	case "metadata":
		return o.Metadata
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "networkAccessPolicyTags":
		return o.NetworkAccessPolicyTags
	case "normalizedTags":
		return o.NormalizedTags
	case "organizationalMetadata":
		return o.OrganizationalMetadata
	case "protected":
		return o.Protected
	case "serviceCertificateValidity":
		return o.ServiceCertificateValidity
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	case "zoning":
		return o.Zoning
	}

	return nil
}

// NamespaceAttributesMap represents the map of attribute for Namespace.
var NamespaceAttributesMap = map[string]elemental.AttributeSpecification{
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
	"JWTCertificateType": {
		AllowedChoices: []string{"RSA", "EC", "None"},
		ConvertedName:  "JWTCertificateType",
		DefaultValue:   NamespaceJWTCertificateTypeNone,
		Description: `JWTCertificateType defines the JWT signing certificate that must be created
for this namespace. If the type is none no certificate will be created.`,
		Exposed: true,
		Name:    "JWTCertificateType",
		Stored:  true,
		Type:    "enum",
	},
	"JWTCertificates": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "JWTCertificates",
		Description: `JWTCertificates hold the certificates used to sign tokens for this namespace.
This is map indexed by the ID of the certificate.`,
		Exposed:  true,
		Name:     "JWTCertificates",
		ReadOnly: true,
		Stored:   true,
		SubType:  "map[string]string",
		Type:     "external",
	},
	"SSHCA": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SSHCA",
		Description:    `The SSH certificate authority used by the namespace.`,
		Name:           "SSHCA",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"SSHCAEnabled": {
		AllowedChoices: []string{},
		ConvertedName:  "SSHCAEnabled",
		Description: `If ` + "`" + `true` + "`" + `, an SSH certificate authority (CA) will be generated for the
namespace. This CA can be deployed in SSH server to validate SSH certificates
issued by the controller.`,
		Exposed:   true,
		Name:      "SSHCAEnabled",
		Orderable: true,
		Stored:    true,
		Type:      "boolean",
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
	"AssociatedLocalCAID": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedLocalCAID",
		Description:    `AssociatedLocalCAID holds the remote ID of the certificate authority to use.`,
		Name:           "associatedLocalCAID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"AssociatedSSHCAID": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedSSHCAID",
		Description:    `The remote ID of the SSH certificate authority to use.`,
		Exposed:        true,
		Name:           "associatedSSHCAID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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
	"CustomZoning": {
		AllowedChoices: []string{},
		ConvertedName:  "CustomZoning",
		CreationOnly:   true,
		Description: `Defines if the namespace should inherit its parent zone. If this property is set
to ` + "`" + `false` + "`" + `,  the ` + "`" + `zoning` + "`" + ` property will be ignored and the namespace will have
the same zone as its parent.`,
		Exposed: true,
		Name:    "customZoning",
		Stored:  true,
		Type:    "boolean",
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
	"LocalCA": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LocalCA",
		Description:    `The certificate authority used by this namespace.`,
		Name:           "localCA",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"LocalCAEnabled": {
		AllowedChoices: []string{},
		ConvertedName:  "LocalCAEnabled",
		Description: `Defines if the namespace should use a local certificate
authority (CA). Switching it off and on again will regenerate a new CA.`,
		Exposed:   true,
		Name:      "localCAEnabled",
		Orderable: true,
		Stored:    true,
		Type:      "boolean",
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
		AllowedChars:   `^[a-zA-Z0-9-_/]+$`,
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		CreationOnly:   true,
		Description:    `The name of the namespace.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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
	"NetworkAccessPolicyTags": {
		AllowedChoices: []string{},
		ConvertedName:  "NetworkAccessPolicyTags",
		Deprecated:     true,
		Description: `List of tags that will be added to every ` + "`" + `or` + "`" + ` clause of all network access
policies in the namespace and its children.`,
		Exposed: true,
		Name:    "networkAccessPolicyTags",
		Stored:  true,
		SubType: "string",
		Type:    "list",
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
	"OrganizationalMetadata": {
		AllowedChoices: []string{},
		ConvertedName:  "OrganizationalMetadata",
		Description: `List of tags that describe this namespace. All organizational tags are
automatically passed to policeable objects (e.g. processing units, external
networks, defenders) during their creation.`,
		Exposed: true,
		Name:    "organizationalMetadata",
		Stored:  true,
		SubType: "string",
		Type:    "list",
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
	"ServiceCertificateValidity": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceCertificateValidity",
		DefaultValue:   "168h",
		Deprecated:     true,
		Description:    `This flag is deprecated and has no incidence.`,
		Exposed:        true,
		Name:           "serviceCertificateValidity",
		Stored:         true,
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
	"Zoning": {
		AllowedChoices: []string{},
		ConvertedName:  "Zoning",
		CreationOnly:   true,
		Description:    `Defines what zone the namespace should live in.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zoning",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
}

// NamespaceLowerCaseAttributesMap represents the map of attribute for Namespace.
var NamespaceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"jwtcertificatetype": {
		AllowedChoices: []string{"RSA", "EC", "None"},
		ConvertedName:  "JWTCertificateType",
		DefaultValue:   NamespaceJWTCertificateTypeNone,
		Description: `JWTCertificateType defines the JWT signing certificate that must be created
for this namespace. If the type is none no certificate will be created.`,
		Exposed: true,
		Name:    "JWTCertificateType",
		Stored:  true,
		Type:    "enum",
	},
	"jwtcertificates": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "JWTCertificates",
		Description: `JWTCertificates hold the certificates used to sign tokens for this namespace.
This is map indexed by the ID of the certificate.`,
		Exposed:  true,
		Name:     "JWTCertificates",
		ReadOnly: true,
		Stored:   true,
		SubType:  "map[string]string",
		Type:     "external",
	},
	"sshca": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SSHCA",
		Description:    `The SSH certificate authority used by the namespace.`,
		Name:           "SSHCA",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"sshcaenabled": {
		AllowedChoices: []string{},
		ConvertedName:  "SSHCAEnabled",
		Description: `If ` + "`" + `true` + "`" + `, an SSH certificate authority (CA) will be generated for the
namespace. This CA can be deployed in SSH server to validate SSH certificates
issued by the controller.`,
		Exposed:   true,
		Name:      "SSHCAEnabled",
		Orderable: true,
		Stored:    true,
		Type:      "boolean",
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
	"associatedlocalcaid": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedLocalCAID",
		Description:    `AssociatedLocalCAID holds the remote ID of the certificate authority to use.`,
		Name:           "associatedLocalCAID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"associatedsshcaid": {
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedSSHCAID",
		Description:    `The remote ID of the SSH certificate authority to use.`,
		Exposed:        true,
		Name:           "associatedSSHCAID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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
	"customzoning": {
		AllowedChoices: []string{},
		ConvertedName:  "CustomZoning",
		CreationOnly:   true,
		Description: `Defines if the namespace should inherit its parent zone. If this property is set
to ` + "`" + `false` + "`" + `,  the ` + "`" + `zoning` + "`" + ` property will be ignored and the namespace will have
the same zone as its parent.`,
		Exposed: true,
		Name:    "customZoning",
		Stored:  true,
		Type:    "boolean",
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
	"localca": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LocalCA",
		Description:    `The certificate authority used by this namespace.`,
		Name:           "localCA",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"localcaenabled": {
		AllowedChoices: []string{},
		ConvertedName:  "LocalCAEnabled",
		Description: `Defines if the namespace should use a local certificate
authority (CA). Switching it off and on again will regenerate a new CA.`,
		Exposed:   true,
		Name:      "localCAEnabled",
		Orderable: true,
		Stored:    true,
		Type:      "boolean",
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
		AllowedChars:   `^[a-zA-Z0-9-_/]+$`,
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		CreationOnly:   true,
		Description:    `The name of the namespace.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
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
	"networkaccesspolicytags": {
		AllowedChoices: []string{},
		ConvertedName:  "NetworkAccessPolicyTags",
		Deprecated:     true,
		Description: `List of tags that will be added to every ` + "`" + `or` + "`" + ` clause of all network access
policies in the namespace and its children.`,
		Exposed: true,
		Name:    "networkAccessPolicyTags",
		Stored:  true,
		SubType: "string",
		Type:    "list",
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
	"organizationalmetadata": {
		AllowedChoices: []string{},
		ConvertedName:  "OrganizationalMetadata",
		Description: `List of tags that describe this namespace. All organizational tags are
automatically passed to policeable objects (e.g. processing units, external
networks, defenders) during their creation.`,
		Exposed: true,
		Name:    "organizationalMetadata",
		Stored:  true,
		SubType: "string",
		Type:    "list",
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
	"servicecertificatevalidity": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceCertificateValidity",
		DefaultValue:   "168h",
		Deprecated:     true,
		Description:    `This flag is deprecated and has no incidence.`,
		Exposed:        true,
		Name:           "serviceCertificateValidity",
		Stored:         true,
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
	"zoning": {
		AllowedChoices: []string{},
		ConvertedName:  "Zoning",
		CreationOnly:   true,
		Description:    `Defines what zone the namespace should live in.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zoning",
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
}

// SparseNamespacesList represents a list of SparseNamespaces
type SparseNamespacesList []*SparseNamespace

// Identity returns the identity of the objects in the list.
func (o SparseNamespacesList) Identity() elemental.Identity {

	return NamespaceIdentity
}

// Copy returns a pointer to a copy the SparseNamespacesList.
func (o SparseNamespacesList) Copy() elemental.Identifiables {

	copy := append(SparseNamespacesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseNamespacesList.
func (o SparseNamespacesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseNamespacesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseNamespace))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseNamespacesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseNamespacesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseNamespacesList converted to NamespacesList.
func (o SparseNamespacesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseNamespacesList) Version() int {

	return 1
}

// SparseNamespace represents the sparse version of a namespace.
type SparseNamespace struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// JWTCertificateType defines the JWT signing certificate that must be created
	// for this namespace. If the type is none no certificate will be created.
	JWTCertificateType *NamespaceJWTCertificateTypeValue `json:"JWTCertificateType,omitempty" msgpack:"JWTCertificateType,omitempty" bson:"jwtcertificatetype,omitempty" mapstructure:"JWTCertificateType,omitempty"`

	// JWTCertificates hold the certificates used to sign tokens for this namespace.
	// This is map indexed by the ID of the certificate.
	JWTCertificates *map[string]string `json:"JWTCertificates,omitempty" msgpack:"JWTCertificates,omitempty" bson:"jwtcertificates,omitempty" mapstructure:"JWTCertificates,omitempty"`

	// The SSH certificate authority used by the namespace.
	SSHCA *string `json:"-" msgpack:"-" bson:"sshca,omitempty" mapstructure:"-,omitempty"`

	// If `true`, an SSH certificate authority (CA) will be generated for the
	// namespace. This CA can be deployed in SSH server to validate SSH certificates
	// issued by the controller.
	SSHCAEnabled *bool `json:"SSHCAEnabled,omitempty" msgpack:"SSHCAEnabled,omitempty" bson:"sshcaenabled,omitempty" mapstructure:"SSHCAEnabled,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// AssociatedLocalCAID holds the remote ID of the certificate authority to use.
	AssociatedLocalCAID *string `json:"-" msgpack:"-" bson:"associatedlocalcaid,omitempty" mapstructure:"-,omitempty"`

	// The remote ID of the SSH certificate authority to use.
	AssociatedSSHCAID *string `json:"associatedSSHCAID,omitempty" msgpack:"associatedSSHCAID,omitempty" bson:"associatedsshcaid,omitempty" mapstructure:"associatedSSHCAID,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Defines if the namespace should inherit its parent zone. If this property is set
	// to `false`,  the `zoning` property will be ignored and the namespace will have
	// the same zone as its parent.
	CustomZoning *bool `json:"customZoning,omitempty" msgpack:"customZoning,omitempty" bson:"customzoning,omitempty" mapstructure:"customZoning,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// The certificate authority used by this namespace.
	LocalCA *string `json:"-" msgpack:"-" bson:"localca,omitempty" mapstructure:"-,omitempty"`

	// Defines if the namespace should use a local certificate
	// authority (CA). Switching it off and on again will regenerate a new CA.
	LocalCAEnabled *bool `json:"localCAEnabled,omitempty" msgpack:"localCAEnabled,omitempty" bson:"localcaenabled,omitempty" mapstructure:"localCAEnabled,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// The name of the namespace.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// List of tags that will be added to every `or` clause of all network access
	// policies in the namespace and its children.
	NetworkAccessPolicyTags *[]string `json:"networkAccessPolicyTags,omitempty" msgpack:"networkAccessPolicyTags,omitempty" bson:"networkaccesspolicytags,omitempty" mapstructure:"networkAccessPolicyTags,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// List of tags that describe this namespace. All organizational tags are
	// automatically passed to policeable objects (e.g. processing units, external
	// networks, defenders) during their creation.
	OrganizationalMetadata *[]string `json:"organizationalMetadata,omitempty" msgpack:"organizationalMetadata,omitempty" bson:"organizationalmetadata,omitempty" mapstructure:"organizationalMetadata,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// This flag is deprecated and has no incidence.
	ServiceCertificateValidity *string `json:"serviceCertificateValidity,omitempty" msgpack:"serviceCertificateValidity,omitempty" bson:"servicecertificatevalidity,omitempty" mapstructure:"serviceCertificateValidity,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	// Defines what zone the namespace should live in.
	Zoning *int `json:"zoning,omitempty" msgpack:"zoning,omitempty" bson:"zoning,omitempty" mapstructure:"zoning,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseNamespace returns a new  SparseNamespace.
func NewSparseNamespace() *SparseNamespace {
	return &SparseNamespace{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseNamespace) Identity() elemental.Identity {

	return NamespaceIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseNamespace) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseNamespace) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseNamespace) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseNamespace{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.JWTCertificateType != nil {
		s.JWTCertificateType = o.JWTCertificateType
	}
	if o.JWTCertificates != nil {
		s.JWTCertificates = o.JWTCertificates
	}
	if o.SSHCA != nil {
		s.SSHCA = o.SSHCA
	}
	if o.SSHCAEnabled != nil {
		s.SSHCAEnabled = o.SSHCAEnabled
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedLocalCAID != nil {
		s.AssociatedLocalCAID = o.AssociatedLocalCAID
	}
	if o.AssociatedSSHCAID != nil {
		s.AssociatedSSHCAID = o.AssociatedSSHCAID
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
	if o.CustomZoning != nil {
		s.CustomZoning = o.CustomZoning
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.LocalCA != nil {
		s.LocalCA = o.LocalCA
	}
	if o.LocalCAEnabled != nil {
		s.LocalCAEnabled = o.LocalCAEnabled
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
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NetworkAccessPolicyTags != nil {
		s.NetworkAccessPolicyTags = o.NetworkAccessPolicyTags
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.OrganizationalMetadata != nil {
		s.OrganizationalMetadata = o.OrganizationalMetadata
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.ServiceCertificateValidity != nil {
		s.ServiceCertificateValidity = o.ServiceCertificateValidity
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
	if o.Zoning != nil {
		s.Zoning = o.Zoning
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseNamespace) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseNamespace{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.JWTCertificateType != nil {
		o.JWTCertificateType = s.JWTCertificateType
	}
	if s.JWTCertificates != nil {
		o.JWTCertificates = s.JWTCertificates
	}
	if s.SSHCA != nil {
		o.SSHCA = s.SSHCA
	}
	if s.SSHCAEnabled != nil {
		o.SSHCAEnabled = s.SSHCAEnabled
	}
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedLocalCAID != nil {
		o.AssociatedLocalCAID = s.AssociatedLocalCAID
	}
	if s.AssociatedSSHCAID != nil {
		o.AssociatedSSHCAID = s.AssociatedSSHCAID
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
	if s.CustomZoning != nil {
		o.CustomZoning = s.CustomZoning
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.LocalCA != nil {
		o.LocalCA = s.LocalCA
	}
	if s.LocalCAEnabled != nil {
		o.LocalCAEnabled = s.LocalCAEnabled
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
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NetworkAccessPolicyTags != nil {
		o.NetworkAccessPolicyTags = s.NetworkAccessPolicyTags
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.OrganizationalMetadata != nil {
		o.OrganizationalMetadata = s.OrganizationalMetadata
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.ServiceCertificateValidity != nil {
		o.ServiceCertificateValidity = s.ServiceCertificateValidity
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
	if s.Zoning != nil {
		o.Zoning = s.Zoning
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseNamespace) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseNamespace) ToPlain() elemental.PlainIdentifiable {

	out := NewNamespace()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.JWTCertificateType != nil {
		out.JWTCertificateType = *o.JWTCertificateType
	}
	if o.JWTCertificates != nil {
		out.JWTCertificates = *o.JWTCertificates
	}
	if o.SSHCA != nil {
		out.SSHCA = *o.SSHCA
	}
	if o.SSHCAEnabled != nil {
		out.SSHCAEnabled = *o.SSHCAEnabled
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedLocalCAID != nil {
		out.AssociatedLocalCAID = *o.AssociatedLocalCAID
	}
	if o.AssociatedSSHCAID != nil {
		out.AssociatedSSHCAID = *o.AssociatedSSHCAID
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.CustomZoning != nil {
		out.CustomZoning = *o.CustomZoning
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.LocalCA != nil {
		out.LocalCA = *o.LocalCA
	}
	if o.LocalCAEnabled != nil {
		out.LocalCAEnabled = *o.LocalCAEnabled
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NetworkAccessPolicyTags != nil {
		out.NetworkAccessPolicyTags = *o.NetworkAccessPolicyTags
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.OrganizationalMetadata != nil {
		out.OrganizationalMetadata = *o.OrganizationalMetadata
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.ServiceCertificateValidity != nil {
		out.ServiceCertificateValidity = *o.ServiceCertificateValidity
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
	if o.Zoning != nil {
		out.Zoning = *o.Zoning
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseNamespace) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseNamespace) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseNamespace) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseNamespace) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseNamespace) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseNamespace) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseNamespace) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseNamespace) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseNamespace) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseNamespace) SetDescription(description string) {

	o.Description = &description
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseNamespace) GetMetadata() (out []string) {

	if o.Metadata == nil {
		return
	}

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseNamespace) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseNamespace) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseNamespace) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseNamespace) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseNamespace) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseNamespace) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseNamespace) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseNamespace) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseNamespace) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseNamespace) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseNamespace) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseNamespace) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseNamespace) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseNamespace) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseNamespace) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseNamespace) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseNamespace) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseNamespace) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseNamespace) SetZone(zone int) {

	o.Zone = &zone
}

// GetZoning returns the Zoning of the receiver.
func (o *SparseNamespace) GetZoning() (out int) {

	if o.Zoning == nil {
		return
	}

	return *o.Zoning
}

// SetZoning sets the property Zoning of the receiver using the address of the given value.
func (o *SparseNamespace) SetZoning(zoning int) {

	o.Zoning = &zoning
}

// DeepCopy returns a deep copy if the SparseNamespace.
func (o *SparseNamespace) DeepCopy() *SparseNamespace {

	if o == nil {
		return nil
	}

	out := &SparseNamespace{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseNamespace.
func (o *SparseNamespace) DeepCopyInto(out *SparseNamespace) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseNamespace: %s", err))
	}

	*out = *target.(*SparseNamespace)
}

type mongoAttributesNamespace struct {
	ID                         bson.ObjectId                    `bson:"_id,omitempty"`
	JWTCertificateType         NamespaceJWTCertificateTypeValue `bson:"jwtcertificatetype"`
	JWTCertificates            map[string]string                `bson:"jwtcertificates"`
	SSHCA                      string                           `bson:"sshca"`
	SSHCAEnabled               bool                             `bson:"sshcaenabled"`
	Annotations                map[string][]string              `bson:"annotations"`
	AssociatedLocalCAID        string                           `bson:"associatedlocalcaid"`
	AssociatedSSHCAID          string                           `bson:"associatedsshcaid"`
	AssociatedTags             []string                         `bson:"associatedtags"`
	CreateIdempotencyKey       string                           `bson:"createidempotencykey"`
	CreateTime                 time.Time                        `bson:"createtime"`
	CustomZoning               bool                             `bson:"customzoning"`
	Description                string                           `bson:"description"`
	LocalCA                    string                           `bson:"localca"`
	LocalCAEnabled             bool                             `bson:"localcaenabled"`
	Metadata                   []string                         `bson:"metadata"`
	MigrationsLog              map[string]string                `bson:"migrationslog,omitempty"`
	Name                       string                           `bson:"name"`
	Namespace                  string                           `bson:"namespace"`
	NetworkAccessPolicyTags    []string                         `bson:"networkaccesspolicytags"`
	NormalizedTags             []string                         `bson:"normalizedtags"`
	OrganizationalMetadata     []string                         `bson:"organizationalmetadata"`
	Protected                  bool                             `bson:"protected"`
	ServiceCertificateValidity string                           `bson:"servicecertificatevalidity"`
	UpdateIdempotencyKey       string                           `bson:"updateidempotencykey"`
	UpdateTime                 time.Time                        `bson:"updatetime"`
	ZHash                      int                              `bson:"zhash"`
	Zone                       int                              `bson:"zone"`
	Zoning                     int                              `bson:"zoning"`
}
type mongoAttributesSparseNamespace struct {
	ID                         bson.ObjectId                     `bson:"_id,omitempty"`
	JWTCertificateType         *NamespaceJWTCertificateTypeValue `bson:"jwtcertificatetype,omitempty"`
	JWTCertificates            *map[string]string                `bson:"jwtcertificates,omitempty"`
	SSHCA                      *string                           `bson:"sshca,omitempty"`
	SSHCAEnabled               *bool                             `bson:"sshcaenabled,omitempty"`
	Annotations                *map[string][]string              `bson:"annotations,omitempty"`
	AssociatedLocalCAID        *string                           `bson:"associatedlocalcaid,omitempty"`
	AssociatedSSHCAID          *string                           `bson:"associatedsshcaid,omitempty"`
	AssociatedTags             *[]string                         `bson:"associatedtags,omitempty"`
	CreateIdempotencyKey       *string                           `bson:"createidempotencykey,omitempty"`
	CreateTime                 *time.Time                        `bson:"createtime,omitempty"`
	CustomZoning               *bool                             `bson:"customzoning,omitempty"`
	Description                *string                           `bson:"description,omitempty"`
	LocalCA                    *string                           `bson:"localca,omitempty"`
	LocalCAEnabled             *bool                             `bson:"localcaenabled,omitempty"`
	Metadata                   *[]string                         `bson:"metadata,omitempty"`
	MigrationsLog              *map[string]string                `bson:"migrationslog,omitempty"`
	Name                       *string                           `bson:"name,omitempty"`
	Namespace                  *string                           `bson:"namespace,omitempty"`
	NetworkAccessPolicyTags    *[]string                         `bson:"networkaccesspolicytags,omitempty"`
	NormalizedTags             *[]string                         `bson:"normalizedtags,omitempty"`
	OrganizationalMetadata     *[]string                         `bson:"organizationalmetadata,omitempty"`
	Protected                  *bool                             `bson:"protected,omitempty"`
	ServiceCertificateValidity *string                           `bson:"servicecertificatevalidity,omitempty"`
	UpdateIdempotencyKey       *string                           `bson:"updateidempotencykey,omitempty"`
	UpdateTime                 *time.Time                        `bson:"updatetime,omitempty"`
	ZHash                      *int                              `bson:"zhash,omitempty"`
	Zone                       *int                              `bson:"zone,omitempty"`
	Zoning                     *int                              `bson:"zoning,omitempty"`
}
