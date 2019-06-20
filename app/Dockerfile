FROM golang:alpine
WORKDIR /go/src/github.com/

COPY postgres-int.go postgres-int.go 

RUN go get && \
    go build -o postgres-int .

CMD /app/postgres-int