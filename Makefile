all:
	@mkdir -p ./bin
	@go build -o ./bin/api
	env PG_HOSTNAME=localhost:5432 PG_USERNAME=postgres PG_PASSWORD= API_PORT=8000 \
	REDIS_HOSTNAME=localhost:6379 REDIS_PASSWORD= ./bin/api
