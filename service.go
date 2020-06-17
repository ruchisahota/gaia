package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ServiceTLSTypeValue represents the possible values for attribute "TLSType".
type ServiceTLSTypeValue string

const (
	// ServiceTLSTypeAporeto represents the value Aporeto.
	ServiceTLSTypeAporeto ServiceTLSTypeValue = "Aporeto"

	// ServiceTLSTypeExternal represents the value External.
	ServiceTLSTypeExternal ServiceTLSTypeValue = "External"

	// ServiceTLSTypeLetsEncrypt represents the value LetsEncrypt.
	ServiceTLSTypeLetsEncrypt ServiceTLSTypeValue = "LetsEncrypt"

	// ServiceTLSTypeNone represents the value None.
	ServiceTLSTypeNone ServiceTLSTypeValue = "None"
)

// ServiceAuthorizationTypeValue represents the possible values for attribute "authorizationType".
type ServiceAuthorizationTypeValue string

const (
	// ServiceAuthorizationTypeJWT represents the value JWT.
	ServiceAuthorizationTypeJWT ServiceAuthorizationTypeValue = "JWT"

	// ServiceAuthorizationTypeMTLS represents the value MTLS.
	ServiceAuthorizationTypeMTLS ServiceAuthorizationTypeValue = "MTLS"

	// ServiceAuthorizationTypeNone represents the value None.
	ServiceAuthorizationTypeNone ServiceAuthorizationTypeValue = "None"

	// ServiceAuthorizationTypeOIDC represents the value OIDC.
	ServiceAuthorizationTypeOIDC ServiceAuthorizationTypeValue = "OIDC"
)

// ServiceTypeValue represents the possible values for attribute "type".
type ServiceTypeValue string

const (
	// ServiceTypeHTTP represents the value HTTP.
	ServiceTypeHTTP ServiceTypeValue = "HTTP"

	// ServiceTypeKubernetesSecrets represents the value KubernetesSecrets.
	ServiceTypeKubernetesSecrets ServiceTypeValue = "KubernetesSecrets"

	// ServiceTypeTCP represents the value TCP.
	ServiceTypeTCP ServiceTypeValue = "TCP"

	// ServiceTypeVaultSecrets represents the value VaultSecrets.
	ServiceTypeVaultSecrets ServiceTypeValue = "VaultSecrets"
)

// ServiceIdentity represents the Identity of the object.
var ServiceIdentity = elemental.Identity{
	Name:     "service",
	Category: "services",
	Package:  "squall",
	Private:  false,
}

// ServicesList represents a list of Services
type ServicesList []*Service

// Identity returns the identity of the objects in the list.
func (o ServicesList) Identity() elemental.Identity {

	return ServiceIdentity
}

// Copy returns a pointer to a copy the ServicesList.
func (o ServicesList) Copy() elemental.Identifiables {

	copy := append(ServicesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ServicesList.
func (o ServicesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ServicesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Service))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ServicesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ServicesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the ServicesList converted to SparseServicesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ServicesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseServicesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseService)
	}

	return out
}

// Version returns the version of the content.
func (o ServicesList) Version() int {

	return 1
}

// Service represents the model of a service
type Service struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// The list of IP addresses where the service can be accessed. This is an optional
	// attribute and
	// is only required if no host names are provided. The system will automatically
	// resolve IP
	// addresses from host names otherwise.
	IPs []string `json:"IPs" msgpack:"IPs" bson:"ips" mapstructure:"IPs,omitempty"`

	// PEM-encoded certificate that will be used to validate the user's JSON web token
	// (JWT)
	// in HTTP requests. This is an optional field, needed only if the
	// `authorizationType`
	// is set to `JWT`.
	JWTSigningCertificate string `json:"JWTSigningCertificate" msgpack:"JWTSigningCertificate" bson:"jwtsigningcertificate" mapstructure:"JWTSigningCertificate,omitempty"`

	// PEM-encoded certificate authority to use to verify client certificates. This
	// only applies
	// if `authorizationType` is set to `MTLS`. If it is not set, Segment's public
	// signing
	// certificate authority will be used.
	MTLSCertificateAuthority string `json:"MTLSCertificateAuthority" msgpack:"MTLSCertificateAuthority" bson:"mtlscertificateauthority" mapstructure:"MTLSCertificateAuthority,omitempty"`

	// This is an advanced setting. Optional OIDC callback URL. If you don't set it,
	// Segment will autodiscover it. It will be
	// `https://<hosts[0]|IPs[0]>/aporeto/oidc/callback`.
	OIDCCallbackURL string `json:"OIDCCallbackURL" msgpack:"OIDCCallbackURL" bson:"oidccallbackurl" mapstructure:"OIDCCallbackURL,omitempty"`

	// OIDC Client ID. Only has effect if the `authorizationType` is set to `OIDC`.
	OIDCClientID string `json:"OIDCClientID" msgpack:"OIDCClientID" bson:"oidcclientid" mapstructure:"OIDCClientID,omitempty"`

	// OIDC Client Secret. Only has effect if the `authorizationType` is set to `OIDC`.
	OIDCClientSecret string `json:"OIDCClientSecret" msgpack:"OIDCClientSecret" bson:"oidcclientsecret" mapstructure:"OIDCClientSecret,omitempty"`

	// OIDC discovery endpoint. Only has effect if the `authorizationType`
	// is set to `OIDC`.
	OIDCProviderURL string `json:"OIDCProviderURL" msgpack:"OIDCProviderURL" bson:"oidcproviderurl" mapstructure:"OIDCProviderURL,omitempty"`

	// Configures the scopes you want to request from the OIDC provider. Only has
	// effect
	// if `authorizationType` is set to `OIDC`.
	OIDCScopes []string `json:"OIDCScopes" msgpack:"OIDCScopes" bson:"oidcscopes" mapstructure:"OIDCScopes,omitempty"`

	// PEM-encoded certificate to expose to the clients for TLS. Only has effect and
	// required if `TLSType` is set to `External`.
	TLSCertificate string `json:"TLSCertificate" msgpack:"TLSCertificate" bson:"tlscertificate" mapstructure:"TLSCertificate,omitempty"`

	// PEM-encoded certificate key associated with `TLSCertificate`. Only has effect
	// and
	// required if `TLSType` is set to `External`.
	TLSCertificateKey string `json:"TLSCertificateKey" msgpack:"TLSCertificateKey" bson:"tlscertificatekey" mapstructure:"TLSCertificateKey,omitempty"`

	// Set how to provide a server certificate to the service.
	//
	// - `Aporeto`: Generate a certificate issued from the Segment public CA.
	// - `LetsEncrypt`: Issue a certificate from Let's Encrypt.
	// - `External`: : Let you define your own certificate and key to use.
	// - `None`: : TLS is disabled (not recommended).
	TLSType ServiceTLSTypeValue `json:"TLSType" msgpack:"TLSType" bson:"tlstype" mapstructure:"TLSType,omitempty"`

	// This is a set of all API tags for matching in the DB.
	AllAPITags []string `json:"-" msgpack:"-" bson:"allapitags" mapstructure:"-,omitempty"`

	// This is a set of all selector tags for matching in the DB.
	AllServiceTags []string `json:"-" msgpack:"-" bson:"allservicetags" mapstructure:"-,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// Defines if the object is archived.
	Archived bool `json:"-" msgpack:"-" bson:"archived" mapstructure:"-,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Defines the user authorization type that should be used.
	//
	// - `None` (default): No authorization.
	// - `JWT`:  Configures a simple JWT verification from the HTTP `Authorization`
	// header.
	// - `OIDC`: Configures OIDC authorization. You must then set
	// `OIDCClientID`,`OIDCClientSecret`, `OIDCProviderURL`.
	// - `MTLS`: Configures client certificate authorization. Then you can optionally
	// use `MTLSCertificateAuthority`, otherwise Segment's public signing certificate
	// will be used.
	AuthorizationType ServiceAuthorizationTypeValue `json:"authorizationType" msgpack:"authorizationType" bson:"authorizationtype" mapstructure:"authorizationType,omitempty"`

	// Defines a list of mappings between claims and HTTP headers. When these mappings
	// are defined,
	// the defender will copy the values of the claims to the corresponding HTTP
	// headers.
	ClaimsToHTTPHeaderMappings []*ClaimMapping `json:"claimsToHTTPHeaderMappings" msgpack:"claimsToHTTPHeaderMappings" bson:"claimstohttpheadermappings" mapstructure:"claimsToHTTPHeaderMappings,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Defines if the property is disabled.
	Disabled bool `json:"disabled" msgpack:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// Resolves the API endpoints that the service is exposing. Only valid during
	// policy rendering.
	Endpoints []*Endpoint `json:"endpoints" msgpack:"endpoints" bson:"-" mapstructure:"endpoints,omitempty"`

	// Contains a tag expression that will determine which APIs a service is exposing.
	// The APIs can be defined as the `RESTAPISpec` or similar specifications for other
	// layer 7 protocols.
	ExposedAPIs [][]string `json:"exposedAPIs" msgpack:"exposedAPIs" bson:"exposedapis" mapstructure:"exposedAPIs,omitempty"`

	// The port that the service can be accessed on. Note that this is different from
	// the
	// `port` attribute that describes the port that the service is actually listening
	// on.
	// For example if a load balancer is used, the `exposedPort` is the port that the
	// load
	// balancer is listening for the service, whereas the port that the implementation
	// is
	// listening can be different.
	ExposedPort int `json:"exposedPort" msgpack:"exposedPort" bson:"exposedport" mapstructure:"exposedPort,omitempty"`

	// Indicates that the exposed service is TLS. This means that the defender has to
	// initiate a
	// TLS session in order to forward traffic to the service.
	ExposedServiceIsTLS bool `json:"exposedServiceIsTLS" msgpack:"exposedServiceIsTLS" bson:"exposedserviceistls" mapstructure:"exposedServiceIsTLS,omitempty"`

	// Indicates if this is an external service.
	External bool `json:"external" msgpack:"external" bson:"external" mapstructure:"external,omitempty"`

	// The host names that the service can be accessed on.
	Hosts []string `json:"hosts" msgpack:"hosts" bson:"hosts" mapstructure:"hosts,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// The port that the implementation of the service is listening to. It can be
	// different than
	// `exposedPort`. This is needed for port mapping use cases where there are private
	// and
	// public ports.
	Port int `json:"port" msgpack:"port" bson:"port" mapstructure:"port,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// A new virtual port that the service can be accessed on, using HTTPS. Since the
	// defender
	// transparently inserts TLS in the application path, you might want to declare a
	// new port
	// where the defender listens for TLS. However, the application does not need to be
	// modified
	// and the defender will map the traffic to the correct application port. This
	// useful when
	// an application is being accessed from a public network.
	PublicApplicationPort int `json:"publicApplicationPort" msgpack:"publicApplicationPort" bson:"publicapplicationport" mapstructure:"publicApplicationPort,omitempty"`

	// If this is set, the user will be redirected to that URL in case of any
	// authorization
	// failure, allowing you to provide a nice message to the user. The query parameter
	// `?failure_message=<message>` will be added to that URL explaining the possible
	// reasons
	// of the failure.
	RedirectURLOnAuthorizationFailure string `json:"redirectURLOnAuthorizationFailure" msgpack:"redirectURLOnAuthorizationFailure" bson:"redirecturlonauthorizationfailure" mapstructure:"redirectURLOnAuthorizationFailure,omitempty"`

	// A tag or tag expression that identifies the processing unit that implements this
	// particular service.
	Selectors [][]string `json:"selectors" msgpack:"selectors" bson:"selectors" mapstructure:"selectors,omitempty"`

	// PEM-encoded certificate authorities to trust when additional hops are needed. It
	// must be
	// set if the service must reach a service marked as `external` or must go through
	// an
	// additional TLS termination point like a layer 7 load balancer.
	TrustedCertificateAuthorities string `json:"trustedCertificateAuthorities" msgpack:"trustedCertificateAuthorities" bson:"trustedcertificateauthorities" mapstructure:"trustedCertificateAuthorities,omitempty"`

	// Type of service.
	Type ServiceTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

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

