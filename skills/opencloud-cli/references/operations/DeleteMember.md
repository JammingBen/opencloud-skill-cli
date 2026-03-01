# DELETE /v1.0/groups/{group-id}/members/{directory-object-id}/$ref

**Resource:** [group](../resources/group.md)
**Delete member from a group**
**Operation ID:** `DeleteMember`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `group-id` | path | string | Yes | key: id of group |
| `directory-object-id` | path | string | Yes | key: id of group member to remove |
| `If-Match` | header | string | No | ETag |

## Responses

| Status | Description |
|--------|-------------|
| 204 | Success |
| default | (reference) |

