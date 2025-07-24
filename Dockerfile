FROM golang:1.24-alpine AS builder

# temp comment
# RUN apk add --no-cache build-base  

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /app 

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]