package main

import (
	"database/sql"
	"log"

	"github.com/adictya/medicist-backend/api"
	db "github.com/adictya/medicist-backend/db/sqlc"

	_ "github.com/lib/pq"
)

const(
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/root?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main(){

	conn, err := sql.Open(dbDriver,dbSource)
	if err != nil {
		log.Fatal("Connection to database failed:",err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil{
		log.Fatal("Cannot start server", err)
	}
}

