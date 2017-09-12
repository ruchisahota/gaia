package barretmodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(APICertIdentity)
	elemental.RegisterIdentity(TokenIdentity)
	elemental.RegisterIdentity(RootIdentity)
}

// ModelVersion returns the current version of the model
func ModelVersion() float64 { return 1 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {
	case APICertIdentity.Name:
		return NewAPICert()
	case TokenIdentity.Name:
		return NewToken()
	case RootIdentity.Name:
		return NewRoot()
	default:
		return nil
	}
}

// IdentifiableForCategory returns a new instance of the Identifiable for the given category name.
func IdentifiableForCategory(category string) elemental.Identifiable {

	switch category {
	case APICertIdentity.Category:
		return NewAPICert()
	case TokenIdentity.Category:
		return NewToken()
	case RootIdentity.Category:
		return NewRoot()
	default:
		return nil
	}
}

// ContentIdentifiableForIdentity returns a new instance of a ContentIdentifiable for the given identity name.
func ContentIdentifiableForIdentity(identity string) elemental.ContentIdentifiable {

	switch identity {
	case APICertIdentity.Name:
		return &APICertsList{}
	case TokenIdentity.Name:
		return &TokensList{}
	default:
		return nil
	}
}

// ContentIdentifiableForCategory returns a new instance of a ContentIdentifiable for the given category name.
func ContentIdentifiableForCategory(category string) elemental.ContentIdentifiable {

	switch category {
	case APICertIdentity.Category:
		return &APICertsList{}
	case TokenIdentity.Category:
		return &TokensList{}
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		APICertIdentity,
		TokenIdentity,
		RootIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{
	"apicerts": APICertIdentity,
	"apis":     APICertIdentity,
	"apicert":  APICertIdentity,
	"api":      APICertIdentity,
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
	case TokenIdentity:
		return []string{}
	case RootIdentity:
		return []string{}
	}

	return nil
}
