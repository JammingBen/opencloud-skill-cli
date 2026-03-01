# application

**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | Yes | The unique identifier for the object. 12345678-9abc-def0-1234-56789abcde. The value of the ID property is often, but not exclusively, in the form of a GUID. The value should be treated as an opaque identifier and not based in being a GUID. Null values are not allowed. Read-only. |
| `appRoles` | appRole[] | No | The collection of roles defined for the application. With app role assignments, these roles can be assigned to users, groups, or service principals associated with other applications. Not nullable. |
| `displayName` | string | No | The display name for the application. |

