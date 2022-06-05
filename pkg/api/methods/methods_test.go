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

package methods

import (
	"testing"

	"carbonaut.cloud/carbonaut/pkg/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Test that the routes are getting initialized right
func TestAddSubRoutes(t *testing.T) {
	port := RouteTesting(t, []routes.IRoutes{barRoutes{}, gooRoutes{}, aRoutes{}})
	assert.NotEmpty(t, port)
	for route, expectedResponse := range map[string]string{
		"bar":     "bar",
		"bar/foo": "foo",
		"goo":     "goo",
		"a":       "a",
		"a/b":     "b",
		"a/b/c":   "c",
		"a/b/d":   "d",
		"a/e":     "e",
	} {
		VerifyGetResponse(t, port, route, expectedResponse)
	}
}

// Build API to test if routes are getting build correct

type barRoutes struct{}

func (c barRoutes) GetPrefix() string {
	return "bar"
}

func (c barRoutes) RouteSubGroups() []routes.IRoutes {
	return []routes.IRoutes{fooRoutes{}}
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

func (c fooRoutes) RouteSubGroups() []routes.IRoutes {
	return []routes.IRoutes{}
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

func (c gooRoutes) RouteSubGroups() []routes.IRoutes {
	return []routes.IRoutes{}
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

func (c aRoutes) RouteSubGroups() []routes.IRoutes {
	return []routes.IRoutes{bRoutes{}, eRoutes{}}
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

func (c bRoutes) RouteSubGroups() []routes.IRoutes {
	return []routes.IRoutes{cRoutes{}, dRoutes{}}
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

func (c cRoutes) RouteSubGroups() []routes.IRoutes {
	return []routes.IRoutes{}
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

func (c dRoutes) RouteSubGroups() []routes.IRoutes {
	return []routes.IRoutes{}
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

func (c eRoutes) RouteSubGroups() []routes.IRoutes {
	return []routes.IRoutes{}
}

func (c eRoutes) AddRoutes(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("e")
	})
}
