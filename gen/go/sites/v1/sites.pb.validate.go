// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: sites/v1/sites.proto

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

// Validate checks the field values on Site with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Site) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Site with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SiteMultiError, or nil if none found.
func (m *Site) ValidateAll() error {
	return m.validate(true)
}

func (m *Site) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SiteId

	// no validation rules for Name

	// no validation rules for Description

	for idx, item := range m.GetDevices() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SiteValidationError{
						field:  fmt.Sprintf("Devices[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SiteValidationError{
						field:  fmt.Sprintf("Devices[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SiteValidationError{
					field:  fmt.Sprintf("Devices[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetSlices() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SiteValidationError{
						field:  fmt.Sprintf("Slices[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SiteValidationError{
						field:  fmt.Sprintf("Slices[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SiteValidationError{
					field:  fmt.Sprintf("Slices[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SiteMultiError(errors)
	}

	return nil
}

// SiteMultiError is an error wrapping multiple validation errors returned by
// Site.ValidateAll() if the designated constraints aren't met.
type SiteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SiteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SiteMultiError) AllErrors() []error { return m }

// SiteValidationError is the validation error returned by Site.Validate if the
// designated constraints aren't met.
type SiteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SiteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SiteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SiteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SiteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SiteValidationError) ErrorName() string { return "SiteValidationError" }

// Error satisfies the builtin error interface
func (e SiteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSite.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SiteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SiteValidationError{}

// Validate checks the field values on GetSitesResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetSitesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSitesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSitesResponseMultiError, or nil if none found.
func (m *GetSitesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSitesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetSites() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetSitesResponseValidationError{
						field:  fmt.Sprintf("Sites[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetSitesResponseValidationError{
						field:  fmt.Sprintf("Sites[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetSitesResponseValidationError{
					field:  fmt.Sprintf("Sites[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetSitesResponseMultiError(errors)
	}

	return nil
}

// GetSitesResponseMultiError is an error wrapping multiple validation errors
// returned by GetSitesResponse.ValidateAll() if the designated constraints
// aren't met.
type GetSitesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSitesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSitesResponseMultiError) AllErrors() []error { return m }

// GetSitesResponseValidationError is the validation error returned by
// GetSitesResponse.Validate if the designated constraints aren't met.
type GetSitesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSitesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSitesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSitesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSitesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSitesResponseValidationError) ErrorName() string { return "GetSitesResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetSitesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSitesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSitesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSitesResponseValidationError{}

// Validate checks the field values on GetSitesRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetSitesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSitesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSitesRequestMultiError, or nil if none found.
func (m *GetSitesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSitesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for EnterpriseId

	if len(errors) > 0 {
		return GetSitesRequestMultiError(errors)
	}

	return nil
}

// GetSitesRequestMultiError is an error wrapping multiple validation errors
// returned by GetSitesRequest.ValidateAll() if the designated constraints
// aren't met.
type GetSitesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSitesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSitesRequestMultiError) AllErrors() []error { return m }

// GetSitesRequestValidationError is the validation error returned by
// GetSitesRequest.Validate if the designated constraints aren't met.
type GetSitesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSitesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSitesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSitesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSitesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSitesRequestValidationError) ErrorName() string { return "GetSitesRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetSitesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSitesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSitesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSitesRequestValidationError{}
