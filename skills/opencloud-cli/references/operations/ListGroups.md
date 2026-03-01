# GET /v1.0/groups

**Resource:** [groups](../resources/groups.md)
**Get entities from groups**
**Operation ID:** `ListGroups`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `$orderby` | query | string[] | No | Order items by property values |
| `$select` | query | string[] | No | Select properties to be returned |
| `$expand` | query | string[] | No | Expand related entities |

## Responses

| Status | Description |
|--------|-------------|
| 200 | Retrieved entities |
| default | (reference) |

