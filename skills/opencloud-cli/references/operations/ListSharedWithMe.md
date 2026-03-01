# GET /v1beta1/me/drive/sharedWithMe

**Resource:** [me.drive](../resources/me-drive.md)
**Get a list of driveItem objects shared with the owner of a drive.**
**Operation ID:** `ListSharedWithMe`

The `driveItems` returned from the `sharedWithMe` method always include the `remoteItem` facet that indicates they are items from a different drive.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `$expand` | query | string[] | No | Expand related entities |

## Responses

| Status | Description |
|--------|-------------|
| 200 | OK |
| default | (reference) |

