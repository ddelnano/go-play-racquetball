#!/bin/bash

flags = '-short'
if [ "$TRAVIS_TAG" ]; then
    echo "Release being made... Running all tests (including integration)"
    flags = ''
fi

go test ./... $flags -v
