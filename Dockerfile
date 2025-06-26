# Create a build image
FROM golang:1.24-alpine AS build

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build

FROM alpine AS release

WORKDIR /app

COPY --from=build /app/mcauth .
COPY config.yml icon.png ./

ENTRYPOINT [ "/app/mcauth" ]