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
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	Port         int
	Password     string
	Host         string
	User         string
	DatabaseName string
	SSLMode      SSLMode
}

type SQLiteConfig struct {
	DatabaseFileName string
}

type SSLMode string

const (
	SSLModeDisable = "disable"
	SSLModeEnable  = "enable"
)

// DatabaseDriver
type DatabaseDriver string

func ConnectToSQLite(cfg *SQLiteConfig) (ICarbonDB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DatabaseFileName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return carbonDB{db}, nil
}

// Connect establishes a connection to the provided database configuration
// The database can hosted locally (or somewhere else)
// Locally you can use the officially postgres container image
// 1. podman run -it --rm -p 127.0.0.1:5432:5432/tcp -e POSTGRES_PASSWORD=test postgres
// 2. psql -d postgres -h localhost -U postgres
// 3. enter password: test
// Setting the same information in PostgresConfig to connect to the local hosted database
func ConnectToPostgres(cfg *PostgresConfig) (ICarbonDB, error) {
	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
			cfg.Host, cfg.User, cfg.Password, cfg.DatabaseName, cfg.Port, cfg.SSLMode,
		)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return carbonDB{db}, nil
}

type ICarbonDB interface {
	Get(id uint) (*Emissions, error)
	Save(emissions *Emissions) error
	List(offset, limit int) ([]*Emissions, error)
	Delete(id uint) error
	Migrate() error
	SearchByResourceName(q string, offset, limit int) ([]*Emissions, error)
}

type carbonDB struct {
	db *gorm.DB
}

func (d carbonDB) Get(id uint) (*Emissions, error) {
	e := &Emissions{}
	err := d.db.Where(`id = ?`, id).Find(e).Error
	return e, err
}

func (d carbonDB) Save(emissions *Emissions) error {
	return d.db.Save(emissions).Error
}

func (d carbonDB) List(offset, limit int) ([]*Emissions, error) {
	var l []*Emissions
	err := d.db.Offset(offset).Limit(limit).Find(&l).Error
	return l, err
}

func (d carbonDB) Delete(id uint) error {
	return d.db.Delete(&Emissions{ID: id}).Error
}

func (d carbonDB) Migrate() error {
	return d.db.AutoMigrate(&Emissions{})
}

func (d carbonDB) SearchByResourceName(q string, offset, limit int) ([]*Emissions, error) {
	var l []*Emissions
	err := d.db.Where(`resource_name like ?`, "%"+q+"%").
		Offset(offset).Limit(limit).Find(&l).Error
	return l, err
}
