# opencloud-skill-cli

An [OpenCloud](https://opencloud.eu) CLI application that can be used by local AI assistants and agents via skills. It allows AI to interact with OpenCloud resources using the [LibreGraph API](https://github.com/opencloud-eu/libre-graph-api).

### Features

- Seamless integration with local AI assistants and agents via skills.
- Efficient agent token usage thanks to TOON and skill reference data.
- Fetching and managing OpenCloud resources.
- Authentication with OpenCloud using OIDC.
- Single binary application for easy installation and usage.

## Getting started

### Install CLI

```sh
go install github.com/JammingBen/opencloud-skill-cli
```

### Login

You need to login to your OpenCloud server to use the CLI or the skill. You can do this using the `login` command:

```sh
oc-cli login -s https://your-opencloud-server.com
```

When working with a local server that uses self-signed certificates, you can use the `--insecure` flag to skip TLS verification:

```sh
oc-cli login -s https://host.docker.internal:9200 --insecure
```

### Install Skill

Install the skill so your local AI assistant or agent can use it to interact with OpenCloud. You can use the provided `install-skill` command, or download the skill manually from the [Github Releases page](https://github.com/JammingBen/opencloud-skill-cli/releases) (see the artifacts of a release).

```sh
oc-cli install-skill --agent=claude-code
oc-cli install-skill --agent=github-copilot
oc-cli install-skill --agent=codex
oc-cli install-skill --agent=open-code
oc-cli install-skill --agent=gemini-cli
```

## Development

```sh
# install Go dependencies
make tidy

# login to OpenCloud server
make login # defaults to https://host.docker.internal:9200

# OR:
make login SERVER_URL=https://cloud.opencloud.eu # specify own URL

# run cli
go run cmd/oc-cli/*.go --help
```

### Generating skill reference data

Skill reference data can be generated using [openapi-to-skills](https://github.com/neutree-ai/openapi-to-skills/tree/main). This will generate markdown files for all resources, operations and schemas defined in the [OpenCloud OpenAPI spec](https://github.com/opencloud-eu/libre-graph-api/blob/main/api/openapi-spec/v1.0.yaml). The generated files will be moved to `skills/opencloud-cli/references`.

Make sure you have `npx` installed.

```sh
make generate-skill-references
```

## Why not use an MCP server?

The OpenCloud CLI is designed to offer a simple and efficient way for local AI assistants and agents to interact with OpenCloud resources. It does not require the overhead of running an MCP server, and can be easily installed and used on any machine.

Additionally, skills are usually more efficient than MCP servers in terms of token usage because they don't need to retrieve tool definitions initially. The CLI's skill reference structure as well as the TOON format used for responses allow for efficient retrieval of necessary information.
