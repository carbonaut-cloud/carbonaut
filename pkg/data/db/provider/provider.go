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

package provider

import (
	"fmt"

	"carbonaut.cloud/carbonaut/pkg/data/db/methods"
	"carbonaut.cloud/carbonaut/pkg/data/db/provider/postgres"
	"carbonaut.cloud/carbonaut/pkg/data/db/provider/sqlite"
)

type Config struct {
	Name           string `validate:"nonzero"`
	PostgresConfig postgres.Config
	SqliteConfig   sqlite.Config
}

type IProvider interface {
	Connect() (methods.ICarbonDB, error)
	ValidateConfig() error
}

func ResolveProvider(c *Config) (IProvider, error) {
	switch c.Name {
	case sqlite.Name:
		if err := c.SqliteConfig.ValidateConfig(); err != nil {
			return nil, err
		}
		return &c.SqliteConfig, nil
	case postgres.Name:
		if err := c.PostgresConfig.ValidateConfig(); err != nil {
			return nil, err
		}
		return &c.PostgresConfig, nil
	default:
		return nil, fmt.Errorf("specified provider %s is not supported", c.Name)
	}
}
