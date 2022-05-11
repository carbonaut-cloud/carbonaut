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
	"carbonaut.cloud/carbonaut/pkg/data/importer"
	"github.com/spf13/cobra"
)

// TODO: import data to your carbonaut db

var dataImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import data to carbonaut",
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunDataImport()
	},
}

// RunDataImport run cli command `carbonaut data import``
func RunDataImport() error {
	_, err := importer.DataImporter(&importer.InImporter{})
	return err
}
