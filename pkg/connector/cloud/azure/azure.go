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

package azure

import (
	"fmt"

	"carbonaut.cloud/carbonaut/pkg/connector/provider"
	"carbonaut.cloud/carbonaut/pkg/data/models"
)

type Config struct{}

var Provider = provider.ImplementedProvider{
	Name:    "azure",
	Methods: implProvider{},
}

type Record struct{}

type implProvider struct{}

func (p implProvider) ImportCsvFile(filepath string) ([]*models.Emissions, error) {
	return nil, fmt.Errorf("not implemented yet")
}

func (p implProvider) ImportCsv(data []byte) ([]*models.Emissions, error) {
	return nil, fmt.Errorf("not implemented yet")
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
