// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EventRelayConfigCreateResponse event relay config create response
// Example: isResource
//
// swagger:model EventRelayConfigCreateResponse
type EventRelayConfigCreateResponse struct {

	// id
	// Example: 1234
	ID string `json:"id,omitempty"`
}

// Validate validates this event relay config create response
func (m *EventRelayConfigCreateResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this event relay config create response based on context it is used
func (m *EventRelayConfigCreateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventRelayConfigCreateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventRelayConfigCreateResponse) UnmarshalBinary(b []byte) error {
	var res EventRelayConfigCreateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
