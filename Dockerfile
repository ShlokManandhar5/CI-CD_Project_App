FROM golang:1.26.5-alpine AS builder

WORKDIR /app

COPY main.go .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o go-app main.go

FROM scratch

COPY --from=builder /app/go-app /go-app

EXPOSE 8080

ENTRYPOINT ["/go-app"]
