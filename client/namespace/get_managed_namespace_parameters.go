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

// NewGetManagedNamespaceParams creates a new GetManagedNamespaceParams object
// with the default values initialized.
func NewGetManagedNamespaceParams() *GetManagedNamespaceParams {
	var ()
	return &GetManagedNamespaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetManagedNamespaceParamsWithTimeout creates a new GetManagedNamespaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetManagedNamespaceParamsWithTimeout(timeout time.Duration) *GetManagedNamespaceParams {
	var ()
	return &GetManagedNamespaceParams{

		timeout: timeout,
	}
}

// NewGetManagedNamespaceParamsWithContext creates a new GetManagedNamespaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetManagedNamespaceParamsWithContext(ctx context.Context) *GetManagedNamespaceParams {
	var ()
	return &GetManagedNamespaceParams{

		Context: ctx,
	}
}

// NewGetManagedNamespaceParamsWithHTTPClient creates a new GetManagedNamespaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetManagedNamespaceParamsWithHTTPClient(client *http.Client) *GetManagedNamespaceParams {
	var ()
	return &GetManagedNamespaceParams{
		HTTPClient: client,
	}
}

/*GetManagedNamespaceParams contains all the parameters to send to the API endpoint
for the get managed namespace operation typically these are written to a http.Request
*/
type GetManagedNamespaceParams struct {

	/*Clustername*/
	Clustername string
	/*Customer*/
	Customer string
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get managed namespace params
func (o *GetManagedNamespaceParams) WithTimeout(timeout time.Duration) *GetManagedNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get managed namespace params
func (o *GetManagedNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get managed namespace params
func (o *GetManagedNamespaceParams) WithContext(ctx context.Context) *GetManagedNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get managed namespace params
func (o *GetManagedNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get managed namespace params
func (o *GetManagedNamespaceParams) WithHTTPClient(client *http.Client) *GetManagedNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get managed namespace params
func (o *GetManagedNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClustername adds the clustername to the get managed namespace params
func (o *GetManagedNamespaceParams) WithClustername(clustername string) *GetManagedNamespaceParams {
	o.SetClustername(clustername)
	return o
}

// SetClustername adds the clustername to the get managed namespace params
func (o *GetManagedNamespaceParams) SetClustername(clustername string) {
	o.Clustername = clustername
}

// WithCustomer adds the customer to the get managed namespace params
func (o *GetManagedNamespaceParams) WithCustomer(customer string) *GetManagedNamespaceParams {
	o.SetCustomer(customer)
	return o
}

// SetCustomer adds the customer to the get managed namespace params
func (o *GetManagedNamespaceParams) SetCustomer(customer string) {
	o.Customer = customer
}

// WithName adds the name to the get managed namespace params
func (o *GetManagedNamespaceParams) WithName(name string) *GetManagedNamespaceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get managed namespace params
func (o *GetManagedNamespaceParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetManagedNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clustername
	if err := r.SetPathParam("clustername", o.Clustername); err != nil {
		return err
	}

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
