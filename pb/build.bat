protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative ./hello/hello_grpc.proto
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative ./person/person.proto
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative ./home/home.proto
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative ./product/product.proto
protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  --grpc-gateway_out .  --grpc-gateway_opt paths=source_relative   ./gateway/hello_world.proto