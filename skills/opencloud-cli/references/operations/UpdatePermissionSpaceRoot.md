# PATCH /v1beta1/drives/{drive-id}/root/permissions/{perm-id}

**Resource:** [drives.root](../resources/drives-root.md)
**Update sharing permission**
**Operation ID:** `UpdatePermissionSpaceRoot`

Update the properties of a sharing permission by patching the permission resource.

Only the `roles`, `expirationDateTime` and `password` properties can be modified this way.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |
| `perm-id` | path | string | Yes | key: id of permission |

## Request Body

New property values

**Required:** Yes

**Content Types:** `application/json`

**Schema:** [permission](../schemas/permission/permission.md)

## Responses

| Status | Description |
|--------|-------------|
| 200 | Updated permission |
| default | (reference) |

**Success Response Schema:**

[permission](../schemas/permission/permission.md)

