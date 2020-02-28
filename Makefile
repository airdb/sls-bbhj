SRC = $(shell go env GOPATH)
PWD = $(shell pwd)

GITHUB = $(shell echo $(PWD) | sed "s@$(SRC)/src/@@")
REPO = $(shell echo $(GITHUB) | sed "s@github.com/@@")

all:
	echo $(GITHUB)
	echo $(REPO)
	docker build -t $(REPO) . --build-arg GITHUB=$(GITHUB) --build-arg BUILDDIR=/go/src/$(GITHUB)

fmt:
	gofmt -s -w .

push:
	docker push $(REPO)
	gofmt -s -w .

exec bash:
	docker run -it $(REPO) bash
