FROM golang:alpine AS build

WORKDIR /go/src/github.com/DamonWright101

RUN apk add --update alpine-sdk 

RUN go get github.com/DamonWright101/go-postgres-integration && \
    cd go-postgres-integration/app && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags '-X github.com/DamonWright101/go-postgres-integration -w -extldflags "-static"'

FROM scratch
COPY --from=build /go/src/github.com/DamonWright101/go-postgres-integration /go-postgres-integration

ENTRYPOINT ["/go-postgres-integration"]