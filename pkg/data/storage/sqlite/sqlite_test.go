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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	negCfg = []Config{
		{},
		{FileName: "db file does not exist"},
		{FileName: "testdata/not-sqlite.yml"},
		{FileName: "//not-sqlite.yml", AutoCreate: true},
	}
	posCfg = []Config{{
		FileName: "testdata/emptytest.db",
	}, {
		FileName:   "testdata/does-not-exist.db",
		AutoCreate: true,
	}}
)

func TestConnectNeg(t *testing.T) {
	for _, cfg := range negCfg {
		db, err := cfg.Connect()
		assert.Error(t, err, fmt.Sprintf("expected error for database file %s", cfg.FileName))
		assert.Nil(t, db, fmt.Sprintf("expected nil for database file %s", cfg.FileName))
	}
}

func TestConnectPos(t *testing.T) {
	for _, cfg := range posCfg {
		if cfg.AutoCreate {
			defer cfg.Destroy()
		}
		db, err := cfg.Connect()
		assert.NoError(t, err, fmt.Sprintf("no error expected for database file %s", cfg.FileName))
		assert.NotNil(t, db, fmt.Sprintf("not nil expected for database file %s", cfg.FileName))
	}
}

func TestDestroyNeg(t *testing.T) {
	cfg := Config{FileName: "//wrong"}
	cfg.Destroy()
}
