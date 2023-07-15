// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NamedCredentialMetadata named credential metadata
//
// swagger:model NamedCredentialMetadata
type NamedCredentialMetadata struct {

	// allow merge fields in body
	// Example: false
	AllowMergeFieldsInBody bool `json:"allowMergeFieldsInBody,omitempty"`

	// allow merge fields in header
	// Example: false
	AllowMergeFieldsInHeader bool `json:"allowMergeFieldsInHeader,omitempty"`

	// endpoint
	// Example: arn:aws:US-EAST-1:XXXXXXXXXX
	// Required: true
	Endpoint *string `json:"endpoint"`

	// generate authorization header
	// Example: true
	GenerateAuthorizationHeader bool `json:"generateAuthorizationHeader,omitempty"`

	// label
	// Example: AWS US-East-1
	// Required: true
	Label *string `json:"label"`

	// principal type
	// Example: Anonymous
	// Required: true
	// Enum: [Anonymous NamedUser PerUser]
	PrincipalType *string `json:"principalType"`

	// protocol
	// Example: NoAuthentication
	// Required: true
	// Enum: [AwsSv4 Jwt JwtExchange NoAuthentication Oauth Password]
	Protocol *string `json:"protocol"`
}

// Validate validates this named credential metadata
func (m *NamedCredentialMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrincipalType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NamedCredentialMetadata) validateEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("endpoint", "body", m.Endpoint); err != nil {
		return err
	}

	return nil
}

func (m *NamedCredentialMetadata) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

var namedCredentialMetadataTypePrincipalTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Anonymous","NamedUser","PerUser"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		namedCredentialMetadataTypePrincipalTypePropEnum = append(namedCredentialMetadataTypePrincipalTypePropEnum, v)
	}
}

const (

	// NamedCredentialMetadataPrincipalTypeAnonymous captures enum value "Anonymous"
	NamedCredentialMetadataPrincipalTypeAnonymous string = "Anonymous"

	// NamedCredentialMetadataPrincipalTypeNamedUser captures enum value "NamedUser"
	NamedCredentialMetadataPrincipalTypeNamedUser string = "NamedUser"

	// NamedCredentialMetadataPrincipalTypePerUser captures enum value "PerUser"
	NamedCredentialMetadataPrincipalTypePerUser string = "PerUser"
)

// prop value enum
func (m *NamedCredentialMetadata) validatePrincipalTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, namedCredentialMetadataTypePrincipalTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NamedCredentialMetadata) validatePrincipalType(formats strfmt.Registry) error {

	if err := validate.Required("principalType", "body", m.PrincipalType); err != nil {
		return err
	}

	// value enum
	if err := m.validatePrincipalTypeEnum("principalType", "body", *m.PrincipalType); err != nil {
		return err
	}

	return nil
}

var namedCredentialMetadataTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AwsSv4","Jwt","JwtExchange","NoAuthentication","Oauth","Password"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		namedCredentialMetadataTypeProtocolPropEnum = append(namedCredentialMetadataTypeProtocolPropEnum, v)
	}
}

const (

	// NamedCredentialMetadataProtocolAwsSv4 captures enum value "AwsSv4"
	NamedCredentialMetadataProtocolAwsSv4 string = "AwsSv4"

	// NamedCredentialMetadataProtocolJwt captures enum value "Jwt"
	NamedCredentialMetadataProtocolJwt string = "Jwt"

	// NamedCredentialMetadataProtocolJwtExchange captures enum value "JwtExchange"
	NamedCredentialMetadataProtocolJwtExchange string = "JwtExchange"

	// NamedCredentialMetadataProtocolNoAuthentication captures enum value "NoAuthentication"
	NamedCredentialMetadataProtocolNoAuthentication string = "NoAuthentication"

	// NamedCredentialMetadataProtocolOauth captures enum value "Oauth"
	NamedCredentialMetadataProtocolOauth string = "Oauth"

	// NamedCredentialMetadataProtocolPassword captures enum value "Password"
	NamedCredentialMetadataProtocolPassword string = "Password"
)

// prop value enum
func (m *NamedCredentialMetadata) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, namedCredentialMetadataTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NamedCredentialMetadata) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this named credential metadata based on context it is used
func (m *NamedCredentialMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NamedCredentialMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamedCredentialMetadata) UnmarshalBinary(b []byte) error {
	var res NamedCredentialMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
