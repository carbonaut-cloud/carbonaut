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

package sqlite

import (
	"errors"
	"fmt"
	"os"

	"carbonaut.cloud/carbonaut/pkg/data/db/methods"
	validator "gopkg.in/validator.v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	DatabaseFileName string `validate:"nonzero"`
}

const Name = "sqlite"

func (c *Config) Connect() (methods.ICarbonDB, error) {
	if err := c.ValidateConfig(); err != nil {
		return nil, err
	}
	// open connection to db file
	db, err := gorm.Open(sqlite.Open(c.DatabaseFileName), &gorm.Config{})
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

func (c *Config) ValidateConfig() error {
	// validate input if `validate:xxx` is specified - see https://github.com/go-validator/validator
	if err := validator.Validate(c); err != nil {
		return fmt.Errorf("provided configuration is not valid, %v", err)
	}
	if _, err := os.Stat(c.DatabaseFileName); errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}
