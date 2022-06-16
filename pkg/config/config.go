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

package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"carbonaut.cloud/carbonaut/pkg/connector"
	"carbonaut.cloud/carbonaut/pkg/data"
	"carbonaut.cloud/carbonaut/pkg/util"
	"github.com/caarlos0/env"
	"github.com/go-playground/validator/v10"
	"github.com/mcuadros/go-defaults"
	"gopkg.in/yaml.v3"
)

// configProvider define how the decode the provided configuration
type configProvider string

// FileConfigMedium The configuration gets provided by file
const (
	FileConfigMedium    configProvider = "file"
	DefaultConfigMedium configProvider = "default"
)

// NOTE: it would be possible to define remote kv-stores later that hold the configuration: https://github.com/spf13/viper#remote-keyvalue-store-support

// Entire carbonaut configuration file which is largely specified in sub structs
type CarbonConfig struct {
	Connector connector.Config `yaml:"connector"`
	Data      data.Config      `yaml:"data"`
}

type ReadFileIn struct {
	ConfigFilePath string `env:"CONFIG_FILE_PATH" default:"./carbonconfig.yaml" validate:"file"`
}

// VerifyConfig check if the carbonaut configuration is valid
func VerifyConfig(cfg *CarbonConfig) error {
	if err := validator.New().Struct(cfg); err != nil {
		return fmt.Errorf("invalid configuration, %v", err)
	}
	return nil
}

// ReadFile read the carbonaut configuration file
func ReadFile() (*CarbonConfig, error) {
	util.Log.Info().Msg("read carbonaut configuration file")
	r := &ReadFileIn{}
	if err := env.Parse(r); err != nil {
		return nil, fmt.Errorf("could not get environment variables %v", err)
	}
	defaults.SetDefaults(r)
	if err := validator.New().Struct(r); err != nil {
		return nil, fmt.Errorf("invalid configuration, %v", err)
	}
	configFileBytes, err := ioutil.ReadFile(r.ConfigFilePath)
	if err != nil {
		return nil, fmt.Errorf("config file not found at %s, %v", r.ConfigFilePath, err)
	}

	carbonConfig := CarbonConfig{}
	if err := yaml.Unmarshal(configFileBytes, &carbonConfig); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	if err := VerifyConfig(&carbonConfig); err != nil {
		return nil, err
	}
	return &carbonConfig, nil
}

type WriteFileIn struct {
	ConfigFilePath string       `default:"./carbonconfig.yaml"`
	Config         CarbonConfig `validate:"required"`
	AutoCreate     bool
}

// WriteFile create the carbonaut configuration file
func WriteFile(in *WriteFileIn) error {
	util.Log.Info().Msg("write carbonaut configuration file")
	defaults.SetDefaults(in)
	if err := validator.New().Struct(in); err != nil {
		return fmt.Errorf("invalid input to write the carbonaut config, %v", err)
	}
	configFileBytes, err := yaml.Marshal(in.Config)
	if err != nil {
		return fmt.Errorf("could not marshal config struct to bytes, %v", err)
	}
	// check if the file exists
	if _, err := os.Stat(in.ConfigFilePath); errors.Is(err, os.ErrNotExist) {
		util.Log.Info().Msg("carbonaut configuration file not found")
		if in.AutoCreate {
			util.Log.Info().Msg("auto create new carbonaut configuration file")
			f, err := os.OpenFile(in.ConfigFilePath, os.O_WRONLY|os.O_CREATE, 0o666)
			if err != nil {
				return fmt.Errorf("could not open or create a file at %s : %w", in.ConfigFilePath, err)
			}
			f.Close()
		} else {
			return fmt.Errorf("unable to write configuration file to not existing file, enable auto create carbonaut configuration to auto create the file")
		}
	}
	if err := ioutil.WriteFile(in.ConfigFilePath, configFileBytes, 0o600); err != nil {
		return fmt.Errorf("could not write config file, %v", err)
	}
	return nil
}
