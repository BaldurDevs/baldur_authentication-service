FROM alpine:latest

RUN mkdir /app
RUN go build -o authApp ./cmd/main.go

COPY authApp /app

CMD ["app/authApp"]
