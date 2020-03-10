// Code generated by go-swagger; DO NOT EDIT.

package mobile_deobfuscation_and_symbolication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// DeleteSymfileReader is a Reader for the DeleteSymfile structure.
type DeleteSymfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSymfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSymfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSymfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSymfileNoContent creates a DeleteSymfileNoContent with default headers values
func NewDeleteSymfileNoContent() *DeleteSymfileNoContent {
	return &DeleteSymfileNoContent{}
}

/*DeleteSymfileNoContent handles this case with default header values.

Success. The file has been deleted successfully. Response doesn't have a body.
*/
type DeleteSymfileNoContent struct {
}

func (o *DeleteSymfileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName}][%d] deleteSymfileNoContent ", 204)
}

func (o *DeleteSymfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSymfileBadRequest creates a DeleteSymfileBadRequest with default headers values
func NewDeleteSymfileBadRequest() *DeleteSymfileBadRequest {
	return &DeleteSymfileBadRequest{}
}

/*DeleteSymfileBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type DeleteSymfileBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *DeleteSymfileBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName}][%d] deleteSymfileBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSymfileBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *DeleteSymfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
