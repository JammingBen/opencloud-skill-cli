# PATCH /v1.0/users/{user-id}

**Resource:** [user](../resources/user.md)
**Update entity in users**
**Operation ID:** `UpdateUser`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `user-id` | path | string | Yes | key: id of user |

## Request Body

New property values

**Required:** Yes

**Content Types:** `application/json`

**Schema:** [userUpdate](../schemas/userUpdate/userUpdate.md)

## Responses

| Status | Description |
|--------|-------------|
| 200 | Success |
| default | (reference) |

**Success Response Schema:**

[user](../schemas/user/user.md)

