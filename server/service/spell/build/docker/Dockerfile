# build image
FROM golang:1.15.7-alpine3.12 as builder

# enable https:// repositories
RUN apk update \
    && apk add curl \
    && apk add bash \
    && apk add git

# app home
ENV APP_SERVER_HOME /holyragingmages

# copy application
WORKDIR $APP_SERVER_HOME
COPY . .

# build API server
WORKDIR $APP_SERVER_HOME/service/spell
RUN go build -o /go/bin/hrm-spell-server ./cmd/server
RUN go build -o /go/bin/hrm-spell-cli ./cmd/cli

# final image
FROM golang:1.15.7-alpine3.12

RUN apk update \
    && apk add curl \
    && apk add bash

COPY --from=builder /go/bin/hrm-spell-server /go/bin
COPY --from=builder /go/bin/hrm-spell-cli /go/bin
COPY --from=builder /holyragingmages/service/spell/build/docker/entrypoint.sh .
COPY --from=builder /holyragingmages/schema/docs ./schema/docs

RUN pwd
RUN ls -la

# entrypoint
ENTRYPOINT [ "./entrypoint.sh" ]
