.PHONY: build build-all clean vet fmt run

GOCMD=GO111MODULE=on go
BINARY=bin/`basename $(CURDIR)`

# Build info
GO_VERSION=`go version | sed 's/go version //g'`
BUILD_GIT_HASH=`git rev-parse --short HEAD`
BUILD_TIME=`date '+%Y-%m-%d %H:%M:%S'`
BUILD_LDFLAGS=

build:
	@echo "build go binary..."
	${GOCMD} build -gcflags '-N -l' -ldflags="${BUILD_LDFLAGS}" -o ${BINARY} ${GOARGS} server/main.go
	@echo "build go binary done"

build-all:
	@echo "build all go files..."
	${GOCMD} build ${GOARGS} ./...
	@echo "build all go files done"

clean:
	@echo "clean go binary..."
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	@echo "clean go binary done"

vet:
	@echo "vet go files..."
	${GOCMD} vet ${GOARGS} ./...
	@echo "vet go files done"

fmt:
	@echo "fmt go files..."
	${GOCMD} fmt ${GOARGS} ./...
	@echo "fmt go files done"

run:
	@echo "run server..."
	${GOCMD} run ${GOARGS} server/main.go -f etc/config.yaml || true
