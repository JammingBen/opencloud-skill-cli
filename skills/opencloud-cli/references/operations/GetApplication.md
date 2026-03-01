# GET /v1.0/applications/{application-id}

**Resource:** [applications](../resources/applications.md)
**Get application by id**
**Operation ID:** `GetApplication`

## Parameters

| Name | In | Type | Required | Description |
|------|------|------|----------|-------------|
| `application-id` | path | string | Yes | key: id of application |

## Responses

| Status | Description |
|--------|-------------|
| 200 | OK |
| default | (reference) |

**Success Response Schema:**

[application](../schemas/application/application.md)

