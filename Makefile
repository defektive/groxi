
build: cmd/
	goreleaser release --snapshot --clean

upx:: build
	upx --best dist/groxi*/groxi*

dist:: clean build upx

release::
	goreleaser release

clean: dist/
	rm -rf dist/*