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

var (
	NegFolderPath = fmt.Sprintf("%s/%s", TestFolder, NegFolder)
	NegConfigs    = []GetCarbonautConfigIn{
		{ConfigMedium: configProvider(GenericDoesNotExistString), FilePath: ""},
		{ConfigMedium: FileConfigMedium, FilePath: fmt.Sprintf("%s/%s.%s", GenericDoesNotExistString, GenericDoesNotExistString, YAML)},
		{ConfigMedium: FileConfigMedium, FilePath: fmt.Sprintf("%s/%s.%s", NegFolderPath, "minimal", GenericDoesNotExistString)},
		{ConfigMedium: FileConfigMedium, FilePath: fmt.Sprintf("%s/%s", NegFolderPath, GenericDoesNotExistString)},
	}
	PosFolderPath = fmt.Sprintf("%s/%s", TestFolder, PosFolder)
	PosConfigs    = []GetCarbonautConfigIn{
		{ConfigMedium: DefaultConfigMedium},
	}
)

func TestGetCarbonautConfigNeg(t *testing.T) {
	for i := range NegConfigs {
		cfg, err := GetCarbonautConfig(&NegConfigs[i])
		assert.Error(t, err)
		assert.Nil(t, cfg)
	}
}

func TestGetCarbonautConfigPos(t *testing.T) {
	for i := range PosConfigs {
		cfg, err := GetCarbonautConfig(&PosConfigs[i])
		assert.NoError(t, err)
		assert.NotNil(t, cfg)
	}
}

func TestGetCarbonautConfigNegFile(t *testing.T) {
	folderPath := fmt.Sprintf("%s/%s", TestFolder, NegFolder)
	files, err := ioutil.ReadDir(folderPath)
	assert.NoError(t, err)
	for _, f := range files {
		cfg, err := GetCarbonautConfig(&GetCarbonautConfigIn{
			ConfigMedium: FileConfigMedium,
			FilePath:     fmt.Sprintf("%s/%s.%s", folderPath, f.Name(), YAML),
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
			FilePath:     fmt.Sprintf("%s/%s.%s", folderPath, f.Name(), YAML),
		})
		assert.NoError(t, err, fmt.Sprintf("no error expected for for test file %s/%s", folderPath, f.Name()))
		assert.NotNil(t, cfg)
	}
}

func TestSplitFilePathToConfigPos(t *testing.T) {
	testMap := map[string]fileConfig{
		"./my": {
			FileName: "",
			FileType: "/my",
			FilePath: "",
		},
		"./my.db": {
			FileName: "my",
			FileType: "db",
			FilePath: ".",
		},
		"./my/my.db": {
			FileName: "my",
			FileType: "db",
			FilePath: "./my",
		},
		"./my/..my.myTale324": {
			FileName: "..my",
			FileType: "myTale324",
			FilePath: "./my",
		},
	}
	for inputString, expectedOutStruct := range testMap {
		cfg, err := splitFilePathToPieces(inputString)
		assert.NoError(t, err)
		assert.Equal(t, cfg.FileName, expectedOutStruct.FileName)
		assert.Equal(t, cfg.FilePath, expectedOutStruct.FilePath)
		assert.Equal(t, cfg.FileType, expectedOutStruct.FileType)
	}
}

func TestSplitFilePathToConfigNeg(t *testing.T) {
	testMap := []string{"my", "my/dw"}
	for _, inputString := range testMap {
		cfg, err := splitFilePathToPieces(inputString)
		assert.Error(t, err)
		assert.Nil(t, cfg)
	}
}
