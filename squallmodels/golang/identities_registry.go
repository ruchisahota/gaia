package squallmodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(APICheckIdentity)
	elemental.RegisterIdentity(MapNodeIdentity)
	elemental.RegisterIdentity(DependencyMapSubviewIdentity)
	elemental.RegisterIdentity(APIAuthorizationPolicyIdentity)
	elemental.RegisterIdentity(SyscallAccessIdentity)
	elemental.RegisterIdentity(ComputedPolicyIdentity)
	elemental.RegisterIdentity(TagIdentity)
	elemental.RegisterIdentity(MapEdgeIdentity)
	elemental.RegisterIdentity(FilePathIdentity)
	elemental.RegisterIdentity(FileAccessIdentity)
	elemental.RegisterIdentity(NamespaceIdentity)
	elemental.RegisterIdentity(IntegrationIdentity)
	elemental.RegisterIdentity(PolicyRuleIdentity)
	elemental.RegisterIdentity(ExternalServiceIdentity)
	elemental.RegisterIdentity(PolicyIdentity)
	elemental.RegisterIdentity(FlowStatisticIdentity)
	elemental.RegisterIdentity(ServerProfileIdentity)
	elemental.RegisterIdentity(ServerPolicyIdentity)
	elemental.RegisterIdentity(ComputedDependencyMapViewIdentity)
	elemental.RegisterIdentity(SystemCallIdentity)
	elemental.RegisterIdentity(FileAccessPolicyIdentity)
	elemental.RegisterIdentity(RenderedPolicyIdentity)
	elemental.RegisterIdentity(NamespaceMappingPolicyIdentity)
	elemental.RegisterIdentity(ProcessingUnitIdentity)
	elemental.RegisterIdentity(DependencyMapViewIdentity)
	elemental.RegisterIdentity(DependencyMapIdentity)
	elemental.RegisterIdentity(VulnerabilityIdentity)
	elemental.RegisterIdentity(ServerIdentity)
	elemental.RegisterIdentity(ActivityIdentity)
	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(NetworkAccessPolicyIdentity)
}

// ModelVersion returns the current version of the model
func ModelVersion() float64 { return 1.0 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {
	case APICheckIdentity.Name:
		return NewAPICheck()
	case MapNodeIdentity.Name:
		return NewMapNode()
	case DependencyMapSubviewIdentity.Name:
		return NewDependencyMapSubview()
	case APIAuthorizationPolicyIdentity.Name:
		return NewAPIAuthorizationPolicy()
	case SyscallAccessIdentity.Name:
		return NewSyscallAccess()
	case ComputedPolicyIdentity.Name:
		return NewComputedPolicy()
	case TagIdentity.Name:
		return NewTag()
	case MapEdgeIdentity.Name:
		return NewMapEdge()
	case FilePathIdentity.Name:
		return NewFilePath()
	case FileAccessIdentity.Name:
		return NewFileAccess()
	case NamespaceIdentity.Name:
		return NewNamespace()
	case IntegrationIdentity.Name:
		return NewIntegration()
	case PolicyRuleIdentity.Name:
		return NewPolicyRule()
	case ExternalServiceIdentity.Name:
		return NewExternalService()
	case PolicyIdentity.Name:
		return NewPolicy()
	case FlowStatisticIdentity.Name:
		return NewFlowStatistic()
	case ServerProfileIdentity.Name:
		return NewServerProfile()
	case ServerPolicyIdentity.Name:
		return NewServerPolicy()
	case ComputedDependencyMapViewIdentity.Name:
		return NewComputedDependencyMapView()
	case SystemCallIdentity.Name:
		return NewSystemCall()
	case FileAccessPolicyIdentity.Name:
		return NewFileAccessPolicy()
	case RenderedPolicyIdentity.Name:
		return NewRenderedPolicy()
	case NamespaceMappingPolicyIdentity.Name:
		return NewNamespaceMappingPolicy()
	case ProcessingUnitIdentity.Name:
		return NewProcessingUnit()
	case DependencyMapViewIdentity.Name:
		return NewDependencyMapView()
	case DependencyMapIdentity.Name:
		return NewDependencyMap()
	case VulnerabilityIdentity.Name:
		return NewVulnerability()
	case ServerIdentity.Name:
		return NewServer()
	case ActivityIdentity.Name:
		return NewActivity()
	case RootIdentity.Name:
		return NewRoot()
	case NetworkAccessPolicyIdentity.Name:
		return NewNetworkAccessPolicy()
	default:
		return nil
	}
}

