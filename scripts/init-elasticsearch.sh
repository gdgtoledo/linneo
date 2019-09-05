#!/bin/sh

SCRIPTS_DIR=$(dirname "$0")
ROOT_DIR="${SCRIPTS_DIR}/.."

function initIndices() {
    # create index
    curl -X PUT "localhost:9200/plants?pretty"

    # create index mapping
    curl -X PUT "localhost:9200/plants/_mapping?pretty" \
        -H 'Content-Type: application/json' \
        -d @${ROOT_DIR}/mapping.json
}

function run() {
    docker-compose -f ${ROOT_DIR}/docker/docker-compose.yml up -d
}

function main() {
    run

    # wait for elasticsearch to start up
    sleep 60

    initIndices
}

main "$@"
