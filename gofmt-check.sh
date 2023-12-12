#!/bin/bash

BAD_FILES="$(gofmt -l . )"
if [[ ! -z "$BAD_FILES" ]]; then
    echo "$BAD_FILES"
    exit 1
fi

