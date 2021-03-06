# Custom Protoc

A docker image for Protobuf with Golang.

The base image contains
- [Protoc](https://github.com/protocolbuffers/protobuf)
- [Protoc-gen-go](https://github.com/golang/protobuf)

This image can be extended to include other plugins.
- `withvalidator`: contains [envoyproxy/protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate).

Heavily inspired by [znly/docker-protobuf](https://github.com/znly/docker-protobuf/blob/master/Dockerfile)
