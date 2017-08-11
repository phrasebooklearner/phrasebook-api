#!/bin/bash

# Create a vendor directory if not exists
mkdir -p /go/src/phrasebook-api/vendor

# Copy dependencies from the cache
cp -r /vendor/* /go/src/phrasebook-api/vendor

# Run build.sh arguments
exec $@