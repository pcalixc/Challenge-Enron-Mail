# Golang image
FROM golang:1.21.1-alpine

# Env vars
ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GO_ENV=production

## Project setup

# Create app directory inside container
RUN mkdir app

# cd /app
WORKDIR /app

COPY go.mod go.sum ./
#RUN go mod tidy


# Download all dependencies

RUN go mod download


# Copies all files from local to container
COPY . .

# Build
RUN go build -o main .

# Run
CMD ["/app/main"]

