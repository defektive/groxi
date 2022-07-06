GOOSES = linux darwin windows
GOARCHS = amd64
VARIANTS = cmd/*
MOD = recon
OS = linux
ARCH = amd64
NAME = groxi
build: cmd/
	rm -rf dist; \
	for var in $(VARIANTS); do \
		for os in $(GOOSES); do \
			for arch in $(GOARCHS); do \
			  cd $$var;\
				echo $$var $$os $$arch; \
				CGO_ENABLED=0 GOOS=$$os GOARCH=$$arch go build -ldflags="-s -w" -o ../../dist/$$os/$$arch/$$(echo $$var|sed 's/variants/bin/'); \
				cd - ; \
			done \
		done \
	done

upx: cmd/
	for var in $(VARIANTS); do \
		for os in $(GOOSES); do \
			for arch in $(GOARCHS); do \
				echo $$var $$os $$arch; \
				mkdir -p dist/$$os/$$arch/upx; \
				upx --best -o$$(echo dist/$$os/$$arch/$$var|sed 's/variants/upx/') dist/$$os/$$arch/$$(echo $$var|sed 's/variants/bin/'); \
			done \
		done \
	done \

dist:: build upx
