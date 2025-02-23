// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: contrib/envoy/extensions/filters/common/sentinel/v3/config_new.proto

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

// Validate checks the field values on CommonConfigNew with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CommonConfigNew) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommonConfigNew with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommonConfigNewMultiError, or nil if none found.
func (m *CommonConfigNew) ValidateAll() error {
	return m.validate(true)
}

func (m *CommonConfigNew) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetAppName()) < 1 {
		err := CommonConfigNewValidationError{
			field:  "AppName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for SystemMetricIntervalMs

	// no validation rules for LogDir

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CommonConfigNewValidationError{
						field:  fmt.Sprintf("Rules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CommonConfigNewValidationError{
						field:  fmt.Sprintf("Rules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommonConfigNewValidationError{
					field:  fmt.Sprintf("Rules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CommonConfigNewMultiError(errors)
	}

	return nil
}

// CommonConfigNewMultiError is an error wrapping multiple validation errors
// returned by CommonConfigNew.ValidateAll() if the designated constraints
// aren't met.
type CommonConfigNewMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommonConfigNewMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommonConfigNewMultiError) AllErrors() []error { return m }

// CommonConfigNewValidationError is the validation error returned by
// CommonConfigNew.Validate if the designated constraints aren't met.
type CommonConfigNewValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonConfigNewValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonConfigNewValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonConfigNewValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonConfigNewValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonConfigNewValidationError) ErrorName() string { return "CommonConfigNewValidationError" }

// Error satisfies the builtin error interface
func (e CommonConfigNewValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonConfigNew.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonConfigNewValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonConfigNewValidationError{}

// Validate checks the field values on RateLimitStrategy with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RateLimitStrategy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RateLimitStrategy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RateLimitStrategyMultiError, or nil if none found.
func (m *RateLimitStrategy) ValidateAll() error {
	return m.validate(true)
}

func (m *RateLimitStrategy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for LimitMode

	if m.GetThreshold() < 0 {
		err := RateLimitStrategyValidationError{
			field:  "Threshold",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Burst

	if all {
		switch v := interface{}(m.GetStatDuration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RateLimitStrategyValidationError{
					field:  "StatDuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RateLimitStrategyValidationError{
					field:  "StatDuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RateLimitStrategyValidationError{
				field:  "StatDuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RateLimitStrategyMultiError(errors)
	}

	return nil
}

// RateLimitStrategyMultiError is an error wrapping multiple validation errors
// returned by RateLimitStrategy.ValidateAll() if the designated constraints
// aren't met.
type RateLimitStrategyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RateLimitStrategyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RateLimitStrategyMultiError) AllErrors() []error { return m }

// RateLimitStrategyValidationError is the validation error returned by
// RateLimitStrategy.Validate if the designated constraints aren't met.
type RateLimitStrategyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitStrategyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitStrategyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitStrategyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitStrategyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitStrategyValidationError) ErrorName() string {
	return "RateLimitStrategyValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitStrategyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitStrategy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitStrategyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitStrategyValidationError{}

// Validate checks the field values on ConcurrencyLimitStrategy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ConcurrencyLimitStrategy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConcurrencyLimitStrategy with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConcurrencyLimitStrategyMultiError, or nil if none found.
func (m *ConcurrencyLimitStrategy) ValidateAll() error {
	return m.validate(true)
}

func (m *ConcurrencyLimitStrategy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for LimitMode

	if m.GetMaxConcurrency() < 0 {
		err := ConcurrencyLimitStrategyValidationError{
			field:  "MaxConcurrency",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ConcurrencyLimitStrategyMultiError(errors)
	}

	return nil
}

// ConcurrencyLimitStrategyMultiError is an error wrapping multiple validation
// errors returned by ConcurrencyLimitStrategy.ValidateAll() if the designated
// constraints aren't met.
type ConcurrencyLimitStrategyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConcurrencyLimitStrategyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConcurrencyLimitStrategyMultiError) AllErrors() []error { return m }

// ConcurrencyLimitStrategyValidationError is the validation error returned by
// ConcurrencyLimitStrategy.Validate if the designated constraints aren't met.
type ConcurrencyLimitStrategyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConcurrencyLimitStrategyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConcurrencyLimitStrategyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConcurrencyLimitStrategyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConcurrencyLimitStrategyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConcurrencyLimitStrategyValidationError) ErrorName() string {
	return "ConcurrencyLimitStrategyValidationError"
}

// Error satisfies the builtin error interface
func (e ConcurrencyLimitStrategyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConcurrencyLimitStrategy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConcurrencyLimitStrategyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConcurrencyLimitStrategyValidationError{}

// Validate checks the field values on CircuitBreakerStrategy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CircuitBreakerStrategy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CircuitBreakerStrategy with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CircuitBreakerStrategyMultiError, or nil if none found.
func (m *CircuitBreakerStrategy) ValidateAll() error {
	return m.validate(true)
}

func (m *CircuitBreakerStrategy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Strategy

	if val := m.GetTriggerRatio(); val < 0 || val > 100 {
		err := CircuitBreakerStrategyValidationError{
			field:  "TriggerRatio",
			reason: "value must be inside range [0, 100]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetStatDuration() == nil {
		err := CircuitBreakerStrategyValidationError{
			field:  "StatDuration",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetRecoveryTimeout() == nil {
		err := CircuitBreakerStrategyValidationError{
			field:  "RecoveryTimeout",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetMinRequestAmount() < 0 {
		err := CircuitBreakerStrategyValidationError{
			field:  "MinRequestAmount",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetSlowCondition()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CircuitBreakerStrategyValidationError{
					field:  "SlowCondition",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CircuitBreakerStrategyValidationError{
					field:  "SlowCondition",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSlowCondition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CircuitBreakerStrategyValidationError{
				field:  "SlowCondition",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetErrorCondition()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CircuitBreakerStrategyValidationError{
					field:  "ErrorCondition",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CircuitBreakerStrategyValidationError{
					field:  "ErrorCondition",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetErrorCondition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CircuitBreakerStrategyValidationError{
				field:  "ErrorCondition",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CircuitBreakerStrategyMultiError(errors)
	}

	return nil
}

// CircuitBreakerStrategyMultiError is an error wrapping multiple validation
// errors returned by CircuitBreakerStrategy.ValidateAll() if the designated
// constraints aren't met.
type CircuitBreakerStrategyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CircuitBreakerStrategyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CircuitBreakerStrategyMultiError) AllErrors() []error { return m }

// CircuitBreakerStrategyValidationError is the validation error returned by
// CircuitBreakerStrategy.Validate if the designated constraints aren't met.
type CircuitBreakerStrategyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CircuitBreakerStrategyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CircuitBreakerStrategyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CircuitBreakerStrategyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CircuitBreakerStrategyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CircuitBreakerStrategyValidationError) ErrorName() string {
	return "CircuitBreakerStrategyValidationError"
}

// Error satisfies the builtin error interface
func (e CircuitBreakerStrategyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCircuitBreakerStrategy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CircuitBreakerStrategyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CircuitBreakerStrategyValidationError{}

// Validate checks the field values on FallbackAction with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FallbackAction) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FallbackAction with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FallbackActionMultiError,
// or nil if none found.
func (m *FallbackAction) ValidateAll() error {
	return m.validate(true)
}

func (m *FallbackAction) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for BodyEncoding

	// no validation rules for ResponseStatuscode

	if len(m.GetResponseContentBody()) > 1024 {
		err := FallbackActionValidationError{
			field:  "ResponseContentBody",
			reason: "value length must be at most 1024 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetResponseRedirectUrl()) > 1024 {
		err := FallbackActionValidationError{
			field:  "ResponseRedirectUrl",
			reason: "value length must be at most 1024 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ResponseAdditionalHeaders

	if len(errors) > 0 {
		return FallbackActionMultiError(errors)
	}

	return nil
}

// FallbackActionMultiError is an error wrapping multiple validation errors
// returned by FallbackAction.ValidateAll() if the designated constraints
// aren't met.
type FallbackActionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FallbackActionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FallbackActionMultiError) AllErrors() []error { return m }

// FallbackActionValidationError is the validation error returned by
// FallbackAction.Validate if the designated constraints aren't met.
type FallbackActionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FallbackActionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FallbackActionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FallbackActionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FallbackActionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FallbackActionValidationError) ErrorName() string { return "FallbackActionValidationError" }

// Error satisfies the builtin error interface
func (e FallbackActionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFallbackAction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FallbackActionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FallbackActionValidationError{}

// Validate checks the field values on FaultToleranceRule with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FaultToleranceRule) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FaultToleranceRule with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FaultToleranceRuleMultiError, or nil if none found.
func (m *FaultToleranceRule) ValidateAll() error {
	return m.validate(true)
}

func (m *FaultToleranceRule) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetTargetResourceNames()) > 100 {
		err := FaultToleranceRuleValidationError{
			field:  "TargetResourceNames",
			reason: "value must contain no more than 100 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetRateLimitStrategies()) > 10 {
		err := FaultToleranceRuleValidationError{
			field:  "RateLimitStrategies",
			reason: "value must contain no more than 10 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetRateLimitStrategies() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FaultToleranceRuleValidationError{
						field:  fmt.Sprintf("RateLimitStrategies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FaultToleranceRuleValidationError{
						field:  fmt.Sprintf("RateLimitStrategies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FaultToleranceRuleValidationError{
					field:  fmt.Sprintf("RateLimitStrategies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(m.GetConcurrencyLimitStrategies()) > 10 {
		err := FaultToleranceRuleValidationError{
			field:  "ConcurrencyLimitStrategies",
			reason: "value must contain no more than 10 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetConcurrencyLimitStrategies() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FaultToleranceRuleValidationError{
						field:  fmt.Sprintf("ConcurrencyLimitStrategies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FaultToleranceRuleValidationError{
						field:  fmt.Sprintf("ConcurrencyLimitStrategies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FaultToleranceRuleValidationError{
					field:  fmt.Sprintf("ConcurrencyLimitStrategies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(m.GetCircuitBreakerStrategies()) > 10 {
		err := FaultToleranceRuleValidationError{
			field:  "CircuitBreakerStrategies",
			reason: "value must contain no more than 10 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetCircuitBreakerStrategies() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FaultToleranceRuleValidationError{
						field:  fmt.Sprintf("CircuitBreakerStrategies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FaultToleranceRuleValidationError{
						field:  fmt.Sprintf("CircuitBreakerStrategies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FaultToleranceRuleValidationError{
					field:  fmt.Sprintf("CircuitBreakerStrategies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.GetAction() == nil {
		err := FaultToleranceRuleValidationError{
			field:  "Action",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetAction()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FaultToleranceRuleValidationError{
					field:  "Action",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FaultToleranceRuleValidationError{
					field:  "Action",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAction()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FaultToleranceRuleValidationError{
				field:  "Action",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return FaultToleranceRuleMultiError(errors)
	}

	return nil
}

// FaultToleranceRuleMultiError is an error wrapping multiple validation errors
// returned by FaultToleranceRule.ValidateAll() if the designated constraints
// aren't met.
type FaultToleranceRuleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FaultToleranceRuleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FaultToleranceRuleMultiError) AllErrors() []error { return m }

// FaultToleranceRuleValidationError is the validation error returned by
// FaultToleranceRule.Validate if the designated constraints aren't met.
type FaultToleranceRuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FaultToleranceRuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FaultToleranceRuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FaultToleranceRuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FaultToleranceRuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FaultToleranceRuleValidationError) ErrorName() string {
	return "FaultToleranceRuleValidationError"
}

// Error satisfies the builtin error interface
func (e FaultToleranceRuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFaultToleranceRule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FaultToleranceRuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FaultToleranceRuleValidationError{}

// Validate checks the field values on
// CircuitBreakerStrategy_CircuitBreakerSlowCondition with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CircuitBreakerStrategy_CircuitBreakerSlowCondition) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// CircuitBreakerStrategy_CircuitBreakerSlowCondition with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in
// CircuitBreakerStrategy_CircuitBreakerSlowConditionMultiError, or nil if
// none found.
func (m *CircuitBreakerStrategy_CircuitBreakerSlowCondition) ValidateAll() error {
	return m.validate(true)
}

func (m *CircuitBreakerStrategy_CircuitBreakerSlowCondition) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetMaxAllowedRt() == nil {
		err := CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError{
			field:  "MaxAllowedRt",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CircuitBreakerStrategy_CircuitBreakerSlowConditionMultiError(errors)
	}

	return nil
}

// CircuitBreakerStrategy_CircuitBreakerSlowConditionMultiError is an error
// wrapping multiple validation errors returned by
// CircuitBreakerStrategy_CircuitBreakerSlowCondition.ValidateAll() if the
// designated constraints aren't met.
type CircuitBreakerStrategy_CircuitBreakerSlowConditionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CircuitBreakerStrategy_CircuitBreakerSlowConditionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CircuitBreakerStrategy_CircuitBreakerSlowConditionMultiError) AllErrors() []error { return m }

// CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError is the
// validation error returned by
// CircuitBreakerStrategy_CircuitBreakerSlowCondition.Validate if the
// designated constraints aren't met.
type CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError) ErrorName() string {
	return "CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError"
}

// Error satisfies the builtin error interface
func (e CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCircuitBreakerStrategy_CircuitBreakerSlowCondition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CircuitBreakerStrategy_CircuitBreakerSlowConditionValidationError{}

// Validate checks the field values on
// CircuitBreakerStrategy_CircuitBreakerErrorCondition with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CircuitBreakerStrategy_CircuitBreakerErrorCondition) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// CircuitBreakerStrategy_CircuitBreakerErrorCondition with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in
// CircuitBreakerStrategy_CircuitBreakerErrorConditionMultiError, or nil if
// none found.
func (m *CircuitBreakerStrategy_CircuitBreakerErrorCondition) ValidateAll() error {
	return m.validate(true)
}

func (m *CircuitBreakerStrategy_CircuitBreakerErrorCondition) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CircuitBreakerStrategy_CircuitBreakerErrorConditionMultiError(errors)
	}

	return nil
}

// CircuitBreakerStrategy_CircuitBreakerErrorConditionMultiError is an error
// wrapping multiple validation errors returned by
// CircuitBreakerStrategy_CircuitBreakerErrorCondition.ValidateAll() if the
// designated constraints aren't met.
type CircuitBreakerStrategy_CircuitBreakerErrorConditionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CircuitBreakerStrategy_CircuitBreakerErrorConditionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CircuitBreakerStrategy_CircuitBreakerErrorConditionMultiError) AllErrors() []error { return m }

// CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError is the
// validation error returned by
// CircuitBreakerStrategy_CircuitBreakerErrorCondition.Validate if the
// designated constraints aren't met.
type CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError) ErrorName() string {
	return "CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError"
}

// Error satisfies the builtin error interface
func (e CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCircuitBreakerStrategy_CircuitBreakerErrorCondition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CircuitBreakerStrategy_CircuitBreakerErrorConditionValidationError{}

// Validate checks the field values on CircuitBreakerStrategy_RecoveryStrategy
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *CircuitBreakerStrategy_RecoveryStrategy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// CircuitBreakerStrategy_RecoveryStrategy with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// CircuitBreakerStrategy_RecoveryStrategyMultiError, or nil if none found.
func (m *CircuitBreakerStrategy_RecoveryStrategy) ValidateAll() error {
	return m.validate(true)
}

func (m *CircuitBreakerStrategy_RecoveryStrategy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProbeCount

	if len(errors) > 0 {
		return CircuitBreakerStrategy_RecoveryStrategyMultiError(errors)
	}

	return nil
}

// CircuitBreakerStrategy_RecoveryStrategyMultiError is an error wrapping
// multiple validation errors returned by
// CircuitBreakerStrategy_RecoveryStrategy.ValidateAll() if the designated
// constraints aren't met.
type CircuitBreakerStrategy_RecoveryStrategyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CircuitBreakerStrategy_RecoveryStrategyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CircuitBreakerStrategy_RecoveryStrategyMultiError) AllErrors() []error { return m }

// CircuitBreakerStrategy_RecoveryStrategyValidationError is the validation
// error returned by CircuitBreakerStrategy_RecoveryStrategy.Validate if the
// designated constraints aren't met.
type CircuitBreakerStrategy_RecoveryStrategyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CircuitBreakerStrategy_RecoveryStrategyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CircuitBreakerStrategy_RecoveryStrategyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CircuitBreakerStrategy_RecoveryStrategyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CircuitBreakerStrategy_RecoveryStrategyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CircuitBreakerStrategy_RecoveryStrategyValidationError) ErrorName() string {
	return "CircuitBreakerStrategy_RecoveryStrategyValidationError"
}

// Error satisfies the builtin error interface
func (e CircuitBreakerStrategy_RecoveryStrategyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCircuitBreakerStrategy_RecoveryStrategy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CircuitBreakerStrategy_RecoveryStrategyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CircuitBreakerStrategy_RecoveryStrategyValidationError{}
