build: 
	rm -rf bin
	@echo "building api"
	mkdir bin
	go build -o bin/api cmd/main.go
	@echo "build successfully done"

run: build
	@echo "starting api"
	./bin/api

docker_up: docker_down
	sudo systemctl start docker
	sudo docker-compose up -d

docker_down:
	sudo docker-compose down