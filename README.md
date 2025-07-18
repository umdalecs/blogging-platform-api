# Golang Blogging Platform API

This is my golang solution to https://roadmap.sh/projects/blogging-platform-api.

## Features

- *Postgres Storage*: The app stores posts in postgresql database
- *Search filter*: You can easily filter posts using query params

## Running application

### 1. Clone the repo

```bash
git clone https://github.com/umdalecs/blogging-platform-api
cd blogging-platform-api
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Environment variables

```bash
cp .env.example .env
```

> [!Note]
> This proyect depends on postgres storage and database should be initialized
> using the posts ddl at `./database_files/posts.sql`

### 4. Build and run the application

```bash
go build -o out/ .
./out/blogging-platform-api
```

## Usage

### Get posts
```bash
curl http://localhost:8080/api/v1/posts?term=
```

### Get single post
```bash
curl http://localhost:8080/api/v1/posts/{id}
```

### Create
```bash
curl -X POST http://localhost:8080/api/v1/posts \
-H "Content-Type: application/json" \
-d '{
  "title": "My First Blog Post",
  "content": "This is the content of my first blog post.",
  "category": "Technology",
  "tags": ["Tech", "Programming"]
}'
```

### Update
```bash
curl -X PUT http://localhost:8080/api/v1/posts/{id} \
-H "Content-Type: application/json" \
-d '{
  "title": "My Updated Blog Post",
  "content": "This is the updated content of my first blog post.",
  "category": "Technology",
  "tags": ["Tech", "Programming"]
}'
```

### Delete
```bash
curl -X DELETE http://localhost:8080/api/v1/posts/{id}
```
