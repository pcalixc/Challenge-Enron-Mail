FROM golang:1.21.5-alpine

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

COPY go.mod .

# Download all dependencies
RUN go mod download

# Declare custom args
ARG ZS_USER \
    ZS_PASSWORD \
    ZS_BASE_URL \
    API_PORT 
    

ENV ZS_USER=${ZS_USER} \
    ZS_PASSWORD=${ZS_PASSWORD} \
    ZS_BASE_URL=${ZS_BASE_URL} \
    API_PORT=/${API_PORT}

# Copies all files from local to container
COPY . .

# Build
RUN go build -o api .

# Run
CMD ["/app/api"]