package types

import (
	"fmt"
	"net/http"

	"go.aporeto.io/elemental"
)

// AuditProfileRuleList is a list of AuditProfileRules
type AuditProfileRuleList []*AuditProfileRule

// Validate will validate all rules in the list
func (a AuditProfileRuleList) Validate() error {

	errs := []error{}

	for _, r := range a {
		if r == nil {
			errs = append(errs, elemental.NewError("Validation Error", "One entry is nil", "gaia", http.StatusUnprocessableEntity))
			continue
		}

		if err := r.Validate(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return elemental.NewErrors(errs...)
	}

	return nil
}

// AuditProfileRuleType specifies the audit rule type.
type AuditProfileRuleType int

// The rule types supported by this package.
const (
	DeleteAllRuleType      AuditProfileRuleType = iota + 1 // DeleteAllRule
	FileWatchRuleType                                      // FileWatchRule
	AppendSyscallRuleType                                  // SyscallRule
	PrependSyscallRuleType                                 // SyscallRule
)

// AuditProfileRule is a generic audit rule
type AuditProfileRule struct {
	Type     AuditProfileRuleType `msgpack:"type" json:"type"`
	Files    *FileWatchRule       `msgpack:"files,omitempty" json:"files,omitempty"`
	Syscalls *SyscallRule         `msgpack:"syscalls,omitempty" json:"syscalls,omitempty"`
}

// Validate validates an audit rule
func (a *AuditProfileRule) Validate() error {

	if a.Type < DeleteAllRuleType || a.Type > PrependSyscallRuleType {
		return elemental.NewError("Validation Error", "Invalid rule type", "elemental", http.StatusUnprocessableEntity)
	}

	switch a.Type {
	case AppendSyscallRuleType, PrependSyscallRuleType:
		if a.Syscalls == nil {
			return elemental.NewError("Validation Error", "Nil syscall rule not allowed", "elemental", http.StatusUnprocessableEntity)
		}
		return a.Syscalls.Validate()
	case FileWatchRuleType:
		if a.Files == nil || a.Syscalls != nil {
			return elemental.NewError("Validation Error", "Nil file watch rule not allowed", "elemental", http.StatusUnprocessableEntity)
		}
		return a.Files.Validate()
	case DeleteAllRuleType:
		if a.Files != nil || a.Syscalls != nil {
			return elemental.NewError("Validation Error", "Delete rule must not have syscall or file", "elemental", http.StatusUnprocessableEntity)
		}
	default:
		return elemental.NewError("Validation Error", "Invalid rule type", "elemental", http.StatusUnprocessableEntity)
	}

	return nil
}

// FileWatchRule is used to audit access to particular files or directories
// that you may be interested in.
type FileWatchRule struct {
	Path        string                 `msgpack:"path" bson:"path" json:"path"`
	Permissions []AuditFilePermissions `msgpack:"permissions" bson:"permissions" json:"permissions"`
}

// Validate validates the filewathc rule.
func (r *FileWatchRule) Validate() error {

	if err := elemental.ValidateMaximumLength("path", r.Path, 128, false); err != nil {
		return err
	}

	if len(r.Permissions) > 4 {
		err := elemental.NewError("Validation Error", "Invalid Permissions", "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": "path"}
		return err
	}

	if len(r.Permissions) == 0 {
		r.Permissions = []AuditFilePermissions{AuditFilePermissionsRead, AuditFilePermissionsWrite, AuditFilePermissionsExecute, AuditFilePermissionsAttribute}
	}

	for _, p := range r.Permissions {
		if err := p.Validate("permissions"); err != nil {
			return err
		}
	}

	return nil
}

// SyscallRule is used to audit invocations of specific syscalls.
type SyscallRule struct {
	List     AuditFilterListType   `msgpack:"list" json:"list"`
	Action   AuditFilterActionType `msgpack:"action" json:"action"`
	Filters  []AuditFilterSpec     `msgpack:"filters" json:"filters"`
	Syscalls []AuditSystemCallType `msgpack:"syscalls" json:"syscalls"`
}

// Validate validates the filewathc rule.
func (r *SyscallRule) Validate() error {

	if err := r.List.Validate("list"); err != nil {
		return err
	}

	if err := r.Action.Validate("action"); err != nil {
		return err
	}

	for _, f := range r.Filters {
		if err := f.Validate(); err != nil {
			return err
		}
	}

	for _, s := range r.Syscalls {
		if err := s.Validate("syscalls"); err != nil {
			return err
		}
	}

	return nil
}

// AuditFilterKind specifies a type of filter to apply to a syscall rule.
type AuditFilterKind uint8

// The type of filters that can be applied.
const (
	AuditFilterKindInterFieldFilter AuditFilterKind = iota + 1 // Inter-field comparison filtering (-C).
	AuditFilteRKindValueFilter                                 // Filtering based on values (-F).
)

// AuditFilterKindFromInt converts an int to an AuditFilterKind.
func AuditFilterKindFromInt(value int) (AuditFilterKind, error) {
	switch value {
	case 1:
		return AuditFilterKindInterFieldFilter, nil
	case 2:
		return AuditFilteRKindValueFilter, nil
	default:
		return AuditFilterKindInterFieldFilter, fmt.Errorf("unable to convert AuditFilterKind %d", value)
	}
}

// AuditFilterSpec defines a filter to apply to a syscall rule.
type AuditFilterSpec struct {
	Kind       AuditFilterKind     `msgpack:"kind" json:"kind"`
	LHS        AuditFilterType     `msgpack:"lhs" json:"lhs"`
	Comparator AuditFilterOperator `msgpack:"comparator" json:"comparator"`
	RHS        string              `msgpack:"rhs" json:"rhs"`
}

// Validate validates and AuditFilterSpec
func (f *AuditFilterSpec) Validate() error {

	if err := f.LHS.Validate(""); err != nil {
		return err
	}

	if err := f.Comparator.Validate(""); err != nil {
		return err
	}

	if f.Kind > AuditFilteRKindValueFilter {
		return elemental.NewError("Validation Error", "Unknown filter type", "elemental", http.StatusUnprocessableEntity)
	}

	return nil
}

func (f *AuditFilterSpec) String() string {
	return fmt.Sprintf("%v %v %v", f.LHS, f.Comparator, f.RHS)
}
