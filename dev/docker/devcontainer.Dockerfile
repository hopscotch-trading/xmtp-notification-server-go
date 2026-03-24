FROM golang:1.25

# Install golangci-lint
RUN curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.1.6

# Install buf CLI
RUN go install github.com/bufbuild/buf/cmd/buf@latest

# Add shellcheck and jq
RUN apt-get update && apt-get install -y \
    shellcheck \
    jq
