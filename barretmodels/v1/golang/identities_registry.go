package barretmodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(APICertIdentity)
	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(EnforcerCertIdentity)
}

// ModelVersion returns the current version of the model
func ModelVersion() float64 { return 1 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {
	case APICertIdentity.Name:
		return NewAPICert()
	case RootIdentity.Name:
		return NewRoot()
	case EnforcerCertIdentity.Name:
		return NewEnforcerCert()
	default:
		return nil
	}
}

// IdentifiableForCategory returns a new instance of the Identifiable for the given category name.
func IdentifiableForCategory(category string) elemental.Identifiable {

	switch category {
	case APICertIdentity.Category:
		return NewAPICert()
	case RootIdentity.Category:
		return NewRoot()
	case EnforcerCertIdentity.Category:
		return NewEnforcerCert()
	default:
		return nil
	}
}

// ContentIdentifiableForIdentity returns a new instance of a ContentIdentifiable for the given identity name.
func ContentIdentifiableForIdentity(identity string) elemental.ContentIdentifiable {

	switch identity {
	case APICertIdentity.Name:
		return &APICertsList{}
	case EnforcerCertIdentity.Name:
		return &EnforcerCertsList{}
	default:
		return nil
	}
}

// ContentIdentifiableForCategory returns a new instance of a ContentIdentifiable for the given category name.
func ContentIdentifiableForCategory(category string) elemental.ContentIdentifiable {

	switch category {
	case APICertIdentity.Category:
		return &APICertsList{}
	case EnforcerCertIdentity.Category:
		return &EnforcerCertsList{}
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		APICertIdentity,
		RootIdentity,
		EnforcerCertIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{
	"apicerts":  APICertIdentity,
	"apis":      APICertIdentity,
	"apicert":   APICertIdentity,
	"api":       APICertIdentity,
	"enforcers": EnforcerCertIdentity,
	"enforcer":  EnforcerCertIdentity,
}

// IdentityFromAlias returns the Identity associated to the given alias.
func IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case APICertIdentity:
		return []string{
			"apicerts",
			"apis",
			"apicert",
			"api",
		}
	case RootIdentity:
		return []string{}
	case EnforcerCertIdentity:
		return []string{
			"enforcers",
			"enforcer",
		}
	}

	return nil
}
