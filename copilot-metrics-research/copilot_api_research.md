# GitHub Copilot API Research Notes

## Date: December 5, 2025

---

## API Endpoints Summary

There are **3 main API categories** for GitHub Copilot data:

### 1. Copilot Usage Metrics Reports (Aggregated + User-Level)
**Primary Endpoints (recommended)**:
- **Enterprise (28-day latest)**: `GET /enterprises/{enterprise}/copilot/metrics/reports/enterprise-28-day/latest`
- **Enterprise (1-day)**: `GET /enterprises/{enterprise}/copilot/metrics/reports/enterprise-1-day?day=YYYY-MM-DD`
- **Enterprise users (28-day latest)**: `GET /enterprises/{enterprise}/copilot/metrics/reports/users-28-day/latest`
- **Enterprise users (1-day)**: `GET /enterprises/{enterprise}/copilot/metrics/reports/users-1-day?day=YYYY-MM-DD`
- **Organization (28-day latest)**: `GET /orgs/{org}/copilot/metrics/reports/organization-28-day/latest`
- **Organization (1-day)**: `GET /orgs/{org}/copilot/metrics/reports/organization-1-day?day=YYYY-MM-DD`
- **Organization users (28-day latest)**: `GET /orgs/{org}/copilot/metrics/reports/users-28-day/latest`
- **Organization users (1-day)**: `GET /orgs/{org}/copilot/metrics/reports/users-1-day?day=YYYY-MM-DD`

**Key Features**:
- **Two-step flow**: request report metadata → download signed JSON files
- **Aggregated and user-level** usage reports
- Reports generated daily and available for **up to 1 year**
- Signed URLs expire (download promptly)
- **Expanded telemetry** (IDE agents, edit modes, models, languages, LOC)
- **PERFECT for repo/team-level analysis (Option B)**

**Prerequisite**: Enterprise policy "Copilot usage metrics" must be enabled.

**Report Metadata Response** (download links):
```json
{
  "download_links": [
    "https://example.com/copilot-usage-report-1.json",
    "https://example.com/copilot-usage-report-2.json"
  ],
  "report_start_day": "2025-07-01",
  "report_end_day": "2025-07-28"
}
```

**Report File Structure** (per day record):
```json
{
  "day": "2024-06-24",
  "total_active_users": 24,
  "total_engaged_users": 20,
  "daily_active_users": 24,
  "weekly_active_users": 20,
  "monthly_active_users": 40,
  "user_initiated_interaction_count": 120,
  "code_generation_activity_count": 350,
  "code_acceptance_activity_count": 180,
  "loc_suggested_to_add_sum": 500,
  "loc_suggested_to_delete_sum": 120,
  "loc_added_sum": 260,
  "loc_deleted_sum": 80
}
```

**Authentication**:
- Enterprise reports: PAT (classic) `manage_billing:copilot` or `read:enterprise`
- Org reports: PAT (classic) `read:org`
- Fine-grained: "Enterprise Copilot metrics" or "Organization Copilot metrics" (read)

---

### 2. Copilot Metrics API (Legacy - Deprecated)
**Deprecated Endpoints**:
- `GET /orgs/{org}/copilot/metrics`
- `GET /enterprises/{enterprise}/copilot/metrics`

**Status**: Sunsets April 2, 2026. Migrate to the usage metrics reports above.

---

### 3. Copilot User Management API (Seat/Adoption data)
**Endpoints**:
- `GET /orgs/{org}/copilot/billing` - Org-level seat summary
- `GET /orgs/{org}/copilot/billing/seats` - List all seat assignments
- `GET /orgs/{org}/members/{username}/copilot` - Individual user seat details

**Seat Summary Response**:
```json
{
  "seat_breakdown": {
    "total": 12,
    "added_this_cycle": 9,
    "pending_invitation": 0,
    "pending_cancellation": 0,
    "active_this_cycle": 12,
    "inactive_this_cycle": 11
  },
  "seat_management_setting": "assign_selected",
  "ide_chat": "enabled",
  "platform_chat": "enabled",
  "cli": "enabled",
  "public_code_suggestions": "block",
  "plan_type": "business"
}
```

**Individual Seat Response**:
```json
{
  "created_at": "2021-08-03T18:00:00-06:00",
  "updated_at": "2021-09-23T15:00:00-06:00",
  "pending_cancellation_date": null,
  "last_activity_at": "2021-10-14T00:53:32-06:00",
  "last_activity_editor": "vscode/1.77.3/copilot/1.86.82",
  "plan_type": "business",
  "assignee": {
    "login": "octocat",
    "id": 1
  },
  "assigning_team": {
    "name": "Justice League",
    "slug": "justice-league"
  }
}
```

---

## Key Metrics Available

### IDE Code Completions
| Metric | Description |
|--------|-------------|
| `total_code_suggestions` | Number of suggestions shown |
| `total_code_acceptances` | Number of suggestions accepted |
| `total_code_lines_suggested` | Lines of code suggested |
| `total_code_lines_accepted` | Lines of code accepted |
| `total_engaged_users` | Users who interacted |

### IDE Chat
| Metric | Description |
|--------|-------------|
| `total_chats` | Total chat conversations |
| `total_chat_insertion_events` | Code inserted from chat |
| `total_chat_copy_events` | Code copied from chat |

### GitHub.com Chat
| Metric | Description |
|--------|-------------|
| `total_chats` | Chats on github.com |
| `total_engaged_users` | Users engaged |

### Pull Requests (Copilot for PRs)
| Metric | Description |
|--------|-------------|
| `total_pr_summaries_created` | PR summaries generated |
| `total_engaged_users` | Users using PR features |
| **By Repository** | Breakdown available per repo! |

---

## Critical Insight: Repository-Level Data Available! 🎯

The `copilot_dotcom_pull_requests` section includes **repository breakdown**:
```json
"repositories": [
  {
    "name": "demo/repo1",
    "total_engaged_users": 8,
    "models": [{ "total_pr_summaries_created": 6 }]
  }
]
```

This is PERFECT for Option B (repo-level correlation)!

---

## Implementation Considerations

### For Option B (Repo/Project-Level Analysis):

1. **Primary Data Source**: Copilot usage metrics reports (enterprise/org)
  - Two-step flow: request report → download JSON
  - Use 1-day reports for incremental collection
  - Use 28-day latest for backfill or validation

2. **Scope Definition**: 
   - Scope = Organization (or Team within org)
   - Store daily aggregates with date range

3. **Copilot Implementation Date Tracking**:
   - Use seat management API to find earliest `created_at` for seats
   - Or allow manual configuration of "Copilot rollout date"

4. **Correlation Strategy**:
   - Join Copilot daily metrics with DevLake's `project_pr_metrics`
   - Compare:
     - PR cycle time before/after Copilot adoption date
     - Deployment frequency trends
     - Code review time changes

### Required Permissions
- Enterprise reports: `manage_billing:copilot` or `read:enterprise`
- Org reports: `read:org`
- Org owner or billing manager access needed

---

## Next Steps

1. ✅ Test API with `octodemo` org using provided GH_PAT
2. Design data model based on metrics response structure
3. Determine how to store/flatten the nested JSON structure
4. Plan Grafana dashboard SQL for before/after comparisons
