DATE=$(shell date)
GOVERSION=$(shell go version | awk '{print $$3}')
BUILDVERSION=$(shell git describe --tags | awk '{print $$1}')
TAG=$(shell cat .version)
GOARCH= \
				amd64 \
				arm64

build:  linux darwin tar sha

linux: linux-amd64 linux-arm64

darwin: darwin-amd64 darwin-arm64

sha: sha-amd64 sha-arm64

tar: tar-amd64 tar-arm64

linux-%:
	go build -ldflags="-X 'main.BuildVersion=${BUILDVERSION}' -X 'main.BuildDate=${DATE}' -X 'main.Platform=linux/$*' -X 'main.GoVersion=${GOVERSION}'" -o kubectl_nodeinfo_linux_$*/kubectl-nodeinfo

darwin-%:
	go build -ldflags="-X 'main.BuildVersion=${BUILDVERSION}' -X 'main.BuildDate=${DATE}' -X 'main.Platform=darwin/$*' -X 'main.GoVersion=${GOVERSION}'" -o kubectl_nodeinfo_darwin_$*/kubectl-nodeinfo

sha-%:
	echo $(shell openssl sha256 < kubectl_nodeinfo_linux_$*.tar.gz | awk '{print $$2}')	kubectl_nodeinfo_linux_$* > kubectl_nodeinfo_linux_$*.sha256
	echo $(shell openssl sha256 < kubectl_nodeinfo_darwin_$*.tar.gz | awk '{print $$2}')	kubectl_nodeinfo_darwin_$* > kubectl_nodeinfo_darwin_$*.sha256

tar-%:
	cp LICENSE kubectl_nodeinfo_linux_$* && cp LICENSE kubectl_nodeinfo_darwin_$*
	cd kubectl_nodeinfo_linux_$* && tar czf kubectl_nodeinfo_linux_$*.tar.gz kubectl-nodeinfo LICENSE && mv kubectl_nodeinfo_linux_$*.tar.gz ../
	cd kubectl_nodeinfo_darwin_$* && tar czf kubectl_nodeinfo_darwin_$*.tar.gz kubectl-nodeinfo LICENSE && mv kubectl_nodeinfo_darwin_$*.tar.gz ../

clean:
	@rm -rf kubectl_nodeinfo*

release:
	gh release create ${TAG} kubectl_nodeinfo*.tar.gz kubectl_nodeinfo*sha256 -F CHANGELOG/CHANGELOG-${TAG}.md

tag:
	git tag -a ${TAG} -m "version ${TAG}"
	git push origin ${TAG}
