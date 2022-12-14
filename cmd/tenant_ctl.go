/*
 * Copyright (C) 2022 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */

package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"intel/amber/tac/v1/config"
	"intel/amber/tac/v1/constants"
	"intel/amber/tac/v1/utils"
	"os"
)

var (
	apiKey string
)

// tenantCmd represents the base command when called without any subcommands
var tenantCmd = &cobra.Command{
	Use:   constants.RootCmd,
	Short: "Tenant CLI used to run the tasks for tenant admin/user",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the tenantCmd.
func Execute() {
	logFile, err := os.OpenFile(constants.LogFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, constants.DefaultFilePermission)
	if err != nil {
		fmt.Println("Error opening/creating log file: " + err.Error())
		os.Exit(1)
	}
	tenantCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		configValues, err := config.LoadConfiguration()
		if err != nil {
			if err := utils.SetUpLogs(logFile, constants.DefaultLogLevel); err != nil {
				return err
			}
		} else {
			if err := utils.SetUpLogs(logFile, configValues.LogLevel); err != nil {
				return err
			}
		}
		return nil
	}
	err = tenantCmd.Execute()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}
