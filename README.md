Proto packages for Presidio API
===========================

[![Build Status](https://travis-ci.org/google/go-genproto.svg?branch=master)](https://travis-ci.org/presid-io/presidio-genproto.svg?branch=master)

This repository contains the generated GO and Python packages for Presidio API, and the generated [gRPC](https://grpc.io) code necessary for interacting with Presidio's gRPC services
APIs.

### Generate Docs using [protoc-gen-doc](https://github.com/pseudomuto/protoc-gen-doc) ###

```
docker run --rm \
  -v $(pwd)/doc:/out \
  -v $(pwd)/src:/protos \
  pseudomuto/protoc-gen-doc --doc_opt=markdown,proto.md
```
