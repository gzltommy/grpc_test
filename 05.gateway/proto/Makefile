# Makefile
.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
       helloworld/*.proto

# 这里有个坑指定了 -I 参数后，不能使用类似 ./04.trace/proto/hello/*.proto 这样带通配符 * 的方式.
# 解决方式：到搜索路径下去执行 protoc 命令，这样就可以不使用 -I 参数。
