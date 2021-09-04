FROM node:14.15 as build-step
RUN mkdir -p /app
WORKDIR /app
COPY package.json /app
COPY . /app
RUN npm install -g @ionic/cli
RUN npm install -g @angular/cli
RUN npm install
EXPOSE 8100
CMD ["ng","serve","--host", "0.0.0.0", "--no-open", "--port", "8100"]
