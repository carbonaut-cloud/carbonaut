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
	"fmt"
	"os"

	"github.com/caarlos0/env"
	"github.com/go-playground/validator/v10"
	"github.com/mcuadros/go-defaults"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Log = GetLogger(&LogConfig{})

type LogConfig struct {
	LoggerOutputFormatting loggerFormats `default:"console"`
	LogLevel               string        `env:"LOG_LEVEL" default:"info" validate:"oneof='info' 'debug' 'error' 'fatal' 'warning'"`
}

type loggerFormats string

const (
	FormatJSON    = "json"
	FormatConsole = "console"
)

func GetLogger(in *LogConfig) *zerolog.Logger {
	logger, err := getLogger(in)
	if err != nil {
		log.Err(err)
	}
	return logger
}

func getLogger(in *LogConfig) (*zerolog.Logger, error) {
	if err := env.Parse(in); err != nil {
		return nil, fmt.Errorf("could not get environment variables %v", err)
	}
	defaults.SetDefaults(in)
	if err := validator.New().Struct(in); err != nil {
		return nil, fmt.Errorf("invalid configuration, %v", err)
	}
	logLevel, err := zerolog.ParseLevel(in.LogLevel)
	if err != nil {
		return nil, err
	}
	switch in.LoggerOutputFormatting {
	case FormatJSON:
		logger := log.Logger.With().Logger().Level(logLevel)
		return &logger, nil
	case FormatConsole:
		logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Logger().Level(logLevel)
		return &logger, nil
	default:
		return nil, fmt.Errorf("specified logger formatting type %s is not supported", in.LoggerOutputFormatting)
	}
}
