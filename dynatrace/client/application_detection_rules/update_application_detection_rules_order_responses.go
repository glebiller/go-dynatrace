// Code generated by go-swagger; DO NOT EDIT.

package application_detection_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// UpdateApplicationDetectionRulesOrderReader is a Reader for the UpdateApplicationDetectionRulesOrder structure.
type UpdateApplicationDetectionRulesOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateApplicationDetectionRulesOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateApplicationDetectionRulesOrderNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateApplicationDetectionRulesOrderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateApplicationDetectionRulesOrderNoContent creates a UpdateApplicationDetectionRulesOrderNoContent with default headers values
func NewUpdateApplicationDetectionRulesOrderNoContent() *UpdateApplicationDetectionRulesOrderNoContent {
	return &UpdateApplicationDetectionRulesOrderNoContent{}
}

/*UpdateApplicationDetectionRulesOrderNoContent handles this case with default header values.

Success. Application detection rules have been reordered. Response doesn't have a body.
*/
type UpdateApplicationDetectionRulesOrderNoContent struct {
}

func (o *UpdateApplicationDetectionRulesOrderNoContent) Error() string {
	return fmt.Sprintf("[PUT /applicationDetectionRules/order][%d] updateApplicationDetectionRulesOrderNoContent ", 204)
}

func (o *UpdateApplicationDetectionRulesOrderNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateApplicationDetectionRulesOrderBadRequest creates a UpdateApplicationDetectionRulesOrderBadRequest with default headers values
func NewUpdateApplicationDetectionRulesOrderBadRequest() *UpdateApplicationDetectionRulesOrderBadRequest {
	return &UpdateApplicationDetectionRulesOrderBadRequest{}
}

/*UpdateApplicationDetectionRulesOrderBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type UpdateApplicationDetectionRulesOrderBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *UpdateApplicationDetectionRulesOrderBadRequest) Error() string {
	return fmt.Sprintf("[PUT /applicationDetectionRules/order][%d] updateApplicationDetectionRulesOrderBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateApplicationDetectionRulesOrderBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *UpdateApplicationDetectionRulesOrderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}