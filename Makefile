BIN = terraform-provider-iris
LDFLAGS += -X main.version=$$(git describe --always --abbrev=40 --dirty)

all: build

.PHONY: build
build: mod fmt lint vet $(BIN)

.PHONY: mod
mod:
	go mod download
	go mod tidy

.PHONY: fmt
fmt:
	go fmt ./iris .

.PHONY: lint
lint:
	go run golang.org/x/lint/golint -set_exit_status ./iris .

.PHONY: vet
vet:
	go vet ./iris .

.PHONY: $(BIN)
$(BIN):
	go build -ldflags "${LDFLAGS}"

.PHONY: install
install:
	go install -ldflags "${LDFLAGS}"

.PHONY: clean
clean:
	rm -f $(BIN)
