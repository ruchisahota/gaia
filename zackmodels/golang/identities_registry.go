package zackmodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(ReportIdentity)
	elemental.RegisterIdentity(RootIdentity)
}

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {
	case ReportIdentity.Name:
		return NewReport()
	case RootIdentity.Name:
		return NewRoot()
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		ReportIdentity,
		RootIdentity,
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
	case ReportIdentity:
		return []string{}
	case RootIdentity:
		return []string{}
	}

	return nil
}
