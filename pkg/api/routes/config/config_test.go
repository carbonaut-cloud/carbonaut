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

// func TestRoutesPos(t *testing.T) {
// 	// check if the routes can be connected without any error (this test does not test call api endpoints)
// 	methods.RouteTesting(t, []routes.IRoutes{Routes{}})
// }

// func TestConfigEndpointsPos(t *testing.T) {
// 	app := fiber.New()
// 	apiGroup := app.Group("api")
// 	port, err := freeport.GetFreePort()
// 	assert.NoError(t, err, "could not find a free port")
// 	r := Routes{}
// 	r.AddRoutes(apiGroup)
// 	r.RouteSubGroups()
// 	s := r.GetPrefix()
// 	assert.NotEmpty(t, s, "prefix should not be empty")
// 	go func() {
// 		if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
// 			assert.NoError(t, err)
// 		}
// 	}()

// 	err = loadHandler(app.AcquireCtx(&fasthttp.RequestCtx{}))
// 	assert.NoError(t, err)
// 	err = validateHandler(app.AcquireCtx(&fasthttp.RequestCtx{}))
// 	assert.NoError(t, err)
// 	err = describeHandler(app.AcquireCtx(&fasthttp.RequestCtx{}))
// 	assert.NoError(t, err)
// }

// // https://github.com/gofiber/recipes/blob/master/unit-test/main.go
// func TestIndexRoute(t *testing.T) {
// 	tests := []struct {
// 		description   string
// 		route         string
// 		expectedError bool
// 		expectedCode  int
// 		expectedBody  string
// 	}{
// 		{
// 			description:   "index route",
// 			route:         "/",
// 			expectedError: false,
// 			expectedCode:  200,
// 			expectedBody:  "OK",
// 		},
// 		{
// 			description:   "non existing route",
// 			route:         "/i-dont-exist",
// 			expectedError: false,
// 			expectedCode:  404,
// 			expectedBody:  "Cannot GET /i-dont-exist",
// 		},
// 	}

// 	app := fiber.New()

// 	for _, test := range tests {
// 		req, _ := http.NewRequest(
// 			"GET",
// 			test.route,
// 			nil,
// 		)
// 		res, err := app.Test(req, -1)
// 		assert.Equalf(t, test.expectedError, err != nil, test.description)
// if test.expectedError {
// 			continue
// 		}
// 		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
// 		body, err := ioutil.ReadAll(res.Body)
// 		assert.Nilf(t, err, test.description)
// 		assert.Equalf(t, test.expectedBody, string(body), test.description)
// 	}
// }
