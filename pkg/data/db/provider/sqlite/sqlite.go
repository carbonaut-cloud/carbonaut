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
)

type Config struct {
	DatabaseFileName string
}

type P struct {
	Name string
}

var Provider = P{
	Name: "sqlite",
}

// func (p P) Connect(cfg interface{}) (methods.ICarbonDB, error) {
// 	if _, err := os.Stat(cfg.DatabaseFileName); errors.Is(err, os.ErrNotExist) {
// 		return nil, err
// 	}
// 	// open connection to db file
// 	db, err := gorm.Open(sqlite.Open(cfg.DatabaseFileName), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	// migrate tables
// 	var carbonDB methods.CarbonDB
// 	carbonDB.Init(db)
// 	if err := carbonDB.Migrate(); err != nil {
// 		return nil, err
// 	}
// 	return carbonDB, nil
// }

func (p P) Validate(cfg interface{}) error {
	return fmt.Errorf("not implemented yet")
}
