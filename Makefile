run:
	go run ./cmd/hub-notifications/main.go
mongo:
	sudo docker-compose up -d
genproto:
	protoc -I . notifications.proto --go_out=plugins=grpc:. 
 
cert:
	cd cert; bash gen.sh; cd ..

.PHONY: gen clean cert
