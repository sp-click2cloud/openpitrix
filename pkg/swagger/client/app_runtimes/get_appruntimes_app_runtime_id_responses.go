// Code generated by go-swagger; DO NOT EDIT.

package app_runtimes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/pkg/swagger/models"
)

// GetAppruntimesAppRuntimeIDReader is a Reader for the GetAppruntimesAppRuntimeID structure.
type GetAppruntimesAppRuntimeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppruntimesAppRuntimeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAppruntimesAppRuntimeIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetAppruntimesAppRuntimeIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAppruntimesAppRuntimeIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAppruntimesAppRuntimeIDOK creates a GetAppruntimesAppRuntimeIDOK with default headers values
func NewGetAppruntimesAppRuntimeIDOK() *GetAppruntimesAppRuntimeIDOK {
	return &GetAppruntimesAppRuntimeIDOK{}
}

/*GetAppruntimesAppRuntimeIDOK handles this case with default header values.

An app runtime
*/
type GetAppruntimesAppRuntimeIDOK struct {
	Payload *models.AppRuntime
}

func (o *GetAppruntimesAppRuntimeIDOK) Error() string {
	return fmt.Sprintf("[GET /appruntimes/{appRuntimeId}][%d] getAppruntimesAppRuntimeIdOK  %+v", 200, o.Payload)
}

func (o *GetAppruntimesAppRuntimeIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppRuntime)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppruntimesAppRuntimeIDNotFound creates a GetAppruntimesAppRuntimeIDNotFound with default headers values
func NewGetAppruntimesAppRuntimeIDNotFound() *GetAppruntimesAppRuntimeIDNotFound {
	return &GetAppruntimesAppRuntimeIDNotFound{}
}

/*GetAppruntimesAppRuntimeIDNotFound handles this case with default header values.

The app runtime does not exist.
*/
type GetAppruntimesAppRuntimeIDNotFound struct {
}

func (o *GetAppruntimesAppRuntimeIDNotFound) Error() string {
	return fmt.Sprintf("[GET /appruntimes/{appRuntimeId}][%d] getAppruntimesAppRuntimeIdNotFound ", 404)
}

func (o *GetAppruntimesAppRuntimeIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAppruntimesAppRuntimeIDInternalServerError creates a GetAppruntimesAppRuntimeIDInternalServerError with default headers values
func NewGetAppruntimesAppRuntimeIDInternalServerError() *GetAppruntimesAppRuntimeIDInternalServerError {
	return &GetAppruntimesAppRuntimeIDInternalServerError{}
}

/*GetAppruntimesAppRuntimeIDInternalServerError handles this case with default header values.

An unexpected error occured.
*/
type GetAppruntimesAppRuntimeIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetAppruntimesAppRuntimeIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /appruntimes/{appRuntimeId}][%d] getAppruntimesAppRuntimeIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAppruntimesAppRuntimeIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
