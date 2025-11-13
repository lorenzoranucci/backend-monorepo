# syntax=docker/dockerfile:1

ARG BASE_IMAGE=golang:1.25.0

FROM $BASE_IMAGE AS build-stage

ARG DOMAIN
ARG APP

WORKDIR /srv

COPY go.work /srv/go.work
COPY shared/go-lib /srv/shared/go-lib
COPY $DOMAIN/shared/go-lib /srv/$DOMAIN/shared/go-lib
COPY $DOMAIN/$APP /srv/$DOMAIN/$APP

WORKDIR /srv/$DOMAIN/$APP

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -trimpath -o bin/appbin
RUN CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -trimpath -o bin/appbin-arm64

FROM gcr.io/distroless/static-debian11:nonroot

ARG DOMAIN
ARG APP

WORKDIR /srv

COPY --from=build-stage /srv/${DOMAIN}/${APP}/bin/appbin ./appbin
COPY --from=build-stage /srv/${DOMAIN}/${APP}/bin/appbin-arm64 ./appbin-arm64

ENTRYPOINT ["/srv/appbin"]
