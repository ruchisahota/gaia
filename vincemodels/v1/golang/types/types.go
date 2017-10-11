package types

// ServiceParameter defines a parameter for the service.
type ServiceParameter struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Key         string `json:"key"`
	Value       string `json:"value"`
	Env         string `json:"-"`
	Secret      bool   `json:"secret"`
	MountPath   string `json:"-"`
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

	return copy
}
