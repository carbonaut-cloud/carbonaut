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

import (
	"testing"

	"carbonaut.cloud/carbonaut/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestDataExporterNeg(t *testing.T) {
	logger, err := util.GetDefaultZapCfg().Build()
	defer func() {
		err = logger.Sync()
	}()
	dbConnection, err := Connect(&InDbConnection{
		DatabaseFileName: "emptytest.db",
	}, logger.Sugar())
	assert.NoError(t, err)
	assert.NotNil(t, dbConnection)
	assert.NotEmpty(t, dbConnection.Name())
}
