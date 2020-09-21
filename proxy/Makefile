BINARY_NAME=proxy
IMAGE_NAME=hkccr.ccs.tencentyun.com/tcp-proxy/tcp-proxy:latest


GOFLAG=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
GOBUILD=${GOFLAG} go build
GOCLEAN=go clean
GOTEST=go test

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	./$(BINARY_NAME)

docker-build:
	docker build -t $(IMAGE_NAME) .

docker-push:
	docker push $(IMAGE_NAME)
