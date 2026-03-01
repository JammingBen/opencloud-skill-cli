# DELETE /v1beta1/drives/{drive-id}/items/{item-id}

**Resource:** [driveItem](../resources/driveItem.md)
**Delete a DriveItem.**
**Operation ID:** `DeleteDriveItem`

Delete a DriveItem by using its ID.

Deleting items using this method moves the items to the recycle bin instead of permanently deleting the item.

Mounted shares in the share jail are unmounted. The `@client.synchronize` property of the `driveItem` in the [sharedWithMe](#/me.drive/ListSharedWithMe) endpoint will change to false.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |
| `item-id` | path | string | Yes | key: id of item |

## Responses

| Status | Description |
|--------|-------------|
| 204 | Success |
| default | (reference) |

