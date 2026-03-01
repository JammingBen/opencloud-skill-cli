# quota

Optional. Information about the drive's storage space quota. Read-only.

**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `deleted` | integer (int64) | No | Total space consumed by files in the recycle bin, in bytes. Read-only. |
| `remaining` | integer (int64) | No | Total space remaining before reaching the quota limit, in bytes. Read-only. |
| `state` | string | No | Enumeration value that indicates the state of the storage space. Either "normal", "nearing", "critical" or "exceeded". Read-only. |
| `total` | integer (int64) | No | Total allowed storage space, in bytes. Read-only. |
| `used` | integer (int64) | No | Total space used, in bytes. Read-only. |

