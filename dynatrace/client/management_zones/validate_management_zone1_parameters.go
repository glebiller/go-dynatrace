// Code generated by go-swagger; DO NOT EDIT.

package management_zones

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

// NewValidateManagementZone1Params creates a new ValidateManagementZone1Params object
// with the default values initialized.
func NewValidateManagementZone1Params() *ValidateManagementZone1Params {
	var ()
	return &ValidateManagementZone1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateManagementZone1ParamsWithTimeout creates a new ValidateManagementZone1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateManagementZone1ParamsWithTimeout(timeout time.Duration) *ValidateManagementZone1Params {
	var ()
	return &ValidateManagementZone1Params{

		timeout: timeout,
	}
}

// NewValidateManagementZone1ParamsWithContext creates a new ValidateManagementZone1Params object
// with the default values initialized, and the ability to set a context for a request
func NewValidateManagementZone1ParamsWithContext(ctx context.Context) *ValidateManagementZone1Params {
	var ()
	return &ValidateManagementZone1Params{

		Context: ctx,
	}
}

// NewValidateManagementZone1ParamsWithHTTPClient creates a new ValidateManagementZone1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateManagementZone1ParamsWithHTTPClient(client *http.Client) *ValidateManagementZone1Params {
	var ()
	return &ValidateManagementZone1Params{
		HTTPClient: client,
	}
}

/*ValidateManagementZone1Params contains all the parameters to send to the API endpoint
for the validate management zone 1 operation typically these are written to a http.Request
*/
type ValidateManagementZone1Params struct {

	/*Body
	  The JSON body of the request, containing the management zone parameters to validate.

	*/
	Body *dynatrace.ManagementZone
	/*ID
	  The ID of the management zone to validate.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate management zone 1 params
func (o *ValidateManagementZone1Params) WithTimeout(timeout time.Duration) *ValidateManagementZone1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate management zone 1 params
func (o *ValidateManagementZone1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate management zone 1 params
func (o *ValidateManagementZone1Params) WithContext(ctx context.Context) *ValidateManagementZone1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate management zone 1 params
func (o *ValidateManagementZone1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate management zone 1 params
func (o *ValidateManagementZone1Params) WithHTTPClient(client *http.Client) *ValidateManagementZone1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate management zone 1 params
func (o *ValidateManagementZone1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate management zone 1 params
func (o *ValidateManagementZone1Params) WithBody(body *dynatrace.ManagementZone) *ValidateManagementZone1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate management zone 1 params
func (o *ValidateManagementZone1Params) SetBody(body *dynatrace.ManagementZone) {
	o.Body = body
}

// WithID adds the id to the validate management zone 1 params
func (o *ValidateManagementZone1Params) WithID(id string) *ValidateManagementZone1Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the validate management zone 1 params
func (o *ValidateManagementZone1Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateManagementZone1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
