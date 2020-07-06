.PHONY: all
all: local-dev

.PHONY: clean
clean:
	-kubectl delete -f deploy/ --ignore-not-found
	-kubectl delete -f deploy/crds/ --ignore-not-found
	-kubectl delete namespace tekton-pipelines --ignore-not-found

.PHONY: dev-setup-local
dev-setup-local:
	kubectl apply -f deploy/kubernetes/crds/operator_v1alpha1_addon_crd.yaml
	kubectl apply -f deploy/kubernetes/crds/operator_v1alpha1_pipeline_crd.yaml
	kubectl apply -f deploy/openshift/crds/operator_v1alpha1_config_crd.yaml

.PHONY: local-dev
local-dev: clean dev-setup-local
	GO111MODULE=on operator-sdk run --local --watch-namespace "" --operator-flags '--zap-encoder=console'

.PHONY: update-deps dev-setup
dev-setup:
	kubectl create namespace tekton-pipelines
	kubectl apply -f deploy/kubernetes/crds/operator_v1alpha1_addon_crd.yaml
	kubectl apply -f deploy/kubernetescrds/operator_v1alpha1_pipeline_crd.yaml
	kubectl apply -f deploy/openshift/crds/operator_v1alpha1_config_crd.yaml
	kubectl apply -f deploy/service_account.yaml
	kubectl apply -f deploy/role.yaml
	kubectl apply -f deploy/role_binding.yaml

.PHONY: update-deps
update-deps:
	GO111MODULE=on go mod tidy

.PHONY: local-test-e2e
local-test-e2e:
	GO111MODULE=on \
	operator-sdk test local ./test/e2e  \
	--up-local \
	--watch-namespace "" \
	--operator-namespace operators \
	--debug  \
	--verbose
