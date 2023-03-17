FROM node:14.18 as web-builder
RUN yarn config set registry https://registry.npm.taobao.org/
WORKDIR /apps
COPY . .
WORKDIR /apps/src
RUN make ui


FROM golang:1.18 AS builder
RUN ln -s /usr/bin/python3 /usr/bin/python
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
COPY --from=web-builder /apps/ /apps/
WORKDIR /apps/src
RUN make server
RUN make enterprise

FROM centos:7
WORKDIR /apps
COPY --from=builder /apps/src/bin/enterprise/cmdb /apps/cmdb