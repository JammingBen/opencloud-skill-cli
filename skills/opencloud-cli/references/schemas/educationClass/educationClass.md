# educationClass

And extension of group representing a class or course

**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | No | Read-only. |
| `description` | string | No | An optional description for the group. Returned by default. |
| `displayName` | string | Yes | The display name for the group. This property is required when a group is created and cannot be cleared during updates. Returned by default. Supports $search and $orderBy. |
| `members` | user[] | No | Users and groups that are members of this group. HTTP Methods: GET (supported for all groups), Nullable. Supports $expand. |
| `members@odata.bind` | string[] | No | A list of member references to the members to be added. Up to 20 members can be added with a single request |
| `classification` | enum: class, course | Yes | Classification of the group, i.e. "class" or "course" |
| `externalId` | string | No | An external unique ID for the class |

