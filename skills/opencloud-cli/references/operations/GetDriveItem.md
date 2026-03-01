# GET /v1beta1/drives/{drive-id}/items/{item-id}

**Resource:** [driveItem](../resources/driveItem.md)
**Get a DriveItem.**
**Operation ID:** `GetDriveItem`

Get a DriveItem by using its ID.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |
| `item-id` | path | string | Yes | key: id of item |

## Responses

| Status | Description |
|--------|-------------|
| 200 | Retrieved driveItem |
| default | (reference) |

**Success Response Schema:**

[driveItem](../schemas/driveItem/driveItem.md)

