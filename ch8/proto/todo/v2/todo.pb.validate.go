// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: todo/v2/todo.proto

package v2

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

// Validate checks the field values on Task with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Task) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Task with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TaskMultiError, or nil if none found.
func (m *Task) ValidateAll() error {
	return m.validate(true)
}

func (m *Task) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Description

	// no validation rules for Done

	if all {
		switch v := interface{}(m.GetDueDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TaskValidationError{
					field:  "DueDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TaskValidationError{
					field:  "DueDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDueDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskValidationError{
				field:  "DueDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TaskMultiError(errors)
	}

	return nil
}

// TaskMultiError is an error wrapping multiple validation errors returned by
// Task.ValidateAll() if the designated constraints aren't met.
type TaskMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TaskMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TaskMultiError) AllErrors() []error { return m }

// TaskValidationError is the validation error returned by Task.Validate if the
// designated constraints aren't met.
type TaskValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskValidationError) ErrorName() string { return "TaskValidationError" }

// Error satisfies the builtin error interface
func (e TaskValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTask.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskValidationError{}

// Validate checks the field values on AddTaskRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddTaskRequestMultiError,
// or nil if none found.
func (m *AddTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		err := AddTaskRequestValidationError{
			field:  "Description",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if t := m.GetDueDate(); t != nil {
		ts, err := t.AsTime(), t.CheckValid()
		if err != nil {
			err = AddTaskRequestValidationError{
				field:  "DueDate",
				reason: "value is not a valid timestamp",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			now := time.Now()

			if ts.Sub(now) <= 0 {
				err := AddTaskRequestValidationError{
					field:  "DueDate",
					reason: "value must be greater than now",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if len(errors) > 0 {
		return AddTaskRequestMultiError(errors)
	}

	return nil
}

// AddTaskRequestMultiError is an error wrapping multiple validation errors
// returned by AddTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type AddTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddTaskRequestMultiError) AllErrors() []error { return m }

// AddTaskRequestValidationError is the validation error returned by
// AddTaskRequest.Validate if the designated constraints aren't met.
type AddTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddTaskRequestValidationError) ErrorName() string { return "AddTaskRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddTaskRequestValidationError{}

// Validate checks the field values on AddTaskResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddTaskResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddTaskResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddTaskResponseMultiError, or nil if none found.
func (m *AddTaskResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AddTaskResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return AddTaskResponseMultiError(errors)
	}

	return nil
}

// AddTaskResponseMultiError is an error wrapping multiple validation errors
// returned by AddTaskResponse.ValidateAll() if the designated constraints
// aren't met.
type AddTaskResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddTaskResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddTaskResponseMultiError) AllErrors() []error { return m }

// AddTaskResponseValidationError is the validation error returned by
// AddTaskResponse.Validate if the designated constraints aren't met.
type AddTaskResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddTaskResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddTaskResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddTaskResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddTaskResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddTaskResponseValidationError) ErrorName() string { return "AddTaskResponseValidationError" }

// Error satisfies the builtin error interface
func (e AddTaskResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddTaskResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddTaskResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddTaskResponseValidationError{}

// Validate checks the field values on ListTasksRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListTasksRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTasksRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTasksRequestMultiError, or nil if none found.
func (m *ListTasksRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTasksRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListTasksRequestValidationError{
					field:  "Mask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListTasksRequestValidationError{
					field:  "Mask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListTasksRequestValidationError{
				field:  "Mask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListTasksRequestMultiError(errors)
	}

	return nil
}

// ListTasksRequestMultiError is an error wrapping multiple validation errors
// returned by ListTasksRequest.ValidateAll() if the designated constraints
// aren't met.
type ListTasksRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTasksRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTasksRequestMultiError) AllErrors() []error { return m }

// ListTasksRequestValidationError is the validation error returned by
// ListTasksRequest.Validate if the designated constraints aren't met.
type ListTasksRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTasksRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTasksRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTasksRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTasksRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTasksRequestValidationError) ErrorName() string { return "ListTasksRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListTasksRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTasksRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTasksRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTasksRequestValidationError{}

// Validate checks the field values on ListTasksResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListTasksResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTasksResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTasksResponseMultiError, or nil if none found.
func (m *ListTasksResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTasksResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListTasksResponseValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListTasksResponseValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListTasksResponseValidationError{
				field:  "Task",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Overdue

	if len(errors) > 0 {
		return ListTasksResponseMultiError(errors)
	}

	return nil
}

// ListTasksResponseMultiError is an error wrapping multiple validation errors
// returned by ListTasksResponse.ValidateAll() if the designated constraints
// aren't met.
type ListTasksResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTasksResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTasksResponseMultiError) AllErrors() []error { return m }

// ListTasksResponseValidationError is the validation error returned by
// ListTasksResponse.Validate if the designated constraints aren't met.
type ListTasksResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTasksResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTasksResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTasksResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTasksResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTasksResponseValidationError) ErrorName() string {
	return "ListTasksResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListTasksResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTasksResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTasksResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTasksResponseValidationError{}

// Validate checks the field values on UpdateTasksRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateTasksRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateTasksRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateTasksRequestMultiError, or nil if none found.
func (m *UpdateTasksRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTasksRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateTasksRequestValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateTasksRequestValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateTasksRequestValidationError{
				field:  "Task",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateTasksRequestMultiError(errors)
	}

	return nil
}

// UpdateTasksRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateTasksRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateTasksRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTasksRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTasksRequestMultiError) AllErrors() []error { return m }

// UpdateTasksRequestValidationError is the validation error returned by
// UpdateTasksRequest.Validate if the designated constraints aren't met.
type UpdateTasksRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTasksRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTasksRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTasksRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTasksRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTasksRequestValidationError) ErrorName() string {
	return "UpdateTasksRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTasksRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTasksRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTasksRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTasksRequestValidationError{}

// Validate checks the field values on UpdateTasksResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateTasksResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateTasksResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateTasksResponseMultiError, or nil if none found.
func (m *UpdateTasksResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateTasksResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateTasksResponseMultiError(errors)
	}

	return nil
}

// UpdateTasksResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateTasksResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateTasksResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateTasksResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateTasksResponseMultiError) AllErrors() []error { return m }

// UpdateTasksResponseValidationError is the validation error returned by
// UpdateTasksResponse.Validate if the designated constraints aren't met.
type UpdateTasksResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateTasksResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateTasksResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateTasksResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateTasksResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateTasksResponseValidationError) ErrorName() string {
	return "UpdateTasksResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateTasksResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateTasksResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateTasksResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateTasksResponseValidationError{}

// Validate checks the field values on DeleteTasksRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteTasksRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTasksRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTasksRequestMultiError, or nil if none found.
func (m *DeleteTasksRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTasksRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteTasksRequestMultiError(errors)
	}

	return nil
}

// DeleteTasksRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteTasksRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteTasksRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTasksRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTasksRequestMultiError) AllErrors() []error { return m }

// DeleteTasksRequestValidationError is the validation error returned by
// DeleteTasksRequest.Validate if the designated constraints aren't met.
type DeleteTasksRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTasksRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTasksRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTasksRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTasksRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTasksRequestValidationError) ErrorName() string {
	return "DeleteTasksRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTasksRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTasksRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTasksRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTasksRequestValidationError{}

// Validate checks the field values on DeleteTasksResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteTasksResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTasksResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTasksResponseMultiError, or nil if none found.
func (m *DeleteTasksResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTasksResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteTasksResponseMultiError(errors)
	}

	return nil
}

// DeleteTasksResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteTasksResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteTasksResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTasksResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTasksResponseMultiError) AllErrors() []error { return m }

// DeleteTasksResponseValidationError is the validation error returned by
// DeleteTasksResponse.Validate if the designated constraints aren't met.
type DeleteTasksResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTasksResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTasksResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTasksResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTasksResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTasksResponseValidationError) ErrorName() string {
	return "DeleteTasksResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTasksResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTasksResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTasksResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTasksResponseValidationError{}
