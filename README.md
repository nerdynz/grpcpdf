# jeevesrpcpdf



### Generate protoc
```protoc -I makepdf makepdf/makepdf.proto --go_out=plugins=grpc:makepdf ```



### Quick notes
-DOCKER

__Build:__ `docker build ./ -t jeevesgrpc`

__Run:__ `docker run -p 127.0.0.1:5532:3456/tcp -e GRPC_PORT=5532 jeevesgrpc:latest`