# opencloud-skill-cli

An OpenCloud CLI application that can be used by local AI assistants and agents via skills. It allows AI to interact with OpenCloud resources and perform various operations.

### Features

- Seamless integration with local AI assistants and agents via skills.
- Efficient agent token usage, especially compared to tools from MCP servers.
- Fetching and managing OpenCloud resources.
- Authentication with OpenCloud using OAuth2.
- Single binary application for easy installation and usage.

### Requirements

- Go 1.24 or later
- Claude Code, GitHub Copilot, or any other local coding agent.

## Getting started

### Install CLI

```sh
go install github.com/JammingBen/opencloud-skill-cli
oc-cli login
```

### Install Skill

```sh
oc-cli install-skills
```
