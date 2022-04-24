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
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestConnectToSQLite(t *testing.T) {
	// 1. set up test env
	sourceFile := "emptytest.db"
	destinationFile := "emptytest2.db"
	// copy file from source to target destination, clean testing file
	input, err := ioutil.ReadFile(sourceFile)
	assert.NoError(t, err)
	err = ioutil.WriteFile(destinationFile, input, 0o600)
	assert.NoError(t, err)

	// 2. test
	db, err := ConnectToSQLite(&SQLiteConfig{
		DatabaseFileName: destinationFile,
	})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	assert.NoError(t, db.Migrate())

	// 3. clean up test file
	assert.NoError(t, os.Remove(destinationFile))
}

// CarbonDBTestSuite is used for testing with mocks
type CarbonDBTestSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	carbonautDB ICarbonDB
}

func TestCarbonDB(t *testing.T) {
	suite.Run(t, new(CarbonDBTestSuite))
}

// SetupTest gets called automatically before each test of the suite to setup the test environment
func (s *CarbonDBTestSuite) SetupTest() {
	var (
		db  *sql.DB
		err error
	)
	// create new mocked database
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)
	// open connection to mocked database
	s.DB, err = gorm.Open(postgres.New(postgres.Config{Conn: db}))
	require.NoError(s.T(), err)
	s.carbonautDB = carbonDB{
		db: s.DB.Debug(),
	}
}

// AfterTest gets called automatically after each test of the suite to clean up the test environment / add checks
func (s *CarbonDBTestSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

// test Get(id uint) (*Emissions, error)
func (s *CarbonDBTestSuite) TestCarbonDBGet() {
	e := Emissions{
		ID:            1,
		ResourceName:  "somename",
		ResourceOwner: "user",
		MTCO2e:        0.2,
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "emissions" WHERE id = $1`)).
		WithArgs(e.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "resource_name", "resource_owner", "mtco2e"}).
			AddRow(fmt.Sprintf("%d", e.ID), e.ResourceName, e.ResourceOwner, e.MTCO2e))

	res, err := s.carbonautDB.Get(e.ID)
	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&e, res))
}

// Delete(id uint) error
func (s *CarbonDBTestSuite) TestCarbonDBDelete() {
	e := Emissions{
		ID:            1,
		ResourceName:  "somename",
		ResourceOwner: "user",
		MTCO2e:        0.2,
	}
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`DELETE FROM "emissions" WHERE id = $1`)).
		WithArgs(e.ID)

	err := s.carbonautDB.Delete(e.ID)
	require.NoError(s.T(), err)
}

// Test func: List(offset, limit int) ([]*Emissions, error)
func (s *CarbonDBTestSuite) TestCarbonDBList() {
	e := Emissions{
		ID:            1,
		ResourceName:  "somename",
		ResourceOwner: "user",
		MTCO2e:        0.2,
	}
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "emissions" LIMIT 10`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "resource_name", "resource_owner", "mtco2e"}).
			AddRow(fmt.Sprintf("%d", e.ID), e.ResourceName, e.ResourceOwner, e.MTCO2e))

	l, err := s.carbonautDB.List(0, 10)
	require.Len(s.T(), l, 1)
	require.Equal(s.T(), e.ID, l[0].ID)
	require.NoError(s.T(), err)
}

// Test func: Migrate() error
func (s *CarbonDBTestSuite) TestCarbonDBMigrate() {
	s.mock.ExpectExec(regexp.QuoteMeta(
		`SELECT count(*) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA() AND table_name = %1 AND table_type = %2`)).
		WithArgs("emissions", "BASE TABLE")
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`CREATE TABLE "emissions" ("id" bigserial,"resource_name" text,"resource_owner" text,"mtco2e" decimal,PRIMARY KEY ("id"))`)).
		WillReturnRows()
	s.mock.ExpectCommit()
	err := s.carbonautDB.Migrate()
	require.NoError(s.T(), err)
}

// SearchByResourceName(q string, offset, limit int) ([]*Emissions, error)
func (s *CarbonDBTestSuite) TestCarbonDBSearchByResourceName() {
	e := Emissions{
		ID:            1,
		ResourceName:  "somename",
		ResourceOwner: "user",
		MTCO2e:        0.2,
	}
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "emissions" WHERE resource_name like $1 LIMIT 10 OFFSET 1`)).
		WithArgs(e.ResourceName).WillReturnRows(sqlmock.NewRows([]string{"id", "resource_name", "resource_owner", "mtco2e"}).
		AddRow(fmt.Sprintf("%d", e.ID), e.ResourceName, e.ResourceOwner, e.MTCO2e))

	_, err := s.carbonautDB.SearchByResourceName(e.ResourceName, 1, 10)
	require.NoError(s.T(), err)
}

// Save(emissions *Emissions) error
func (s *CarbonDBTestSuite) TestCarbonDBSave() {
	e := Emissions{
		ID:            1,
		ResourceName:  "somename",
		ResourceOwner: "user",
		MTCO2e:        0.2,
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "emissions" VALUES ($1,$2,$3,$4)`)).
		WithArgs(e.ID, e.ResourceName, e.ResourceOwner, e.MTCO2e).
		WillReturnRows(sqlmock.NewRows([]string{"id", "resource_name", "resource_owner", "mtco2e"}).
			AddRow(fmt.Sprintf("%d", e.ID), e.ResourceName, e.ResourceOwner, e.MTCO2e))
	s.mock.ExpectCommit()

	err := s.carbonautDB.Save(&e)
	require.NoError(s.T(), err)
}
