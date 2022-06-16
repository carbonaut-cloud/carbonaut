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

package connector

import (
	"carbonaut.cloud/carbonaut/pkg/api/routes"
	"github.com/gofiber/fiber/v2"
)

type Routes struct{}

func (c Routes) GetPrefix() string {
	return "connector"
}

func (c Routes) RouteSubGroups() []routes.IRoutes {
	return []routes.IRoutes{ConnectRoutes{}}
}

func (c Routes) AddRoutes(r fiber.Router) {
	r.Get("/connections", connectionsHandler)
}

// @description WIP, list carbonaut data provider connections
// @Success 200 {string} config
// @Tags connector
// @Router /api/v1/connector/connections [get]
func connectionsHandler(c *fiber.Ctx) error {
	// TODO: not implemented
	return c.SendString("wip, not implemented")
}

//
// Connect
//

type ConnectRoutes struct{}

func (c ConnectRoutes) GetPrefix() string {
	return "connect"
}

func (c ConnectRoutes) RouteSubGroups() []routes.IRoutes {
	return []routes.IRoutes{}
}

func (c ConnectRoutes) AddRoutes(r fiber.Router) {
	r.Post("/aws", connectAwsHandler)
	r.Post("/azure", connectAzureHandler)
	r.Post("/gcp", connectGcpHandler)
}

// @description WIP, connect to aws data source
// @Success 200 {string} config
// @Tags connector
// @Router /api/v1/connector/connect/aws [post]
func connectAwsHandler(c *fiber.Ctx) error {
	// TODO: not implemented
	return c.SendString("wip, not implemented")
}

// @description WIP, connect to azure data source
// @Success 200 {string} config
// @Tags connector
// @Router /api/v1/connector/connect/azure [post]
func connectAzureHandler(c *fiber.Ctx) error {
	// TODO: not implemented
	return c.SendString("wip, not implemented")
}

// @description WIP, connect to gcp data source
// @Success 200 {string} config
// @Tags connector
// @Router /api/v1/connector/connect/gcp [post]
func connectGcpHandler(c *fiber.Ctx) error {
	// TODO: not implemented
	return c.SendString("wip, not implemented")
}
