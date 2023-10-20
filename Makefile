hello:
	echo "Hello, World"

wails-build:
	wails build -clean

wails-check:
	wails doctor

install: wails-check wails-build
	cp -rf build/bin/Go\ Architect.app /Applications
