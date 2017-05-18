package squallmodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(APICheckIdentity)
	elemental.RegisterIdentity(NamespaceMappingPolicyIdentity)
	elemental.RegisterIdentity(APIAuthorizationPolicyIdentity)
	elemental.RegisterIdentity(SyscallAccessIdentity)
	elemental.RegisterIdentity(SystemCallIdentity)
	elemental.RegisterIdentity(TagIdentity)
	elemental.RegisterIdentity(EnforcerIdentity)
	elemental.RegisterIdentity(PokeIdentity)
	elemental.RegisterIdentity(FilePathIdentity)
	elemental.RegisterIdentity(FileAccessIdentity)
	elemental.RegisterIdentity(NamespaceIdentity)
	elemental.RegisterIdentity(IntegrationIdentity)
	elemental.RegisterIdentity(PolicyRuleIdentity)
	elemental.RegisterIdentity(TabulatedIdentity)
	elemental.RegisterIdentity(ExternalServiceIdentity)
	elemental.RegisterIdentity(RoleIdentity)
	elemental.RegisterIdentity(PolicyIdentity)
	elemental.RegisterIdentity(FlowStatisticIdentity)
	elemental.RegisterIdentity(SuggestedPolicyIdentity)
	elemental.RegisterIdentity(QuotaPolicyIdentity)
	elemental.RegisterIdentity(FileAccessPolicyIdentity)
	elemental.RegisterIdentity(EnforcerProfileIdentity)
	elemental.RegisterIdentity(RenderedPolicyIdentity)
	elemental.RegisterIdentity(ProcessingUnitIdentity)
	elemental.RegisterIdentity(DependencyMapIdentity)
	elemental.RegisterIdentity(VulnerabilityIdentity)
	elemental.RegisterIdentity(EnforcerProfileMappingPolicyIdentity)
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
	case NamespaceMappingPolicyIdentity.Name:
		return NewNamespaceMappingPolicy()
	case APIAuthorizationPolicyIdentity.Name:
		return NewAPIAuthorizationPolicy()
	case SyscallAccessIdentity.Name:
		return NewSyscallAccess()
	case SystemCallIdentity.Name:
		return NewSystemCall()
	case TagIdentity.Name:
		return NewTag()
	case EnforcerIdentity.Name:
		return NewEnforcer()
	case PokeIdentity.Name:
		return NewPoke()
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
	case TabulatedIdentity.Name:
		return NewTabulated()
	case ExternalServiceIdentity.Name:
		return NewExternalService()
	case RoleIdentity.Name:
		return NewRole()
	case PolicyIdentity.Name:
		return NewPolicy()
	case FlowStatisticIdentity.Name:
		return NewFlowStatistic()
	case SuggestedPolicyIdentity.Name:
		return NewSuggestedPolicy()
	case QuotaPolicyIdentity.Name:
		return NewQuotaPolicy()
	case FileAccessPolicyIdentity.Name:
		return NewFileAccessPolicy()
	case EnforcerProfileIdentity.Name:
		return NewEnforcerProfile()
	case RenderedPolicyIdentity.Name:
		return NewRenderedPolicy()
	case ProcessingUnitIdentity.Name:
		return NewProcessingUnit()
	case DependencyMapIdentity.Name:
		return NewDependencyMap()
	case VulnerabilityIdentity.Name:
		return NewVulnerability()
	case EnforcerProfileMappingPolicyIdentity.Name:
		return NewEnforcerProfileMappingPolicy()
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

