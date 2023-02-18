sh gen_proto.sh

go install .
protoc -I=./proto --go_out=../ --my-option_out=../ --proto_path=test_proto test_option.proto