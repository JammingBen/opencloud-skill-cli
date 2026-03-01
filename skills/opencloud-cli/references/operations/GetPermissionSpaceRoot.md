# GET /v1beta1/drives/{drive-id}/root/permissions/{perm-id}

**Resource:** [drives.root](../resources/drives-root.md)
**Get a single sharing permission for the root item of a drive**
**Operation ID:** `GetPermissionSpaceRoot`

Return the effective sharing permission for a particular permission resource.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |
| `perm-id` | path | string | Yes | key: id of permission |

## Responses

| Status | Description |
|--------|-------------|
| 200 | Retrieved resource |
| default | (reference) |

**Success Response Schema:**

[permission](../schemas/permission/permission.md)

