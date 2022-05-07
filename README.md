# fimproto
Contains all shared proto files and generated code

To build the proto files:
```
protoc -I . -I %USERPROFILE%\go\src\github.com\envoyproxy\protoc-gen-validate --go_out=":./proto" --go-grpc_out=":./proto" --validate_out="lang=go:./proto" --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./proto/fim.proto
```

