# 用于非GoLand环境的编译

AppName:=pgo-demo
GOPATH:=$(shell pwd)

.PHONY: build start stop update init

build:
	GOPATH=${GOPATH} go build -o bin/${AppName} src/Main/main.go

start: build
	bin/$(AppName)

stop:
	killall $(AppName)

update:
	GOPATH=${GOPATH} cd src && glide update

init:
	[]
