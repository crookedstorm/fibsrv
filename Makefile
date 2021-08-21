.PHONY: build down up test
MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MKFILE_DIR := $(dir $(MKFILE_PATH))
RELEASE_DIR := ${MKFILE_DIR}bin

RELEASE?=0.0.1

build: fibsrv

fibsrv:
	cd ${MKFILE_DIR} && CGO_ENABLED=0 go build -v -trimpath -ldflags "-s -w" \
	-o ${RELEASE_DIR}/$@ ${MKFILE_DIR}cmd/$@/

docker_build:
	docker build -t crookedstorm/fibsrv:${RELEASE} -f deployments/Dockerfile .

clean:
	@rm -f ${MKFILE_DIR}bin/*

up:
	docker-compose -f deployments/docker-compose.yml up -d

down:
	docker-compose -f deployments/docker-compose.yml down

test:
	go test ./...
