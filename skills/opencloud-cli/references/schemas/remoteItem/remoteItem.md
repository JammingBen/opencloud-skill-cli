# remoteItem

Remote item data, if the item is shared from a drive other than the one being accessed. Read-only.

**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `createdBy` | [identitySet](identitySet.md) | No |  |
| `createdDateTime` | string (date-time) | No | Date and time of item creation. Read-only. |
| `file` | [openGraphFile](openGraphFile.md) | No |  |
| `fileSystemInfo` | [fileSystemInfo](fileSystemInfo.md) | No |  |
| `folder` | [folder](folder.md) | No |  |
| `driveAlias` | string | No | The drive alias can be used in clients to make the urls user friendly. Example: 'personal/einstein'. This will be used to resolve to the correct driveID. |
| `path` | string | No | The relative path of the item in relation to its drive root. |
| `rootId` | string | No | Unique identifier for the drive root of this item. Read-only. |
| `id` | string | No | Unique identifier for the remote item in its drive. Read-only. |
| `image` | [image](image.md) | No |  |
| `lastModifiedBy` | [identitySet](identitySet.md) | No |  |
| `lastModifiedDateTime` | string (date-time) | No | Date and time the item was last modified. Read-only. |
| `name` | string | No | Optional. Filename of the remote item. Read-only. |
| `eTag` | string | No | ETag for the item. Read-only. |
| `cTag` | string | No | An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only. |
| `parentReference` | [itemReference](itemReference.md) | No |  |
| `permissions` | permission[] | No | The set of permissions for the item. Read-only. Nullable. |
| `size` | integer (int64) | No | Size of the remote item. Read-only. |
| `specialFolder` | [specialFolder](specialFolder.md) | No |  |
| `webDavUrl` | string | No | DAV compatible URL for the item. |
| `webUrl` | string | No | URL that displays the resource in the browser. Read-only. |

