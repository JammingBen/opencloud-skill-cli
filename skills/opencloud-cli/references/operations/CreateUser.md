# POST /v1.0/users

**Resource:** [users](../resources/users.md)
**Add new entity to users**
**Operation ID:** `CreateUser`

## Request Body

New entity

**Required:** Yes

**Content Types:** `application/json`

**Schema:** [user](../schemas/user/user.md)

## Responses

| Status | Description |
|--------|-------------|
| 201 | Created entity |
| default | (reference) |

**Success Response Schema:**

[user](../schemas/user/user.md)

