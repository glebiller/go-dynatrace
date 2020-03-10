// Code generated by go-swagger; DO NOT EDIT.

package maintenance_windows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// CreateMaintenanceWindowReader is a Reader for the CreateMaintenanceWindow structure.
type CreateMaintenanceWindowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMaintenanceWindowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateMaintenanceWindowCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateMaintenanceWindowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateMaintenanceWindowCreated creates a CreateMaintenanceWindowCreated with default headers values
func NewCreateMaintenanceWindowCreated() *CreateMaintenanceWindowCreated {
	return &CreateMaintenanceWindowCreated{}
}

/*CreateMaintenanceWindowCreated handles this case with default header values.

Success. The new maintenance window has been created. The response body contains its ID.
*/
type CreateMaintenanceWindowCreated struct {
	Payload *dynatrace.EntityShortRepresentation
}

func (o *CreateMaintenanceWindowCreated) Error() string {
	return fmt.Sprintf("[POST /maintenanceWindows][%d] createMaintenanceWindowCreated  %+v", 201, o.Payload)
}

func (o *CreateMaintenanceWindowCreated) GetPayload() *dynatrace.EntityShortRepresentation {
	return o.Payload
}

func (o *CreateMaintenanceWindowCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.EntityShortRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMaintenanceWindowBadRequest creates a CreateMaintenanceWindowBadRequest with default headers values
func NewCreateMaintenanceWindowBadRequest() *CreateMaintenanceWindowBadRequest {
	return &CreateMaintenanceWindowBadRequest{}
}

/*CreateMaintenanceWindowBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type CreateMaintenanceWindowBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *CreateMaintenanceWindowBadRequest) Error() string {
	return fmt.Sprintf("[POST /maintenanceWindows][%d] createMaintenanceWindowBadRequest  %+v", 400, o.Payload)
}

func (o *CreateMaintenanceWindowBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *CreateMaintenanceWindowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
