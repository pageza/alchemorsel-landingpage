#!/bin/sh
set -e

# Assume external PostgreSQL service handles database initialization and persistence

# Start supervisord which will run both PostgreSQL and the backend.
exec supervisord -n -c /app/supervisord.conf 
