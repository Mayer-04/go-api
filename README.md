# Go-mongo-rest-api 🐹🍃

This project is an API REST developed in the Go programming language, making use of its **standard library** which offers improvements in routing and other features since *version 1.22* The API implements CRUD operations (Create, Read, Update, Delete) to interact with a **MongoDB** database. Additionally, the **validator** package is used for data validations.

## Features

- Implementation of a RESTful API with enhanced routing available from Go *1.22*.
- Use of MongoDB as a database to store and manage data.
- Data validation through the **validator** package to maintain data integrity and consistency.

## Installation and Requirements

### Requirements

- Go *1.22* or higher installed on the system.
- MongoDB installed and configured locally or access to a remote instance.

### Installation

1.Clone the repository:

```bash
git clone https://github.com/Mayer-04/mongo-api-go.git
```

2.Install project dependencies:

```bash
go mod tidy
```

3.Clone the **.env.example** file to **.env** to configure the environment variables. Database Credentials
4.Start the application:

```bash
go run cmd/api-server/main.go
```

## Usage

Once the application is up and running, you can interact with the REST API using any HTTP client. Here are some examples of how you can perform CRUD operations:

- **Create a new product:**

```bash
curl -X POST http://localhost:8080/api/users -d '{"name": "TV", "description": "samsung tv", "price": 1.234, "category": "technology", "stock": 23, "image": "tv.jpg"}'
```

- **Read all products:**

```bash
curl -X GET http://localhost:8080/api/users
```

- **Read product by id:**

```bash
curl -X GET http://localhost:8080/api/users/{id}
```

- **Update product information:**

```bash
curl -X PUT http://localhost:8080/api/users/{id} -d '{"name": "TV", "description": "samsung tv", "price": 1.234, "category": "technology", "stock": 23, "image": "tv.jpg"}'
```

- **Delete a product:**

```bash
curl -X DELETE http://localhost:8080/api/users/{id}
```
