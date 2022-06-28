#!/usr/bin/env bash
set -eo pipefail

export BUNDLE_IMG=icr.io/continuous-delivery/pipeline/olm/iks/tekton-operator-bundle:0.60.0
export CATALOG_IMG=icr.io/continuous-delivery/pipeline/olm/iks/tekton-operator-index:1.0.0

docker build -f bundle.Dockerfile -t ${BUNDLE_IMG} .
docker push $BUNDLE_IMG
operator-sdk bundle validate $BUNDLE_IMG

rm -rf catalog
mkdir -p catalog
cp catalog-common/common.yaml catalog/operator.yaml
opm render $BUNDLE_IMG --output yaml >> catalog/operator.yaml
opm validate catalog

docker build . -f catalog.Dockerfile -t $CATALOG_IMG
docker push $CATALOG_IMG