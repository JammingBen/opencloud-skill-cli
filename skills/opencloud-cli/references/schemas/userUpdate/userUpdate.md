# userUpdate

Represents updates to an Active Directory user object.

**Type:** object

## Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `id` | string | No | Read-only. |
| `accountEnabled` | boolean | No | Set to "true" when the account is enabled. |
| `appRoleAssignments` | appRoleAssignment[] | No | The apps and app roles which this user has been assigned. |
| `displayName` | string | No | The name displayed in the address book for the user. This value is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates. Returned by default. Supports $orderby. |
| `drives` | drive[] | No | A collection of drives available for this user. Read-only. |
| `drive` | [drive](drive.md) | No |  |
| `identities` | objectIdentity[] | No | Identities associated with this account. |
| `mail` | string | No | The SMTP address for the user, for example, 'jeff@contoso.opencloud.com'. Returned by default. |
| `memberOf` | group[] | No | Groups that this user is a member of. HTTP Methods: GET (supported for all groups). Read-only. Nullable. Supports $expand. |
| `onPremisesSamAccountName` | string | No | Contains the on-premises SAM account name synchronized from the on-premises directory. |
| `passwordProfile` | [passwordProfile](passwordProfile.md) | No |  |
| `surname` | string | No | The user's surname (family name or last name). Returned by default. |
| `givenName` | string | No | The user's givenName. Returned by default. |
| `userType` | string | No | The user`s type. This can be either "Member" for regular user, "Guest" for guest users or "Federated" for users imported from a federated instance. |
| `preferredLanguage` | [language](language.md) | No |  |
| `signInActivity` | [signInActivity](signInActivity.md) | No |  |

