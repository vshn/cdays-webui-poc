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
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"github.com/vshn/cdays-webui-poc/client/cluster"
	"github.com/vshn/cdays-webui-poc/client/namespace"

	httptransport "github.com/go-openapi/runtime/client"
	apiclient "github.com/vshn/cdays-webui-poc/client"
	cls "github.com/vshn/cdays-webui-poc/pkg/client/cluster"
	ns "github.com/vshn/cdays-webui-poc/pkg/client/namespace"
)

var (
	config appConfig
)

type appConfig struct {
	Listen string
	API    string
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

	// Cluster management routes
	e.GET("/clusters", func(c echo.Context) error {
		resp, err := client.Cluster.GetAllClusters(cluster.NewGetAllClustersParams())
		if err != nil {
			log.Fatal(err)
		}
		return c.Render(http.StatusOK, "clusters", echo.Map{
			"title":    "APPUiO Management API - Clusters",
			"clusters": resp.Payload,
			"add": func(a, b int) int {
				return a + b
			},
		})
	})
	e.GET("/clscud", func(c echo.Context) error {
		return c.Render(http.StatusOK, "clscud", echo.Map{
			"title":       "APPUiO Management API - Create Managed Cluster",
			"action":      "Create",
			"message":     nil,
			"messagetype": nil,
		})
	})
	e.POST("/clscud", func(c echo.Context) error {
		return cls.CreateCluster(c, client)
	})

	// Namespace management routes
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
			"title":       "APPUiO Management API - Create Managed Namespace",
			"action":      "Create",
			"message":     nil,
			"messagetype": nil,
		})
	})
	e.POST("/nscud", func(c echo.Context) error {
		return ns.CreateNamespace(c, client)
	})
	e.GET("/nscud/delete/:name", func(c echo.Context) error {
		return ns.DeleteNamespace(c, client)
	})
	e.GET("/nscud/update/:customer/:name", func(c echo.Context) error {
		params := namespace.NewGetManagedNamespaceParams()
		params.SetCustomer(c.Param("customer"))
		params.SetName(c.Param("name"))
		resp, err := client.Namespace.GetManagedNamespace(params)

		if err != nil {
			log.Fatal(err)
		}
		return c.Render(http.StatusOK, "nscud", echo.Map{
			"title":       "APPUiO Management API - Update Managed Namespace",
			"action":      "Update",
			"message":     nil,
			"messagetype": nil,
			"namespace":   resp.Payload,
		})
	})
	e.POST("/nscud/update/:customer/:name", func(c echo.Context) error {
		return ns.UpdateNamespace(c, client)
	})

	// Start server
	e.Logger.Fatal(e.Start(config.Listen))
}
