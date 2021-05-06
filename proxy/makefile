build:
	GOOS=linux GOARCH=amd64 go build -o proxy_linux
	docker build -t zy84338719/proxy:latest -t zy84338719/proxy:0.1 .
start:
	docker-compose up -d
stop:
	docker-compose down
dev:
	docker-compose up