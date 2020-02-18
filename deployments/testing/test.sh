docker build -t testing-mongo mongo

docker stack deploy -c docker-compose.yml test