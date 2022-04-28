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

package config

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TestFolder                = "test"
	PosFolder                 = "pos"
	NegFolder                 = "neg"
	YAML                      = "yaml"
	GenericDoesNotExistString = "does-not-exist"
)

func TestGetCarbonautConfigNegInvalidConfigMedium(t *testing.T) {
	cfg, err := GetCarbonautConfig(&GetCarbonautConfigIn{
		ConfigMedium:     configMedium(GenericDoesNotExistString),
		ConfigMediumFile: &FileMediumConfig{},
	})
	assert.Error(t, err)
	assert.Nil(t, cfg)
}

func TestGetCarbonautConfigNegNoFile(t *testing.T) {
	folderPathDoesNotExist := fmt.Sprintf("%s/%s", TestFolder, GenericDoesNotExistString)
	cfg, err := GetCarbonautConfig(&GetCarbonautConfigIn{
		ConfigMedium: FileConfigMedium,
		ConfigMediumFile: &FileMediumConfig{
			FileName: GenericDoesNotExistString,
			FileType: YAML,
			FilePath: folderPathDoesNotExist,
		},
	})
	assert.Error(t, err)
	assert.Nil(t, cfg)
}

func TestGetCarbonautConfigNegWrongFileType(t *testing.T) {
	posFolderPath := fmt.Sprintf("%s/%s", TestFolder, PosFolder)
	cfg, err := GetCarbonautConfig(&GetCarbonautConfigIn{
		ConfigMedium: FileConfigMedium,
		ConfigMediumFile: &FileMediumConfig{
			FileName: "minimal",
			FileType: GenericDoesNotExistString,
			FilePath: posFolderPath,
		},
	})
	assert.Error(t, err)
	assert.Nil(t, cfg)
}

func TestGetCarbonautConfigNegFile(t *testing.T) {
	folderPath := fmt.Sprintf("%s/%s", TestFolder, NegFolder)
	files, err := ioutil.ReadDir(folderPath)
	assert.NoError(t, err)
	for _, f := range files {
		cfg, err := GetCarbonautConfig(&GetCarbonautConfigIn{
			ConfigMedium: FileConfigMedium,
			ConfigMediumFile: &FileMediumConfig{
				FileName: f.Name(),
				FileType: YAML,
				FilePath: folderPath,
			},
		})
		assert.Error(t, err, fmt.Sprintf("expect an error for test file %s/%s", folderPath, f.Name()))
		assert.Nil(t, cfg)
	}
}

func TestGetCarbonautConfigPosFile(t *testing.T) {
	folderPath := fmt.Sprintf("%s/%s", TestFolder, PosFolder)
	files, err := ioutil.ReadDir(folderPath)
	assert.NoError(t, err)
	for _, f := range files {
		cfg, err := GetCarbonautConfig(&GetCarbonautConfigIn{
			ConfigMedium: FileConfigMedium,
			ConfigMediumFile: &FileMediumConfig{
				FileName: f.Name(),
				FileType: YAML,
				FilePath: folderPath,
			},
		})
		assert.NoError(t, err, fmt.Sprintf("no error expected for for test file %s/%s", folderPath, f.Name()))
		assert.NotNil(t, cfg)
	}
}
