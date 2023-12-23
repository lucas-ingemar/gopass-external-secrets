FROM golang:1.21-alpine as builder
WORKDIR /gopass-api

# RUN go install github.com/gopasspw/gopass-jsonapi@latest
# RUN go install github.com/gopasspw/gopass@latest

# Install dependencies required to git clone.
RUN apk update && \
    apk add --update git && \
    apk add --update make && \
    apk add --update openssh


RUN git clone https://github.com/gopasspw/gopass-jsonapi.git
WORKDIR /gopass-api/gopass-jsonapi
RUN make build


WORKDIR /gopass-api
RUN git clone https://github.com/gopasspw/gopass.git
WORKDIR /gopass-api/gopass
RUN make build

# COPY go.sum ./
# COPY go.mod ./

# RUN go mod download

# COPY cmd ./cmd
# COPY i18n ./i18n
# COPY internal ./internal
# COPY migrations ./migrations

# RUN go build -o /go_app cmd/main.go

FROM vladgh/gpg
COPY --from=builder /gopass-api/gopass-jsonapi/gopass-jsonapi /bin/gopass-jsonapi
COPY --from=builder /gopass-api/gopass/gopass/                /bin/gopass
# COPY --from=amacneil/dbmate:1.13 /usr/local/bin/dbmate /usr/local/bin/dbmate
#

RUN apk update && \
    apk add --update git && \
    apk add --update openssh
# apk add --update systemctl && \

WORKDIR /rundir

# COPY i18n ./i18n
# COPY migrations ./migrations
COPY entrypoint.sh ./entrypoint.sh

RUN chmod 777 entrypoint.sh

# EXPOSE 9876

ENTRYPOINT [ "sh", "entrypoint.sh" ]
# RUN gpg

# CMD [ "ls" ]
