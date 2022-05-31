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

package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLoggerPos(t *testing.T) {
	configs := []LogConfig{{}, {LoggerOutputFormatting: FormatJSON, LogLevel: "info"}}
	for i := range configs {
		logger := GetLogger(&configs[i])
		assert.NotNil(t, logger)
		logger, err := getLogger(&configs[i])
		assert.NoError(t, err)
		assert.NotNil(t, logger)
		logger.Info().Msg("Does not error")
	}
}

func TestGetLoggerNeg(t *testing.T) {
	configs := []LogConfig{
		{LoggerOutputFormatting: "does not exist"},
		{LoggerOutputFormatting: FormatJSON, LogLevel: "indigo"},
	}
	for i := range configs {
		logger := GetLogger(&configs[i])
		assert.Nil(t, logger)
		logger, err := getLogger(&configs[i])
		assert.Error(t, err)
		assert.Nil(t, logger)
	}
}
