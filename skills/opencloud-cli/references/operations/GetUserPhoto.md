# GET /v1.0/users/{user-id}/photo/$value

**Resource:** [user.photo](../resources/user-photo.md)
**Get the photo of a user**
**Operation ID:** `GetUserPhoto`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `user-id` | path | string | Yes | key: id or name of user |

## Responses

| Status | Description |
|--------|-------------|
| 200 | Retrieved photo |
| default | (reference) |

