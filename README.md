# Go net/http Blogging Platform API

This is my golang solution to https://roadmap.sh/projects/blogging-platform-api.

## Features

- *Mysql Storage*: The app stores posts in mysql database
- *Search filter*: You can easily filter posts using query params

## Running application

### 1. Clone the repo

```bash
git clone https://github.com/umdalecs/blogging-platform-api
cd blogging-platform-api
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Environment variables

```bash
cp .env.example .env
```

> [!Note]
> Don't forget, this proyect depends on mysql storage and database should be initialized

### 4. Build and run the application

```bash
go build -o out/blogging-platform-api ./cmd
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
curl -X POST http://localhost:8080/api/v1/posts -d\
-H "Content-Type: application/json" \
'{\
  "title": "My First Blog Post",\
  "content": "This is the content of my first blog post.",\
  "category": "Technology",\
  "tags": ["Tech", "Programming"]\
}'
```

### Update
```bash
curl -X PUT http://localhost:8080/api/v1/posts/{id} -d\
-H "Content-Type: application/json" \
'{\
  "title": "My Updated Blog Post",\
  "content": "This is the updated content of my first blog post.",\
  "category": "Technology",\
  "tags": ["Tech", "Programming"]\
}'
```

### Delete
```bash
curl -X DELETE http://localhost:8080/api/v1/posts/{id}
```
