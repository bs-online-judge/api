all:
	@mkdir -p ./bin
	@go build -o ./bin/api
	env PG_ADDR=localhost:5432 PG_USER=postgres PG_PASSWORD= API_PORT=8000 ./bin/api
