// Code generated by go-swagger; DO NOT EDIT.

package event_relay_config

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

	"vestahealthcare/models"
)

// NewCreateEventRelayConfigParams creates a new CreateEventRelayConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateEventRelayConfigParams() *CreateEventRelayConfigParams {
	return &CreateEventRelayConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEventRelayConfigParamsWithTimeout creates a new CreateEventRelayConfigParams object
// with the ability to set a timeout on a request.
func NewCreateEventRelayConfigParamsWithTimeout(timeout time.Duration) *CreateEventRelayConfigParams {
	return &CreateEventRelayConfigParams{
		timeout: timeout,
	}
}

// NewCreateEventRelayConfigParamsWithContext creates a new CreateEventRelayConfigParams object
// with the ability to set a context for a request.
func NewCreateEventRelayConfigParamsWithContext(ctx context.Context) *CreateEventRelayConfigParams {
	return &CreateEventRelayConfigParams{
		Context: ctx,
	}
}

// NewCreateEventRelayConfigParamsWithHTTPClient creates a new CreateEventRelayConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateEventRelayConfigParamsWithHTTPClient(client *http.Client) *CreateEventRelayConfigParams {
	return &CreateEventRelayConfigParams{
		HTTPClient: client,
	}
}

/*
CreateEventRelayConfigParams contains all the parameters to send to the API endpoint

	for the create event relay config operation.

	Typically these are written to a http.Request.
*/
type CreateEventRelayConfigParams struct {

	// Body.
	Body *models.EventRelayConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create event relay config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEventRelayConfigParams) WithDefaults() *CreateEventRelayConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create event relay config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateEventRelayConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create event relay config params
func (o *CreateEventRelayConfigParams) WithTimeout(timeout time.Duration) *CreateEventRelayConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create event relay config params
func (o *CreateEventRelayConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create event relay config params
func (o *CreateEventRelayConfigParams) WithContext(ctx context.Context) *CreateEventRelayConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create event relay config params
func (o *CreateEventRelayConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create event relay config params
func (o *CreateEventRelayConfigParams) WithHTTPClient(client *http.Client) *CreateEventRelayConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create event relay config params
func (o *CreateEventRelayConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create event relay config params
func (o *CreateEventRelayConfigParams) WithBody(body *models.EventRelayConfig) *CreateEventRelayConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create event relay config params
func (o *CreateEventRelayConfigParams) SetBody(body *models.EventRelayConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEventRelayConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}