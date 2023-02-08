# ResetPasswordPost Endpoint

This endpoint handles user password reset requests.

## Request

The endpoint accepts a JSON request with the following fields:

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `PasswordResetToken` | string | Yes | A unique token sent to the user's email to reset their password |
| `Password` | string | Yes | The new password to be set |

## Responses

The endpoint returns a JSON response with the following fields:

### Success

The request was successful, and the password has been reset.

| HTTP Status | Response |
|-------------|----------|
| 200 OK | `Password successfully reset` |

### Errors

The request failed for the following reasons:

| HTTP Status | Response |
|-------------|----------|
| 400 Bad Request | `Invalid Reset token or user not verified`