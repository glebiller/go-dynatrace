// Code generated by go-swagger; DO NOT EDIT.

package r_u_m_metrics

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

// NewUpdateCalculatedMetricsRumParams creates a new UpdateCalculatedMetricsRumParams object
// with the default values initialized.
func NewUpdateCalculatedMetricsRumParams() *UpdateCalculatedMetricsRumParams {
	var ()
	return &UpdateCalculatedMetricsRumParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCalculatedMetricsRumParamsWithTimeout creates a new UpdateCalculatedMetricsRumParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCalculatedMetricsRumParamsWithTimeout(timeout time.Duration) *UpdateCalculatedMetricsRumParams {
	var ()
	return &UpdateCalculatedMetricsRumParams{

		timeout: timeout,
	}
}

// NewUpdateCalculatedMetricsRumParamsWithContext creates a new UpdateCalculatedMetricsRumParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCalculatedMetricsRumParamsWithContext(ctx context.Context) *UpdateCalculatedMetricsRumParams {
	var ()
	return &UpdateCalculatedMetricsRumParams{

		Context: ctx,
	}
}

// NewUpdateCalculatedMetricsRumParamsWithHTTPClient creates a new UpdateCalculatedMetricsRumParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCalculatedMetricsRumParamsWithHTTPClient(client *http.Client) *UpdateCalculatedMetricsRumParams {
	var ()
	return &UpdateCalculatedMetricsRumParams{
		HTTPClient: client,
	}
}

/*UpdateCalculatedMetricsRumParams contains all the parameters to send to the API endpoint
for the update calculated metrics rum operation typically these are written to a http.Request
*/
type UpdateCalculatedMetricsRumParams struct {

	/*Body
	  JSON body of the request containing updated definition of the calculated RUM metric.

	*/
	Body *dynatrace.RumMetricUpdate
	/*MetricKey
	  The ID of the calculated RUM metric to update.

	*/
	MetricKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update calculated metrics rum params
func (o *UpdateCalculatedMetricsRumParams) WithTimeout(timeout time.Duration) *UpdateCalculatedMetricsRumParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update calculated metrics rum params
func (o *UpdateCalculatedMetricsRumParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update calculated metrics rum params
func (o *UpdateCalculatedMetricsRumParams) WithContext(ctx context.Context) *UpdateCalculatedMetricsRumParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update calculated metrics rum params
func (o *UpdateCalculatedMetricsRumParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update calculated metrics rum params
func (o *UpdateCalculatedMetricsRumParams) WithHTTPClient(client *http.Client) *UpdateCalculatedMetricsRumParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update calculated metrics rum params
func (o *UpdateCalculatedMetricsRumParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update calculated metrics rum params
func (o *UpdateCalculatedMetricsRumParams) WithBody(body *dynatrace.RumMetricUpdate) *UpdateCalculatedMetricsRumParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update calculated metrics rum params
func (o *UpdateCalculatedMetricsRumParams) SetBody(body *dynatrace.RumMetricUpdate) {
	o.Body = body
}

// WithMetricKey adds the metricKey to the update calculated metrics rum params
func (o *UpdateCalculatedMetricsRumParams) WithMetricKey(metricKey string) *UpdateCalculatedMetricsRumParams {
	o.SetMetricKey(metricKey)
	return o
}

// SetMetricKey adds the metricKey to the update calculated metrics rum params
func (o *UpdateCalculatedMetricsRumParams) SetMetricKey(metricKey string) {
	o.MetricKey = metricKey
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCalculatedMetricsRumParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param metricKey
	if err := r.SetPathParam("metricKey", o.MetricKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}