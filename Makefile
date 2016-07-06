PWD := `pwd`
build:
	docker run --rm -v $(PWD):/go/src/github.com/ddelnano/racquetball -w /go/src/github.com/ddelnano/racquetball golang:1.6.2-alpine go build -v
