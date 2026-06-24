FROM golang:1.26.1

WORKDIR /app

COPY go.mod go.sum ./

# Download all modules
RUN go mod download

COPY . .

# Build the application
RUN go build -o main ./cmd/server

CMD ["server"]