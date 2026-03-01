# unifiedRoleDefinition

A role definition is a collection of permissions in libre graph listing the operations that can be performed
and the resources against which they can performed.


**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `description` | string | No | The description for the unifiedRoleDefinition. |
| `displayName` | string | No | The display name for the unifiedRoleDefinition. Required. Supports $filter (`eq`, `in`). |
| `id` | string | No | The unique identifier for the role definition. Key, not nullable, Read-only. Inherited from entity. Supports $filter (`eq`, `in`). |
| `rolePermissions` | unifiedRolePermission[] | No | List of permissions included in the role. |
| `@libre.graph.weight` | integer | No | When presenting a list of roles the weight can be used to order them in a meaningful way.
Lower weight gets higher precedence. So content with lower weight will come first. If set,
weights should be non-zero, as 0 is interpreted as an unset weight.
 |

