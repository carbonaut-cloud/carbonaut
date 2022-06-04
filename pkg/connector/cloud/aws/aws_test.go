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

package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImportCsvFileNeg(t *testing.T) {
	c := implProvider{}
	data, err := c.ImportCsvFile("")
	assert.Error(t, err)
	assert.Nil(t, data)
}

func TestImportCsvNeg(t *testing.T) {
	c := implProvider{}
	data, err := c.ImportCsv([]byte{})
	assert.Error(t, err)
	assert.Nil(t, data)
}

func TestExportAllToCsvNeg(t *testing.T) {
	c := implProvider{}
	data, err := c.ExportAllToCsv()
	assert.Error(t, err)
	assert.Nil(t, data)
}

func TestPullDataNeg(t *testing.T) {
	c := implProvider{}
	data, err := c.PullData()
	assert.Error(t, err)
	assert.Nil(t, data)
}

func TestConnectNeg(t *testing.T) {
	c := implProvider{}
	err := c.Connect()
	assert.Error(t, err)
}

func TestStatusNeg(t *testing.T) {
	c := implProvider{}
	s, err := c.Status()
	assert.Error(t, err)
	assert.Empty(t, s)
}
