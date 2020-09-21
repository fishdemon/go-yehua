# 这是一个 echo demo 的构建文件
BINARY_NAME=echo-demo

GO_FLAG=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
GO_BUILD=${GO_FLAG} go build
GO_CLEAN=go clean
GO_TEST=go test

all: begin build end

begin:
	@ echo starting compile
build:
	${GO_BUILD} -o ${BINARY_NAME} main.go
end:
	@ echo complete compile...