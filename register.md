# Register Endpoint

The `RegisterPost` endpoint handles user registration requests. 

## Request

The endpoint accepts JSON payload with the following parameters:

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| username  | string | Yes | The desired username of the user. |
| email     | string | Yes | The email address of the user. |
| password  | string | Yes | The password of the user. |

Example request body:

```
{
"username": "john.doe",
"email": "john.doe@example.com",
"password": "secretpassword"
}
```

## Response

The endpoint returns a JSON response with the following properties:

| Property | Type | Description |
|----------|------|-------------|
| message  | string | Success message if the account was created successfully. |
| error    | string | Error message if there was a problem creating the account. |

### Success Response

HTTP 200 OK is returned with the following JSON payload:

```
{
"message": "Your account has been successfully created"
}
```

### Error Responses

| HTTP Status Code | Error Type | Description |
|------------------|------------|-------------|
| 400 Bad Request | Validation Error | Occurs when the input data does not meet the requirements |
| 400 Bad Request | Duplicate Data Error | Occurs when the username or email is already taken |
| 500 Internal Server Error | Hash Password Error | Occurs when the password hash generation fails |
| 500 Internal Server Error | Generate UUID Error | Occurs when generating a unique identifier fails |
| 500 Internal Server Error | Save User Error | Occurs when saving the user to the database fails |
| 500 Internal Server Error | Email Error | Occurs when sending an email fails |