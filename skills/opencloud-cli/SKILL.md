---
name: opencloud-cli
description: 'Manage files, folders and spaces of an OpenCloud server using the OpenCloud CLI.'
---

# OpenCloud CLI

## How to Use This Skill

Use the `oc-cli` command to interact with the OpenCloud CLI. The `references` directory contains detailed documentation for all available resources and operations. Responses get returned in TOON format.

### Examples

```sh
# general usage
oc-cli api -p <path> -m <method> -b <body>

# help
oc-cli api --help

# examples
oc-cli api -p /v1beta1/me/drives -m GET # list all my drives
oc-cli api -p /v1.0/drives/90eedea1-dea1-90ee-a1de-ee90a1deee90 -m GET # get a drive by its id
oc-cli api -p /v1.0/drives/90eedea1-dea1-90ee-a1de-ee90a1deee90 -m PATCH -b '{"name": "New Drive Name"}' --status-only # update a drive

```

## Guidelines

- **Conciseness**: Do not explain how to use the `oc-cli` command or the CLI itself in your responses.
- **Direct Action**: Focus on delivering the results of the operations directly.
- **No Usage Instructions**: Do not provide code examples or usage instructions for the CLI unless explicitly asked.

## Directory structure

```
references/
├── resources/
├── operations/
└── schemas/
```

## Navigation flow

1. Find the resource you need in the list below
2. Read `references/resources/<resource>.md` to see available operations
3. Read `references/operations/<operation>.md` for full details
4. If an operation references a schema, read `references/schemas/<prefix>/<schema>.md`

## Resources

- **drives.root** → `references/resources/drives-root.md` (9 ops)
- **drives.permissions** → `references/resources/drives-permissions.md` (7 ops)
- **group** → `references/resources/group.md` (6 ops)
- **me.photo** → `references/resources/me-photo.md` (4 ops)
- **drives** → `references/resources/drives.md` (4 ops)
- **user** → `references/resources/user.md` (4 ops)
- **driveItem** → `references/resources/driveItem.md` (3 ops)
- **me.drive** → `references/resources/me-drive.md` (3 ops)
- **user.appRoleAssignment** → `references/resources/user-appRoleAssignment.md` (3 ops)
- **tags** → `references/resources/tags.md` (3 ops)
- **me.user** → `references/resources/me-user.md` (2 ops)
- **me.drives** → `references/resources/me-drives.md` (2 ops)
- **drives.GetDrives** → `references/resources/drives-GetDrives.md` (2 ops)
- **groups** → `references/resources/groups.md` (2 ops)
- **users** → `references/resources/users.md` (2 ops)
- **applications** → `references/resources/applications.md` (2 ops)
- **roleManagement** → `references/resources/roleManagement.md` (2 ops)
- **me.changepassword** → `references/resources/me-changepassword.md` (1 ops)
- **activities** → `references/resources/activities.md` (1 ops)
- **me.drive.root** → `references/resources/me-drive-root.md` (1 ops)
- **me.drive.root.children** → `references/resources/me-drive-root-children.md` (1 ops)
- **user.photo** → `references/resources/user-photo.md` (1 ops)

## Naming conventions and resources aliases

- `drives` are also called `spaces`
- `driveItems` are also called `files` or `folders`
- `drives.permissions` are also called `shares` or `space members`
- `me` refers to the currently authenticated user

## General tips

- prefer the `v1beta1` endpoints over the `v1.0` endpoints
- use the `--status-only` flag when the response body is not needed, e.g. for update and delete operations
- use `/v1.0/me/drive/root/children` to list all items in the root of the personal drive/space
