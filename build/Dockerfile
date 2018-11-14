FROM alpine:3.8

RUN apk upgrade; apk add ca-certificates --no-cache; apk update --no-cache

USER nobody

ADD build/_output/bin/vault-secret-operator /usr/local/bin/vault-secret-operator
