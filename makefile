all: go.mod
	@go generate
	@go build

go.mod:
	@git fetch
	@git submodule init
	@git submodule update

	@go mod init
	@go mod tidy

	@go get github.com/mjibson/esc
	@go get github.com/gdamore/tcell
	@go get github.com/yuin/gopher-lua
	@go get gopkg.in/yaml.v2
