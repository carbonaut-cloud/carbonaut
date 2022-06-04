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

package connector

import (
	"fmt"

	"carbonaut.cloud/carbonaut/pkg/connector/cloud/gcp"
	"carbonaut.cloud/carbonaut/pkg/connector/provider"
	"carbonaut.cloud/carbonaut/pkg/data/methods"
)

type Config struct {
	Providers []provider.Config
}

//
// Connect a cloud provider to carbonaut to pull data
//

// This function establishes a connection to the specified provider to pull data until the Connect func stops
func Connect(c Config) error {
	return fmt.Errorf("not implemented yet")
}

//
// Import provider data into carbonaut
//

type (
	ImportIn struct {
		ProviderName provider.Name
		ImportType   importType
		FilePath     string
		DB           methods.ICarbonDB
	}
	importType string
)

const FileImport = "file"

var SupportedImportTypes = []importType{FileImport}

func ImportData(in ImportIn) error {
	p, err := getProvider(in.ProviderName)
	if err != nil {
		return err
	}
	switch in.ImportType {
	case FileImport:
		e, err := p.ImportCsvFile(in.FilePath)
		if err != nil {
			return err
		}
		if err := in.DB.BatchSave(e); err != nil {
			return err
		}
	default:
		return fmt.Errorf("could not resolve specified import type, supported import types: %v", SupportedImportTypes)
	}
	return nil
}

func getProvider(providerName provider.Name) (provider.CloudProvider, error) {
	switch providerName {
	case gcp.Provider.Name:
		return gcp.Provider.Methods, nil
	default:
		return nil, fmt.Errorf("could not resolve specified provider")
	}
}
