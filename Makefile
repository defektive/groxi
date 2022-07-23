
build: cmd/
	goreleaser release --snapshot --rm-dist

upx:: build
	upx --best dist/groxi*/groxi*

dist:: clean build upx

clean: dist/
	rm -rf dist/*