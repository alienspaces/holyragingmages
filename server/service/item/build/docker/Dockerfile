# build image
FROM golang:1.13.1-alpine3.10 as builder

# enable https:// repositories
RUN apk update \
    && apk add curl \
    && apk add bash \
    && apk add git

# app home
ENV APP_HOME /holyragingmages

# copy application
WORKDIR $APP_HOME
COPY . .

# build API server
WORKDIR $APP_HOME/service/item
RUN go build -o /go/bin/hrm-item-server ./cmd/server
RUN go build -o /go/bin/hrm-item-cli ./cmd/cli

# final image
FROM golang:1.13.1-alpine3.10

RUN apk update \
    && apk add curl \
    && apk add bash

COPY --from=builder /go/bin/hrm-item-server /go/bin
COPY --from=builder /go/bin/hrm-item-cli /go/bin
COPY --from=builder /holyragingmages/service/item/build/docker/entrypoint.sh .

RUN pwd
RUN ls -la

# entrypoint
ENTRYPOINT [ "./entrypoint.sh" ]
