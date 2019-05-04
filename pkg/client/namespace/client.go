//
// Copyright (c) 2019, VSHN AG, info@vshn.ch
// Licensed under "BSD 3-Clause". See LICENSE file.
//
//

package namespace

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/swag"
	"github.com/labstack/echo"
	apiclient "github.com/vshn/cdays-webui-poc/client"
	"github.com/vshn/cdays-webui-poc/client/cluster"
	"github.com/vshn/cdays-webui-poc/client/namespace"
	models "github.com/vshn/cdays-webui-poc/models"
)

func CreateNamespace(c echo.Context, client *apiclient.Webapi, clusters *cluster.GetAllClustersOK) error {
	nsName := c.FormValue("nsName")
	nsCustomer := c.FormValue("nsCustomer")
	nsDescription := c.FormValue("nsDescription")
	clusterName := c.Param("name")

	params := namespace.NewCreateManagedNamespaceParams()
	params.SetCustomer(nsCustomer)
	params.SetClustername(clusterName)
	params.SetBody(&models.Namespace{
		Description: nsDescription,
		Name:        swag.String(nsName),
		Customer:    swag.String(nsCustomer),
	})

	_, err := client.Namespace.CreateManagedNamespace(params)

	// TODO API doesnt report errors
	msg := ""
	msgtype := ""
	if err != nil {
		msg = fmt.Sprintf("Failed! %v", err)
		msgtype = "danger"
	} else {
		msg = "Namespace " + nsName + " on " + clusterName + " successfully created"
		msgtype = "success"
	}

	return c.Render(http.StatusCreated, "nscud", echo.Map{
		"title":       "APPUiO Management API - Created Namespace " + nsName,
		"clusters":    clusters.Payload,
		"cluster":     c.Param("name"),
		"message":     msg,
		"messagetype": msgtype,
	})
}

func DeleteNamespace(c echo.Context, client *apiclient.Webapi, clusters *cluster.GetAllClustersOK) error {
	clusterName := c.Param("name")
	nsName := c.Param("nsname")
	customerName := c.Param("customer")

	params := namespace.NewDeleteManagedNamespaceParams()
	params.SetClustername(clusterName)
	params.SetCustomer(customerName)
	params.SetName(nsName)

	_, err := client.Namespace.DeleteManagedNamespace(params)

	// TODO API doesnt report errors
	msg := ""
	msgtype := ""
	if err != nil {
		msg = fmt.Sprintf("Failed! %v", err)
		msgtype = "danger"
	} else {
		msg = "Namespace " + nsName + " on " + clusterName + " successfully deleted"
		msgtype = "success"
	}

	return c.Render(http.StatusOK, "nscud", echo.Map{
		"title":       "APPUiO Management API - Deleted Namespace",
		"clusters":    clusters.Payload,
		"cluster":     clusterName,
		"message":     msg,
		"messagetype": msgtype,
	})
}

func UpdateNamespace(c echo.Context, client *apiclient.Webapi, clusters *cluster.GetAllClustersOK) error {
	clusterName := c.Param("name")
	nsName := c.Param("nsname")
	nsCustomer := c.Param("customer")
	nsDescription := c.FormValue("nsDescription")

	params := namespace.NewUpdateManagedNamespaceParams()
	params.SetClustername(clusterName)
	params.SetName(nsName)
	params.SetCustomer(nsCustomer)
	params.SetBody(&models.Namespace{
		Description: nsDescription,
		Name:        swag.String(nsName),
		Customer:    swag.String(nsCustomer),
	})

	_, err := client.Namespace.UpdateManagedNamespace(params)

	// TODO API doesnt report errors
	msg := ""
	msgtype := ""
	if err != nil {
		msg = fmt.Sprintf("Failed! %v", err)
		msgtype = "danger"
	} else {
		msg = "Namespace " + nsName + " on " + clusterName + " successfully updated"
		msgtype = "success"
	}

	return c.Render(http.StatusOK, "nscud", echo.Map{
		"title":       "APPUiO Management API - Updated Namespace",
		"clusters":    clusters.Payload,
		"cluster":     clusterName,
		"message":     msg,
		"messagetype": msgtype,
	})
}
