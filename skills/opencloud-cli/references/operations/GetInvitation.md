# GET /v1.0/invitations/{invitation-id}

**Resource:** [invitations](../resources/invitations.md)
**Get an invitation by key**
**Operation ID:** `GetInvitation`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `invitation-id` | path | string | Yes | key: id of invitation |

## Responses

| Status | Description |
|--------|-------------|
| 200 | Retrieved invitation |

**Success Response Schema:**

[invitation](../schemas/invitation/invitation.md)

