FROM golang:alpine
WORKDIR /app
COPY . .
RUN go build -o main cmd/default/main.go

EXPOSE 8080
ENV PORT=8080
CMD [ "/app/main" ]
