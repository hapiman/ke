FROM fuyufjh/go-alpine
WORKDIR $GOPATH/src/hapiman/ke
COPY . $GOPATH/src/hapiman/ke
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add -U tzdata
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
  && echo 'Asia/Shanghai' >/etc/timezone
EXPOSE 6600
CMD ["./ke"]
