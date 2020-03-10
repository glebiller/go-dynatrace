// Code generated by go-swagger; DO NOT EDIT.

package alerting_profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// ValidateCreateAlertingProfileReader is a Reader for the ValidateCreateAlertingProfile structure.
type ValidateCreateAlertingProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateCreateAlertingProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewValidateCreateAlertingProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateCreateAlertingProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateCreateAlertingProfileNoContent creates a ValidateCreateAlertingProfileNoContent with default headers values
func NewValidateCreateAlertingProfileNoContent() *ValidateCreateAlertingProfileNoContent {
	return &ValidateCreateAlertingProfileNoContent{}
}

/*ValidateCreateAlertingProfileNoContent handles this case with default header values.

Validated. The submitted alerting profile is valid. Response doesn't have a body
*/
type ValidateCreateAlertingProfileNoContent struct {
}

func (o *ValidateCreateAlertingProfileNoContent) Error() string {
	return fmt.Sprintf("[POST /alertingProfiles/validator][%d] validateCreateAlertingProfileNoContent ", 204)
}

func (o *ValidateCreateAlertingProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateCreateAlertingProfileBadRequest creates a ValidateCreateAlertingProfileBadRequest with default headers values
func NewValidateCreateAlertingProfileBadRequest() *ValidateCreateAlertingProfileBadRequest {
	return &ValidateCreateAlertingProfileBadRequest{}
}

/*ValidateCreateAlertingProfileBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type ValidateCreateAlertingProfileBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *ValidateCreateAlertingProfileBadRequest) Error() string {
	return fmt.Sprintf("[POST /alertingProfiles/validator][%d] validateCreateAlertingProfileBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateCreateAlertingProfileBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *ValidateCreateAlertingProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
