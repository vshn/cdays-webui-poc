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

// GetManagedNamespacesReader is a Reader for the GetManagedNamespaces structure.
type GetManagedNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetManagedNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetManagedNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetManagedNamespacesOK creates a GetManagedNamespacesOK with default headers values
func NewGetManagedNamespacesOK() *GetManagedNamespacesOK {
	return &GetManagedNamespacesOK{}
}

/*GetManagedNamespacesOK handles this case with default header values.

returns a list of namespaces
*/
type GetManagedNamespacesOK struct {
	Payload []*models.Namespace
}

func (o *GetManagedNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /namespaces][%d] getManagedNamespacesOK  %+v", 200, o.Payload)
}

func (o *GetManagedNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}