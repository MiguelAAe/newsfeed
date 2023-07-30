FROM golang:1-alpine3.18
WORKDIR /app
COPY . .
RUN go build -o newsfeed cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=0 /app ./
EXPOSE 9200
CMD ["./newsfeed"]