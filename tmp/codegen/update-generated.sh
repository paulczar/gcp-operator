#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

vendor/k8s.io/code-generator/generate-groups.sh \
deepcopy \
github.com/paulczar/gcp-operator/pkg/generated \
github.com/paulczar/gcp-operator/pkg/apis \
cloud:v1alpha1 \
--go-header-file "./tmp/codegen/boilerplate.go.txt"
