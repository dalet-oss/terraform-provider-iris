BIN = terraform-provider-iris
LDFLAGS += -X main.version=$$(git describe --always --abbrev=40 --dirty)

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: all
all: mod fmt lint vet $(BIN) ; @

.PHONY: get-sdk
get-sdk:
	$Q rm -rf models sdk
	$Q cp -rf ../iris-daemon/models .
	$Q cp -rf ../iris-daemon/pkg/iris/sdk .
	$Q find sdk -name "*.go" -exec sed -i "s%bitbucket.org/ooyalaflex/iris-daemon/models%github.com/dalet-oss/terraform-provider-iris/models%g" {} +
	$Q find sdk -name "*.go" -exec sed -i "s%bitbucket.org/ooyalaflex/iris-daemon/pkg/iris/sdk%github.com/dalet-oss/terraform-provider-iris/sdk%g" {} +

.PHONY: mod
mod: ; $(info $(M) collecting modules…) @
	$Q go mod download
	$Q go mod tidy

.PHONY: fmt
fmt: ; $(info $(M) formatting code…) @
	$Q go fmt ./iris .

.PHONY: lint
lint: ; $(info $(M) running linter…) @
	$Q go run golang.org/x/lint/golint -set_exit_status ./iris .

.PHONY: vet
vet: ; $(info $(M) running vetter…) @
	$Q go vet ./iris .

.PHONY: $(BIN)
$(BIN): ; $(info $(M) building terraform provider plugin…) @
	$Q go build -ldflags "${LDFLAGS}"

.PHONY: install
install: ; $(info $(M) installing terraform provider plugin…) @
	$Q go install -ldflags "${LDFLAGS}"

.PHONY: clean
clean: ; $(info $(M) cleanup…) @
	$Q rm -f $(BIN)
