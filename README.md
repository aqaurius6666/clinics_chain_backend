1. Run server
Option 1: Go to internal and type 'go run .'
Option 2: Go to internal and type 'go build -o ./out/go-example' --> then type './out/go-example'

2. Migrate struct from proto
- Go to proto, type 'protoc --go-grpc_out=. *.proto --go_out=. *.proto'