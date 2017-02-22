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

// IdentityForAlias returns the Identity associated to the given alias
func IdentityForAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}
