#!/bin/bash -x

errcount=0

error_handler () {
    echo "Trapped error - ${1:-"Unknown Error"}" 1>&2
    (( errcount++ ))       # or (( errcount += $? ))
}

trap error_handler ERR

git checkout -b build
git remote set-url self git@github.com:ozym/test.git
openssl aes-256-cbc -k "$travis_key_password" -d -md sha256 -a -in travis_key.enc -out travis_key
echo "Host github.com" > ~/.ssh/config
echo "  IdentityFile  $(pwd)/travis_key" >> ~/.ssh/config
chmod 400 travis_key
echo "github.com ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEAq2A7hRGmdnm9tUDbO9IDSwBK6TbQa+PXYPCPy6rbTrTtw7PHkccKrpp0yVhp5HdEIcKr6pLlVDBfOLX9QUsyCOV0wzfjIJNlGEYsdlLJizHhbn2mUjvSAHQqZETYP81eFzLQNnPHt4EVVUh7VfDESU84KezmD5QlWpXLmvU31/yMf+Se8xhHTvKSCZIFImWwoG6mbUoWf9nzpIoaSjB+weqqUUmpaaasXVal72J+UX2B+2RPW3RcT0eOzQgqlJL3RKrTJvdsjE3JEAvGq3lGHSZXy28G3skua2SmVi/w4yCE6gbODqnTWlg7+wC604ydGXA8VJiS5ap43JXiUFFAaQ==" > ~/.ssh/known_hosts

go get github.com/GeoNet/delta/meta
go get gopkg.in/yaml.v2

mkdir -p files/ntripcaster
go build ./code/cmd/ntripcaster-config || exit 255

./ntripcaster-config -base testdata -input testdata -output files/ntripcaster/test.yaml

git add files/ntripcaster/test.yaml
git commit -m 'auto update [skip travis]'
git remote -v 
git push self build:master

exit $errcount

# vim: tabstop=4 expandtab shiftwidth=4 softtabstop=4
