# POST /v1.0/users/{user-id}/appRoleAssignments

**Resource:** [user.appRoleAssignment](../resources/user-appRoleAssignment.md)
**Grant an appRoleAssignment to a user**
**Operation ID:** `user.CreateAppRoleAssignments`

Use this API to assign a global role to a user. To grant an app role assignment to a user, you need three identifiers:
* `principalId`: The `id` of the user to whom you are assigning the app role.
* `resourceId`: The `id` of the resource `servicePrincipal` or `application` that has defined the app role.
* `appRoleId`: The `id` of the `appRole` (defined on the resource service principal or application) to assign to the user.


## Request Body

New app role assignment value

**Required:** Yes

**Content Types:** `application/json`

**Schema:** [appRoleAssignment](../schemas/appRoleAssignment/appRoleAssignment.md)

## Responses

| Status | Description |
|--------|-------------|
| 200 | Created new app role assignment. |
| default | (reference) |

**Success Response Schema:**

[appRoleAssignment](../schemas/appRoleAssignment/appRoleAssignment.md)

