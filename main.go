package main

import (
	"chat-app/auth-service/api"
	db "chat-app/auth-service/db/sqlc"
	chat "chat-app/chat-service"
	"chat-app/util"
	"database/sql"
	"flag"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	server := flag.String("server", "", "http,websocket")
	flag.Parse()
	fmt.Println("Distributed Chat App v0.01")

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := sql.Open(config.DBDRIVER, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	auth := db.NewAuth(conn)

	if *server == "auth" {
		runGinServer(config, auth)
	} else if *server == "chat" {
		runChatServer(config)
	} else {
		log.Fatal().Msg("invalid server type")
	}
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Printf("Db migrated successfully")
}

func runGinServer(config util.Config, auth db.Auth) {

	runDBMigration(config.MigrationURL, config.DBSource)

	server, err := api.NewServer(config, auth)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}

func runChatServer(config util.Config) {
	log.Info().Msg("Starting chat server on :8082")
	chat.ConnectRedis(&config)
	chat.StartWebsocketServer()
}
