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

import "context"

type (
	// Info that is set by init to carbonaut; like log level, config file path, log output, ...
	staticCtxKey string
	carbonautCtx struct {
		ctx context.Context
	}
)

const (
	LogLevel       staticCtxKey = "log-level"
	ConfigFilePath staticCtxKey = "config-filepath"
)

// Runtime context that is set upon initialization
var Carbonaut = carbonautCtx{ctx: context.Background()}

func (c *carbonautCtx) Clear() {
	c.ctx = context.Background()
}

func (c *carbonautCtx) Set(key, val any) {
	c.ctx = context.WithValue(c.ctx, key, val)
}

func (c *carbonautCtx) Get(key any) any {
	return c.ctx.Value(key)
}

func (c *carbonautCtx) GetStr(key any) string {
	return c.Get(key).(string)
}

func (c *carbonautCtx) GetLogLevel() string {
	return c.GetStr(LogLevel)
}
