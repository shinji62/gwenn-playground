all: test compile
compile: linux32 linux64 darwin64

test:
	ginkgo -r -v .

linux32:
	 CGO_ENABLED=0 GOARCH=386 GOOS=linux go build -o dist/linux/386/gwenn-playground
linux64:
	 CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o dist/linux/amd64/gwenn-playground_linux_amd64
darwin64:
	GOARCH=amd64 GOOS=darwin go build -o dist/darwin/amd64/gwenn-playground_darwin_amd64

clean:
	-rm -rf dist/*
	-rm -rf *.prof
