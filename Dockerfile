FROM golang:1.21-alpine as builder
WORKDIR /gopass-api

# Install dependencies required to git clone.
RUN apk update && \
    apk add --update git && \
    apk add --update make && \
    apk add --update openssh


WORKDIR /gopass-api
RUN git clone https://github.com/gopasspw/gopass.git
WORKDIR /gopass-api/gopass
RUN make build

WORKDIR /gopass-external-secrets
COPY go.sum ./
COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal

RUN go mod download
RUN go build -o ./gopass-external-secrets cmd/gopass-external-secrets/main.go


FROM vladgh/gpg
COPY --from=builder /gopass-api/gopass/gopass/                          /bin/gopass
COPY --from=builder /gopass-external-secrets/gopass-external-secrets    /rundir/gopass-external-secrets

RUN apk update && \
    apk add --update git && \
    apk add --update openssh

WORKDIR /rundir

COPY entrypoint.sh ./entrypoint.sh

RUN chmod 777 entrypoint.sh

ENTRYPOINT [ "sh", "entrypoint.sh" ]
