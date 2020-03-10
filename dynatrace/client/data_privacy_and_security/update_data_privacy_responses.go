// Code generated by go-swagger; DO NOT EDIT.

package data_privacy_and_security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// UpdateDataPrivacyReader is a Reader for the UpdateDataPrivacy structure.
type UpdateDataPrivacyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDataPrivacyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateDataPrivacyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDataPrivacyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateDataPrivacyNoContent creates a UpdateDataPrivacyNoContent with default headers values
func NewUpdateDataPrivacyNoContent() *UpdateDataPrivacyNoContent {
	return &UpdateDataPrivacyNoContent{}
}

/*UpdateDataPrivacyNoContent handles this case with default header values.

Success. The configuration is updated. Response does not have a body.
*/
type UpdateDataPrivacyNoContent struct {
}

func (o *UpdateDataPrivacyNoContent) Error() string {
	return fmt.Sprintf("[PUT /dataPrivacy][%d] updateDataPrivacyNoContent ", 204)
}

func (o *UpdateDataPrivacyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDataPrivacyBadRequest creates a UpdateDataPrivacyBadRequest with default headers values
func NewUpdateDataPrivacyBadRequest() *UpdateDataPrivacyBadRequest {
	return &UpdateDataPrivacyBadRequest{}
}

/*UpdateDataPrivacyBadRequest handles this case with default header values.

Failed. The input is invalid. See the response body for details.
*/
type UpdateDataPrivacyBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *UpdateDataPrivacyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dataPrivacy][%d] updateDataPrivacyBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDataPrivacyBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *UpdateDataPrivacyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}