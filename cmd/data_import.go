// Copyright 2022 The Carbonaut Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dataImportCmd = &cobra.Command{
	Use:              "import",
	Short:            "Import data into carbonaut",
	SilenceUsage:     true,
	TraverseChildren: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunDataImport()
	},
}

func RunDataImport() error {
	return fmt.Errorf("data import gcp command is not implemented yet")
}

func init() {
	dataImportCmd.AddCommand(dataImportGcpCmd)
}

// GCP Import command

var dataImportGcpCmd = &cobra.Command{
	Use:              "gcp",
	Short:            "Import GCP data into carbonaut",
	SilenceUsage:     true,
	TraverseChildren: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunDataImportGcp()
	},
}

func RunDataImportGcp() error {
	return fmt.Errorf("data import gcp command is not implemented yet")
}
