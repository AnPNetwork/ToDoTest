package main

import (
	"flag"
	"log"
	"test/controller"
	"test/domain"
	"test/repository"
)

func main() {
	var (
		settings domain.Settings
	)

	// CMD args

	// HTTP
	flag.StringVar(&settings.Http.Host, "http_host", "0.0.0.0", "host http server")
	httpPort := flag.Int("http_port", 3000, "port http server")

	// DB
	flag.StringVar(&settings.DB.Host, "db_host", "localhost", "host db server")
	dbPort := flag.Int("db_port", 5432, "port db server")
	flag.StringVar(&settings.DB.Name, "db_name", "site", "db name PostgreSQL")
	flag.StringVar(&settings.DB.User, "db_user", "site", "user PostgreSQL")
	flag.StringVar(&settings.DB.Password, "db_pass", "site", "password PostgreSQL")
	flag.StringVar(&settings.DB.SSL, "db_ssl", "disable", "ssl PostgreSQL")

	flag.Parse()

	if httpPort != nil && *httpPort > 0 && *httpPort <= 65535 {
		settings.Http.Port = uint16(*httpPort)
	}

	if dbPort != nil && *dbPort > 0 && *dbPort <= 65535 {
		settings.DB.Port = uint16(*dbPort)
	}

	db, err := repository.NewConnectionDB(settings.DB)
	if err != nil {
		log.Print(err)
		return
	}

	err = db.CreateTable()
	if err != nil {
		if err != nil {
			log.Print(err)
			return
		}
	}

	httpServer := controller.NewServerHttp(settings.Http, db)

	if err := httpServer.Start(); err != nil {
		log.Print(err)
	}
}
