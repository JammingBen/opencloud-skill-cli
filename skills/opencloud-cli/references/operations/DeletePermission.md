# DELETE /v1beta1/drives/{drive-id}/items/{item-id}/permissions/{perm-id}

**Resource:** [drives.permissions](../resources/drives-permissions.md)
**Remove access to a DriveItem**
**Operation ID:** `DeletePermission`

Remove access to a DriveItem.

Only sharing permissions that are not inherited can be deleted. The `inheritedFrom` property must be `null`.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |
| `item-id` | path | string | Yes | key: id of item |
| `perm-id` | path | string | Yes | key: id of permission |

## Responses

| Status | Description |
|--------|-------------|
| 204 | Success |
| default | (reference) |

