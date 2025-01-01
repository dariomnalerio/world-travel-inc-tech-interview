FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux go build -o main ./cmd/api 

FROM scratch

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]