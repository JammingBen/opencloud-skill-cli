# driveItem

Represents a resource inside a drive. Read-only.

**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | No | Read-only. |
| `createdBy` | [identitySet](identitySet.md) | No |  |
| `createdDateTime` | string (date-time) | No | Date and time of item creation. Read-only. |
| `description` | string | No | Provides a user-visible description of the item. Optional. |
| `eTag` | string | No | ETag for the item. Read-only. |
| `lastModifiedBy` | [identitySet](identitySet.md) | No |  |
| `lastModifiedDateTime` | string (date-time) | No | Date and time the item was last modified. Read-only. |
| `name` | string | No | The name of the item. Read-write. |
| `parentReference` | [itemReference](itemReference.md) | No |  |
| `webUrl` | string | No | URL that displays the resource in the browser. Read-only. |
| `content` | string (base64url) | No | The content stream, if the item represents a file. |
| `cTag` | string | No | An eTag for the content of the item. This eTag is not changed if only the metadata is changed. Note This property is not returned if the item is a folder. Read-only. |
| `deleted` | [deleted](deleted.md) | No |  |
| `file` | [openGraphFile](openGraphFile.md) | No |  |
| `fileSystemInfo` | [fileSystemInfo](fileSystemInfo.md) | No |  |
| `folder` | [folder](folder.md) | No |  |
| `image` | [image](image.md) | No |  |
| `photo` | [photo](photo.md) | No |  |
| `location` | [geoCoordinates](geoCoordinates.md) | No |  |
| `thumbnails` | thumbnailSet[] | No | Collection containing ThumbnailSet objects associated with the item. Read-only. Nullable. |
| `root` | [root](root.md) | No |  |
| `trash` | [trash](trash.md) | No |  |
| `specialFolder` | [specialFolder](specialFolder.md) | No |  |
| `remoteItem` | [remoteItem](remoteItem.md) | No |  |
| `size` | integer (int64) | No | Size of the item in bytes. Read-only. |
| `webDavUrl` | string | No | WebDAV compatible URL for the item. Read-only. |
| `children` | driveItem[] | No | Collection containing Item objects for the immediate children of Item. Only items representing folders have children. Read-only. Nullable. |
| `permissions` | permission[] | No | The set of permissions for the item. Read-only. Nullable. |
| `audio` | [audio](audio.md) | No |  |
| `video` | [video](video.md) | No |  |
| `@client.synchronize` | boolean | No | Indicates if the item is synchronized with the underlying storage provider. Read-only. |
| `@UI.Hidden` | boolean | No | Properties or facets (see UI.Facet) annotated with this term will not be rendered if the annotation evaluates to true. Users can set this to hide permissions. |

