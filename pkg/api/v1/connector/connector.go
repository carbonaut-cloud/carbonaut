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
	"carbonaut.cloud/carbonaut/pkg/api/models"
	"github.com/gofiber/fiber/v2"
)

type Routes struct{}

func (c Routes) GetPrefix() string {
	return "connector"
}

func (c Routes) RouteSubGroups() []models.IRoutes {
	return []models.IRoutes{}
}

func (c Routes) AddRoutes(r fiber.Router) {
	r.Get("/*", helloConnectorAPIHandler)
}

// @description Dummy connector API test
// @Success 200 {string} connector
// @Tags connector
// @Router /api/v1/connector/ [get]
func helloConnectorAPIHandler(c *fiber.Ctx) error {
	return c.SendString("connector")
}
