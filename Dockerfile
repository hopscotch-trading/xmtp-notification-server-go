# BUILD IMAGE --------------------------------------------------------
FROM golang:1.24-alpine AS builder

# Get build tools and required header files
RUN apk add --no-cache build-base

WORKDIR /app
COPY . .

ARG GIT_COMMIT=unknown
ARG XMTP_GO_CLIENT_VERSION=unknown
RUN go build \
    -o bin/notifications-server \
    -ldflags="-X 'main.GitCommit=$GIT_COMMIT' -X 'main.XMTPGoClientVersion=$XMTP_GO_CLIENT_VERSION'" \
    cmd/server/main.go

# ACTUAL IMAGE -------------------------------------------------------

FROM alpine:3

LABEL maintainer="engineering@xmtp.com"
LABEL source="https://github.com/xmtp/example-notification-server-go"
LABEL description="XMTP Example Notification Server"

# color, nocolor, json
ENV GOLOG_LOG_FMT=nocolor
RUN addgroup --system --gid 1001 go
RUN adduser --system --uid 1001 -G go go

COPY --from=builder --chown=go:go /app/bin/notifications-server /usr/bin/

USER go

EXPOSE 5556

ENTRYPOINT ["/usr/bin/notifications-server"]

CMD ["--help"]
