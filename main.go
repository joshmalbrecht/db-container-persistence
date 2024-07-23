package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v4/pgxpool"
)

const DB_NAME = "books"
const DB_HOSTNAME = "localhost"
const DB_USER = "admin"
const DB_PASSWORD = "admin"
const DB_PORT = "5432"
const MIGRATION_DIRECTORY = "file://db"

const RECORDS_TO_CREATE = 10

const CHARSET = "abcdefghijklmnopqrstuvwxyz"

func main() {

	println("attempting to establish database connection")

	databaseUrl := "postgres://" + DB_USER + ":" + DB_PASSWORD + "@" + DB_HOSTNAME + ":" + DB_PORT + "/" + DB_NAME + "?sslmode=disable"
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		panic("unable to connect to database: " + err.Error())
	}

	defer dbPool.Close()

	config, err := pgx.ParseConnectionString(dbPool.Config().ConnString())
	if err != nil {
		panic("unable to parse connection string: " + err.Error())
	}

	db := stdlib.OpenDB(config)
	driverInstance, err := postgres.WithInstance(db, new(postgres.Config))
	if err != nil {
		panic("unable to get driver instance: " + err.Error())
	}

	println("database connection established")

	migrateInstance, err := migrate.NewWithDatabaseInstance(MIGRATION_DIRECTORY, DB_NAME, driverInstance)
	if err != nil {
		panic("unable to get database migration instance: " + err.Error())
	}

	err = migrateInstance.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic("failure to run database migration: " + err.Error())
	}

	println("database migration completed")

	for i := 0; i < RECORDS_TO_CREATE; i++ {
		name := StringWithCharset(10)

		_, err := db.Exec(`INSERT INTO books (name, author, genre) VALUES (($1), ($2), ($3))`, StringWithCharset(10), StringWithCharset(10), StringWithCharset(10))
		if err != nil {
			panic("failure to insert into table")
		}

		println("created book with name: " + name)
	}
}

func StringWithCharset(length int) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = CHARSET[seededRand.Intn(len(CHARSET))]
	}
	return string(b)
}
