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

	"carbonaut.cloud/carbonaut/pkg/api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
)

func verifyGetResponse(t *testing.T, p int, route, expectedBodyResult string) {
	respBar, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/%s/", p, route))
	assert.NoError(t, err)
	respBarBody, err := ioutil.ReadAll(respBar.Body)
	assert.NoError(t, err)
	assert.Equal(t, expectedBodyResult, string(respBarBody))
}

// Test that the routes are getting initialized right
func TestAddSubRoutes(t *testing.T) {
	port, err := freeport.GetFreePort()
	assert.NoError(t, err, "could not find a free port")
	app := fiber.New()
	addSubRoutes([]models.IRoutes{barRoutes{}, gooRoutes{}, aRoutes{}}, app)
	go func() {
		err := app.Listen(fmt.Sprintf(":%d", port))
		assert.NoError(t, err)
	}()
	time.Sleep(time.Millisecond * 20)

	// To test this a test api is getting build up
	// Structure:
	// 	GET bar/ -> "bar"
	// 	GET bar/foo -> "foo"
	// 	GET goo/ -> "goo"
	// 	GET a/ -> "a"
	// 	GET a/b -> "b"
	// 	GET a/b/c -> "c"
	// 	GET a/b/d -> "d"
	// 	GET a/e -> "e"
	routeStructure := map[string]string{
		"bar":     "bar",
		"bar/foo": "foo",
		"goo":     "goo",
		"a":       "a",
		"a/b":     "b",
		"a/b/c":   "c",
		"a/b/d":   "d",
		"a/e":     "e",
	}

	for route, expectedResponse := range routeStructure {
		verifyGetResponse(t, port, route, expectedResponse)
	}
}

// Build API to test if routes are getting build correct

type barRoutes struct{}

func (c barRoutes) GetPrefix() string {
	return "bar"
}

func (c barRoutes) RouteSubGroups() []models.IRoutes {
	return []models.IRoutes{fooRoutes{}}
}

func (c barRoutes) AddRoutes(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("bar")
	})
}

type fooRoutes struct{}

func (c fooRoutes) GetPrefix() string {
	return "foo"
}

func (c fooRoutes) RouteSubGroups() []models.IRoutes {
	return []models.IRoutes{}
}

func (c fooRoutes) AddRoutes(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("foo")
	})
}

type gooRoutes struct{}

func (c gooRoutes) GetPrefix() string {
	return "goo"
}

func (c gooRoutes) RouteSubGroups() []models.IRoutes {
	return []models.IRoutes{}
}

func (c gooRoutes) AddRoutes(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("goo")
	})
}

type aRoutes struct{}

func (c aRoutes) GetPrefix() string {
	return "a"
}

func (c aRoutes) RouteSubGroups() []models.IRoutes {
	return []models.IRoutes{bRoutes{}, eRoutes{}}
}

func (c aRoutes) AddRoutes(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("a")
	})
}

type bRoutes struct{}

func (c bRoutes) GetPrefix() string {
	return "b"
}

func (c bRoutes) RouteSubGroups() []models.IRoutes {
	return []models.IRoutes{cRoutes{}, dRoutes{}}
}

func (c bRoutes) AddRoutes(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("b")
	})
}

type cRoutes struct{}

func (c cRoutes) GetPrefix() string {
	return "c"
}

func (c cRoutes) RouteSubGroups() []models.IRoutes {
	return []models.IRoutes{}
}

func (c cRoutes) AddRoutes(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("c")
	})
}

type dRoutes struct{}

func (c dRoutes) GetPrefix() string {
	return "d"
}

func (c dRoutes) RouteSubGroups() []models.IRoutes {
	return []models.IRoutes{}
}

func (c dRoutes) AddRoutes(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("d")
	})
}

type eRoutes struct{}

func (c eRoutes) GetPrefix() string {
	return "e"
}

func (c eRoutes) RouteSubGroups() []models.IRoutes {
	return []models.IRoutes{}
}

func (c eRoutes) AddRoutes(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("e")
	})
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

func TestStartGracefullShudownPos(t *testing.T) {
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
