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

package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var ePos = Emissions{
	ID:              1,
	ResourceName:    "ec2",
	ResourceOwner:   "user1",
	ResourceType:    "compute",
	Provider:        "aws",
	MTCO2e:          0.001,
	Location:        "eu-west-1",
	ProviderVersion: 1,
	Timestamp:       time.Now(),
	CreatedAt:       time.Now(),
	UpdatedAt:       time.Now(),
}

func TestEmissionMarshalingPos(t *testing.T) {
	d, err := ePos.Marshal()
	assert.NoError(t, err)
	assert.NotNil(t, d)
	e, err := UnmarshalEmissions(d)
	assert.NoError(t, err)
	assert.Equal(t, ePos.Location, e.Location)
	assert.Equal(t, ePos.MTCO2e, e.MTCO2e)
	assert.Equal(t, ePos.Provider, e.Provider)
}

func TestGetEmissionsTestDataPos(t *testing.T) {
	e, err := GetEmissionsTestData()
	assert.NoError(t, err)
	assert.NotEmpty(t, e.ID)
}

func TestGetEmissionsTestDataSetsPos(t *testing.T) {
	posList := []int{0, 10, 100}
	for i := range posList {
		e, err := GetEmissionsTestDataSets(posList[i])
		assert.NoError(t, err)
		assert.Equal(t, posList[i], len(e))
	}
}

func TestGetEmissionsTestDataSetsNeg(t *testing.T) {
	posList := -1
	e, err := GetEmissionsTestDataSets(posList)
	assert.NoError(t, err)
	assert.Empty(t, e)
}
