package vincemodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// KubernetesclusterIdentity represents the Identity of the object
var KubernetesclusterIdentity = elemental.Identity{
	Name:     "kubernetescluster",
	Category: "kubernetesclusters",
}

// KubernetesclustersList represents a list of Kubernetesclusters
type KubernetesclustersList []*Kubernetescluster

// ContentIdentity returns the identity of the objects in the list.
func (o KubernetesclustersList) ContentIdentity() elemental.Identity {

	return KubernetesclusterIdentity
}

// List converts the object to an elemental.IdentifiablesList.
func (o KubernetesclustersList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o KubernetesclustersList) DefaultOrder() []string {

	return []string{}
}

// Kubernetescluster represents the model of a kubernetescluster
type Kubernetescluster struct {
	// Link to the certificate created for this cluster
	CertificateID string `json:"certificateID" bson:"certificateid"`

	// base64 of the .tar.gz file thjat contains all the .YAMLs files needed to create the aporeto side on your kubernetes Cluster
	KubernetesDefinitions string `json:"kubernetesDefinitions" bson:"kubernetesdefinitions"`

	// The name of your cluster
	Name string `json:"name" bson:"name"`

	// Link to your namespace
	NamespaceID string `json:"namespaceID" bson:"namespaceid"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewKubernetescluster returns a new *Kubernetescluster
func NewKubernetescluster() *Kubernetescluster {

	return &Kubernetescluster{
		ModelVersion: 1.0,
	}
}

// Identity returns the Identity of the object.
func (o *Kubernetescluster) Identity() elemental.Identity {

	return KubernetesclusterIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Kubernetescluster) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Kubernetescluster) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *Kubernetescluster) Version() float64 {

	return 1.0
}

// DefaultOrder returns the list of default ordering fields.
func (o *Kubernetescluster) DefaultOrder() []string {

	return []string{}
}

func (o *Kubernetescluster) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Kubernetescluster) Validate() error {

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
func (*Kubernetescluster) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := KubernetesclusterAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return KubernetesclusterLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Kubernetescluster) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return KubernetesclusterAttributesMap
}

// KubernetesclusterAttributesMap represents the map of attribute for Kubernetescluster.
var KubernetesclusterAttributesMap = map[string]elemental.AttributeSpecification{
	"CertificateID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Link to the certificate created for this cluster`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "certificateID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"KubernetesDefinitions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `base64 of the .tar.gz file thjat contains all the .YAMLs files needed to create the aporeto side on your kubernetes Cluster`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "kubernetesDefinitions",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `The name of your cluster`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"NamespaceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Link to your namespace `,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "namespaceID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}

// KubernetesclusterLowerCaseAttributesMap represents the map of attribute for Kubernetescluster.
var KubernetesclusterLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"certificateid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Link to the certificate created for this cluster`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "certificateID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"kubernetesdefinitions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `base64 of the .tar.gz file thjat contains all the .YAMLs files needed to create the aporeto side on your kubernetes Cluster`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "kubernetesDefinitions",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `The name of your cluster`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"namespaceid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Link to your namespace `,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "namespaceID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}
