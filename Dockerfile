FROM golang:1.6

WORKDIR /src
COPY racquetball racquetball

ENTRYPOINT ["./racquetball"]
