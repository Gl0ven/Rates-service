# Rates service

The Grpc rates service returns the lowest "ask" price and the highest "bid" 
with a unix timestamp at the rate of "USDT-RUB".

## Targets

- make docker-build - build a docker image with the app
- make run - run the app in docker compose
    flags:
    dbname - change db name
    dbuser - change db user
    dbpassword - change db password
- make test - local run unit-tests