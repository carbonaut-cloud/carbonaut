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
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	Use:   "carbonaut",
	Short: "Run commands against your Carbonaut deployments",
	Long:  "CarbonCtl controls your Carbonaut deployments like the main carbonaut pod, the cloud-connector, dashboards, databases and other components",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		cfg := Config{
			log: logger.Sugar(),
		}
		return Run(&cfg)
	},
}

type Config struct {
	log *zap.SugaredLogger
}

// Execute executes the ci-reporter root command.
func Execute() error {
	return rootCmd.Execute()
}

// Run executes the main logic
func Run(cfg *Config) error {
	cfg.log.Info("Run root carbonaut command")
	return nil
}
