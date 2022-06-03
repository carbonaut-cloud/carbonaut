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
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	negTestFile          = "./testdata/not-a-csv.json"
	doesNotExistTestdata = "does-not-exist"
	testDataFileLength   = 265
	testDataFile         = "./testdata/gcp-carbon-data-extract-1.csv"
)

func TestImportCsvFilePos(t *testing.T) {
	c := Config{}
	r, err := c.ImportCsvFile(testDataFile)
	assert.NoError(t, err)
	assert.Len(t, r, testDataFileLength)
}

func TestImportCsvFileNeg(t *testing.T) {
	c := Config{}
	r, err := c.ImportCsvFile(doesNotExistTestdata)
	assert.Error(t, err)
	assert.Nil(t, r)
	r, err = c.ImportCsvFile(negTestFile)
	assert.Error(t, err)
	assert.Nil(t, r)
}

func TestImportCsvPos(t *testing.T) {
	in, err := os.Open(testDataFile)
	assert.NoError(t, err)
	defer in.Close()
	c := Config{}
	data, err := ioutil.ReadAll(in)
	assert.NoError(t, err)
	r, err := c.ImportCsv(data)
	assert.NoError(t, err)
	assert.Len(t, r, testDataFileLength)
}

func TestImportCsvNeg(t *testing.T) {
	in, err := os.Open(negTestFile)
	assert.NoError(t, err)
	defer in.Close()
	c := Config{}
	data, err := ioutil.ReadAll(in)
	assert.NoError(t, err)
	r, err := c.ImportCsv(data)
	assert.Error(t, err)
	assert.Nil(t, r)
}

func TestPullDataNeg(t *testing.T) {
	c := Config{}
	data, err := c.PullData()
	assert.Error(t, err)
	assert.Nil(t, data)
}
