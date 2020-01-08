# Presidio Protobuf installation
This document describes how to install the specific protobuf versions used in Presidio

### Python
Install specific grpcio-tools and proto versions:

``` sh
pip install grpcio-tools==1.11.1
pip install googleapis-common-protos==1.5.0
```

 
### Go
1. Navigate to `/Users/<user>/go/src/github.com/golang/protobuf`

2. Run

``` sh
git checkout v1.2.0

curl -OL https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-osx-x86_64.zip
unzip protoc-3.6.1-osx-x86_64.zip    -d {$GOPATH}

```
