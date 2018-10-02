.PHONY: bundle build clean copy compile-app compress

clean:
	rm -rf build bundle.zip
	mkdir build

copy:
	cp opdl build/onepiece-dl

build: 
	go build

compile-app:
	osacompile -l JavaScript -o build/onepiece-dl.app opdl.js

compress:
	zip -r bundle.zip build

bundle: clean build compile-app copy compress