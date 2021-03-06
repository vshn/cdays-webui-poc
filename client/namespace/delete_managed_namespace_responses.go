// Code generated by go-swagger; DO NOT EDIT.

package namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vshn/cdays-webui-poc/models"
)

// DeleteManagedNamespaceReader is a Reader for the DeleteManagedNamespace structure.
type DeleteManagedNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteManagedNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteManagedNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteManagedNamespaceOK creates a DeleteManagedNamespaceOK with default headers values
func NewDeleteManagedNamespaceOK() *DeleteManagedNamespaceOK {
	return &DeleteManagedNamespaceOK{}
}

/*DeleteManagedNamespaceOK handles this case with default header values.

deletes a single namespace
*/
type DeleteManagedNamespaceOK struct {
	Payload *models.Namespace
}

func (o *DeleteManagedNamespaceOK) Error() string {
	return fmt.Sprintf("[DELETE /{clustername}/namespace/{customer}/{name}][%d] deleteManagedNamespaceOK  %+v", 200, o.Payload)
}

func (o *DeleteManagedNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Namespace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
