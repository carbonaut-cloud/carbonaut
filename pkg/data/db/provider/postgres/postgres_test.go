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
	"testing"

	"github.com/stretchr/testify/assert"
)

var negCfg = []Config{
	{
		DatabaseName: "postgres-neg1",
	},
	{
		Port:         5430,
		Password:     "some-password",
		Host:         "127.0.0.1",
		User:         "test",
		DatabaseName: "postgres",
		SSLMode:      "enable",
	},
}

func TestConnectNeg(t *testing.T) {
	for _, cfg := range negCfg {
		db, err := cfg.Connect()
		assert.Error(t, err, fmt.Sprintf("expected error for database file %s", cfg.DatabaseName))
		assert.Nil(t, db, fmt.Sprintf("expected nil for database file %s", cfg.DatabaseName))
	}
}

func TestConnString(t *testing.T) {
	c := Config{
		Port:         5432,
		Password:     "some-password",
		Host:         "127.0.0.1",
		User:         "test",
		DatabaseName: "postgres",
		SSLMode:      "enable",
	}

	str := c.connString()
	assert.Equal(t, fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		c.Host, c.User, c.Password, c.DatabaseName, c.Port, c.SSLMode,
	), str)
}
