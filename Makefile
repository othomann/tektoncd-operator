.PHONY: all
all: local-dev

.PHONY: clean-k8s
clean-k8s:
	-kubectl delete -f deploy/kubernetes/ --ignore-not-found
	-kubectl delete -f deploy/kubernetes/crds/ --ignore-not-found
	-kubectl delete namespace tekton-pipelines --ignore-not-found

.PHONY: dev-setup-local-k8s
dev-setup-local-k8s:
	kubectl apply -f deploy/kubernetes/crds/operator_v1alpha1_addon_crd.yaml
	kubectl apply -f deploy/kubernetes/crds/operator_v1alpha1_pipeline_crd.yaml

.PHONY: clean-openshift
clean-openshift:
	-kubectl delete -f deploy/openshift/ --ignore-not-found
	-kubectl delete -f deploy/openshift/crds/ --ignore-not-found
	-kubectl delete namespace openshift-pipelines --ignore-not-found

.PHONY: dev-setup-local-openshift
dev-setup-local-openshift:
	kubectl apply -f deploy/openshift/crds/operator_v1alpha1_config_crd.yaml

#.PHONY: local-dev-k8s
#local-dev-k8s: clean-k8s dev-setup-local-k8s
#	GO111MODULE=on operator-sdk run --local --watch-namespace "" --operator-flags '--zap-encoder=console'

.PHONY: local-dev-k8s
local-dev-k8s: clean-k8s dev-setup-local-k8s
	-rm cmd/manager
	sleep 10
	ln -s $(shell pwd)/cmd/kubernetes-manager cmd/manager
	GO111MODULE=on operator-sdk run --local --watch-namespace "" --operator-flags '--zap-encoder=console'

.PHONY: local-dev-openshift
local-dev-openshift: clean-openshift dev-setup-local-openshift
	-rm cmd/manager
	sleep 10
	ln -s $(shell pwd)/cmd/openshift-manager cmd/manager
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
