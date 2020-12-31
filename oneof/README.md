# Protobuf Oneofs

Examples for Protobuf `Oneof` fields.

## Generate

```bash
# From the current folder (if using my docker protoc image.)
$ docker run --rm -v $(pwd)/proto:/proto -v $GOPATH/src:/go-out  krishnaiyer/protoc-go  oneof.proto
```
