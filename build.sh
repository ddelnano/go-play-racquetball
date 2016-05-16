#!/bin/bash

body='{
"request": {
  "branch":"master",
  "config": {
    "script": "go test ./... -v"
  }
}}'

loggedIn=$(travis whoami)

if [ $? -eq 1 ]
then 
    echo "You are not logged in"
    exit 1
fi

token=$(travis token)

curl -s -X POST \
  -H "Content-Type: application/json" \
  -H "Accept: application/json" \
  -H "Travis-API-Version: 3" \
  -H "Authorization: token $token" \
  -d "$body" \
  https://api.travis-ci.org/repo/ddelnano%2Fgo-play-racquetball/requests
