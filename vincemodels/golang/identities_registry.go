package vincemodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(AccountIdentity)
	elemental.RegisterIdentity(ActivateIdentity)
	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(CheckIdentity)
	elemental.RegisterIdentity(CertificateIdentity)
}

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {
	case AccountIdentity.Name:
		return NewAccount()
	case ActivateIdentity.Name:
		return NewActivate()
	case RootIdentity.Name:
		return NewRoot()
	case CheckIdentity.Name:
		return NewCheck()
	case CertificateIdentity.Name:
		return NewCertificate()
	default:
		return nil
	}
}

// ContentIdentifiableForIdentity returns a new instance of a ContentIdentifiable for the given identity name.
func ContentIdentifiableForIdentity(identity string) elemental.ContentIdentifiable {

	switch identity {
	case AccountIdentity.Name:
		return &AccountsList{}
	case ActivateIdentity.Name:
		return &ActivatesList{}
	case CheckIdentity.Name:
		return &ChecksList{}
	case CertificateIdentity.Name:
		return &CertificatesList{}
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		AccountIdentity,
		ActivateIdentity,
		RootIdentity,
		CheckIdentity,
		CertificateIdentity,
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
	case AccountIdentity:
		return []string{}
	case ActivateIdentity:
		return []string{}
	case RootIdentity:
		return []string{}
	case CheckIdentity:
		return []string{}
	case CertificateIdentity:
		return []string{}
	}

	return nil
}
