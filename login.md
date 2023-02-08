# Login Endpoint

This endpoint is used to handle user login requests.

## Request

The following table describes the required parameters for a login request:

| Parameter | Type | Description |
|-----------|------|-------------|
| email | string | The email address of the user. |
| username | string | The username of the user. |
| password | string | The password of the user. |

## Response

A successful request returns a JSON object with the following properties:

| Property | Type | Description |
|----------|------|-------------|
| jwtToken | string | A JSON Web Token that can be used to authenticate subsequent requests. |
| refreshToken | string | A refresh token that can be used to generate a new JWT token. |

A failed request returns a JSON object with an error property that describes the error.