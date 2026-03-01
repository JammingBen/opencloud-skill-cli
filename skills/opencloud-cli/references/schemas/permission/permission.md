# permission

The Permission resource provides information about a sharing permission granted for a DriveItem resource.

### Remarks

The Permission resource uses *facets* to provide information about the kind of permission represented by the resource.

Permissions with a `link` facet represent sharing links created on the item. Sharing links contain a unique token that provides access to the item for anyone with the link.

Permissions with a `invitation` facet represent permissions added by inviting specific users or groups to have access to the file.


**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | No | The unique identifier of the permission among all permissions on the item. Read-only. |
| `hasPassword` | boolean | No | Indicates whether the password is set for this permission. This property only
appears in the response. Optional. Read-only.
 |
| `expirationDateTime` | string (date-time) | No | An optional expiration date which limits the permission in time. |
| `createdDateTime` | string (date-time) | No | An optional creation date. Libregraph only. |
| `grantedToV2` | [sharePointIdentitySet](sharePointIdentitySet.md) | No |  |
| `link` | [sharingLink](sharingLink.md) | No |  |
| `roles` | string[] | No |  |
| `grantedToIdentities` | identitySet[] | No | For link type permissions, the details of the identity to whom permission was granted. This could be used to grant access to a an external user that can be identified by email, aka guest accounts. |
| `@libre.graph.permissions.actions` | string[] | No | Use this to create a permission with custom actions. |
| `invitation` | [sharingInvitation](sharingInvitation.md) | No |  |

