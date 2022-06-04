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

	"carbonaut.cloud/carbonaut/pkg/connector/provider"
	"carbonaut.cloud/carbonaut/pkg/data/models"
	"github.com/bxcodec/faker/v3"
	"github.com/gocarina/gocsv"
)

type Config struct{}

var Provider = provider.ImplementedProvider{
	Name:    "gcp",
	Methods: implProvider{},
}

type Record struct {
	UsageMonth            string `csv:"usage_month" faker:"oneof: 01/01/2001, 02/02/2002, 03/03/2003"`
	BillingAccount        string `csv:"billing_account_id"`
	ProjectNumber         string `csv:"project.number"`
	ProjectID             string `csv:"project.id"`
	ServiceID             string `csv:"service.id"`
	ServiceDescription    string `csv:"service.description"`
	LocationLocation      string `csv:"location.location" faker:"oneof: us, eu, asia, africa"`
	LocationRegion        string `csv:"location.region" faker:"oneof: us, eu, asia, africa"`
	CarbonFootprintKgCO2e string `csv:"carbon_footprint_kgCO2e" faker:"oneof: 0.003, 0.004, 0.0001, 0.0003"`
	CarbonModelVersion    string `csv:"carbon_model_version" faker:"oneof: 1, 2, 3"`
}

type implProvider struct{}

func (p implProvider) ImportCsvFile(filepath string) ([]*models.Emissions, error) {
	r, err := readToRecords(filepath)
	if err != nil {
		return nil, err
	}
	e, err := translateRecords(r)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (p implProvider) ImportCsv(data []byte) ([]*models.Emissions, error) {
	clients := []*Record{}
	if err := gocsv.UnmarshalBytes(data, &clients); err != nil {
		return nil, err
	}
	e, err := translateRecords(clients)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (p implProvider) ExportAllToCsv() ([]byte, error) {
	return nil, fmt.Errorf("not implemented yet")
}

func (p implProvider) PullData() ([]byte, error) {
	return nil, fmt.Errorf("not implemented yet")
}

func (p implProvider) Connect() error {
	return fmt.Errorf("not implemented yet")
}

func (p implProvider) Status() (string, error) {
	return "", fmt.Errorf("not implemented yet")
}

func readToRecords(filepath string) ([]*Record, error) {
	in, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer in.Close()
	r := []*Record{}
	if err := gocsv.UnmarshalFile(in, &r); err != nil {
		return nil, err
	}
	return r, nil
}

func GetRecordTestData() (*Record, error) {
	e := &Record{}
	err := faker.FakeData(e)
	return e, err
}
