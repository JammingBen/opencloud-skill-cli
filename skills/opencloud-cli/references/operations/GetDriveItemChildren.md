# GET /v1.0/drives/{drive-id}/items/{item-id}/children

**Resource:** [driveItem](../resources/driveItem.md)
**List children of a DriveItem**
**Operation ID:** `GetDriveItemChildren`

List the children of the item identified by `item-id` in the drive identified by `drive-id`. The item must exist and be a folder.

Modeled on the MS Graph list driveItem children endpoint
(https://learn.microsoft.com/en-us/graph/api/driveitem-list-children).

This endpoint also accepts the MS Graph colon-syntax URL forms:

    GET /v1.0/drives/{drive-id}/root:/{path}:/children
    GET /v1.0/drives/{drive-id}/items/{item-id}:/{path}:/children

OpenAPI cannot express the colon-delimited path segment, so these URL forms are not represented as separate operations in this specification. The server still accepts them, resolves `:/{path}:` as the parent item, and lists its children.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |
| `item-id` | path | string | Yes | key: id of item |

## Responses

| Status | Description |
|--------|-------------|
| 200 | Retrieved resource list |
| default | (reference) |

**Success Response Schema** (inline):

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `value` | driveItem[] | No |  |
| `@odata.nextLink` | string | No |  |

