// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: contrib/envoy/extensions/filters/http/sentinel/v3/sentinel_new.proto

package sentinelv3

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

// Validate checks the field values on SentinelNew with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SentinelNew) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SentinelNew with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SentinelNewMultiError, or
// nil if none found.
func (m *SentinelNew) ValidateAll() error {
	return m.validate(true)
}

func (m *SentinelNew) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCommonConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SentinelNewValidationError{
					field:  "CommonConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SentinelNewValidationError{
					field:  "CommonConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCommonConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SentinelNewValidationError{
				field:  "CommonConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		err := SentinelNewValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SentinelNewValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SentinelNewValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SentinelNewValidationError{
				field:  "Status",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetResponseHeadersToAdd()) > 10 {
		err := SentinelNewValidationError{
			field:  "ResponseHeadersToAdd",
			reason: "value must contain no more than 10 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetResponseHeadersToAdd() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SentinelNewValidationError{
						field:  fmt.Sprintf("ResponseHeadersToAdd[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SentinelNewValidationError{
						field:  fmt.Sprintf("ResponseHeadersToAdd[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SentinelNewValidationError{
					field:  fmt.Sprintf("ResponseHeadersToAdd[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SentinelNewMultiError(errors)
	}

	return nil
}

// SentinelNewMultiError is an error wrapping multiple validation errors
// returned by SentinelNew.ValidateAll() if the designated constraints aren't met.
type SentinelNewMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SentinelNewMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SentinelNewMultiError) AllErrors() []error { return m }

// SentinelNewValidationError is the validation error returned by
// SentinelNew.Validate if the designated constraints aren't met.
type SentinelNewValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SentinelNewValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SentinelNewValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SentinelNewValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SentinelNewValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SentinelNewValidationError) ErrorName() string { return "SentinelNewValidationError" }

// Error satisfies the builtin error interface
func (e SentinelNewValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSentinelNew.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SentinelNewValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SentinelNewValidationError{}
