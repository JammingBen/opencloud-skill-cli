# video

The video resource groups video-related data items into a single structure.

If a driveItem has a non-null video facet, the item represents a video file. The properties of the video resource are populated by extracting metadata from the file.


**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `audioBitsPerSample` | integer (int32) | No | Number of audio bits per sample. |
| `audioChannels` | integer (int32) | No | Number of audio channels. |
| `audioFormat` | string | No | Name of the audio format (AAC, MP3, etc.). |
| `audioSamplesPerSecond` | integer (int32) | No | Number of audio samples per second. |
| `bitrate` | integer (int32) | No | Bit rate of the video in bits per second. |
| `duration` | integer (int64) | No | Duration of the file in milliseconds. |
| `fourCC` | string | No | \"Four character code\" name of the video format. |
| `frameRate` | number (double) | No | Frame rate of the video. |
| `height` | integer (int32) | No | Height of the video, in pixels. |
| `width` | integer (int32) | No | Width of the video, in pixels. |

