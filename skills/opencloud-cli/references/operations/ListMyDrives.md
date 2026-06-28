# GET /v1.0/me/drives

**Resource:** [me.drives](../resources/me-drives.md)
**Get all drives where the current user is a regular member of**
**Operation ID:** `ListMyDrives`

## Responses

| Status | Description |
|--------|-------------|
| 200 | Retrieved spaces |
| default | (reference) |

**Success Response Schema** (inline):

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `value` | drive[] | No |  |
| `@odata.nextLink` | string | No |  |

