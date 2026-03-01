# sharingLink

The `SharingLink` resource groups link-related data items into a single structure.

If a `permission` resource has a non-null `sharingLink` facet, the permission represents a sharing link (as opposed to permissions granted to a person or group).


**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | [sharingLinkType](sharingLinkType.md) | No |  |
| `preventsDownload` | boolean | No | If `true` then the user can only use this link to view the item on the web, and cannot use it to download the contents of the item. |
| `webUrl` | string | No | A URL that opens the item in the browser on the website. |
| `@libre.graph.displayName` | string | No | Provides a user-visible display name of the link. Optional. Libregraph only. |
| `@libre.graph.quickLink` | boolean | No | The quicklink property can be assigned to only one link per resource. A quicklink can be used in the clients to provide a one-click copy to clipboard action. Optional. Libregraph only. |

