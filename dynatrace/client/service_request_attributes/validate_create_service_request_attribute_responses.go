// Code generated by go-swagger; DO NOT EDIT.

package service_request_attributes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// ValidateCreateServiceRequestAttributeReader is a Reader for the ValidateCreateServiceRequestAttribute structure.
type ValidateCreateServiceRequestAttributeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateCreateServiceRequestAttributeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewValidateCreateServiceRequestAttributeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateCreateServiceRequestAttributeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateCreateServiceRequestAttributeNoContent creates a ValidateCreateServiceRequestAttributeNoContent with default headers values
func NewValidateCreateServiceRequestAttributeNoContent() *ValidateCreateServiceRequestAttributeNoContent {
	return &ValidateCreateServiceRequestAttributeNoContent{}
}

/*ValidateCreateServiceRequestAttributeNoContent handles this case with default header values.

Validated. The submitted configuration is valid. Response does not have a body.
*/
type ValidateCreateServiceRequestAttributeNoContent struct {
}

func (o *ValidateCreateServiceRequestAttributeNoContent) Error() string {
	return fmt.Sprintf("[POST /service/requestAttributes/validator][%d] validateCreateServiceRequestAttributeNoContent ", 204)
}

func (o *ValidateCreateServiceRequestAttributeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateCreateServiceRequestAttributeBadRequest creates a ValidateCreateServiceRequestAttributeBadRequest with default headers values
func NewValidateCreateServiceRequestAttributeBadRequest() *ValidateCreateServiceRequestAttributeBadRequest {
	return &ValidateCreateServiceRequestAttributeBadRequest{}
}

/*ValidateCreateServiceRequestAttributeBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type ValidateCreateServiceRequestAttributeBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *ValidateCreateServiceRequestAttributeBadRequest) Error() string {
	return fmt.Sprintf("[POST /service/requestAttributes/validator][%d] validateCreateServiceRequestAttributeBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateCreateServiceRequestAttributeBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *ValidateCreateServiceRequestAttributeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
