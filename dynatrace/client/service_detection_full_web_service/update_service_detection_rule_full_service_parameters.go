// Code generated by go-swagger; DO NOT EDIT.

package service_detection_full_web_service

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

// NewUpdateServiceDetectionRuleFullServiceParams creates a new UpdateServiceDetectionRuleFullServiceParams object
// with the default values initialized.
func NewUpdateServiceDetectionRuleFullServiceParams() *UpdateServiceDetectionRuleFullServiceParams {
	var ()
	return &UpdateServiceDetectionRuleFullServiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateServiceDetectionRuleFullServiceParamsWithTimeout creates a new UpdateServiceDetectionRuleFullServiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateServiceDetectionRuleFullServiceParamsWithTimeout(timeout time.Duration) *UpdateServiceDetectionRuleFullServiceParams {
	var ()
	return &UpdateServiceDetectionRuleFullServiceParams{

		timeout: timeout,
	}
}

// NewUpdateServiceDetectionRuleFullServiceParamsWithContext creates a new UpdateServiceDetectionRuleFullServiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateServiceDetectionRuleFullServiceParamsWithContext(ctx context.Context) *UpdateServiceDetectionRuleFullServiceParams {
	var ()
	return &UpdateServiceDetectionRuleFullServiceParams{

		Context: ctx,
	}
}

// NewUpdateServiceDetectionRuleFullServiceParamsWithHTTPClient creates a new UpdateServiceDetectionRuleFullServiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateServiceDetectionRuleFullServiceParamsWithHTTPClient(client *http.Client) *UpdateServiceDetectionRuleFullServiceParams {
	var ()
	return &UpdateServiceDetectionRuleFullServiceParams{
		HTTPClient: client,
	}
}

/*UpdateServiceDetectionRuleFullServiceParams contains all the parameters to send to the API endpoint
for the update service detection rule full service operation typically these are written to a http.Request
*/
type UpdateServiceDetectionRuleFullServiceParams struct {

	/*Body
	  The JSON body of the request containing updated parameters of the service detection rule.

	*/
	Body *dynatrace.FullWebServiceRule
	/*ID
	  The ID of the rule to be updated.

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update service detection rule full service params
func (o *UpdateServiceDetectionRuleFullServiceParams) WithTimeout(timeout time.Duration) *UpdateServiceDetectionRuleFullServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update service detection rule full service params
func (o *UpdateServiceDetectionRuleFullServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update service detection rule full service params
func (o *UpdateServiceDetectionRuleFullServiceParams) WithContext(ctx context.Context) *UpdateServiceDetectionRuleFullServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update service detection rule full service params
func (o *UpdateServiceDetectionRuleFullServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update service detection rule full service params
func (o *UpdateServiceDetectionRuleFullServiceParams) WithHTTPClient(client *http.Client) *UpdateServiceDetectionRuleFullServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update service detection rule full service params
func (o *UpdateServiceDetectionRuleFullServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update service detection rule full service params
func (o *UpdateServiceDetectionRuleFullServiceParams) WithBody(body *dynatrace.FullWebServiceRule) *UpdateServiceDetectionRuleFullServiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update service detection rule full service params
func (o *UpdateServiceDetectionRuleFullServiceParams) SetBody(body *dynatrace.FullWebServiceRule) {
	o.Body = body
}

// WithID adds the id to the update service detection rule full service params
func (o *UpdateServiceDetectionRuleFullServiceParams) WithID(id strfmt.UUID) *UpdateServiceDetectionRuleFullServiceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update service detection rule full service params
func (o *UpdateServiceDetectionRuleFullServiceParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateServiceDetectionRuleFullServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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