# build stage
FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build
EXPOSE 8100
CMD ["yarn","serve","--host", "0.0.0.0", "--no-open", "--port", "8100"]