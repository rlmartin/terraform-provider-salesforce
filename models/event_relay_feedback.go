// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EventRelayFeedback event relay feedback
// Example: isResource
//
// swagger:model EventRelayFeedback
type EventRelayFeedback struct {

	// event relay number
	// Example: 00000001
	// Required: true
	EventRelayNumber *string `json:"EventRelayNumber"`

	// Id
	// Example: 1234
	ID string `json:"Id,omitempty"`

	// is deleted
	// Example: false
	// Required: true
	IsDeleted *bool `json:"IsDeleted"`

	// remote resource
	// Example: aws.partner/salesforce.com/000000000000/111111111111
	RemoteResource string `json:"RemoteResource,omitempty"`

	// status
	// Example: STOPPED
	// Required: true
	Status *string `json:"Status"`
}

// Validate validates this event relay feedback
func (m *EventRelayFeedback) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventRelayNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventRelayFeedback) validateEventRelayNumber(formats strfmt.Registry) error {

	if err := validate.Required("EventRelayNumber", "body", m.EventRelayNumber); err != nil {
		return err
	}

	return nil
}

func (m *EventRelayFeedback) validateIsDeleted(formats strfmt.Registry) error {

	if err := validate.Required("IsDeleted", "body", m.IsDeleted); err != nil {
		return err
	}

	return nil
}

func (m *EventRelayFeedback) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("Status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this event relay feedback based on context it is used
func (m *EventRelayFeedback) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventRelayFeedback) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventRelayFeedback) UnmarshalBinary(b []byte) error {
	var res EventRelayFeedback
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
