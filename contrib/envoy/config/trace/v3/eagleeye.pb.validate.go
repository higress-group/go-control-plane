// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: contrib/envoy/config/trace/v3/eagleeye.proto

package tracev3

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

// Validate checks the field values on EagleEyeConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EagleEyeConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EagleEyeConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EagleEyeConfigMultiError,
// or nil if none found.
func (m *EagleEyeConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *EagleEyeConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TurnOnRpcLog

	// no validation rules for TurnOnBusinessLog

	if val := m.GetSamplingInterval(); val < 1 || val > 9000 {
		err := EagleEyeConfigValidationError{
			field:  "SamplingInterval",
			reason: "value must be inside range [1, 9000]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for PassUserData

	if len(errors) > 0 {
		return EagleEyeConfigMultiError(errors)
	}

	return nil
}

// EagleEyeConfigMultiError is an error wrapping multiple validation errors
// returned by EagleEyeConfig.ValidateAll() if the designated constraints
// aren't met.
type EagleEyeConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EagleEyeConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EagleEyeConfigMultiError) AllErrors() []error { return m }

// EagleEyeConfigValidationError is the validation error returned by
// EagleEyeConfig.Validate if the designated constraints aren't met.
type EagleEyeConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EagleEyeConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EagleEyeConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EagleEyeConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EagleEyeConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EagleEyeConfigValidationError) ErrorName() string { return "EagleEyeConfigValidationError" }

// Error satisfies the builtin error interface
func (e EagleEyeConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEagleEyeConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EagleEyeConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EagleEyeConfigValidationError{}
