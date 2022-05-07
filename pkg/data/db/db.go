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

package db

// A ORM library is used to connect to the database: see https://gorm.io/docs/

import (
	"fmt"

	"carbonaut.cloud/carbonaut/pkg/data/db/methods"
	"carbonaut.cloud/carbonaut/pkg/data/db/provider/postgres"
	"carbonaut.cloud/carbonaut/pkg/data/db/provider/sqlite"
)

func ValidateConfig(cfg *Config) error {
	return fmt.Errorf("provided database configuration is invalid")
}

type Config struct {
	Provider string `validate:"nonzero"`
	// Provider config gets dynamically set over the provider string specification
	ProviderConfig interface{}
}

// establish a connection to the configured database
func Connect(cfg *Config) (methods.ICarbonDB, error) {
	provider, err := resolveProvider(cfg)
	if err != nil {
		return nil, err
	}
	provider.Connect(cfg)
	return nil, fmt.Errorf("not implemented")
}

// resolve which database provider is specified in the configuration
func resolveProvider(cfg *Config) (IProvider, error) {
	switch cfg.Provider {
	case sqlite.Provider.Name:
		fmt.Println("OS X.")
	case postgres.Provider.Name:
		fmt.Println("OS X.")
		// postgres.Provider.Connect(cfg.ProviderConfig)
	default:
		return nil, fmt.Errorf("specified provider %s is not supported", cfg.Provider)
	}
	return nil, fmt.Errorf("not implemented")
}

// DatabaseDriver
type DatabaseDriver string

type IProvider interface {
	Connect(cfg interface{}) (methods.ICarbonDB, error)
	Validate(cfg interface{}) error
}
