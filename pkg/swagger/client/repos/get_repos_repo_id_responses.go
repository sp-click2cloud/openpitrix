// Code generated by go-swagger; DO NOT EDIT.

package repos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/pkg/swagger/models"
)

// GetReposRepoIDReader is a Reader for the GetReposRepoID structure.
type GetReposRepoIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReposRepoIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReposRepoIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetReposRepoIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetReposRepoIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReposRepoIDOK creates a GetReposRepoIDOK with default headers values
func NewGetReposRepoIDOK() *GetReposRepoIDOK {
	return &GetReposRepoIDOK{}
}

/*GetReposRepoIDOK handles this case with default header values.

A repo
*/
type GetReposRepoIDOK struct {
	Payload *models.Repo
}

func (o *GetReposRepoIDOK) Error() string {
	return fmt.Sprintf("[GET /repos/{repoId}][%d] getReposRepoIdOK  %+v", 200, o.Payload)
}

func (o *GetReposRepoIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Repo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReposRepoIDNotFound creates a GetReposRepoIDNotFound with default headers values
func NewGetReposRepoIDNotFound() *GetReposRepoIDNotFound {
	return &GetReposRepoIDNotFound{}
}

/*GetReposRepoIDNotFound handles this case with default header values.

The repo does not exist.
*/
type GetReposRepoIDNotFound struct {
}

func (o *GetReposRepoIDNotFound) Error() string {
	return fmt.Sprintf("[GET /repos/{repoId}][%d] getReposRepoIdNotFound ", 404)
}

func (o *GetReposRepoIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReposRepoIDInternalServerError creates a GetReposRepoIDInternalServerError with default headers values
func NewGetReposRepoIDInternalServerError() *GetReposRepoIDInternalServerError {
	return &GetReposRepoIDInternalServerError{}
}

/*GetReposRepoIDInternalServerError handles this case with default header values.

An unexpected error occured.
*/
type GetReposRepoIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetReposRepoIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /repos/{repoId}][%d] getReposRepoIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReposRepoIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
