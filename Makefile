include .env

.PHONY: run-database
run-database:
	docker run \
	-d \
	--name db_echoServer \
	-p 3306:3306 \
	-e MYSQL_ROOT_PASSWORD=${DB_ROOT_PASSWORD} \
	-e MYSQL_DATABASE=${DB_NAME} \
	-e MYSQL_USER=${DB_USER} \
	-e MYSQL_PASSWORD=${DB_PASSWORD} \
	mysql:latest