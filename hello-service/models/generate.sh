protoc -I/usr/local/include -I. \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --js_out=import_style=commonjs,binary:. --go_out=plugins=grpc:. service.proto
