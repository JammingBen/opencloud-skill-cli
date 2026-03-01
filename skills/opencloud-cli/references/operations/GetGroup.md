# GET /v1.0/groups/{group-id}

**Resource:** [group](../resources/group.md)
**Get entity from groups by key**
**Operation ID:** `GetGroup`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `group-id` | path | string | Yes | key: id or name of group |
| `$select` | query | string[] | No | Select properties to be returned |
| `$expand` | query | string[] | No | Expand related entities |

## Responses

| Status | Description |
|--------|-------------|
| 200 | Retrieved entity |
| default | (reference) |

**Success Response Schema:**

[group](../schemas/group/group.md)

