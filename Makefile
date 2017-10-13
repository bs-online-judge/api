all:
	mkdir -p ./bin
	go build -o ./bin/api
	./bin/api
