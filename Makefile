DATE=$(shell date)
GOVERSION=$(shell go version | awk '{print $$3}')
BUILDVERSION=$(shell git describe --tags | awk '{print $$1}')
GOARCH=amd64

build: linux darwin sha tar

linux:
	go build -ldflags="-X 'main.BuildVersion=${BUILDVERSION}' -X 'main.BuildDate=${DATE}' -X 'main.Platform=linux/${GOARCH}' -X 'main.GoVersion=${GOVERSION}'" -o kubectl_nodeinfo_linux_${GOARCH}/kubectl-nodeinfo

darwin:
	go build -ldflags="-X 'main.BuildVersion=${BUILDVERSION}' -X 'main.BuildDate=${DATE}' -X 'main.Platform=darwin/${GOARCH}' -X 'main.GoVersion=${GOVERSION}'" -o kubectl_nodeinfo_darwin_${GOARCH}/kubectl-nodeinfo

sha:
	echo $(shell openssl sha256 < kubectl_nodeinfo_linux_${GOARCH}/kubectl-nodeinfo | awk '{print $$2}')	kubectl_nodeinfo_linux_${GOARCH} > kubectl_nodeinfo_linux_${GOARCH}.sha256
	echo $(shell openssl sha256 < kubectl_nodeinfo_darwin_${GOARCH}/kubectl-nodeinfo | awk '{print $$2}')	kubectl_nodeinfo_darwin_${GOARCH} >> kubectl_nodeinfo_darwin_${GOARCH}.sha256

tar:
	tar czf kubectl_nodeinfo_linux_${GOARCH}.tar.gz kubectl_nodeinfo_linux_${GOARCH}/kubectl-nodeinfo
	tar czf kubectl_nodeinfo_darwin_${GOARCH}.tar.gz kubectl_nodeinfo_darwin_${GOARCH}/kubectl-nodeinfo
