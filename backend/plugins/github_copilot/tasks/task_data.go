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
	"time"

	"github.com/apache/incubator-devlake/core/errors"
	helper "github.com/apache/incubator-devlake/helpers/pluginhelper/api"
	"github.com/apache/incubator-devlake/plugins/github_copilot/models"
)

// GithubCopilotOptions holds the task options
type GithubCopilotOptions struct {
	ConnectionId     uint64   `json:"connectionId" mapstructure:"connectionId"`
	OrganizationName string   `json:"organizationName" mapstructure:"organizationName"`
	EnterpriseName   string   `json:"enterpriseName,omitempty" mapstructure:"enterpriseName,omitempty"`
	ScopeConfigId    uint64   `json:"scopeConfigId,omitempty" mapstructure:"scopeConfigId,omitempty"`
	Since            *time.Time `json:"since,omitempty" mapstructure:"since,omitempty"`
}

// GithubCopilotTaskData holds the task execution context
type GithubCopilotTaskData struct {
	Options   *GithubCopilotOptions
	ApiClient *helper.ApiAsyncClient
}

// DecodeAndValidateTaskOptions decodes and validates task options
func DecodeAndValidateTaskOptions(options map[string]interface{}) (*GithubCopilotOptions, errors.Error) {
	var op GithubCopilotOptions
	if err := helper.Decode(options, &op, nil); err != nil {
		return nil, err
	}
	if op.OrganizationName == "" {
		return nil, errors.BadInput.New("organizationName is required")
	}
	return &op, nil
}

// NewGithubCopilotApiClient creates a new API client for GitHub Copilot
func NewGithubCopilotApiClient(taskCtx helper.TaskContext, connection *models.GithubCopilotConnection) (*helper.ApiAsyncClient, errors.Error) {
	// GitHub API base URL
	apiUrl := connection.Endpoint
	if apiUrl == "" {
		apiUrl = "https://api.github.com"
	}

	// Create auth headers
	headers := make(map[string]string)
	if connection.Token != "" {
		headers["Authorization"] = "Bearer " + connection.Token
	}
	headers["Accept"] = "application/vnd.github+json"
	headers["X-GitHub-Api-Version"] = "2022-11-28"

	// Prepare HTTP client config
	apiClient, err := helper.NewApiClientFromConnection(taskCtx.GetContext(), taskCtx, connection)
	if err != nil {
		return nil, err
	}

	// Create async client with rate limiting
	asyncClient, err := helper.CreateAsyncApiClient(
		taskCtx,
		apiClient,
		&helper.ApiRateLimitCalculator{
			UserRateLimitPerHour: 5000, // GitHub API rate limit
		},
	)
	if err != nil {
		return nil, err
	}

	return asyncClient, nil
}

// GithubCopilotUsageResponse represents the API response structure
type GithubCopilotUsageResponse []struct {
	Day                      string `json:"day"`
	TotalSeats               int    `json:"total_seats,omitempty"`
	TotalActiveUsers         int    `json:"total_active_users,omitempty"`
	TotalEngagedUsers        int    `json:"total_engaged_users,omitempty"`
	TotalSuggestionsCount    int    `json:"total_suggestions_count,omitempty"`
	TotalAcceptancesCount    int    `json:"total_acceptances_count,omitempty"`
	TotalLinesSuggested      int    `json:"total_lines_suggested,omitempty"`
	TotalLinesAccepted       int    `json:"total_lines_accepted,omitempty"`
	TotalActiveUsersChat     int    `json:"total_active_users_chat,omitempty"`
	TotalActiveChatSessions  int    `json:"total_active_chat_sessions,omitempty"`
	Breakdown                []struct {
		Language        string `json:"language,omitempty"`
		Editor          string `json:"editor,omitempty"`
		SuggestionsCount int   `json:"suggestions_count,omitempty"`
		AcceptancesCount int   `json:"acceptances_count,omitempty"`
		LinesSuggested   int   `json:"lines_suggested,omitempty"`
		LinesAccepted    int   `json:"lines_accepted,omitempty"`
		ActiveUsers      int   `json:"active_users,omitempty"`
	} `json:"breakdown,omitempty"`
}

// ConvertToUsageModel converts API response to database model
func ConvertToUsageModel(connectionId uint64, orgName string, item interface{}) (*models.GithubCopilotUsage, errors.Error) {
	data, err := json.Marshal(item)
	if err != nil {
		return nil, errors.Convert(err)
	}
	
	var response struct {
		Day                      string `json:"day"`
		TotalSeats               int    `json:"total_seats,omitempty"`
		TotalActiveUsers         int    `json:"total_active_users,omitempty"`
		TotalEngagedUsers        int    `json:"total_engaged_users,omitempty"`
		TotalSuggestionsCount    int    `json:"total_suggestions_count,omitempty"`
		TotalAcceptancesCount    int    `json:"total_acceptances_count,omitempty"`
		TotalLinesSuggested      int    `json:"total_lines_suggested,omitempty"`
		TotalLinesAccepted       int    `json:"total_lines_accepted,omitempty"`
		TotalActiveUsersChat     int    `json:"total_active_users_chat,omitempty"`
		TotalActiveChatSessions  int    `json:"total_active_chat_sessions,omitempty"`
		Breakdown                json.RawMessage `json:"breakdown,omitempty"`
	}
	
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, errors.Convert(err)
	}
	
	day, err := time.Parse("2006-01-02", response.Day)
	if err != nil {
		return nil, errors.Convert(err)
	}
	
	usage := &models.GithubCopilotUsage{
		ConnectionId:            connectionId,
		OrganizationName:        orgName,
		Day:                     day,
		TotalSeats:              response.TotalSeats,
		TotalActiveUsers:        response.TotalActiveUsers,
		TotalEngagedUsers:       response.TotalEngagedUsers,
		TotalSuggestionsCount:   response.TotalSuggestionsCount,
		TotalAcceptancesCount:   response.TotalAcceptancesCount,
		TotalLinesSuggested:     response.TotalLinesSuggested,
		TotalLinesAccepted:      response.TotalLinesAccepted,
		TotalActiveUsersChat:    response.TotalActiveUsersChat,
		TotalActiveChatSessions: response.TotalActiveChatSessions,
	}
	
	if len(response.Breakdown) > 0 {
		usage.LanguageBreakdown = string(response.Breakdown)
		usage.EditorBreakdown = string(response.Breakdown)
	}
	
	return usage, nil
}
