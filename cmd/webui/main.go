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

func createNamespace(c echo.Context) error {
	nsName := c.FormValue("nsName")
	return c.Render(http.StatusCreated, "createns", echo.Map{
		"title":   "APPUiO Management API - Created Namespace " + nsName,
		"message": "Namespace " + nsName + " successfully created",
	})
}

func main() {

	// Setup WebAPI Client
	client := apiclient.New(httptransport.New("127.0.0.1:8080", "", nil), strfmt.Default)

	// Initiate and configure Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		Browse: false,
	}))
	e.Renderer = echoview.Default()

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", echo.Map{
			"title": "APPUiO Management API",
		})
	})
	e.GET("/clusters", func(c echo.Context) error {
		return c.Render(http.StatusOK, "clusters", echo.Map{
			"title": "APPUiO Management API - Clusters",
		})
	})
	e.GET("/namespaces", func(c echo.Context) error {
		resp, err := client.Operations.GetManagedNamespaces(operations.NewGetManagedNamespacesParams())
		if err != nil {
			log.Fatal(err)
		}
		return c.Render(http.StatusOK, "namespaces", echo.Map{
			"title":      "APPUiO Management API - Namespaces",
			"namespaces": resp.Payload,
			"add": func(a, b int) int {
				return a + b
			},
		})
	})
	e.GET("/createns", func(c echo.Context) error {
		return c.Render(http.StatusOK, "createns", echo.Map{
			"title":   "APPUiO Management API - Create Namespace",
			"message": nil,
		})
	})
	e.POST("/createns", createNamespace)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
