# Verify Email 

Verifies an email using the provided `verificationCode` as a query parameter in the URL.

#### Endpoint

```
GET /verify-email?verificationCode=<verificationCode>
```

#### URL Parameters

| Parameter | Type | Description |
| --- | --- | --- |
| verificationCode | string | Required. The verification code used to verify the email. |

#### Success Response

Status Code: 200 OK

```
{
"message": "User marked as verified successfully!"
}
```

#### Error Responses


The following table lists the possible error responses for the `GET /verify-email` endpoint:

| HTTP Status Code | Error Code | Description |
| --- | --- | --- |
| 500 Internal Server Error | `Failed to verify the user` | The verification failed due to a server-side error. |
| 500 Internal Server Error | `Invalid verification code` | The provided `verificationCode` is invalid. |


