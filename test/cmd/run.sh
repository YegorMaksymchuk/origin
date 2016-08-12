#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

OS_ROOT=$(dirname "${BASH_SOURCE}")/../..
source "${OS_ROOT}/hack/lib/init.sh"
os::log::stacktrace::install
trap os::test::junit::reconcile_output EXIT

os::test::junit::declare_suite_start "cmd/run"
# This test validates the value of --image for oc run
os::cmd::expect_success_and_text 'oc run newdcforimage --image=validimagevalue' 'deploymentconfig "newdcforimage" created'
os::cmd::expect_failure_and_text 'oc run newdcforimage2 --image="InvalidImageValue0192"' 'error: Invalid image name "InvalidImageValue0192": invalid reference format'
echo "oc run: ok"
os::test::junit::declare_suite_end
