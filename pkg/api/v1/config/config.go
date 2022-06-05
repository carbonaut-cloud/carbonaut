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

package config

import (
	"carbonaut.cloud/carbonaut/pkg/api/routes"
	"github.com/gofiber/fiber/v2"
)

type Routes struct{}

func (c Routes) GetPrefix() string {
	return "config"
}

func (c Routes) RouteSubGroups() []routes.IRoutes {
	return []routes.IRoutes{}
}

func (c Routes) AddRoutes(r fiber.Router) {
	r.Post("/validate", validateHandler)
	r.Get("/describe", describeHandler)
	r.Put("/load", loadHandler)
}

// @description WIP, validate provided carbonaut configuration
// @Success 200 {string} config
// @Tags config
// @Router /api/v1/config/validate [post]
func validateHandler(c *fiber.Ctx) error {
	// TODO: not implemented
	return c.SendString("wip, not implemented")
}

// @description WIP, describe current carbonaut configuration
// @Success 200 {string} config
// @Tags config
// @Router /api/v1/config/describe [get]
func describeHandler(c *fiber.Ctx) error {
	// TODO: not implemented
	return c.SendString("wip, not implemented")
}

// @description WIP, update carbonaut configuration
// @Success 200 {string} config
// @Tags config
// @Router /api/v1/config/load [put]
func loadHandler(c *fiber.Ctx) error {
	// TODO: not implemented
	return c.SendString("wip, not implemented")
}
