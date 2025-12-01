# account-service
## Making Requests
### Create User

POST /user

Create a new user with a username and password.

Request Body (JSON)
````
{
"username": "alice",
"password": "mypassword"
}
````
Example Python Request
````
import requests

url = "http://localhost:8080/user"
data = {
"username": "alice",
"password": "mypassword"
}

response = requests.post(url, json=data)
print(response.status_code)
print(response.json())
````
Example Response (201 Created)
````
{
"id": 1,
"username": "alice",
"created_at": "2025-01-01T12:00:00Z"
}
````
### Get User by ID

GET /user?id={id}

Example Python Request
````
import requests

response = requests.get("http://localhost:8080/user?id=1")
print(response.status_code)
print(response.json())
````
Example Response (200 OK)
````
{
"id": 1,
"username": "alice",
"created_at": "2025-01-01T12:00:00Z"
}
````
Errors

| Status | Meaning  |
|--------|-------------|
| 400    | missing or invalid id |
| 404    | user not found |


### Get User by Username

GET /user?username={username}

Example Python Request
````
import requests

response = requests.get("http://localhost:8080/user?username=alice")
print(response.status_code)
print(response.json())

Example Response (200 OK)
{
"id": 1,
"username": "alice",
"created_at": "2025-01-01T12:00:00Z"
}
````

### User Login

POST /user/login

Login requires username + password.
If successful, a JWT token is returned.

Request Body (JSON)
````
{
"username": "alice",
"password": "mypassword"
}
````
Example Python Request
````
import requests

url = "http://localhost:8080/user/login"
data = {
"username": "alice",
"password": "mypassword"
}

response = requests.post(url, json=data)
print(response.status_code)
print(response.json())
````
Example Response (200 OK)
````
{
"token": "eyJhbGciOiJIUzI1NiIs...",
"message": "Logged in successfully"
}
````
Invalid Credentials Response (401)
````
{
"error": "Login failed. Check your credentials."
}
````
🔑 JWT Token Structure

Tokens include:
````
{
"username": "alice",
"user_id": 1,
"exp": 1735758123
}
````

Signing algorithm: HS256

Expiration: 24 hours


