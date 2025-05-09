FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o server .

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 3000
CMD ["./server"]
