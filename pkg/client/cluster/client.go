//
// Copyright (c) 2019, VSHN AG, info@vshn.ch
// Licensed under "BSD 3-Clause". See LICENSE file.
//
//

package cluster

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/swag"
	"github.com/labstack/echo"
	apiclient "github.com/vshn/cdays-webui-poc/client"
	"github.com/vshn/cdays-webui-poc/client/cluster"
	"github.com/vshn/cdays-webui-poc/models"
)

func CreateCluster(c echo.Context, client *apiclient.Webapi) error {
	clusterName := c.FormValue("clusterName")
	clusterURL := c.FormValue("clusterURL")
	clusterToken := c.FormValue("clusterToken")

	params := cluster.NewRegisterClusterParams()
	params.SetBody(&models.Cluster{
		Name:  swag.String(clusterName),
		Token: clusterToken,
		URL:   swag.String(clusterURL),
	})

	_, err := client.Cluster.RegisterCluster(params)

	// TODO API doesnt report errors
	msg := ""
	msgtype := ""
	if err != nil {
		msg = fmt.Sprintf("Failed! %v", err)
		msgtype = "danger"
	} else {
		msg = "Cluster " + clusterName + " successfully registered"
		msgtype = "success"
	}

	return c.Render(http.StatusCreated, "clscud", echo.Map{
		"title":       "APPUiO Management API - Registered Cluster " + clusterName,
		"message":     msg,
		"messagetype": msgtype,
	})
}
