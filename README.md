# opencloud-skill-cli

An [OpenCloud](https://opencloud.eu) CLI application that can be used by local AI assistants and agents via skills. It allows AI to interact with OpenCloud resources and perform various operations.

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
oc-cli login
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

Skill reference data can be generated using [openapi-to-skills](https://github.com/neutree-ai/openapi-to-skills/tree/main). This will generate markdown files for all resources, operations and schemas defined in the [OpenAPI spec](https://github.com/opencloud-eu/libre-graph-api/blob/main/api/openapi-spec/v1.0.yaml). The generated files will be moved to `skills/opencloud-cli/references`.

Make sure you have `npx` installed.

```sh
make generate-skill-references
```
