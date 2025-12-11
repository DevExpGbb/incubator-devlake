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

// CopilotOrgMetrics represents aggregated Copilot usage metrics for an organization
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
