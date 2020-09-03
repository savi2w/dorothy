FROM node:14.9.0-alpine3.12
LABEL maintainer="weslenng"

WORKDIR /dorothy/faker-http
COPY faker-http/package.json ./
COPY faker-http/yarn.lock ./
RUN yarn --pure-lockfile

COPY faker-http/. .

CMD ["yarn", "start"]
