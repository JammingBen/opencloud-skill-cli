# invitation

Represents an invitation to a drive item.


**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `invitedUserDisplayName` | string | No | The display name of the user being invited. |
| `invitedUserEmailAddress` | string | No | The email address of the user being invited. Required. |
| `invitedUserMessageInfo` | [invitedUserMessageInfo](invitedUserMessageInfo.md) | No |  |
| `sendInvitationMessage` | boolean | No | Indicates whether an invitation message should be sent to the user. |
| `inviteRedirectUrl` | string | No | The URL to which the user is redirected after accepting the invitation. Required. |
| `inviteRedeemUrl` | string | No | The URL that the user can use to redeem the invitation. Read-only. |
| `status` | string | No | The status of the invitation. Read-only. |
| `invitedUser` | [user](user.md) | No |  |
| `invitedUserType` | string | No | The type of user being invited. |

