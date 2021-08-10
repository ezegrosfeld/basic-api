FROM golang:1.16-alpine as builder

WORKDIR /app 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o api cmd/api/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/api /usr/bin/

ENTRYPOINT ["api"]