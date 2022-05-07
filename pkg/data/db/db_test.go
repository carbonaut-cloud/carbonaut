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

// func TestValidateConfig(t *testing.T) {
// 	err := ValidateConfig(&Config{
// 		Provider:       "",
// 		PostgresConfig: PostgresConfig{Port: 0, Password: "", Host: "", User: "", DatabaseName: "", SSLMode: ""},
// 		SQLiteConfig:   SQLiteConfig{DatabaseFileName: ""},
// 	})
// 	assert.Error(t, err)
// }

// //
// // Database tests
// //
// //

// // CarbonDBTestSuite is used for testing with mocks
// type CarbonDBTestSuite struct {
// 	suite.Suite
// 	DB                *gorm.DB
// 	carbonautDB       ICarbonDB
// 	dbDestinationFile string
// }

// //
// // Tests to establish a db connection
// //

// func TestConnectToSqliteNeg(t *testing.T) {
// 	db, err := ConnectToSQLite(&SQLiteConfig{
// 		DatabaseFileName: "db file does not exist",
// 	})
// 	assert.Error(t, err)
// 	assert.Nil(t, db)
// }

// func TestConnectToSqliteNoAccessNeg(t *testing.T) {
// 	db, err := ConnectToSQLite(&SQLiteConfig{
// 		DatabaseFileName: "testdata/not-a-db",
// 	})
// 	assert.Error(t, err)
// 	assert.Nil(t, db)
// }

// func TestConnectToSqliteWrongSchemaNeg(t *testing.T) {
// 	db, err := ConnectToSQLite(&SQLiteConfig{
// 		DatabaseFileName: "testdata/not-sqlite.yml",
// 	})
// 	assert.Error(t, err)
// 	assert.Nil(t, db)
// }

// func TestConnectToPostgresNeg(t *testing.T) {
// 	// injecting a empty config with default val's should fail
// 	db, err := ConnectToPostgres(&PostgresConfig{})
// 	assert.Error(t, err)
// 	assert.Nil(t, db)
// }
