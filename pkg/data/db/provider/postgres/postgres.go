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

package postgres

import (
	"fmt"

	"carbonaut.cloud/carbonaut/pkg/data/db/methods"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Port         int
	Password     string
	Host         string
	User         string
	DatabaseName string
	SSLMode      SSLMode
}

type SSLMode string

const (
	SSLModeDisable = "disable"
	SSLModeEnable  = "enable"
)

type P struct {
	Name string
}

var Provider = P{
	Name: "postgres",
}

// Connect establishes a connection to the provided database configuration
// The database can hosted locally (or somewhere else)
// Locally you can use the officially postgres container image
// 1. podman run -it --rm -p 127.0.0.1:5432:5432/tcp -e POSTGRES_PASSWORD=test postgres
// 2. psql -d postgres -h localhost -U postgres
// 3. enter password: test
// Setting the same information in PostgresConfig to connect to the local hosted database
func (p P) Connect(cfg Config) (methods.ICarbonDB, error) {
	// open connection to db
	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
			cfg.Host, cfg.User, cfg.Password, cfg.DatabaseName, cfg.Port, cfg.SSLMode,
		)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// migrate tables
	var carbonDB methods.CarbonDB
	carbonDB.Init(db)
	if err := carbonDB.Migrate(); err != nil {
		return nil, err
	}
	return carbonDB, nil
}

func (p P) Validate(cfg *Config) error {
	return fmt.Errorf("not implemented yet")
}
