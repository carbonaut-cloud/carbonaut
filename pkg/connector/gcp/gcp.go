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

package gcp

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type Record struct {
	UsageMonth            string `csv:"usage_month"`
	BillingAccount        string `csv:"billing_account_id"`
	ProjectNumber         string `csv:"project.number"`
	ProjectID             string `csv:"project.id"`
	ServiceID             string `csv:"service.id"`
	ServiceDescription    string `csv:"service.description"`
	LocationLocation      string `csv:"location.location"`
	LocationRegion        string `csv:"location.region"`
	CarbonFootprintKgCO2e string `csv:"carbon_footprint_kgCO2e"`
	CarbonModelVersion    string `csv:"carbon_model_version"`
}

type Config struct{}

func (c Config) ImportCsvFile(filepath string) ([]*Record, error) {
	in, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer in.Close()

	clients := []*Record{}

	if err := gocsv.UnmarshalFile(in, &clients); err != nil {
		return nil, err
	}
	return clients, nil
}

func (c Config) ImportCsv(data []byte) ([]*Record, error) {
	clients := []*Record{}
	if err := gocsv.UnmarshalBytes(data, &clients); err != nil {
		return nil, err
	}
	return clients, nil
}

func (c Config) PullData() ([]byte, error) {
	return nil, fmt.Errorf("not implemented yet")
}
