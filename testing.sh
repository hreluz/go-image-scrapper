#!/bin/sh

for d in $(go list ./...); do
	echo "Testing package $d"
	go test -v $d
done
