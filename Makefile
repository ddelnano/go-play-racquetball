PWD := `pwd`
build:
	docker run --rm -v $(PWD):/go/src/github.com/ddelnano/racquetball -w /go/src/github.com/ddelnano/racquetball golang:1.7 go build -v
