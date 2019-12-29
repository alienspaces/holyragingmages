#!/usr/bin/env bash

curl localhost:8080/templates

for s in foo bar baz ; do curl -d"{\"test\":\"$s\"}" localhost:8080/templates ; done
