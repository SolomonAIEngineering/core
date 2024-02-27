// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user_service/v1/roles.proto

package user_servicev1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Role with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Role) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Role with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RoleMultiError, or nil if none found.
func (m *Role) ValidateAll() error {
	return m.validate(true)
}

func (m *Role) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Type

	// no validation rules for CanCreateUsers

	// no validation rules for CanReadUsers

	// no validation rules for CanUpdateUsers

	// no validation rules for CanDeleteUsers

	// no validation rules for CanCreateProjects

	// no validation rules for CanReadProjects

	// no validation rules for CanUpdateProjects

	// no validation rules for CanDeleteProjects

	// no validation rules for CanCreateReports

	// no validation rules for CanReadReports

	// no validation rules for CanUpdateReports

	// no validation rules for CanDeleteReports

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RoleValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RoleValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RoleValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RoleValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RoleValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RoleValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetAuditLog() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RoleValidationError{
						field:  fmt.Sprintf("AuditLog[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RoleValidationError{
						field:  fmt.Sprintf("AuditLog[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RoleValidationError{
					field:  fmt.Sprintf("AuditLog[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return RoleMultiError(errors)
	}

	return nil
}

// RoleMultiError is an error wrapping multiple validation errors returned by
// Role.ValidateAll() if the designated constraints aren't met.
type RoleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RoleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RoleMultiError) AllErrors() []error { return m }

// RoleValidationError is the validation error returned by Role.Validate if the
// designated constraints aren't met.
type RoleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoleValidationError) ErrorName() string { return "RoleValidationError" }

// Error satisfies the builtin error interface
func (e RoleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRole.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoleValidationError{}

// Validate checks the field values on RoleAuditEvents with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RoleAuditEvents) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RoleAuditEvents with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RoleAuditEventsMultiError, or nil if none found.
func (m *RoleAuditEvents) ValidateAll() error {
	return m.validate(true)
}

func (m *RoleAuditEvents) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Action

	// no validation rules for PerformedBy

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RoleAuditEventsValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RoleAuditEventsValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RoleAuditEventsValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ClientIp

	// no validation rules for UserAgent

	// no validation rules for Context

	if len(errors) > 0 {
		return RoleAuditEventsMultiError(errors)
	}

	return nil
}

// RoleAuditEventsMultiError is an error wrapping multiple validation errors
// returned by RoleAuditEvents.ValidateAll() if the designated constraints
// aren't met.
type RoleAuditEventsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RoleAuditEventsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RoleAuditEventsMultiError) AllErrors() []error { return m }

// RoleAuditEventsValidationError is the validation error returned by
// RoleAuditEvents.Validate if the designated constraints aren't met.
type RoleAuditEventsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoleAuditEventsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoleAuditEventsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoleAuditEventsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoleAuditEventsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoleAuditEventsValidationError) ErrorName() string { return "RoleAuditEventsValidationError" }

// Error satisfies the builtin error interface
func (e RoleAuditEventsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoleAuditEvents.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoleAuditEventsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoleAuditEventsValidationError{}
