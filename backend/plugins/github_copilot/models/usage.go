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

package models

import (
	"time"

	"github.com/apache/incubator-devlake/core/models/common"
)

// GithubCopilotUsage stores Copilot usage metrics for users
type GithubCopilotUsage struct {
	common.NoPKModel
	ConnectionId     uint64    `gorm:"primaryKey"`
	OrganizationName string    `gorm:"primaryKey;type:varchar(255)"`
	Day              time.Time `gorm:"primaryKey;type:date"`
	
	// User information
	TotalSeats            int `json:"total_seats"`
	TotalActiveUsers      int `json:"total_active_users"`
	TotalEngagedUsers     int `json:"total_engaged_users"`
	
	// Code suggestions metrics
	TotalSuggestionsCount    int `json:"total_suggestions_count"`
	TotalAcceptancesCount    int `json:"total_acceptances_count"`
	TotalLinesSuggested      int `json:"total_lines_suggested"`
	TotalLinesAccepted       int `json:"total_lines_accepted"`
	TotalActiveUsersChat     int `json:"total_active_users_chat"`
	TotalActiveChatSessions  int `json:"total_active_chat_sessions"`
	
	// Language breakdown (stored as JSON)
	LanguageBreakdown string `gorm:"type:text" json:"language_breakdown,omitempty"`
	// Editor breakdown (stored as JSON)
	EditorBreakdown   string `gorm:"type:text" json:"editor_breakdown,omitempty"`
	
	common.NoPKModel `json:"-" mapstructure:"-" swaggerignore:"true"`
}

func (GithubCopilotUsage) TableName() string {
	return "_tool_github_copilot_usage"
}
