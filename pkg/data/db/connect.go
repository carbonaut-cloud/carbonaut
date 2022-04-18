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

// A ORM library is used to connect to the database: see https://gorm.io/docs/

import (
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type InDbConnection struct {
	DatabaseFileName string
}

//  TODO: connect to database
func Connect(in *InDbConnection, log *zap.SugaredLogger) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(in.DatabaseFileName), &gorm.Config{})
}
