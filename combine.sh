protoc \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 -I. --include_imports --include_source_info \
  --descriptor_set_out=proto.pb ./service/models/service.proto ./hello-service/models/service.proto
