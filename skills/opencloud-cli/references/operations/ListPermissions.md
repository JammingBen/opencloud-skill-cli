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

