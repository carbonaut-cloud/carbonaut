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

package methods

import (
	"io/ioutil"
	"os"
	"testing"

	"carbonaut.cloud/carbonaut/pkg/data/db/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CarbonDBTestSuite struct {
	suite.Suite
	DB                *gorm.DB
	carbonautDB       ICarbonDB
	dbDestinationFile string
}

// TestCarbonDB runs the entire 'CarbonDBTestSuite' test suite
func TestCarbonDB(t *testing.T) {
	suite.Run(t, new(CarbonDBTestSuite))
}

// SetupTest gets called automatically before each test of the suite to setup the test environment
func (s *CarbonDBTestSuite) SetupTest() {
	sourceFile := "testdata/emptytest.db"
	destinationFile := "testdata/emptytest2.db"
	// copy file from source to target destination, clean testing file
	input, err := ioutil.ReadFile(sourceFile)
	assert.NoError(s.T(), err)
	err = ioutil.WriteFile(destinationFile, input, 0o600)
	assert.NoError(s.T(), err)

	db, err := gorm.Open(sqlite.Open(destinationFile), &gorm.Config{})
	var carbonDB CarbonDB
	carbonDB.Init(db)
	assert.NoError(s.T(), err)

	s.dbDestinationFile = destinationFile
	assert.NotNil(s.T(), carbonDB)
	s.carbonautDB = carbonDB
}

// AfterTest gets called automatically after each test of the suite to clean up the test environment / add checks
func (s *CarbonDBTestSuite) AfterTest(_, _ string) {
	assert.NoError(s.T(), os.Remove(s.dbDestinationFile), "clean up database file")
}

// mock test: Get(id uint) (*Emissions, error)
func (s *CarbonDBTestSuite) TestCarbonDBGet() {
	e := models.Emissions{
		ID:            1,
		ResourceName:  "somename",
		ResourceOwner: "user",
		MTCO2e:        0.2,
	}

	_, err := s.carbonautDB.Get(e.ID)
	assert.NoError(s.T(), err)
}

// mock test: Delete(id uint) error
func (s *CarbonDBTestSuite) TestCarbonDBDelete() {
	e := models.Emissions{
		ID:            1,
		ResourceName:  "somename",
		ResourceOwner: "user",
		MTCO2e:        0.2,
	}

	err := s.carbonautDB.Delete(e.ID)
	assert.NoError(s.T(), err)
}

// mock test: List(offset, limit int) ([]*Emissions, error)
func (s *CarbonDBTestSuite) TestCarbonDBList() {
	_, err := s.carbonautDB.List(0, 10)
	assert.NoError(s.T(), err)
}

// mock test: Migrate() error
func (s *CarbonDBTestSuite) TestCarbonDBMigrate() {
	err := s.carbonautDB.Migrate()
	assert.NoError(s.T(), err)
}

// mock test: SearchByResourceName(q string, offset, limit int) ([]*Emissions, error)
func (s *CarbonDBTestSuite) TestCarbonDBSearchByResourceName() {
	e := models.Emissions{
		ID:            1,
		ResourceName:  "somename",
		ResourceOwner: "user",
		MTCO2e:        0.2,
	}

	_, err := s.carbonautDB.SearchByResourceName(e.ResourceName, 1, 10)
	assert.NoError(s.T(), err)
}

// mock test: Save(emissions *Emissions) error
func (s *CarbonDBTestSuite) TestCarbonDBSave() {
	e := models.Emissions{
		ID:            1,
		ResourceName:  "somename",
		ResourceOwner: "user",
		MTCO2e:        0.2,
	}

	err := s.carbonautDB.Save(&e)
	assert.NoError(s.T(), err)
}
