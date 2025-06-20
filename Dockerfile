FROM golang:1.24.4-alpine3.21 AS builder

COPY . /userDomain/source/
WORKDIR /userDomain/source/

RUN go mod download
RUN go build -o ./bin/userServer cmd/http/main.go

FROM alpine:3.21

RUN apk add --no-cache postgresql-client

WORKDIR /root/
COPY --from=builder /userDomain/source/bin/userServer .

COPY wait-for-db.sh .
RUN chmod +x wait-for-db.sh

COPY config/local.yaml ./config/local.yaml
ENV CONFIG_PATH=/root/config/local.yaml

CMD ["./wait-for-db.sh","db", "user", "12345", "userdb", "./userServer"]