# GET /v1beta1/me/drive/sharedByMe

**Resource:** [me.drive](../resources/me-drive.md)
**Get a list of driveItem objects shared by the current user.**
**Operation ID:** `ListSharedByMe`

The `driveItems` returned from the `sharedByMe` method always include the `permissions` relation that indicates they are shared items.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `$expand` | query | string[] | No | Expand related entities |

## Responses

| Status | Description |
|--------|-------------|
| 200 | OK |
| default | (reference) |

