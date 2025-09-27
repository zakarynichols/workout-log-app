# Stage 1: Build frontend with Node
FROM node:22.18.0-alpine AS build
WORKDIR /app
COPY ./frontend/package.json ./frontend/package-lock.json* ./
RUN npm ci
COPY ./frontend .
RUN npm run build

# Stage 2: Nginx to serve frontend + proxy API
FROM nginx:stable-alpine
# Clean default html
RUN rm -rf /usr/share/nginx/html/*
# Copy built frontend
COPY --from=build /app/dist /usr/share/nginx/html
# Copy our custom nginx.conf
COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
