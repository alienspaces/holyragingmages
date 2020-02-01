#!/usr/bin/env bash

COMMAND=$1

echo "=> (entrypoint) Command ${COMMAND}"

if [ "$COMMAND" == "run" ]; then

    # run
    echo "=> (entrypoint) Executing run command"

else

    # user command
    echo "=> (entrypoint) Executing user command $*"

    exec "$@"
fi
