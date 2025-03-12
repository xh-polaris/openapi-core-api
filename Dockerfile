FROM golang:1.22-alpine AS builder

LABEL stage=gobuilder

WORKDIR /build

RUN apk update --no-cache && apk add --no-cache tzdata

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN sh ./build.sh

FROM alpine

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai

ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /build/output /app

CMD ["sh", "./bootstrap.sh"]
