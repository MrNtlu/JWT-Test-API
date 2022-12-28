# Test API

This is a test REST API for JWT Login.

JWT Token will be refreshed in every 30 seconds.

[Base URL](https://jwt-test-api.onrender.com)

```
Endpoints /api ->
  - /auth ->
    - POST /login
    - POST /register
    - POST /logout
    - GET /refresh
  - /user ->
    - GET /info
```

## Run Locally

If you want to run this backend in your local machine, please follow these instructions.

1. Download source code, it'll have some missing points for security reasons.
2. Create your own JWT Secret Key from this link. https://jwt.io/
3. Create your own MongoDB Database. I've used MongoDB Atlas.
4. Create .env file.
5. You'll need to add 5 environment variables,
    - BASE_URI = "YOUR URI, in my case it's jwt-test-api.onrender.com"
    - MONGO_ATLAS_URI = "Local or Hosted MongoDB URI"
    - JWT_SECRET_KEY = "Your JWT Secret Key"
    - ENV = "Production" This one is optional, if you are going to work locally, you don't need this.

That's it.