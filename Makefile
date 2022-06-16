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

verify: swag verify-go-mod verify-git verify-build verify-lint verify-test-unit

verify-build:
	./scripts/verify-build.sh

verify-lint:
	./scripts/verify-golangci-lint.sh

verify-test-unit:
	./scripts/verify-test-go.sh

verify-go-mod:
	go vet ./...
	go mod tidy -compat=1.18

verify-git:
	git diff --exit-code

#

build-podman:
	podman build -f Containerfile -t carbonaut:latest .

run-podman:
	podman run -p 3000:3000 carbonaut:latest

build-docker:
	docker build -f Containerfile -t carbonaut:latest .

run-docker:
	docker run -p 3000:3000 carbonaut:latest

verify-container-image:
	trivy image carbonaut:latest

upgrade:
	go get -u -t ./...

install: 
	# install swagger tool to compile swagger carbonaut api definition 
	go install github.com/swaggo/swag/cmd/swag@v1.8.1
	go get ./...

swag:
	swag init --dir "./pkg/api/"
