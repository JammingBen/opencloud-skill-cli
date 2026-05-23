# motionPhoto

Motion Photo metadata. A Motion Photo is a still image with a short video clip appended to
the end of the file. The presence of this facet on a driveItem indicates that the item is
a Motion Photo; absence indicates it is not.

Based on the Google Motion Photo format v1.0 specification:
https://developer.android.com/media/platform/motion-photo-format


**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `version` | integer (int32) | No | The file format version of the Motion Photo. Currently always 1. Read-only. |
| `presentationTimestampUs` | integer (int64) | No | Presentation timestamp in microseconds of the video frame that corresponds to the still
image. A value of -1 indicates unspecified. If absent, readers should use a timestamp
near the middle of the video track. Read-only.
 |
| `videoSize` | integer (int64) | No | Size in bytes of the embedded video portion of the file. The video is appended at the
end of the file, so clients can fetch it with a Range request:
`Range: bytes=<fileSize - videoSize>-`. Read-only.
 |

