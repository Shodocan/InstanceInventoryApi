#/bin/sh

docker network create -d bridge inventory
docker run -d --rm -e MONGO_INITDB_DATABASE="Instance" --name inventory-mongo --network=inventory walissoncasonatto/instance-inventory-db
docker run --rm  -e DB_HOST="inventory-mongo" --name inventory-tests --network=inventory walissoncasonatto/instance-inventory-api