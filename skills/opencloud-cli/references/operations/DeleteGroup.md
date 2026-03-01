# DELETE /v1.0/groups/{group-id}

**Resource:** [group](../resources/group.md)
**Delete entity from groups**
**Operation ID:** `DeleteGroup`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `group-id` | path | string | Yes | key: id of group |
| `If-Match` | header | string | No | ETag |

## Responses

| Status | Description |
|--------|-------------|
| 204 | Success |
| default | (reference) |

