FROM golang:1.26.1
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application
RUN go build -o /app/server ./cmd/server

CMD ["/app/server"]