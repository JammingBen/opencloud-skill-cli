# photo

The photo resource provides photo and camera properties, for example, EXIF metadata, on a driveItem.


**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `cameraMake` | string | No | Camera manufacturer. Read-only. |
| `cameraModel` | string | No | Camera model. Read-only. |
| `exposureDenominator` | number (double) | No | The denominator for the exposure time fraction from the camera. Read-only. |
| `exposureNumerator` | number (double) | No | The numerator for the exposure time fraction from the camera. Read-only. |
| `fNumber` | number (double) | No | The F-stop value from the camera. Read-only. |
| `focalLength` | number (double) | No | The focal length from the camera. Read-only. |
| `iso` | integer (int32) | No | The ISO value from the camera. Read-only. |
| `orientation` | integer (int16) | No | The orientation value from the camera. Read-only. |
| `takenDateTime` | string (date-time) | No | Represents the date and time the photo was taken. Read-only. |

