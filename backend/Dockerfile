# cursor--Dockerize: Multi-stage Dockerfile for building a self-contained landing page image.

# Build stage: compile the backend Go binary.
FROM golang:1.18-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final image: install PostgreSQL, supervisord, and nginx; copy static assets.
FROM alpine:latest
# cursor--Add nginx installation along with PostgreSQL, its client, and supervisor.
RUN apk add --no-cache postgresql postgresql-client supervisor nginx

WORKDIR /app

# Copy the built Go binary.
COPY --from=builder /app/main .

# Copy static files into nginx's html folder.
COPY frontend/index.html /usr/share/nginx/html/index.html
COPY frontend/main.js /usr/share/nginx/html/main.js
COPY frontend/styles.css /usr/share/nginx/html/styles.css

# Copy entrypoint script and supervisord configuration.
COPY backend/entrypoint.sh /app/entrypoint.sh
COPY backend/supervisord.conf /app/supervisord.conf
RUN chmod +x /app/entrypoint.sh

# Expose port 8080 for the backend and 80 for nginx.
EXPOSE 8080
EXPOSE 80

# Set environment variables for database connectivity.
ENV DB_HOST=127.0.0.1
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD=postgres
ENV DB_NAME=alchemorsel_db

ENTRYPOINT ["/app/entrypoint.sh"] 