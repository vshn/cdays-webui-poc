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

// NewGetManagedNamespacesParams creates a new GetManagedNamespacesParams object
// with the default values initialized.
func NewGetManagedNamespacesParams() *GetManagedNamespacesParams {

	return &GetManagedNamespacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetManagedNamespacesParamsWithTimeout creates a new GetManagedNamespacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetManagedNamespacesParamsWithTimeout(timeout time.Duration) *GetManagedNamespacesParams {

	return &GetManagedNamespacesParams{

		timeout: timeout,
	}
}

// NewGetManagedNamespacesParamsWithContext creates a new GetManagedNamespacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetManagedNamespacesParamsWithContext(ctx context.Context) *GetManagedNamespacesParams {

	return &GetManagedNamespacesParams{

		Context: ctx,
	}
}

// NewGetManagedNamespacesParamsWithHTTPClient creates a new GetManagedNamespacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetManagedNamespacesParamsWithHTTPClient(client *http.Client) *GetManagedNamespacesParams {

	return &GetManagedNamespacesParams{
		HTTPClient: client,
	}
}

/*GetManagedNamespacesParams contains all the parameters to send to the API endpoint
for the get managed namespaces operation typically these are written to a http.Request
*/
type GetManagedNamespacesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get managed namespaces params
func (o *GetManagedNamespacesParams) WithTimeout(timeout time.Duration) *GetManagedNamespacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get managed namespaces params
func (o *GetManagedNamespacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get managed namespaces params
func (o *GetManagedNamespacesParams) WithContext(ctx context.Context) *GetManagedNamespacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get managed namespaces params
func (o *GetManagedNamespacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get managed namespaces params
func (o *GetManagedNamespacesParams) WithHTTPClient(client *http.Client) *GetManagedNamespacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get managed namespaces params
func (o *GetManagedNamespacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetManagedNamespacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
