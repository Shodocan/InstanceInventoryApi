#!/bin/sh

BUILD=$1

if [ -z $BUILD ]
then
    echo "must specify a version release.sh VERSION"
fi

if [ -z "/bin/app" ]
then
    echo "aborting"
    exit 1
fi

cp ./bin/app ./deployments
docker build -t walissoncasonatto/instance-inventory-api:${BUILD} deployments

rm-rf ./deployments/app

docker build -t walissoncasonatto/instance-inventory-db:${BUILD} deployments/local/mongo

if [ -z ${DOCKER_HUB_LOGIN} ]
then
    echo "Docker hub credentials not set skipping publish"
    exit 0
fi

docker login -u ${DOCKER_HUB_LOGIN} -p ${DOCKER_HUB_PASS}

docker push walissoncasonatto/instance-inventory-db:${BUILD}
docker push walissoncasonatto/instance-inventory-api:${BUILD}