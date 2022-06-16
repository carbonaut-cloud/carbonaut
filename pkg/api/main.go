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

// The API uses fiber https://docs.gofiber.io/
// To document the api properly the library https://github.com/swaggo/swag/ is used

import (
	"fmt"

	// This is needed to initialize the swagger API
	"carbonaut.cloud/carbonaut/docs"
	"carbonaut.cloud/carbonaut/pkg/api/methods"
	"carbonaut.cloud/carbonaut/pkg/api/routes"
	"carbonaut.cloud/carbonaut/pkg/api/routes/config"
	"carbonaut.cloud/carbonaut/pkg/api/routes/connector"
	"carbonaut.cloud/carbonaut/pkg/api/routes/data"
	"carbonaut.cloud/carbonaut/pkg/util"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/caarlos0/env"
	"github.com/gofiber/fiber/v2"
	"github.com/mcuadros/go-defaults"
)

// Config sets configuration file information that gets read by the pkg/config package
// To set values in the configuration file use the path 'api.version'.
type Config struct {
	Port int `default:"3000" env:"API_PORT"`
}

// all r in sub packages
var r = []routes.IRoutes{
	connector.Routes{},
	config.Routes{},
	data.Routes{},
}

type CarbonautAPI struct {
	app *fiber.App
}

const Version = "v0.0.1"

// @title Carbonaut API
// @version 0.0.1
// @description This API is used to interact with Carbonaut resources
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func (a *CarbonautAPI) Start(c *Config) error {
	log := util.Log
	log.Info().Msg("Start Carbonaut API")
	if err := env.Parse(c); err != nil {
		return fmt.Errorf("could not get environment variables %v", err)
	}
	defaults.SetDefaults(c)
	app := fiber.New()
	a.app = app
	v := app.Group(fmt.Sprintf("/api/%s", Version))

	// Add swagger information
	docs.SwaggerInfo.Version = Version

	addBasePathRoutes(v)
	// Add routes to the API Server
	methods.AddSubRoutes(r, v)
	// Add a 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	log.Info().Msgf("Swagger endpoint: http://127.0.0.1:%d/api/%s/swagger/index.html", c.Port, Version)
	if err := app.Listen(fmt.Sprintf(":%d", c.Port)); err != nil {
		return err
	}
	return nil
}

// Shutdown the API
func (a CarbonautAPI) Shutdown() error {
	return a.app.Shutdown()
}

func addBasePathRoutes(r fiber.Router) {
	r.Get("/status", statusHandler)
	r.Post("/init", initHandler)
	// host a add swagger web UI
	r.Get("/swagger/*", swagger.HandlerDefault)
}

const statusOK = "Carbonaut API is running!"

// @description Carbonaut Status Endpoint
// @Success 200 {string} Carbonaut API is running!
// @Router /api/v1/status/ [get]
func statusHandler(c *fiber.Ctx) error {
	return c.SendString(statusOK)
}

// @description Initialize carbonaut to be fully functional
// @Success 200 {string} Carbonaut initialized!
// @Router /api/v1/init [post]
func initHandler(c *fiber.Ctx) error {
	return c.SendString("wip, not implemented")
}
