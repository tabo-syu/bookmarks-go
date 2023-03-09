####################### Build stage #######################
FROM golang:1.20.2-alpine3.16 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/lambda_server/main.go

####################### Run stage #######################
FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/main .

COPY root.crt .

EXPOSE 8080

CMD [ "/app/main" ]
