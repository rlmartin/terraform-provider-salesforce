// Code generated by go-swagger; DO NOT EDIT.

package platform_event_channel_member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetPlatformEventChannelMemberParams creates a new GetPlatformEventChannelMemberParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPlatformEventChannelMemberParams() *GetPlatformEventChannelMemberParams {
	return &GetPlatformEventChannelMemberParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlatformEventChannelMemberParamsWithTimeout creates a new GetPlatformEventChannelMemberParams object
// with the ability to set a timeout on a request.
func NewGetPlatformEventChannelMemberParamsWithTimeout(timeout time.Duration) *GetPlatformEventChannelMemberParams {
	return &GetPlatformEventChannelMemberParams{
		timeout: timeout,
	}
}

// NewGetPlatformEventChannelMemberParamsWithContext creates a new GetPlatformEventChannelMemberParams object
// with the ability to set a context for a request.
func NewGetPlatformEventChannelMemberParamsWithContext(ctx context.Context) *GetPlatformEventChannelMemberParams {
	return &GetPlatformEventChannelMemberParams{
		Context: ctx,
	}
}

// NewGetPlatformEventChannelMemberParamsWithHTTPClient creates a new GetPlatformEventChannelMemberParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPlatformEventChannelMemberParamsWithHTTPClient(client *http.Client) *GetPlatformEventChannelMemberParams {
	return &GetPlatformEventChannelMemberParams{
		HTTPClient: client,
	}
}

/*
GetPlatformEventChannelMemberParams contains all the parameters to send to the API endpoint

	for the get platform event channel member operation.

	Typically these are written to a http.Request.
*/
type GetPlatformEventChannelMemberParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get platform event channel member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlatformEventChannelMemberParams) WithDefaults() *GetPlatformEventChannelMemberParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get platform event channel member params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPlatformEventChannelMemberParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get platform event channel member params
func (o *GetPlatformEventChannelMemberParams) WithTimeout(timeout time.Duration) *GetPlatformEventChannelMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get platform event channel member params
func (o *GetPlatformEventChannelMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get platform event channel member params
func (o *GetPlatformEventChannelMemberParams) WithContext(ctx context.Context) *GetPlatformEventChannelMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get platform event channel member params
func (o *GetPlatformEventChannelMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get platform event channel member params
func (o *GetPlatformEventChannelMemberParams) WithHTTPClient(client *http.Client) *GetPlatformEventChannelMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get platform event channel member params
func (o *GetPlatformEventChannelMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get platform event channel member params
func (o *GetPlatformEventChannelMemberParams) WithID(id string) *GetPlatformEventChannelMemberParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get platform event channel member params
func (o *GetPlatformEventChannelMemberParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlatformEventChannelMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Id
	if err := r.SetPathParam("Id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
