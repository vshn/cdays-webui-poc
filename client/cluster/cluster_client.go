// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new cluster API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cluster API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAllClusters get all clusters API
*/
func (a *Client) GetAllClusters(params *GetAllClustersParams) (*GetAllClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllClusters",
		Method:             "GET",
		PathPattern:        "/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllClustersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllClustersOK), nil

}

/*
RegisterCluster register cluster API
*/
func (a *Client) RegisterCluster(params *RegisterClusterParams) (*RegisterClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "registerCluster",
		Method:             "PUT",
		PathPattern:        "/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RegisterClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegisterClusterOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
