# POST /v1beta1/drives/{drive-id}/root/permissions/{perm-id}/setPassword

**Resource:** [drives.root](../resources/drives-root.md)
**Set sharing link password for the root item of a drive**
**Operation ID:** `SetPermissionPasswordSpaceRoot`

Set the password of a sharing permission.

Only the `password` property can be modified this way.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |
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

