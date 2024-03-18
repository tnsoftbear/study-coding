generate:
	@mkdir -p ./internal/generated
	@swagger generate server -f ./api/spec.yml -t ./internal/generated --exclude-main
build:
	@go build cmd/collector.go
run:
	sudo ./collector --tls-certificate /mnt/e/dev/certs/auction.lv.crt --tls-key /mnt/e/dev/certs/auction.lv.key
runwin:
	./collector.exe --tls-certificate e:/dev/certs/auction.lv.crt --tls-key e:/dev/certs/auction.lv.key