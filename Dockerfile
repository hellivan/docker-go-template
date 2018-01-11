FROM alpine:3.7

COPY ./go-template /usr/bin/go-template

ENTRYPOINT ["/usr/bin/go-template"]
