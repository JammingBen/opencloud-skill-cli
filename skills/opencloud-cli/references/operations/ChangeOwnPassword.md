# POST /v1.0/me/changePassword

**Resource:** [me.changepassword](../resources/me-changepassword.md)
**Change your own password**
**Operation ID:** `ChangeOwnPassword`

## Request Body

Password change request

**Required:** Yes

**Content Types:** `application/json`

**Schema** (inline):

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `currentPassword` | string | Yes |  |
| `newPassword` | string | Yes |  |

## Responses

| Status | Description |
|--------|-------------|
| 204 | Success |
| default | (reference) |

