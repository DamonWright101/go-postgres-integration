FROM golang:alpine AS build

WORKDIR /go/src/github.com/DamonWright101

RUN apk add --update alpine-sdk 

RUN go get github.com/DamonWright101/go-postgres-integration && \
    cd go-postgres-integration && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build
#     go build -ldflags '-w -extldflags "-static"'

RUN ls -lsa

# FROM scratch
# COPY --from=build /go/src/github.com/DamonWright101/go-postgres-integration /go-postgres-integration

FROM golang:alpine

COPY --from=build /go/src/github.com/DamonWright101/go-postgres-integration /go-postgres-integration

ENTRYPOINT ["/go-postgres-integration"]