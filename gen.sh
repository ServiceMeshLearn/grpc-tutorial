protoc -I proto/ \
    -I${PROTO_PATH} \
    --go_out=plugins=grpc:proto \
    proto/api.proto
