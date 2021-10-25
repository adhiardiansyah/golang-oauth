FROM golang:1.17

WORKDIR /app

COPY go.mod ./
COPY *.go ./
COPY auth ./auth
COPY handler ./handler
COPY helper ./helper
COPY templates ./templates
COPY user ./user

RUN go build -o /golang-oauth

EXPOSE 3000

CMD ["/golang-oauth"]