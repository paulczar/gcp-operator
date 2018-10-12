FROM alpine:3.6

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN adduser -D gcp-operator
USER gcp-operator

ADD tmp/_output/bin/gcp-operator /usr/local/bin/gcp-operator