// IdentifiableForCategory returns a new instance of the Identifiable for the given category name.
func IdentifiableForCategory(category string) elemental.Identifiable {

	switch category {
	case APICheckIdentity.Category:
		return NewAPICheck()
	case NamespaceMappingPolicyIdentity.Category:
		return NewNamespaceMappingPolicy()
	case APIAuthorizationPolicyIdentity.Category:
		return NewAPIAuthorizationPolicy()
	case SyscallAccessIdentity.Category:
		return NewSyscallAccess()
	case SystemCallIdentity.Category:
		return NewSystemCall()
	case TagIdentity.Category:
		return NewTag()
	case EnforcerIdentity.Category:
		return NewEnforcer()
	case PokeIdentity.Category:
		return NewPoke()
	case FilePathIdentity.Category:
		return NewFilePath()
	case FileAccessIdentity.Category:
		return NewFileAccess()
	case NamespaceIdentity.Category:
		return NewNamespace()
	case IntegrationIdentity.Category:
		return NewIntegration()
	case PolicyRuleIdentity.Category:
		return NewPolicyRule()
	case TabulatedIdentity.Category:
		return NewTabulated()
	case ExternalServiceIdentity.Category:
		return NewExternalService()
	case RoleIdentity.Category:
		return NewRole()
	case PolicyIdentity.Category:
		return NewPolicy()
	case FlowStatisticIdentity.Category:
		return NewFlowStatistic()
	case SuggestedPolicyIdentity.Category:
		return NewSuggestedPolicy()
	case QuotaPolicyIdentity.Category:
		return NewQuotaPolicy()
	case FileAccessPolicyIdentity.Category:
		return NewFileAccessPolicy()
	case EnforcerProfileIdentity.Category:
		return NewEnforcerProfile()
	case RenderedPolicyIdentity.Category:
		return NewRenderedPolicy()
	case ProcessingUnitIdentity.Category:
		return NewProcessingUnit()
	case DependencyMapIdentity.Category:
		return NewDependencyMap()
	case VulnerabilityIdentity.Category:
		return NewVulnerability()
	case EnforcerProfileMappingPolicyIdentity.Category:
		return NewEnforcerProfileMappingPolicy()
	case ActivityIdentity.Category:
		return NewActivity()
	case RootIdentity.Category:
		return NewRoot()
	case NetworkAccessPolicyIdentity.Category:
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
	case NamespaceMappingPolicyIdentity.Name:
		return &NamespaceMappingPoliciesList{}
	case APIAuthorizationPolicyIdentity.Name:
		return &APIAuthorizationPoliciesList{}
	case SyscallAccessIdentity.Name:
		return &SyscallAccessList{}
	case SystemCallIdentity.Name:
		return &SystemCallsList{}
	case TagIdentity.Name:
		return &TagsList{}
	case EnforcerIdentity.Name:
		return &EnforcersList{}
	case PokeIdentity.Name:
		return &PokesList{}
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
	case TabulatedIdentity.Name:
		return &TabulatedsList{}
	case ExternalServiceIdentity.Name:
		return &ExternalServicesList{}
	case RoleIdentity.Name:
		return &RolesList{}
	case PolicyIdentity.Name:
		return &PoliciesList{}
	case FlowStatisticIdentity.Name:
		return &FlowStatisticsList{}
	case SuggestedPolicyIdentity.Name:
		return &SuggestedPoliciesList{}
	case QuotaPolicyIdentity.Name:
		return &QuotaPoliciesList{}
	case FileAccessPolicyIdentity.Name:
		return &FileAccessPoliciesList{}
	case EnforcerProfileIdentity.Name:
		return &EnforcerProfilesList{}
	case RenderedPolicyIdentity.Name:
		return &RenderedPoliciesList{}
	case ProcessingUnitIdentity.Name:
		return &ProcessingUnitsList{}
	case DependencyMapIdentity.Name:
		return &DependencyMapsList{}
	case VulnerabilityIdentity.Name:
		return &VulnerabilitiesList{}
	case EnforcerProfileMappingPolicyIdentity.Name:
		return &EnforcerProfileMappingPoliciesList{}
	case ActivityIdentity.Name:
		return &ActivitiesList{}
	case NetworkAccessPolicyIdentity.Name:
		return &NetworkAccessPoliciesList{}
	default:
		return nil
	}
}

