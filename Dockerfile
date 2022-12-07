
FROM golang:1.19.3-alpine

RUN mkdir /app

WORKDIR /app

COPY go.mod ./ \ && go.sum ./
RUN go mod download
COPY . . 


FROM golang:1.19.3-alpine

RUN go build -o /better-schedule cmd/main.go

EXPOSE 8080

CMD [ "/better-schedule" ]
