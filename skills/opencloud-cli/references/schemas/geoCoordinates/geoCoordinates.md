# geoCoordinates

The GeoCoordinates resource provides geographic coordinates and elevation of a location based on metadata contained within the file.
If a DriveItem has a non-null location facet, the item represents a file with a known location associated with it.


**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `altitude` | number (double) | No | The altitude (height), in feet, above sea level for the item. Read-only. |
| `latitude` | number (double) | No | The latitude, in decimal, for the item. Read-only. |
| `longitude` | number (double) | No | The longitude, in decimal, for the item. Read-only. |

