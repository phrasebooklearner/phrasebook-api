#!/bin/bash

set -o pipefail; go test -v $TEST_NAMESPACE 2>&1 | go-junit-report | tee $TEST_REPORT_PATH