# POST /v1beta1/drives/{drive-id}/root/children

**Resource:** [drives.root](../resources/drives-root.md)
**Create a new DriveItem at the drive root**
**Operation ID:** `CreateDriveItem`

Create a new folder or DriveItem in a Drive with the drive root as the parent.

Modeled on the MS Graph create driveItem endpoint
(https://learn.microsoft.com/en-us/graph/api/driveitem-post-children).

The request body must specify exactly one of `folder` (set to `{}` to create a folder), `file` (to create a file item), or `remoteItem` (to mount a shared item; see [sharedWithMe](#/me.drive/ListSharedWithMe) for obtaining the source `remoteItem.id`). Requests with none of these, or with more than one, return 400. Mounting a share changes the `@client.synchronize` property of the `driveItem` in [sharedWithMe](#/me.drive/ListSharedWithMe) to true.

The `@libre.graph.conflictBehavior` query parameter controls what happens if a child with the same name already exists.

This endpoint also accepts the MS Graph colon-syntax URL form:

    POST /v1beta1/drives/{drive-id}/root:/{path}:/children

OpenAPI cannot express the colon-delimited path segment, so this URL form is not represented as a separate operation in this specification. The server still accepts it, resolves `:/{path}:` as the parent of the new item, and applies `@libre.graph.missingParentsBehavior` to decide whether to create missing intermediate folders.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |
| `@libre.graph.conflictBehavior` | query | enum: fail, replace | No | Controls what happens when a child with the same name already exists. `fail` (default) returns 409; `replace` overwrites the existing item. MS Graph's `rename` value is not supported.
 |
| `@libre.graph.missingParentsBehavior` | query | enum: fail, create | No | Controls what happens when a colon-syntax URL refers to a path whose intermediate folders don't all exist yet. `fail` (default) returns 404; `create` creates the missing intermediate folders before creating the final item. Only meaningful for colon-syntax URLs; ignored otherwise.
 |

## Request Body

In the request body, provide a JSON object describing the new driveItem. Must specify exactly one of `folder`, `file`, or `remoteItem`. For mount-share, see [sharedWithMe](#/me.drive/ListSharedWithMe) for obtaining the source `remoteItem.id` and `permission` id.

**Content Types:** `application/json`

**Schema:** [driveItem](../schemas/driveItem/driveItem.md)

## Responses

| Status | Description |
|--------|-------------|
| 200 | Response |
| default | (reference) |

**Success Response Schema:**

[driveItem](../schemas/driveItem/driveItem.md)

