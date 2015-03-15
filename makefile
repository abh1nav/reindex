version=0.1.0

.PHONY: all

# To cross compile for linux on mac, build go-linux cross compiler first using
# cd /usr/local/go/src
# sudo GOOS=linux GOARCH=amd64 CGO_ENABLED=0 ./make.bash --no-clean

all:
	@echo "make <cmd>"
	@echo ""
	@echo "commands:"
	@echo "  test        - standard go test"
	@echo "  build       - build the dist binary"
	@echo "  clean       - clean the dist build"
	@echo "  install     - standard go install"
	@echo ""
	@echo "  tools       - go gets a bunch of tools for dev"
	@echo "  deps        - pull and setup dependencies"
	@echo "  update_deps - update deps lock file"

build: clean
	@go build ./...
	@mkdir -p ./bin
	@rm -f ./bin/*
	@go build -o ./bin/reindex-$(version).bin main.go

clean:
	@rm -rf ./bin

coverage:
	@go test -cover -v ./...

deps:
	@glock sync -n github.com/abh1nav/reindex < Glockfile

install:
	@go install ./...

test:
	@go test ./... | grep -v "no test files" | sort -r

tools:
	go get github.com/robfig/glock

update_deps:
	@glock save -n github.com/abh1nav/reindex > Glockfile