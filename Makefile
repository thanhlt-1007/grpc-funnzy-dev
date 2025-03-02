compile:
	rm -rf pbs;
	mkdir -p pbs;
	protoc --proto_path=protos protos/*.proto \
		--go_out=pbs --go_opt=paths=source_relative \
		--go-grpc_out=pbs --go-grpc_opt=paths=source_relative;