// NewService returns a new *Service
func NewService() *Service {

	return &Service{
		ModelVersion:               1,
		AllAPITags:                 []string{},
		Annotations:                map[string][]string{},
		AllServiceTags:             []string{},
		AuthorizationType:          ServiceAuthorizationTypeNone,
		ExposedAPIs:                [][]string{},
		ExposedServiceIsTLS:        false,
		External:                   false,
		Hosts:                      []string{},
		Endpoints:                  []*Endpoint{},
		ClaimsToHTTPHeaderMappings: []*ClaimMapping{},
		AssociatedTags:             []string{},
		NormalizedTags:             []string{},
		OIDCScopes:                 []string{},
		Selectors:                  [][]string{},
		MigrationsLog:              map[string]string{},
		Type:                       ServiceTypeHTTP,
		IPs:                        []string{},
		Metadata:                   []string{},
		TLSType:                    ServiceTLSTypeAporeto,
	}
}

// Identity returns the Identity of the object.
func (o *Service) Identity() elemental.Identity {

	return ServiceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Service) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Service) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Service) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesService{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.IPs = o.IPs
	s.JWTSigningCertificate = o.JWTSigningCertificate
	s.MTLSCertificateAuthority = o.MTLSCertificateAuthority
	s.OIDCCallbackURL = o.OIDCCallbackURL
	s.OIDCClientID = o.OIDCClientID
	s.OIDCClientSecret = o.OIDCClientSecret
	s.OIDCProviderURL = o.OIDCProviderURL
	s.OIDCScopes = o.OIDCScopes
	s.TLSCertificate = o.TLSCertificate
	s.TLSCertificateKey = o.TLSCertificateKey
	s.TLSType = o.TLSType
	s.AllAPITags = o.AllAPITags
	s.AllServiceTags = o.AllServiceTags
	s.Annotations = o.Annotations
	s.Archived = o.Archived
	s.AssociatedTags = o.AssociatedTags
	s.AuthorizationType = o.AuthorizationType
	s.ClaimsToHTTPHeaderMappings = o.ClaimsToHTTPHeaderMappings
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.Description = o.Description
	s.Disabled = o.Disabled
	s.ExposedAPIs = o.ExposedAPIs
	s.ExposedPort = o.ExposedPort
	s.ExposedServiceIsTLS = o.ExposedServiceIsTLS
	s.External = o.External
	s.Hosts = o.Hosts
	s.Metadata = o.Metadata
	s.MigrationsLog = o.MigrationsLog
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.Port = o.Port
	s.Protected = o.Protected
	s.PublicApplicationPort = o.PublicApplicationPort
	s.RedirectURLOnAuthorizationFailure = o.RedirectURLOnAuthorizationFailure
	s.Selectors = o.Selectors
	s.TrustedCertificateAuthorities = o.TrustedCertificateAuthorities
	s.Type = o.Type
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Service) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesService{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.IPs = s.IPs
	o.JWTSigningCertificate = s.JWTSigningCertificate
	o.MTLSCertificateAuthority = s.MTLSCertificateAuthority
	o.OIDCCallbackURL = s.OIDCCallbackURL
	o.OIDCClientID = s.OIDCClientID
	o.OIDCClientSecret = s.OIDCClientSecret
	o.OIDCProviderURL = s.OIDCProviderURL
	o.OIDCScopes = s.OIDCScopes
	o.TLSCertificate = s.TLSCertificate
	o.TLSCertificateKey = s.TLSCertificateKey
	o.TLSType = s.TLSType
	o.AllAPITags = s.AllAPITags
	o.AllServiceTags = s.AllServiceTags
	o.Annotations = s.Annotations
	o.Archived = s.Archived
	o.AssociatedTags = s.AssociatedTags
	o.AuthorizationType = s.AuthorizationType
	o.ClaimsToHTTPHeaderMappings = s.ClaimsToHTTPHeaderMappings
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.Description = s.Description
	o.Disabled = s.Disabled
	o.ExposedAPIs = s.ExposedAPIs
	o.ExposedPort = s.ExposedPort
	o.ExposedServiceIsTLS = s.ExposedServiceIsTLS
	o.External = s.External
	o.Hosts = s.Hosts
	o.Metadata = s.Metadata
	o.MigrationsLog = s.MigrationsLog
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.Port = s.Port
	o.Protected = s.Protected
	o.PublicApplicationPort = s.PublicApplicationPort
	o.RedirectURLOnAuthorizationFailure = s.RedirectURLOnAuthorizationFailure
	o.Selectors = s.Selectors
	o.TrustedCertificateAuthorities = s.TrustedCertificateAuthorities
	o.Type = s.Type
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Service) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Service) BleveType() string {

	return "service"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Service) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *Service) Doc() string {

	return `Defines a generic service object at layer 4 or layer 7 that encapsulates the
description of a micro-service. A service exposes APIs and can be implemented
through third party entities (such as a cloud provider) or through  processing
units.`
}

func (o *Service) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *Service) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *Service) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetArchived returns the Archived of the receiver.
func (o *Service) GetArchived() bool {

	return o.Archived
}

// SetArchived sets the property Archived of the receiver using the given value.
func (o *Service) SetArchived(archived bool) {

	o.Archived = archived
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Service) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *Service) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *Service) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *Service) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Service) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *Service) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *Service) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *Service) SetDescription(description string) {

	o.Description = description
}

