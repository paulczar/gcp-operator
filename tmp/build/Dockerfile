FROM alpine:3.6

RUN adduser -D gcp-operator
USER gcp-operator

ADD tmp/_output/bin/gcp-operator /usr/local/bin/gcp-operator
