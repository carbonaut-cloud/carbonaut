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
	"fmt"
	"testing"

	"carbonaut.cloud/carbonaut/pkg/api/methods"
	"carbonaut.cloud/carbonaut/pkg/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestRoutesPos(t *testing.T) {
	// check if the routes can be connected without any error (this test does not test call api endpoints)
	methods.RouteTesting(t, []routes.IRoutes{Routes{}})
}

func TestConnectorEndpointsPos(t *testing.T) {
	app := fiber.New()
	apiGroup := app.Group("api")
	port, err := freeport.GetFreePort()
	assert.NoError(t, err, "could not find a free port")
	r := Routes{}
	r.AddRoutes(apiGroup)
	r.RouteSubGroups()
	s := r.GetPrefix()
	assert.NotEmpty(t, s, "prefix should not be empty")
	go func() {
		if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
			assert.NoError(t, err)
		}
	}()

	err = connectAwsHandler(app.AcquireCtx(&fasthttp.RequestCtx{}))
	assert.NoError(t, err)
	err = connectAzureHandler(app.AcquireCtx(&fasthttp.RequestCtx{}))
	assert.NoError(t, err)
	err = connectGcpHandler(app.AcquireCtx(&fasthttp.RequestCtx{}))
	assert.NoError(t, err)
	err = connectionsHandler(app.AcquireCtx(&fasthttp.RequestCtx{}))
	assert.NoError(t, err)
}
