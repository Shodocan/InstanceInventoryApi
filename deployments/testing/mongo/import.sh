#!/bin/sh

mongoimport --db $MONGO_INITDB_DATABASE --collection instances --file /data/data.json --jsonArray