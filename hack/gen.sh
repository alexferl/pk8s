#!/usr/bin/env bash

# Check if the version is provided as an argument
if [ -z "$1" ]; then
  echo "Usage: $0 <K8S_VERSION>"
  exit 1
fi

K8S_VERSION=$1

rm -rf k8s && \
openapi-generator generate \
    -i https://raw.githubusercontent.com/kubernetes/kubernetes/${K8S_VERSION}/api/openapi-spec/swagger.json \
    -g go \
    -o ./k8s \
    --generate-alias-as-model \
    --model-package models \
    --global-property models,modelDocs=false,supportingFiles=utils.go \
    --template-dir ./hack/custom-templates \
    --package-name k8s && \
go run hack/rename_k8s_types.go && \
go run hack/gen_api_version_and_kind.go -version ${K8S_VERSION} && \
goimports -w .
