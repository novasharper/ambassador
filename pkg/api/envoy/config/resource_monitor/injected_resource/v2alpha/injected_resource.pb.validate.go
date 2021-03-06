// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/resource_monitor/injected_resource/v2alpha/injected_resource.proto

package envoy_config_resource_monitor_injected_resource_v2alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _injected_resource_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on InjectedResourceConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *InjectedResourceConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetFilename()) < 1 {
		return InjectedResourceConfigValidationError{
			field:  "Filename",
			reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// InjectedResourceConfigValidationError is the validation error returned by
// InjectedResourceConfig.Validate if the designated constraints aren't met.
type InjectedResourceConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InjectedResourceConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InjectedResourceConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InjectedResourceConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InjectedResourceConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InjectedResourceConfigValidationError) ErrorName() string {
	return "InjectedResourceConfigValidationError"
}

// Error satisfies the builtin error interface
func (e InjectedResourceConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInjectedResourceConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InjectedResourceConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InjectedResourceConfigValidationError{}
