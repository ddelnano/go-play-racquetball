FROM golang:1.6.2-alpine

WORKDIR /src

COPY reservation.json reservation.json
COPY racquetball racquetball

ENTRYPOINT ["./racquetball"]
