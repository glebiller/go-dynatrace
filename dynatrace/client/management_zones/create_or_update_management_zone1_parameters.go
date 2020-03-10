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

// NewCreateOrUpdateManagementZone1Params creates a new CreateOrUpdateManagementZone1Params object
// with the default values initialized.
func NewCreateOrUpdateManagementZone1Params() *CreateOrUpdateManagementZone1Params {
	var ()
	return &CreateOrUpdateManagementZone1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrUpdateManagementZone1ParamsWithTimeout creates a new CreateOrUpdateManagementZone1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateOrUpdateManagementZone1ParamsWithTimeout(timeout time.Duration) *CreateOrUpdateManagementZone1Params {
	var ()
	return &CreateOrUpdateManagementZone1Params{

		timeout: timeout,
	}
}

// NewCreateOrUpdateManagementZone1ParamsWithContext creates a new CreateOrUpdateManagementZone1Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateOrUpdateManagementZone1ParamsWithContext(ctx context.Context) *CreateOrUpdateManagementZone1Params {
	var ()
	return &CreateOrUpdateManagementZone1Params{

		Context: ctx,
	}
}

// NewCreateOrUpdateManagementZone1ParamsWithHTTPClient creates a new CreateOrUpdateManagementZone1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateOrUpdateManagementZone1ParamsWithHTTPClient(client *http.Client) *CreateOrUpdateManagementZone1Params {
	var ()
	return &CreateOrUpdateManagementZone1Params{
		HTTPClient: client,
	}
}

/*CreateOrUpdateManagementZone1Params contains all the parameters to send to the API endpoint
for the create or update management zone 1 operation typically these are written to a http.Request
*/
type CreateOrUpdateManagementZone1Params struct {

	/*Body
	  The JSON body of the request, containing updated parameters of the management zone.

	*/
	Body *dynatrace.ManagementZone
	/*ID
	 The ID of the management zone to update.

	If you set the ID in the body as well, it must match this ID.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create or update management zone 1 params
func (o *CreateOrUpdateManagementZone1Params) WithTimeout(timeout time.Duration) *CreateOrUpdateManagementZone1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create or update management zone 1 params
func (o *CreateOrUpdateManagementZone1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create or update management zone 1 params
func (o *CreateOrUpdateManagementZone1Params) WithContext(ctx context.Context) *CreateOrUpdateManagementZone1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create or update management zone 1 params
func (o *CreateOrUpdateManagementZone1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create or update management zone 1 params
func (o *CreateOrUpdateManagementZone1Params) WithHTTPClient(client *http.Client) *CreateOrUpdateManagementZone1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create or update management zone 1 params
func (o *CreateOrUpdateManagementZone1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create or update management zone 1 params
func (o *CreateOrUpdateManagementZone1Params) WithBody(body *dynatrace.ManagementZone) *CreateOrUpdateManagementZone1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create or update management zone 1 params
func (o *CreateOrUpdateManagementZone1Params) SetBody(body *dynatrace.ManagementZone) {
	o.Body = body
}

// WithID adds the id to the create or update management zone 1 params
func (o *CreateOrUpdateManagementZone1Params) WithID(id string) *CreateOrUpdateManagementZone1Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the create or update management zone 1 params
func (o *CreateOrUpdateManagementZone1Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrUpdateManagementZone1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
