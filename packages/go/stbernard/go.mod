// Copyright 2023 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
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
//
// SPDX-License-Identifier: Apache-2.0

module github.com/specterops/bloodhound/packages/go/stbernard

go 1.23.0

toolchain go1.23.8

require (
	github.com/Masterminds/semver/v3 v3.2.1
	github.com/gofrs/uuid v4.4.0+incompatible
	github.com/specterops/bloodhound/bhlog v0.0.0-00010101000000-000000000000
	github.com/specterops/bloodhound/slicesext v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.10.0
	golang.org/x/mod v0.21.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/tools v0.26.0
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/specterops/bloodhound/bhlog => ../bhlog
	github.com/specterops/bloodhound/slicesext => ../slicesext
)
