package vincemodels

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
	// Main Relationship for account
	//
	AccountMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
	}

	relationshipsRegistry[elemental.IdentityFromName("account")] = AccountMainRelationship

	//
	// Main Relationship for activate
	//
	ActivateMainRelationship := &elemental.Relationship{}

	relationshipsRegistry[elemental.IdentityFromName("activate")] = ActivateMainRelationship

	//
	// Main Relationship for certificate
	//
	CertificateMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	relationshipsRegistry[elemental.IdentityFromName("certificate")] = CertificateMainRelationship

	//
	// Main Relationship for passwordreset
	//
	PasswordResetMainRelationship := &elemental.Relationship{}

	relationshipsRegistry[elemental.IdentityFromName("passwordreset")] = PasswordResetMainRelationship

	//
	// Main Relationship for plan
	//
	PlanMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
	}

	// Children relationship for plans in plan
	PlanMainRelationship.AddChild(
		elemental.IdentityFromName("plan"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)
	relationshipsRegistry[elemental.IdentityFromName("plan")] = PlanMainRelationship

	//
	// Main Relationship for awsaccount
	//
	AWSAccountMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsDelete:   true,
	}

	relationshipsRegistry[elemental.IdentityFromName("awsaccount")] = AWSAccountMainRelationship

	//
	// Main Relationship for root
	//
	RootMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	// Children relationship for accounts in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("account"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for activate in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("activate"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for awsaccounts in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("awsaccount"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for certificates in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("certificate"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for check in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("check"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for passwordreset in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("passwordreset"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for plans in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("plan"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)
	relationshipsRegistry[elemental.IdentityFromName("root")] = RootMainRelationship

	//
	// Main Relationship for check
	//
	CheckMainRelationship := &elemental.Relationship{}

	relationshipsRegistry[elemental.IdentityFromName("check")] = CheckMainRelationship

}
