#!/bin/sh

docker run -p 8080:8080 -v "${PWD}/src:/app/src" -e ENVIRONMENT=dev app2 