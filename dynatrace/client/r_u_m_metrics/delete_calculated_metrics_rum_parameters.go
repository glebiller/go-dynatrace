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
)

// NewDeleteCalculatedMetricsRumParams creates a new DeleteCalculatedMetricsRumParams object
// with the default values initialized.
func NewDeleteCalculatedMetricsRumParams() *DeleteCalculatedMetricsRumParams {
	var ()
	return &DeleteCalculatedMetricsRumParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCalculatedMetricsRumParamsWithTimeout creates a new DeleteCalculatedMetricsRumParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCalculatedMetricsRumParamsWithTimeout(timeout time.Duration) *DeleteCalculatedMetricsRumParams {
	var ()
	return &DeleteCalculatedMetricsRumParams{

		timeout: timeout,
	}
}

// NewDeleteCalculatedMetricsRumParamsWithContext creates a new DeleteCalculatedMetricsRumParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCalculatedMetricsRumParamsWithContext(ctx context.Context) *DeleteCalculatedMetricsRumParams {
	var ()
	return &DeleteCalculatedMetricsRumParams{

		Context: ctx,
	}
}

// NewDeleteCalculatedMetricsRumParamsWithHTTPClient creates a new DeleteCalculatedMetricsRumParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCalculatedMetricsRumParamsWithHTTPClient(client *http.Client) *DeleteCalculatedMetricsRumParams {
	var ()
	return &DeleteCalculatedMetricsRumParams{
		HTTPClient: client,
	}
}

/*DeleteCalculatedMetricsRumParams contains all the parameters to send to the API endpoint
for the delete calculated metrics rum operation typically these are written to a http.Request
*/
type DeleteCalculatedMetricsRumParams struct {

	/*MetricKey
	  The ID of the calculated RUM metric to delete.

	*/
	MetricKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete calculated metrics rum params
func (o *DeleteCalculatedMetricsRumParams) WithTimeout(timeout time.Duration) *DeleteCalculatedMetricsRumParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete calculated metrics rum params
func (o *DeleteCalculatedMetricsRumParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete calculated metrics rum params
func (o *DeleteCalculatedMetricsRumParams) WithContext(ctx context.Context) *DeleteCalculatedMetricsRumParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete calculated metrics rum params
func (o *DeleteCalculatedMetricsRumParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete calculated metrics rum params
func (o *DeleteCalculatedMetricsRumParams) WithHTTPClient(client *http.Client) *DeleteCalculatedMetricsRumParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete calculated metrics rum params
func (o *DeleteCalculatedMetricsRumParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMetricKey adds the metricKey to the delete calculated metrics rum params
func (o *DeleteCalculatedMetricsRumParams) WithMetricKey(metricKey string) *DeleteCalculatedMetricsRumParams {
	o.SetMetricKey(metricKey)
	return o
}

// SetMetricKey adds the metricKey to the delete calculated metrics rum params
func (o *DeleteCalculatedMetricsRumParams) SetMetricKey(metricKey string) {
	o.MetricKey = metricKey
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCalculatedMetricsRumParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param metricKey
	if err := r.SetPathParam("metricKey", o.MetricKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
