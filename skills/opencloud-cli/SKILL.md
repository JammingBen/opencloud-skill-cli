---
name: opencloud-cli
description: 'Manage files, folders and spaces of an OpenCloud server using the OpenCloud CLI.'
---

# OpenCloud CLI

Manage files, folders and spaces of an OpenCloud server using the OpenCloud CLI (`oc-cli`). Responses get returned in TOON format.

## How to Use This Skill

Use the `oc-cli api` command to interact with the API. The `references` directory contains detailed documentation for all available resources and operations.

### Examples

```sh
# list all my drives (v1beta1)
oc-cli api -p '/v1beta1/me/drives' -m GET

# get a specific drive by id (v1.0)
oc-cli api -p '/v1.0/drives/90eedea1-dea1-90ee-a1de-ee90a1deee90' -m GET

# update a drive name (v1.0) and only return status
oc-cli api -p '/v1.0/drives/90eedea1-dea1-90ee-a1de-ee90a1deee90' -m PATCH -b '{"name": "New Name"}' --status-only

# help
oc-cli api --help
```

## Guidelines & Operational Rules

### Response Style

- **Conciseness**: Do not explain how to use the `oc-cli` command or the CLI itself.
- **Direct Action**: Focus on delivering the results of the operations directly.
- **No Usage Instructions**: Do not provide code examples for the CLI unless explicitly asked.

### Technical Requirements

- **Versioning**: **Always prefer `v1beta1`** endpoints over `v1.0` if both are available.
- **Shell Safety**: Always wrap API paths and body data in single quotes (e.g., `-p '/v1.0/me' -b '{"name": "foo"}'`).
- **Optimization**: Use the `--status-only` flag when the response body isn't needed (e.g., DELETE, or PATCH when only confirmation is required).
- **Recipient Type**: When inviting users or groups by `objectId`, the `@libre.graph.recipient.type` property **must** be set to either `"user"` or `"group"`.
- **Link Passwords**: If creating links fails with an error stating that password protection is enforced, provide a password via the `password` field (8+ characters, including uppercase, lowercase, numbers, and special characters).

### Workflow Patterns & Tips

- **Search-Then-Act**: If an operation requires an `objectId` (e.g. for a user or group when creating a space member or share) that you don't have, search for it first:
  - `-p '/v1.0/users' -q '$search="alice"'`
  - `-p '/v1.0/groups' -q '$search="marketing"'`
- **Root Personal Drive**: Use `'/v1.0/me/drive/root/children'` to list all items in the root of the personal drive/space.
- **Drives/Spaces**: "Drives" are also referred to as "Spaces". They represent a container for files and folders.
- **DriveItems/Files/Folders**: "DriveItems" are also referred to as "Files" or "Folders".
- **Permissions/Shares/Space Members**: "Permissions" are also referred to as "Shares" (on normal files and folders) or "Space Members" (on spaces).
- **Managing Space Members**: When managing members (=permissions) for a project space, you must use the `v1beta1/drives/{driveId}/root/...` endpoints:
  - **List members**: Use the `'/v1beta1/drives/{driveId}/root/permissions'` path.
  - **Add members**: Use the `'/v1beta1/drives/{driveId}/root/invite'` path.
  - **Remove/Update members**: Use the `'/v1beta1/drives/{driveId}/root/permissions/{permissionId}'` path.

## Navigation flow

1.  **Identify Resource**: Find the relevant resource in the [Resources](#resources) section.
2.  **Check Resource Docs**: Read `references/resources/<resource>.md` to see available operations.
3.  **Read Operation Details**: Read `references/operations/<operation>.md` for specific path parameters and request body requirements.
4.  **Review Schemas**: If an operation references a schema, check `references/schemas/<prefix>/<schema>.md` for the full data structure.

## Resources

### Files, Folders & Spaces

- **drives.root** → `references/resources/drives-root.md`
- **drives.permissions** → `references/resources/drives-permissions.md`
- **me.drive** → `references/resources/me-drive.md`
- **me.drive.root** → `references/resources/me-drive-root.md`
- **me.drive.root.children** → `references/resources/me-drive-root-children.md`
- **me.drives** → `references/resources/me-drives.md`
- **drives.GetDrives** → `references/resources/drives-GetDrives.md`
- **drives** → `references/resources/drives.md`
- **driveItem** → `references/resources/driveItem.md`

### Users, Groups & Management

- **user** → `references/resources/user.md`
- **users** → `references/resources/users.md`
- **group** → `references/resources/group.md`
- **groups** → `references/resources/groups.md`
- **me.user** → `references/resources/me-user.md`
- **me.changepassword** → `references/resources/me-changepassword.md`
- **me.photo** → `references/resources/me-photo.md`
- **user.photo** → `references/resources/user-photo.md`
- **user.appRoleAssignment** → `references/resources/user-appRoleAssignment.md`

### Extras & Insights

- **tags** → `references/resources/tags.md`
- **activities** → `references/resources/activities.md`
- **applications** → `references/resources/applications.md`
- **roleManagement** → `references/resources/roleManagement.md`

## Directory Structure

```
references/
├── resources/
├── operations/
└── schemas/
```
