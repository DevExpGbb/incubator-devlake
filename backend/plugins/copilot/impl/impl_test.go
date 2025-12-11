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

package impl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopilot_Name(t *testing.T) {
	plugin := Copilot{}
	assert.Equal(t, "copilot", plugin.Name())
}

func TestCopilot_Description(t *testing.T) {
	plugin := Copilot{}
	assert.NotEmpty(t, plugin.Description())
}

func TestCopilot_GetTablesInfo(t *testing.T) {
	plugin := Copilot{}
	tables := plugin.GetTablesInfo()
	assert.Equal(t, 5, len(tables), "Should have 5 tables: Connection, Scope, OrgMetrics, LanguageMetrics, Seat")

	// Verify all tables have non-empty names
	for _, table := range tables {
		assert.NotEmpty(t, table.TableName())
	}
}

func TestCopilot_MigrationScripts(t *testing.T) {
	plugin := Copilot{}
	scripts := plugin.MigrationScripts()
	assert.NotEmpty(t, scripts, "Should have at least one migration script")
}

func TestCopilot_SubTaskMetas(t *testing.T) {
	plugin := Copilot{}
	metas := plugin.SubTaskMetas()
	// Initially empty since we haven't added any subtasks yet
	assert.Equal(t, 0, len(metas), "Should have no subtasks in the foundational plumbing")
}

func TestCopilot_RootPkgPath(t *testing.T) {
	plugin := Copilot{}
	assert.Equal(t, "github.com/apache/incubator-devlake/plugins/copilot", plugin.RootPkgPath())
}
