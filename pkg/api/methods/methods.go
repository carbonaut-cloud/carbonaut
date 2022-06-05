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
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"carbonaut.cloud/carbonaut/pkg/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
)

func AddSubRoutes(r []routes.IRoutes, routeGroup fiber.Router) {
	// if there aren't any routes left break
	if len(r) == 0 {
		return
	}
	subGroup := routeGroup.Group(r[0].GetPrefix())
	r[0].AddRoutes(subGroup)
	// recursion to add routes of subgroups (e.g. api/v1/data/db)
	AddSubRoutes(r[0].RouteSubGroups(), subGroup)
	// remove first element from the list
	AddSubRoutes(append(r[:0], r[1:]...), routeGroup)
}

func VerifyGetResponse(t *testing.T, p int, route, expectedBodyResult string) {
	respBar, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/%s/", p, route))
	assert.NoError(t, err)
	respBarBody, err := ioutil.ReadAll(respBar.Body)
	assert.NoError(t, err)
	assert.Equal(t, expectedBodyResult, string(respBarBody))
}

func RouteTesting(t *testing.T, r []routes.IRoutes) int {
	port, err := freeport.GetFreePort()
	assert.NoError(t, err, "could not find a free port")
	app := fiber.New()
	AddSubRoutes(r, app)
	go func() {
		err := app.Listen(fmt.Sprintf(":%d", port))
		assert.NoError(t, err)
	}()
	time.Sleep(time.Millisecond * 20)
	return port
}
