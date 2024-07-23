# Database Container With Persistent Data

## Start the Postgres database container

Run the following command in the root directory to start the Postgres container.

`docker-compose up -d`

## Run main.go

Run main.go using `go run main.go` and this will run a database migration on the books table and populate the table with some random entries.

## Verify the data is in the table

Connect to the database and verify that the `books` table has the random data.

`psql -h localhost -p 5432 -d books -U admin`

## Destroy the database container

`docker-compose down`

## Verify that the container is destroyed

`sudo docker ps -a`

## Verify that the volume is created

`sudo docker volume ls`

You can also inspect the volume using 

`sudo docker volume inspect <VOLUME NAME>`

## Create a new database container

`docker-compose up -d`

## Verify that the data still exists

`psql -h localhost -p 5432 -d books -U admin`
