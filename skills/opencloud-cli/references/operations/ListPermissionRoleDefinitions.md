# GET /v1beta1/roleManagement/permissions/roleDefinitions

**Resource:** [roleManagement](../resources/roleManagement.md)
**List roleDefinitions**
**Operation ID:** `ListPermissionRoleDefinitions`

Get a list of `unifiedRoleDefinition` objects for the permissions provider. This list determines the roles that can be selected when creating sharing invites.


## Responses

| Status | Description |
|--------|-------------|
| 200 | A list of permission roles than can be used when sharing with users or groups. |
| default | (reference) |

**Success Response Schema:**

Array of [unifiedRoleDefinition](../schemas/unifiedRoleDefinition/unifiedRoleDefinition.md)

