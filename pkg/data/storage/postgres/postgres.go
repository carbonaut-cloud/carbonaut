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

	"carbonaut.cloud/carbonaut/pkg/data/methods"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Port         int `default:"5432"`
	Password     string
	Host         string `default:"127.0.0.1"`
	User         string
	DatabaseName string `default:"postgres"`
	SSLMode      string `default:"disable"`
}

const Name = "postgres"

// Connect establishes a connection to the provided database configuration
// The database can hosted locally (or somewhere else)
// Locally you can use the officially postgres container image
// 1. podman run -it --rm -p 127.0.0.1:5432:5432/tcp -e POSTGRES_PASSWORD=test postgres
// 2. psql -d postgres -h localhost -U postgres
// 3. enter password: test
// Setting the same information in PostgresConfig to connect to the local hosted database
func (c *Config) Connect() (methods.ICarbonDB, error) {
	var carbonDB methods.CarbonDB
	if err := c.ValidateConfig(); err != nil {
		return nil, err
	}
	// open connection to db
	db, err := gorm.Open(postgres.Open(c.connString()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// migrate tables
	carbonDB.Init(db)
	if err := carbonDB.Migrate(); err != nil {
		return nil, err
	}
	return carbonDB, nil
}

func (c *Config) ValidateConfig() error {
	if c.DatabaseName == "" || c.Host == "" || c.Password == "" || c.User == "" || c.Port == 0 {
		return fmt.Errorf("configuration invalid required parameters are not set")
	}
	if c.SSLMode != "enable" && c.SSLMode != "disable" && c.SSLMode != "" {
		return fmt.Errorf("SSL Mode must match 'enable' or 'disable' key")
	}
	return nil
}

func (c *Config) connString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		c.Host, c.User, c.Password, c.DatabaseName, c.Port, c.SSLMode,
	)
}
