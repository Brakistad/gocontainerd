# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.245.2/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.19, 1.18, 1-bullseye, 1.19-bullseye, 1.18-bullseye, 1-buster, 1.19-buster, 1.18-buster
ARG VARIANT="1.19-bullseye"
FROM golang:1.19.1-buster

# install curl
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends curl

WORKDIR /app

COPY docker /docker

COPY main.go /app/main.go
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum

RUN go mod download

RUN chmod +x /docker/entrypoint.sh

RUN chmod +x /docker/start.sh

ENTRYPOINT ["/docker/entrypoint.sh"]

CMD ["start_app"]