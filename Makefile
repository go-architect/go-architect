.PHONY: wails-install wails-check wails-build install-macos wails-build-windows build-linux build-windows

wails-build: wails-install
	CGO_ENABLED=1 wails build -clean

wails-build-windows: wails-install
	CGO_ENABLED=1 wails build -platform=windows -clean

wails-check: wails-install
	wails doctor

wails-install:
	go install github.com/wailsapp/wails/v2/cmd/wails@latest

install-macos: wails-check wails-build
	@cp -rf build/bin/Go\ Architect.app /Applications

build-linux: wails-check wails-build
	@cp build/bin/go-architect .

build-windows: wails-check wails-build-windows
	@cp build/bin/go-architect.exe .

ifeq ($(OS),Windows_NT)
    current_os := Windows
else
    current_os := $(shell uname -s)
endif

install:
	@echo "Detected OS: $(current_os)"
ifeq ($(current_os),Darwin)
	@make install-macos
	@echo "Your Go-Architect executable was installed in '/Applications' folder"
else ifeq ($(current_os),Linux)
	@make build-linux
	@echo "Your Go-Architect executable was generated in current folder"
else ifeq ($(current_os),Windows)
	@make build-windows
	@echo "Your Go-Architect executable was generated in current folder"
else
	@echo "Cannot install, Unknown OS: $(current_os)"
endif