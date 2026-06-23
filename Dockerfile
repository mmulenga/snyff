FROM golang:1.26.1

WORKDIR /app

COPY . .

# Download all modules
RUN go mod download

# Build the application
RUN go build -o main ./cmd/server

CMD ["server"]