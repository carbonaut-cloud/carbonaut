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

package main

import (
	"carbonaut.cloud/carbonaut/pkg/api"
	"carbonaut.cloud/carbonaut/pkg/util"
	"github.com/rs/zerolog/log"
)

func main() {
	l := util.GetLogger(&util.LogConfig{})
	l.Info().Msg("Applying defaults to the configuration to look up the carbonaut config file")
	a := api.CarbonautAPI{}
	if err := a.Start(&api.Config{}); err != nil {
		log.Fatal().Err(err)
	}
}
