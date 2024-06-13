.PHONY: all
all: build

.PHONY: build
build:
	@mkdir -p build
	@go build -o build/server ./cmd/server

.PHONY: clean
clean:
	@rm -rf build

.PHONY: dev
dev:
	@air

.PHONY: run
run:
	@go run ./cmd/server

.PHONY: fmt
fmt:
	@goimports-reviser -format -rm-unused -recursive -use-cache ./...

.PHONY: test
test:
	@go test -v ./{cmd,server}/... $(ARGS)

.PHONY: vet
vet:
	@go vet -v ./{cmd,server}/...

.PHONY: docs
docs:
	@godoc -http=:3029 -play

.PHONY: update
update:
	@go get -u -t -v ./{cmd,server}/...
	@make tidy

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: tools
tools:
	@go install github.com/cosmtrek/air@latest
	@go install github.com/incu6us/goimports-reviser/v3@latest
	@go install golang.org/x/tools/cmd/godoc@latest
