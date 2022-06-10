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

package connector

import (
	"io/ioutil"
	"os"
	"testing"

	"carbonaut.cloud/carbonaut/pkg/connector/cloud/gcp"
	"carbonaut.cloud/carbonaut/pkg/connector/provider"
	"carbonaut.cloud/carbonaut/pkg/data"
	"carbonaut.cloud/carbonaut/pkg/data/methods"
	"carbonaut.cloud/carbonaut/pkg/data/storage"
	"carbonaut.cloud/carbonaut/pkg/data/storage/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	posGCPFilePath = "./testdata/gcp-carbon-data-extract-1.csv"
	negFilePath    = "./testdata/not-a-csv.json"
	doesNotExist   = "does-not-exist"
)

// This function establishes a connection to the specified provider to pull data until the Connect func stops
func TestConnectNeg(t *testing.T) {
	assert.Error(t, Connect(Config{
		Providers: []provider.Config{},
	}))
}

func TestGetProviderPos(t *testing.T) {
	gcpProvider, err := getProvider(gcp.Provider.Name)
	assert.NotNil(t, gcpProvider)
	assert.NoError(t, err)
}

func TestGetProviderNeg(t *testing.T) {
	gcpProvider, err := getProvider(provider.Name("does-not-exist"))
	assert.Nil(t, gcpProvider)
	assert.Error(t, err)
}

// Tests with DB

type ConnectorDBTestSuite struct {
	suite.Suite
	carbonautDB       methods.ICarbonDB
	dbDestinationFile string
}

// TestCarbonDB runs the entire 'CarbonDBTestSuite' test suite
func TestCarbonDB(t *testing.T) {
	suite.Run(t, new(ConnectorDBTestSuite))
}

// SetupTest gets called automatically before each test of the suite to setup the test environment
func (s *ConnectorDBTestSuite) SetupTest() {
	sourceFile := "./testdata/empty.db"
	destinationFile := "./testdata/empty2.db"
	// copy file from source to target destination, clean testing file
	input, err := ioutil.ReadFile(sourceFile)
	assert.NoError(s.T(), err)
	err = ioutil.WriteFile(destinationFile, input, 0o600)
	assert.NoError(s.T(), err)
	db, err := data.Connect(&data.Config{
		Storage: storage.Config{
			ProviderName: sqlite.Name,
			SqliteConfig: sqlite.Config{FileName: destinationFile, AutoCreate: true},
		},
	})
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), db)
	s.dbDestinationFile = destinationFile
	s.carbonautDB = db
}

// AfterTest gets called automatically after each test of the suite to clean up the test environment / add checks
func (s *ConnectorDBTestSuite) AfterTest(_, _ string) {
	assert.NoError(s.T(), os.Remove(s.dbDestinationFile), "clean up database file")
}

func (s *ConnectorDBTestSuite) TestImportDataPos() {
	getFileBytes := func(t *testing.T, filePath string) []byte {
		in, err := os.Open(filePath)
		assert.NoError(t, err)
		defer in.Close()
		data, err := ioutil.ReadAll(in)
		assert.NoError(t, err)
		return data
	}
	impInputs := []ImportIn{{
		ProviderName: gcp.Provider.Name,
		ImportType:   CsvRawImport,
		Data:         getFileBytes(s.T(), posGCPFilePath),
		DB:           s.carbonautDB,
	}}
	for i := range impInputs {
		err := ImportData(impInputs[i])
		assert.NoError(s.T(), err)
	}
}

func (s *ConnectorDBTestSuite) TestImportDataNeg() {
	impInputs := []ImportIn{{
		ProviderName: gcp.Provider.Name,
		ImportType:   CsvRawImport,
		DB:           s.carbonautDB,
	}, {
		ProviderName: gcp.Provider.Name,
		ImportType:   importType(doesNotExist),
		DB:           s.carbonautDB,
	}, {
		ProviderName: provider.Name(doesNotExist),
		ImportType:   CsvRawImport,
		Data:         []byte{1, 0, 1},
		DB:           s.carbonautDB,
	}}
	for i := range impInputs {
		err := ImportData(impInputs[i])
		assert.Error(s.T(), err)
	}
}
