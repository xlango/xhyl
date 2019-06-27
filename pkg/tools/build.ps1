$GOOS="linux"

go build -o ../deploy/main ./
cp config.json ../deploy/