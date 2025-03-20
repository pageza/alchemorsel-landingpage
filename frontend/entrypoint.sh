#!/bin/sh
set -e

echo "Checking for certificate files in /etc/letsencrypt..."
if [ ! -f /etc/letsencrypt/live/alchemorsel.com/fullchain.pem ]; then
  echo "ERROR: Valid CA-signed certificates not found at /etc/letsencrypt."
  echo "Please generate your SSL certificates (using Certbot or another ACME client) on your host,"
  echo "and ensure they are stored at the mounted directory so that /etc/letsencrypt/live/alchemorsel.com/fullchain.pem exists."
  exit 1
fi

echo "Copying certificate files from /etc/letsencrypt to /etc/nginx/ssl..."
mkdir -p /etc/nginx/ssl
# Copy using -L so that symlinks are followed and the actual file content is copied.
cp -L /etc/letsencrypt/live/alchemorsel.com/fullchain.pem /etc/nginx/ssl/fullchain.pem
cp -L /etc/letsencrypt/live/alchemorsel.com/privkey.pem /etc/nginx/ssl/privkey.pem

echo "Starting Nginx..."
exec nginx -g "daemon off;" 