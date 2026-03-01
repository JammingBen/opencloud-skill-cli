# GET /v1beta1/drives/{drive-id}/items/{item-id}/permissions/{perm-id}

**Resource:** [drives.permissions](../resources/drives-permissions.md)
**Get sharing permission for a file or folder**
**Operation ID:** `GetPermission`

Return the effective sharing permission for a particular permission resource.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |
| `item-id` | path | string | Yes | key: id of item |
| `perm-id` | path | string | Yes | key: id of permission |

## Responses

| Status | Description |
|--------|-------------|
| 200 | Retrieved resource |
| default | (reference) |

**Success Response Schema:**

[permission](../schemas/permission/permission.md)

