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
	"intel/amber/tac/v1/client/pms"
	"intel/amber/tac/v1/config"
	"intel/amber/tac/v1/constants"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/cobra"
)

// getPoliciesCmd represents the getPolicies command
var getPoliciesCmd = &cobra.Command{
	Use:   constants.PolicyCmd,
	Short: "Get list of policies or specific policy user a tenant",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Info("list policies called")
		response, err := getPolicies(cmd)
		if err != nil {
			return err
		}
		fmt.Println("Policies: \n\n", response)
		return nil
	},
}

func init() {
	listCmd.AddCommand(getPoliciesCmd)

	getPoliciesCmd.Flags().StringVarP(&apiKey, constants.ApiKeyParamName, "a", "", "API key to be used to connect to amber services")
	getPoliciesCmd.Flags().StringP(constants.PolicyIdParamName, "p", "", "Path of the file containing the policy to be uploaded")
	getPoliciesCmd.MarkFlagRequired(constants.ApiKeyParamName)
}

func getPolicies(cmd *cobra.Command) (string, error) {

	configValues, err := config.LoadConfiguration()
	if err != nil {
		return "", err
	}
	client := &http.Client{
		Timeout: time.Duration(configValues.HTTPClientTimeout) * time.Second,
	}

	pmsUrl, err := url.Parse(configValues.AmberBaseUrl + constants.PmsBaseUrl)
	if err != nil {
		return "", err
	}

	policyIdString, err := cmd.Flags().GetString(constants.PolicyIdParamName)
	if err != nil {
		return "", err
	}

	pmsClient := pms.NewPmsClient(client, pmsUrl, uuid.Nil, apiKey)

	var responseBytes []byte
	if policyIdString == "" {
		response, err := pmsClient.SearchPolicy()
		if err != nil {
			return "", err
		}
		responseBytes, err = json.MarshalIndent(response, "", "  ")
		if err != nil {
			return "", err
		}
	} else {
		policyId, err := uuid.Parse(policyIdString)
		if err != nil {
			return "", errors.Wrap(err, "Invalid policy id provided")
		}
		response, err := pmsClient.GetPolicy(policyId)
		if err != nil {
			return "", err
		}

		responseBytes, err = json.MarshalIndent(response, "", "  ")
		if err != nil {
			return "", err
		}
	}

	return string(responseBytes), nil
}
