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

function main() {
    initIndices
}

main "$@"
