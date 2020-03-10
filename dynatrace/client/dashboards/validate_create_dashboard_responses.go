// Code generated by go-swagger; DO NOT EDIT.

package dashboards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// ValidateCreateDashboardReader is a Reader for the ValidateCreateDashboard structure.
type ValidateCreateDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateCreateDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewValidateCreateDashboardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewValidateCreateDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateCreateDashboardNoContent creates a ValidateCreateDashboardNoContent with default headers values
func NewValidateCreateDashboardNoContent() *ValidateCreateDashboardNoContent {
	return &ValidateCreateDashboardNoContent{}
}

/*ValidateCreateDashboardNoContent handles this case with default header values.

Validated. The submitted dashboard is valid. The response doesn't have a body.
*/
type ValidateCreateDashboardNoContent struct {
}

func (o *ValidateCreateDashboardNoContent) Error() string {
	return fmt.Sprintf("[POST /dashboards/validator][%d] validateCreateDashboardNoContent ", 204)
}

func (o *ValidateCreateDashboardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateCreateDashboardBadRequest creates a ValidateCreateDashboardBadRequest with default headers values
func NewValidateCreateDashboardBadRequest() *ValidateCreateDashboardBadRequest {
	return &ValidateCreateDashboardBadRequest{}
}

/*ValidateCreateDashboardBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type ValidateCreateDashboardBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *ValidateCreateDashboardBadRequest) Error() string {
	return fmt.Sprintf("[POST /dashboards/validator][%d] validateCreateDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *ValidateCreateDashboardBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *ValidateCreateDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
