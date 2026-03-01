# POST /v1beta1/drives/{drive-id}/items/{item-id}/permissions/{perm-id}/setPassword

**Resource:** [drives.permissions](../resources/drives-permissions.md)
**Set sharing link password**
**Operation ID:** `SetPermissionPassword`

Set the password of a sharing permission.

Only the `password` property can be modified this way.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |
| `item-id` | path | string | Yes | key: id of item |
| `perm-id` | path | string | Yes | key: id of permission |

## Request Body

New password value

**Required:** Yes

**Content Types:** `application/json`

**Schema:** [sharingLinkPassword](../schemas/sharingLinkPassword/sharingLinkPassword.md)

## Responses

| Status | Description |
|--------|-------------|
| 200 | Updated permission |
| default | (reference) |

**Success Response Schema:**

[permission](../schemas/permission/permission.md)

