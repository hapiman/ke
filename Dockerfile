FROM fuyufjh/go-alpine
WORKDIR $GOPATH/src/hapiman/ke
COPY . $GOPATH/src/hapiman/ke

EXPOSE 6600
CMD ["./ke"]
