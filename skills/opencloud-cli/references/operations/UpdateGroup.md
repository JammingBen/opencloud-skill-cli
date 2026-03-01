# PATCH /v1.0/groups/{group-id}

**Resource:** [group](../resources/group.md)
**Update entity in groups**
**Operation ID:** `UpdateGroup`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `group-id` | path | string | Yes | key: id of group |

## Request Body

New property values

**Required:** Yes

**Content Types:** `application/json`

**Schema:** [group](../schemas/group/group.md)

## Responses

| Status | Description |
|--------|-------------|
| 204 | Success |
| default | (reference) |

