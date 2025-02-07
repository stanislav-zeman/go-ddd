FROM golang:latest AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN mage build

FROM alpine:latest AS runner

RUN apk update && apk upgrade

ARG BIN_PATH
ARG CONFIG_PATH

WORKDIR /app

COPY --from=builder /app/${BIN_PATH} service
COPY --from=builder /app/${CONFIG_PATH} config.yaml

ENTRYPOINT [ "service" ]
CMD ["--config", "config.yaml"]
