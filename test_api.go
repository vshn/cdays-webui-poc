package main

import (
	"fmt"
	"log"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiclient "github.com/vshn/cdays-webui-poc/client"
	"github.com/vshn/cdays-webui-poc/client/operations"
)

func main() {

	// create the transport
	transport := httptransport.New("127.0.0.1:8080", "", nil)

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	// to override the host for the default client
	// apiclient.Default.SetTransport(transport)

	// make the request to get all items
	//resp, err := client.Operations.All(operations.AllParams{})
	resp, err := client.Operations.GetManagedNamespaces(operations.NewGetManagedNamespacesParams())

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%#v\n", resp.Payload)

	for _, ns := range resp.Payload {
		fmt.Printf("%v\n", *ns.Name)
	}
}