// GetDisabled returns the Disabled of the receiver.
func (o *Service) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the given value.
func (o *Service) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetMetadata returns the Metadata of the receiver.
func (o *Service) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *Service) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *Service) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *Service) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetName returns the Name of the receiver.
func (o *Service) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *Service) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *Service) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Service) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Service) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *Service) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *Service) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *Service) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *Service) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *Service) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Service) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *Service) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *Service) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Service) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Service) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Service) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Service) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseService{
			ID:                                &o.ID,
			IPs:                               &o.IPs,
			JWTSigningCertificate:             &o.JWTSigningCertificate,
			MTLSCertificateAuthority:          &o.MTLSCertificateAuthority,
			OIDCCallbackURL:                   &o.OIDCCallbackURL,
			OIDCClientID:                      &o.OIDCClientID,
			OIDCClientSecret:                  &o.OIDCClientSecret,
			OIDCProviderURL:                   &o.OIDCProviderURL,
			OIDCScopes:                        &o.OIDCScopes,
			TLSCertificate:                    &o.TLSCertificate,
			TLSCertificateKey:                 &o.TLSCertificateKey,
			TLSType:                           &o.TLSType,
			AllAPITags:                        &o.AllAPITags,
			AllServiceTags:                    &o.AllServiceTags,
			Annotations:                       &o.Annotations,
			Archived:                          &o.Archived,
			AssociatedTags:                    &o.AssociatedTags,
			AuthorizationType:                 &o.AuthorizationType,
			ClaimsToHTTPHeaderMappings:        &o.ClaimsToHTTPHeaderMappings,
			CreateIdempotencyKey:              &o.CreateIdempotencyKey,
			CreateTime:                        &o.CreateTime,
			Description:                       &o.Description,
			Disabled:                          &o.Disabled,
			Endpoints:                         &o.Endpoints,
			ExposedAPIs:                       &o.ExposedAPIs,
			ExposedPort:                       &o.ExposedPort,
			ExposedServiceIsTLS:               &o.ExposedServiceIsTLS,
			External:                          &o.External,
			Hosts:                             &o.Hosts,
			Metadata:                          &o.Metadata,
			MigrationsLog:                     &o.MigrationsLog,
			Name:                              &o.Name,
			Namespace:                         &o.Namespace,
			NormalizedTags:                    &o.NormalizedTags,
			Port:                              &o.Port,
			Protected:                         &o.Protected,
			PublicApplicationPort:             &o.PublicApplicationPort,
			RedirectURLOnAuthorizationFailure: &o.RedirectURLOnAuthorizationFailure,
			Selectors:                         &o.Selectors,
			TrustedCertificateAuthorities:     &o.TrustedCertificateAuthorities,
			Type:                              &o.Type,
			UpdateIdempotencyKey:              &o.UpdateIdempotencyKey,
			UpdateTime:                        &o.UpdateTime,
			ZHash:                             &o.ZHash,
			Zone:                              &o.Zone,
		}
	}

	sp := &SparseService{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "IPs":
			sp.IPs = &(o.IPs)
		case "JWTSigningCertificate":
			sp.JWTSigningCertificate = &(o.JWTSigningCertificate)
		case "MTLSCertificateAuthority":
			sp.MTLSCertificateAuthority = &(o.MTLSCertificateAuthority)
		case "OIDCCallbackURL":
			sp.OIDCCallbackURL = &(o.OIDCCallbackURL)
		case "OIDCClientID":
			sp.OIDCClientID = &(o.OIDCClientID)
		case "OIDCClientSecret":
			sp.OIDCClientSecret = &(o.OIDCClientSecret)
		case "OIDCProviderURL":
			sp.OIDCProviderURL = &(o.OIDCProviderURL)
		case "OIDCScopes":
			sp.OIDCScopes = &(o.OIDCScopes)
		case "TLSCertificate":
			sp.TLSCertificate = &(o.TLSCertificate)
		case "TLSCertificateKey":
			sp.TLSCertificateKey = &(o.TLSCertificateKey)
		case "TLSType":
			sp.TLSType = &(o.TLSType)
		case "allAPITags":
			sp.AllAPITags = &(o.AllAPITags)
		case "allServiceTags":
			sp.AllServiceTags = &(o.AllServiceTags)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "archived":
			sp.Archived = &(o.Archived)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "authorizationType":
			sp.AuthorizationType = &(o.AuthorizationType)
		case "claimsToHTTPHeaderMappings":
			sp.ClaimsToHTTPHeaderMappings = &(o.ClaimsToHTTPHeaderMappings)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "disabled":
			sp.Disabled = &(o.Disabled)
		case "endpoints":
			sp.Endpoints = &(o.Endpoints)
		case "exposedAPIs":
			sp.ExposedAPIs = &(o.ExposedAPIs)
		case "exposedPort":
			sp.ExposedPort = &(o.ExposedPort)
		case "exposedServiceIsTLS":
			sp.ExposedServiceIsTLS = &(o.ExposedServiceIsTLS)
		case "external":
			sp.External = &(o.External)
		case "hosts":
			sp.Hosts = &(o.Hosts)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "port":
			sp.Port = &(o.Port)
		case "protected":
			sp.Protected = &(o.Protected)
		case "publicApplicationPort":
			sp.PublicApplicationPort = &(o.PublicApplicationPort)
		case "redirectURLOnAuthorizationFailure":
			sp.RedirectURLOnAuthorizationFailure = &(o.RedirectURLOnAuthorizationFailure)
		case "selectors":
			sp.Selectors = &(o.Selectors)
		case "trustedCertificateAuthorities":
			sp.TrustedCertificateAuthorities = &(o.TrustedCertificateAuthorities)
		case "type":
			sp.Type = &(o.Type)
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

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *Service) EncryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if o.TLSCertificateKey, err = encrypter.EncryptString(o.TLSCertificateKey); err != nil {
		return fmt.Errorf("unable to encrypt attribute 'TLSCertificateKey' for 'Service' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *Service) DecryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if o.TLSCertificateKey, err = encrypter.DecryptString(o.TLSCertificateKey); err != nil {
		return fmt.Errorf("unable to decrypt attribute 'TLSCertificateKey' for 'Service' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// Patch apply the non nil value of a *SparseService to the object.
func (o *Service) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseService)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.IPs != nil {
		o.IPs = *so.IPs
	}
	if so.JWTSigningCertificate != nil {
		o.JWTSigningCertificate = *so.JWTSigningCertificate
	}
	if so.MTLSCertificateAuthority != nil {
		o.MTLSCertificateAuthority = *so.MTLSCertificateAuthority
	}
	if so.OIDCCallbackURL != nil {
		o.OIDCCallbackURL = *so.OIDCCallbackURL
	}
	if so.OIDCClientID != nil {
		o.OIDCClientID = *so.OIDCClientID
	}
	if so.OIDCClientSecret != nil {
		o.OIDCClientSecret = *so.OIDCClientSecret
	}
	if so.OIDCProviderURL != nil {
		o.OIDCProviderURL = *so.OIDCProviderURL
	}
	if so.OIDCScopes != nil {
		o.OIDCScopes = *so.OIDCScopes
	}
	if so.TLSCertificate != nil {
		o.TLSCertificate = *so.TLSCertificate
	}
	if so.TLSCertificateKey != nil {
		o.TLSCertificateKey = *so.TLSCertificateKey
	}
	if so.TLSType != nil {
		o.TLSType = *so.TLSType
	}
	if so.AllAPITags != nil {
		o.AllAPITags = *so.AllAPITags
	}
	if so.AllServiceTags != nil {
		o.AllServiceTags = *so.AllServiceTags
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.Archived != nil {
		o.Archived = *so.Archived
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.AuthorizationType != nil {
		o.AuthorizationType = *so.AuthorizationType
	}
	if so.ClaimsToHTTPHeaderMappings != nil {
		o.ClaimsToHTTPHeaderMappings = *so.ClaimsToHTTPHeaderMappings
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
	if so.Disabled != nil {
		o.Disabled = *so.Disabled
	}
	if so.Endpoints != nil {
		o.Endpoints = *so.Endpoints
	}
	if so.ExposedAPIs != nil {
		o.ExposedAPIs = *so.ExposedAPIs
	}
	if so.ExposedPort != nil {
		o.ExposedPort = *so.ExposedPort
	}
	if so.ExposedServiceIsTLS != nil {
		o.ExposedServiceIsTLS = *so.ExposedServiceIsTLS
	}
	if so.External != nil {
		o.External = *so.External
	}
	if so.Hosts != nil {
		o.Hosts = *so.Hosts
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
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Port != nil {
		o.Port = *so.Port
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.PublicApplicationPort != nil {
		o.PublicApplicationPort = *so.PublicApplicationPort
	}
	if so.RedirectURLOnAuthorizationFailure != nil {
		o.RedirectURLOnAuthorizationFailure = *so.RedirectURLOnAuthorizationFailure
	}
	if so.Selectors != nil {
		o.Selectors = *so.Selectors
	}
	if so.TrustedCertificateAuthorities != nil {
		o.TrustedCertificateAuthorities = *so.TrustedCertificateAuthorities
	}
	if so.Type != nil {
		o.Type = *so.Type
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

// DeepCopy returns a deep copy if the Service.
func (o *Service) DeepCopy() *Service {

	if o == nil {
		return nil
	}

	out := &Service{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Service.
func (o *Service) DeepCopyInto(out *Service) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Service: %s", err))
	}

	*out = *target.(*Service)
}

// Validate valides the current information stored into the structure.
func (o *Service) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("TLSType", string(o.TLSType), []string{"Aporeto", "LetsEncrypt", "External", "None"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("authorizationType", string(o.AuthorizationType), []string{"None", "JWT", "OIDC", "MTLS"}, false); err != nil {
		errors = errors.Append(err)
	}

	for _, sub := range o.ClaimsToHTTPHeaderMappings {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	for _, sub := range o.Endpoints {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if err := ValidateTagsExpression("exposedAPIs", o.ExposedAPIs); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("exposedPort", o.ExposedPort); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumInt("exposedPort", o.ExposedPort, int(65535), false); err != nil {
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

	if err := elemental.ValidateMaximumInt("port", o.Port, int(65535), false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumInt("publicApplicationPort", o.PublicApplicationPort, int(65535), false); err != nil {
		errors = errors.Append(err)
	}

	if err := ValidateTagsExpression("selectors", o.Selectors); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"HTTP", "TCP", "KubernetesSecrets", "VaultSecrets"}, false); err != nil {
		errors = errors.Append(err)
	}

	// Custom object validation.
	if err := ValidateServiceEntity(o); err != nil {
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
func (*Service) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ServiceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ServiceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Service) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ServiceAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Service) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "IPs":
		return o.IPs
	case "JWTSigningCertificate":
		return o.JWTSigningCertificate
	case "MTLSCertificateAuthority":
		return o.MTLSCertificateAuthority
	case "OIDCCallbackURL":
		return o.OIDCCallbackURL
	case "OIDCClientID":
		return o.OIDCClientID
	case "OIDCClientSecret":
		return o.OIDCClientSecret
	case "OIDCProviderURL":
		return o.OIDCProviderURL
	case "OIDCScopes":
		return o.OIDCScopes
	case "TLSCertificate":
		return o.TLSCertificate
	case "TLSCertificateKey":
		return o.TLSCertificateKey
	case "TLSType":
		return o.TLSType
	case "allAPITags":
		return o.AllAPITags
	case "allServiceTags":
		return o.AllServiceTags
	case "annotations":
		return o.Annotations
	case "archived":
		return o.Archived
	case "associatedTags":
		return o.AssociatedTags
	case "authorizationType":
		return o.AuthorizationType
	case "claimsToHTTPHeaderMappings":
		return o.ClaimsToHTTPHeaderMappings
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "disabled":
		return o.Disabled
	case "endpoints":
		return o.Endpoints
	case "exposedAPIs":
		return o.ExposedAPIs
	case "exposedPort":
		return o.ExposedPort
	case "exposedServiceIsTLS":
		return o.ExposedServiceIsTLS
	case "external":
		return o.External
	case "hosts":
		return o.Hosts
	case "metadata":
		return o.Metadata
	case "migrationsLog":
		return o.MigrationsLog
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "port":
		return o.Port
	case "protected":
		return o.Protected
	case "publicApplicationPort":
		return o.PublicApplicationPort
	case "redirectURLOnAuthorizationFailure":
		return o.RedirectURLOnAuthorizationFailure
	case "selectors":
		return o.Selectors
	case "trustedCertificateAuthorities":
		return o.TrustedCertificateAuthorities
	case "type":
		return o.Type
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

// ServiceAttributesMap represents the map of attribute for Service.
var ServiceAttributesMap = map[string]elemental.AttributeSpecification{
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
	"IPs": {
		AllowedChoices: []string{},
		ConvertedName:  "IPs",
		Description: `The list of IP addresses where the service can be accessed. This is an optional
attribute and
is only required if no host names are provided. The system will automatically
resolve IP
addresses from host names otherwise.`,
		Exposed: true,
		Name:    "IPs",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"JWTSigningCertificate": {
		AllowedChoices: []string{},
		ConvertedName:  "JWTSigningCertificate",
		Description: `PEM-encoded certificate that will be used to validate the user's JSON web token
(JWT)
in HTTP requests. This is an optional field, needed only if the
` + "`" + `authorizationType` + "`" + `
is set to ` + "`" + `JWT` + "`" + `.`,
		Exposed: true,
		Name:    "JWTSigningCertificate",
		Stored:  true,
		Type:    "string",
	},
	"MTLSCertificateAuthority": {
		AllowedChoices: []string{},
		ConvertedName:  "MTLSCertificateAuthority",
		Description: `PEM-encoded certificate authority to use to verify client certificates. This
only applies
if ` + "`" + `authorizationType` + "`" + ` is set to ` + "`" + `MTLS` + "`" + `. If it is not set, Segment's public
signing
certificate authority will be used.`,
		Exposed: true,
		Name:    "MTLSCertificateAuthority",
		Stored:  true,
		Type:    "string",
	},
	"OIDCCallbackURL": {
		AllowedChoices: []string{},
		ConvertedName:  "OIDCCallbackURL",
		Description: `This is an advanced setting. Optional OIDC callback URL. If you don't set it,
Segment will autodiscover it. It will be
` + "`" + `https://<hosts[0]|IPs[0]>/aporeto/oidc/callback` + "`" + `.`,
		Exposed: true,
		Name:    "OIDCCallbackURL",
		Stored:  true,
		Type:    "string",
	},
	"OIDCClientID": {
		AllowedChoices: []string{},
		ConvertedName:  "OIDCClientID",
		Description:    `OIDC Client ID. Only has effect if the ` + "`" + `authorizationType` + "`" + ` is set to ` + "`" + `OIDC` + "`" + `.`,
		Exposed:        true,
		Name:           "OIDCClientID",
		Stored:         true,
		Type:           "string",
	},
	"OIDCClientSecret": {
		AllowedChoices: []string{},
		ConvertedName:  "OIDCClientSecret",
		Description:    `OIDC Client Secret. Only has effect if the ` + "`" + `authorizationType` + "`" + ` is set to ` + "`" + `OIDC` + "`" + `.`,
		Exposed:        true,
		Name:           "OIDCClientSecret",
		Stored:         true,
		Type:           "string",
	},
	"OIDCProviderURL": {
		AllowedChoices: []string{},
		ConvertedName:  "OIDCProviderURL",
		Description: `OIDC discovery endpoint. Only has effect if the ` + "`" + `authorizationType` + "`" + `
is set to ` + "`" + `OIDC` + "`" + `.`,
		Exposed: true,
		Name:    "OIDCProviderURL",
		Stored:  true,
		Type:    "string",
	},
	"OIDCScopes": {
		AllowedChoices: []string{},
		ConvertedName:  "OIDCScopes",
		Description: `Configures the scopes you want to request from the OIDC provider. Only has
effect
if ` + "`" + `authorizationType` + "`" + ` is set to ` + "`" + `OIDC` + "`" + `.`,
		Exposed: true,
		Name:    "OIDCScopes",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"TLSCertificate": {
		AllowedChoices: []string{},
		ConvertedName:  "TLSCertificate",
		Description: `PEM-encoded certificate to expose to the clients for TLS. Only has effect and
required if ` + "`" + `TLSType` + "`" + ` is set to ` + "`" + `External` + "`" + `.`,
		Exposed: true,
		Name:    "TLSCertificate",
		Stored:  true,
		Type:    "string",
	},
	"TLSCertificateKey": {
		AllowedChoices: []string{},
		ConvertedName:  "TLSCertificateKey",
		Description: `PEM-encoded certificate key associated with ` + "`" + `TLSCertificate` + "`" + `. Only has effect
and
required if ` + "`" + `TLSType` + "`" + ` is set to ` + "`" + `External` + "`" + `.`,
		Encrypted: true,
		Exposed:   true,
		Name:      "TLSCertificateKey",
		Stored:    true,
		Type:      "string",
	},
	"TLSType": {
		AllowedChoices: []string{"Aporeto", "LetsEncrypt", "External", "None"},
		ConvertedName:  "TLSType",
		DefaultValue:   ServiceTLSTypeAporeto,
		Description: `Set how to provide a server certificate to the service.

- ` + "`" + `Aporeto` + "`" + `: Generate a certificate issued from the Segment public CA.
- ` + "`" + `LetsEncrypt` + "`" + `: Issue a certificate from Let's Encrypt.
- ` + "`" + `External` + "`" + `: : Let you define your own certificate and key to use.
- ` + "`" + `None` + "`" + `: : TLS is disabled (not recommended).`,
		Exposed: true,
		Name:    "TLSType",
		Stored:  true,
		Type:    "enum",
	},
	"AllAPITags": {
		AllowedChoices: []string{},
		ConvertedName:  "AllAPITags",
		Description:    `This is a set of all API tags for matching in the DB.`,
		Name:           "allAPITags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"AllServiceTags": {
		AllowedChoices: []string{},
		ConvertedName:  "AllServiceTags",
		Description:    `This is a set of all selector tags for matching in the DB.`,
		Name:           "allServiceTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"Archived": {
		AllowedChoices: []string{},
		ConvertedName:  "Archived",
		Description:    `Defines if the object is archived.`,
		Getter:         true,
		Name:           "archived",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
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
	"AuthorizationType": {
		AllowedChoices: []string{"None", "JWT", "OIDC", "MTLS"},
		ConvertedName:  "AuthorizationType",
		DefaultValue:   ServiceAuthorizationTypeNone,
		Description: `Defines the user authorization type that should be used.

- ` + "`" + `None` + "`" + ` (default): No authorization.
- ` + "`" + `JWT` + "`" + `:  Configures a simple JWT verification from the HTTP ` + "`" + `Authorization` + "`" + `
header.
- ` + "`" + `OIDC` + "`" + `: Configures OIDC authorization. You must then set
` + "`" + `OIDCClientID` + "`" + `,` + "`" + `OIDCClientSecret` + "`" + `, ` + "`" + `OIDCProviderURL` + "`" + `.
- ` + "`" + `MTLS` + "`" + `: Configures client certificate authorization. Then you can optionally
use ` + "`" + `MTLSCertificateAuthority` + "`" + `, otherwise Segment's public signing certificate
will be used.`,
		Exposed: true,
		Name:    "authorizationType",
		Stored:  true,
		Type:    "enum",
	},
	"ClaimsToHTTPHeaderMappings": {
		AllowedChoices: []string{},
		ConvertedName:  "ClaimsToHTTPHeaderMappings",
		Description: `Defines a list of mappings between claims and HTTP headers. When these mappings
are defined,
the defender will copy the values of the claims to the corresponding HTTP
headers.`,
		Exposed: true,
		Name:    "claimsToHTTPHeaderMappings",
		Stored:  true,
		SubType: "claimmapping",
		Type:    "refList",
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
	"Disabled": {
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Defines if the property is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Endpoints": {
		AllowedChoices: []string{},
		ConvertedName:  "Endpoints",
		Description: `Resolves the API endpoints that the service is exposing. Only valid during
policy rendering.`,
		Exposed:  true,
		Name:     "endpoints",
		ReadOnly: true,
		SubType:  "endpoint",
		Type:     "refList",
	},
	"ExposedAPIs": {
		AllowedChoices: []string{},
		ConvertedName:  "ExposedAPIs",
		Description: `Contains a tag expression that will determine which APIs a service is exposing.
The APIs can be defined as the ` + "`" + `RESTAPISpec` + "`" + ` or similar specifications for other
layer 7 protocols.`,
		Exposed: true,
		Name:    "exposedAPIs",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"ExposedPort": {
		AllowedChoices: []string{},
		ConvertedName:  "ExposedPort",
		Description: `The port that the service can be accessed on. Note that this is different from
the
` + "`" + `port` + "`" + ` attribute that describes the port that the service is actually listening
on.
For example if a load balancer is used, the ` + "`" + `exposedPort` + "`" + ` is the port that the
load
balancer is listening for the service, whereas the port that the implementation
is
listening can be different.`,
		Exposed:  true,
		MaxValue: 65535,
		Name:     "exposedPort",
		Required: true,
		Stored:   true,
		Type:     "integer",
	},
	"ExposedServiceIsTLS": {
		AllowedChoices: []string{},
		ConvertedName:  "ExposedServiceIsTLS",
		Description: `Indicates that the exposed service is TLS. This means that the defender has to
initiate a
TLS session in order to forward traffic to the service.`,
		Exposed:    true,
		Filterable: true,
		Name:       "exposedServiceIsTLS",
		Orderable:  true,
		Stored:     true,
		Type:       "boolean",
	},
	"External": {
		AllowedChoices: []string{},
		ConvertedName:  "External",
		Description:    `Indicates if this is an external service.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "external",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"Hosts": {
		AllowedChoices: []string{},
		ConvertedName:  "Hosts",
		Description:    `The host names that the service can be accessed on.`,
		Exposed:        true,
		Name:           "hosts",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"Port": {
		AllowedChoices: []string{},
		ConvertedName:  "Port",
		Description: `The port that the implementation of the service is listening to. It can be
different than
` + "`" + `exposedPort` + "`" + `. This is needed for port mapping use cases where there are private
and
public ports.`,
		Exposed:  true,
		MaxValue: 65535,
		Name:     "port",
		Stored:   true,
		Type:     "integer",
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
	"PublicApplicationPort": {
		AllowedChoices: []string{},
		ConvertedName:  "PublicApplicationPort",
		Description: `A new virtual port that the service can be accessed on, using HTTPS. Since the
defender
transparently inserts TLS in the application path, you might want to declare a
new port
where the defender listens for TLS. However, the application does not need to be
modified
and the defender will map the traffic to the correct application port. This
useful when
an application is being accessed from a public network.`,
		Exposed:  true,
		MaxValue: 65535,
		Name:     "publicApplicationPort",
		Stored:   true,
		Type:     "integer",
	},
	"RedirectURLOnAuthorizationFailure": {
		AllowedChoices: []string{},
		ConvertedName:  "RedirectURLOnAuthorizationFailure",
		Description: `If this is set, the user will be redirected to that URL in case of any
authorization
failure, allowing you to provide a nice message to the user. The query parameter
` + "`" + `?failure_message=<message>` + "`" + ` will be added to that URL explaining the possible
reasons
of the failure.`,
		Exposed: true,
		Name:    "redirectURLOnAuthorizationFailure",
		Stored:  true,
		Type:    "string",
	},
	"Selectors": {
		AllowedChoices: []string{},
		ConvertedName:  "Selectors",
		Description: `A tag or tag expression that identifies the processing unit that implements this
particular service.`,
		Exposed: true,
		Name:    "selectors",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"TrustedCertificateAuthorities": {
		AllowedChoices: []string{},
		ConvertedName:  "TrustedCertificateAuthorities",
		Description: `PEM-encoded certificate authorities to trust when additional hops are needed. It
must be
set if the service must reach a service marked as ` + "`" + `external` + "`" + ` or must go through
an
additional TLS termination point like a layer 7 load balancer.`,
		Exposed: true,
		Name:    "trustedCertificateAuthorities",
		Stored:  true,
		Type:    "string",
	},
	"Type": {
		AllowedChoices: []string{"HTTP", "TCP", "KubernetesSecrets", "VaultSecrets"},
		ConvertedName:  "Type",
		DefaultValue:   ServiceTypeHTTP,
		Description:    `Type of service.`,
		Exposed:        true,
		Name:           "type",
		Stored:         true,
		Type:           "enum",
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

// ServiceLowerCaseAttributesMap represents the map of attribute for Service.
var ServiceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"ips": {
		AllowedChoices: []string{},
		ConvertedName:  "IPs",
		Description: `The list of IP addresses where the service can be accessed. This is an optional
attribute and
is only required if no host names are provided. The system will automatically
resolve IP
addresses from host names otherwise.`,
		Exposed: true,
		Name:    "IPs",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"jwtsigningcertificate": {
		AllowedChoices: []string{},
		ConvertedName:  "JWTSigningCertificate",
		Description: `PEM-encoded certificate that will be used to validate the user's JSON web token
(JWT)
in HTTP requests. This is an optional field, needed only if the
` + "`" + `authorizationType` + "`" + `
is set to ` + "`" + `JWT` + "`" + `.`,
		Exposed: true,
		Name:    "JWTSigningCertificate",
		Stored:  true,
		Type:    "string",
	},
	"mtlscertificateauthority": {
		AllowedChoices: []string{},
		ConvertedName:  "MTLSCertificateAuthority",
		Description: `PEM-encoded certificate authority to use to verify client certificates. This
only applies
if ` + "`" + `authorizationType` + "`" + ` is set to ` + "`" + `MTLS` + "`" + `. If it is not set, Segment's public
signing
certificate authority will be used.`,
		Exposed: true,
		Name:    "MTLSCertificateAuthority",
		Stored:  true,
		Type:    "string",
	},
	"oidccallbackurl": {
		AllowedChoices: []string{},
		ConvertedName:  "OIDCCallbackURL",
		Description: `This is an advanced setting. Optional OIDC callback URL. If you don't set it,
Segment will autodiscover it. It will be
` + "`" + `https://<hosts[0]|IPs[0]>/aporeto/oidc/callback` + "`" + `.`,
		Exposed: true,
		Name:    "OIDCCallbackURL",
		Stored:  true,
		Type:    "string",
	},
	"oidcclientid": {
		AllowedChoices: []string{},
		ConvertedName:  "OIDCClientID",
		Description:    `OIDC Client ID. Only has effect if the ` + "`" + `authorizationType` + "`" + ` is set to ` + "`" + `OIDC` + "`" + `.`,
		Exposed:        true,
		Name:           "OIDCClientID",
		Stored:         true,
		Type:           "string",
	},
	"oidcclientsecret": {
		AllowedChoices: []string{},
		ConvertedName:  "OIDCClientSecret",
		Description:    `OIDC Client Secret. Only has effect if the ` + "`" + `authorizationType` + "`" + ` is set to ` + "`" + `OIDC` + "`" + `.`,
		Exposed:        true,
		Name:           "OIDCClientSecret",
		Stored:         true,
		Type:           "string",
	},
	"oidcproviderurl": {
		AllowedChoices: []string{},
		ConvertedName:  "OIDCProviderURL",
		Description: `OIDC discovery endpoint. Only has effect if the ` + "`" + `authorizationType` + "`" + `
is set to ` + "`" + `OIDC` + "`" + `.`,
		Exposed: true,
		Name:    "OIDCProviderURL",
		Stored:  true,
		Type:    "string",
	},
	"oidcscopes": {
		AllowedChoices: []string{},
		ConvertedName:  "OIDCScopes",
		Description: `Configures the scopes you want to request from the OIDC provider. Only has
effect
if ` + "`" + `authorizationType` + "`" + ` is set to ` + "`" + `OIDC` + "`" + `.`,
		Exposed: true,
		Name:    "OIDCScopes",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"tlscertificate": {
		AllowedChoices: []string{},
		ConvertedName:  "TLSCertificate",
		Description: `PEM-encoded certificate to expose to the clients for TLS. Only has effect and
required if ` + "`" + `TLSType` + "`" + ` is set to ` + "`" + `External` + "`" + `.`,
		Exposed: true,
		Name:    "TLSCertificate",
		Stored:  true,
		Type:    "string",
	},
	"tlscertificatekey": {
		AllowedChoices: []string{},
		ConvertedName:  "TLSCertificateKey",
		Description: `PEM-encoded certificate key associated with ` + "`" + `TLSCertificate` + "`" + `. Only has effect
and
required if ` + "`" + `TLSType` + "`" + ` is set to ` + "`" + `External` + "`" + `.`,
		Encrypted: true,
		Exposed:   true,
		Name:      "TLSCertificateKey",
		Stored:    true,
		Type:      "string",
	},
	"tlstype": {
		AllowedChoices: []string{"Aporeto", "LetsEncrypt", "External", "None"},
		ConvertedName:  "TLSType",
		DefaultValue:   ServiceTLSTypeAporeto,
		Description: `Set how to provide a server certificate to the service.

- ` + "`" + `Aporeto` + "`" + `: Generate a certificate issued from the Segment public CA.
- ` + "`" + `LetsEncrypt` + "`" + `: Issue a certificate from Let's Encrypt.
- ` + "`" + `External` + "`" + `: : Let you define your own certificate and key to use.
- ` + "`" + `None` + "`" + `: : TLS is disabled (not recommended).`,
		Exposed: true,
		Name:    "TLSType",
		Stored:  true,
		Type:    "enum",
	},
	"allapitags": {
		AllowedChoices: []string{},
		ConvertedName:  "AllAPITags",
		Description:    `This is a set of all API tags for matching in the DB.`,
		Name:           "allAPITags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"allservicetags": {
		AllowedChoices: []string{},
		ConvertedName:  "AllServiceTags",
		Description:    `This is a set of all selector tags for matching in the DB.`,
		Name:           "allServiceTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"archived": {
		AllowedChoices: []string{},
		ConvertedName:  "Archived",
		Description:    `Defines if the object is archived.`,
		Getter:         true,
		Name:           "archived",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
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
	"authorizationtype": {
		AllowedChoices: []string{"None", "JWT", "OIDC", "MTLS"},
		ConvertedName:  "AuthorizationType",
		DefaultValue:   ServiceAuthorizationTypeNone,
		Description: `Defines the user authorization type that should be used.

- ` + "`" + `None` + "`" + ` (default): No authorization.
- ` + "`" + `JWT` + "`" + `:  Configures a simple JWT verification from the HTTP ` + "`" + `Authorization` + "`" + `
header.
- ` + "`" + `OIDC` + "`" + `: Configures OIDC authorization. You must then set
` + "`" + `OIDCClientID` + "`" + `,` + "`" + `OIDCClientSecret` + "`" + `, ` + "`" + `OIDCProviderURL` + "`" + `.
- ` + "`" + `MTLS` + "`" + `: Configures client certificate authorization. Then you can optionally
use ` + "`" + `MTLSCertificateAuthority` + "`" + `, otherwise Segment's public signing certificate
will be used.`,
		Exposed: true,
		Name:    "authorizationType",
		Stored:  true,
		Type:    "enum",
	},
	"claimstohttpheadermappings": {
		AllowedChoices: []string{},
		ConvertedName:  "ClaimsToHTTPHeaderMappings",
		Description: `Defines a list of mappings between claims and HTTP headers. When these mappings
are defined,
the defender will copy the values of the claims to the corresponding HTTP
headers.`,
		Exposed: true,
		Name:    "claimsToHTTPHeaderMappings",
		Stored:  true,
		SubType: "claimmapping",
		Type:    "refList",
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
	"disabled": {
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Defines if the property is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"endpoints": {
		AllowedChoices: []string{},
		ConvertedName:  "Endpoints",
		Description: `Resolves the API endpoints that the service is exposing. Only valid during
policy rendering.`,
		Exposed:  true,
		Name:     "endpoints",
		ReadOnly: true,
		SubType:  "endpoint",
		Type:     "refList",
	},
	"exposedapis": {
		AllowedChoices: []string{},
		ConvertedName:  "ExposedAPIs",
		Description: `Contains a tag expression that will determine which APIs a service is exposing.
The APIs can be defined as the ` + "`" + `RESTAPISpec` + "`" + ` or similar specifications for other
layer 7 protocols.`,
		Exposed: true,
		Name:    "exposedAPIs",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"exposedport": {
		AllowedChoices: []string{},
		ConvertedName:  "ExposedPort",
		Description: `The port that the service can be accessed on. Note that this is different from
the
` + "`" + `port` + "`" + ` attribute that describes the port that the service is actually listening
on.
For example if a load balancer is used, the ` + "`" + `exposedPort` + "`" + ` is the port that the
load
balancer is listening for the service, whereas the port that the implementation
is
listening can be different.`,
		Exposed:  true,
		MaxValue: 65535,
		Name:     "exposedPort",
		Required: true,
		Stored:   true,
		Type:     "integer",
	},
	"exposedserviceistls": {
		AllowedChoices: []string{},
		ConvertedName:  "ExposedServiceIsTLS",
		Description: `Indicates that the exposed service is TLS. This means that the defender has to
initiate a
TLS session in order to forward traffic to the service.`,
		Exposed:    true,
		Filterable: true,
		Name:       "exposedServiceIsTLS",
		Orderable:  true,
		Stored:     true,
		Type:       "boolean",
	},
	"external": {
		AllowedChoices: []string{},
		ConvertedName:  "External",
		Description:    `Indicates if this is an external service.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "external",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"hosts": {
		AllowedChoices: []string{},
		ConvertedName:  "Hosts",
		Description:    `The host names that the service can be accessed on.`,
		Exposed:        true,
		Name:           "hosts",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"port": {
		AllowedChoices: []string{},
		ConvertedName:  "Port",
		Description: `The port that the implementation of the service is listening to. It can be
different than
` + "`" + `exposedPort` + "`" + `. This is needed for port mapping use cases where there are private
and
public ports.`,
		Exposed:  true,
		MaxValue: 65535,
		Name:     "port",
		Stored:   true,
		Type:     "integer",
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
	"publicapplicationport": {
		AllowedChoices: []string{},
		ConvertedName:  "PublicApplicationPort",
		Description: `A new virtual port that the service can be accessed on, using HTTPS. Since the
defender
transparently inserts TLS in the application path, you might want to declare a
new port
where the defender listens for TLS. However, the application does not need to be
modified
and the defender will map the traffic to the correct application port. This
useful when
an application is being accessed from a public network.`,
		Exposed:  true,
		MaxValue: 65535,
		Name:     "publicApplicationPort",
		Stored:   true,
		Type:     "integer",
	},
	"redirecturlonauthorizationfailure": {
		AllowedChoices: []string{},
		ConvertedName:  "RedirectURLOnAuthorizationFailure",
		Description: `If this is set, the user will be redirected to that URL in case of any
authorization
failure, allowing you to provide a nice message to the user. The query parameter
` + "`" + `?failure_message=<message>` + "`" + ` will be added to that URL explaining the possible
reasons
of the failure.`,
		Exposed: true,
		Name:    "redirectURLOnAuthorizationFailure",
		Stored:  true,
		Type:    "string",
	},
	"selectors": {
		AllowedChoices: []string{},
		ConvertedName:  "Selectors",
		Description: `A tag or tag expression that identifies the processing unit that implements this
particular service.`,
		Exposed: true,
		Name:    "selectors",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"trustedcertificateauthorities": {
		AllowedChoices: []string{},
		ConvertedName:  "TrustedCertificateAuthorities",
		Description: `PEM-encoded certificate authorities to trust when additional hops are needed. It
must be
set if the service must reach a service marked as ` + "`" + `external` + "`" + ` or must go through
an
additional TLS termination point like a layer 7 load balancer.`,
		Exposed: true,
		Name:    "trustedCertificateAuthorities",
		Stored:  true,
		Type:    "string",
	},
	"type": {
		AllowedChoices: []string{"HTTP", "TCP", "KubernetesSecrets", "VaultSecrets"},
		ConvertedName:  "Type",
		DefaultValue:   ServiceTypeHTTP,
		Description:    `Type of service.`,
		Exposed:        true,
		Name:           "type",
		Stored:         true,
		Type:           "enum",
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

// SparseServicesList represents a list of SparseServices
type SparseServicesList []*SparseService

// Identity returns the identity of the objects in the list.
func (o SparseServicesList) Identity() elemental.Identity {

	return ServiceIdentity
}

// Copy returns a pointer to a copy the SparseServicesList.
func (o SparseServicesList) Copy() elemental.Identifiables {

	copy := append(SparseServicesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseServicesList.
func (o SparseServicesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseServicesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseService))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseServicesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseServicesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseServicesList converted to ServicesList.
func (o SparseServicesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseServicesList) Version() int {

	return 1
}

// SparseService represents the sparse version of a service.
type SparseService struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// The list of IP addresses where the service can be accessed. This is an optional
	// attribute and
	// is only required if no host names are provided. The system will automatically
	// resolve IP
	// addresses from host names otherwise.
	IPs *[]string `json:"IPs,omitempty" msgpack:"IPs,omitempty" bson:"ips,omitempty" mapstructure:"IPs,omitempty"`

	// PEM-encoded certificate that will be used to validate the user's JSON web token
	// (JWT)
	// in HTTP requests. This is an optional field, needed only if the
	// `authorizationType`
	// is set to `JWT`.
	JWTSigningCertificate *string `json:"JWTSigningCertificate,omitempty" msgpack:"JWTSigningCertificate,omitempty" bson:"jwtsigningcertificate,omitempty" mapstructure:"JWTSigningCertificate,omitempty"`

	// PEM-encoded certificate authority to use to verify client certificates. This
	// only applies
	// if `authorizationType` is set to `MTLS`. If it is not set, Segment's public
	// signing
	// certificate authority will be used.
	MTLSCertificateAuthority *string `json:"MTLSCertificateAuthority,omitempty" msgpack:"MTLSCertificateAuthority,omitempty" bson:"mtlscertificateauthority,omitempty" mapstructure:"MTLSCertificateAuthority,omitempty"`

	// This is an advanced setting. Optional OIDC callback URL. If you don't set it,
	// Segment will autodiscover it. It will be
	// `https://<hosts[0]|IPs[0]>/aporeto/oidc/callback`.
	OIDCCallbackURL *string `json:"OIDCCallbackURL,omitempty" msgpack:"OIDCCallbackURL,omitempty" bson:"oidccallbackurl,omitempty" mapstructure:"OIDCCallbackURL,omitempty"`

	// OIDC Client ID. Only has effect if the `authorizationType` is set to `OIDC`.
	OIDCClientID *string `json:"OIDCClientID,omitempty" msgpack:"OIDCClientID,omitempty" bson:"oidcclientid,omitempty" mapstructure:"OIDCClientID,omitempty"`

	// OIDC Client Secret. Only has effect if the `authorizationType` is set to `OIDC`.
	OIDCClientSecret *string `json:"OIDCClientSecret,omitempty" msgpack:"OIDCClientSecret,omitempty" bson:"oidcclientsecret,omitempty" mapstructure:"OIDCClientSecret,omitempty"`

	// OIDC discovery endpoint. Only has effect if the `authorizationType`
	// is set to `OIDC`.
	OIDCProviderURL *string `json:"OIDCProviderURL,omitempty" msgpack:"OIDCProviderURL,omitempty" bson:"oidcproviderurl,omitempty" mapstructure:"OIDCProviderURL,omitempty"`

	// Configures the scopes you want to request from the OIDC provider. Only has
	// effect
	// if `authorizationType` is set to `OIDC`.
	OIDCScopes *[]string `json:"OIDCScopes,omitempty" msgpack:"OIDCScopes,omitempty" bson:"oidcscopes,omitempty" mapstructure:"OIDCScopes,omitempty"`

	// PEM-encoded certificate to expose to the clients for TLS. Only has effect and
	// required if `TLSType` is set to `External`.
	TLSCertificate *string `json:"TLSCertificate,omitempty" msgpack:"TLSCertificate,omitempty" bson:"tlscertificate,omitempty" mapstructure:"TLSCertificate,omitempty"`

	// PEM-encoded certificate key associated with `TLSCertificate`. Only has effect
	// and
	// required if `TLSType` is set to `External`.
	TLSCertificateKey *string `json:"TLSCertificateKey,omitempty" msgpack:"TLSCertificateKey,omitempty" bson:"tlscertificatekey,omitempty" mapstructure:"TLSCertificateKey,omitempty"`

	// Set how to provide a server certificate to the service.
	//
	// - `Aporeto`: Generate a certificate issued from the Segment public CA.
	// - `LetsEncrypt`: Issue a certificate from Let's Encrypt.
	// - `External`: : Let you define your own certificate and key to use.
	// - `None`: : TLS is disabled (not recommended).
	TLSType *ServiceTLSTypeValue `json:"TLSType,omitempty" msgpack:"TLSType,omitempty" bson:"tlstype,omitempty" mapstructure:"TLSType,omitempty"`

	// This is a set of all API tags for matching in the DB.
	AllAPITags *[]string `json:"-" msgpack:"-" bson:"allapitags,omitempty" mapstructure:"-,omitempty"`

	// This is a set of all selector tags for matching in the DB.
	AllServiceTags *[]string `json:"-" msgpack:"-" bson:"allservicetags,omitempty" mapstructure:"-,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// Defines if the object is archived.
	Archived *bool `json:"-" msgpack:"-" bson:"archived,omitempty" mapstructure:"-,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// Defines the user authorization type that should be used.
	//
	// - `None` (default): No authorization.
	// - `JWT`:  Configures a simple JWT verification from the HTTP `Authorization`
	// header.
	// - `OIDC`: Configures OIDC authorization. You must then set
	// `OIDCClientID`,`OIDCClientSecret`, `OIDCProviderURL`.
	// - `MTLS`: Configures client certificate authorization. Then you can optionally
	// use `MTLSCertificateAuthority`, otherwise Segment's public signing certificate
	// will be used.
	AuthorizationType *ServiceAuthorizationTypeValue `json:"authorizationType,omitempty" msgpack:"authorizationType,omitempty" bson:"authorizationtype,omitempty" mapstructure:"authorizationType,omitempty"`

	// Defines a list of mappings between claims and HTTP headers. When these mappings
	// are defined,
	// the defender will copy the values of the claims to the corresponding HTTP
	// headers.
	ClaimsToHTTPHeaderMappings *[]*ClaimMapping `json:"claimsToHTTPHeaderMappings,omitempty" msgpack:"claimsToHTTPHeaderMappings,omitempty" bson:"claimstohttpheadermappings,omitempty" mapstructure:"claimsToHTTPHeaderMappings,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Defines if the property is disabled.
	Disabled *bool `json:"disabled,omitempty" msgpack:"disabled,omitempty" bson:"disabled,omitempty" mapstructure:"disabled,omitempty"`

	// Resolves the API endpoints that the service is exposing. Only valid during
	// policy rendering.
	Endpoints *[]*Endpoint `json:"endpoints,omitempty" msgpack:"endpoints,omitempty" bson:"-" mapstructure:"endpoints,omitempty"`

	// Contains a tag expression that will determine which APIs a service is exposing.
	// The APIs can be defined as the `RESTAPISpec` or similar specifications for other
	// layer 7 protocols.
	ExposedAPIs *[][]string `json:"exposedAPIs,omitempty" msgpack:"exposedAPIs,omitempty" bson:"exposedapis,omitempty" mapstructure:"exposedAPIs,omitempty"`

	// The port that the service can be accessed on. Note that this is different from
	// the
	// `port` attribute that describes the port that the service is actually listening
	// on.
	// For example if a load balancer is used, the `exposedPort` is the port that the
	// load
	// balancer is listening for the service, whereas the port that the implementation
	// is
	// listening can be different.
	ExposedPort *int `json:"exposedPort,omitempty" msgpack:"exposedPort,omitempty" bson:"exposedport,omitempty" mapstructure:"exposedPort,omitempty"`

	// Indicates that the exposed service is TLS. This means that the defender has to
	// initiate a
	// TLS session in order to forward traffic to the service.
	ExposedServiceIsTLS *bool `json:"exposedServiceIsTLS,omitempty" msgpack:"exposedServiceIsTLS,omitempty" bson:"exposedserviceistls,omitempty" mapstructure:"exposedServiceIsTLS,omitempty"`

	// Indicates if this is an external service.
	External *bool `json:"external,omitempty" msgpack:"external,omitempty" bson:"external,omitempty" mapstructure:"external,omitempty"`

	// The host names that the service can be accessed on.
	Hosts *[]string `json:"hosts,omitempty" msgpack:"hosts,omitempty" bson:"hosts,omitempty" mapstructure:"hosts,omitempty"`

	// Contains tags that can only be set during creation, must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// The port that the implementation of the service is listening to. It can be
	// different than
	// `exposedPort`. This is needed for port mapping use cases where there are private
	// and
	// public ports.
	Port *int `json:"port,omitempty" msgpack:"port,omitempty" bson:"port,omitempty" mapstructure:"port,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// A new virtual port that the service can be accessed on, using HTTPS. Since the
	// defender
	// transparently inserts TLS in the application path, you might want to declare a
	// new port
	// where the defender listens for TLS. However, the application does not need to be
	// modified
	// and the defender will map the traffic to the correct application port. This
	// useful when
	// an application is being accessed from a public network.
	PublicApplicationPort *int `json:"publicApplicationPort,omitempty" msgpack:"publicApplicationPort,omitempty" bson:"publicapplicationport,omitempty" mapstructure:"publicApplicationPort,omitempty"`

	// If this is set, the user will be redirected to that URL in case of any
	// authorization
	// failure, allowing you to provide a nice message to the user. The query parameter
	// `?failure_message=<message>` will be added to that URL explaining the possible
	// reasons
	// of the failure.
	RedirectURLOnAuthorizationFailure *string `json:"redirectURLOnAuthorizationFailure,omitempty" msgpack:"redirectURLOnAuthorizationFailure,omitempty" bson:"redirecturlonauthorizationfailure,omitempty" mapstructure:"redirectURLOnAuthorizationFailure,omitempty"`

	// A tag or tag expression that identifies the processing unit that implements this
	// particular service.
	Selectors *[][]string `json:"selectors,omitempty" msgpack:"selectors,omitempty" bson:"selectors,omitempty" mapstructure:"selectors,omitempty"`

	// PEM-encoded certificate authorities to trust when additional hops are needed. It
	// must be
	// set if the service must reach a service marked as `external` or must go through
	// an
	// additional TLS termination point like a layer 7 load balancer.
	TrustedCertificateAuthorities *string `json:"trustedCertificateAuthorities,omitempty" msgpack:"trustedCertificateAuthorities,omitempty" bson:"trustedcertificateauthorities,omitempty" mapstructure:"trustedCertificateAuthorities,omitempty"`

	// Type of service.
	Type *ServiceTypeValue `json:"type,omitempty" msgpack:"type,omitempty" bson:"type,omitempty" mapstructure:"type,omitempty"`

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

// NewSparseService returns a new  SparseService.
func NewSparseService() *SparseService {
	return &SparseService{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseService) Identity() elemental.Identity {

	return ServiceIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseService) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseService) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseService) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseService{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.IPs != nil {
		s.IPs = o.IPs
	}
	if o.JWTSigningCertificate != nil {
		s.JWTSigningCertificate = o.JWTSigningCertificate
	}
	if o.MTLSCertificateAuthority != nil {
		s.MTLSCertificateAuthority = o.MTLSCertificateAuthority
	}
	if o.OIDCCallbackURL != nil {
		s.OIDCCallbackURL = o.OIDCCallbackURL
	}
	if o.OIDCClientID != nil {
		s.OIDCClientID = o.OIDCClientID
	}
	if o.OIDCClientSecret != nil {
		s.OIDCClientSecret = o.OIDCClientSecret
	}
	if o.OIDCProviderURL != nil {
		s.OIDCProviderURL = o.OIDCProviderURL
	}
	if o.OIDCScopes != nil {
		s.OIDCScopes = o.OIDCScopes
	}
	if o.TLSCertificate != nil {
		s.TLSCertificate = o.TLSCertificate
	}
	if o.TLSCertificateKey != nil {
		s.TLSCertificateKey = o.TLSCertificateKey
	}
	if o.TLSType != nil {
		s.TLSType = o.TLSType
	}
	if o.AllAPITags != nil {
		s.AllAPITags = o.AllAPITags
	}
	if o.AllServiceTags != nil {
		s.AllServiceTags = o.AllServiceTags
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.Archived != nil {
		s.Archived = o.Archived
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.AuthorizationType != nil {
		s.AuthorizationType = o.AuthorizationType
	}
	if o.ClaimsToHTTPHeaderMappings != nil {
		s.ClaimsToHTTPHeaderMappings = o.ClaimsToHTTPHeaderMappings
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
	if o.Disabled != nil {
		s.Disabled = o.Disabled
	}
	if o.ExposedAPIs != nil {
		s.ExposedAPIs = o.ExposedAPIs
	}
	if o.ExposedPort != nil {
		s.ExposedPort = o.ExposedPort
	}
	if o.ExposedServiceIsTLS != nil {
		s.ExposedServiceIsTLS = o.ExposedServiceIsTLS
	}
	if o.External != nil {
		s.External = o.External
	}
	if o.Hosts != nil {
		s.Hosts = o.Hosts
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
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.Port != nil {
		s.Port = o.Port
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.PublicApplicationPort != nil {
		s.PublicApplicationPort = o.PublicApplicationPort
	}
	if o.RedirectURLOnAuthorizationFailure != nil {
		s.RedirectURLOnAuthorizationFailure = o.RedirectURLOnAuthorizationFailure
	}
	if o.Selectors != nil {
		s.Selectors = o.Selectors
	}
	if o.TrustedCertificateAuthorities != nil {
		s.TrustedCertificateAuthorities = o.TrustedCertificateAuthorities
	}
	if o.Type != nil {
		s.Type = o.Type
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
func (o *SparseService) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseService{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.IPs != nil {
		o.IPs = s.IPs
	}
	if s.JWTSigningCertificate != nil {
		o.JWTSigningCertificate = s.JWTSigningCertificate
	}
	if s.MTLSCertificateAuthority != nil {
		o.MTLSCertificateAuthority = s.MTLSCertificateAuthority
	}
	if s.OIDCCallbackURL != nil {
		o.OIDCCallbackURL = s.OIDCCallbackURL
	}
	if s.OIDCClientID != nil {
		o.OIDCClientID = s.OIDCClientID
	}
	if s.OIDCClientSecret != nil {
		o.OIDCClientSecret = s.OIDCClientSecret
	}
	if s.OIDCProviderURL != nil {
		o.OIDCProviderURL = s.OIDCProviderURL
	}
	if s.OIDCScopes != nil {
		o.OIDCScopes = s.OIDCScopes
	}
	if s.TLSCertificate != nil {
		o.TLSCertificate = s.TLSCertificate
	}
	if s.TLSCertificateKey != nil {
		o.TLSCertificateKey = s.TLSCertificateKey
	}
	if s.TLSType != nil {
		o.TLSType = s.TLSType
	}
	if s.AllAPITags != nil {
		o.AllAPITags = s.AllAPITags
	}
	if s.AllServiceTags != nil {
		o.AllServiceTags = s.AllServiceTags
	}
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.Archived != nil {
		o.Archived = s.Archived
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.AuthorizationType != nil {
		o.AuthorizationType = s.AuthorizationType
	}
	if s.ClaimsToHTTPHeaderMappings != nil {
		o.ClaimsToHTTPHeaderMappings = s.ClaimsToHTTPHeaderMappings
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
	if s.Disabled != nil {
		o.Disabled = s.Disabled
	}
	if s.ExposedAPIs != nil {
		o.ExposedAPIs = s.ExposedAPIs
	}
	if s.ExposedPort != nil {
		o.ExposedPort = s.ExposedPort
	}
	if s.ExposedServiceIsTLS != nil {
		o.ExposedServiceIsTLS = s.ExposedServiceIsTLS
	}
	if s.External != nil {
		o.External = s.External
	}
	if s.Hosts != nil {
		o.Hosts = s.Hosts
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
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.Port != nil {
		o.Port = s.Port
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.PublicApplicationPort != nil {
		o.PublicApplicationPort = s.PublicApplicationPort
	}
	if s.RedirectURLOnAuthorizationFailure != nil {
		o.RedirectURLOnAuthorizationFailure = s.RedirectURLOnAuthorizationFailure
	}
	if s.Selectors != nil {
		o.Selectors = s.Selectors
	}
	if s.TrustedCertificateAuthorities != nil {
		o.TrustedCertificateAuthorities = s.TrustedCertificateAuthorities
	}
	if s.Type != nil {
		o.Type = s.Type
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
func (o *SparseService) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseService) ToPlain() elemental.PlainIdentifiable {

	out := NewService()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.IPs != nil {
		out.IPs = *o.IPs
	}
	if o.JWTSigningCertificate != nil {
		out.JWTSigningCertificate = *o.JWTSigningCertificate
	}
	if o.MTLSCertificateAuthority != nil {
		out.MTLSCertificateAuthority = *o.MTLSCertificateAuthority
	}
	if o.OIDCCallbackURL != nil {
		out.OIDCCallbackURL = *o.OIDCCallbackURL
	}
	if o.OIDCClientID != nil {
		out.OIDCClientID = *o.OIDCClientID
	}
	if o.OIDCClientSecret != nil {
		out.OIDCClientSecret = *o.OIDCClientSecret
	}
	if o.OIDCProviderURL != nil {
		out.OIDCProviderURL = *o.OIDCProviderURL
	}
	if o.OIDCScopes != nil {
		out.OIDCScopes = *o.OIDCScopes
	}
	if o.TLSCertificate != nil {
		out.TLSCertificate = *o.TLSCertificate
	}
	if o.TLSCertificateKey != nil {
		out.TLSCertificateKey = *o.TLSCertificateKey
	}
	if o.TLSType != nil {
		out.TLSType = *o.TLSType
	}
	if o.AllAPITags != nil {
		out.AllAPITags = *o.AllAPITags
	}
	if o.AllServiceTags != nil {
		out.AllServiceTags = *o.AllServiceTags
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.Archived != nil {
		out.Archived = *o.Archived
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.AuthorizationType != nil {
		out.AuthorizationType = *o.AuthorizationType
	}
	if o.ClaimsToHTTPHeaderMappings != nil {
		out.ClaimsToHTTPHeaderMappings = *o.ClaimsToHTTPHeaderMappings
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
	if o.Disabled != nil {
		out.Disabled = *o.Disabled
	}
	if o.Endpoints != nil {
		out.Endpoints = *o.Endpoints
	}
	if o.ExposedAPIs != nil {
		out.ExposedAPIs = *o.ExposedAPIs
	}
	if o.ExposedPort != nil {
		out.ExposedPort = *o.ExposedPort
	}
	if o.ExposedServiceIsTLS != nil {
		out.ExposedServiceIsTLS = *o.ExposedServiceIsTLS
	}
	if o.External != nil {
		out.External = *o.External
	}
	if o.Hosts != nil {
		out.Hosts = *o.Hosts
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
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Port != nil {
		out.Port = *o.Port
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.PublicApplicationPort != nil {
		out.PublicApplicationPort = *o.PublicApplicationPort
	}
	if o.RedirectURLOnAuthorizationFailure != nil {
		out.RedirectURLOnAuthorizationFailure = *o.RedirectURLOnAuthorizationFailure
	}
	if o.Selectors != nil {
		out.Selectors = *o.Selectors
	}
	if o.TrustedCertificateAuthorities != nil {
		out.TrustedCertificateAuthorities = *o.TrustedCertificateAuthorities
	}
	if o.Type != nil {
		out.Type = *o.Type
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

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *SparseService) EncryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if *o.TLSCertificateKey, err = encrypter.EncryptString(*o.TLSCertificateKey); err != nil {
		return fmt.Errorf("unable to encrypt attribute 'TLSCertificateKey' for 'SparseService' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *SparseService) DecryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if *o.TLSCertificateKey, err = encrypter.DecryptString(*o.TLSCertificateKey); err != nil {
		return fmt.Errorf("unable to decrypt attribute 'TLSCertificateKey' for 'SparseService' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseService) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseService) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetArchived returns the Archived of the receiver.
func (o *SparseService) GetArchived() (out bool) {

	if o.Archived == nil {
		return
	}

	return *o.Archived
}

// SetArchived sets the property Archived of the receiver using the address of the given value.
func (o *SparseService) SetArchived(archived bool) {

	o.Archived = &archived
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseService) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseService) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseService) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseService) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseService) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseService) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseService) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseService) SetDescription(description string) {

	o.Description = &description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SparseService) GetDisabled() (out bool) {

	if o.Disabled == nil {
		return
	}

	return *o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the address of the given value.
func (o *SparseService) SetDisabled(disabled bool) {

	o.Disabled = &disabled
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseService) GetMetadata() (out []string) {

	if o.Metadata == nil {
		return
	}

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseService) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseService) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseService) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetName returns the Name of the receiver.
func (o *SparseService) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseService) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseService) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseService) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseService) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseService) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseService) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseService) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseService) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseService) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseService) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseService) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseService) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseService) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseService) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseService) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseService.
func (o *SparseService) DeepCopy() *SparseService {

	if o == nil {
		return nil
	}

	out := &SparseService{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseService.
func (o *SparseService) DeepCopyInto(out *SparseService) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseService: %s", err))
	}

	*out = *target.(*SparseService)
}

type mongoAttributesService struct {
	ID                                bson.ObjectId                 `bson:"_id,omitempty"`
	IPs                               []string                      `bson:"ips"`
	JWTSigningCertificate             string                        `bson:"jwtsigningcertificate"`
	MTLSCertificateAuthority          string                        `bson:"mtlscertificateauthority"`
	OIDCCallbackURL                   string                        `bson:"oidccallbackurl"`
	OIDCClientID                      string                        `bson:"oidcclientid"`
	OIDCClientSecret                  string                        `bson:"oidcclientsecret"`
	OIDCProviderURL                   string                        `bson:"oidcproviderurl"`
	OIDCScopes                        []string                      `bson:"oidcscopes"`
	TLSCertificate                    string                        `bson:"tlscertificate"`
	TLSCertificateKey                 string                        `bson:"tlscertificatekey"`
	TLSType                           ServiceTLSTypeValue           `bson:"tlstype"`
	AllAPITags                        []string                      `bson:"allapitags"`
	AllServiceTags                    []string                      `bson:"allservicetags"`
	Annotations                       map[string][]string           `bson:"annotations"`
	Archived                          bool                          `bson:"archived"`
	AssociatedTags                    []string                      `bson:"associatedtags"`
	AuthorizationType                 ServiceAuthorizationTypeValue `bson:"authorizationtype"`
	ClaimsToHTTPHeaderMappings        []*ClaimMapping               `bson:"claimstohttpheadermappings"`
	CreateIdempotencyKey              string                        `bson:"createidempotencykey"`
	CreateTime                        time.Time                     `bson:"createtime"`
	Description                       string                        `bson:"description"`
	Disabled                          bool                          `bson:"disabled"`
	ExposedAPIs                       [][]string                    `bson:"exposedapis"`
	ExposedPort                       int                           `bson:"exposedport"`
	ExposedServiceIsTLS               bool                          `bson:"exposedserviceistls"`
	External                          bool                          `bson:"external"`
	Hosts                             []string                      `bson:"hosts"`
	Metadata                          []string                      `bson:"metadata"`
	MigrationsLog                     map[string]string             `bson:"migrationslog,omitempty"`
	Name                              string                        `bson:"name"`
	Namespace                         string                        `bson:"namespace"`
	NormalizedTags                    []string                      `bson:"normalizedtags"`
	Port                              int                           `bson:"port"`
	Protected                         bool                          `bson:"protected"`
	PublicApplicationPort             int                           `bson:"publicapplicationport"`
	RedirectURLOnAuthorizationFailure string                        `bson:"redirecturlonauthorizationfailure"`
	Selectors                         [][]string                    `bson:"selectors"`
	TrustedCertificateAuthorities     string                        `bson:"trustedcertificateauthorities"`
	Type                              ServiceTypeValue              `bson:"type"`
	UpdateIdempotencyKey              string                        `bson:"updateidempotencykey"`
	UpdateTime                        time.Time                     `bson:"updatetime"`
	ZHash                             int                           `bson:"zhash"`
	Zone                              int                           `bson:"zone"`
}
type mongoAttributesSparseService struct {
	ID                                bson.ObjectId                  `bson:"_id,omitempty"`
	IPs                               *[]string                      `bson:"ips,omitempty"`
	JWTSigningCertificate             *string                        `bson:"jwtsigningcertificate,omitempty"`
	MTLSCertificateAuthority          *string                        `bson:"mtlscertificateauthority,omitempty"`
	OIDCCallbackURL                   *string                        `bson:"oidccallbackurl,omitempty"`
	OIDCClientID                      *string                        `bson:"oidcclientid,omitempty"`
	OIDCClientSecret                  *string                        `bson:"oidcclientsecret,omitempty"`
	OIDCProviderURL                   *string                        `bson:"oidcproviderurl,omitempty"`
	OIDCScopes                        *[]string                      `bson:"oidcscopes,omitempty"`
	TLSCertificate                    *string                        `bson:"tlscertificate,omitempty"`
	TLSCertificateKey                 *string                        `bson:"tlscertificatekey,omitempty"`
	TLSType                           *ServiceTLSTypeValue           `bson:"tlstype,omitempty"`
	AllAPITags                        *[]string                      `bson:"allapitags,omitempty"`
	AllServiceTags                    *[]string                      `bson:"allservicetags,omitempty"`
	Annotations                       *map[string][]string           `bson:"annotations,omitempty"`
	Archived                          *bool                          `bson:"archived,omitempty"`
	AssociatedTags                    *[]string                      `bson:"associatedtags,omitempty"`
	AuthorizationType                 *ServiceAuthorizationTypeValue `bson:"authorizationtype,omitempty"`
	ClaimsToHTTPHeaderMappings        *[]*ClaimMapping               `bson:"claimstohttpheadermappings,omitempty"`
	CreateIdempotencyKey              *string                        `bson:"createidempotencykey,omitempty"`
	CreateTime                        *time.Time                     `bson:"createtime,omitempty"`
	Description                       *string                        `bson:"description,omitempty"`
	Disabled                          *bool                          `bson:"disabled,omitempty"`
	ExposedAPIs                       *[][]string                    `bson:"exposedapis,omitempty"`
	ExposedPort                       *int                           `bson:"exposedport,omitempty"`
	ExposedServiceIsTLS               *bool                          `bson:"exposedserviceistls,omitempty"`
	External                          *bool                          `bson:"external,omitempty"`
	Hosts                             *[]string                      `bson:"hosts,omitempty"`
	Metadata                          *[]string                      `bson:"metadata,omitempty"`
	MigrationsLog                     *map[string]string             `bson:"migrationslog,omitempty"`
	Name                              *string                        `bson:"name,omitempty"`
	Namespace                         *string                        `bson:"namespace,omitempty"`
	NormalizedTags                    *[]string                      `bson:"normalizedtags,omitempty"`
	Port                              *int                           `bson:"port,omitempty"`
	Protected                         *bool                          `bson:"protected,omitempty"`
	PublicApplicationPort             *int                           `bson:"publicapplicationport,omitempty"`
	RedirectURLOnAuthorizationFailure *string                        `bson:"redirecturlonauthorizationfailure,omitempty"`
	Selectors                         *[][]string                    `bson:"selectors,omitempty"`
	TrustedCertificateAuthorities     *string                        `bson:"trustedcertificateauthorities,omitempty"`
	Type                              *ServiceTypeValue              `bson:"type,omitempty"`
	UpdateIdempotencyKey              *string                        `bson:"updateidempotencykey,omitempty"`
	UpdateTime                        *time.Time                     `bson:"updatetime,omitempty"`
	ZHash                             *int                           `bson:"zhash,omitempty"`
	Zone                              *int                           `bson:"zone,omitempty"`
}
