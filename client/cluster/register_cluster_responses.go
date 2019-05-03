// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vshn/cdays-webui-poc/models"
)

// RegisterClusterReader is a Reader for the RegisterCluster structure.
type RegisterClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRegisterClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRegisterClusterOK creates a RegisterClusterOK with default headers values
func NewRegisterClusterOK() *RegisterClusterOK {
	return &RegisterClusterOK{}
}

/*RegisterClusterOK handles this case with default header values.

registers a new cluster
*/
type RegisterClusterOK struct {
	Payload *models.Cluster
}

func (o *RegisterClusterOK) Error() string {
	return fmt.Sprintf("[PUT /cluster][%d] registerClusterOK  %+v", 200, o.Payload)
}

func (o *RegisterClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
