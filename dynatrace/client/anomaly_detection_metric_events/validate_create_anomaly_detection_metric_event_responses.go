// Code generated by go-swagger; DO NOT EDIT.

package anomaly_detection_metric_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// ValidateCreateAnomalyDetectionMetricEventReader is a Reader for the ValidateCreateAnomalyDetectionMetricEvent structure.
type ValidateCreateAnomalyDetectionMetricEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateCreateAnomalyDetectionMetricEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewValidateCreateAnomalyDetectionMetricEventNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateCreateAnomalyDetectionMetricEventBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateCreateAnomalyDetectionMetricEventNoContent creates a ValidateCreateAnomalyDetectionMetricEventNoContent with default headers values
func NewValidateCreateAnomalyDetectionMetricEventNoContent() *ValidateCreateAnomalyDetectionMetricEventNoContent {
	return &ValidateCreateAnomalyDetectionMetricEventNoContent{}
}

/*ValidateCreateAnomalyDetectionMetricEventNoContent handles this case with default header values.

Validated. The submitted configuration is valid. Response doesn't have a body.
*/
type ValidateCreateAnomalyDetectionMetricEventNoContent struct {
}

func (o *ValidateCreateAnomalyDetectionMetricEventNoContent) Error() string {
	return fmt.Sprintf("[POST /anomalyDetection/metricEvents/validator][%d] validateCreateAnomalyDetectionMetricEventNoContent ", 204)
}

func (o *ValidateCreateAnomalyDetectionMetricEventNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateCreateAnomalyDetectionMetricEventBadRequest creates a ValidateCreateAnomalyDetectionMetricEventBadRequest with default headers values
func NewValidateCreateAnomalyDetectionMetricEventBadRequest() *ValidateCreateAnomalyDetectionMetricEventBadRequest {
	return &ValidateCreateAnomalyDetectionMetricEventBadRequest{}
}

/*ValidateCreateAnomalyDetectionMetricEventBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type ValidateCreateAnomalyDetectionMetricEventBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *ValidateCreateAnomalyDetectionMetricEventBadRequest) Error() string {
	return fmt.Sprintf("[POST /anomalyDetection/metricEvents/validator][%d] validateCreateAnomalyDetectionMetricEventBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateCreateAnomalyDetectionMetricEventBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *ValidateCreateAnomalyDetectionMetricEventBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}