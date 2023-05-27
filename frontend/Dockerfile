# build stage
# FROM node:lts-alpine as build-stage
# WORKDIR /app
# COPY package*.json ./
# RUN npm install
# COPY . .
# EXPOSE 8081
# RUN npm run dev

FROM node:lts-alpine
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 8081
CMD ["npm", "run", "dev"]

# production stage
# FROM nginx:stable-alpine as production-stage
# COPY --from=build-stage /app/dist /usr/share/nginx/html
# CMD ["nginx", "-g", "daemon off;"]

# FROM node:16-slim

# WORKDIR /app

# RUN apt-get update

# COPY ./package*.json /app/
# RUN npm install

COPY . /app
EXPOSE 8081