FROM golang:1.21.3-alpine AS builder

RUN set -xe && \
    apk update && apk upgrade && \
    apk add --no-cache make git curl gcc g++ && \
    apk add --no-cache git ca-certificates

WORKDIR /app
COPY go.mod go.sum ./

ENV GO111MODULE=on
ENV CGO_ENABLED=1
RUN go mod download

COPY . .
WORKDIR /app
RUN make test && make build && \
    cp build/nftique /nftique

FROM alpine:3.9
COPY --from=builder /nftique /nftique

CMD ["/nftique"]

EXPOSE 8080