# Endpoint: RefreshPost

Refreshes a JWT for a user with a valid refresh token.

## Request

Method | Endpoint | Description
-------|----------|------------
POST   | /refresh | Refreshes a JWT for a user with a valid refresh token

### Request Body

Field | Type | Required | Description
------|------|----------|------------
refreshToken | string | Yes | Refresh token of the user

## Response

HTTP Code | Response Body | Description
----------|---------------|------------
200       | `{"message": "Success", "data": {"jwt": "<jwt>"}}` | If the JWT was successfully refreshed
400       | `{"message": "Invalid refresh token"}` | If the provided refresh token is invalid
500       | `{"message": "Failed to find user in database"}` | If there was a problem finding the user in the database
500       | `{"message": "Failed to generate JWT"}` | If there was a problem generating the JWT

### Response Body

Field   | Type   | Description
--------|--------|------------
message | string | Status message
data    | object | JSON object with the refreshed JWT

## Example

Request:

```
POST /refresh
Content-Type: application/x-www-form-urlencoded

refreshToken=<refresh_token>
```

Response:

```
HTTP/1.1 200 OK
Content-Type: application/json

{"message": "Success", "data": {"jwt": "<jwt>"}}
```
