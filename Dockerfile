FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/default/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/public ./public

EXPOSE 8080
ENV PORT=8080
CMD [ "/app/main" ]
