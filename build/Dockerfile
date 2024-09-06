# Build
# golang:1.19-buster
FROM golang:1.19-alpine AS build
ARG GOPROXY=https://goproxy.cn,direct
WORKDIR /app

RUN apk add \
  bash \
  git \
  make

COPY . .

RUN make prepare compile

RUN ls -R /app

# Deploy

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/output/ /app/

EXPOSE 3636
EXPOSE 3635
EXPOSE 3634

CMD ["/app/bin/jt808-server-go"]
