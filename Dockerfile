FROM golang:1.19.3

RUN mkdir /app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . . 


RUN go build -o /better-schedule cmd/main.go

EXPOSE 8080

CMD [ "/better-schedule" ]
