/*
 * Copyright (C) 2022 Intel Corporation
 * SPDX-License-Identifier: BSD-3-Clause
 */

package cmd

import (
	"github.com/stretchr/testify/assert"
	"intel/amber/tac/v1/constants"
	"intel/amber/tac/v1/test"
	"testing"
)

func TestCreateTagCmd(t *testing.T) {
	server := test.MockServer(t)
	defer server.Close()
	test.SetupMockConfiguration(server.URL, tempConfigFile)

	tt := []struct {
		args        []string
		wantErr     bool
		description string
	}{
		{
			args:    []string{constants.CreateCmd, constants.TagCmd, "-a", "abc", "-n", "Test Tag"},
			wantErr: false,
		},
	}

	createCmd.AddCommand(createTagCmd)
	tenantCmd.AddCommand(createCmd)

	for _, tc := range tt {
		_, err := execute(t, tenantCmd, tc.args)

		if tc.wantErr == true {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
