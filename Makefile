GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

VERSION=15.0.0
UNICODE_DATA=UnicodeData.txt
UNIHAN_DATA=Unihan.zip

data:
	go run cmd/ucd-build-unicodedata/main.go -data https://www.unicode.org/Public/$(VERSION)/ucd/$(UNICODE_DATA) > unicodedata/unicodedata.go
	go run cmd/ucd-build-unihan/main.go -data https://www.unicode.org/Public/$(VERSION)/ucd/$(UNIHAN_DATA) > unihan/unihan.go

tools:
	@make cli

cli:
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/ucd cmd/ucd/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/ucd-dump cmd/ucd-dump/main.go
	go build -mod $(GOMOD) -ldflags="-s -w" -o bin/ucd-server cmd/ucd-server/main.go
