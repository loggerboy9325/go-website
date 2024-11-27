# Stage 1: Build the Go application with Node.js and npm
FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk update && apk add --no-cache make nodejs npm

COPY . ./
RUN make install
RUN make build
RUN > /app/.env

# Stage 2: Create a minimal image with only the binary and environment
FROM scratch

COPY --from=builder /website_go-htmx-temple /website_go-htmx-temple
COPY --from=builder /app/.env .env

EXPOSE 3000

ENTRYPOINT ["./website_go-htmx-temple"]

