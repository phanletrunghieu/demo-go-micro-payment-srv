gen-protoc:
	protoc --micro_out=. --go_out=. proto/payment/payment.proto

build: gen-protoc
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/payment-srv ./cmd/main.go

dockerize:
	docker build -t payment-srv .

run:
	docker run --rm -p 50052:50052 \
		-e MICRO_SERVER_ADDRESS=:50052 \
		payment-srv