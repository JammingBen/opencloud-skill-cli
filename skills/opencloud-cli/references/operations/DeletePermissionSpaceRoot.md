# DELETE /v1beta1/drives/{drive-id}/root/permissions/{perm-id}

**Resource:** [drives.root](../resources/drives-root.md)
**Remove access to a Drive**
**Operation ID:** `DeletePermissionSpaceRoot`

Remove access to the root item of a drive.

Only sharing permissions that are not inherited can be deleted. The `inheritedFrom` property must be `null`.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |
| `perm-id` | path | string | Yes | key: id of permission |

## Responses

| Status | Description |
|--------|-------------|
| 204 | Success |
| default | (reference) |

