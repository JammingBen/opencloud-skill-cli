# POST /v1.0/groups/{group-id}/members/$ref

**Resource:** [group](../resources/group.md)
**Add a member to a group**
**Operation ID:** `AddMember`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `group-id` | path | string | Yes | key: id of group |

## Request Body

Object to be added as member

**Required:** Yes

**Content Types:** `application/json`

## Responses

| Status | Description |
|--------|-------------|
| 204 | Success |
| default | (reference) |

