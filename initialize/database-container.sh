# create docker network
# docker network create authservice
# docker network create apiservice

#After that, create database schema for dbs
docker run --name apiservice-db -p 3306:3306 --volume "authservice-db":/tmp -e MYSQL_ROOT_PASSWORD=1 -d mysql
docker run --name authservice-db -p 3307:3306 --volume "authservice-db":/tmp -e MYSQL_ROOT_PASSWORD=1 -d mysql

# Go to db console:  "docker exec -it [container name] bash -l"