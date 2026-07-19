# POST /v1beta1/drives/{drive-id}/items/{item-id}/children

**Resource:** [driveItem](../resources/driveItem.md)
**Create a new DriveItem under a parent item**
**Operation ID:** `CreateChildDriveItem`

Create a new folder or DriveItem in a Drive with the specified parent item. The parent must exist and be a folder.

Modeled on the MS Graph create driveItem endpoint
(https://learn.microsoft.com/en-us/graph/api/driveitem-post-children).
Identical request and response shape to the [drive-root variant](#/drives.root/CreateDriveItem), just with an explicit parent item id rather than the drive root.

The request body must specify exactly one of `folder` (set to `{}` to create a folder) or `file` (to create a file item). Requests with none of these, or with both, return 400. The `@libre.graph.conflictBehavior` query parameter controls what happens if a child with the same name already exists.

This endpoint also accepts the MS Graph colon-syntax URL form:

    POST /v1beta1/drives/{drive-id}/items/{item-id}:/{path}:/children

OpenAPI cannot express the colon-delimited path segment, so this URL form is not represented as a separate operation in this specification. The server still accepts it, resolves `:/{path}:` as the parent of the new item (relative to `item-id`), and applies `@libre.graph.missingParentsBehavior` to decide whether to create missing intermediate folders.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |
| `item-id` | path | string | Yes | key: id of item |
| `@libre.graph.conflictBehavior` | query | enum: fail, replace | No | Controls what happens when a child with the same name already exists. `fail` (default) returns 409; `replace` overwrites the existing item. MS Graph's `rename` value is not supported.
 |
| `@libre.graph.missingParentsBehavior` | query | enum: fail, create | No | Controls what happens when a colon-syntax URL refers to a path whose intermediate folders don't all exist yet. `fail` (default) returns 404; `create` creates the missing intermediate folders before creating the final item. Only meaningful for colon-syntax URLs; ignored otherwise.
 |

## Request Body

In the request body, provide a JSON object describing the new driveItem. Must specify exactly one of `folder` or `file`.

**Required:** Yes

**Content Types:** `application/json`

**Schema:** [driveItem](../schemas/driveItem/driveItem.md)

## Responses

| Status | Description |
|--------|-------------|
| 200 | The created DriveItem. |
| 400 | (reference) |
| 403 | (reference) |
| 404 | (reference) |
| 409 | (reference) |
| default | (reference) |

**Success Response Schema:**

[driveItem](../schemas/driveItem/driveItem.md)