// ContentIdentifiableForIdentity returns a new instance of a ContentIdentifiable for the given identity name.
func ContentIdentifiableForIdentity(identity string) elemental.ContentIdentifiable {

	switch identity {
	case APICheckIdentity.Name:
		return &APIChecksList{}
	case MapNodeIdentity.Name:
		return &MapNodesList{}
	case DependencyMapSubviewIdentity.Name:
		return &DependencyMapSubviewsList{}
	case APIAuthorizationPolicyIdentity.Name:
		return &APIAuthorizationPoliciesList{}
	case SyscallAccessIdentity.Name:
		return &SyscallAccessList{}
	case ComputedPolicyIdentity.Name:
		return &ComputedPoliciesList{}
	case TagIdentity.Name:
		return &TagsList{}
	case MapEdgeIdentity.Name:
		return &MapEdgesList{}
	case FilePathIdentity.Name:
		return &FilePathsList{}
	case FileAccessIdentity.Name:
		return &FileAccessList{}
	case NamespaceIdentity.Name:
		return &NamespacesList{}
	case IntegrationIdentity.Name:
		return &IntegrationsList{}
	case PolicyRuleIdentity.Name:
		return &PolicyRulesList{}
	case ExternalServiceIdentity.Name:
		return &ExternalServicesList{}
	case PolicyIdentity.Name:
		return &PoliciesList{}
	case FlowStatisticIdentity.Name:
		return &FlowStatisticsList{}
	case ServerProfileIdentity.Name:
		return &ServerProfilesList{}
	case ServerPolicyIdentity.Name:
		return &ServerPoliciesList{}
	case ComputedDependencyMapViewIdentity.Name:
		return &ComputedDependencyMapViewsList{}
	case SystemCallIdentity.Name:
		return &SystemCallsList{}
	case FileAccessPolicyIdentity.Name:
		return &FileAccessPoliciesList{}
	case RenderedPolicyIdentity.Name:
		return &RenderedPoliciesList{}
	case NamespaceMappingPolicyIdentity.Name:
		return &NamespaceMappingPoliciesList{}
	case ProcessingUnitIdentity.Name:
		return &ProcessingUnitsList{}
	case DependencyMapViewIdentity.Name:
		return &DependencyMapViewsList{}
	case DependencyMapIdentity.Name:
		return &DependencyMapsList{}
	case VulnerabilityIdentity.Name:
		return &VulnerabilitiesList{}
	case ServerIdentity.Name:
		return &ServersList{}
	case ActivityIdentity.Name:
		return &ActivitiesList{}
	case NetworkAccessPolicyIdentity.Name:
		return &NetworkAccessPoliciesList{}
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		APICheckIdentity,
		MapNodeIdentity,
		DependencyMapSubviewIdentity,
		APIAuthorizationPolicyIdentity,
		SyscallAccessIdentity,
		ComputedPolicyIdentity,
		TagIdentity,
		MapEdgeIdentity,
		FilePathIdentity,
		FileAccessIdentity,
		NamespaceIdentity,
		IntegrationIdentity,
		PolicyRuleIdentity,
		ExternalServiceIdentity,
		PolicyIdentity,
		FlowStatisticIdentity,
		ServerProfileIdentity,
		ServerPolicyIdentity,
		ComputedDependencyMapViewIdentity,
		SystemCallIdentity,
		FileAccessPolicyIdentity,
		RenderedPolicyIdentity,
		NamespaceMappingPolicyIdentity,
		ProcessingUnitIdentity,
		DependencyMapViewIdentity,
		DependencyMapIdentity,
		VulnerabilityIdentity,
		ServerIdentity,
		ActivityIdentity,
		RootIdentity,
		NetworkAccessPolicyIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{
	"apiauths":   APIAuthorizationPolicyIdentity,
	"apiauth":    APIAuthorizationPolicyIdentity,
	"fps":        FilePathIdentity,
	"fp":         FilePathIdentity,
	"ns":         NamespaceIdentity,
	"extsrvs":    ExternalServiceIdentity,
	"extsrv":     ExternalServiceIdentity,
	"srvpols":    ServerPolicyIdentity,
	"srvpol":     ServerPolicyIdentity,
	"rpols":      RenderedPolicyIdentity,
	"rpol":       RenderedPolicyIdentity,
	"nsmaps":     NamespaceMappingPolicyIdentity,
	"nsmap":      NamespaceMappingPolicyIdentity,
	"nspolicies": NamespaceMappingPolicyIdentity,
	"nspolicy":   NamespaceMappingPolicyIdentity,
	"pus":        ProcessingUnitIdentity,
	"pu":         ProcessingUnitIdentity,
	"vulns":      VulnerabilityIdentity,
	"vul":        VulnerabilityIdentity,
	"netpols":    NetworkAccessPolicyIdentity,
	"netpol":     NetworkAccessPolicyIdentity,
}

// IdentityFromAlias returns the Identity associated to the given alias.
func IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case APICheckIdentity:
		return []string{}
	case MapNodeIdentity:
		return []string{}
	case DependencyMapSubviewIdentity:
		return []string{}
	case APIAuthorizationPolicyIdentity:
		return []string{
			"apiauths",
			"apiauth",
		}
	case SyscallAccessIdentity:
		return []string{}
	case ComputedPolicyIdentity:
		return []string{}
	case TagIdentity:
		return []string{}
	case MapEdgeIdentity:
		return []string{}
	case FilePathIdentity:
		return []string{
			"fps",
			"fp",
		}
	case FileAccessIdentity:
		return []string{}
	case NamespaceIdentity:
		return []string{
			"ns",
		}
	case IntegrationIdentity:
		return []string{}
	case PolicyRuleIdentity:
		return []string{}
	case ExternalServiceIdentity:
		return []string{
			"extsrvs",
			"extsrv",
		}
	case PolicyIdentity:
		return []string{}
	case FlowStatisticIdentity:
		return []string{}
	case ServerProfileIdentity:
		return []string{}
	case ServerPolicyIdentity:
		return []string{
			"srvpols",
			"srvpol",
		}
	case ComputedDependencyMapViewIdentity:
		return []string{}
	case SystemCallIdentity:
		return []string{}
	case FileAccessPolicyIdentity:
		return []string{}
	case RenderedPolicyIdentity:
		return []string{
			"rpols",
			"rpol",
		}
	case NamespaceMappingPolicyIdentity:
		return []string{
			"nsmaps",
			"nsmap",
			"nspolicies",
			"nspolicy",
		}
	case ProcessingUnitIdentity:
		return []string{
			"pus",
			"pu",
		}
	case DependencyMapViewIdentity:
		return []string{}
	case DependencyMapIdentity:
		return []string{}
	case VulnerabilityIdentity:
		return []string{
			"vulns",
			"vul",
		}
	case ServerIdentity:
		return []string{}
	case ActivityIdentity:
		return []string{}
	case RootIdentity:
		return []string{}
	case NetworkAccessPolicyIdentity:
		return []string{
			"netpols",
			"netpol",
		}
	}

	return nil
}
