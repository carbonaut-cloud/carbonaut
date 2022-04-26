# Copyright 2022 The Carbonaut Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

install: 
	./scripts/install-go-dependecies.sh

verify: verify-go-mod verify-build verify-lint verify-test-unit verify-git

verify-build:
	./scripts/verify-build.sh

verify-lint:
	./scripts/verify-golangci-lint.sh

verify-test-unit:
	./scripts/verify-test-go.sh

verify-go-mod:
	./scripts/verify-go-mod.sh

verify-git:
	./scripts/verify-git.sh

build:
	echo "TODO: build go binaries"
	echo "TODO: compile protobuf API code"