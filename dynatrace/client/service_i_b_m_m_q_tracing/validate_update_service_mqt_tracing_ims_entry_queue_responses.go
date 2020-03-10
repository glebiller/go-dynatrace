// Code generated by go-swagger; DO NOT EDIT.

package service_i_b_m_m_q_tracing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// ValidateUpdateServiceMqtTracingImsEntryQueueReader is a Reader for the ValidateUpdateServiceMqtTracingImsEntryQueue structure.
type ValidateUpdateServiceMqtTracingImsEntryQueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateUpdateServiceMqtTracingImsEntryQueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewValidateUpdateServiceMqtTracingImsEntryQueueNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateUpdateServiceMqtTracingImsEntryQueueBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateUpdateServiceMqtTracingImsEntryQueueNoContent creates a ValidateUpdateServiceMqtTracingImsEntryQueueNoContent with default headers values
func NewValidateUpdateServiceMqtTracingImsEntryQueueNoContent() *ValidateUpdateServiceMqtTracingImsEntryQueueNoContent {
	return &ValidateUpdateServiceMqtTracingImsEntryQueueNoContent{}
}

/*ValidateUpdateServiceMqtTracingImsEntryQueueNoContent handles this case with default header values.

Validated. The submitted configuration is valid. Response does not have a body.
*/
type ValidateUpdateServiceMqtTracingImsEntryQueueNoContent struct {
}

func (o *ValidateUpdateServiceMqtTracingImsEntryQueueNoContent) Error() string {
	return fmt.Sprintf("[POST /service/ibmMQTracing/imsEntryQueue/{id}/validator][%d] validateUpdateServiceMqtTracingImsEntryQueueNoContent ", 204)
}

func (o *ValidateUpdateServiceMqtTracingImsEntryQueueNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateUpdateServiceMqtTracingImsEntryQueueBadRequest creates a ValidateUpdateServiceMqtTracingImsEntryQueueBadRequest with default headers values
func NewValidateUpdateServiceMqtTracingImsEntryQueueBadRequest() *ValidateUpdateServiceMqtTracingImsEntryQueueBadRequest {
	return &ValidateUpdateServiceMqtTracingImsEntryQueueBadRequest{}
}

/*ValidateUpdateServiceMqtTracingImsEntryQueueBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type ValidateUpdateServiceMqtTracingImsEntryQueueBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *ValidateUpdateServiceMqtTracingImsEntryQueueBadRequest) Error() string {
	return fmt.Sprintf("[POST /service/ibmMQTracing/imsEntryQueue/{id}/validator][%d] validateUpdateServiceMqtTracingImsEntryQueueBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateUpdateServiceMqtTracingImsEntryQueueBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *ValidateUpdateServiceMqtTracingImsEntryQueueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
