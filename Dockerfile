FROM golang:1.22.4-alpine

WORKDIR /rates

COPY . .

RUN go mod init test & go mod tidy

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN GRPC_HEALTH_PROBE_VERSION=v0.4.28 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

RUN go build -C /rates/cmd -o main

EXPOSE ${APP_PORT}
