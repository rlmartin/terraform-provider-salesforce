// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PlatformEventChannelMemberUpdateRequest platform event channel member update request
// Example: isResource
//
// swagger:model PlatformEventChannelMemberUpdateRequest
type PlatformEventChannelMemberUpdateRequest struct {

	// full name
	// Example: My_Channel__chn
	FullName string `json:"FullName,omitempty"`

	// metadata
	Metadata *PlatformEventChannelMemberMetadata `json:"Metadata,omitempty"`
}

// Validate validates this platform event channel member update request
func (m *PlatformEventChannelMemberUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlatformEventChannelMemberUpdateRequest) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Metadata")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this platform event channel member update request based on the context it is used
func (m *PlatformEventChannelMemberUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlatformEventChannelMemberUpdateRequest) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {

		if swag.IsZero(m.Metadata) { // not required
			return nil
		}

		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlatformEventChannelMemberUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlatformEventChannelMemberUpdateRequest) UnmarshalBinary(b []byte) error {
	var res PlatformEventChannelMemberUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}