// ContentIdentifiableForCategory returns a new instance of a ContentIdentifiable for the given category name.
func ContentIdentifiableForCategory(category string) elemental.ContentIdentifiable {

	switch category {
	case APICheckIdentity.Category:
		return &APIChecksList{}
	case NamespaceMappingPolicyIdentity.Category:
		return &NamespaceMappingPoliciesList{}
	case APIAuthorizationPolicyIdentity.Category:
		return &APIAuthorizationPoliciesList{}
	case SyscallAccessIdentity.Category:
		return &SyscallAccessList{}
	case SystemCallIdentity.Category:
		return &SystemCallsList{}
	case TagIdentity.Category:
		return &TagsList{}
	case EnforcerIdentity.Category:
		return &EnforcersList{}
	case PokeIdentity.Category:
		return &PokesList{}
	case FilePathIdentity.Category:
		return &FilePathsList{}
	case FileAccessIdentity.Category:
		return &FileAccessList{}
	case NamespaceIdentity.Category:
		return &NamespacesList{}
	case IntegrationIdentity.Category:
		return &IntegrationsList{}
	case PolicyRuleIdentity.Category:
		return &PolicyRulesList{}
	case TabulatedIdentity.Category:
		return &TabulatedsList{}
	case ExternalServiceIdentity.Category:
		return &ExternalServicesList{}
	case RoleIdentity.Category:
		return &RolesList{}
	case PolicyIdentity.Category:
		return &PoliciesList{}
	case FlowStatisticIdentity.Category:
		return &FlowStatisticsList{}
	case SuggestedPolicyIdentity.Category:
		return &SuggestedPoliciesList{}
	case QuotaPolicyIdentity.Category:
		return &QuotaPoliciesList{}
	case FileAccessPolicyIdentity.Category:
		return &FileAccessPoliciesList{}
	case EnforcerProfileIdentity.Category:
		return &EnforcerProfilesList{}
	case RenderedPolicyIdentity.Category:
		return &RenderedPoliciesList{}
	case ProcessingUnitIdentity.Category:
		return &ProcessingUnitsList{}
	case DependencyMapIdentity.Category:
		return &DependencyMapsList{}
	case VulnerabilityIdentity.Category:
		return &VulnerabilitiesList{}
	case EnforcerProfileMappingPolicyIdentity.Category:
		return &EnforcerProfileMappingPoliciesList{}
	case ActivityIdentity.Category:
		return &ActivitiesList{}
	case NetworkAccessPolicyIdentity.Category:
		return &NetworkAccessPoliciesList{}
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		APICheckIdentity,
		NamespaceMappingPolicyIdentity,
		APIAuthorizationPolicyIdentity,
		SyscallAccessIdentity,
		SystemCallIdentity,
		TagIdentity,
		EnforcerIdentity,
		PokeIdentity,
		FilePathIdentity,
		FileAccessIdentity,
		NamespaceIdentity,
		IntegrationIdentity,
		PolicyRuleIdentity,
		TabulatedIdentity,
		ExternalServiceIdentity,
		RoleIdentity,
		PolicyIdentity,
		FlowStatisticIdentity,
		SuggestedPolicyIdentity,
		QuotaPolicyIdentity,
		FileAccessPolicyIdentity,
		EnforcerProfileIdentity,
		RenderedPolicyIdentity,
		ProcessingUnitIdentity,
		DependencyMapIdentity,
		VulnerabilityIdentity,
		EnforcerProfileMappingPolicyIdentity,
		ActivityIdentity,
		RootIdentity,
		NetworkAccessPolicyIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{
	"nsmaps":     NamespaceMappingPolicyIdentity,
	"nsmap":      NamespaceMappingPolicyIdentity,
	"nspolicies": NamespaceMappingPolicyIdentity,
	"nspolicy":   NamespaceMappingPolicyIdentity,
	"apiauth":    APIAuthorizationPolicyIdentity,
	"apiauths":   APIAuthorizationPolicyIdentity,
	"fps":        FilePathIdentity,
	"fp":         FilePathIdentity,
	"ns":         NamespaceIdentity,
	"table":      TabulatedIdentity,
	"tables":     TabulatedIdentity,
	"tabs":       TabulatedIdentity,
	"tab":        TabulatedIdentity,
	"extsrv":     ExternalServiceIdentity,
	"extsrvs":    ExternalServiceIdentity,
	"suggs":      SuggestedPolicyIdentity,
	"sugg":       SuggestedPolicyIdentity,
	"sugpols":    SuggestedPolicyIdentity,
	"sugpol":     SuggestedPolicyIdentity,
	"quotapols":  QuotaPolicyIdentity,
	"quotapol":   QuotaPolicyIdentity,
	"quotas":     QuotaPolicyIdentity,
	"quota":      QuotaPolicyIdentity,
	"profiles":   EnforcerProfileIdentity,
	"profile":    EnforcerProfileIdentity,
	"rpols":      RenderedPolicyIdentity,
	"rpol":       RenderedPolicyIdentity,
	"pus":        ProcessingUnitIdentity,
	"pu":         ProcessingUnitIdentity,
	"depmaps":    DependencyMapIdentity,
	"depmap":     DependencyMapIdentity,
	"vulns":      VulnerabilityIdentity,
	"vul":        VulnerabilityIdentity,
	"srvpol":     EnforcerProfileMappingPolicyIdentity,
	"srvpols":    EnforcerProfileMappingPolicyIdentity,
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
	case NamespaceMappingPolicyIdentity:
		return []string{
			"nsmaps",
			"nsmap",
			"nspolicies",
			"nspolicy",
		}
	case APIAuthorizationPolicyIdentity:
		return []string{
			"apiauth",
			"apiauths",
		}
	case SyscallAccessIdentity:
		return []string{}
	case SystemCallIdentity:
		return []string{}
	case TagIdentity:
		return []string{}
	case EnforcerIdentity:
		return []string{}
	case PokeIdentity:
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
	case TabulatedIdentity:
		return []string{
			"table",
			"tables",
			"tabs",
			"tab",
		}
	case ExternalServiceIdentity:
		return []string{
			"extsrv",
			"extsrvs",
		}
	case RoleIdentity:
		return []string{}
	case PolicyIdentity:
		return []string{}
	case FlowStatisticIdentity:
		return []string{}
	case SuggestedPolicyIdentity:
		return []string{
			"suggs",
			"sugg",
			"sugpols",
			"sugpol",
		}
	case QuotaPolicyIdentity:
		return []string{
			"quotapols",
			"quotapol",
			"quotas",
			"quota",
		}
	case FileAccessPolicyIdentity:
		return []string{}
	case EnforcerProfileIdentity:
		return []string{
			"profiles",
			"profile",
		}
	case RenderedPolicyIdentity:
		return []string{
			"rpols",
			"rpol",
		}
	case ProcessingUnitIdentity:
		return []string{
			"pus",
			"pu",
		}
	case DependencyMapIdentity:
		return []string{
			"depmaps",
			"depmap",
		}
	case VulnerabilityIdentity:
		return []string{
			"vulns",
			"vul",
		}
	case EnforcerProfileMappingPolicyIdentity:
		return []string{
			"srvpol",
			"srvpols",
		}
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
