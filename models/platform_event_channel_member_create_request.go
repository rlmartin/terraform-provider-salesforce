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

// PlatformEventChannelMemberCreateRequest platform event channel member create request
// Example: isResource
//
// swagger:model PlatformEventChannelMemberCreateRequest
type PlatformEventChannelMemberCreateRequest struct {

	// full name
	// Example: My_Channel__chn_MyEvent
	// Required: true
	FullName *string `json:"FullName"`

	// metadata
	// Required: true
	Metadata *PlatformEventChannelMemberMetadata `json:"Metadata"`
}

// Validate validates this platform event channel member create request
func (m *PlatformEventChannelMemberCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFullName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlatformEventChannelMemberCreateRequest) validateFullName(formats strfmt.Registry) error {

	if err := validate.Required("FullName", "body", m.FullName); err != nil {
		return err
	}

	return nil
}

func (m *PlatformEventChannelMemberCreateRequest) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("Metadata", "body", m.Metadata); err != nil {
		return err
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

// ContextValidate validate this platform event channel member create request based on the context it is used
func (m *PlatformEventChannelMemberCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlatformEventChannelMemberCreateRequest) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {

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
func (m *PlatformEventChannelMemberCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlatformEventChannelMemberCreateRequest) UnmarshalBinary(b []byte) error {
	var res PlatformEventChannelMemberCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
