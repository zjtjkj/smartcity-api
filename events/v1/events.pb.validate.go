// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: events/v1/events.proto

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

// Validate checks the field values on ReceiveEventRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReceiveEventRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReceiveEventRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReceiveEventRequestMultiError, or nil if none found.
func (m *ReceiveEventRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ReceiveEventRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetId()) != 36 {
		err := ReceiveEventRequestValidationError{
			field:  "Id",
			reason: "value length must be 36 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetMid() <= 0 {
		err := ReceiveEventRequestValidationError{
			field:  "Mid",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetType()) < 1 {
		err := ReceiveEventRequestValidationError{
			field:  "Type",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCreated()) < 1 {
		err := ReceiveEventRequestValidationError{
			field:  "Created",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetImage()) < 1 {
		err := ReceiveEventRequestValidationError{
			field:  "Image",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Finished

	if len(m.GetObjects()) < 1 {
		err := ReceiveEventRequestValidationError{
			field:  "Objects",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetObjects() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ReceiveEventRequestValidationError{
						field:  fmt.Sprintf("Objects[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ReceiveEventRequestValidationError{
						field:  fmt.Sprintf("Objects[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReceiveEventRequestValidationError{
					field:  fmt.Sprintf("Objects[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ReceiveEventRequestMultiError(errors)
	}

	return nil
}

// ReceiveEventRequestMultiError is an error wrapping multiple validation
// errors returned by ReceiveEventRequest.ValidateAll() if the designated
// constraints aren't met.
type ReceiveEventRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReceiveEventRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReceiveEventRequestMultiError) AllErrors() []error { return m }

// ReceiveEventRequestValidationError is the validation error returned by
// ReceiveEventRequest.Validate if the designated constraints aren't met.
type ReceiveEventRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceiveEventRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceiveEventRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceiveEventRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceiveEventRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceiveEventRequestValidationError) ErrorName() string {
	return "ReceiveEventRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ReceiveEventRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceiveEventRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceiveEventRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceiveEventRequestValidationError{}

// Validate checks the field values on ReceiveEventReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ReceiveEventReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReceiveEventReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReceiveEventReplyMultiError, or nil if none found.
func (m *ReceiveEventReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ReceiveEventReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ReceiveEventReplyMultiError(errors)
	}

	return nil
}

// ReceiveEventReplyMultiError is an error wrapping multiple validation errors
// returned by ReceiveEventReply.ValidateAll() if the designated constraints
// aren't met.
type ReceiveEventReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReceiveEventReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReceiveEventReplyMultiError) AllErrors() []error { return m }

// ReceiveEventReplyValidationError is the validation error returned by
// ReceiveEventReply.Validate if the designated constraints aren't met.
type ReceiveEventReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceiveEventReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceiveEventReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceiveEventReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceiveEventReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceiveEventReplyValidationError) ErrorName() string {
	return "ReceiveEventReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ReceiveEventReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceiveEventReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceiveEventReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceiveEventReplyValidationError{}

// Validate checks the field values on ReceiveEventRequest_Point with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReceiveEventRequest_Point) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReceiveEventRequest_Point with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReceiveEventRequest_PointMultiError, or nil if none found.
func (m *ReceiveEventRequest_Point) ValidateAll() error {
	return m.validate(true)
}

func (m *ReceiveEventRequest_Point) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if val := m.GetX(); val < 0 || val > 10000 {
		err := ReceiveEventRequest_PointValidationError{
			field:  "X",
			reason: "value must be inside range [0, 10000]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetY(); val < 0 || val > 10000 {
		err := ReceiveEventRequest_PointValidationError{
			field:  "Y",
			reason: "value must be inside range [0, 10000]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ReceiveEventRequest_PointMultiError(errors)
	}

	return nil
}

// ReceiveEventRequest_PointMultiError is an error wrapping multiple validation
// errors returned by ReceiveEventRequest_Point.ValidateAll() if the
// designated constraints aren't met.
type ReceiveEventRequest_PointMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReceiveEventRequest_PointMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReceiveEventRequest_PointMultiError) AllErrors() []error { return m }

// ReceiveEventRequest_PointValidationError is the validation error returned by
// ReceiveEventRequest_Point.Validate if the designated constraints aren't met.
type ReceiveEventRequest_PointValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceiveEventRequest_PointValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceiveEventRequest_PointValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceiveEventRequest_PointValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceiveEventRequest_PointValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceiveEventRequest_PointValidationError) ErrorName() string {
	return "ReceiveEventRequest_PointValidationError"
}

// Error satisfies the builtin error interface
func (e ReceiveEventRequest_PointValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceiveEventRequest_Point.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceiveEventRequest_PointValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceiveEventRequest_PointValidationError{}

// Validate checks the field values on ReceiveEventRequest_Property with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReceiveEventRequest_Property) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReceiveEventRequest_Property with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReceiveEventRequest_PropertyMultiError, or nil if none found.
func (m *ReceiveEventRequest_Property) ValidateAll() error {
	return m.validate(true)
}

func (m *ReceiveEventRequest_Property) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := ReceiveEventRequest_PropertyValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetValue()) < 1 {
		err := ReceiveEventRequest_PropertyValidationError{
			field:  "Value",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ReceiveEventRequest_PropertyMultiError(errors)
	}

	return nil
}

// ReceiveEventRequest_PropertyMultiError is an error wrapping multiple
// validation errors returned by ReceiveEventRequest_Property.ValidateAll() if
// the designated constraints aren't met.
type ReceiveEventRequest_PropertyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReceiveEventRequest_PropertyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReceiveEventRequest_PropertyMultiError) AllErrors() []error { return m }

// ReceiveEventRequest_PropertyValidationError is the validation error returned
// by ReceiveEventRequest_Property.Validate if the designated constraints
// aren't met.
type ReceiveEventRequest_PropertyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceiveEventRequest_PropertyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceiveEventRequest_PropertyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceiveEventRequest_PropertyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceiveEventRequest_PropertyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceiveEventRequest_PropertyValidationError) ErrorName() string {
	return "ReceiveEventRequest_PropertyValidationError"
}

// Error satisfies the builtin error interface
func (e ReceiveEventRequest_PropertyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceiveEventRequest_Property.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceiveEventRequest_PropertyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceiveEventRequest_PropertyValidationError{}

// Validate checks the field values on ReceiveEventRequest_Object with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReceiveEventRequest_Object) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReceiveEventRequest_Object with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReceiveEventRequest_ObjectMultiError, or nil if none found.
func (m *ReceiveEventRequest_Object) ValidateAll() error {
	return m.validate(true)
}

func (m *ReceiveEventRequest_Object) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := ReceiveEventRequest_ObjectValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Aid

	if len(m.GetPoints()) != 2 {
		err := ReceiveEventRequest_ObjectValidationError{
			field:  "Points",
			reason: "value must contain exactly 2 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetPoints() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ReceiveEventRequest_ObjectValidationError{
						field:  fmt.Sprintf("Points[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ReceiveEventRequest_ObjectValidationError{
						field:  fmt.Sprintf("Points[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReceiveEventRequest_ObjectValidationError{
					field:  fmt.Sprintf("Points[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetProperties() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ReceiveEventRequest_ObjectValidationError{
						field:  fmt.Sprintf("Properties[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ReceiveEventRequest_ObjectValidationError{
						field:  fmt.Sprintf("Properties[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReceiveEventRequest_ObjectValidationError{
					field:  fmt.Sprintf("Properties[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ReceiveEventRequest_ObjectMultiError(errors)
	}

	return nil
}

// ReceiveEventRequest_ObjectMultiError is an error wrapping multiple
// validation errors returned by ReceiveEventRequest_Object.ValidateAll() if
// the designated constraints aren't met.
type ReceiveEventRequest_ObjectMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReceiveEventRequest_ObjectMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReceiveEventRequest_ObjectMultiError) AllErrors() []error { return m }

// ReceiveEventRequest_ObjectValidationError is the validation error returned
// by ReceiveEventRequest_Object.Validate if the designated constraints aren't met.
type ReceiveEventRequest_ObjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceiveEventRequest_ObjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceiveEventRequest_ObjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceiveEventRequest_ObjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceiveEventRequest_ObjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceiveEventRequest_ObjectValidationError) ErrorName() string {
	return "ReceiveEventRequest_ObjectValidationError"
}

// Error satisfies the builtin error interface
func (e ReceiveEventRequest_ObjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceiveEventRequest_Object.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceiveEventRequest_ObjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceiveEventRequest_ObjectValidationError{}
