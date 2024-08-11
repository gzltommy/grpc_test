#all: proto docker-build docker-run
all: proto build-server build-client

.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	    ./07.load_balance/proto/*.proto

#proto:
#	protoc -I=$(SearchPath) \
#       --go_out=$(SearchPath) --go_opt=paths=source_relative \
#       --go-grpc_out=$(SearchPath) --go-grpc_opt=paths=source_relative \
#       --grpc-gateway_out=$(SearchPath) --grpc-gateway_opt=paths=source_relative \
#       helloworld/hello_world.proto helloworld/common.proto

# 这里有个坑指定了 -I 参数后，不能使用类似 ./04.trace/proto/hello/*.proto 这样带通配符 * 的方式.
# 解决方式：到搜索路径下去执行 protoc 命令，这样就可以不使用 -I 参数。

.PHONY: build-server
build-server:
	docker build -t server:latest \
				 -f Dockerfile.server \
				 --network=host \
				 .

.PHONY: build-client
build-client:
	docker build -t client:latest \
				 -f Dockerfile.client \
				 --network=host \
				 .

.PHONY: run-server1
run-server1:
	docker run --name server1 --rm server

.PHONY: run-server2
run-server2:
	docker run --name server2 --rm server

.PHONY: run-client
run-client:
	docker run --name client --rm client

#.PHONY: build
#build:
#	go build -o SERVER ./cmd/api
#
#.PHONY: run
#run:
#	./SERVER start -c ./settings/gzl2.json
#
#.PHONY: cyc
#cyc:
#	go-cyclic run --dir .
#
#.PHONY: data
#data:
#	docker-compose up -d