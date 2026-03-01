# POST /v1beta1/drives/{drive-id}/root/createLink

**Resource:** [drives.root](../resources/drives-root.md)
**Create a sharing link for the root item of a Drive**
**Operation ID:** `CreateLinkSpaceRoot`

You can use the createLink action to share a driveItem via a sharing link.

The response will be a permission object with the link facet containing the created link details.

## Link types

For now, The following values are allowed for the type parameter.

| Value          | Display name      | Description                                                     |
| -------------- | ----------------- | --------------------------------------------------------------- |
| view           | View              | Creates a read-only link to the driveItem.                      |
| upload         | Upload            | Creates a read-write link to the folder driveItem.              |
| edit           | Edit              | Creates a read-write link to the driveItem.                     |
| createOnly     | File Drop         | Creates an upload-only link to the folder driveItem.            |
| blocksDownload | Secure View       | Creates a read-only link that blocks download to the driveItem. |


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |

## Request Body

In the request body, provide a JSON object with the following parameters.

**Content Types:** `application/json`

**Schema:** [driveItemCreateLink](../schemas/driveItemCreateLink/driveItemCreateLink.md)

## Responses

| Status | Description |
|--------|-------------|
| 200 | Response |
| 207 | Partial success response TODO |
| default | (reference) |

**Success Response Schema:**

[permission](../schemas/permission/permission.md)

