FROM golang:1.15.0-alpine3.12
LABEL maintainer="Weslen Guerreiro <savi2w@icloud.com>"

WORKDIR /dorothy
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build

CMD ["./dorothy"]
