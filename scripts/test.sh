#!/bin/sh


export GO111MODULE=on

# stop old mongo instances

MONGO_DOCKER=$(docker ps | grep inventory-mongo)
if [  -z "$MONGO_DOCKER" ]
then
    echo "No Mongo Running"
else
    docker stop inventory-mongo
fi

docker network create -d bridge inventory
# build new mongo
docker build -t inventory-mongo deployments/local/mongo

docker run -d --rm -e MONGO_INITDB_DATABASE="Instance" --name inventory-mongo --network=inventory inventory-mongo

#point tests to mongo

docker build -t inventory-tests deployments/local/tests

docker run -e DB_HOST="inventory-mongo" --name inventory-tests --network=inventory inventory-tests

EXIT=$(docker inspect inventory-tests --format='{{.State.ExitCode}}')

docker cp inventory-tests:/go/src/github.com/Shodocan/InstanceInventoryApi/coverage.out .

docker rm inventory-tests

docker stop inventory-mongo

exit $EXIT