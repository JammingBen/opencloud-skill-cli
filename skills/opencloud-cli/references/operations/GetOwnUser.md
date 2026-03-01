# GET /v1.0/me

**Resource:** [me.user](../resources/me-user.md)
**Get current user**
**Operation ID:** `GetOwnUser`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `$expand` | query | string[] | No | Expand related entities |

## Responses

| Status | Description |
|--------|-------------|
| 200 | Retrieved entity |
| default | (reference) |

**Success Response Schema:**

[user](../schemas/user/user.md)

