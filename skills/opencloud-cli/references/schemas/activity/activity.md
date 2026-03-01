# activity

Represents activity.

**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | Yes | Activity ID. |
| `times` | object | Yes |  |
| `template` | object | Yes |  |

## Nested Fields

### `times`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `recordedTime` | string (date-time) | Yes | Timestamp of the activity. |

### `template`

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `message` | string | Yes | Activity description. |
| `variables` | object | No | Activity description variables. |

