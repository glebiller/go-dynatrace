// Code generated by go-swagger; DO NOT EDIT.

package service_request_naming

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// NewUpdateServiceRequestNamingParams creates a new UpdateServiceRequestNamingParams object
// with the default values initialized.
func NewUpdateServiceRequestNamingParams() *UpdateServiceRequestNamingParams {
	var ()
	return &UpdateServiceRequestNamingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateServiceRequestNamingParamsWithTimeout creates a new UpdateServiceRequestNamingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateServiceRequestNamingParamsWithTimeout(timeout time.Duration) *UpdateServiceRequestNamingParams {
	var ()
	return &UpdateServiceRequestNamingParams{

		timeout: timeout,
	}
}

// NewUpdateServiceRequestNamingParamsWithContext creates a new UpdateServiceRequestNamingParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateServiceRequestNamingParamsWithContext(ctx context.Context) *UpdateServiceRequestNamingParams {
	var ()
	return &UpdateServiceRequestNamingParams{

		Context: ctx,
	}
}

// NewUpdateServiceRequestNamingParamsWithHTTPClient creates a new UpdateServiceRequestNamingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateServiceRequestNamingParamsWithHTTPClient(client *http.Client) *UpdateServiceRequestNamingParams {
	var ()
	return &UpdateServiceRequestNamingParams{
		HTTPClient: client,
	}
}

/*UpdateServiceRequestNamingParams contains all the parameters to send to the API endpoint
for the update service request naming operation typically these are written to a http.Request
*/
type UpdateServiceRequestNamingParams struct {

	/*Body
	  The JSON body of the request containing updated parameters of the request naming.

	*/
	Body *dynatrace.RequestNaming
	/*ID
	 The ID of the request naming to be updated.

	The ID of the request naming in the body of the request must match this ID.

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update service request naming params
func (o *UpdateServiceRequestNamingParams) WithTimeout(timeout time.Duration) *UpdateServiceRequestNamingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update service request naming params
func (o *UpdateServiceRequestNamingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update service request naming params
func (o *UpdateServiceRequestNamingParams) WithContext(ctx context.Context) *UpdateServiceRequestNamingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update service request naming params
func (o *UpdateServiceRequestNamingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update service request naming params
func (o *UpdateServiceRequestNamingParams) WithHTTPClient(client *http.Client) *UpdateServiceRequestNamingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update service request naming params
func (o *UpdateServiceRequestNamingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update service request naming params
func (o *UpdateServiceRequestNamingParams) WithBody(body *dynatrace.RequestNaming) *UpdateServiceRequestNamingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update service request naming params
func (o *UpdateServiceRequestNamingParams) SetBody(body *dynatrace.RequestNaming) {
	o.Body = body
}

// WithID adds the id to the update service request naming params
func (o *UpdateServiceRequestNamingParams) WithID(id strfmt.UUID) *UpdateServiceRequestNamingParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update service request naming params
func (o *UpdateServiceRequestNamingParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServiceRequestNamingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
