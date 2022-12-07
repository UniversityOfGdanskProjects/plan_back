FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod ./ 
COPY go.sum ./
RUN go mod download
COPY . . 


FROM golang:1.19.3-alpine

RUN go build -o /better-schedule cmd/main.go

FROM alpine
WORKDIR /app

COPY --from=builder /app/better-schedule /app/better-schedule

CMD ["./better-schedule"]
