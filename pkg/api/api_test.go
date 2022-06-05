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

package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func verifyGetResponse(t *testing.T, p int, route, expectedBodyResult string) {
	respBar, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/%s/", p, route))
	assert.NoError(t, err)
	respBarBody, err := ioutil.ReadAll(respBar.Body)
	assert.NoError(t, err)
	assert.Equal(t, expectedBodyResult, string(respBarBody))
}

// test if the carbonaut api gets started without any errors

func TestStartPos(t *testing.T) {
	port1, err := freeport.GetFreePort()
	assert.NoError(t, err, "could not find a free port")
	port2, err := freeport.GetFreePort()
	assert.NoError(t, err, "could not find a free port")
	posCfg := []Config{{
		Version: "v1",
		Port:    port1,
	}, {
		Version: "v1",
		Port:    port2,
	}}
	for _, c := range posCfg {
		go func(c Config) {
			api := CarbonautAPI{}
			err := api.Start(&c)
			assert.NoError(t, err)
			err = initHandler(api.app.AcquireCtx(&fasthttp.RequestCtx{}))
			assert.NoError(t, err)
			err = statusHandler(api.app.AcquireCtx(&fasthttp.RequestCtx{}))
			assert.NoError(t, err)
		}(c)
		time.Sleep(time.Millisecond * 50)
		routeStructure := map[string]string{
			fmt.Sprintf("api/%s/status", c.Version): statusOK,
		}

		for route, expectedResponse := range routeStructure {
			verifyGetResponse(t, c.Port, route, expectedResponse)
		}
	}
}

func TestStart404Pos(t *testing.T) {
	port, err := freeport.GetFreePort()
	assert.NoError(t, err, "could not find a free port")
	c := Config{
		Version: "v1",
		Port:    port,
	}
	go func(c Config) {
		api := CarbonautAPI{}
		err := api.Start(&c)
		assert.NoError(t, err)
	}(c)
	time.Sleep(time.Millisecond * 10)

	for route, expectedResponse := range map[string]string{"does-not-exist-route": "Not Found"} {
		verifyGetResponse(t, c.Port, route, expectedResponse)
	}
}

func TestStartPortBlockedNeg(t *testing.T) {
	port, err := freeport.GetFreePort()
	assert.NoError(t, err, "could not find a free port")
	c := Config{
		Version: "v1",
		Port:    port,
	}
	// first one does not throw an error
	go func(c Config) {
		api := CarbonautAPI{}
		err := api.Start(&c)
		assert.NoError(t, err)
	}(c)

	time.Sleep(time.Millisecond * 10)

	// first one does not throw an error
	func(c Config) {
		api := CarbonautAPI{}
		err := api.Start(&c)
		assert.Error(t, err)
	}(c)
}

func TestStartGracefullyShutdownPos(t *testing.T) {
	port, err := freeport.GetFreePort()
	assert.NoError(t, err, "could not find a free port")
	c := Config{
		Version: "v1",
		Port:    port,
	}
	api := CarbonautAPI{}
	// first one does not throw an error
	go func(c Config) {
		err := api.Start(&c)
		assert.NoError(t, err)
	}(c)

	time.Sleep(time.Millisecond * 10)
	err = api.Shutdown()
	assert.NoError(t, err)
}

func TestStartNeg(t *testing.T) {
	port, err := freeport.GetFreePort()
	assert.NoError(t, err, "could not find a free port")
	posCfg := []Config{{}, {
		Version: "v-1",
		Port:    port,
	}, {
		Version: "v1",
		Port:    -1,
	}}
	for _, c := range posCfg {
		go func(c Config) {
			api := CarbonautAPI{}
			err := api.Start(&c)
			assert.Error(t, err)
		}(c)
	}
}
