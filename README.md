# Company Management REST API

This project implements a REST API to manage companies. The API allows clients to create, read, update, and delete
company records.

## Prerequisites

- [Docker](https://www.docker.com/get-started) installed
- [Docker Compose](https://docs.docker.com/compose/install/) installed
- [Make](https://www.gnu.org/software/make/) installed

## Installation and Setup

1. **Clone the repository:**

    ```sh
    gh repo clone akolybelnikov/xm-exercise
    cd xm-exercise
    ```

2. **Set up the development environment:**

    ```sh
    make setup
    ```

## Configuration

The application uses environment variables and static YML files for configuration. As a minimum, the `APP_ENV`
environment variable must be set to `dev` for the configuration to be loaded correctly in the development environment.

The static configuration files are located in the `config` directory. The configuration files are loaded based on the
name of the file and the value of the `APP_ENV` environment variable. For example, the `config/dev.yml` file will be
loaded when the `APP_ENV` environment variable is set to `dev`.

The environment variables are loaded from a `.env` file in the root directory of the project. The `.env` file is not
committed to the repository and must be created manually.

## Running the Project Locally

1. **Run the application:**

    ```sh
    make run
    ```

   This will build and start the application. Make sure to have a Postgres database running locally.

2. **Lint the code:**

    ```sh
    make lint
    ```

   This will run the linter to enforce coding standards.

## Running the Project with Docker Compose

This will initialize a Docker container with the application running inside it, and a Postgres database for data
storage.
The application will be accessible at `http://localhost:8080`. The Postgres database will be initialized with the DDL.

1. **Start the Docker container:**

    ```sh
    docker compose up -d
    ```
2. **Stop the Docker container:**

    ```sh
    docker compose down
    ```

## Testing the Application

1. **Run all the tests:**

    ```sh
    make test
    ```

   This will run the all the tests in the project, unit and integration tests. Make sure to have a Postgres database
   running locally.

2. **Run only the unit tests:**

    ```sh
    make test-unit
    ```
   This will run the unit tests in the project, if you prefer to run only the unit tests.

## Usage

Once the application is running, you can interact with the API to perform operations on company records. Typical
operations include:

- **Login to the application to obtain a valid JWT token**
- **Create a new company record**
- **Retrieve company information**
- **Update existing company records**
- **Delete company records**

Refer to the API documentation or implemented endpoints in the source code for detailed usage instructions and available
routes.

## Authentication

The application uses JWT tokens for authentication. The tokens are generated when a user logs in to the application.
The tokens are required to access all the routes.

In order to log in to the application, you need to send a POST request to the `/login` endpoint with the following
payload:
   
```json
{
  "username": "admin",
  "password": "admin"
}
```

For example with curl:

 ```sh
    curl -X POST -v http://localhost:8080/login -d '{"username": "admin", "password": "admin"}'
 ```

The response will contain the JWT token that you can use to access the protected routes,
For example:
```sh
    curl -X GET http://localhost:8080/api/v1/companies/some-uuid-id -H "Authorization: Bearer your_jwt_token_here"
```

## Logging

The application uses the built-in `log` package for logging. By default, the logs will be printed to the standard
output.

## Linting

The project employs the `golangci-lint` linter to enforce coding standards and prohibit the use of certain print
functions.
Make sure to follow these linting rules to maintain code quality and consistency.
