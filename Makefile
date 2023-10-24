.PHONY: wails-build wails-check install-macos

wails-build: wails-install
	wails build -clean

wails-check: wails-install
	wails doctor

wails-install:
	go install github.com/wailsapp/wails/v2/cmd/wails@latest

install-macos: wails-check wails-build
	@cp -rf build/bin/Go\ Architect.app /Applications

build-linux: wails-check wails-build
	@cp build/bin/Go\ Architect .

build-windows: wails-check wails-build
	@cp build/bin/Go\ Architect.exe .

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
