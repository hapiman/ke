CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -v -a -installsuffix cgo -o ke .
docker build -t ke .
docker-compose -f docker-compose.yaml  up -d
