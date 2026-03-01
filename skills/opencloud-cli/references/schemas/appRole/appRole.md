# appRole

**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `allowedMemberTypes` | string[] | No | Specifies whether this app role can be assigned to users and groups (by setting to ['User']), to other application's (by setting to ['Application'], or both (by setting to ['User', 'Application']). App roles supporting assignment to other applications' service principals are also known as application permissions. The 'Application' value is only supported for app roles defined on application entities. |
| `description` | string | No | The description for the app role. This is displayed when the app role is being assigned and, if the app role functions as an application permission, during  consent experiences. |
| `displayName` | string | No | Display name for the permission that appears in the app role assignment and consent experiences. |
| `id` | string (uuid) | Yes | Unique role identifier inside the appRoles collection. When creating a new app role, a new GUID identifier must be provided. |

