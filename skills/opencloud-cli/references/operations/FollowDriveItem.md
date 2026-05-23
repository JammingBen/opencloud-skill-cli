# POST /v1.0/me/drive/items/{item-id}/follow

**Resource:** [me.drive](../resources/me-drive.md)
**Follow a DriveItem**
**Operation ID:** `FollowDriveItem`

Follow a DriveItem.

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `item-id` | path | string | Yes | key: id of item |

## Responses

| Status | Description |
|--------|-------------|
| 201 | Created |
| default | (reference) |

**Success Response Schema:**

[driveItem](../schemas/driveItem/driveItem.md)

