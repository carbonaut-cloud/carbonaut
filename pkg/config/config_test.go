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
	"os"
	"testing"

	"github.com/mcuadros/go-defaults"
	"github.com/stretchr/testify/assert"
)

const (
	folderPath  = "../../test/data/config"
	NotExistent = "does-not-exist"
	ConfigPath  = "CONFIG_FILE_PATH"
)

var (
	NegConfigs = []map[string]string{
		{ConfigPath: fmt.Sprintf("%s/neg/%s", folderPath, "neg-no-yaml-3.yml")},
		{ConfigPath: fmt.Sprintf("%s/neg/%s", folderPath, "pos-empty.yml")},
	}
	PosConfigs = []map[string]string{
		{ConfigPath: fmt.Sprintf("%s/pos/%s", folderPath, "pos-empty.yml")},
		{ConfigPath: fmt.Sprintf("%s/pos/%s", folderPath, "pos-minimal.yml")},
	}
)

// check if test files are available
func TestCheckTestData(t *testing.T) {
	f, err := os.Stat(folderPath)
	assert.NoError(t, err)
	assert.Equal(t, "config", f.Name())
}

func TestReadFileNeg(t *testing.T) {
	for i := range NegConfigs {
		t.Log(NegConfigs[i][ConfigPath])
		// configure environment variables
		configPathEnvBefore := os.Getenv(ConfigPath)
		assert.NoError(t, os.Setenv(ConfigPath, NegConfigs[i][ConfigPath]))
		// run test
		cfg, err := ReadFile()
		assert.Error(t, err)
		assert.Nil(t, cfg)
		// clean up
		assert.NoError(t, os.Setenv(ConfigPath, configPathEnvBefore))
	}
}

func TestReadFilePos(t *testing.T) {
	for i := range PosConfigs {
		t.Log(PosConfigs[i][ConfigPath])
		// configure environment variables
		configPathEnvBefore := os.Getenv(ConfigPath)
		assert.NoError(t, os.Setenv(ConfigPath, PosConfigs[i][ConfigPath]))
		// run test
		cfg, err := ReadFile()
		assert.NoError(t, err)
		assert.NotNil(t, cfg)
		// clean up
		assert.NoError(t, os.Setenv(ConfigPath, configPathEnvBefore))
	}
}

func TestWriteFilePos(t *testing.T) {
	c := CarbonConfig{}
	defaults.SetDefaults(&c)
	err := WriteFile(&WriteFileIn{
		ConfigFilePath: "./tmp.yml",
		Config:         c,
		AutoCreate:     true,
	})
	assert.NoError(t, err)
	assert.NoError(t, os.Remove("./tmp.yml"), "removing/ cleaning up the generated test file should not error^")
}

func TestWriteFileNeg(t *testing.T) {
	c := CarbonConfig{}
	defaults.SetDefaults(&c)
	configs := []WriteFileIn{{
		ConfigFilePath: "./tmp.yml",
		Config:         c,
		AutoCreate:     false,
	}, {
		ConfigFilePath: "/usr/tmp.yml",
		Config:         c,
		AutoCreate:     true,
	}}
	for i := range configs {
		err := WriteFile(&configs[i])
		assert.Error(t, err)
	}
}
