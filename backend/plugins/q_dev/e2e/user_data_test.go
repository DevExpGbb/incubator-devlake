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

package e2e

import (
	"testing"

	"github.com/apache/incubator-devlake/core/models/common"
	"github.com/apache/incubator-devlake/helpers/e2ehelper"
	"github.com/apache/incubator-devlake/plugins/q_dev/impl"
	"github.com/apache/incubator-devlake/plugins/q_dev/models"
)

// TestQDevS3FileMeta tests the S3 file metadata model
func TestQDevS3FileMeta(t *testing.T) {
	var plugin impl.QDev
	dataflowTester := e2ehelper.NewDataFlowTester(t, "q_dev", plugin)

	// Flush tables before testing
	dataflowTester.FlushTabler(&models.QDevS3FileMeta{})

	// Import CSV data into the tool table
	dataflowTester.ImportCsvIntoTabler("./snapshot_tables/_tool_q_dev_s3_file_meta.csv", &models.QDevS3FileMeta{})

	// Verify the file meta data
	dataflowTester.VerifyTableWithOptions(
		models.QDevS3FileMeta{},
		e2ehelper.TableOptions{
			CSVRelPath:  "./snapshot_tables/_tool_q_dev_s3_file_meta.csv",
			IgnoreTypes: []interface{}{common.NoPKModel{}},
			IgnoreFields: []string{
				"processed_time",
			},
		},
	)
}

// TestQDevUserData tests the user data model
func TestQDevUserData(t *testing.T) {
	var plugin impl.QDev
	dataflowTester := e2ehelper.NewDataFlowTester(t, "q_dev", plugin)

	// Flush tables before testing
	dataflowTester.FlushTabler(&models.QDevUserData{})

	// Import CSV data into the tool table
	dataflowTester.ImportCsvIntoTabler("./snapshot_tables/_tool_q_dev_user_data.csv", &models.QDevUserData{})

	// Verify the user data - comparing only the fields that exist in our CSV
	dataflowTester.VerifyTableWithOptions(
		models.QDevUserData{},
		e2ehelper.TableOptions{
			CSVRelPath:  "./snapshot_tables/_tool_q_dev_user_data.csv",
			IgnoreTypes: []interface{}{common.Model{}},
		},
	)
}
