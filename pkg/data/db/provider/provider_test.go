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
	"testing"

	"carbonaut.cloud/carbonaut/pkg/data/db/provider/postgres"
	"carbonaut.cloud/carbonaut/pkg/data/db/provider/sqlite"
	"github.com/stretchr/testify/assert"
)

var (
	negCfg = []Config{{
		Name:         sqlite.Name,
		SqliteConfig: sqlite.Config{DatabaseFileName: "testdata/does-not-exist-db"},
	}, {
		Name: postgres.Name,
	}, {
		Name:         "does-not-exist-provider",
		SqliteConfig: sqlite.Config{DatabaseFileName: "testdata/does-not-exist-db"},
	}}
	posCfg = []Config{{
		Name:         sqlite.Name,
		SqliteConfig: sqlite.Config{DatabaseFileName: "testdata/emptytest.db"},
	}, {
		Name: postgres.Name,
		PostgresConfig: postgres.Config{
			Port:         5430,
			Password:     "some-password",
			Host:         "127.0.0.1",
			User:         "test",
			DatabaseName: "postgres",
			SSLMode:      "enable",
		},
	}}
)

func TestResolveProviderPos(t *testing.T) {
	for i := range posCfg {
		db, err := ResolveProvider(&posCfg[i])
		assert.NoError(t, err, fmt.Sprintf("expected error for provider %s on execution %d", posCfg[i].Name, i))
		assert.NotNil(t, db, fmt.Sprintf("expected nil for provider %s on execution %d", posCfg[i].Name, i))
	}
}

func TestResolveProviderNeg(t *testing.T) {
	for i := range negCfg {
		db, err := ResolveProvider(&negCfg[i])
		assert.Error(t, err, fmt.Sprintf("expected error for provider %s on execution %d", negCfg[i].Name, i))
		assert.Nil(t, db, fmt.Sprintf("expected nil for provider %s on execution %d", negCfg[i].Name, i))
	}
}
