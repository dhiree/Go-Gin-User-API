# User Management API

This project is a simple User Management API built with Go and the Gin framework. It allows you to create, retrieve, update, and delete user records in a MongoDB database.

## Features

- **Create User**: Add a new user to the database.
- **Get User by ID**: Retrieve a user by their unique ID.
- **Update User**: Modify details of an existing user.
- **Delete User**: Remove a user from the database.
- **Get All Users**: Retrieve a list of all users.

## Prerequisites

- Go 1.18 or later
- MongoDB
- Gin framework

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/dhiree/Go-Gin-User-API.git
    cd yourrepository
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Set up your MongoDB connection:

    - Edit the `config/config.go` file with your MongoDB connection details.

## Usage

1. Run the application:

    ```bash
    go run main.go
    ```

2. The server will start on `http://localhost:8080`.

## API Endpoints

- **POST /users**: Create a new user
- **GET /users/:id**: Get a user by ID
- **PUT /users/:id**: Update an existing user
- **DELETE /users/:id**: Delete a user by ID
- **GET /users**: Get all users

## Example Requests

1. **Create User**

    ```bash
    curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"name": "Dheerendra Singh", "email": "bhandaridheere@gmail.com", "password": "password123"}'
    ```

2. **Get User by ID**

    ```bash
    curl http://localhost:8080/users/60d5f483e3d3f4e8d4c1b8e9
    ```

3. **Update User**

    ```bash
    curl -X PUT http://localhost:8080/users/60d5f483e3d3f4e8d4c1b8e9 -H "Content-Type: application/json" -d '{"name": "Dheerendra Singh"}'
    ```

4. **Delete User**

    ```bash
    curl -X DELETE http://localhost:8080/users/60d5f483e3d3f4e8d4c1b8e9
    ```

5. **Get All Users**

    ```bash
    curl http://localhost:8080/users
    ```

