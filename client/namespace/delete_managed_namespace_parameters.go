// Code generated by go-swagger; DO NOT EDIT.

package namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteManagedNamespaceParams creates a new DeleteManagedNamespaceParams object
// with the default values initialized.
func NewDeleteManagedNamespaceParams() *DeleteManagedNamespaceParams {
	var ()
	return &DeleteManagedNamespaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteManagedNamespaceParamsWithTimeout creates a new DeleteManagedNamespaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteManagedNamespaceParamsWithTimeout(timeout time.Duration) *DeleteManagedNamespaceParams {
	var ()
	return &DeleteManagedNamespaceParams{

		timeout: timeout,
	}
}

// NewDeleteManagedNamespaceParamsWithContext creates a new DeleteManagedNamespaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteManagedNamespaceParamsWithContext(ctx context.Context) *DeleteManagedNamespaceParams {
	var ()
	return &DeleteManagedNamespaceParams{

		Context: ctx,
	}
}

// NewDeleteManagedNamespaceParamsWithHTTPClient creates a new DeleteManagedNamespaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteManagedNamespaceParamsWithHTTPClient(client *http.Client) *DeleteManagedNamespaceParams {
	var ()
	return &DeleteManagedNamespaceParams{
		HTTPClient: client,
	}
}

/*DeleteManagedNamespaceParams contains all the parameters to send to the API endpoint
for the delete managed namespace operation typically these are written to a http.Request
*/
type DeleteManagedNamespaceParams struct {

	/*Customer*/
	Customer string
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete managed namespace params
func (o *DeleteManagedNamespaceParams) WithTimeout(timeout time.Duration) *DeleteManagedNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete managed namespace params
func (o *DeleteManagedNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete managed namespace params
func (o *DeleteManagedNamespaceParams) WithContext(ctx context.Context) *DeleteManagedNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete managed namespace params
func (o *DeleteManagedNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete managed namespace params
func (o *DeleteManagedNamespaceParams) WithHTTPClient(client *http.Client) *DeleteManagedNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete managed namespace params
func (o *DeleteManagedNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomer adds the customer to the delete managed namespace params
func (o *DeleteManagedNamespaceParams) WithCustomer(customer string) *DeleteManagedNamespaceParams {
	o.SetCustomer(customer)
	return o
}

// SetCustomer adds the customer to the delete managed namespace params
func (o *DeleteManagedNamespaceParams) SetCustomer(customer string) {
	o.Customer = customer
}

// WithName adds the name to the delete managed namespace params
func (o *DeleteManagedNamespaceParams) WithName(name string) *DeleteManagedNamespaceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete managed namespace params
func (o *DeleteManagedNamespaceParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteManagedNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param customer
	if err := r.SetPathParam("customer", o.Customer); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}