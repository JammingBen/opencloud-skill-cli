# PATCH /v1.0/drives/{drive-id}

**Resource:** [drives](../resources/drives.md)
**Update the drive**
**Operation ID:** `UpdateDrive`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `drive-id` | path | string | Yes | key: id of drive |

## Request Body

New space values

**Required:** Yes

**Content Types:** `application/json`

**Schema:** [driveUpdate](../schemas/driveUpdate/driveUpdate.md)

## Responses

| Status | Description |
|--------|-------------|
| 200 | Success |
| default | (reference) |

**Success Response Schema:**

[drive](../schemas/drive/drive.md)

