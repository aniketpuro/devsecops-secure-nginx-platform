SHELL := /bin/bash
IMAGE_NAME ?= devsecops-secure-nginx-platform
DOCKERFILE ?= docker/Dockerfile
VERSION ?= $(shell cat VERSION)
GIT_SHA ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo dev)
IMAGE_TAG ?= $(VERSION)-$(GIT_SHA)
HELM_CHART ?= helm/devsecops-secure-nginx-platform
NAMESPACE ?= devsecops-secure-nginx

.PHONY: help test scan docker-build docker-run sbom helm-template k8s-apply clean

help:
	@echo "Available targets:"
	@echo "  test          Run local tests"
	@echo "  scan          Run local security scans"
	@echo "  docker-build  Build the hardened container image"
	@echo "  docker-run    Run the container locally"
	@echo "  sbom          Generate an SBOM with Syft"
	@echo "  helm-template Render the Helm chart"
	@echo "  k8s-apply     Apply raw Kubernetes manifests"
	@echo "  clean         Remove generated artifacts"

test:
	./scripts/test.sh

scan:
	./scripts/scan.sh

docker-build:
	docker build -f $(DOCKERFILE) -t $(IMAGE_NAME):$(IMAGE_TAG) -t $(IMAGE_NAME):$(GIT_SHA) .

docker-run:
	docker run --rm -p 8080:8080 --read-only --tmpfs /tmp --tmpfs /var/cache/nginx --tmpfs /var/run $(IMAGE_NAME):$(IMAGE_TAG)

sbom:
	mkdir -p sbom
	syft dir:. -o cyclonedx-json=sbom/$(IMAGE_NAME)-$(IMAGE_TAG).cdx.json

helm-template:
	helm template $(IMAGE_NAME) $(HELM_CHART) --namespace $(NAMESPACE)

k8s-apply:
	kubectl apply -f kubernetes/
	kubectl apply -f security/kyverno/

clean:
	rm -rf sbom
