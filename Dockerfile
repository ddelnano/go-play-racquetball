FROM golang:1.7

WORKDIR /src
COPY racquetball racquetball

ENTRYPOINT ["./racquetball"]
