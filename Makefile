PROTO_DIR := api/proto
GEN_DIR := api/gen/go/product

.PHONY: gen clean run

gen:
	protoc --proto_path=$(PROTO_DIR) \
		--go_out=$(GEN_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/*.proto

clean:
	rm -f $(GEN_DIR)/*.pb.go

run:
	go run cmd/server/main.go