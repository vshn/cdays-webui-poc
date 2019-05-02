// Code generated by go-swagger; DO NOT EDIT.

package namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new namespace API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for namespace API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateManagedNamespace create managed namespace API
*/
func (a *Client) CreateManagedNamespace(params *CreateManagedNamespaceParams) (*CreateManagedNamespaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateManagedNamespaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createManagedNamespace",
		Method:             "PUT",
		PathPattern:        "/namespace/{customer}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateManagedNamespaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateManagedNamespaceOK), nil

}

/*
DeleteManagedNamespace delete managed namespace API
*/
func (a *Client) DeleteManagedNamespace(params *DeleteManagedNamespaceParams) (*DeleteManagedNamespaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteManagedNamespaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteManagedNamespace",
		Method:             "DELETE",
		PathPattern:        "/namespace/{customer}/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteManagedNamespaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteManagedNamespaceOK), nil

}

/*
GetManagedNamespace get managed namespace API
*/
func (a *Client) GetManagedNamespace(params *GetManagedNamespaceParams) (*GetManagedNamespaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetManagedNamespaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getManagedNamespace",
		Method:             "GET",
		PathPattern:        "/namespace/{customer}/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetManagedNamespaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetManagedNamespaceOK), nil

}

/*
GetManagedNamespaces get managed namespaces API
*/
func (a *Client) GetManagedNamespaces(params *GetManagedNamespacesParams) (*GetManagedNamespacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetManagedNamespacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getManagedNamespaces",
		Method:             "GET",
		PathPattern:        "/namespaces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetManagedNamespacesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetManagedNamespacesOK), nil

}

/*
UpdateManagedNamespace update managed namespace API
*/
func (a *Client) UpdateManagedNamespace(params *UpdateManagedNamespaceParams) (*UpdateManagedNamespaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateManagedNamespaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateManagedNamespace",
		Method:             "POST",
		PathPattern:        "/namespace/{customer}/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateManagedNamespaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateManagedNamespaceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
