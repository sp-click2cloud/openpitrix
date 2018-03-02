// Code generated by go-swagger; DO NOT EDIT.

package runtime_env_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DeleteRuntimeEnvReader is a Reader for the DeleteRuntimeEnv structure.
type DeleteRuntimeEnvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRuntimeEnvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteRuntimeEnvOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRuntimeEnvOK creates a DeleteRuntimeEnvOK with default headers values
func NewDeleteRuntimeEnvOK() *DeleteRuntimeEnvOK {
	return &DeleteRuntimeEnvOK{}
}

/*DeleteRuntimeEnvOK handles this case with default header values.

DeleteRuntimeEnvOK delete runtime env o k
*/
type DeleteRuntimeEnvOK struct {
	Payload *models.OpenpitrixDeleteRuntimeEnvResponse
}

func (o *DeleteRuntimeEnvOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/runtime_envs][%d] deleteRuntimeEnvOK  %+v", 200, o.Payload)
}

func (o *DeleteRuntimeEnvOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDeleteRuntimeEnvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}