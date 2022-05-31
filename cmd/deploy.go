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
	"carbonaut.cloud/carbonaut/pkg/api"
	"carbonaut.cloud/carbonaut/pkg/config"
	"carbonaut.cloud/carbonaut/pkg/util"
	"github.com/mcuadros/go-defaults"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var cfg = config.GetCarbonautConfigIn{}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy carbonaut",
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunDeploy()
	},
}

// RunDeploy run cli command `carbonaut deploy``
func RunDeploy() error {
	c, err := config.GetCarbonautConfig(&cfg)
	if err != nil {
		log.Err(err)
		return err
	}
	a := api.CarbonautAPI{}
	if err := a.Start(&c.API); err != nil {
		log.Err(err)
		return err
	}
	return nil
}

func init() {
	l := util.GetLogger(&util.LogConfig{})
	l.Info().Msg("Applying defaults to the configuration to look up the carbonaut config file")
	defaults.SetDefaults(&cfg)
	deployCmd.PersistentFlags().StringVarP(&cfg.FilePath, "config", "c", "", "Specify where to find the configuration file")
	if cfg.FilePath != "" {
		cfg.ConfigMedium = config.FileConfigMedium
	}
}
