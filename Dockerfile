FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main ./cmd/main

FROM scratch

WORKDIR /app

COPY --from=builder /app/main /usr/bin/

ENTRYPOINT ["main"]
