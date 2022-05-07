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
	"carbonaut.cloud/carbonaut/pkg/data/db/models"
	"gorm.io/gorm"
)

type ICarbonDB interface {
	Get(id uint) (*models.Emissions, error)
	Save(emissions *models.Emissions) error
	List(offset, limit int) ([]*models.Emissions, error)
	Delete(id uint) error
	Migrate() error
	SearchByResourceName(q string, offset, limit int) ([]*models.Emissions, error)
}

type CarbonDB struct {
	db *gorm.DB
}

// initialize database struct
func (d *CarbonDB) Init(db *gorm.DB) {
	d.db = db
}

func (d CarbonDB) Get(id uint) (*models.Emissions, error) {
	e := &models.Emissions{}
	err := d.db.Where(`id = ?`, id).Find(e).Error
	return e, err
}

func (d CarbonDB) Save(emissions *models.Emissions) error {
	return d.db.Save(emissions).Error
}

func (d CarbonDB) List(offset, limit int) ([]*models.Emissions, error) {
	var l []*models.Emissions
	err := d.db.Offset(offset).Limit(limit).Find(&l).Error
	return l, err
}

func (d CarbonDB) Delete(id uint) error {
	return d.db.Delete(&models.Emissions{ID: id}).Error
}

func (d CarbonDB) Migrate() error {
	return d.db.AutoMigrate(&models.Emissions{})
}

func (d CarbonDB) SearchByResourceName(q string, offset, limit int) ([]*models.Emissions, error) {
	var l []*models.Emissions
	err := d.db.Where(`resource_name like ?`, "%"+q+"%").
		Offset(offset).Limit(limit).Find(&l).Error
	return l, err
}
