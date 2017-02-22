package midgardmodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(IssueIdentity)
	elemental.RegisterIdentity(AuthIdentity)
}

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {
	case RootIdentity.Name:
		return NewRoot()
	case IssueIdentity.Name:
		return NewIssue()
	case AuthIdentity.Name:
		return NewAuth()
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		RootIdentity,
		IssueIdentity,
		AuthIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{}

// IdentityFromAlias returns the Identity associated to the given alias.
func IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case RootIdentity:
		return []string{}
	case IssueIdentity:
		return []string{}
	case AuthIdentity:
		return []string{}
	}

	return nil
}
