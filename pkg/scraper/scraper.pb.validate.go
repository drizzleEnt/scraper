// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: scraper.proto

package scraper_v1

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

// Validate checks the field values on ScrapData with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ScrapData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ScrapData with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ScrapDataMultiError, or nil
// if none found.
func (m *ScrapData) ValidateAll() error {
	return m.validate(true)
}

func (m *ScrapData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Text

	if len(errors) > 0 {
		return ScrapDataMultiError(errors)
	}

	return nil
}

// ScrapDataMultiError is an error wrapping multiple validation errors returned
// by ScrapData.ValidateAll() if the designated constraints aren't met.
type ScrapDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ScrapDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ScrapDataMultiError) AllErrors() []error { return m }

// ScrapDataValidationError is the validation error returned by
// ScrapData.Validate if the designated constraints aren't met.
type ScrapDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ScrapDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ScrapDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ScrapDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ScrapDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ScrapDataValidationError) ErrorName() string { return "ScrapDataValidationError" }

// Error satisfies the builtin error interface
func (e ScrapDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sScrapData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ScrapDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ScrapDataValidationError{}

// Validate checks the field values on ScrapUrl with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ScrapUrl) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ScrapUrl with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ScrapUrlMultiError, or nil
// if none found.
func (m *ScrapUrl) ValidateAll() error {
	return m.validate(true)
}

func (m *ScrapUrl) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Url

	if len(errors) > 0 {
		return ScrapUrlMultiError(errors)
	}

	return nil
}

// ScrapUrlMultiError is an error wrapping multiple validation errors returned
// by ScrapUrl.ValidateAll() if the designated constraints aren't met.
type ScrapUrlMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ScrapUrlMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ScrapUrlMultiError) AllErrors() []error { return m }

// ScrapUrlValidationError is the validation error returned by
// ScrapUrl.Validate if the designated constraints aren't met.
type ScrapUrlValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ScrapUrlValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ScrapUrlValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ScrapUrlValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ScrapUrlValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ScrapUrlValidationError) ErrorName() string { return "ScrapUrlValidationError" }

// Error satisfies the builtin error interface
func (e ScrapUrlValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sScrapUrl.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ScrapUrlValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ScrapUrlValidationError{}
