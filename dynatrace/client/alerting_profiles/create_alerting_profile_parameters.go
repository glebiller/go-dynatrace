// Code generated by go-swagger; DO NOT EDIT.

package alerting_profiles

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

// NewCreateAlertingProfileParams creates a new CreateAlertingProfileParams object
// with the default values initialized.
func NewCreateAlertingProfileParams() *CreateAlertingProfileParams {
	var ()
	return &CreateAlertingProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAlertingProfileParamsWithTimeout creates a new CreateAlertingProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAlertingProfileParamsWithTimeout(timeout time.Duration) *CreateAlertingProfileParams {
	var ()
	return &CreateAlertingProfileParams{

		timeout: timeout,
	}
}

// NewCreateAlertingProfileParamsWithContext creates a new CreateAlertingProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAlertingProfileParamsWithContext(ctx context.Context) *CreateAlertingProfileParams {
	var ()
	return &CreateAlertingProfileParams{

		Context: ctx,
	}
}

// NewCreateAlertingProfileParamsWithHTTPClient creates a new CreateAlertingProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAlertingProfileParamsWithHTTPClient(client *http.Client) *CreateAlertingProfileParams {
	var ()
	return &CreateAlertingProfileParams{
		HTTPClient: client,
	}
}

/*CreateAlertingProfileParams contains all the parameters to send to the API endpoint
for the create alerting profile operation typically these are written to a http.Request
*/
type CreateAlertingProfileParams struct {

	/*Body
	  The JSON body of the request. Contains parameters of the new alerting profile.

	*/
	Body *dynatrace.AlertingProfile

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create alerting profile params
func (o *CreateAlertingProfileParams) WithTimeout(timeout time.Duration) *CreateAlertingProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create alerting profile params
func (o *CreateAlertingProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create alerting profile params
func (o *CreateAlertingProfileParams) WithContext(ctx context.Context) *CreateAlertingProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create alerting profile params
func (o *CreateAlertingProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create alerting profile params
func (o *CreateAlertingProfileParams) WithHTTPClient(client *http.Client) *CreateAlertingProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create alerting profile params
func (o *CreateAlertingProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create alerting profile params
func (o *CreateAlertingProfileParams) WithBody(body *dynatrace.AlertingProfile) *CreateAlertingProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create alerting profile params
func (o *CreateAlertingProfileParams) SetBody(body *dynatrace.AlertingProfile) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAlertingProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
