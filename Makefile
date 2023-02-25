WHAT ?= app

DOCKERHUB ?= dkr.xargs.dev

OCI_BUILDER ?= docker
OCI_REPO ?= oci.xargs.dev/handroll
OCI_TAG ?= $(shell bin/teval query --values ./apps/${WHAT}/values.yaml .images.handroll.tag)

bin/teval:
	go build -o bin/teval ./teval

.PHONY: dist
dist: bin/teval
dist:
	rm -rf dist/${WHAT}
	bin/teval run \
		--values ./apps/${WHAT}/values.yaml \
		--dist dist/${WHAT} \
		oci/Dockerfile kubernetes/*.yaml
	rsync --archive ./apps/${WHAT}/overlay/ dist/${WHAT} 2>/dev/null || true

.PHONY: oci
oci: dist
oci:
	${OCI_BUILDER} build \
		--tag ${OCI_REPO}/${WHAT}:${OCI_TAG} \
		--build-arg DOCKERHUB=${DOCKERHUB} \
		--file ./dist/${WHAT}/oci/Dockerfile \
		--push ./oci

.PHONY: deploy
deploy: dist
deploy:
	kubectl apply -f ./dist/${WHAT}/kubernetes/
