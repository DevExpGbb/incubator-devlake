# GitHub Copilot Dashboards

This directory contains Grafana dashboards for visualizing GitHub Copilot usage metrics.

## Dashboards

### 1. GithubCopilot.json - GitHub Copilot User Data Dashboard

Main dashboard for monitoring GitHub Copilot usage across your organization.

#### Panels

##### Overview Section
- **Overall Usage Statistics** - Key metrics summary showing active users, accepted lines, and acceptance rates

##### New Metric Cards (Row 2)
- **Average Acceptance Rate** - Shows the average acceptance rate of AI suggestions over the selected time period
  - Color-coded thresholds: Red (<30%), Yellow (30-60%), Green (>60%)
  - Displayed as percentage
  
- **Engaged Users %** - Percentage of users who accepted at least one suggestion
  - Shows delta/change from previous time period
  - Color-coded thresholds: Red (<40%), Yellow (40-70%), Green (>70%)
  
- **Editor Distribution** - Pie chart showing distribution of GitHub Copilot usage across different editors/IDEs
  - **Shows percentages when highlighted** (hover tooltips)
  - Legend displays both raw values and percentages
  - Labels show percentages on the chart

##### Time Series Section
- **Daily AI Code Line Changes** - Track code generation trends over time
- **Daily AI Interactions** - Monitor user engagement with AI features
- **Code Review Metrics** - Code review findings and fixes
- **Daily AI Suggestion Acceptance Rate** - Acceptance rate trends

##### Details Section
- **User Interactions** - Detailed table of per-user metrics

### 2. GithubCopilotDORA.json - GitHub Copilot + DORA Correlation Dashboard

Correlates GitHub Copilot usage with DORA (DevOps Research and Assessment) metrics to understand the impact of AI-assisted development.

#### Panels
- **GitHub Copilot Active Users** - Number of users actively using GitHub Copilot
- **Total AI Code Lines** - Total lines of code generated/accepted
- **AI Acceptance Rate** - Overall acceptance rate stat
- **Total Deployments** - DORA deployment frequency
- **Median Lead Time** - DORA lead time metric
- **Change Failure Rate** - DORA quality metric
- **AI Code Generation vs Lead Time Trend** - Correlation chart
- **AI Acceptance Rate vs Deployment Frequency** - Correlation chart
- **AI Test Generation vs Change Failure Rate** - Quality correlation
- **GitHub Copilot Users vs Code Review Findings** - Code quality impact
- **Monthly Comparison Table** - Comprehensive month-over-month comparison

## Data Requirements

These dashboards expect data in the `_tool_github_copilot_user_data` table with the following fields:

- `user_id` - Unique user identifier
- `display_name` - User's display name
- `date` - Date of the metrics
- `editor_name` - Name of the editor/IDE used (e.g., "VS Code", "IntelliJ IDEA", "Vim")
- `inline_acceptance_count` - Number of inline suggestions accepted
- `inline_suggestions_count` - Total number of inline suggestions shown
- `inline_ai_code_lines` - Lines of code from accepted inline suggestions
- `chat_ai_code_lines` - Lines of code from chat interactions
- `code_fix_acceptance_event_count` - Code fix acceptances
- `code_fix_generation_event_count` - Code fix generations
- `code_review_findings_count` - Code review findings
- Other metrics similar to Q Dev structure

## Installation

1. Copy the dashboard JSON files to your Grafana dashboards directory
2. Import them through the Grafana UI:
   - Go to Dashboards → Import
   - Upload the JSON file or paste the JSON content
   - Select your MySQL datasource
   - Click Import

## Configuration

- **Datasource**: The dashboards use MySQL as the datasource
- **Time Range**: Use the time picker in the top-right to select your analysis period
- **Variables**: If you have multiple projects, you can add project filters via dashboard variables

## Notes

- All "Copilot" references have been updated to "GitHub Copilot" for clarity
- The Editor Distribution panel is configured to show percentages in tooltips and legend
- Acceptance rate cards use color thresholds to quickly identify healthy vs. low adoption
- The Engaged Users % card shows trend comparison to help track adoption over time
