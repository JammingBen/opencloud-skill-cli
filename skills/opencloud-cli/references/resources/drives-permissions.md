# drives.permissions

## Operations

| Method | Path | Summary | Details |
|--------|------|---------|----------|
| POST | `/v1beta1/drives/{drive-id}/items/{item-id}/createLink` | Create a sharing link for a DriveItem | [View](../operations/CreateLink.md) |
| POST | `/v1beta1/drives/{drive-id}/items/{item-id}/invite` | Send a sharing invitation | [View](../operations/Invite.md) |
| GET | `/v1beta1/drives/{drive-id}/items/{item-id}/permissions` | List the effective sharing permissions on a driveItem. | [View](../operations/ListPermissions.md) |
| GET | `/v1beta1/drives/{drive-id}/items/{item-id}/permissions/{perm-id}` | Get sharing permission for a file or folder | [View](../operations/GetPermission.md) |
| DELETE | `/v1beta1/drives/{drive-id}/items/{item-id}/permissions/{perm-id}` | Remove access to a DriveItem | [View](../operations/DeletePermission.md) |
| PATCH | `/v1beta1/drives/{drive-id}/items/{item-id}/permissions/{perm-id}` | Update sharing permission | [View](../operations/UpdatePermission.md) |
| POST | `/v1beta1/drives/{drive-id}/items/{item-id}/permissions/{perm-id}/setPassword` | Set sharing link password | [View](../operations/SetPermissionPassword.md) |
