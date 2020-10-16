FROM node:14.9.0-alpine3.12
LABEL maintainer="weslenng"

WORKDIR /dorothy/faker-http
COPY faker-http/package.json ./
COPY faker-http/package-lock.json ./
RUN npm install

COPY faker-http/. .

CMD ["npm", "run", "start"]
