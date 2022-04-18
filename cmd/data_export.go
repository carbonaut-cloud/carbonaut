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
	"carbonaut.cloud/carbonaut/pkg/data/exporter"
	"github.com/spf13/cobra"
)

// TODO: export data from your carbonaut db

var dataExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export data from carbonaut",
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunDataExport()
	},
}

// RunDataExport run cli command `carbonaut data export``
func RunDataExport() error {
	out, err := exporter.DataExporter(&exporter.InExporter{})
	if err != nil {
		return err
	}
	cfg.logger.Info(out)
	return nil
}
