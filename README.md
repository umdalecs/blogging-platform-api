# Go net/http Blogging Platform API

This is my golang solution to https://roadmap.sh/projects/blogging-platform-api, i tried to minimize dependencies using only the `net/http` standard library.

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

I hardcode some default values in `config/configuration.go` file but
can be easily overriden copying the .env example and populating it with your data

```bash
cp .env.example .env
```

### 4. Build and run the application

```bash
go build -o out/blogging-platform-api ./cmd
./out/blogging-platform-api
```
