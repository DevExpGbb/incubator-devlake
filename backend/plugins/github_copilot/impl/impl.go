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
	coreModels "github.com/apache/incubator-devlake/core/models"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/helpers/pluginhelper/subtaskmeta/sorter"
	"github.com/apache/incubator-devlake/plugins/github_copilot/api"
	"github.com/apache/incubator-devlake/plugins/github_copilot/models"
	"github.com/apache/incubator-devlake/plugins/github_copilot/models/migrationscripts"
	"github.com/apache/incubator-devlake/plugins/github_copilot/tasks"
)

var _ interface {
	plugin.PluginMeta
	plugin.PluginInit
	plugin.PluginTask
	plugin.PluginModel
	plugin.PluginMigration
	plugin.PluginSource
	plugin.DataSourcePluginBlueprintV200
	plugin.CloseablePluginTask
} = (*GithubCopilot)(nil)

type GithubCopilot struct{}

func init() {
	// check subtask meta loop when init subtask meta
	if _, err := sorter.NewDependencySorter(tasks.SubTaskMetaList).Sort(); err != nil {
		panic(err)
	}
}

func (p GithubCopilot) Init(basicRes context.BasicRes) errors.Error {
	api.Init(basicRes, p)
	return nil
}

func (p GithubCopilot) Connection() dal.Tabler {
	return &models.GithubCopilotConnection{}
}

func (p GithubCopilot) Scope() plugin.ToolLayerScope {
	return &models.GithubCopilotOrganization{}
}

func (p GithubCopilot) ScopeConfig() dal.Tabler {
	return &models.GithubCopilotScopeConfig{}
}

func (p GithubCopilot) MakeDataSourcePipelinePlanV200(
	connectionId uint64,
	scopes []*coreModels.BlueprintScope,
) (coreModels.PipelinePlan, []plugin.Scope, errors.Error) {
	return api.MakePipelinePlanV200(p.SubTaskMetas(), connectionId, scopes)
}

func (p GithubCopilot) GetTablesInfo() []dal.Tabler {
	return []dal.Tabler{
		&models.GithubCopilotConnection{},
		&models.GithubCopilotOrganization{},
		&models.GithubCopilotScopeConfig{},
		&models.GithubCopilotUsage{},
	}
}

func (p GithubCopilot) Description() string {
	return "Collect and enrich GitHub Copilot usage metrics"
}

func (p GithubCopilot) Name() string {
	return "github_copilot"
}

func (p GithubCopilot) SubTaskMetas() []plugin.SubTaskMeta {
	list, err := sorter.NewDependencySorter(tasks.SubTaskMetaList).Sort()
	if err != nil {
		panic(err)
	}
	return list
}

func (p GithubCopilot) PrepareTaskData(taskCtx plugin.TaskContext, options map[string]interface{}) (interface{}, errors.Error) {
	logger := taskCtx.GetLogger()
	logger.Debug("%v", options)
	op, err := tasks.DecodeAndValidateTaskOptions(options)
	if err != nil {
		return nil, err
	}
	if op.ConnectionId == 0 {
		return nil, errors.BadInput.New("connectionId is invalid")
	}
	connection := &models.GithubCopilotConnection{}
	connectionHelper := helper.NewConnectionHelper(
		taskCtx,
		nil,
		p.Name(),
	)
	err = connectionHelper.FirstById(connection, op.ConnectionId)
	if err != nil {
		return nil, errors.BadInput.Wrap(err, "connection not found")
	}

	apiClient, err := tasks.NewGithubCopilotApiClient(taskCtx, connection)
	if err != nil {
		return nil, err
	}

	return &tasks.GithubCopilotTaskData{
		Options:   op,
		ApiClient: apiClient,
	}, nil
}

func (p GithubCopilot) RootPkgPath() string {
	return "github.com/apache/incubator-devlake/plugins/github_copilot"
}

func (p GithubCopilot) MigrationScripts() []plugin.MigrationScript {
	return migrationscripts.All()
}

func (p GithubCopilot) Close(taskCtx plugin.TaskContext) errors.Error {
	data, ok := taskCtx.GetData().(*tasks.GithubCopilotTaskData)
	if !ok {
		return errors.Default.New("GetData failed when try to close")
	}
	data.ApiClient.Release()
	return nil
}
