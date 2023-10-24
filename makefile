.PHONY: all disasm

all: go.mod disasm
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

disasm: /usr/local/bin/wla-gb
	@cd oracles-disasm && $(MAKE) seasons

/usr/local/bin/wla-gb:
	@echo "need sudo to install wla-dx"
	sudo curl https://wiki.zeldahacking.net/files/oracles-disasm/wla-gb -o /usr/local/bin/wla-gb
	sudo curl https://wiki.zeldahacking.net/files/oracles-disasm/wlalink -o /usr/local/bin/wlalink
	sudo chmod +x /usr/local/bin/wla*
