# audio

The Audio resource groups audio-related properties on an item into a single structure.

If a DriveItem has a non-null audio facet, the item represents an audio file. The properties of the Audio resource are populated by extracting metadata from the file.


**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `album` | string | No | The title of the album for this audio file. |
| `albumArtist` | string | No | The artist named on the album for the audio file. |
| `artist` | string | No | The performing artist for the audio file. |
| `bitrate` | integer (int64) | No | Bitrate expressed in kbps. |
| `composers` | string | No | The name of the composer of the audio file. |
| `copyright` | string | No | Copyright information for the audio file. |
| `disc` | integer (int16) | No | The number of the disc this audio file came from. |
| `discCount` | integer (int16) | No | The total number of discs in this album. |
| `duration` | integer (int64) | No | Duration of the audio file, expressed in milliseconds |
| `genre` | string | No | The genre of this audio file. |
| `hasDrm` | boolean | No | Indicates if the file is protected with digital rights management. |
| `isVariableBitrate` | boolean | No | Indicates if the file is encoded with a variable bitrate. |
| `title` | string | No | The title of the audio file. |
| `track` | integer (int32) | No | The number of the track on the original disc for this audio file. |
| `trackCount` | integer (int32) | No | The total number of tracks on the original disc for this audio file. |
| `year` | integer (int32) | No | The year the audio file was recorded. |

