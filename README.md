# BlazeStack

BlazeStack is a small, container-friendly demo stack composed of:
- blazestack-frontend-web: Next.js (App Router, TS, Tailwind, MUI) web UI
- blazestack-ms-incidents: Go (Gin + GORM + Postgres) microservice for incidents and auth
- PostgreSQL: stateful DB with simple DDL seeded at startup (db_init/ddl.sql)

The repository includes a docker-compose.yml to spin up PostgreSQL. Frontend and backend can be run locally for fast development.


## Architecture overview
- Frontend (Next.js):
  - Dev server on http://localhost:3000
  - Calls the API at http://localhost:8080 (see src/app/login/page.tsx and IncidentsTable.tsx)
- Backend (Go, Gin):
  - Listens on 127.0.0.1:8080 by default
  - Depends on PostgreSQL
  - Requires environment variables for DB and JWT configuration
- Database (Postgres 16):
  - Exposed via POSTGRES_PORT (default recommended: 5432)
  - Initializes schema from db_init/ddl.sql on first run


## Prerequisites
- Docker and Docker Compose
- Node.js 20+ and a package manager (npm, pnpm, yarn, or bun)
- Go 1.24+


## Quick start (recommended)
This starts only PostgreSQL with Docker. Run the backend and frontend locally.

1) Create a .env file in the repo root with Postgres variables used by docker-compose:

```
# docker-compose (database) environment
POSTGRES_PORT=5432
POSTGRES_USER=blaze
POSTGRES_PASSWORD=blaze
POSTGRES_DB=blazestack
```

2) Start the database:

```
docker compose up -d
```

- The DB will be available on localhost:5432 and initialized with the incidents table from db_init/ddl.sql.

3) Prepare environment for the Go service (blazestack-ms-incidents/.env):

```
# Server
PORT=8080

# Database (match docker-compose values)
DB_HOST=localhost
DB_PORT=5432
DB_USERNAME=blaze
DB_PASSWORD=blaze
DB_NAME=blazestack
DB_SSL=disable
DB_MIN_CONNECTIONS=1
DB_MAX_CONNECTIONS=5
DB_LOGGER=true

# Auth
JWT_SECRET=dev_secret_change_me
```

4) Run the Go service:

```
cd blazestack-ms-incidents
go mod download
go run .
```

- The API exposes:
  - GET / -> { "message": "All right" }
  - GET /ping -> "pong"
  - POST /api/v1/auth/login
  - GET /api/v1/auth/profile (Authorization: Bearer <token>)
  - GET /api/v1/incidents (Authorization: Bearer <token>)
  - POST /api/v1/incidents (Authorization: Bearer <token>)

5) Run the frontend:

```
cd blazestack-frontend-web
npm install
npm run dev
```

- Open http://localhost:3000
- The login form posts to http://localhost:8080/api/v1/auth/login. After login, the app requests the profile and then lists incidents.


## Environment variables
- Database (docker-compose expects in repo root .env):
  - POSTGRES_PORT (required)
  - POSTGRES_USER (required)
  - POSTGRES_PASSWORD (required)
  - POSTGRES_DB (required)
- Backend (blazestack-ms-incidents/.env):
  - PORT (default 8080)
  - DB_HOST, DB_PORT, DB_USERNAME, DB_PASSWORD, DB_NAME (required)
  - DB_SSL (e.g., disable)
  - DB_MIN_CONNECTIONS, DB_MAX_CONNECTIONS (optional)
  - DB_LOGGER (optional)
  - JWT_SECRET (required)


## Troubleshooting
- Database cannot connect:
  - Ensure docker compose is up and POSTGRES_PORT is not in use.
  - Verify blazestack-ms-incidents/.env matches your DB settings.
- Migrations/DDL:
  - Initial schema is created from db_init/ddl.sql by the postgres container at first startup.
- CORS:
  - The API enables CORS for all origins via gin-contrib/cors; localhost:3000 should work out of the box.
- Static export in Next.js:
  - next.config.ts uses output: 'export' and unoptimized images to support static export if you later run npm run build. For dev, npm run dev is enough.

## Project layout
- blazestack-frontend-web: Next.js 15, React 19
- blazestack-ms-incidents: Go Gin service, GORM Postgres
- db_init: SQL DDL applied at DB container initialization
- docker-compose.yml: Database service definition

## Tradeoffs/assumptions
- Using postgresql to persist data (easy with docker compose)
- Using GORM for ORM
- Using Gin for API framework
- Using Next.js for frontend
- Using Docker for local development
- Using Docker Compose for local development
- Using JWT for authentication
- Using Tailwind for styling
- Using Postman for API testing

## AI Ussage
- For small components in the frontend
- For DDL generation based on DB schema
- For demo records in the DB

## Base
- This backend is based on a personal boilerplate project in golang: https://github.com/jessusandres/golang-boilerplate

## Nexts Steps
Because this is a demo stack and the development time limit is a maximum of 60 minutes, I will implement this work afterward:
- Modal for incident creation
- Pagination
- Search
- Improve frontend styling, layout, accessibility, etc.
- Improve backend logging, serialization, and ORM usage including the repository pattern.
- Unit and Integration tests
- Dockerize the backend
- Dockerize the frontend
- CI/CD
- Monitoring
- Secrets for the environment
- Secrets for the database

## License
MIT (or project-specific; update as appropriate)
