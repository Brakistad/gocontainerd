# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.245.2/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.19, 1.18, 1-bullseye, 1.19-bullseye, 1.18-bullseye, 1-buster, 1.19-buster, 1.18-buster
ARG VARIANT="1.19-bullseye"
FROM golang:1.19.1-buster as base

# install curl
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends curl

WORKDIR /app

FROM base as dev

ARG USER_UID
ARG USER_GID
ARG USER_NAME=''
ARG USER_HOME=''
ARG USER_SHELL=''
ARG DOCKER_GID=''

ENV USER_UID=$USER_UID
ENV USER_GID=$USER_GID

# Change user ID, group ID and user name if a USER_NAME is supplied, this ensures that the ownership
# of bind mounted volumes is the same as on the host machine.
# This is needed for Linux environments or WSL 2
RUN export uname=${USER_NAME:-dev} && \
    export gname=${GROUP_NAME:-${USER_NAME:-dev}} && \
    export home_dir=${USER_HOME:-/home/$uname} && \
    if [ -n "$USER_UID" ] || [ -n "$USER_GID" ] || [ "$uname" != "root" ]; then \
        if [ "$uname" = "root" ]; then \
            echo "Cannot change uid/gid when USER_NAME is root" >&2 && \
            exit 1; \
        fi && \
        echo "Creating new user $uname uid=${USER_UID-1000} gid=${GROUP_UID-1000}" && \
        groupadd -g ${USER_GID:-1000} $gname && \
        useradd -m -d "/home/$uname" -g ${gname} -G sudo -u ${USER_UID-1000} ${uname}; \
    fi && \
    if [ -n "${USER_SHELL}" ]; then \
        echo "Changing default shell to ${USER_SHELL} for user $uname" && \
        chsh -s "${USER_SHELL}" $uname; \
    fi

# Install a good environment for development, including extra packages
# Force /etc/profile to run for bash without requirement --login option
# iproute2: Provides `ip` command
RUN export uname=${USER_NAME:-dev} && \
    apt-get update --allow-releaseinfo-change && \
    apt-get install -yq \
        iproute2 \
        bash-completion git vim curl gnupg2 nano less ripgrep wget procps lsof \
        apt-transport-https ca-certificates software-properties-common \
        rsync sudo unzip && \
    mkdir -p /etc/bash_completion.d && \
    echo ". /etc/profile" >> $HOME/.bashrc && \
    if [ "$uname" != "root" ]; then \
        # Allow sudo to be used without password \
        echo '%sudo ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers && \
        adduser $uname sudo; \
    fi

# Install docker inside container, for development usage
COPY --from=library/docker:latest /usr/local/bin/docker /usr/bin/
RUN export docker_gid=${DOCKER_GID:-} && \
    groupadd -g ${docker_gid:-1001} docker && \
    usermod -a -G docker ${USER_NAME:-dev}

# install docker-compose inside container, for development usage
RUN curl -L \
    "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" \
    -o /usr/local/bin/docker-compose \
    && chmod +x /usr/local/bin/docker-compose

COPY docker /docker

COPY main.go /app/main.go
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum

RUN go mod download

RUN chmod +x /docker/entrypoint.sh

RUN chmod +x /docker/start.sh

ENV PATH=$PATH:${USER_HOME:-/home/${USER_NAME:-dev}}/.local/bin:${USER_HOME:-/home/${USER_NAME:-dev}}/.yarn/bin

# Change owner of $GOPATH to user
RUN chown -R ${USER_NAME:-dev}:${GROUP_NAME:-${USER_NAME:-dev}} $GOPATH

# Switch to non-root user
USER ${USER_NAME:-dev}

# Installing Go tools
RUN go install -v golang.org/x/tools/gopls@latest \
    && go install -v github.com/go-delve/delve/cmd/dlv@latest \
    && go install -v honnef.co/go/tools/cmd/staticcheck@latest

ENTRYPOINT ["/docker/entrypoint.sh"]