FROM docker.io/library/golang:1.24.4 AS build

ENV CGO_ENABLED=0\
    GOOS=linux\
    GOARCH=amd64

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o blogging-platform-api ./cmd

FROM docker.io/library/alpine:latest

WORKDIR /root/

COPY --from=build /src .

ENV MYSQL_DB_USER=
ENV MYSQL_DB_PASSWD=
ENV MYSQL_DB_NAME=
ENV MYSQL_DB_ADDR=

EXPOSE 8080

CMD ["./blogging-platform-api"]
