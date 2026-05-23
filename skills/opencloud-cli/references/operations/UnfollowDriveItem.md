# DELETE /v1.0/me/drive/following/{item-id}

**Resource:** [me.drive](../resources/me-drive.md)
**Unfollow a DriveItem**
**Operation ID:** `UnfollowDriveItem`

Unfollow a DriveItem.

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `item-id` | path | string | Yes | key: id of item |

## Responses

| Status | Description |
|--------|-------------|
| 204 | No Content |
| default | (reference) |

