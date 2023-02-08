# Check Username Endpoint

## Get a user's availability by checking their username

### Request

`GET /check-username?username=<username>`

### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| username | string | Required. The username to check availability for |

### Response

Success:
```
HTTP/1.1 200 OK
Content-Type: application/json

{
"message": "Username available",
"available": true
}
```

or

```
HTTP/1.1 200 OK
Content-Type: application/json

{
"message": "Username not available",
"available": false
}
```

### Error Responses

| HTTP Code | Description |
| --------- | ----------- |
| 400 Bad Request | If the query parameter `username` is missing |
| 500 Internal Server Error | If an error occurs while checking the username in the database |

### Response

Bad Request:

```
HTTP/1.1 400 Bad Request
Content-Type: application/json

{
"message": "Username is required"
}
```

Internal Server Error:

```
HTTP/1.1 500 Internal Server Error
Content-Type: application/json

{
"message": "An error occurred while checking the username",
"error": "<error message>"
}
```