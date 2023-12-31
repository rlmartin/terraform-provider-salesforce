// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EventRelayFeedbackLookup event relay feedback lookup
// Example: isResource
//
// swagger:model EventRelayFeedbackLookup
type EventRelayFeedbackLookup struct {

	// done
	// Example: true
	Done bool `json:"done,omitempty"`

	// Necessary here only for Terraform reasons
	// Required: true
	Q *string `json:"q"`

	// records
	Records []*Record `json:"records"`

	// total size
	// Example: 1
	TotalSize int32 `json:"totalSize,omitempty"`
}

// Validate validates this event relay feedback lookup
func (m *EventRelayFeedbackLookup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQ(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventRelayFeedbackLookup) validateQ(formats strfmt.Registry) error {

	if err := validate.Required("q", "body", m.Q); err != nil {
		return err
	}

	return nil
}

func (m *EventRelayFeedbackLookup) validateRecords(formats strfmt.Registry) error {
	if swag.IsZero(m.Records) { // not required
		return nil
	}

	for i := 0; i < len(m.Records); i++ {
		if swag.IsZero(m.Records[i]) { // not required
			continue
		}

		if m.Records[i] != nil {
			if err := m.Records[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this event relay feedback lookup based on the context it is used
func (m *EventRelayFeedbackLookup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventRelayFeedbackLookup) contextValidateRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Records); i++ {

		if m.Records[i] != nil {

			if swag.IsZero(m.Records[i]) { // not required
				return nil
			}

			if err := m.Records[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("records" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventRelayFeedbackLookup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventRelayFeedbackLookup) UnmarshalBinary(b []byte) error {
	var res EventRelayFeedbackLookup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
