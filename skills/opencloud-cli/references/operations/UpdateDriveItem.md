# PATCH /v1beta1/drives/{drive-id}/items/{item-id}

**Resource:** [driveItem](../resources/driveItem.md)
**Update a DriveItem.**
**Operation ID:** `UpdateDriveItem`

Update a DriveItem.

The request body must include a JSON object with the properties to update.
Only the properties that are provided will be updated.

Currently it supports updating the following properties:

* `@UI.Hidden` - Hides the item from the UI.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |
| `item-id` | path | string | Yes | key: id of item |

## Request Body

DriveItem properties to update

**Required:** Yes

**Content Types:** `application/json`

**Schema:** [driveItem](../schemas/driveItem/driveItem.md)

## Responses

| Status | Description |
|--------|-------------|
| 200 | Success |
| default | (reference) |

**Success Response Schema:**

[driveItem](../schemas/driveItem/driveItem.md)

