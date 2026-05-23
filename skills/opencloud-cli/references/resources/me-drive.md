# me.drive

## Operations

| Method | Path | Summary | Details |
|--------|------|---------|----------|
| GET | `/v1.0/me/drive` | Get personal space for user | [View](../operations/GetHome.md) |
| GET | `/v1beta1/me/drive/sharedByMe` | Get a list of driveItem objects shared by the current user. | [View](../operations/ListSharedByMe.md) |
| POST | `/v1.0/me/drive/items/{item-id}/follow` | Follow a DriveItem | [View](../operations/FollowDriveItem.md) |
| DELETE | `/v1.0/me/drive/following/{item-id}` | Unfollow a DriveItem | [View](../operations/UnfollowDriveItem.md) |
| GET | `/v1beta1/me/drive/sharedWithMe` | Get a list of driveItem objects shared with the owner of a drive. | [View](../operations/ListSharedWithMe.md) |
