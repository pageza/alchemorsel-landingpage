#!/bin/sh
set -e

# cursor--Initialize PostgreSQL data directory if not already initialized.
if [ ! -f /var/lib/postgresql/data/PG_VERSION ]; then
  echo "cursor--Initializing PostgreSQL database..."
  mkdir -p /var/lib/postgresql/data
  chown -R postgres:postgres /var/lib/postgresql/data
  su - postgres -c "initdb -D /var/lib/postgresql/data"
  
  # Start PostgreSQL temporarily in the background for initialization.
  su - postgres -c "pg_ctl -D /var/lib/postgresql/data -w start"
  
  # Set postgres password and create the application database.
  psql -U postgres -c "ALTER USER postgres WITH PASSWORD 'postgres';"
  psql -U postgres -c "CREATE DATABASE alchemorsel_db;"
  
  su - postgres -c "pg_ctl -D /var/lib/postgresql/data -m fast -w stop"
fi

echo "cursor--Applying database migrations..."
psql -U postgres -d alchemorsel_db -f /app/database.sql

# Start supervisord which will run both PostgreSQL and the backend.
exec supervisord -n -c /app/supervisord.conf 
