#!/usr/bin/env bash

source ${BASH_SOURCE%/*}/start-services || status=$?

if [ $status -ne 0 ]; then
    echo "Failed start-services" >&2
    exit $status
fi


# required for all environments
MODULES=( \
    'common/database' \
    'common/env' \
    'common/logger' \
    'common/model' \
    'common/repository' \
    'common/service' \
);

for MODULE in ${MODULES[@]}; do

    cd $MODULE
    go test -v -count 1 .
    cd -

done
