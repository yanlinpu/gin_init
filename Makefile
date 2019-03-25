# SHELL := /bin/bash
.PHONY: build clean dev test

# 全局环境变量设置
export GINENV=development
export GINURL=http://localhost:5555/api/v1/
all: clean

depInstall:
	brew install dep
depInit:
	dep init

# make depAdd src="github.com/xxxx/xxxxx@=0.4.3"
# ifeq 顶格写 ifeq ("", "") else endif
depAdd:
ifeq ("$(src)", "")
	@echo "src can not be null, like github.com/xxxx/xxxxx@=0.4.3"
else
	dep ensure -add $(src)
endif

proInit:dep
	dep ensure -v

build:
	swag init
	go build -v .

swag:
	swag init

test: export GINENV=testing
test:
	go test -v -coverprofile=cover.out ./...
	go tool cover -func=cover.out
	# 执行可视化查看 go tool cover -html=cover.out
	go tool cover -html=cover.out

# 不同target下设置不同环境变量 必须单独一行
# dev: export GINENV=development
dev: swag
	go run main.go

clean:
	rm -rf main
	go clean -i .

# npm install -g swagger-markdow
swag-md:
	swagger-markdown -i docs/swagger/swagger.yaml

# api.1
ApiTest:
	curl -X GET "$(GINURL)get-test?page=1&pageSize=20" \
	-H "accept: application/json" \
	-H "Content-Type: application/json" \

ApiPost:
	curl -X POST "$(GINURL)post-test" \
	-H "accept: application/json" \
	-H "Content-Type: application/json" \
	-d '{ \
		"page":1, \
		"pageSize":20 \
	}'
