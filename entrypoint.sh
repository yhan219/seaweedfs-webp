#!/bin/bash
ARGS=''
if [[ -n "$volumeServer" ]]; then
	ARGS="-volumeServer=$volumeServer"
fi
exec go run src/main.go $@ ${ARGS}