
FROM golang:1.19.3-alpine3.11.3

RUN mkdir /app

WORKDIR /app

COPY go.mod ./ \ && go.sum ./
RUN go mod download
COPY . . 

FROM alpine:3.11.3
RUN go build -o /better-schedule cmd/main.go

EXPOSE 8080

CMD [ "/better-schedule" ]
