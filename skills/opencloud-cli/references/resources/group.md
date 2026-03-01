# group

## Operations

| Method | Path | Summary | Details |
|--------|------|---------|----------|
| GET | `/v1.0/groups/{group-id}` | Get entity from groups by key | [View](../operations/GetGroup.md) |
| DELETE | `/v1.0/groups/{group-id}` | Delete entity from groups | [View](../operations/DeleteGroup.md) |
| PATCH | `/v1.0/groups/{group-id}` | Update entity in groups | [View](../operations/UpdateGroup.md) |
| GET | `/v1.0/groups/{group-id}/members` | Get a list of the group's direct members | [View](../operations/ListMembers.md) |
| POST | `/v1.0/groups/{group-id}/members/$ref` | Add a member to a group | [View](../operations/AddMember.md) |
| DELETE | `/v1.0/groups/{group-id}/members/{directory-object-id}/$ref` | Delete member from a group | [View](../operations/DeleteMember.md) |
