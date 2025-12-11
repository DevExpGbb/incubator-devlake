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

package archived

import (
	"time"

	"github.com/apache/incubator-devlake/core/models/common"
	"github.com/apache/incubator-devlake/core/plugin"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
)

type CopilotConnection struct {
	helper.RestConnection `mapstructure:",squash"`
	common.Model
}

func (CopilotConnection) TableName() string {
	return "_tool_copilot_connections"
}

type CopilotScope struct {
	common.Scope `mapstructure:",squash"`
	Organization string `json:"organization" mapstructure:"organization" gorm:"primaryKey;type:varchar(255)"`
	Name         string `json:"name" mapstructure:"name" gorm:"type:varchar(255)"`
}

func (CopilotScope) TableName() string {
	return "_tool_copilot_scopes"
}

func (s CopilotScope) ScopeId() string {
	return s.Organization
}

func (s CopilotScope) ScopeName() string {
	if s.Name != "" {
		return s.Name
	}
	return s.Organization
}

func (s CopilotScope) ScopeFullName() string {
	return s.Organization
}

func (s CopilotScope) ScopeParams() interface{} {
	return nil
}

var _ plugin.ToolLayerScope = (*CopilotScope)(nil)

type CopilotOrgMetrics struct {
	common.NoPKModel
	ConnectionId         uint64    `json:"connection_id" gorm:"primaryKey;type:BIGINT"`
	Organization         string    `json:"organization" gorm:"primaryKey;type:varchar(255)"`
	Date                 time.Time `json:"date" gorm:"primaryKey;type:date"`
	TotalSuggestionsCount int      `json:"total_suggestions_count"`
	TotalAcceptancesCount int      `json:"total_acceptances_count"`
	TotalLinesSuggested  int       `json:"total_lines_suggested"`
	TotalLinesAccepted   int       `json:"total_lines_accepted"`
	TotalActiveUsers     int       `json:"total_active_users"`
	TotalChatAcceptances int       `json:"total_chat_acceptances"`
	TotalChatTurns       int       `json:"total_chat_turns"`
	TotalActiveChatUsers int       `json:"total_active_chat_users"`
}

func (CopilotOrgMetrics) TableName() string {
	return "_tool_copilot_org_metrics"
}

type CopilotLanguageMetrics struct {
	common.NoPKModel
	ConnectionId         uint64    `json:"connection_id" gorm:"primaryKey;type:BIGINT"`
	Organization         string    `json:"organization" gorm:"primaryKey;type:varchar(255)"`
	Date                 time.Time `json:"date" gorm:"primaryKey;type:date"`
	Language             string    `json:"language" gorm:"primaryKey;type:varchar(100)"`
	TotalSuggestionsCount int      `json:"total_suggestions_count"`
	TotalAcceptancesCount int      `json:"total_acceptances_count"`
	TotalLinesSuggested  int       `json:"total_lines_suggested"`
	TotalLinesAccepted   int       `json:"total_lines_accepted"`
	TotalActiveUsers     int       `json:"total_active_users"`
}

func (CopilotLanguageMetrics) TableName() string {
	return "_tool_copilot_language_metrics"
}

type CopilotSeat struct {
	common.NoPKModel
	ConnectionId      uint64     `json:"connection_id" gorm:"primaryKey;type:BIGINT"`
	Organization      string     `json:"organization" gorm:"primaryKey;type:varchar(255)"`
	Assignee          string     `json:"assignee" gorm:"primaryKey;type:varchar(255)"`
	AssigningTeam     string     `json:"assigning_team" gorm:"type:varchar(255)"`
	CreatedAt         time.Time  `json:"created_at" gorm:"type:datetime"`
	UpdatedAt         time.Time  `json:"updated_at" gorm:"type:datetime"`
	LastActivityAt    *time.Time `json:"last_activity_at" gorm:"type:datetime"`
	LastActivityEditor string    `json:"last_activity_editor" gorm:"type:varchar(100)"`
}

func (CopilotSeat) TableName() string {
	return "_tool_copilot_seats"
}
