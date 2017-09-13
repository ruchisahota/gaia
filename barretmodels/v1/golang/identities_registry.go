package barretmodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(TokenIdentity)
	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(CheckIdentity)
	elemental.RegisterIdentity(CertificateIdentity)
	elemental.RegisterIdentity(RevocationIdentity)
}

// ModelVersion returns the current version of the model
func ModelVersion() float64 { return 1 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {
	case TokenIdentity.Name:
		return NewToken()
	case RootIdentity.Name:
		return NewRoot()
	case CheckIdentity.Name:
		return NewCheck()
	case CertificateIdentity.Name:
		return NewCertificate()
	case RevocationIdentity.Name:
		return NewRevocation()
	default:
		return nil
	}
}

// IdentifiableForCategory returns a new instance of the Identifiable for the given category name.
func IdentifiableForCategory(category string) elemental.Identifiable {

	switch category {
	case TokenIdentity.Category:
		return NewToken()
	case RootIdentity.Category:
		return NewRoot()
	case CheckIdentity.Category:
		return NewCheck()
	case CertificateIdentity.Category:
		return NewCertificate()
	case RevocationIdentity.Category:
		return NewRevocation()
	default:
		return nil
	}
}

// ContentIdentifiableForIdentity returns a new instance of a ContentIdentifiable for the given identity name.
func ContentIdentifiableForIdentity(identity string) elemental.ContentIdentifiable {

	switch identity {
	case TokenIdentity.Name:
		return &TokensList{}
	case CheckIdentity.Name:
		return &ChecksList{}
	case CertificateIdentity.Name:
		return &CertificatesList{}
	case RevocationIdentity.Name:
		return &RevocationsList{}
	default:
		return nil
	}
}

// ContentIdentifiableForCategory returns a new instance of a ContentIdentifiable for the given category name.
func ContentIdentifiableForCategory(category string) elemental.ContentIdentifiable {

	switch category {
	case TokenIdentity.Category:
		return &TokensList{}
	case CheckIdentity.Category:
		return &ChecksList{}
	case CertificateIdentity.Category:
		return &CertificatesList{}
	case RevocationIdentity.Category:
		return &RevocationsList{}
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		TokenIdentity,
		RootIdentity,
		CheckIdentity,
		CertificateIdentity,
		RevocationIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{
	"api":      CertificateIdentity,
	"apicert":  CertificateIdentity,
	"apis":     CertificateIdentity,
	"apicerts": CertificateIdentity,
}

// IdentityFromAlias returns the Identity associated to the given alias.
func IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case TokenIdentity:
		return []string{}
	case RootIdentity:
		return []string{}
	case CheckIdentity:
		return []string{}
	case CertificateIdentity:
		return []string{
			"api",
			"apicert",
			"apis",
			"apicerts",
		}
	case RevocationIdentity:
		return []string{}
	}

	return nil
}
