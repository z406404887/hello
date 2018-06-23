## Makefile


.PHONY: setup
setup: ## Install all the build and lint dependencies
	go get -u github.com/alecthomas/gometalinter
	go get -u github.com/golang/dep/cmd/dep
	gometalinter --install --update
	@$(MAKE) dep

.PHONY: dep
dep : ## run dep ensure and prune
	dep ensure
	dep prune

.PHONY: fmt
fmt:
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do goimports -w "$$file"; done

.PHONY: lint
lint: ## Run all the linters
	gometalinter --vendor --disable-all \
	--enable=deadcode \
	--enable=ineffassign \
	--enable=gosimple \
	--enable=staticcheck \
	--enable=gofmt \
	--enable=goimports \
	--enable=misspell \
	--enable=errcheck \
	--enable=vet \
	--enable=vetshadow \
	--deadline=10m \
	./...

.PHONY: build
build: gateway login dbserver manager robot game


.PHONY: gateway 
gateway:
	go build -o ./bin/gateway ./cmd/gateway/main.go 

.PHONY: login
login:
	go build -o ./bin/login ./cmd/login/main.go 

.PHONY: dbserver
dbserver:
	go build -o ./bin/dbserver ./cmd/dbserver/main.go 

.PHONY: manager
manager:
	go build -o ./bin/manager ./cmd/manager/main.go 

.PHONY: game
game:
	go build -o ./bin/game ./cmd/robot/main.go 

.PHONY: robot
robot:
	go build -o ./bin/robot ./cmd/robot/main.go 

.PHONY:clean
clean: ## Remove temporary files
	go clean

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := build
