#!/usr/bin/env bash

#export uname=${USER_NAME:-}
#export gname=${GROUP_NAME:-${USER_NAME:-}}
#export home_dir=${USER_HOME:-/home/$uname}
# change to use arguments instead of environment variables
export uname=$1
export gname=$2
export home_dir=$3

# this is how to use the arguments
# ./set_env.sh user group /home/user

# Change user ID, group ID and user name if a USER_NAME is supplied, this ensures that the ownership
# of bind mounted volumes is the same as on the host machine.
# This is needed for Linux environments or WSL 2
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
        rsync sudo unzip postgresql-client && \
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