VERSION=13.0.0
UNICODE_DATA=UnicodeData-13.0.0d6.txt
UNIHAN_DATA=Unihan-13.0.0d5.zip

fmt:
	go fmt *.go
	go fmt unicodedata/*.go
	go fmt unihan/*.go
	go fmt cmd/*.go

data:
	go run cmd/ucd-build-unicodedata/main.go -data https://www.unicode.org/Public/$(VERSION)/ucd/$(UNICODE_DATA) > unicodedata/unicodedata.go
	go run cmd/ucd-build-unihan/main.go -data https://www.unicode.org/Public/$(VERSION)/ucd/$(UNIHAN_DATA) > unihan/unihan.go

tools:	
	go build -o bin/ucd cmd/ucd/main.go
	go build -o bin/ucd-dump cmd/ucd-dump/main.go
	go build -o bin/ucd-server cmd/ucd-server/main.go
