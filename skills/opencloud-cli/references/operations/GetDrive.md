# GET /v1.0/drives/{drive-id}

**Resource:** [drives](../resources/drives.md)
**Get drive by id**
**Operation ID:** `GetDrive`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |

## Responses

| Status | Description |
|--------|-------------|
| 200 | Retrieved drive |
| default | (reference) |

**Success Response Schema:**

[drive](../schemas/drive/drive.md)

