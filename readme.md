# Item API

This is an API built with Golang Gin framework and GORM for database access. Redis is also used for caching to improve performance. Repository pattern is implemented to separate business logic from data access.

## Installation

1. Clone this repository:
```bash
git clone https://github.com/yosikez/item-api.git
```
2. Create a .env file based on the .env.example file and set the necessary configuration variables.
3. Start the containers using Docker Compose
```bash
docker compose up --build
```

## Usage

Once the containers are running, you can interact with the API using the following endpoints:

- `GET /items` : Retrieve all items
- `GET /items/:id` : Retrieve an item by ID
- `POST /items` : Create a new item
- `PUT /items/:id` : Update an existing item by ID
- `DELETE /items/:id` : Delete an item by ID