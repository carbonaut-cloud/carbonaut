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
	"fmt"

	"carbonaut.cloud/carbonaut/pkg/api"
	"carbonaut.cloud/carbonaut/pkg/connector"
	"carbonaut.cloud/carbonaut/pkg/data"
	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
	"gopkg.in/validator.v2"
)

// configProvider define how the decode the provided configuration
type configProvider string

// FileConfigMedium The configuration gets provided by file
const FileConfigMedium configProvider = "file"

// NOTE: it would be possible to define remote kv-stores later that hold the configuration: https://github.com/spf13/viper#remote-keyvalue-store-support

// Used to configure where to find the configuration file and how to parse it
type GetCarbonautConfigIn struct {
	ConfigMedium     configProvider `default:"file"`
	ConfigMediumFile *FileMediumConfig
}

type FileMediumConfig struct {
	FileName string `default:"carbonconfig"`
	FileType string `default:"yaml"`
	FilePath string
}

// Entire carbonaut configuration file which is largely specified in sub structs
type CarbonConfig struct {
	Version  string `default:"v1"`
	LogLevel string `default:"info" validate:"regexp=^info|debug|error|fatal|warning$"`

	API       api.Config       `mapstructure:"api"`
	Connector connector.Config `mapstructure:"connector"`
	Data      data.Config      `mapstructure:"data"`
}

// GetCarbonautConfig get the carbonaut configuration
func GetCarbonautConfig(in *GetCarbonautConfigIn) (*CarbonConfig, error) {
	defaults.SetDefaults(in)
	// get configuration from configured config medium
	carbonConfig, err := func() (*CarbonConfig, error) {
		switch in.ConfigMedium {
		case FileConfigMedium:
			return getFileConfiguration(in.ConfigMediumFile)
		default:
			return nil, fmt.Errorf("config medium %s not supported", in.ConfigMedium)
		}
	}()
	if err != nil {
		return nil, err
	}

	// apply defaults to empty configuration fields
	defaults.SetDefaults(carbonConfig)

	// validate input if `validate:xxx` is specified - see https://github.com/go-validator/validator
	if err := validator.Validate(carbonConfig); err != nil {
		return nil, fmt.Errorf("provided configuration is not valid, %v", err)
	}
	return carbonConfig, nil
}

// getFileConfiguration get the carbonaut configuration stored in a local file
func getFileConfiguration(in *FileMediumConfig) (*CarbonConfig, error) {
	viper.SetConfigName(in.FileName)
	viper.SetConfigType(in.FileType)
	viper.AddConfigPath(in.FilePath)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("config file %s not found at %s, %v", in.FileName, in.FilePath, err)
		}
		return nil, fmt.Errorf("unable to read in config file %s at %s, %v", in.FileName, in.FilePath, err)
	}

	carbonConfig := CarbonConfig{}
	if err := viper.Unmarshal(&carbonConfig); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return &carbonConfig, nil
}
