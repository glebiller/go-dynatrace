// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/Kissy/go-dynatrace/dynatrace"
)

// UploadPluginReader is a Reader for the UploadPlugin structure.
type UploadPluginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadPluginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUploadPluginCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUploadPluginBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUploadPluginCreated creates a UploadPluginCreated with default headers values
func NewUploadPluginCreated() *UploadPluginCreated {
	return &UploadPluginCreated{}
}

/*UploadPluginCreated handles this case with default header values.

Success. Plugin has been uploaded. Response contains the ID of the plugin.
*/
type UploadPluginCreated struct {
	Payload *dynatrace.EntityShortRepresentation
}

func (o *UploadPluginCreated) Error() string {
	return fmt.Sprintf("[POST /plugins][%d] uploadPluginCreated  %+v", 201, o.Payload)
}

func (o *UploadPluginCreated) GetPayload() *dynatrace.EntityShortRepresentation {
	return o.Payload
}

func (o *UploadPluginCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.EntityShortRepresentation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadPluginBadRequest creates a UploadPluginBadRequest with default headers values
func NewUploadPluginBadRequest() *UploadPluginBadRequest {
	return &UploadPluginBadRequest{}
}

/*UploadPluginBadRequest handles this case with default header values.

Failed. The input is invalid.
*/
type UploadPluginBadRequest struct {
	Payload *dynatrace.ErrorEnvelope
}

func (o *UploadPluginBadRequest) Error() string {
	return fmt.Sprintf("[POST /plugins][%d] uploadPluginBadRequest  %+v", 400, o.Payload)
}

func (o *UploadPluginBadRequest) GetPayload() *dynatrace.ErrorEnvelope {
	return o.Payload
}

func (o *UploadPluginBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dynatrace.ErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
