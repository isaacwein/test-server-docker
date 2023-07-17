FROM golang:latest AS builder
LABEL authors="isaac"

WORKDIR /app

COPY ./ ./

RUN go mod tidy



RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/main


FROM scratch
EXPOSE 8080
WORKDIR /app
COPY --from=builder /app/bin/main .
CMD ["--host=127.0.0.1", "--port=8080"]
ENTRYPOINT ["./main"]