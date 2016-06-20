#!/bin/bash
set -ev

flags='-short'
if [ "$TRAVIS_TAG" ]; then
    echo "Release being made... Running all tests (including integration)"
    flags=''
fi

go test ./... $flags -v

if [ "$TRAVIS_TAG" ]; then
    docker build -t ddelnano/go-play-racq:$TRAVIS_TAG .
    docker push 
fi
