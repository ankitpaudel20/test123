FROM golang:1.24.0-alpine3.20 AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev make

COPY go.mod go.sum Makefile ./
RUN make tidy

COPY . .

# RUN CGO_ENABLED=0 GOOS=linux go build -o /app/final_app ./cmd/web
RUN make build

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/final_app .

COPY --from=builder /app/assets ./assets

# Run the application
ENTRYPOINT [ "make", "run" ]