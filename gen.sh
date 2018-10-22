protoc -I proto/ \
    -I${PROTO_PATH} \
    --go_out=plugins=grpc:proto \
    --grpc-gateway_out=:proto \
    proto/api.proto
