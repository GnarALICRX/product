

.PHONY: proto
proto:
	E:\gopractice/gomicro/git.imooc.com/cap1573/product/proto/product
    docker run --rm -v /e/gopractice/gomicro/git.imooc.com/cap1573/product:/e/gopractice/gomicro/git.imooc.com/cap1573/product -w /e/gopractice/gomicro/git.imooc.com/cap1573/product -e ICODE=2D169C17E77AE695 cap1573/cap-protoc -I /e/gopractice/gomicro/git.imooc.com/cap1573/product --go_out=/e/gopractice/gomicro/git.imooc.com/cap1573/product /e/gopractice/gomicro/git.imooc.com/cap1573/product/proto/product/product.proto --micro_out=/e/gopractice/gomicro/git.imooc.com/cap1573/product

.PHONY: build
build: 

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o product-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t product-service:latest
