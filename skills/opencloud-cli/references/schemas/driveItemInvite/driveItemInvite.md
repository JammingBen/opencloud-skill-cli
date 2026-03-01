# driveItemInvite

**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `recipients` | driveRecipient[] | No | A collection of recipients who will receive access and the sharing invitation. Currently, only internal users or groups are supported. |
| `roles` | string[] | No | Specifies the roles that are to be granted to the recipients of the sharing invitation. |
| `@libre.graph.permissions.actions` | string[] | No | Specifies the actions that are to be granted to the recipients of the sharing invitation, in effect creating a custom role. |
| `expirationDateTime` | string (date-time) | No | Specifies the dateTime after which the permission expires. |

