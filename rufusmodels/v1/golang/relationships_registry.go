package rufusmodels

import "github.com/aporeto-inc/elemental"

const nodocString = "[nodoc]"

var relationshipsRegistry elemental.RelationshipsRegistry

// Relationships returns the model relationships.
func Relationships() elemental.RelationshipsRegistry {

	return relationshipsRegistry
}

func init() {
	relationshipsRegistry = elemental.RelationshipsRegistry{}

	//
	// Main Relationship for remoteprocessor
	//
	RemoteProcessorMainRelationship := &elemental.Relationship{}

	relationshipsRegistry[elemental.IdentityFromName("remoteprocessor")] = RemoteProcessorMainRelationship

	//
	// Main Relationship for root
	//
	RootMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
	}

	// Children relationship for remoteprocessors in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("remoteprocessor"),
		&elemental.Relationship{
			AllowsCreate: true,
		},
	)
	relationshipsRegistry[elemental.IdentityFromName("root")] = RootMainRelationship

}
