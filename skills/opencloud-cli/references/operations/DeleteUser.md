# DELETE /v1.0/users/{user-id}

**Resource:** [user](../resources/user.md)
**Delete entity from users**
**Operation ID:** `DeleteUser`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `user-id` | path | string | Yes | key: id or name of user |
| `If-Match` | header | string | No | ETag |

## Responses

| Status | Description |
|--------|-------------|
| 204 | Success |
| default | (reference) |

