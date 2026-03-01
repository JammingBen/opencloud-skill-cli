# driveRecipient

Represents a person, group, or other recipient to share a drive item with using the invite action.

When using invite to add permissions, the `driveRecipient` object would specify the `email`, `alias`,
or `objectId` of the recipient. Only one of these values is required; multiple values are not accepted.


**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `objectId` | string | No | The unique identifier for the recipient in the directory. |
| `@libre.graph.recipient.type` | string | No | When the recipient is referenced by objectId this annotation is used to differentiate `user` and `group` recipients. |

