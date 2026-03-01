# GET /v1.0/groups/{group-id}/members

**Resource:** [group](../resources/group.md)
**Get a list of the group's direct members**
**Operation ID:** `ListMembers`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `group-id` | path | string | Yes | key: id or name of group |

## Responses

| Status | Description |
|--------|-------------|
| 200 | Retrieved group members |
| default | (reference) |

