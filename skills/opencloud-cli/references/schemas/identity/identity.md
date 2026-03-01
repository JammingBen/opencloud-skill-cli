# identity

**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `displayName` | string | Yes | The identity's display name. Note that this may not always be available or up to date. For example, if a user changes their display name, the API may show the new value in a future response, but the items associated with the user won't show up as having changed when using delta. |
| `id` | string | No | Unique identifier for the identity. |
| `@libre.graph.userType` | string | No | The type of the identity. This can be either "Member" for regular user, "Guest" for guest users or "Federated" for users imported from a federated instance. Can be used by clients to indicate the type of user. For more details, clients should look up and cache the user at the /users endpoint. |

