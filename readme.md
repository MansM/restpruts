# Restpruts
this is a hobby thing that I am using to learn a bit about go

docker-compose file should be in ${GOPATH}

## database

Just for fun I am trying out cockroachdb, I am using the insecure flag, 
so dont be stupid and never use this anywhere near production!


### database setup:

docker exec -it golang_db_1 ./cockroach user set gorm
docker exec -it golang_db_1 ./cockroach sql -e 'CREATE DATABASE events'
docker exec -it golang_db_1 ./cockroach sql -e 'GRANT ALL ON DATABASE events TO gorm'

docker exec -it golang_db_1 ./cockroach sql -d events -e "INSERT INTO events (name) VALUES ('meetup')"

dont forget to set pass