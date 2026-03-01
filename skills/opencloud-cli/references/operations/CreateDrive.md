# POST /v1.0/drives

**Resource:** [drives](../resources/drives.md)
**Create a new drive of a specific type**
**Operation ID:** `CreateDrive`

## Request Body

New space property values

**Required:** Yes

**Content Types:** `application/json`

**Schema:** [drive](../schemas/drive/drive.md)

## Responses

| Status | Description |
|--------|-------------|
| 201 | Created |
| default | (reference) |

**Success Response Schema:**

[drive](../schemas/drive/drive.md)

