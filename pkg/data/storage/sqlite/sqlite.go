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

	"carbonaut.cloud/carbonaut/pkg/data/methods"
	"carbonaut.cloud/carbonaut/pkg/util"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	FileName   string `default:"tmp.db"`
	AutoCreate bool   `default:"true"`
}

const Name = "sqlite"

func (c *Config) Connect() (methods.ICarbonDB, error) {
	if err := c.ValidateConfig(); err != nil {
		return nil, err
	}
	// create an empty sqlite database file if it does not exist
	if _, err := os.Stat(c.FileName); errors.Is(err, os.ErrNotExist) && c.AutoCreate {
		log.Info().Msg("Local database file not found, auto create sqlite database file")
		file, err := os.Create(c.FileName)
		if err != nil {
			return nil, err
		}
		file.Close()
		log.Info().Msgf("Local database file %s created", c.FileName)
	}
	// open connection to db file
	db, err := gorm.Open(sqlite.Open(c.FileName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// migrate tables
	var carbonDB methods.CarbonDB
	// init db model and migrate tables
	if err := carbonDB.Init(db); err != nil {
		return nil, err
	}
	return carbonDB, nil
}

func (c *Config) ValidateConfig() error {
	if c.FileName == "" {
		return fmt.Errorf("database filename is not set")
	}
	if _, err := os.Stat(c.FileName); errors.Is(err, os.ErrNotExist) && !c.AutoCreate {
		return err
	}
	return nil
}

func (c *Config) Destroy() {
	if err := os.Remove(c.FileName); err != nil {
		util.Log.Err(err)
	}
}
