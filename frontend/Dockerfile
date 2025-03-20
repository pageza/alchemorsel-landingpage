# cursor--Dockerize: Dockerfile for serving the static frontend with Nginx.
FROM nginx:alpine
# Copy static files into nginx html folder.
COPY index.html /usr/share/nginx/html/index.html
COPY styles.css /usr/share/nginx/html/styles.css
COPY main.js /usr/share/nginx/html/main.js

EXPOSE 80 