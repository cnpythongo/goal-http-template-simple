ARG GOAL_APP
FROM golang:1.19.13-alpine3.18 AS build-env
ARG GOAL_APP
ADD . /app
WORKDIR /app
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct
ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0
RUN go mod download \
    && time go build -o goal-${GOAL_APP} cmd/${GOAL_APP}/main.go


FROM alpine
ARG GOAL_APP
WORKDIR /app
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && apk del tzdata \
    && mkdir settings runtime
COPY --from=build-env /app/goal-${GOAL_APP} /app/goal
CMD ["echo", "hello goal app"]
