# driveItemCreateLink

**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `type` | [sharingLinkType](sharingLinkType.md) | No |  |
| `expirationDateTime` | string (date-time) | No | Optional. A String with format of yyyy-MM-ddTHH:mm:ssZ of DateTime indicates the expiration time of the permission. |
| `password` | string | No | Optional.The password of the sharing link that is set by the creator. |
| `displayName` | string | No | Provides a user-visible display name of the link. Optional. Libregraph only. |
| `@libre.graph.quickLink` | boolean | No | The quicklink property can be assigned to only one link per resource. A quicklink can be used in the clients to provide a one-click copy to clipboard action. Optional. Libregraph only. |

