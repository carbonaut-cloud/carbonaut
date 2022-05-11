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

	"carbonaut.cloud/carbonaut/pkg/api/models"
	"carbonaut.cloud/carbonaut/pkg/api/v1/config_api"
	"carbonaut.cloud/carbonaut/pkg/api/v1/connector_api"
	"carbonaut.cloud/carbonaut/pkg/api/v1/data_api"
	"github.com/gofiber/fiber/v2"
)

// Config sets configuration file information that gets read by the pkg/config package
// To set values in the configuration file use the path 'api.version'.
type Config struct {
	Version string `default:"v1" validate:"regexp=^v1$"`
	Port    int    `default:"3000"`
}

// all routes in subpackages
var routes = []models.IRoutes{
	connector_api.Routes{},
	config_api.Routes{},
	data_api.Routes{},
}

// Start API server
func Start(c *Config) {
	app := fiber.New()
	v := app.Group(fmt.Sprintf("/api/%s", c.Version))
	addBasePathRoutes(v)
	// Add routes to the API Server
	for _, r := range routes {
		// add recursion to build up API
		group := v.Group(r.GetPrefix())
		r.AddRoutes(group)
	}
	// Add a 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	app.Listen(fmt.Sprintf(":%d", c.Port))
}

func addBasePathRoutes(r fiber.Router) {
	r.Get("/*", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Base API!")
	})
}
