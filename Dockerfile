FROM golang:alpine AS build-env

WORKDIR /app
ADD . /app
RUN go build -o gh-ingestor ./cmd/

FROM alpine

RUN apk update && \
    apk add ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=build-env /app/gh-ingestor /app

ENV GIN_MODE=release
ENV PORT=3000

EXPOSE 3000
ENTRYPOINT ["./gh-ingestor"]