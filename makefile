# 用于非GoLand环境的编译
# GOPATH=$(pwd) go build -o bin/pgo-demo src/Main/main.go

AppName:=pgo-demo
GOPATH:=$(shell pwd)

.PHONY: build start stop

build:
	GOPATH=${GOPATH} go build -o bin/${AppName} src/Main/main.go

start:
	bin/$(AppName)


