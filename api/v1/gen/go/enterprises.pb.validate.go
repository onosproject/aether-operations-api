// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: enterprises.proto

package v1

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

// Validate checks the field values on Enterprise with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Enterprise) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Enterprise with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EnterpriseMultiError, or
// nil if none found.
func (m *Enterprise) ValidateAll() error {
	return m.validate(true)
}

func (m *Enterprise) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ID

	// no validation rules for Description

	// no validation rules for DisplayName

	for idx, item := range m.GetApplications() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EnterpriseValidationError{
						field:  fmt.Sprintf("Applications[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EnterpriseValidationError{
						field:  fmt.Sprintf("Applications[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EnterpriseValidationError{
					field:  fmt.Sprintf("Applications[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return EnterpriseMultiError(errors)
	}

	return nil
}

// EnterpriseMultiError is an error wrapping multiple validation errors
// returned by Enterprise.ValidateAll() if the designated constraints aren't met.
type EnterpriseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EnterpriseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EnterpriseMultiError) AllErrors() []error { return m }

// EnterpriseValidationError is the validation error returned by
// Enterprise.Validate if the designated constraints aren't met.
type EnterpriseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnterpriseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnterpriseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnterpriseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnterpriseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnterpriseValidationError) ErrorName() string { return "EnterpriseValidationError" }

// Error satisfies the builtin error interface
func (e EnterpriseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnterprise.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnterpriseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnterpriseValidationError{}

// Validate checks the field values on Enterprises with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Enterprises) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Enterprises with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EnterprisesMultiError, or
// nil if none found.
func (m *Enterprises) ValidateAll() error {
	return m.validate(true)
}

func (m *Enterprises) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetEnterprises() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EnterprisesValidationError{
						field:  fmt.Sprintf("Enterprises[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EnterprisesValidationError{
						field:  fmt.Sprintf("Enterprises[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EnterprisesValidationError{
					field:  fmt.Sprintf("Enterprises[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return EnterprisesMultiError(errors)
	}

	return nil
}

// EnterprisesMultiError is an error wrapping multiple validation errors
// returned by Enterprises.ValidateAll() if the designated constraints aren't met.
type EnterprisesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EnterprisesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EnterprisesMultiError) AllErrors() []error { return m }

// EnterprisesValidationError is the validation error returned by
// Enterprises.Validate if the designated constraints aren't met.
type EnterprisesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnterprisesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnterprisesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnterprisesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnterprisesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnterprisesValidationError) ErrorName() string { return "EnterprisesValidationError" }

// Error satisfies the builtin error interface
func (e EnterprisesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnterprises.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnterprisesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnterprisesValidationError{}
