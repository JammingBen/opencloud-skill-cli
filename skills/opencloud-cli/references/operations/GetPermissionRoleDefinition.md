# GET /v1beta1/roleManagement/permissions/roleDefinitions/{role-id}

**Resource:** [roleManagement](../resources/roleManagement.md)
**Get unifiedRoleDefinition**
**Operation ID:** `GetPermissionRoleDefinition`

Read the properties and relationships of a `unifiedRoleDefinition` object.


## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `role-id` | path | string | Yes | key: id of roleDefinition |

## Responses

| Status | Description |
|--------|-------------|
| 200 | OK |
| default | (reference) |

**Success Response Schema:**

[unifiedRoleDefinition](../schemas/unifiedRoleDefinition/unifiedRoleDefinition.md)

