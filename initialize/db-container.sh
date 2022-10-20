# create docker network
# docker network create authservice
# docker network create apiservice

#After that, create database schema for dbs
docker run --name apiservice-db -p 5432:5432 --volume "authservice-db":/tmp -e POSTGRES_PASSWORD=1 -d postgres
docker run --name authservice-db -p 5433:5432 --volume "authservice-db":/tmp -e POSTGRES_PASSWORD=1 -d postgres

# Go to db console:  "docker exec -it [container name] bash -l"