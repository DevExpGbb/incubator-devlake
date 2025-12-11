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

// CopilotSeat represents a Copilot seat assignment
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
