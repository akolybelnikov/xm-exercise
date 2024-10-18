# Company Management REST API

This project implements a REST API to manage companies. The API allows clients to create, read, update, and delete company records.

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


## Building the Project

1. **Build the project:**

   ```sh
   make build
   ```

2. **Run the application:**

    ```sh
    make run
    ```

    This will start the application inside the Docker container.
3. **Run the tests:**

    ```sh
    make test
    ```

    This will run the unit tests and generate a coverage report.
4. **Lint the code:**

    ```sh
    make lint
    ```

    This will run the linter to enforce coding standards.
5. **Clean up the project:**

    ```sh
    make clean
    ```

    This will remove the project build.

## Stopping the Application

To stop the running application, use:

```sh
make stop
```

## Configuration

Make sure to customize any configuration parameters based on your environment and requirements. Common configurations might include database connection strings, server ports, and logging levels.

## Usage

Once the application is running, you can interact with the API to perform operations on company records. Typical operations include:

- **Create a new company record**
- **Retrieve company information**
- **Update existing company records**
- **Delete company records**

Refer to the API documentation or implemented endpoints in the source code for detailed usage instructions and available routes.

## Logging

The application uses the built-in `log` package for logging. By default, the logs will be printed to the standard output.

## Linting

The project employs the `golangci-lint` linter to enforce coding standards and prohibit the use of certain print functions.
Make sure to follow these linting rules to maintain code quality and consistency.
