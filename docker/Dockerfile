FROM golang:1.15.0-alpine3.12
LABEL maintainer="weslenng"

WORKDIR /dorothy
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build

CMD ["./dorothy"]
