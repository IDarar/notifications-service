run:
	go run ./cmd/hub-notifications/main.go
genproto:
	sudo docker-compose up -d
cert:
	cd cert; bash gen.sh; cd ..

.PHONY: gen clean cert
