package barretmodels

import "github.com/aporeto-inc/elemental"

const nodocString = "[nodoc]" // nolint: varcheck

var relationshipsRegistry elemental.RelationshipsRegistry

// Relationships returns the model relationships.
func Relationships() elemental.RelationshipsRegistry {

	return relationshipsRegistry
}

func init() {
	relationshipsRegistry = elemental.RelationshipsRegistry{}

	relationshipsRegistry[APICertIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
	}
	relationshipsRegistry[EnforcerCertIdentity] = &elemental.Relationship{
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsDelete: map[string]bool{
			"root": true,
		},
		AllowsRetrieve: map[string]bool{
			"root": true,
		},
		AllowsRetrieveMany: map[string]bool{
			"root": true,
		},
		AllowsInfo: map[string]bool{
			"root": true,
		},
	}
	relationshipsRegistry[RootIdentity] = &elemental.Relationship{}
	relationshipsRegistry[TriremeCertIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
	}
}
