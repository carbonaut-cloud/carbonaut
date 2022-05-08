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
	"carbonaut.cloud/carbonaut/pkg/data/db/methods"
	"carbonaut.cloud/carbonaut/pkg/data/db/provider"
)

// Provider config gets dynamically set over the provider string specification
type Config struct {
	Provider provider.Config
}

// establish a connection to the configured database
func Connect(cfg *Config) (methods.ICarbonDB, error) {
	p, err := provider.ResolveProvider(cfg.Provider)
	if err != nil {
		return nil, err
	}
	return p.Connect()
}
