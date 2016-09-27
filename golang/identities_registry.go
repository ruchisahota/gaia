package gaia

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(MyNamespaceIdentity)
	elemental.RegisterIdentity(MapNodeIdentity)
	elemental.RegisterIdentity(LayerIdentity)
	elemental.RegisterIdentity(DependencyMapSubviewIdentity)
	elemental.RegisterIdentity(APIAuthorizationPolicyIdentity)
	elemental.RegisterIdentity(ImageIdentity)
	elemental.RegisterIdentity(ComputedPolicyIdentity)
	elemental.RegisterIdentity(TagIdentity)
	elemental.RegisterIdentity(MapEdgeIdentity)
	elemental.RegisterIdentity(CertificateIdentity)
	elemental.RegisterIdentity(FilePathIdentity)
	elemental.RegisterIdentity(NamespaceIdentity)
	elemental.RegisterIdentity(PolicyRuleIdentity)
	elemental.RegisterIdentity(ExternalServiceIdentity)
	elemental.RegisterIdentity(PolicyIdentity)
	elemental.RegisterIdentity(FlowStatisticIdentity)
	elemental.RegisterIdentity(ComputedDependencyMapViewIdentity)
	elemental.RegisterIdentity(SystemCallIdentity)
	elemental.RegisterIdentity(HealthReportIdentity)
	elemental.RegisterIdentity(FileAccessPolicyIdentity)
	elemental.RegisterIdentity(RenderedPolicyIdentity)
	elemental.RegisterIdentity(NamespaceMappingPolicyIdentity)
	elemental.RegisterIdentity(ProcessingUnitIdentity)
	elemental.RegisterIdentity(AuthenticatorIdentity)
	elemental.RegisterIdentity(DependencyMapViewIdentity)
	elemental.RegisterIdentity(DependencyMapIdentity)
	elemental.RegisterIdentity(VulnerabilityIdentity)
	elemental.RegisterIdentity(ServerIdentity)
	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(NetworkAccessPolicyIdentity)
	elemental.RegisterIdentity(UserIdentity)
}
