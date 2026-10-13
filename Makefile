GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w
CWD=$(shell pwd)

vulnup:
	go install golang.org/x/vuln/cmd/govulncheck@latest

vuln:
	govulncheck -show verbose ./...

VERSION=17.0.0
UNICODE_DATA=UnicodeData.txt
UNIHAN_DATA=Unihan.zip

data:
	@make data-unicode
	@make data-unihan

data-unicode:
	go run cmd/ucd-build-unicodedata/main.go -data https://www.unicode.org/Public/$(VERSION)/ucd/$(UNICODE_DATA) > unicodedata/unicodedata.go

data-unihan:
	go run cmd/ucd-build-unihan/main.go -data https://www.unicode.org/Public/$(VERSION)/ucd/$(UNIHAN_DATA) > unihan/unihan.go

tools:
	@make cli

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/ucd cmd/ucd/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/ucd-dump cmd/ucd-dump/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/ucd-server cmd/ucd-server/main.go
