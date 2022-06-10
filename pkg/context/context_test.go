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

package context

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestContextPos(t *testing.T) {
	assert.Nil(t, Carbonaut.Get(""))
	assert.Nil(t, Carbonaut.Get(""))
}

type CarbonautContextTestSuite struct {
	suite.Suite
	ctx carbonautCtx
}

// TestCarbonDB runs the entire 'CarbonDBTestSuite' test suite
func TestCarbonDB(t *testing.T) {
	suite.Run(t, new(CarbonautContextTestSuite))
}

// SetupTest gets called automatically before each test of the suite to setup the test environment
func (s *CarbonautContextTestSuite) SetupTest() {
	s.ctx = carbonautCtx{}
	s.ctx.Clear()
	assert.NotNil(s.T(), s.ctx)
}

// mock test: (c *carbonautCtx) Clear()
func (s *CarbonautContextTestSuite) TestCtxClear() {
	type k string
	var a k = "a"
	ctxCp := s.ctx
	s.ctx.ctx = context.WithValue(s.ctx.ctx, a, "b")
	assert.NotEqual(s.T(), ctxCp, s.ctx)
	s.ctx.Clear()
	assert.Equal(s.T(), ctxCp, s.ctx)
}

// mock test: (c *carbonautCtx) Set(key any, val any)
func (s *CarbonautContextTestSuite) TestCtxSet() {
	v := map[any]any{"a": 0, 0: "b"}
	for k := range v {
		s.ctx.Set(k, v[k])
		assert.NotNil(s.T(), s.ctx.ctx.Value(k))
	}
	for k := range v {
		assert.Equal(s.T(), v[k], s.ctx.ctx.Value(k))
	}
}

// mock test: (c *carbonautCtx) Get(key any)
func (s *CarbonautContextTestSuite) TestCtxGet() {
	type k string
	var (
		a    k = "a"
		b      = "b"
		info   = "info"
	)
	s.ctx.ctx = context.WithValue(s.ctx.ctx, a, b)
	s.ctx.ctx = context.WithValue(s.ctx.ctx, LogLevel, info)
	assert.Equal(s.T(), b, s.ctx.GetStr(a))
	assert.Equal(s.T(), info, s.ctx.GetLogLevel())
}
