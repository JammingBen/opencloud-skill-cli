# thumbnail

The thumbnail resource type represents a thumbnail for an image, video, document, or any item that has a bitmap representation.


**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `content` | string (base64url) | No | The content stream for the thumbnail. |
| `height` | integer (int32) | No | The height of the thumbnail, in pixels. |
| `sourceItemId` | string | No | The unique identifier of the item that provided the thumbnail. This is only available when a folder thumbnail is requested. |
| `url` | string | No | The URL used to fetch the thumbnail content. |
| `width` | integer (int32) | No | The width of the thumbnail, in pixels. |

