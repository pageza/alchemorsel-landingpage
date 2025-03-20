#!/bin/sh
set -e

echo "Waiting for /etc/letsencrypt/live/alchemorsel.com/fullchain.pem to appear..."
tries=0
while [ ! -f /etc/letsencrypt/live/alchemorsel.com/fullchain.pem ] && [ $tries -lt 10 ]; do
  echo "Certificate not found. Sleeping..."
  sleep 2
  tries=$((tries+1))
done

if [ ! -f /etc/letsencrypt/live/alchemorsel.com/fullchain.pem ]; then
  echo "Certificate file still not found after waiting. Exiting."
  exit 1
fi

echo "Copying certificate files from /etc/letsencrypt to /etc/nginx/ssl..."
mkdir -p /etc/nginx/ssl
# Copy using -L so that symlinks are followed and the actual file content is copied.
cp -L /etc/letsencrypt/live/alchemorsel.com/fullchain.pem /etc/nginx/ssl/fullchain.pem
cp -L /etc/letsencrypt/live/alchemorsel.com/privkey.pem /etc/nginx/ssl/privkey.pem

echo "Starting Nginx..."
exec nginx -g "daemon off;" 