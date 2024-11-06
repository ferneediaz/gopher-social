# Gopher Socials API

Gopher Socials API is a backend service for the Gopher Socials platform, providing endpoints for user authentication, post creation, commenting, and more. It is built with Go and follows best practices for API development, including comprehensive documentation with Swagger.

## Table of Contents

- Features
- Architecture
- Installation
- Usage
- API Documentation
- Database Migrations
- Seeding the Database
- Contributing
- License

## Features

- User registration and authentication
- Create, update, delete, and fetch posts
- Commenting on posts
- Following and unfollowing users
- Fetching user feeds
- Comprehensive API documentation

## Architecture

The project is structured as follows:

.
├── cmd/
│   └── api/
│       ├── api.go
│       ├── errors.go
│       ├── feed.go
│       ├── health.go
│       ├── json.go
│       ├── main.go
│       ├── posts.go
│       └── users.go
├── internal/
│   ├── db/
│   │   ├── db.go
│   │   └── seed.go
│   ├── env/
│   │   └── env.go
│   └── store/
│       ├── comments.go
│       ├── followers.go
│       ├── pagination.go
│       ├── posts.go
│       ├── storage.go
│       └── users.go
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── makefile
├── go.mod
├── go.sum
└── .air.toml

## Installation

1. **Clone the repository:**

```sh
   git clone https://github.com/ferneediaz/gopher-socials.git
   ```
cd gopher-socials

go mod download

## Setup Environment Variables:

Create a .env file based on the .envrc template with the necessary environment variables.
```
make build
```
## Running the API
Start the API server using Air for live reloading
```
air
```
## Running Database Migrations
```
make migrate-up
```
## Create a new migration:

```
make migrate-create <migration_name>
```
##Rollback the last migration:
```
make migrate-down
```
## Seeding database 
```
make seed
```
## API DOCUMENTATION
```
http://localhost:3000/swagger/index.html
```