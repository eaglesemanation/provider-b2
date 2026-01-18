#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
PROVIDER_DIR=$( readlink -f "$SCRIPT_DIR/../.." )

if [ ! -f "$PROVIDER_DIR/examples/namespaced/providerconfig/secret.yaml" ]; then
  >&2 echo "Missing examples/namespaced/providerconfig/secret.yaml, populate with valid B2 credentials for testing"
  exit 1
fi

${KUBECTL} -n crossplane-system apply -f "$PROVIDER_DIR/examples/namespaced/providerconfig/secret.yaml"
${KUBECTL} -n crossplane-system apply -f "$PROVIDER_DIR/examples/namespaced/providerconfig/providerconfig.yaml"
${KUBECTL} -n crossplane-system wait --for=condition=Available deployment --all --timeout=5m
