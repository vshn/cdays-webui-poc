//
// Copyright (c) 2019, VSHN AG, info@vshn.ch
// Licensed under "BSD 3-Clause". See LICENSE file.
//
//

package main

import (
	"log"
	"net/http"

	"github.com/foolin/goview/supports/echoview"
	"github.com/go-openapi/strfmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vshn/cdays-webui-poc/client/operations"

	httptransport "github.com/go-openapi/runtime/client"
	apiclient "github.com/vshn/cdays-webui-poc/client"
)

func main() {

	// Setup WebAPI Client
	client := apiclient.New(httptransport.New("127.0.0.1:8080", "", nil), strfmt.Default)

	resp, err := client.Operations.GetManagedNamespaces(operations.NewGetManagedNamespacesParams())
	if err != nil {
		log.Fatal(err)
	}

	// Initiate and configure Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Renderer = echoview.Default()

	// Routes
	e.GET("/", func(c echo.Context) error {
		//render with master
		return c.Render(http.StatusOK, "index", echo.Map{
			"title":      "APPUiO Management API",
			"namespaces": resp.Payload,
		})
	})

	e.GET("/page", func(c echo.Context) error {
		//render only file, must full name with extension
		return c.Render(http.StatusOK, "page.html", echo.Map{"title": "Page file title!!"})
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
