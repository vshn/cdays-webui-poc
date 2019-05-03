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
	"github.com/vshn/cdays-webui-poc/client/namespace"
	models "github.com/vshn/cdays-webui-poc/models"
)

func CreateNamespace(c echo.Context, client *apiclient.Webapi) error {
	nsName := c.FormValue("nsName")
	nsCustomer := c.FormValue("nsCustomer")
	nsDescription := c.FormValue("nsDescription")

	params := namespace.NewCreateManagedNamespaceParams()
	params.SetCustomer(nsCustomer)
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
		msg = "Namespace " + nsName + " successfully created"
		msgtype = "success"
	}

	return c.Render(http.StatusCreated, "nscud", echo.Map{
		"title":       "APPUiO Management API - Created Namespace " + nsName,
		"message":     msg,
		"messagetype": msgtype,
	})
}

func DeleteNamespace(c echo.Context, client *apiclient.Webapi) error {
	nsName := c.Param("name")

	params := namespace.NewDeleteManagedNamespaceParams()
	params.SetCustomer("mobiliar")
	params.SetName(nsName)

	_, err := client.Namespace.DeleteManagedNamespace(params)

	// TODO API doesnt report errors
	msg := ""
	msgtype := ""
	if err != nil {
		msg = fmt.Sprintf("Failed! %v", err)
		msgtype = "danger"
	} else {
		msg = "Namespace " + nsName + " successfully deleted"
		msgtype = "success"
	}

	return c.Render(http.StatusOK, "nscud", echo.Map{
		"title":       "APPUiO Management API - Deleted Namespace",
		"message":     msg,
		"messagetype": msgtype,
	})
}

func UpdateNamespace(c echo.Context, client *apiclient.Webapi) error {
	nsName := c.Param("name")
	nsCustomer := c.Param("customer")
	nsDescription := c.FormValue("nsDescription")

	params := namespace.NewUpdateManagedNamespaceParams()
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
		msg = "Namespace " + nsName + " successfully updated"
		msgtype = "success"
	}

	return c.Render(http.StatusOK, "nscud", echo.Map{
		"title":       "APPUiO Management API - Updated Namespace",
		"message":     msg,
		"messagetype": msgtype,
	})
}
