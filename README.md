# Rates service

The Grpc rates service returns the lowest "ask" price and the highest "bid" 
with a unix timestamp at the rate of "USDT-RUB".

## Targets

- `make docker-build` - build a docker image with the app

- `make run` - run the app in docker compose

    | Option | Description |
    |:------------|-------------|
    | **`dbname=...`** | set your own db name |
    | **`dbuser=...`** | set your own db user |
    | **`dbpassword=...`** | set your own db password |
    
- `make test` - local run unit-tests