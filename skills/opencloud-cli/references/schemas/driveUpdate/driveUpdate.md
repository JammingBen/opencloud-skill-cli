# driveUpdate

The drive represents an update to a space on the storage.

**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | No | The unique identifier for this drive. |
| `createdBy` | [identitySet](identitySet.md) | No |  |
| `createdDateTime` | string (date-time) | No | Date and time of item creation. Read-only. |
| `description` | string | No | Provides a user-visible description of the item. Optional. |
| `eTag` | string | No | ETag for the item. Read-only. |
| `lastModifiedBy` | [identitySet](identitySet.md) | No |  |
| `lastModifiedDateTime` | string (date-time) | No | Date and time the item was last modified. Read-only. |
| `name` | string | No | The name of the item. Read-write. |
| `parentReference` | [itemReference](itemReference.md) | No |  |
| `webUrl` | string | No | URL that displays the resource in the browser. Read-only. |
| `driveType` | string | No | Describes the type of drive represented by this resource. Values are "personal" for users home spaces, "project", "virtual" or "share". Read-only. |
| `driveAlias` | string | No | The drive alias can be used in clients to make the urls user friendly. Example: 'personal/einstein'. This will be used to resolve to the correct driveID. |
| `owner` | [identitySet](identitySet.md) | No |  |
| `quota` | [quota](quota.md) | No |  |
| `items` | driveItem[] | No | All items contained in the drive. Read-only. Nullable. |
| `root` | [driveItem](driveItem.md) | No |  |
| `special` | driveItem[] | No | A collection of special drive resources. |
| `@libre.graph.hasTrashedItems` | boolean | No | Indicates whether the drive has items in the trash. Read-only. |

