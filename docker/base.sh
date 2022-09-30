#! /usr/bin/env bash
set -eu

# Logs information in a format that matches the Python logging output
# Parameters:
# log_level: INFO, ERROR, CRITICAL or DEBUG
# ...: Log message
logger() {
    local dt=$(date +"%Y-%m-%d %H:%M:%S %z")
    local pid=$(echo $$)
    local level=$1
    local logname="app"
    echo "[${dt}] [${pid}] [${level}] [${logname}]" "${@:2}"
}

logger_prefix() {
    local dt=$(date +"%Y-%m-%d %H:%M:%S %z")
    local pid=$(echo $$)
    local level=$1
    local logname="app"
    echo -n "[${dt}] [${pid}] [${level}] [${logname}]" ""
}

log() {
    logger INFO "$@"
}

log_error() {
    logger ERROR "$@"
}

log_critical() {
    logger CRITICAL "$@"
}

log_debug() {
    logger DEBUG "$@"
}