FROM golang:1.15

WORKDIR /go/src/app
COPY . .
RUN go get -u
RUN go build -o main