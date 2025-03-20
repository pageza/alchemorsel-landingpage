#!/bin/sh
set -e

echo "Copying certificate files from /etc/letsencrypt to /etc/nginx/ssl..."
mkdir -p /etc/nginx/ssl
# Copy using -L so that symlinks are followed and the actual file content is copied.
cp -L /etc/letsencrypt/live/alchemorsel.com/fullchain.pem /etc/nginx/ssl/fullchain.pem
cp -L /etc/letsencrypt/live/alchemorsel.com/privkey.pem /etc/nginx/ssl/privkey.pem

echo "Starting Nginx..."
exec nginx -g "daemon off;" 