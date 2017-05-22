package squallmodels

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
	// Main Relationship for apicheck
	//
	APICheckMainRelationship := &elemental.Relationship{}

	relationshipsRegistry[elemental.IdentityFromName("apicheck")] = APICheckMainRelationship

	//
	// Main Relationship for hookpolicy
	//
	HookPolicyMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	relationshipsRegistry[elemental.IdentityFromName("hookpolicy")] = HookPolicyMainRelationship

	//
	// Main Relationship for apiauthorizationpolicy
	//
	APIAuthorizationPolicyMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	relationshipsRegistry[elemental.IdentityFromName("apiauthorizationpolicy")] = APIAuthorizationPolicyMainRelationship

	//
	// Main Relationship for syscallaccess
	//
	SyscallAccessMainRelationship := &elemental.Relationship{}

	relationshipsRegistry[elemental.IdentityFromName("syscallaccess")] = SyscallAccessMainRelationship

	//
	// Main Relationship for systemcall
	//
	SystemCallMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	relationshipsRegistry[elemental.IdentityFromName("systemcall")] = SystemCallMainRelationship

	//
	// Main Relationship for tag
	//
	TagMainRelationship := &elemental.Relationship{}

	relationshipsRegistry[elemental.IdentityFromName("tag")] = TagMainRelationship

	//
	// Main Relationship for enforcer
	//
	EnforcerMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	// Children relationship for enforcerprofiles in enforcer
	EnforcerMainRelationship.AddChild(
		elemental.IdentityFromName("enforcerprofile"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for poke in enforcer
	EnforcerMainRelationship.AddChild(
		elemental.IdentityFromName("poke"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)
	relationshipsRegistry[elemental.IdentityFromName("enforcer")] = EnforcerMainRelationship

	//
	// Main Relationship for poke
	//
	PokeMainRelationship := &elemental.Relationship{}

	relationshipsRegistry[elemental.IdentityFromName("poke")] = PokeMainRelationship

	//
	// Main Relationship for filepath
	//
	FilePathMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	relationshipsRegistry[elemental.IdentityFromName("filepath")] = FilePathMainRelationship

	//
	// Main Relationship for fileaccess
	//
	FileAccessMainRelationship := &elemental.Relationship{}

	relationshipsRegistry[elemental.IdentityFromName("fileaccess")] = FileAccessMainRelationship

	//
	// Main Relationship for namespace
	//
	NamespaceMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	relationshipsRegistry[elemental.IdentityFromName("namespace")] = NamespaceMainRelationship

	//
	// Main Relationship for integration
	//
	IntegrationMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	relationshipsRegistry[elemental.IdentityFromName("integration")] = IntegrationMainRelationship

	//
	// Main Relationship for policyrule
	//
	PolicyRuleMainRelationship := &elemental.Relationship{}

	relationshipsRegistry[elemental.IdentityFromName("policyrule")] = PolicyRuleMainRelationship

	//
	// Main Relationship for externalservice
	//
	ExternalServiceMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	relationshipsRegistry[elemental.IdentityFromName("externalservice")] = ExternalServiceMainRelationship

	//
	// Main Relationship for role
	//
	RoleMainRelationship := &elemental.Relationship{}

	relationshipsRegistry[elemental.IdentityFromName("role")] = RoleMainRelationship

	//
	// Main Relationship for policy
	//
	PolicyMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsDelete:   true,
	}

	relationshipsRegistry[elemental.IdentityFromName("policy")] = PolicyMainRelationship

	//
	// Main Relationship for flowstatistic
	//
	FlowStatisticMainRelationship := &elemental.Relationship{}

	relationshipsRegistry[elemental.IdentityFromName("flowstatistic")] = FlowStatisticMainRelationship

	//
	// Main Relationship for suggestedpolicy
	//
	SuggestedPolicyMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	relationshipsRegistry[elemental.IdentityFromName("suggestedpolicy")] = SuggestedPolicyMainRelationship

	//
	// Main Relationship for tabulation
	//
	TabulationMainRelationship := &elemental.Relationship{}

	relationshipsRegistry[elemental.IdentityFromName("tabulation")] = TabulationMainRelationship

	//
	// Main Relationship for quotapolicy
	//
	QuotaPolicyMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	relationshipsRegistry[elemental.IdentityFromName("quotapolicy")] = QuotaPolicyMainRelationship

	//
	// Main Relationship for fileaccesspolicy
	//
	FileAccessPolicyMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	relationshipsRegistry[elemental.IdentityFromName("fileaccesspolicy")] = FileAccessPolicyMainRelationship

	//
	// Main Relationship for enforcerprofile
	//
	EnforcerProfileMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	relationshipsRegistry[elemental.IdentityFromName("enforcerprofile")] = EnforcerProfileMainRelationship

	//
	// Main Relationship for renderedpolicy
	//
	RenderedPolicyMainRelationship := &elemental.Relationship{}

	relationshipsRegistry[elemental.IdentityFromName("renderedpolicy")] = RenderedPolicyMainRelationship

	//
	// Main Relationship for namespacemappingpolicy
	//
	NamespaceMappingPolicyMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	relationshipsRegistry[elemental.IdentityFromName("namespacemappingpolicy")] = NamespaceMappingPolicyMainRelationship

	//
	// Main Relationship for processingunit
	//
	ProcessingUnitMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	// Children relationship for fileaccesses in processingunit
	ProcessingUnitMainRelationship.AddChild(
		elemental.IdentityFromName("fileaccess"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for renderedpolicies in processingunit
	ProcessingUnitMainRelationship.AddChild(
		elemental.IdentityFromName("renderedpolicy"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for syscallaccesses in processingunit
	ProcessingUnitMainRelationship.AddChild(
		elemental.IdentityFromName("syscallaccess"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for vulnerabilities in processingunit
	ProcessingUnitMainRelationship.AddChild(
		elemental.IdentityFromName("vulnerability"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)
	relationshipsRegistry[elemental.IdentityFromName("processingunit")] = ProcessingUnitMainRelationship

	//
	// Main Relationship for dependencymap
	//
	DependencyMapMainRelationship := &elemental.Relationship{}

	relationshipsRegistry[elemental.IdentityFromName("dependencymap")] = DependencyMapMainRelationship

	//
	// Main Relationship for vulnerability
	//
	VulnerabilityMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
	}

	// Children relationship for processingunits in vulnerability
	VulnerabilityMainRelationship.AddChild(
		elemental.IdentityFromName("processingunit"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)
	relationshipsRegistry[elemental.IdentityFromName("vulnerability")] = VulnerabilityMainRelationship

	//
	// Main Relationship for enforcerprofilemappingpolicy
	//
	EnforcerProfileMappingPolicyMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	relationshipsRegistry[elemental.IdentityFromName("enforcerprofilemappingpolicy")] = EnforcerProfileMappingPolicyMainRelationship

	//
	// Main Relationship for activity
	//
	ActivityMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
	}

	relationshipsRegistry[elemental.IdentityFromName("activity")] = ActivityMainRelationship

	//
	// Main Relationship for root
	//
	RootMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
	}

	// Children relationship for activities in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("activity"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for apiauthorizationpolicies in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("apiauthorizationpolicy"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for apichecks in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("apicheck"),
		&elemental.Relationship{
			AllowsCreate: true,
		},
	)

	// Children relationship for dependencymaps in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("dependencymap"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for enforcers in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("enforcer"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for enforcerprofiles in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("enforcerprofile"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for enforcerprofilemappingpolicies in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("enforcerprofilemappingpolicy"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for externalservices in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("externalservice"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for fileaccesspolicies in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("fileaccesspolicy"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for filepaths in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("filepath"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for flowstatistics in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("flowstatistic"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for hookpolicies in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("hookpolicy"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for integrations in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("integration"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for namespaces in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("namespace"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for namespacemappingpolicies in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("namespacemappingpolicy"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for networkaccesspolicies in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("networkaccesspolicy"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for policies in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("policy"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for processingunits in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("processingunit"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for quotapolicies in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("quotapolicy"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for roles in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("role"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for suggestedpolicies in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("suggestedpolicy"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for systemcalls in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("systemcall"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for tabulations in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("tabulation"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for tags in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("tag"),
		&elemental.Relationship{
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)

	// Children relationship for vulnerabilities in root
	RootMainRelationship.AddChild(
		elemental.IdentityFromName("vulnerability"),
		&elemental.Relationship{
			AllowsCreate:       true,
			AllowsRetrieveMany: true,
			AllowsInfo:         true,
		},
	)
	relationshipsRegistry[elemental.IdentityFromName("root")] = RootMainRelationship

	//
	// Main Relationship for networkaccesspolicy
	//
	NetworkAccessPolicyMainRelationship := &elemental.Relationship{
		AllowsRetrieve: true,
		AllowsUpdate:   true,
		AllowsDelete:   true,
	}

	relationshipsRegistry[elemental.IdentityFromName("networkaccesspolicy")] = NetworkAccessPolicyMainRelationship

}
