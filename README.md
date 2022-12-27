# Test API

This is a test REST API for JWT Login.

JWT Token will be refreshed in every 30 seconds.

[Base URL](https://jwt-test-api.onrender.com)

Endpoints `/api` ->
  - `/auth` ->
    - POST `/login` 
    - POST `/register`
    - POST `/logout`
    - GET `/refresh`
  - `/user` ->
    - GET `/info`