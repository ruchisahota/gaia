package zackmodels

import "github.com/aporeto-inc/elemental"

var relationshipsRegistry elemental.RelationshipsRegistry

// Relationships returns the model relationships.
func Relationships() elemental.RelationshipsRegistry {

	return relationshipsRegistry
}

func init() {
	relationshipsRegistry = elemental.RelationshipsRegistry{}

	//
	// Main Relationship for report
	//
	ReportMainRelationship := &elemental.Relationship{}

	relationshipsRegistry[elemental.IdentityFromName("report")] = ReportMainRelationship

	//
	// Main Relationship for root
	//
	RootMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	// Children relationship for reports in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("report"),
		&elemental.Relationship{
			AllowsCreate: true,
		},
	)

	relationshipsRegistry[elemental.IdentityFromName("root")] = RootMainRelationship

}
