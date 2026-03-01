# GET /v1.0/users/{user-id}

**Resource:** [user](../resources/user.md)
**Get entity from users by key**
**Operation ID:** `GetUser`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `user-id` | path | string | Yes | key: id or name of user |
| `$select` | query | string[] | No | Select properties to be returned |
| `$expand` | query | string[] | No | Expand related entities |

## Responses

| Status | Description |
|--------|-------------|
| 200 | Retrieved entity |
| default | (reference) |

**Success Response Schema:**

[user](../schemas/user/user.md)

