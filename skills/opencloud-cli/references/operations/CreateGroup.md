# POST /v1.0/groups

**Resource:** [groups](../resources/groups.md)
**Add new entity to groups**
**Operation ID:** `CreateGroup`

## Request Body

New entity

**Required:** Yes

**Content Types:** `application/json`

**Schema:** [group](../schemas/group/group.md)

## Responses

| Status | Description |
|--------|-------------|
| 201 | Created entity |
| default | (reference) |

**Success Response Schema:**

[group](../schemas/group/group.md)

