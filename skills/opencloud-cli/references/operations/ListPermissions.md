# GET /v1beta1/drives/{drive-id}/items/{item-id}/permissions

**Resource:** [drives.permissions](../resources/drives-permissions.md)
**List the effective sharing permissions on a driveItem.**
**Operation ID:** `ListPermissions`

The permissions collection includes potentially sensitive information and may not be available for every caller.

* For the owner of the item, all sharing permissions will be returned. This includes co-owners.
* For a non-owner caller, only the sharing permissions that apply to the caller are returned.
* Sharing permission properties that contain secrets (e.g. `webUrl`) are only returned for callers that are able to create the sharing permission.

All permission objects have an `id`. A permission representing
* a link has the `link` facet filled with details.
* a share has the `roles` property set and the `grantedToV2` property filled with the grant recipient details.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |
| `item-id` | path | string | Yes | key: id of item |

## Responses

| Status | Description |
|--------|-------------|
| 200 | Retrieved resource |
| default | (reference) |

**Success Response Schema** (inline):

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `@libre.graph.permissions.roles.allowedValues` | unifiedRoleDefinition[] | No | A list of role definitions that can be chosen for the resource. |
| `@libre.graph.permissions.actions.allowedValues` | string[] | No | A list of actions that can be chosen for a custom role.

Following the CS3 API we can represent the CS3 permissions by mapping them to driveItem properties or relations like this:
| [CS3 ResourcePermission](https://cs3org.github.io/cs3apis/#cs3.storage.provider.v1beta1.ResourcePermissions) | action | comment |
| ------------------------------------------------------------------------------------------------------------ | ------ | ------- |
| `stat` | `libre.graph/driveItem/basic/read` | `basic` because it does not include versions or trashed items |
| `get_quota` | `libre.graph/driveItem/quota/read` | read only the `quota` property |
| `get_path` | `libre.graph/driveItem/path/read` | read only the `path` property |
| `move` | `libre.graph/driveItem/path/update` | allows updating the `path` property of a CS3 resource |
| `delete` | `libre.graph/driveItem/standard/delete` | `standard` because deleting is a common update operation |
| `list_container` | `libre.graph/driveItem/children/read` | |
| `create_container` | `libre.graph/driveItem/children/create` | |
| `initiate_file_download` | `libre.graph/driveItem/content/read` | `content` is the property read when initiating a download |
| `initiate_file_upload` | `libre.graph/driveItem/upload/create` | `uploads` are a separate property. postprocessing creates the `content` |
| `add_grant` | `libre.graph/driveItem/permissions/create` | |
| `list_grant` | `libre.graph/driveItem/permissions/read` | |
| `update_grant` | `libre.graph/driveItem/permissions/update` | |
| `remove_grant` | `libre.graph/driveItem/permissions/delete` | |
| `deny_grant` | `libre.graph/driveItem/permissions/deny` | uses a non CRUD action `deny` |
| `list_file_versions` | `libre.graph/driveItem/versions/read` | `versions` is a `driveItemVersion` collection |
| `restore_file_version` | `libre.graph/driveItem/versions/update` | the only `update` action is restore |
| `list_recycle` | `libre.graph/driveItem/deleted/read` | reading a driveItem `deleted` property implies listing |
| `restore_recycle_item` | `libre.graph/driveItem/deleted/update` | the only `update` action is restore |
| `purge_recycle` | `libre.graph/driveItem/deleted/delete` | allows purging deleted `driveItems` |
 |
| `value` | permission[] | No |  |
| `@odata.count` | integer | No | The total number of permissions available, only present if the `count` query parameter is set to true. |

