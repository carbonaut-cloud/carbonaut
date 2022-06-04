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

	"github.com/bxcodec/faker/v3"
)

// Emissions define data models
type Emissions struct {
	ID              uint64    `gorm:"primarykey;autoIncrement:true" json:"id"`
	ResourceName    string    `gorm:"column:resource_name" json:"resource_name" faker:"oneof: ec2, eks, s3, ses, sns, sqs, rds, ecs, vpc"`
	ResourceOwner   string    `gorm:"column:resource_owner" json:"resource_owner" faker:"name"`
	ResourceType    string    `gorm:"column:resource_type" json:"resource_type" faker:"oneof: compute, network, ml, serverless, db, api, security"`
	Provider        string    `gorm:"column:provider" json:"provider" faker:"oneof: gcp, aws, azure, kubernetes"`
	MTCO2e          float64   `gorm:"column:mtco2e" json:"mtco2e" faker:"boundary_start=0.00001, boundary_end=0.1"`
	Location        string    `gorm:"column:location" json:"location" faker:"oneof: us, eu, asia, africa"`
	ProviderVersion int       `gorm:"column:provider_version" json:"provider_version"`
	Timestamp       time.Time `gorm:"column:timestamp" json:"timestamp"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (e *Emissions) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

func UnmarshalEmissions(data []byte) (Emissions, error) {
	var e Emissions
	err := json.Unmarshal(data, &e)
	return e, err
}

func GetEmissionsTestDataSets(n int) ([]*Emissions, error) {
	e := []*Emissions{}
	for i := 0; i < n; i++ {
		emission, err := GetEmissionsTestData()
		if err != nil {
			return nil, err
		}
		e = append(e, emission)
	}
	return e, nil
}

func GetEmissionsTestData() (*Emissions, error) {
	e := &Emissions{}
	err := faker.FakeData(e)
	return e, err
}
