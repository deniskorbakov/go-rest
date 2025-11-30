FROM golang:1.24.4 as Modules

ADD go.mod go.sum /m/

RUN cd /m && go mod download

FROM golang:1.24.4 as Builder

COPY --from=Modules /go/pkg /go/pkg

RUN useradd -u 10001 go-http-template

RUN mkdir -p /go-http-template
ADD . /go-http-template
WORKDIR /go-http-template

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o ./bin/go-http-template ./cmd/go-http-template

FROM scratch

COPY --from=Builder /etc/passwd /etc/passwd
USER go-http-template

COPY --from=Builder /go-http-template/bin/go-http-template /go-http-template

CMD ["/go-http-template"]


