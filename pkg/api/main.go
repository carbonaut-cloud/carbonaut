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

	// This is needed to initialize the swagger API
	"carbonaut.cloud/carbonaut/docs"
	"carbonaut.cloud/carbonaut/pkg/api/models"
	"carbonaut.cloud/carbonaut/pkg/api/v1/config"
	"carbonaut.cloud/carbonaut/pkg/api/v1/connector"
	"carbonaut.cloud/carbonaut/pkg/api/v1/data"
	"carbonaut.cloud/carbonaut/pkg/util"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// Config sets configuration file information that gets read by the pkg/config package
// To set values in the configuration file use the path 'api.version'.
type Config struct {
	Version string `default:"v1" validate:"regexp=^v1$"`
	Port    int    `default:"3000"`
}

// all routes in sub packages
var routes = []models.IRoutes{
	connector.Routes{},
	config.Routes{},
	data.Routes{},
}

type CarbonautAPI struct {
	app *fiber.App
}

// @title Carbonaut API
// @version 0.0.1
// @description This API is used to interact with Carbonaut resources
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func (a *CarbonautAPI) Start(c *Config) error {
	log := util.Log
	log.Info().Msg("Start Carbonaut API")
	app := fiber.New()
	a.app = app
	v := app.Group(fmt.Sprintf("/api/%s", c.Version))

	// Add swagger information
	docs.SwaggerInfo.Version = c.Version

	addBasePathRoutes(v)
	// Add routes to the API Server
	addSubRoutes(routes, v)
	// Add a 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	log.Info().Msgf("Swagger endpoint: http://127.0.0.1:%d/api/%s/swagger/index.html", c.Port, c.Version)
	if err := app.Listen(fmt.Sprintf(":%d", c.Port)); err != nil {
		return err
	}
	return nil
}

// Shutdown the API
func (a CarbonautAPI) Shutdown() error {
	return a.app.Shutdown()
}

func addSubRoutes(routes []models.IRoutes, routeGroup fiber.Router) {
	// if there aren't any routes left break
	if len(routes) == 0 {
		return
	}
	subGroup := routeGroup.Group(routes[0].GetPrefix())
	routes[0].AddRoutes(subGroup)
	// recursion to add routes of subgroups (e.g. api/v1/data/db)
	addSubRoutes(routes[0].RouteSubGroups(), subGroup)
	// remove first element from the list
	addSubRoutes(append(routes[:0], routes[1:]...), routeGroup)
}

func addBasePathRoutes(r fiber.Router) {
	r.Get("/status", statusHandler)
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
