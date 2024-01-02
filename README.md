# web-crud

This project was developed as part of the Go programming language training at Alura. It involves the creation of a full-stack CRUD web application using the Go SDL.

# Dependencies

To ensure smooth operation of serenatto, make sure you have the following dependencies installed:

- Go (version 1.21)
- PostgreSQL 16.1

# Usage

1. After installing `go`, clone the project and navigate to the project directory with the following commands:
```
git clone https://gitlab.com/alura-courses-code/golang/web-crud.git
cd web-crud
```

2. Install all dependencies required to run the project, especially `air`:
```
go mod tidy
```

3. Run the project and open `localhost:8080` in your browser using the following command:
```
air
```

**OBS**: Create and configure a PostgreSQL database, and make the connection in a .env.local file. You can use the [.env](https://gitlab.com/alura-courses-code/golang/web-crud/-/blob/main/.env) as an example.

## Contributing

If you wish to contribute to this project, feel free to open a merge request. We welcome all forms of contribution!

## License

This project is licensed under the [MIT](https://gitlab.com/alura-courses-code/golang/web-crud/-/blob/main/LICENSE). Refer to the LICENSE file for more details.
