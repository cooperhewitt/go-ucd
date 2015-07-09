build:
	go run ./bin/ucd-build-unicodedata.go > ./src/org.cooperhewitt/ucd/data/unicodedata/unicodedata.go
	go run ./bin/ucd-build-unihan.go > ./src/org.cooperhewitt/ucd/data/unihan/unihan.go

compile:
	go build -o ./bin/ucd ./bin/ucd.go
	go build -o ./bin/ucd-server ./bin/ucd-server.go