#!/bin/bash

# Install application test dependencies
# All the dependencies are installed into /go/src/phrasebook-api/vendor folder
glide install

# Coping the dependencies to a volume
cp -r /go/src/phrasebook-api/vendor/* /vendor

# Run test.sh arguments
exec $@