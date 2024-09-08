build: 
	rm -rf bin
	@echo "building api"
	mkdir bin
	go build -o bin/api cmd/main.go
	@echo "build successfully done"

run: build
	@echo "starting api"
	./bin/api