#! /bin/sh

cd src
python -m grpc_tools.protoc -I . --python_out=../python --grpc_python_out=../python ./*.proto
protoc -I . --go_out=plugins=grpc:../golang ./*.proto
cd ..