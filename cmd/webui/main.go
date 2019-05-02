//
// Copyright (c) 2019, VSHN AG, info@vshn.ch
// Licensed under "BSD 3-Clause". See LICENSE file.
//
//

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/foolin/goview/supports/echoview"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"github.com/vshn/cdays-webui-poc/client/namespace"

	httptransport "github.com/go-openapi/runtime/client"
	apiclient "github.com/vshn/cdays-webui-poc/client"
	models "github.com/vshn/cdays-webui-poc/models"
)

var (
	config appConfig
)

type appConfig struct {
	Listen string
	API    string
}

func createNamespace(c echo.Context) error {
	nsName := c.FormValue("nsName")
	nsCustomer := c.FormValue("nsCustomer")
	nsDescription := c.FormValue("nsDescription")

	params := namespace.NewCreateManagedNamespaceParams()
	params.SetCustomer(nsCustomer)
	params.SetBody(&models.Namespace{
		Description: nsDescription,
		Name:        swag.String(nsName),
	})

	client := apiclient.New(httptransport.New(config.API, "", nil), strfmt.Default)
	_, err := client.Namespace.CreateManagedNamespace(params)

	// TODO API doesnt report errors
	msg := ""
	if err != nil {
		msg = "Failed!"
	} else {
		msg = "Namespace " + nsName + " successfully created"
	}

	return c.Render(http.StatusCreated, "nscud", echo.Map{
		"title":   "APPUiO Management API - Created Namespace " + nsName,
		"message": msg,
	})
}

func deleteNamespace(c echo.Context) error {
	nsName := c.Param("name")
	msg := "Namespace " + nsName + " successfully deleted"
	return c.Render(http.StatusOK, "nscud", echo.Map{
		"title":   "APPUiO Management API - Deleted Namespace",
		"message": msg,
	})
}

func main() {

	// Initiate and configure Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		Browse: false,
	}))
	e.Renderer = echoview.Default()

	// Read configuration file and set defaults
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("config file reading error: %v", err))
	}

	// unmarshal config into appConfig struct
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(fmt.Errorf("config file parsing error: %v", err))
	}

	if config.API == "" {
		e.Logger.Fatal("API endpoint not configured")
	}
	if config.Listen == "" {
		e.Logger.Info("listen not configured - defaulting to :1323")
		config.Listen = ":1323"
	}

	// Setup WebAPI Client
	client := apiclient.New(httptransport.New(config.API, "", nil), strfmt.Default)

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
		resp, err := client.Namespace.GetManagedNamespaces(namespace.NewGetManagedNamespacesParams())
		if err != nil {
			log.Fatal(err)
		}
		return c.Render(http.StatusOK, "namespaces", echo.Map{
			"title":      "APPUiO Management API - Managed Namespaces",
			"namespaces": resp.Payload,
			"add": func(a, b int) int {
				return a + b
			},
		})
	})
	e.GET("/nscud", func(c echo.Context) error {
		return c.Render(http.StatusOK, "nscud", echo.Map{
			"title":   "APPUiO Management API - Create Managed Namespace",
			"message": nil,
		})
	})
	e.POST("/nscud", createNamespace)
	e.GET("/nscud/delete/:name", deleteNamespace)

	// Start server
	e.Logger.Fatal(e.Start(config.Listen))
}
