// Code generated by go-swagger; DO NOT EDIT.

package instance

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

// NewStartParams creates a new StartParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartParams() *StartParams {
	return &StartParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartParamsWithTimeout creates a new StartParams object
// with the ability to set a timeout on a request.
func NewStartParamsWithTimeout(timeout time.Duration) *StartParams {
	return &StartParams{
		timeout: timeout,
	}
}

// NewStartParamsWithContext creates a new StartParams object
// with the ability to set a context for a request.
func NewStartParamsWithContext(ctx context.Context) *StartParams {
	return &StartParams{
		Context: ctx,
	}
}

// NewStartParamsWithHTTPClient creates a new StartParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartParamsWithHTTPClient(client *http.Client) *StartParams {
	return &StartParams{
		HTTPClient: client,
	}
}

/* StartParams contains all the parameters to send to the API endpoint
   for the start operation.

   Typically these are written to a http.Request.
*/
type StartParams struct {

	/* Name.

	   Instance name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartParams) WithDefaults() *StartParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the start params
func (o *StartParams) WithTimeout(timeout time.Duration) *StartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start params
func (o *StartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start params
func (o *StartParams) WithContext(ctx context.Context) *StartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start params
func (o *StartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start params
func (o *StartParams) WithHTTPClient(client *http.Client) *StartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start params
func (o *StartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the start params
func (o *StartParams) WithName(name string) *StartParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the start params
func (o *StartParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *StartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
