// Code generated by go-swagger; DO NOT EDIT.

package maintenance_windows

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

// NewCreateMaintenanceWindowParams creates a new CreateMaintenanceWindowParams object
// with the default values initialized.
func NewCreateMaintenanceWindowParams() *CreateMaintenanceWindowParams {
	var ()
	return &CreateMaintenanceWindowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMaintenanceWindowParamsWithTimeout creates a new CreateMaintenanceWindowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMaintenanceWindowParamsWithTimeout(timeout time.Duration) *CreateMaintenanceWindowParams {
	var ()
	return &CreateMaintenanceWindowParams{

		timeout: timeout,
	}
}

// NewCreateMaintenanceWindowParamsWithContext creates a new CreateMaintenanceWindowParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMaintenanceWindowParamsWithContext(ctx context.Context) *CreateMaintenanceWindowParams {
	var ()
	return &CreateMaintenanceWindowParams{

		Context: ctx,
	}
}

// NewCreateMaintenanceWindowParamsWithHTTPClient creates a new CreateMaintenanceWindowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMaintenanceWindowParamsWithHTTPClient(client *http.Client) *CreateMaintenanceWindowParams {
	var ()
	return &CreateMaintenanceWindowParams{
		HTTPClient: client,
	}
}

/*CreateMaintenanceWindowParams contains all the parameters to send to the API endpoint
for the create maintenance window operation typically these are written to a http.Request
*/
type CreateMaintenanceWindowParams struct {

	/*Body
	  The JSON body of the request. Contains parameters of the new maintenance window.

	*/
	Body *dynatrace.MaintenanceWindow

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create maintenance window params
func (o *CreateMaintenanceWindowParams) WithTimeout(timeout time.Duration) *CreateMaintenanceWindowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create maintenance window params
func (o *CreateMaintenanceWindowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create maintenance window params
func (o *CreateMaintenanceWindowParams) WithContext(ctx context.Context) *CreateMaintenanceWindowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create maintenance window params
func (o *CreateMaintenanceWindowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create maintenance window params
func (o *CreateMaintenanceWindowParams) WithHTTPClient(client *http.Client) *CreateMaintenanceWindowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create maintenance window params
func (o *CreateMaintenanceWindowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create maintenance window params
func (o *CreateMaintenanceWindowParams) WithBody(body *dynatrace.MaintenanceWindow) *CreateMaintenanceWindowParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create maintenance window params
func (o *CreateMaintenanceWindowParams) SetBody(body *dynatrace.MaintenanceWindow) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMaintenanceWindowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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