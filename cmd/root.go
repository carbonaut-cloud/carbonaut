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
	"carbonaut.cloud/carbonaut/pkg/util"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var cfg *Config

var rootCmd = &cobra.Command{
	Use:   "carbonaut",
	Short: "Run carbonaut commands",
	Long:  "TBD",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		logger, err := util.GetDefaultZapCfg().Build()
		if err != nil {
			return err
		}
		cfg = &Config{
			logger: logger.Sugar(),
		}
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return cfg.logger.Sync()
	},
}

func init() {
	rootCmd.AddCommand(dataCmd)
	dataCmd.AddCommand(dataExportCmd)
	dataCmd.AddCommand(dataImportCmd)
	dataCmd.AddCommand(dataReportCmd)

	rootCmd.AddCommand(deployCmd)
	deployCmd.AddCommand(deployDescribeCmd)
}

type Config struct {
	logger *zap.SugaredLogger
}

// Execute executes the ci-reporter root command.
func Execute() error {
	return rootCmd.Execute()
}
