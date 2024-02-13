#!make
lint:
	gofmt -w -s .
	golangci-lint run services/*
	golangci-lint run main.go

install:
	go get ./...
	go mod tidy

protogen:
	@for proto in services/**/protobuf/*.proto ; do \
		echo "Processing $$proto..."; \
		protoc -I/usr/local/include -I. \
		  -I. \
		  -I./services \
		  -I${GOPATH}/src \
		  -I/Users/triviere/lab/googleapis \
		  --go_out=paths=source_relative:. \
		  --go-grpc_out=paths=source_relative:. \
		  --grpc-gateway_out=logtostderr=true,paths=source_relative:. \
		  $$proto; \
		if [ $$? -ne 0 ]; then \
			echo "Error processing $$proto with go, go-grpc, or grpc-gateway"; \
			exit 1; \
		fi; \
		protoc -I/usr/local/include -I. \
		  -I. \
		  -I./services \
		  -I${GOPATH}/src \
		  -I/Users/triviere/lab/googleapis \
		  --openapiv2_out=logtostderr=true:. \
		  $$proto; \
		if [ $$? -ne 0 ]; then \
			echo "Error processing $$proto with openapiv2"; \
			exit 1; \
		fi; \
	done

clean:
	rm services/**/protobuf/*.pb.go || true
	rm services/**/protobuf/*.pb.gw.go || true
	rm services/**/protobuf/*.swagger.json || true