Proto packages for Presidio API
===========================

[![Build Status](https://dev.azure.com/csedevil/Presidio/_apis/build/status/Presidio-genproto?branchName=master)](https://dev.azure.com/csedevil/Presidio/_build/latest?definitionId=57?branchName=master)

This repository contains the generated GO and Python packages for [Presidio]
(https://github.com/microsoft/presidio) API, and the generated [gRPC](https://grpc.io) code necessary for interacting with Presidio's gRPC services
APIs.


# Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.microsoft.com.

When you submit a pull request, a CLA-bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., label, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.


# Changing Presidio's API
First, set up a Presidio development environment. For more info, see Presidio's development document (https://github.com/microsoft/presidio/edit/master/docs/development.md)

Presidio leverages [protobuf](https://github.com/golang/protobuf) to create API classes and services across multiple environments.

Follow these steps to change Presidio's API:
1. Fork the [presidio-genproto](https://github.com/Microsoft/presidio-genproto) repo into `YOUR_ORG/presidio-genproto`
2. Clone the repo into the `$GOPATH/src/github.com/YOUR_ORG/presidio-genproto` folder
3. Make the desired changes to the .proto files in /src
4. Make sure you have [protobuf](https://github.com/golang/protobuf) installed
5. Generate the Go and Python files. Run the following commands in the `src` folder of `presidio-genproto`:

    ```sh
    python -m grpc_tools.protoc -I . --python_out=../python --grpc_python_out=../python ./*.proto

    protoc -I . --go_out=plugins=grpc:../golang ./*.proto
    ```
    
 5. Copy all the files in the `python` folder into `presidio-analyzer/analyzer`. All generated files end with `*pb2.py` or `*pb2_grpc.py`
 6. Change the constraint on `Gopkg.toml` which directs to the location of `presidio-genproto`
From:

```yaml
[[constraint]]
  branch = "master"
  name = "github.com/Microsoft/presidio-genproto"
```

To:

```yaml
[[constraint]]
  branch = "YOUR_GENPROTO_BRANCH"
  name = "github.com/YOUR_ORG/presidio-genproto"

```
  7. Update `Gopkg.lock` by calling `dep ensure` or `dep ensure --update github.com/YOUR_ORG/presidio-genproto`
  8. Push all the changes (generated python files, `Gopkg.toml` and `Gopkg.lock` into your presidio repo

For more info, see https://grpc.io/docs/tutorials/basic/python.html

