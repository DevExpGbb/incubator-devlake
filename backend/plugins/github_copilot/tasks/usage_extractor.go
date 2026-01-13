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

package tasks

import (
	"encoding/json"

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/github_copilot/models"
)

var _ plugin.SubTaskEntryPoint = ExtractUsage

var ExtractUsageMeta = plugin.SubTaskMeta{
	Name:             "ExtractUsage",
	EntryPoint:       ExtractUsage,
	EnabledByDefault: true,
	Description:      "Extract GitHub Copilot usage data from raw layer",
	DomainTypes:      []string{},
	Dependencies:     []*plugin.SubTaskMeta{&CollectUsageMeta},
}

func ExtractUsage(taskCtx plugin.SubTaskContext) errors.Error {
	data := taskCtx.GetData().(*GithubCopilotTaskData)
	logger := taskCtx.GetLogger()
	logger.Info("Extracting GitHub Copilot usage data")

	extractor, err := helper.NewApiExtractor(helper.ApiExtractorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: GithubCopilotApiParams{
				ConnectionId:     data.Options.ConnectionId,
				OrganizationName: data.Options.OrganizationName,
			},
			Table: RAW_USAGE_TABLE,
		},
		Extract: func(row *helper.RawData) ([]interface{}, errors.Error) {
			var items []json.RawMessage
			err := json.Unmarshal(row.Data, &items)
			if err != nil {
				return nil, errors.Convert(err)
			}

			var results []interface{}
			for _, item := range items {
				usage, err := ConvertToUsageModel(data.Options.ConnectionId, data.Options.OrganizationName, item)
				if err != nil {
					logger.Warn(err, "failed to convert usage data")
					continue
				}
				results = append(results, usage)
			}
			return results, nil
		},
	})

	if err != nil {
		return err
	}

	return extractor.Execute()
}
