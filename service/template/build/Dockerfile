# build image
FROM golang:1.13.1-alpine3.10 as builder

# enable https:// repositories
RUN apk update \
    && apk add curl \
    && apk add bash \
    && apk add git

# app home
ENV APP_HOME /go/service

# copy application
WORKDIR $APP_HOME
COPY . .

# build API server
RUN go build ./cmd/template -o /go/bin/service

# final image
FROM golang:1.13.1-alpine3.10

RUN apk update \
    && apk add curl \
    && apk add bash

COPY --from=builder /go/bin/service /go/bin

# entrypoint
ENTRYPOINT [ "./entrypoint.sh" ]
