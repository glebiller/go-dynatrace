// Code generated by go-swagger; DO NOT EDIT.

package log_monitoring_metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// GetCalculatedMetricsLogReader is a Reader for the GetCalculatedMetricsLog structure.
type GetCalculatedMetricsLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCalculatedMetricsLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCalculatedMetricsLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCalculatedMetricsLogOK creates a GetCalculatedMetricsLogOK with default headers values
func NewGetCalculatedMetricsLogOK() *GetCalculatedMetricsLogOK {
	return &GetCalculatedMetricsLogOK{}
}

/*GetCalculatedMetricsLogOK handles this case with default header values.

successful operation
*/
type GetCalculatedMetricsLogOK struct {
	Payload *dynatrace.LogMetricConfig
}

func (o *GetCalculatedMetricsLogOK) Error() string {
	return fmt.Sprintf("[GET /calculatedMetrics/log/{metricKey}][%d] getCalculatedMetricsLogOK  %+v", 200, o.Payload)
}

func (o *GetCalculatedMetricsLogOK) GetPayload() *dynatrace.LogMetricConfig {
	return o.Payload
}

func (o *GetCalculatedMetricsLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.LogMetricConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}