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
	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/dal"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	"github.com/apache/incubator-devlake/plugins/copilot/api"
	"github.com/apache/incubator-devlake/plugins/copilot/models"
	"github.com/apache/incubator-devlake/plugins/copilot/models/migrationscripts"
)

var _ interface {
	plugin.PluginMeta
	plugin.PluginInit
	plugin.PluginTask
	plugin.PluginApi
	plugin.PluginModel
	plugin.PluginSource
	plugin.PluginMigration
} = (*Copilot)(nil)

type Copilot struct{}

func (p Copilot) Init(basicRes context.BasicRes) errors.Error {
	api.Init(basicRes, p)
	return nil
}

func (p Copilot) GetTablesInfo() []dal.Tabler {
	return []dal.Tabler{
		&models.CopilotConnection{},
	}
}

func (p Copilot) Description() string {
	return "Collect and analyze GitHub Copilot usage metrics"
}

func (p Copilot) Name() string {
	return "copilot"
}

func (p Copilot) Connection() dal.Tabler {
	return &models.CopilotConnection{}
}

func (p Copilot) Scope() plugin.ToolLayerScope {
	return nil
}

func (p Copilot) ScopeConfig() dal.Tabler {
	return nil
}

func (p Copilot) SubTaskMetas() []plugin.SubTaskMeta {
	return []plugin.SubTaskMeta{}
}

func (p Copilot) PrepareTaskData(taskCtx plugin.TaskContext, options map[string]interface{}) (interface{}, errors.Error) {
	return nil, nil
}

func (p Copilot) RootPkgPath() string {
	return "github.com/apache/incubator-devlake/plugins/copilot"
}

func (p Copilot) MigrationScripts() []plugin.MigrationScript {
	return migrationscripts.All()
}

func (p Copilot) ApiResources() map[string]map[string]plugin.ApiResourceHandler {
	return map[string]map[string]plugin.ApiResourceHandler{
		"test": {
			"POST": api.TestConnection,
		},
		"connections": {
			"POST": api.PostConnections,
			"GET":  api.ListConnections,
		},
		"connections/:connectionId": {
			"PATCH":  api.PatchConnection,
			"DELETE": api.DeleteConnection,
			"GET":    api.GetConnection,
		},
	}
}
