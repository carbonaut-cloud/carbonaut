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

package data

import (
	"fmt"

	"carbonaut.cloud/carbonaut/pkg/api/routes"
	"github.com/gofiber/fiber/v2"
)

type Routes struct{}

func (c Routes) GetPrefix() string {
	return "data"
}

func (c Routes) RouteSubGroups() []routes.IRoutes {
	return []routes.IRoutes{ImportRoutes{}, ExportRoutes{}}
}

func (c Routes) AddRoutes(r fiber.Router) {
	r.Post("/storage", createStorageDataHandler)
	r.Get("/storage", describeStorageDataHandler)
}

// @description Configure a connection to storage
// @Success 200 {string} data
// @Tags data
// @Router /api/v1/data/storage [post]
func createStorageDataHandler(c *fiber.Ctx) error {
	// TODO: not implemented
	return c.SendString("wip, not implemented")
}

// @description Describe carbonaut storage connection
// @Success 200 {string} data
// @Tags data
// @Router /api/v1/data/storage [get]
func describeStorageDataHandler(c *fiber.Ctx) error {
	// TODO: not implemented
	return c.SendString("wip, not implemented")
}

//
// Import data
//

type ImportRoutes struct{}

func (c ImportRoutes) GetPrefix() string {
	return "import"
}

func (c ImportRoutes) RouteSubGroups() []routes.IRoutes {
	return []routes.IRoutes{}
}

const importProvider = "provider"

func (c ImportRoutes) AddRoutes(r fiber.Router) {
	r.Post(fmt.Sprintf("/raw/:%s?", importProvider), importRawDataHandler)
	r.Post(fmt.Sprintf("/csv/:%s?", importProvider), importCsvFileDataHandler)
}

// @description Import raw bytes of provider data to carbonaut
// @Success 200 {string} data
// @Tags data
// @Param provider query string true "Used to match provided data format to provider"
// @Router /api/v1/data/import/raw [post]
// @Accept plain
func importRawDataHandler(c *fiber.Ctx) error {
	// TODO: not implemented
	return c.SendString("wip, not implemented")
}

// @description Import data in a csv file to carbonaut
// @Success 200 {string} data
// @Tags data
// @Param provider query string true "Used to match provided data format to provider"
// @Router /api/v1/data/import/csv [post]
// @Accept plain
func importCsvFileDataHandler(c *fiber.Ctx) error {
	// TODO: not implemented
	return c.SendString("wip, not implemented")
}

//
// Export data
//

type ExportRoutes struct{}

func (c ExportRoutes) GetPrefix() string {
	return "export"
}

func (c ExportRoutes) RouteSubGroups() []routes.IRoutes {
	return []routes.IRoutes{}
}

func (c ExportRoutes) AddRoutes(r fiber.Router) {
	r.Get("/", exportDataHandler)
}

// @description Export carbonaut data
// @Success 200 {string} data
// @Tags data
// @Router /api/v1/data/export [get]
func exportDataHandler(c *fiber.Ctx) error {
	// TODO: not implemented
	return c.SendString("wip, not implemented")
}
