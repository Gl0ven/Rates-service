dbname =
dbuser =
dbpassword =

docker-build:
	docker build -t rates:latest .

run:
	if [ "$(dbname)" != "" ]; then \
		sed -i 's/^DB_NAME=.*/DB_NAME=$(dbname)/' .env; \
	fi
	if [ "$(dbuser)" != "" ]; then \
		sed -i 's/^DB_USER=.*/DB_USER=$(dbuser)/' .env; \
	fi
	if [ "$(dbpassword)" != "" ]; then \
		sed -i 's/^DB_PASSWORD=.*/DB_PASSWORD=$(dbpassword)/' .env; \
	fi

	docker compose up -d --build

test:
	go test -v -cover ./...
