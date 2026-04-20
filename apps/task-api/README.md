# task-api

Go REST API for the todo application.

## Requirements

- [Go 1.26+](https://go.dev/dl/)
- [Task](https://taskfile.dev/installation/) - task runner

## Quick start

### 1. Environment variables

Copy `.env.example` in the **project root**:

```bash
cp .env.example .env
```

### 2. Start infrastructure

From the **project root**:

```bash
docker compose up -d
```

### 3. Install tools

```bash12
task tools
```

### 4. Run migrations

```bash
task migrate:up
```

### 5. Start the API

With live reload:
```bash
task dev
```

Without live reload:
```bash
task run
```

API available at `http://localhost:3000/api`

## Swagger UI

`http://localhost:3000/swagger`

Authorization via Keycloak (PKCE):
- client: `swagger-ui`
- test user: `user@example.com` / `user`