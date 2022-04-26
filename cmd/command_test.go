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

package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCommand(t *testing.T) {
	err := Execute()
	assert.NoError(t, err)
}

func TestDataExportCommand(t *testing.T) {
	err := RunDataExport()
	assert.Error(t, err)
}

func TestDataImportCommand(t *testing.T) {
	err := RunDataImport()
	assert.Error(t, err)
}

func TestDataReportCommand(t *testing.T) {
	err := RunDataReport()
	assert.Error(t, err)
}

func TestDataCommand(t *testing.T) {
	err := RunData()
	assert.Error(t, err)
}

func TestDeployCommand(t *testing.T) {
	err := RunDeploy()
	assert.Error(t, err)
}

func TestDeployDescribeCommand(t *testing.T) {
	err := RunDeployDescribe()
	assert.Error(t, err)
}