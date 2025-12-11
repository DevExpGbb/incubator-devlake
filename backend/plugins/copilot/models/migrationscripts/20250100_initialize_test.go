/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package migrationscripts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitTables_Version(t *testing.T) {
	script := &initTables{}
	assert.Equal(t, uint64(20250100000000), script.Version())
}

func TestInitTables_Name(t *testing.T) {
	script := &initTables{}
	assert.Equal(t, "Initialize schema for Copilot plugin", script.Name())
}

func TestAll_ContainsInitTables(t *testing.T) {
	scripts := All()
	assert.NotEmpty(t, scripts, "All() should return at least one migration script")
	assert.Equal(t, 1, len(scripts), "All() should return exactly one migration script for now")
}
