# PATCH /v1.0/me

**Resource:** [me.user](../resources/me-user.md)
**Update the current user**
**Operation ID:** `UpdateOwnUser`

## Request Body

New user values

**Content Types:** `application/json`

**Schema:** [userUpdate](../schemas/userUpdate/userUpdate.md)

## Responses

| Status | Description |
|--------|-------------|
| 200 | Success |
| default | (reference) |

**Success Response Schema:**

[user](../schemas/user/user.md)

