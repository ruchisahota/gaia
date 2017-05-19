package rufusmodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(RemoteProcessorIdentity)
	elemental.RegisterIdentity(RootIdentity)
}

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {
	case RemoteProcessorIdentity.Name:
		return NewRemoteProcessor()
	case RootIdentity.Name:
		return NewRoot()
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		RemoteProcessorIdentity,
		RootIdentity,
	}
}
