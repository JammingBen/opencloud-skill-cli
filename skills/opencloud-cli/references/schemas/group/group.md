# group

**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | No | Read-only. |
| `description` | string | No | An optional description for the group. Returned by default. |
| `displayName` | string | No | The display name for the group. This property is required when a group is created and cannot be cleared during updates. Returned by default. Supports $search and $orderBy. |
| `groupTypes` | string[] | No | Specifies the group types. In MS Graph a group can have multiple types, so this is an array. In libreGraph the possible group types deviate from the MS Graph. The only group type that we currently support is "ReadOnly", which is set for groups that cannot be modified on the current instance. |
| `members` | user[] | No | Users and groups that are members of this group. HTTP Methods: GET (supported for all groups), Nullable. Supports $expand. |
| `members@odata.bind` | string[] | No | A list of member references to the members to be added. Up to 20 members can be added with a single request |

