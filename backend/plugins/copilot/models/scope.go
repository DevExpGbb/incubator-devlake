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
	"github.com/apache/incubator-devlake/core/models/common"
	"github.com/apache/incubator-devlake/core/plugin"
)

// CopilotScope represents a GitHub organization for Copilot metrics collection
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
	return &CopilotScopeParams{
		ConnectionId: s.ConnectionId,
		Organization: s.Organization,
	}
}

type CopilotScopeParams struct {
	ConnectionId uint64 `json:"connection_id"`
	Organization string `json:"organization"`
}

var _ plugin.ToolLayerScope = (*CopilotScope)(nil)
