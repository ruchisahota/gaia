package enum

// EntityStatus defines the status of an entity
type EntityStatus string

const (
	// Active defines that an entity is active
	Active EntityStatus = "Active"

	// Disabled defines that an entity is disabled
	Disabled EntityStatus = "Disabled"

	// Candidate defines that an entity is a candidate and could be either activated or deleted or disabled
	Candidate EntityStatus = "Candidate"
)

// PolicyType defines the type the policy
type PolicyType string

const (
	// APIAuthorization defines that a policy is a APIAuthorization policy
	APIAuthorization PolicyType = "APIAuthorization"

	// Network defines that a policy is a Network policy
	Network PolicyType = "Network"

	// File defines that a policy is a File policy
	File PolicyType = "File"

	// Syscall defines that a policy is a Syscall policy
	Syscall PolicyType = "Syscall"

	// Statistics defines that a policy is a Statistics policy
	Statistics PolicyType = "Statistics"

	// ExtendTags defines that a policy is a ExtendTags policy
	ExtendTags PolicyType = "ExtendTags"

	// NamespaceMapping defines that a policy is a NamespaceMapping policy
	NamespaceMapping PolicyType = "NamespaceMapping"
)

// SeverityLevel defines the severity level of an image or a layer
type SeverityLevel int

const (
	// Unknown defines that severity level is Unknown
	Unknown SeverityLevel = iota
	// Negligible defines that severity level is Negligible
	Negligible
	// Low defines that severity level is Low
	Low
	// Medium defines that severity level is Medium
	Medium
	// High defines that severity level is High
	High
	// Critical defines that severity level is Critical
	Critical
	// Defcon1 defines that severity level is Defcon1
	Defcon1
)

var securityLevels = [...]string{
	"Unknown",
	"Negligible",
	"Low",
	"Medium",
	"High",
	"Critical",
	"Defcon1",
}

func (sl SeverityLevel) String() string { return securityLevels[sl] }

// EntityStatusToString converts EntityStatusToString to its corresponding string value
func EntityStatusToString(entityStatus EntityStatus) string {

	switch entityStatus {
	case Active:
		return "Active"
	case Disabled:
		return "Disabled"
	case Candidate:
		return "Candidate"
	default:
		return ""
	}
}

// PolicyTypeToString converts PolicyTypeToString to its corresponding string value
func PolicyTypeToString(policyType PolicyType) string {

	switch policyType {
	case APIAuthorization:
		return "APIAuthorization"
	case Network:
		return "Network"
	case File:
		return "File"
	case Syscall:
		return "Syscall"
	case Statistics:
		return "Statistics"
	case ExtendTags:
		return "ExtendTags"
	case NamespaceMapping:
		return "NamespaceMapping"
	default:
		return ""
	}
}
