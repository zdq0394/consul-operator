# ! /usr/bin/sh

ifneq ($(shell uname), Darwin)
    EXTLDFLAGS = -extldflags "-static" $(null)
else
    EXTLDFLAGS =
endif

BUILD_NUMBER=$(shell git rev-parse --short HEAD)

BUILD_DATE=$(shelldate +%FT%T%z)

build:
	mkdir -p make/release
	go build -o make/release/consulops github.com/zdq0394/consul-operator/cmd/operator
	chmod 775 make/release/consulops

run:
	make/release/consulops consul --develop --kubeconfig=/root/.kube/config

clean:
	rm -f make/release/consulops
