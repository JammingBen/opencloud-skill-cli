# drives.root

## Operations

| Method | Path | Summary | Details |
|--------|------|---------|----------|
| POST | `/v1beta1/drives/{drive-id}/root/children` | Create a drive item | [View](../operations/CreateDriveItem.md) |
| POST | `/v1beta1/drives/{drive-id}/root/createLink` | Create a sharing link for the root item of a Drive | [View](../operations/CreateLinkSpaceRoot.md) |
| POST | `/v1beta1/drives/{drive-id}/root/invite` | Send a sharing invitation | [View](../operations/InviteSpaceRoot.md) |
| GET | `/v1beta1/drives/{drive-id}/root/permissions` | List the effective permissions on the root item of a drive. | [View](../operations/ListPermissionsSpaceRoot.md) |
| GET | `/v1beta1/drives/{drive-id}/root/permissions/{perm-id}` | Get a single sharing permission for the root item of a drive | [View](../operations/GetPermissionSpaceRoot.md) |
| DELETE | `/v1beta1/drives/{drive-id}/root/permissions/{perm-id}` | Remove access to a Drive | [View](../operations/DeletePermissionSpaceRoot.md) |
| PATCH | `/v1beta1/drives/{drive-id}/root/permissions/{perm-id}` | Update sharing permission | [View](../operations/UpdatePermissionSpaceRoot.md) |
| POST | `/v1beta1/drives/{drive-id}/root/permissions/{perm-id}/setPassword` | Set sharing link password for the root item of a drive | [View](../operations/SetPermissionPasswordSpaceRoot.md) |
