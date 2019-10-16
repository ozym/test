#!/bin/bash -x

errcount=0

error_handler () {
    echo "Trapped error - ${1:-"Unknown Error"}" 1>&2
    (( errcount++ ))       # or (( errcount += $? ))
}

trap error_handler ERR

go get github.com/GeoNet/delta/meta
go get go build gopkg.in/yaml.v2

mkdir -p files/ntripcaster
go build ./code/cmd/ntripcaster-config || exit 255

./ntripcaster-config -base testdata -input testdata -output files/ntripcaster/test.yaml

exit $errcount

# vim: tabstop=4 expandtab shiftwidth=4 softtabstop=4
