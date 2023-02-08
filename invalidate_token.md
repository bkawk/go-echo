# Refresh Delete Endpoint Error Responses

The `RefreshDelete` endpoint may return the following error responses in case of a failure. 

## Responses

The endpoint returns a JSON object with the following fields:

| Field      | Type   | Description                                                           |
|------------|--------|-----------------------------------------------------------------------|
| Message    | string | A string indicating the status of the request. Can be an error message. |

### HTTP 400 Bad Request

When the provided `refreshToken` is invalid, the endpoint returns the following response:


```
HTTP 400 Bad Request

{
"Message": "Invalid refresh token"
}
```

### HTTP 500 Internal Server Error

In case of a failure to update the user in the database, the endpoint returns the following response:


```
HTTP 500 Internal Server Error

{
"Message": "Failed to update user in database"
}
```