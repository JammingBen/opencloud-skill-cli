# DELETE /v1.0/drives/{drive-id}

**Resource:** [drives](../resources/drives.md)
**Delete a specific space**
**Operation ID:** `DeleteDrive`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |
| `If-Match` | header | string | No | ETag |

## Responses

| Status | Description |
|--------|-------------|
| 204 | Success |
| default | (reference) |

