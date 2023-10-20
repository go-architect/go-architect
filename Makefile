.PHONY: wails-build wails-check install-macos

wails-build: wails-install
	wails build -clean

wails-check: wails-install
	wails doctor

wails-install:
	go install github.com/wailsapp/wails/v2/cmd/wails@latest

install-macos: wails-check wails-build
	cp -rf build/bin/Go\ Architect.app /Applications

build-linux: wails-check wails-build
	cp build/bin/Go\ Architect .

build-windows: wails-check wails-build
	cp build/bin/Go\ Architect.exe .

