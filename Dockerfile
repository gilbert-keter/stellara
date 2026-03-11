FROM golang:1.25.0-alpine As build

WORKDIR /app/

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o stellara ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/stellara .
COPY .env .

EXPOSE 8080

CMD ["./stellara"]