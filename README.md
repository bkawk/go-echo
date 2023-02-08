# Go-Echo Boilerplate API

A pre-built REST API project that combines Go and the Echo framework to provide a simple yet robust foundation for developers.

## Introduction

Go-Echo Boilerplate API is designed to be a fast and efficient solution for building RESTful APIs. The [Echo](https://echo.labstack.com/)  framework is known for its performance, scalability, and ease of use. With this boilerplate, you can take advantage of Echo's lightning-fast routing and middleware support to build APIs that can handle large amounts of traffic with ease. Whether you're building a high-performance application or just need to get up and running quickly, Go-Echo provides a foundation that is both fast and reliable. By leveraging the power of Go and the Echo framework, you can ensure that your API will perform at peak efficiency and provide a smooth experience for your users.

## Key Features

- User authentication and authorization
- MongoDB Atlas integration
- Rate limiting
- Body size limit
- Secure middleware
- CORS support
- Email verification
- Password reset

## Requirements
 - Go
 - Echo
 - MongoDB Atlas
 - SMTP Server (e.g. Gmail, SendGrid, etc.)

## Getting Started

1. Clone this repository
2. Make sure you have Go installed on your machine
3. Set up a MongoDB Atlas cluster and obtain your connection string
4. Create a .env file in the root of the project and add the following variables:


 - PORT: The port on which the server will run, e.g. :8000
 - MONGO_URL: The connection string for your MongoDB Atlas cluster, e.g. mongodb+srv://<username>:<password>@cluster0.mongodb.net/test?retryWrites=true&w=majority
 
 - MONGO_DB: The name of the database, e.g. databaseName
 - BCRYPT_PASSWORD: The bcrypt password, e.g. bcryptPassword
 - EMAIL_FROM: The email address from which emails will be sent, e.g. my@email.com
 - SMTP_SERVER: The SMTP server used for sending emails, e.g. smtp.example.com
 - SMTP_PORT: The port used for the SMTP server, e.g. 587
 - SMTP_PASSWORD: The password for the SMTP server, e.g. password
 - JWT_SECRET: The secret used for signing JSON Web Tokens (JWT), e.g. jwtSecret
 - VERIFY_URL: The URL used for email verification, e.g. http://example.com/verify
 - RESET_EMAIL_URL: The URL used for resetting the email, e.g. http://example.com/reset-email


5. Run the following command to install the dependencies:

``` go get ```

6. Run the following command to start the server:

``` go run main.go  ```

## Endpoints
 - [Health Check](health.md): GET /health
 - [Check username availability](username_availability.md): GET /username/:username
 - [Register](register.md): POST /register
 - [Login](login.md): POST /login
 - [Verify Email](verify_email.md): GET /verify/:token
 - [Resend Verification Email](resend_verification.md): POST /resend-verification
 - [Reset Password](password_reset.md): POST /reset-password
 - [Refresh Token](refresh_token.md): POST /refresh
 - [Invalidate Refresh Token](invalidate_token.md): DELETE /refresh
 
 - [Update Profile](update_profile.md): PUT /profile ***
 - [Retrieve Profile](retrieve_profile.md): GET /profile/:user_id ***

## How to Contribute
If you would like to contribute to the project, please follow these steps:

1. Fork the repository
2. Create a new branch for your contribution
3. Commit your changes to the new branch
4. Submit a pull request to the original repository
5. We welcome all contributions, including bug fixes, new features, and documentation 6. improvements. Thank you for your support!
