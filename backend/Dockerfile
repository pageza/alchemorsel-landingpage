# Build stage: compile the backend Go binary.
FROM golang:1.18-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod /go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final image: install PostgreSQL, supervisord, and nginx.
FROM alpine:latest
RUN apk add --no-cache postgresql postgresql-client supervisor nginx

WORKDIR /app

# Copy the built Go binary.
COPY --from=builder /app/main .

# Copy backend config & scripts.
RUN mkdir -p /run/postgresql && chown postgres:postgres /run/postgresql
COPY entrypoint.sh /app/entrypoint.sh
COPY supervisord.conf /app/supervisord.conf
# Ensure the .env file is copied
COPY .env /app/.env
COPY database.sql /app/database.sql
RUN chmod +x /app/entrypoint.sh

# Expose ports.
EXPOSE 8080
EXPOSE 80

# Set environment variables for database connectivity. Use the container name for PostgreSQL.
ENV PG_HOST=localhost
ENV PG_PORT=5432
ENV PG_USER=postgres
ENV PG_PASSWORD=postgres
ENV PG_DB=alchemorsel_db

ENTRYPOINT ["/app/entrypoint.sh"]
