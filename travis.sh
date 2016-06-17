#!/bin/bash

go test ./... -short -v

if [ "$TRAVIS_TAG" ] then;
    echo "This is when a docker image would be built"
else
    echo "This build should not build a docker image"
fi
