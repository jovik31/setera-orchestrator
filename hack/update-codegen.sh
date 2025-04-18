#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail


echo "$(dirname $BASH_SOURCE)"
# get the root directory for the project
SCRIPT_ROOT=$(realpath "$(dirname ${BASH_SOURCE[0]})/..")
echo $SCRIPT_ROOT

# get the codegen package from the vendor directory
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}
echo $CODEGEN_PKG

source "${CODEGEN_PKG}/kube_codegen.sh"

THIS_PKG="github/setera"

ln -s github
trap "rm github" EXIT

kube::codegen::gen_helpers \
   ${SCRIPT_ROOT}/pkg/api \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
    

kube::codegen::gen_client \
    --with-applyconfig \
    --with-watch \
    --output-dir "${SCRIPT_ROOT}/pkg/generated" \
    --output-pkg "${THIS_PKG}/pkg/generated" \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
    "${SCRIPT_ROOT}/pkg/api"