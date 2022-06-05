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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TODO: translate gcp data to unified carbonaut data schema (implement spec connector/translator.go)

// func translateRecords(r []*Record) ([]*models.Emissions, error)
// func translateRecord(r *Record) (*models.Emissions, error)
// func dateFormat(d string) time.Time

var negRecords = []*Record{{}, {CarbonModelVersion: doesNotExistTestdata}, {
	CarbonFootprintKgCO2e: doesNotExistTestdata,
	CarbonModelVersion:    "1",
}}

func TestTranslateRecordPos(t *testing.T) {
	r, err := GetRecordTestData()
	assert.NoError(t, err)
	e, err := translateRecord(r)
	assert.NoError(t, err)
	assert.NotNil(t, e)
	// cross check one of the fields if the decoding succeeded
	assert.Equal(t, dateFormat(r.UsageMonth), e.Timestamp)
}

func TestTranslateRecordNeg(t *testing.T) {
	for i := range negRecords {
		e, err := translateRecord(negRecords[i])
		assert.Error(t, err)
		assert.Nil(t, e)
	}
}

func TestTranslateRecordsPos(t *testing.T) {
	r, err := readToRecords(testDataFile)
	assert.NoError(t, err)
	e, err := translateRecords(r)
	assert.NoError(t, err)
	assert.NotNil(t, e)
	assert.Equal(t, len(r), len(e))
}

func TestTranslateRecordsNeg(t *testing.T) {
	e, err := translateRecords(negRecords)
	assert.Error(t, err)
	assert.Empty(t, e)
}

type dateTranslationData struct {
	s string
	t time.Time
}

func TestDateFormatPos(t *testing.T) {
	data := []dateTranslationData{{
		s: "01/04/2022",
		t: time.Date(2022, 4, 1, 0, 0, 0, 0, &time.Location{}),
	}, {
		s: "01/01/2022",
		t: time.Date(2022, 1, 1, 0, 0, 0, 0, &time.Location{}),
	}, {
		s: "-11/04/2022",
		t: time.Date(2022, 4, 11, 0, 0, 0, 0, &time.Location{}),
	}, {
		s: "ABC11/04/2022/",
		t: time.Date(2022, 4, 11, 0, 0, 0, 0, &time.Location{}),
	}}
	for _, d := range data {
		decodedTime := dateFormat(d.s)
		assert.NotNil(t, decodedTime)
		assert.Equal(t, d.t, decodedTime)
	}
}

func TestDateFormatNeg(t *testing.T) {
	data := []dateTranslationData{
		{s: "", t: emptyDate},
		{s: "01/012/022", t: emptyDate},
		{s: "01/01/ABCD", t: emptyDate},
		{s: "00/01/2022", t: emptyDate},
		{s: "01/00/2022", t: emptyDate},
		{s: "01/01/0000", t: emptyDate},
	}
	for _, d := range data {
		decodedTime := dateFormat(d.s)
		assert.NotNil(t, decodedTime)
		assert.Equal(t, d.t, decodedTime)
	}
}
