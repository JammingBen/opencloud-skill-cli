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

**Schema** (inline):

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `storageLocation` | string | No | the path where the file should be created in the users personal space |

## Responses

| Status | Description |
|--------|-------------|
| 202 | success |
| default | (reference) |

