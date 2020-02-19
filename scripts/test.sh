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

# build new mongo
docker build -t testing-mongo deployments/local/mongo

docker run -p 27017:27017 -d --rm -e MONGO_INITDB_DATABASE="Instance" --name inventory-mongo testing-mongo

#point tests to mongo

export DB_DATABASE=Instance
export DB_PORT=27017
export DB_HOST=$DOCKER_HOST
export LOG_LEVEL=Info

go get -t -v ./...
go test ./... -timeout=2m -parallel=4 -covermode=atomic -coverprofile coverage.out

docker stop inventory-mongo