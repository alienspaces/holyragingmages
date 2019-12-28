#!/usr/bin/env bash

curl localhost:8080/templates

for s in foo bar baz ; do curl -d"{\"s\":\"$s\"}" localhost:8080/templates ; done
