#!/bin/bash

CHECK_JQ=$(which jq)

if [ -z "$CHECK_JQ" ]; then
    echo "Please install jq (mac: brew install jq ;  ubuntu/debian: sudo apt install jq; )"
    exit 1
fi

time curl http://localhost:9191/api/v1/version | jq .
time curl http://localhost:9191/api/v1/callsomething | jq .
