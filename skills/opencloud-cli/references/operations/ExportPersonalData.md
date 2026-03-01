# POST /v1.0/users/{user-id}/exportPersonalData

**Resource:** [user](../resources/user.md)
**export personal data of a user**
**Operation ID:** `ExportPersonalData`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `user-id` | path | string | Yes | key: id or name of user |

## Request Body

destination the file should be created at

**Content Types:** `application/json`

## Responses

| Status | Description |
|--------|-------------|
| 202 | success |
| default | (reference) |

