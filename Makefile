
##@ Help
#[x]Add help target to Makefile

.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)


##@ Manifest Generation

.PHONY: rbac-orchestrator rbac-agent


##@ Development

##General
.PHONY: crds
crds: controller-gen ##Generate Webhook configuration, ClusterRole and CustomResourceDefinition objects
	$(CONTROLLER_GEN) crd paths="./..." output:crd:artifacts:config=config/crd/bases 

.PHONY: rbac-orchestrator
.PHONY: rbac
rbac: controller-gen ##Generate RBAC configuration
	$(CONTROLLER_GEN) rbac:roleName=mycontroller-role clusterName=cluster-role paths=./... output:dir=./config/rbac


.PHONY: fmt
fmt: ## Run go fmt against code
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code
	go vet ./...

.PHONY: lint
lint: golangci-lint ## Run golangci-lint against code
	$(GOLANGCI_LINT) run

##@ Build

.PHONY: build
build: manifests generate fmt vet ## Build manager binary
	go build -o bin/manager main.go

.PHONY: run
run: manifests generate fmt vet ## Run against the configured Kubernetes cluster in ~/.kube/config
	go run ./main.go

33
##Output directories

RBAC_ORCHESTRATOR_DIR = config/rbac/orchestrator
RBAC_AGENT_DIR = config/rbac/agent


##@ Dependencies

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool Binaries
CONTROLLER_GEN ?= $(LOCALBIN)/controller-gen
ENVTEST ?= $(LOCALBIN)/setup-envtest


## Tool Versions
CONTROLLER_TOOLS_VERSION ?= v0.16.4
ENVTEST_VERSION ?= release-0.19

.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN) ## Download controller-gen locally if necessary.
$(CONTROLLER_GEN): $(LOCALBIN)	
	$(call go-install-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen,$(CONTROLLER_TOOLS_VERSION))

.PHONY: envtest
envtest: $(ENVTEST) ## Download setup-envtest locally if necessary.
$(ENVTEST): $(LOCALBIN)
	$(call go-install-tool,$(ENVTEST),sigs.k8s.io/controller-runtime/tools/setup-envtest,$(ENVTEST_VERSION))


# go-install-tool will 'go install' any package with custom target and name of binary, if it doesn't exist.
# $1 - target path with name of the binary
# $2 - package url wich can be installed
# $3 - specific version of the package
define go-install-tool
@[ -f "$(1)-$(3)" ] || { \
set -e; \
package=$(2)@$(3) ;\
echo "Downloading $${package}" ;\
rm -f $(1) || true ;\
GOBIN=$(LOCALBIN) go install $${package} ;\
mv $(1) $(1)-$(3) ;\
} ;\
ln -sf $(1)-$(3) $(1)
endef