## Resend Verification Endpoint

This endpoint is used to resend the verification code to a user's email address.

### Request

| Method | Endpoint |
| ------ | -------- |
| POST   | /resend-verification |

#### Body

| Field  | Type   | Description             |
| ------ | ------ | ----------------------- |
| email  | String | Email of the user       |

#### Example Body

```json
{
  "email": "user@example.com"
}
```

### Response

#### Success

| Status | Message                   |
| ------ | ------------------------ |
| 200    | Verification code sent   |

#### Bad Request

| Status | Message                                                                            |
| ------ | ---------------------------------------------------------------------------------- |
| 400    | Email not found                                                                    |
| 400    | Verification code already sent. Please wait x minutes before trying again.         |

#### Internal Server Error

| Status | Message                           |
| ------ | --------------------------------- |
| 500    | Failed to generate user ID        |
| 500    | Failed to update user record      |
| 500    | An error occured sending email    |