FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
COPY auth ./auth
COPY handler ./handler
COPY helper ./helper
COPY templates ./templates
COPY user ./user

RUN go build -o /golang-oauth

EXPOSE 3000

CMD ["/golang-oauth"]