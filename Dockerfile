FROM golang:1.22.2-alpine AS builder
WORKDIR /app
COPY . .
RUN go build main.go


FROM alpine:3.19.1 as runner
WORKDIR /app
COPY --from=0 /app/main .
COPY ./.env.compose .env.local
CMD ["/app/main"]