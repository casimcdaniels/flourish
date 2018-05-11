package main

import (
	"os"
	"log"
	"github.com/alexsasharegan/dotenv"
	"github.com/jmoiron/sqlx"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"net/http"
	"github.com/casimcdaniels/flourish"
)

/*************************************************************


___________.__                      .__       .__
\_   _____/|  |   ____  __ _________|__| _____|  |__
 |    __)  |  |  /  _ \|  |  \_  __ \  |/  ___/  |  \
 |     \   |  |_(  <_> )  |  /|  | \/  |\___ \|   Y  \
 \___  /   |____/\____/|____/ |__|  |__/____  >___|  /
	 \/                                     \/     \/

					 M
					dM
					MMr
				   4MMML
				   MMMMM                 xf
				  MMMMMM                MM-
	Mh..          +MMMMMM             MMMM
	 MMM           MMMMML           MMMMMh
	  3MMMMx      'MMMMMMf      xnMMMMMMM
	  '*MMMMM      MMMMMM      nMMMMMMPM
		*MMMMMx    MMMMMM\     MMMMMMM=
		   MMMMMM   3MMMM   dMMMMMM
			MMMMMM  MMMMM   MMMMMM       MMMMM
=            *MMMMx  MMMM  dMMMMM     nnMMMMM*
  MMMn...     'MMMMr 'MM   MMMM    nMMMMMMM*M
   M4MMMMnn..   *MMM  MM  MMPM   dMMMMMMM""
	 ^MMMMMMMMx   *ML MM  M*   MMMMMM**M
		*PMMMMMMhn  *x > M   MMMM**""
		   ""**MMMMhx/ h/  =*M
					 3PM%....
				  nPM     M*MMnx



**********************************************************/


func main () {
	rootPath, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	cfg, err := loadCfg(rootPath)

	if err != nil {
		log.Fatal(err)
	}

	db, err := setupDB(cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbSchema)

	if err != nil {
		log.Fatal(err)
	}

	strainRepo := flourish.MysqlStrainRepository{DB: db}
	strainService := flourish.StrainService{Strains: strainRepo}
	router := mux.NewRouter()

	router.HandleFunc("/strains", flourish.CreateStrainEndpoint(strainService)).Methods("POST")
	router.HandleFunc("/strains", flourish.UpdateStrainEndpoint(strainService)).Methods("PATCH")
	router.HandleFunc("/strains", flourish.DeleteStrainEndpoint(strainService)).Methods("DELETE")
	router.HandleFunc("/strains/search", flourish.SearchStrainsEndpoint(strainService)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8888", router))
}

// Application Config

type config struct {
	DbHost string
	DbPort string
	DbUser string
	DbPassword string
	DbSchema string
}

// Loads the config from .env file
func loadCfg (path string) (config, error){
	err := dotenv.Load(path + "/.env")

	if err != nil {
		return config{}, err
	}


	return config{
		DbHost: os.Getenv("FLOURISH_DB_HOST"),
		DbPort: os.Getenv("FLOURISH_DB_PORT"),
		DbUser: os.Getenv("FLOURISH_DB_USER"),
		DbPassword: os.Getenv("FLOURISH_DB_PASSWORD"),
		DbSchema: os.Getenv("FLOURISH_DB_SCHEMA"),
	}, nil
}

// Initializes a sql connection pool using MySQL driver
func setupDB(host, port, user, password, schema string) (*sqlx.DB, error) {
	url := fmt.Sprintf("%s:%s@(%s:%s)/%s", user, password, host, port, schema)
	return sqlx.Connect("mysql", url)
}