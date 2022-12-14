/*
 * Copyright (C) 2022 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */

package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"intel/amber/tac/v1/client/tms"
	"intel/amber/tac/v1/config"
	"intel/amber/tac/v1/constants"
	"intel/amber/tac/v1/models"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/cobra"
)

// createUserCmd represents the createUser command
var createUserCmd = &cobra.Command{
	Use:   constants.UserCmd,
	Short: "Creates a new user under a tenant",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Info("create user called")
		response, err := createUser(cmd)
		if err != nil {
			return err
		}
		fmt.Println("User: \n\n", response)
		return nil
	},
}

func init() {
	createCmd.AddCommand(createUserCmd)

	createUserCmd.Flags().StringVarP(&apiKey, constants.ApiKeyParamName, "a", "", "API key to be used to connect to amber services")
	createUserCmd.Flags().StringP(constants.TenantIdParamName, "t", "", "Id of the tenant for whom the user needs to be created")
	createUserCmd.Flags().StringP(constants.EmailIdParamName, "e", "", "Email id of the tenant user to be created")
	createUserCmd.Flags().StringP(constants.UserRoleParamName, "r", "", "Role of the tenant user to be created, should be one of Tenant Admin/User")
	createUserCmd.MarkFlagRequired(constants.ApiKeyParamName)
	createUserCmd.MarkFlagRequired(constants.EmailIdParamName)
	createUserCmd.MarkFlagRequired(constants.UserRoleParamName)
}

func createUser(cmd *cobra.Command) (string, error) {
	configValues, err := config.LoadConfiguration()
	if err != nil {
		return "", err
	}
	client := &http.Client{
		Timeout: time.Duration(configValues.HTTPClientTimeout) * time.Second,
	}

	tmsUrl, err := url.Parse(configValues.AmberBaseUrl + constants.TmsBaseUrl)
	if err != nil {
		return "", err
	}

	tenantIdString, err := cmd.Flags().GetString(constants.TenantIdParamName)
	if err != nil {
		return "", err
	}

	if tenantIdString == "" {
		tenantIdString = configValues.TenantId
	}

	tenantId, err := uuid.Parse(tenantIdString)
	if err != nil {
		return "", errors.Wrap(err, "Invalid tenant id provided")
	}

	emailId, err := cmd.Flags().GetString(constants.EmailIdParamName)
	if err != nil {
		return "", err
	}

	role, err := cmd.Flags().GetString(constants.UserRoleParamName)
	if err != nil {
		return "", err
	}

	var createUserInfo = &models.CreateTenantUser{
		Email: emailId,
		Role:  role,
	}

	tmsClient := tms.NewTmsClient(client, tmsUrl, tenantId, apiKey)
	response, err := tmsClient.CreateUser(createUserInfo)
	if err != nil {
		return "", err
	}

	responseBytes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		return "", err
	}

	return string(responseBytes), nil
}
