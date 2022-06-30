#!/usr/bin/env bash
set -eo pipefail

export CATALOG_IMG=icr.io/continuous-delivery/pipeline/olm/iks/tekton-operator-index:1.0.0

docker build . -f catalog.Dockerfile -t $CATALOG_IMG
docker push $CATALOG_IMG