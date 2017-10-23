package types

// ServiceParameterType is the type representing the type of a parameter
type ServiceParameterType string

// Various values for ServiceParameterType.
const (
	ServiceParameterTypeBool        ServiceParameterType = "bool"
	ServiceParameterTypeDuration    ServiceParameterType = "duration"
	ServiceParameterTypeEmum        ServiceParameterType = "enum"
	ServiceParameterTypeIntSlice    ServiceParameterType = "intSlice"
	ServiceParameterTypeInt         ServiceParameterType = "int"
	ServiceParameterTypeFloat       ServiceParameterType = "float"
	ServiceParameterTypeFloatSlice  ServiceParameterType = "floatSlice"
	ServiceParameterTypePassword    ServiceParameterType = "password"
	ServiceParameterTypeString      ServiceParameterType = "string"
	ServiceParameterTypeStringSlice ServiceParameterType = "stringSlice"
)

// ServiceParameterLink defines the link of the service parameter.
type ServiceParameterLink int

const (
	// ServiceParameterLinkInstallationSecret defines a link to installation secret.
	ServiceParameterLinkInstallationSecret ServiceParameterLink = iota

	// ServiceParameterLinkInstallationConfigMap defines a link to installation config map.
	ServiceParameterLinkInstallationConfigMap

	// ServiceParameterLinkServiceSecret defines a link to service secret.
	ServiceParameterLinkServiceSecret

	// ServiceParameterLinkServiceConfigMap defines a link to service config map.
	ServiceParameterLinkServiceConfigMap
)

// ServiceParameter defines a parameter for the service.
type ServiceParameter struct {
	Name            string               `json:"name"`
	Description     string               `json:"description"`
	LongDescription string               `json:"longDescription"`
	Key             string               `json:"key"`
	Value           string               `json:"value"`
	Env             string               `json:"-"`
	Secret          bool                 `json:"secret"`
	Type            ServiceParameterType `json:"type"`
	AllowedValues   []interface{}        `json:"allowedValues"`
	DefaultValue    interface{}          `json:"defaultValue"`
	MountPath       string               `json:"-"`
	LinkType        ServiceParameterLink `json:"-"`
	Optional        bool                 `json:"optional"`
}

// NewServiceParameter creates a new parameter.
func NewServiceParameter() *ServiceParameter {

	return &ServiceParameter{}
}

// Copy returns a copy of the current paramter.
func (p *ServiceParameter) Copy() *ServiceParameter {

	copy := NewServiceParameter()
	copy.Name = p.Name
	copy.Description = p.Description
	copy.Key = p.Key
	copy.Value = p.Value
	copy.Env = p.Env
	copy.Secret = p.Secret
	copy.MountPath = p.MountPath
	copy.Type = p.Type
	copy.Optional = p.Optional
	copy.LinkType = p.LinkType
	copy.AllowedValues = append(copy.AllowedValues, copy.AllowedValues...)

	return copy
}

// ServiceRelatedObject defines a related object.
type ServiceRelatedObject struct {
	Namespace string `json:"-"`
	Identity  string `json:"-"`
	ID        string `json:"-"`
}

// NewServiceRelatedObject creates a new related object.
func NewServiceRelatedObject() *ServiceRelatedObject {

	return &ServiceRelatedObject{}
}

// ServiceRelatedObjectOption is to prepare the future :)
type ServiceRelatedObjectOption struct{}
