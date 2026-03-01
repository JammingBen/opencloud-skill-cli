# POST /v1beta1/drives/{drive-id}/root/invite

**Resource:** [drives.root](../resources/drives-root.md)
**Send a sharing invitation**
**Operation ID:** `InviteSpaceRoot`

Sends a sharing invitation for the root of a `drive`. A sharing invitation provides permissions to the
recipients and optionally sends them an email with a sharing link.

The response will be a permission object with the grantedToV2 property containing the created grant details.

## Roles property values
For now, roles are only identified by a uuid. There are no hardcoded aliases like `read` or `write` because role actions can be completely customized.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |

## Request Body

In the request body, provide a JSON object with the following parameters. To create a custom role submit a list of actions instead of roles.

**Content Types:** `application/json`

**Schema:** [driveItemInvite](../schemas/driveItemInvite/driveItemInvite.md)

## Responses

| Status | Description |
|--------|-------------|
| 200 | Response |
| 207 | Partial success response TODO |
| 400 | Bad request |
| default | (reference) |

