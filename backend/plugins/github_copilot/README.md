# GitHub Copilot Plugin

This plugin collects GitHub Copilot usage metrics from GitHub's API and integrates them into Apache DevLake.

## Features

- Collect Copilot usage metrics at organization or enterprise level
- Track acceptance rates, active users, and engaged users
- Analyze usage by programming language and editor/IDE
- Visualize metrics in Grafana dashboards

## Configuration

### Connection Setup

To connect to GitHub Copilot API, you need:

1. A GitHub Personal Access Token with `copilot` scope
2. Organization admin permissions (for organization metrics)
3. Enterprise admin permissions (for enterprise metrics)

### Scope Configuration

Each scope represents a GitHub organization (or enterprise) to collect metrics from.

## Data Models

### Tool Layer Models

- `_tool_github_copilot_connections`: Connection configuration
- `_tool_github_copilot_organizations`: Organization scopes
- `_tool_github_copilot_scope_configs`: Scope configurations
- `_tool_github_copilot_usage`: Daily usage metrics

## API Reference

The plugin uses the following GitHub API endpoints:

- `GET /orgs/{org}/copilot/usage` - Organization-level metrics
- `GET /enterprises/{enterprise}/copilot/usage` - Enterprise-level metrics

See [GitHub Copilot Usage API Documentation](https://docs.github.com/en/rest/copilot/copilot-usage) for details.

## Metrics Collected

- **Total Seats**: Number of Copilot seats assigned
- **Total Active Users**: Users who used Copilot during the period
- **Total Engaged Users**: Users who accepted at least one suggestion
- **Suggestions Count**: Total number of suggestions shown
- **Acceptances Count**: Total number of suggestions accepted
- **Lines Suggested**: Total lines of code suggested
- **Lines Accepted**: Total lines of code accepted
- **Active Chat Users**: Users who used Copilot Chat
- **Active Chat Sessions**: Number of Copilot Chat sessions
- **Language Breakdown**: Usage metrics by programming language
- **Editor Breakdown**: Usage metrics by editor/IDE

## Grafana Dashboards

The plugin includes Grafana dashboards for visualizing:

- Copilot adoption and engagement trends
- Acceptance rate metrics
- Usage distribution by language and editor
- User-level and team-level analytics

## Development

### Building

```bash
make build
```

### Testing

```bash
make unit-test
```

### Running Standalone

```bash
go run github_copilot.go
```

## License

This plugin is licensed under the Apache License 2.0.
