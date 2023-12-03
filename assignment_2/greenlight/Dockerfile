# syntax=docker/dockerfile:1

FROM golang:1.20

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
RUN go get -v -u github.com/golang-migrate/migrate/v4

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . ./

# Copy the wait-for-it script
COPY scripts/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Build
RUN go build -o main ./cmd/api/ && chmod +x bin

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 4000
EXPOSE 5432

# Run
CMD ["migrate -path ./migrations -database postgres://greenlight:olzhas@db:5432/greenlight?sslmode=disable up"]
CMD ["./main"]