# GET /v1.0/users

**Resource:** [users](../resources/users.md)
**Get entities from users**
**Operation ID:** `ListUsers`

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

