# POST /v1beta1/drives/{drive-id}/root/children

**Resource:** [drives.root](../resources/drives-root.md)
**Create a drive item**
**Operation ID:** `CreateDriveItem`

You can use the root childrens endpoint to mount a remoteItem in the share jail. The `@client.synchronize` property of the `driveItem` in the [sharedWithMe](#/me.drive/ListSharedWithMe) endpoint will change to true.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |

## Request Body

In the request body, provide a JSON object with the following parameters. For mounting a share the necessary remoteItem id and permission id can be taken from the [sharedWithMe](#/me.drive/ListSharedWithMe) endpoint.

**Content Types:** `application/json`

**Schema:** [driveItem](../schemas/driveItem/driveItem.md)

## Responses

| Status | Description |
|--------|-------------|
| 200 | Response |
| default | (reference) |

**Success Response Schema:**

[driveItem](../schemas/driveItem/driveItem.md)

