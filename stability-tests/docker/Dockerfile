ARG DOGESILVER_IMAGE
ARG DOGESILVERMINER_IMAGE

FROM ${DOGESILVER_IMAGE} as dogesilverd
FROM ${DOGESILVERMINER_IMAGE} as dogesilverminer

FROM golang:1.23-alpine

RUN mkdir -p /go/src/github.com/dogesilvernet/dogesilverd

WORKDIR /go/src/github.com/dogesilvernet/dogesilverd

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=dogesilverd /app/ /app/
COPY --from=dogesilverminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/dogesilvernet/dogesilverd/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]