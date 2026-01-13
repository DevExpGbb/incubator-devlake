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
	"time"

	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/models/migrationscripts/archived"
	"github.com/apache/incubator-devlake/helpers/migrationhelper"
)

type addInitTables struct{}

type githubCopilotConnection20260113 struct {
	archived.RestConnection `mapstructure:",squash"`
	archived.AccessToken    `mapstructure:",squash"`
}

func (githubCopilotConnection20260113) TableName() string {
	return "_tool_github_copilot_connections"
}

type githubCopilotOrganization20260113 struct {
	archived.ScopeConfig
	OrganizationName string `gorm:"type:varchar(255);index"`
	EnterpriseName   string `gorm:"type:varchar(255)"`
}

func (githubCopilotOrganization20260113) TableName() string {
	return "_tool_github_copilot_organizations"
}

type githubCopilotScopeConfig20260113 struct {
	archived.ScopeConfig
}

func (githubCopilotScopeConfig20260113) TableName() string {
	return "_tool_github_copilot_scope_configs"
}

type githubCopilotUsage20260113 struct {
	archived.NoPKModel
	ConnectionId     uint64    `gorm:"primaryKey"`
	OrganizationName string    `gorm:"primaryKey;type:varchar(255)"`
	Day              time.Time `gorm:"primaryKey;type:date"`
	
	TotalSeats               int `json:"total_seats"`
	TotalActiveUsers         int `json:"total_active_users"`
	TotalEngagedUsers        int `json:"total_engaged_users"`
	TotalSuggestionsCount    int `json:"total_suggestions_count"`
	TotalAcceptancesCount    int `json:"total_acceptances_count"`
	TotalLinesSuggested      int `json:"total_lines_suggested"`
	TotalLinesAccepted       int `json:"total_lines_accepted"`
	TotalActiveUsersChat     int `json:"total_active_users_chat"`
	TotalActiveChatSessions  int `json:"total_active_chat_sessions"`
	LanguageBreakdown        string `gorm:"type:text"`
	EditorBreakdown          string `gorm:"type:text"`
}

func (githubCopilotUsage20260113) TableName() string {
	return "_tool_github_copilot_usage"
}

func (*addInitTables) Up(basicRes context.BasicRes) errors.Error {
	return migrationhelper.AutoMigrateTables(
		basicRes,
		&githubCopilotConnection20260113{},
		&githubCopilotOrganization20260113{},
		&githubCopilotScopeConfig20260113{},
		&githubCopilotUsage20260113{},
	)
}

func (*addInitTables) Version() uint64 {
	return 20260113000001
}

func (*addInitTables) Name() string {
	return "GitHub Copilot init schemas"
}
