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

package models

import (
	"encoding/json"
	"time"
)

// Emissions define data models
type Emissions struct {
	ID              uint64    `gorm:"primarykey;autoIncrement:true" json:"id"`
	ResourceName    string    `gorm:"column:resource_name" json:"resource_name"`
	ResourceOwner   string    `gorm:"column:resource_owner" json:"resource_owner"`
	ResourceType    string    `gorm:"column:resource_type" json:"resource_type"`
	Provider        string    `gorm:"column:provider" json:"provider"`
	MTCO2e          float64   `gorm:"column:mtco2e" json:"mtco2e"`
	Location        string    `gorm:"column:location" json:"location"`
	ProviderVersion int       `gorm:"column:provider_version" json:"provider_version"`
	Timestamp       time.Time `gorm:"column:timestamp" json:"timestamp"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (e *Emissions) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

func UnmarshalEmissions(data []byte) (Emissions, error) {
	var r Emissions
	err := json.Unmarshal(data, &r)
	return r, err
}
