# GET /v1beta1/drives/{drive-id}/items/{item-id}/content

**Resource:** [driveItem](../resources/driveItem.md)
**Download the content of a DriveItem**
**Operation ID:** `GetDriveItemContent`

Download the contents of the primary stream (file) of a driveItem. Only
driveItem objects with a `file` facet can be downloaded.

The response is a `302 Found` redirecting to a pre-authenticated download
URL for the file. This is the same URL that is returned via the
`@microsoft.graph.downloadUrl` instance annotation on the driveItem when
requested via `$select`. Choose between the two based on whether you
want to call the redirecting `/content` endpoint directly (for example,
with a client that follows redirects automatically) or you want to
inspect / schedule / prefetch the URL yourself via the annotation.

The pre-authenticated URL is short-lived and does not require an
`Authorization` header.

To download a partial range of bytes, apply the `Range` header to the
redirect target (the pre-authenticated URL), not to the `/content`
request.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |
| `item-id` | path | string | Yes | key: id of item |

## Responses

| Status | Description |
|--------|-------------|
| 302 | Pre-authenticated redirect to the file content. |
| 404 | The driveItem was not found or is not a file. |
| default | (reference) |

