# GET /v1.0/drives/{drive-id}/root

**Resource:** [drives.root](../resources/drives-root.md)
**Get root from arbitrary space**
**Operation ID:** `GetRoot`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |

## Responses

| Status | Description |
|--------|-------------|
| 200 | Retrieved resource |
| default | (reference) |

**Success Response Schema:**

[driveItem](../schemas/driveItem/driveItem.md)

