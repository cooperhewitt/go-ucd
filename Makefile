data:
	go run -mod vendor cmd/ucd-build-unicodedata/main.go -data https://www.unicode.org/Public/UCD/latest/ucd/UnicodeData.txt > unicodedata/unicodedata.go
	go run -mod vendor cmd/ucd-build-unihan/main.go -data https://www.unicode.org/Public/UCD/latest/ucd/Unihan.zip > unihan/unihan.go

tools:
	@make cli

cli:
	go build -mod vendor -o bin/ucd cmd/ucd/main.go
	go build -mod vendor -o bin/ucd-dump cmd/ucd-dump/main.go
	go build -mod vendor -o bin/ucd-server cmd/ucd-server/main.go
