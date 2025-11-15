# go-crud

> A full-stack CRUD web application built with Go and PostgreSQL.

## About the Project

This project is a full-stack web application that demonstrates Create, Read, Update, and Delete (CRUD) functionality. It was developed as part of a Go programming language training course at Alura.

## Tech Stack

*   [Go](https://golang.org/)
*   [PostgreSQL](https://www.postgresql.org/)
*   [Air](https://github.com/cosmtrek/air) (for live reloading)

## Usage

Below are the instructions for you to set up and run the project locally.

### Prerequisites

You need to have the following software installed:

*   [Go](https://golang.org/dl/) (version 1.21 or higher)
*   [PostgreSQL](https://www.postgresql.org/download/) (version 16.1 or higher)
*   [Air](https://github.com/cosmtrek/air#installation) (for development)

### Installation and Setup

Follow the steps below:

1.  **Clone the repository**
    ```bash
    git clone https://github.com/luizvilasboas/go-crud.git
    ```

2.  **Navigate to the project directory**
    ```bash
    cd go-crud
    ```

3.  **Install Go dependencies**
    ```bash
    go mod tidy
    ```

4.  **Configure the database**
    *   Ensure your PostgreSQL server is running.
    *   Create a new database for the project.
    *   Create a `.env.local` file in the project root (you can use the `.env` file as a template) and add your database connection details.

### Workflow

To run the project with live reloading, use `air`:
```bash
air
```
The application will be available in your browser at `http://localhost:8080`.

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
